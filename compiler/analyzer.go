package compiler

import (
	"fmt"
	"protolsd/parser"
	"strconv"
	"strings"
)

// analyzer is the struct that will be used to analyze the AST and convert it to the ProtoLSD model.
type analyzer struct {
	*parser.BaseProtoLSDListener
	inGlobalScope      bool
	script             *Script
	prevMessage        *Message
	currSerivce        *Service
	visitedMessageDefs map[*parser.MessageDefinitionContext]bool
}

func (a *analyzer) EnterPreprocessorDirective(ctx *parser.PreprocessorDirectiveContext) {
	if !a.inGlobalScope {
		return
	}

	name := ctx.PREPROCESSOR_NAME().GetText()
	params := resolvePreprocessorParameters(ctx.PreprocessorParameters())

	if name == "private-import" {
		if len(params) != 1 {
			panic("invalid number of parameters for #private-import")
		}

		if containsStr(params, "global", true) {
			a.script.AreImportsPrivate = true
		} else if containsStr(params, "unset", true) {
			a.script.AreImportsPrivate = false
		} else {
			panic("invalid parameter for #private-import; expected 'global' or 'unset'")
		}
	} else if name == "alias" {
		if len(params) != 2 {
			panic("invalid number of parameters for #alias")
		}

		baseType := params[0]
		alias := params[1]

		if _, ok := a.script.TypeAliases[alias]; ok {
			panic(fmt.Sprintf("alias '%s' already defined", alias))
		}

		a.script.TypeAliases[alias] = baseType
	} else if name == "decl-type" {
		if len(params) != 1 {
			panic("invalid number of parameters for #decl-type")
		}

		typeDecl := params[0]

		if _, ok := a.script.DeclaredTypes[typeDecl]; ok {
			panic(fmt.Sprintf("type '%s' already declared", typeDecl))
		}

		a.script.DeclaredTypes[typeDecl] = true
	}
}

func (a *analyzer) EnterImportStatement(ctx *parser.ImportStatementContext) {
	if !a.inGlobalScope {
		return
	}

	path := ctx.STRING().GetText()
	path = path[1 : len(path)-1]
	isPrivate := a.script.AreImportsPrivate

	if ctx.PRIVATE() != nil || ctx.BANG() != nil {
		isPrivate = true
	}

	importStmt := &ImportStatement{
		Path:    path,
		Private: isPrivate,
	}

	a.script.Imports = append(a.script.Imports, importStmt)
}

func (a *analyzer) EnterEnumDefinition(ctx *parser.EnumDefinitionContext) {
	name := ctx.IDENTIFIER().GetText()
	isEnumClass := true

	if _, ok := a.script.Enums[name]; ok {
		panic(fmt.Sprintf("Enum '%s' already defined", name))
	}

	for _, preproc := range ctx.AllAnnotationDirective() {
		if strings.EqualFold(preproc.IDENTIFIER().GetText(), "NoEnumClass") {
			isEnumClass = false
			break
		}
	}

	enum := &Enum{
		Name:      name,
		IsClass:   isEnumClass,
		Constants: map[string]int{},
	}

	a.script.Enums[name] = enum
	counter := 0

	for _, field := range ctx.EnumBody().AllEnumField() {
		fieldName := field.IDENTIFIER().GetText()

		fieldValue := 0
		if field.NUMBER() != nil {
			value := field.NUMBER().GetText()
			intVal, err := strconv.Atoi(value)
			if err != nil {
				panic(fmt.Sprintf("invalid number '%s' for Enum field '%s'", value, fieldName))
			}

			fieldValue = intVal
			counter = intVal + 1
		} else {
			fieldValue = counter
			counter++
		}

		if _, ok := enum.Constants[fieldName]; ok {
			panic(fmt.Sprintf("Enum field '%s' already defined in Enum '%s'", fieldName, name))
		}

		enum.Constants[fieldName] = fieldValue
	}
}

