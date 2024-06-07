package compiler

import (
	"io"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

type OutputType string

const (
	OutputTypeLSD      OutputType = "protobuf"
	OutputTypeCompiled OutputType = "compiled"
)

type Config struct {
	InputDir     string                 `toml:"input_dir"`
	OutputDir    *string                `toml:"output_dir"`
	OutputType   *OutputType            `toml:"output_type"`
	Minimize     *bool                  `toml:"minimize"`
	OrderPersist *bool                  `toml:"order_persist"`
	Protobuf     *ConfigProtobufSection `toml:"protobuf"`
}

type ConfigProtobufSection struct {
	Version *string           `toml:"version"`
	AutoDL  *bool             `toml:"auto_dl"`
	Package *string           `toml:"package"`
	Targets []string          `toml:"targets"`
	Options map[string]string `toml:"options"`
}

func LoadConfig(dir string) (*Config, error) {
	file, err := os.Open(path.Join(dir, "protolsd.toml"))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	tomlText, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(tomlText, &config); err != nil {
		return nil, err
	}

	if config.OutputDir == nil {
		config.OutputDir = ptr("dist")
	}

	if config.OutputType == nil {
		config.OutputType = ptr(OutputTypeLSD)
	}

	if config.Minimize == nil {
		config.Minimize = ptr(false)
	}

	if config.OrderPersist == nil {
		config.OrderPersist = ptr(true)
	}

	if config.Protobuf == nil {
		config.Protobuf = &ConfigProtobufSection{}
	}

	if config.Protobuf.Version == nil {
		config.Protobuf.Version = ptr("latest")
	}

	if config.Protobuf.AutoDL == nil {
		config.Protobuf.AutoDL = ptr(true)
	}

	if config.Protobuf.Package == nil {
		config.Protobuf.Package = ptr("main")
	}

	if config.Protobuf.Targets == nil {
		config.Protobuf.Targets = []string{}
	}

	if config.Protobuf.Options == nil {
		config.Protobuf.Options = map[string]string{}
	}

	return &config, nil
}

func GenerateMinimalConfig(dir string) error {
	config := Config{
		InputDir:   dir,
		OutputDir:  ptr("dist"),
		OutputType: ptr(OutputTypeLSD),
		Minimize:   ptr(true),
		Protobuf: &ConfigProtobufSection{
			Version: ptr("latest"),
			AutoDL:  ptr(true),
			Package: ptr("main"),
			Targets: []string{},
			Options: map[string]string{},
		},
	}

	tomlText, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(dir, "protolsd.toml"), tomlText, 0644)
}

func ptr[T any](v T) *T {
	return &v
}
