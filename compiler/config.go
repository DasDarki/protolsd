package compiler

import (
	"fmt"
	"io"
	"os"
	"path"
	"protolsd/protobuf"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml/v2"
)

type OutputType string

const (
	OutputTypeLSD      OutputType = "protobuf"
	OutputTypeCompiled OutputType = "compiled"
)

type Config struct {
	EnvFile      *string                `toml:"env_file"`
	InputDir     string                 `toml:"input_dir"`
	OutputDir    *string                `toml:"output_dir"`
	OutputType   *OutputType            `toml:"output_type"`
	OrderPersist *bool                  `toml:"order_persist"`
	WithGRPC     *bool                  `toml:"with_grpc"`
	Protobuf     *ConfigProtobufSection `toml:"protobuf"`
}

type ConfigProtobufSection struct {
	Version       *string                      `toml:"version"`
	AutoDL        *bool                        `toml:"auto_dl"`
	Package       *string                      `toml:"package"`
	Targets       []string                     `toml:"targets"`
	TargetOptions map[string]map[string]string `toml:"target"`
	Options       map[string]string            `toml:"options"`
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

	if config.OrderPersist == nil {
		config.OrderPersist = ptr(true)
	}

	if config.WithGRPC == nil {
		config.WithGRPC = ptr(false)
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

	if config.Protobuf.Targets != nil {
		for i, target := range config.Protobuf.Targets {
			config.Protobuf.Targets[i] = strings.ToLower(target)

			if !protobuf.IsLanguageTargetSupported(config.Protobuf.Targets[i]) {
				return nil, fmt.Errorf("unsupported target: %s", config.Protobuf.Targets[i])
			}
		}
	}

	if config.Protobuf.TargetOptions == nil {
		config.Protobuf.TargetOptions = map[string]map[string]string{}
	}

	if config.Protobuf.Options == nil {
		config.Protobuf.Options = map[string]string{}
	}

	if config.EnvFile != nil && *config.EnvFile != "" {
		err := godotenv.Load(path.Join(dir, *config.EnvFile))
		if err != nil {
			return nil, err
		}
	}

	resolveConfigEnvVars(&config)

	return &config, nil
}

func GenerateMinimalConfig(dir string) error {
	config := Config{
		InputDir:   "src",
		OutputDir:  ptr("dist"),
		OutputType: ptr(OutputTypeLSD),
		Protobuf: &ConfigProtobufSection{
			Version:       ptr("latest"),
			AutoDL:        ptr(true),
			Package:       ptr("main"),
			Targets:       []string{},
			TargetOptions: map[string]map[string]string{},
			Options:       map[string]string{},
		},
	}

	tomlText, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path.Join(dir, "src"), 0755); err != nil {
		return err
	}

	return os.WriteFile(path.Join(dir, "protolsd.toml"), tomlText, 0644)
}

func resolveConfigEnvVars(config *Config) {
	config.InputDir = *resolveConfigEnvVar(&config.InputDir)
	config.OutputDir = resolveConfigEnvVar(config.OutputDir)

	if config.OutputType != nil {
		config.OutputType = ptr(OutputType(*resolveConfigEnvVar(ptr(string(*config.OutputType)))))
	}

	if config.Protobuf != nil {
		config.Protobuf.Version = resolveConfigEnvVar(config.Protobuf.Version)
		config.Protobuf.Package = resolveConfigEnvVar(config.Protobuf.Package)

		if config.Protobuf.Targets != nil {
			for i, target := range config.Protobuf.Targets {
				config.Protobuf.Targets[i] = *resolveConfigEnvVar(ptr(target))
			}
		}

		if config.Protobuf.TargetOptions != nil {
			for target, opts := range config.Protobuf.TargetOptions {
				for key, val := range opts {
					config.Protobuf.TargetOptions[target][key] = *resolveConfigEnvVar(ptr(val))
				}
			}
		}

		if config.Protobuf.Options != nil {
			for key, val := range config.Protobuf.Options {
				config.Protobuf.Options[key] = *resolveConfigEnvVar(ptr(val))
			}
		}
	}
}

func resolveConfigEnvVar(val *string) *string {
	if val != nil {
		if strings.HasPrefix(*val, "$") {
			return ptr(os.Getenv((*val)[1:]))
		}

		return val
	}

	return nil
}

func ptr[T any](v T) *T {
	return &v
}

func unptr[T any](v *T) T {
	return *v
}
