// Code generated from ProtoLSD.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ProtoLSD

import "github.com/antlr4-go/antlr/v4"

// BaseProtoLSDListener is a complete listener for a parse tree produced by ProtoLSDParser.
type BaseProtoLSDListener struct{}

var _ ProtoLSDListener = &BaseProtoLSDListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProtoLSDListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProtoLSDListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProtoLSDListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProtoLSDListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSemi is called when production semi is entered.
func (s *BaseProtoLSDListener) EnterSemi(ctx *SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *BaseProtoLSDListener) ExitSemi(ctx *SemiContext) {}

// EnterProtoLSD is called when production protoLSD is entered.
func (s *BaseProtoLSDListener) EnterProtoLSD(ctx *ProtoLSDContext) {}

// ExitProtoLSD is called when production protoLSD is exited.
func (s *BaseProtoLSDListener) ExitProtoLSD(ctx *ProtoLSDContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseProtoLSDListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseProtoLSDListener) ExitStatement(ctx *StatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseProtoLSDListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseProtoLSDListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterServiceDefinition is called when production serviceDefinition is entered.
func (s *BaseProtoLSDListener) EnterServiceDefinition(ctx *ServiceDefinitionContext) {}

// ExitServiceDefinition is called when production serviceDefinition is exited.
func (s *BaseProtoLSDListener) ExitServiceDefinition(ctx *ServiceDefinitionContext) {}

// EnterServiceBody is called when production serviceBody is entered.
func (s *BaseProtoLSDListener) EnterServiceBody(ctx *ServiceBodyContext) {}

// ExitServiceBody is called when production serviceBody is exited.
func (s *BaseProtoLSDListener) ExitServiceBody(ctx *ServiceBodyContext) {}

// EnterRpcDefinition is called when production rpcDefinition is entered.
func (s *BaseProtoLSDListener) EnterRpcDefinition(ctx *RpcDefinitionContext) {}

// ExitRpcDefinition is called when production rpcDefinition is exited.
func (s *BaseProtoLSDListener) ExitRpcDefinition(ctx *RpcDefinitionContext) {}

// EnterRpcParameters is called when production rpcParameters is entered.
func (s *BaseProtoLSDListener) EnterRpcParameters(ctx *RpcParametersContext) {}

// ExitRpcParameters is called when production rpcParameters is exited.
func (s *BaseProtoLSDListener) ExitRpcParameters(ctx *RpcParametersContext) {}

// EnterRpcParameter is called when production rpcParameter is entered.
func (s *BaseProtoLSDListener) EnterRpcParameter(ctx *RpcParameterContext) {}

// ExitRpcParameter is called when production rpcParameter is exited.
func (s *BaseProtoLSDListener) ExitRpcParameter(ctx *RpcParameterContext) {}

// EnterRpcParameterModifier is called when production rpcParameterModifier is entered.
func (s *BaseProtoLSDListener) EnterRpcParameterModifier(ctx *RpcParameterModifierContext) {}

// ExitRpcParameterModifier is called when production rpcParameterModifier is exited.
func (s *BaseProtoLSDListener) ExitRpcParameterModifier(ctx *RpcParameterModifierContext) {}

// EnterRpcReturn is called when production rpcReturn is entered.
func (s *BaseProtoLSDListener) EnterRpcReturn(ctx *RpcReturnContext) {}

// ExitRpcReturn is called when production rpcReturn is exited.
func (s *BaseProtoLSDListener) ExitRpcReturn(ctx *RpcReturnContext) {}

// EnterRpcInput is called when production rpcInput is entered.
func (s *BaseProtoLSDListener) EnterRpcInput(ctx *RpcInputContext) {}

// ExitRpcInput is called when production rpcInput is exited.
func (s *BaseProtoLSDListener) ExitRpcInput(ctx *RpcInputContext) {}

// EnterRpcMethodMessage is called when production rpcMethodMessage is entered.
func (s *BaseProtoLSDListener) EnterRpcMethodMessage(ctx *RpcMethodMessageContext) {}

// ExitRpcMethodMessage is called when production rpcMethodMessage is exited.
func (s *BaseProtoLSDListener) ExitRpcMethodMessage(ctx *RpcMethodMessageContext) {}

// EnterRpcSingleMessage is called when production rpcSingleMessage is entered.
func (s *BaseProtoLSDListener) EnterRpcSingleMessage(ctx *RpcSingleMessageContext) {}

// ExitRpcSingleMessage is called when production rpcSingleMessage is exited.
func (s *BaseProtoLSDListener) ExitRpcSingleMessage(ctx *RpcSingleMessageContext) {}

// EnterRpcMessageValue is called when production rpcMessageValue is entered.
func (s *BaseProtoLSDListener) EnterRpcMessageValue(ctx *RpcMessageValueContext) {}

// ExitRpcMessageValue is called when production rpcMessageValue is exited.
func (s *BaseProtoLSDListener) ExitRpcMessageValue(ctx *RpcMessageValueContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BaseProtoLSDListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BaseProtoLSDListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseProtoLSDListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseProtoLSDListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseProtoLSDListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseProtoLSDListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterMessageDefinition is called when production messageDefinition is entered.
func (s *BaseProtoLSDListener) EnterMessageDefinition(ctx *MessageDefinitionContext) {}

// ExitMessageDefinition is called when production messageDefinition is exited.
func (s *BaseProtoLSDListener) ExitMessageDefinition(ctx *MessageDefinitionContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BaseProtoLSDListener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BaseProtoLSDListener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterMessageBodyStatement is called when production messageBodyStatement is entered.
func (s *BaseProtoLSDListener) EnterMessageBodyStatement(ctx *MessageBodyStatementContext) {}

// ExitMessageBodyStatement is called when production messageBodyStatement is exited.
func (s *BaseProtoLSDListener) ExitMessageBodyStatement(ctx *MessageBodyStatementContext) {}

// EnterMessageField is called when production messageField is entered.
func (s *BaseProtoLSDListener) EnterMessageField(ctx *MessageFieldContext) {}

// ExitMessageField is called when production messageField is exited.
func (s *BaseProtoLSDListener) ExitMessageField(ctx *MessageFieldContext) {}

// EnterPreprocessorParameters is called when production preprocessorParameters is entered.
func (s *BaseProtoLSDListener) EnterPreprocessorParameters(ctx *PreprocessorParametersContext) {}

// ExitPreprocessorParameters is called when production preprocessorParameters is exited.
func (s *BaseProtoLSDListener) ExitPreprocessorParameters(ctx *PreprocessorParametersContext) {}

// EnterPreprocessorDirective is called when production preprocessorDirective is entered.
func (s *BaseProtoLSDListener) EnterPreprocessorDirective(ctx *PreprocessorDirectiveContext) {}

// ExitPreprocessorDirective is called when production preprocessorDirective is exited.
func (s *BaseProtoLSDListener) ExitPreprocessorDirective(ctx *PreprocessorDirectiveContext) {}

// EnterAnnotationDirective is called when production annotationDirective is entered.
func (s *BaseProtoLSDListener) EnterAnnotationDirective(ctx *AnnotationDirectiveContext) {}

// ExitAnnotationDirective is called when production annotationDirective is exited.
func (s *BaseProtoLSDListener) ExitAnnotationDirective(ctx *AnnotationDirectiveContext) {}

// EnterOptionalModifier is called when production optionalModifier is entered.
func (s *BaseProtoLSDListener) EnterOptionalModifier(ctx *OptionalModifierContext) {}

// ExitOptionalModifier is called when production optionalModifier is exited.
func (s *BaseProtoLSDListener) ExitOptionalModifier(ctx *OptionalModifierContext) {}

// EnterOptionStatement is called when production optionStatement is entered.
func (s *BaseProtoLSDListener) EnterOptionStatement(ctx *OptionStatementContext) {}

// ExitOptionStatement is called when production optionStatement is exited.
func (s *BaseProtoLSDListener) ExitOptionStatement(ctx *OptionStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BaseProtoLSDListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BaseProtoLSDListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseProtoLSDListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseProtoLSDListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDataTypeText is called when production dataTypeText is entered.
func (s *BaseProtoLSDListener) EnterDataTypeText(ctx *DataTypeTextContext) {}

// ExitDataTypeText is called when production dataTypeText is exited.
func (s *BaseProtoLSDListener) ExitDataTypeText(ctx *DataTypeTextContext) {}
