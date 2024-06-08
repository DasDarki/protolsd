package protobuf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

type Executor struct {
	target    Target
	inputDir  string
	outputDir string
	protocDir string
	grpc      bool
}

func NewExecutor(target Target, inputDir, outputDir, protocDir string, grpc bool) *Executor {
	return &Executor{
		target:    target,
		inputDir:  inputDir,
		outputDir: outputDir,
		protocDir: protocDir,
		grpc:      grpc,
	}
}

func (e *Executor) Execute() error {
	exePath, err := getExecutablePath(e.protocDir)
	if err != nil {
		return err
	}

	incPath := path.Join(e.protocDir, "include")
	if _, err := os.Stat(incPath); err != nil {
		return fmt.Errorf("include directory not found: %s -> %v", incPath, err)
	}

	outPath := e.outputDir
	if _, err := os.Stat(outPath); err == nil {
		if err := os.RemoveAll(outPath); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(outPath, 0755); err != nil {
		return err
	}

	args := []string{
		"--proto_path=" + e.inputDir,
		"--proto_path=" + incPath,
		"--" + string(e.target) + "_out=" + outPath,
	}

	args = append(args, getAdditionalArgs(e.target, e.grpc, outPath)...)

	args = append(args, path.Join(e.inputDir, "*.proto"))

	return runCommand(exePath, args)
}

func runCommand(exePath string, args []string) error {
	cmd := exec.Command(exePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pathEnv := os.Getenv("PATH")
	goBinPath := path.Join(os.Getenv("GOPATH"), "bin")
	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s%c%s", pathEnv, os.PathListSeparator, goBinPath))

	log.Printf("INFO: Running command %s %v", exePath, args)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func getExecutablePath(dir string) (string, error) {
	osn := getOS()
	if osn == "" {
		return "", fmt.Errorf("unsupported platform")
	}

	protoc := "protoc"
	if osn == "win" {
		protoc += ".exe"
	}

	return path.Join(dir, "bin", protoc), nil
}
