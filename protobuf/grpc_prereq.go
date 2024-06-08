package protobuf

import (
	"fmt"
	"os/exec"
)

// Check if go and grpc plugins are installed
func CheckGrpcPrerequisities() error {
	if err := checkGoCLI(); err != nil {
		return err
	}

	// install grpc plugins
	if err := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@v1.28").Run(); err != nil {
		return fmt.Errorf("failed to install grpc plugin: %v", err)
	}

	if err := exec.Command("go", "install", "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2").Run(); err != nil {
		return fmt.Errorf("failed to install grpc plugin: %v", err)
	}

	return nil
}

func checkGoCLI() error {
	if err := exec.Command("go", "version").Run(); err != nil {
		return fmt.Errorf("go not found: %v", err)
	}

	return nil
}