func (a *analyzer) EnterMessageDefinition(ctx *parser.MessageDefinitionContext) {
	if a.visitedMessageDefs[ctx] {
		return
	}

	a.visitedMessageDefs[ctx] = true

	parent := a.prevMessage

	name := ctx.IDENTIFIER().GetText()
	if parent != nil {
		if a.script.Messages[name] != nil {
			panic(fmt.Sprintf("Message '%s' already defined", name))
		}
	}

	message := &Message{
		Name:     name,
		Fields:   map[string]*messageField{},
		Children: map[string]*Message{},
	}

	if parent != nil {
		parent.Children[name] = message
	} else {
		a.script.Messages[name] = message
	}

	for _, stmt := range ctx.MessageBody().AllMessageBodyStatement() {
		if stmt.MessageField() != nil {
			field := stmt.MessageField()

			fieldName := field.IDENTIFIER().GetText()
			fieldType := resolveDataType(field.DataType())
			fieldOrder := -1
			overwrite := false

			var fieldModifier *string
			if field.OptionalModifier() != nil {
				fieldModifier = ptr(field.OptionalModifier().GetText())
			}

			if field.NUMBER() != nil {
				order, err := strconv.Atoi(field.NUMBER().GetText())
				if err != nil {
					panic(fmt.Sprintf("invalid number '%s' for Message field '%s'", field.NUMBER().GetText(), fieldName))
				}

				if order < 0 || (order >= 19000 && order <= 19999) {
					panic(fmt.Sprintf("invalid order '%d' for Message field '%s': must be in the range [0, 18999] or [20000, 536870911]", order, fieldName))
				}

				fieldOrder = order

				if field.BANG() != nil {
					overwrite = true
				}
			}

			if _, ok := message.Fields[fieldName]; ok {
				panic(fmt.Sprintf("field '%s' already defined in Message '%s'", fieldName, name))
			}

			message.Fields[fieldName] = &messageField{
				Name:      fieldName,
				DataType:  fieldType,
				Order:     fieldOrder,
				Modifier:  fieldModifier,
				Overwrite: overwrite,
			}
		} else if stmt.MessageDefinition() != nil {
			a.prevMessage = message

			a.EnterMessageDefinition(stmt.MessageDefinition().(*parser.MessageDefinitionContext))

			a.prevMessage = parent
		}
	}
}

func (a *analyzer) EnterServiceDefinition(ctx *parser.ServiceDefinitionContext) {
	name := ctx.IDENTIFIER().GetText()
	if a.script.Services[name] != nil {
		panic(fmt.Sprintf("Service '%s' already defined", name))
	}

	service := &Service{
		Name: name,
		Rpcs: map[string]*RpcMethod{},
	}

	a.script.Services[name] = service
	a.currSerivce = service

	for _, rpc := range ctx.ServiceBody().AllRpcDefinition() {
		a.EnterRpcDefinition(rpc.(*parser.RpcDefinitionContext))
	}
}

func (a *analyzer) EnterRpcDefinition(ctx *parser.RpcDefinitionContext) {
	if a.currSerivce == nil {
		panic("rpc definition outside of Service definition")
	}

	name := ctx.IDENTIFIER().GetText()
	if a.currSerivce.Rpcs[name] != nil {
		return
	}

	var input *RpcMethodMessage
	if ctx.RpcInput() != nil {
		input = resolveRpcMethodMessage(name+"Request", ctx.RpcInput().RpcMethodMessage())
	}

	var returns *RpcMethodMessage
	if ctx.RpcReturn() != nil {
		returns = resolveRpcMethodMessage(name+"Response", ctx.RpcReturn().RpcMethodMessage())
	}

	if input == nil {
		a.script.UsedEmptyRequest = true
	}

	if returns == nil {
		a.script.UsedEmptyResponse = true
	}

	rpc := &RpcMethod{
		Name:    name,
		Returns: returns,
		Input:   input,
	}

	a.currSerivce.Rpcs[name] = rpc
}

