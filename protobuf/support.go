package protobuf

type Target string

const (
	TargetGo     Target = "go"
	TargetPython Target = "python"
	TargetJava   Target = "java"
	TargetCSharp Target = "csharp"
	TargetCpp    Target = "cpp"
	TargetDart   Target = "dart"
	TargetKotlin Target = "kotlin"
)

var (
	supportedTargets = []string{"go", "python", "java", "csharp", "cpp", "dart", "kotlin"}
)

func IsLanguageTargetSupported(target string) bool {
	for _, t := range supportedTargets {
		if t == target {
			return true
		}
	}

	return false
}

func getAdditionalArgs(target Target, withGrpc bool, outPath string) []string {
	args := []string{}

	if target == TargetGo {
		args = append(args, "--go_opt=paths=source_relative")

		if withGrpc {
			args = append(args, "--go-grpc_opt=paths=source_relative")
			args = append(args, "--go-grpc_out="+outPath)
		}
	}

	return args
}
