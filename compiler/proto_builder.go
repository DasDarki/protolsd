package compiler

import (
	"fmt"
	"sort"
	"strings"
)

type ProtoBuilder struct {
	lines       []string
	indentLevel int
}

type ProtoField struct {
	Type     string
	Number   int
	Modifier string
}

func newProtoBuilder() *ProtoBuilder {
	return &ProtoBuilder{
		lines:       []string{},
		indentLevel: 0,
	}
}

func (pb *ProtoBuilder) indent() {
	pb.indentLevel++
}

func (pb *ProtoBuilder) dedent() {
	if pb.indentLevel > 0 {
		pb.indentLevel--
	}
}

func (pb *ProtoBuilder) addLine(line string) {
	indent := strings.Repeat("  ", pb.indentLevel)
	pb.lines = append(pb.lines, indent+line)
}

func (pb *ProtoBuilder) AddSyntax(syntax string) {
	pb.addLine(fmt.Sprintf(`syntax = "%s";`, syntax))
	pb.addLine("")
}

func (pb *ProtoBuilder) AddPackage(packageName string) {
	pb.addLine(fmt.Sprintf("package %s;", packageName))
	pb.addLine("")
}

func (pb *ProtoBuilder) AddOptions(opts map[string]string) {
	for key, opt := range opts {
		pb.addLine(fmt.Sprintf("option %s = \"%s\";", key, opt))
	}
	pb.addLine("")
}

func (pb *ProtoBuilder) AddImport(importPath string) {
	pb.addLine(fmt.Sprintf(`import "%s";`, importPath))
}

func (pb *ProtoBuilder) AddMessage(name string, fields map[string]ProtoField, nestedCallback func(*ProtoBuilder)) {
	pb.addLine(fmt.Sprintf("Message %s {", name))
	pb.indent()

	if nestedCallback != nil {
		nestedCallback(pb)
	}

	for fieldName, field := range fields {
		pb.addLine(strings.TrimSpace(fmt.Sprintf("%s %s %s = %d;", field.Modifier, field.Type, fieldName, field.Number)))
	}

	pb.dedent()
	pb.addLine("}")
	pb.addLine("")
}

func (pb *ProtoBuilder) AddEnum(name string, values map[int]string) {
	pb.addLine(fmt.Sprintf("Enum %s {", name))
	pb.indent()

	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		pb.addLine(fmt.Sprintf("%s = %d;", values[k], k))
	}

	pb.dedent()
	pb.addLine("}")
	pb.addLine("")
}

func (pb *ProtoBuilder) AddService(name string, methods map[string][2]string) {
	pb.addLine(fmt.Sprintf("Service %s {", name))
	pb.indent()
	for methodName, params := range methods {
		requestType := params[0]
		responseType := params[1]
		pb.addLine(fmt.Sprintf("rpc %s (%s) returns (%s);", methodName, requestType, responseType))
	}
	pb.dedent()
	pb.addLine("}")
	pb.addLine("")
}

func (pb *ProtoBuilder) String() string {
	return strings.Join(pb.lines, "\n")
}