func resolveRpcMethodMessage(name string, ctx parser.IRpcMethodMessageContext) *RpcMethodMessage {
	if ctx.OBJECT() != nil {
		return nil
	}

	if ctx.RpcSingleMessage() != nil && ctx.RpcSingleMessage().RpcMessageValue() != nil {
		return &RpcMethodMessage{
			NeedsAutoWire: true,
			Name:          ptr(ctx.RpcSingleMessage().RpcMessageValue().GetText()),
		}
	}

	if ctx.RpcParameters() == nil {
		return nil
	}

	paramsLen := len(ctx.RpcParameters().AllRpcParameter())

	if paramsLen <= 0 {
		return nil
	}

	rm := &RpcMethodMessage{
		NeedsAutoWire: false,
		Name:          &name,
		Message: &Message{
			Name:     name,
			Children: map[string]*Message{},
			Fields:   map[string]*messageField{},
		},
	}

	counter := 1
	pidx := 1
	var lastDataType *DataType

	for _, param := range ctx.RpcParameters().AllRpcParameter() {
		if lastDataType != nil && (param.DataType() == nil || (param.PLUS() == nil && param.RpcParameterModifier() == nil)) {
			fieldOrder := counter
			counter++

			if fieldOrder >= 19000 && fieldOrder <= 19999 {
				counter += 1000
			}

			fieldName := ""
			if param.IDENTIFIER() != nil {
				fieldName = param.IDENTIFIER().GetText()
			} else if param.DataType() != nil {
				shallowDataType := resolveDataType(param.DataType())
				fieldName = shallowDataType.Text
			}

			rm.Message.Fields[fieldName] = &messageField{
				Name:     fieldName,
				DataType: lastDataType,
				Order:    fieldOrder,
				Modifier: nil,
			}
		} else if param.DataType() != nil {
			dataType := resolveDataType(param.DataType())
			if param.PLUS() != nil {
				lastDataType = dataType
			} else {
				lastDataType = nil
			}

			if param.RpcParameterModifier() == nil && paramsLen == 1 {
				return &RpcMethodMessage{
					NeedsAutoWire: true,
					Name:          ptr(dataType.Text),
				}
			}

			var pname *string
			repitions := 1

			if param.RpcParameterModifier() != nil {
				if param.RpcParameterModifier().NUMBER() != nil {
					i, err := strconv.Atoi(param.RpcParameterModifier().NUMBER().GetText())
					if err != nil {
						panic(fmt.Sprintf("invalid number '%s' for rpc parameter modifier", param.RpcParameterModifier().NUMBER().GetText()))
					}

					repitions = i
				} else if param.RpcParameterModifier().IDENTIFIER() != nil {
					pname = ptr(param.RpcParameterModifier().IDENTIFIER().GetText())
				}
			}

			for i := 0; i < repitions; i++ {
				fieldName := ""
				if pname != nil {
					fieldName = *pname

					if _, ok := rm.Message.Fields[fieldName]; ok {
						panic(fmt.Sprintf("field '%s' already defined in rpc Message '%s'", fieldName, *pname))
					}
				} else {
					fieldName = fmt.Sprintf("val%d", pidx)
					pidx++

					for {
						if _, ok := rm.Message.Fields[fieldName]; ok {
							fieldName = fmt.Sprintf("val%d", pidx)
							pidx++
						} else {
							break
						}
					}
				}

				fieldOrder := counter
				counter++

				if fieldOrder >= 19000 && fieldOrder <= 19999 {
					counter += 1000
				}

				rm.Message.Fields[fieldName] = &messageField{
					Name:     fieldName,
					DataType: dataType,
					Order:    fieldOrder,
					Modifier: nil,
				}
			}
		}
	}

	return rm
}

func resolveDataType(ctx parser.IDataTypeContext) *DataType {
	text := ctx.DataTypeText().GetText()
	isOptional := ctx.QUESTION() != nil
	isArray := ctx.ARRAY() != nil

	return &DataType{
		Text:       text,
		IsOptional: isOptional,
		IsArray:    isArray,
	}
}

func resolvePreprocessorParameters(ctx parser.IPreprocessorParametersContext) []string {
	params := []string{}

	for _, param := range ctx.AllRpcMessageValue() {
		params = append(params, param.GetText())
	}

	return params
}

func containsStr(slice []string, elem string, ignoreCase bool) bool {
	for _, e := range slice {
		if ignoreCase {
			if strings.EqualFold(e, elem) {
				return true
			}
		} else {
			if e == elem {
				return true
			}
		}
	}

	return false
}
