// Code generated from ProtoLSD.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ProtoLSD

import "github.com/antlr4-go/antlr/v4"

// ProtoLSDListener is a complete listener for a parse tree produced by ProtoLSDParser.
type ProtoLSDListener interface {
	antlr.ParseTreeListener

	// EnterSemi is called when entering the semi production.
	EnterSemi(c *SemiContext)

	// EnterProtoLSD is called when entering the protoLSD production.
	EnterProtoLSD(c *ProtoLSDContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterServiceDefinition is called when entering the serviceDefinition production.
	EnterServiceDefinition(c *ServiceDefinitionContext)

	// EnterServiceBody is called when entering the serviceBody production.
	EnterServiceBody(c *ServiceBodyContext)

	// EnterRpcDefinition is called when entering the rpcDefinition production.
	EnterRpcDefinition(c *RpcDefinitionContext)

	// EnterRpcParameters is called when entering the rpcParameters production.
	EnterRpcParameters(c *RpcParametersContext)

	// EnterRpcParameter is called when entering the rpcParameter production.
	EnterRpcParameter(c *RpcParameterContext)

	// EnterRpcParameterModifier is called when entering the rpcParameterModifier production.
	EnterRpcParameterModifier(c *RpcParameterModifierContext)

	// EnterRpcReturn is called when entering the rpcReturn production.
	EnterRpcReturn(c *RpcReturnContext)

	// EnterRpcInput is called when entering the rpcInput production.
	EnterRpcInput(c *RpcInputContext)

	// EnterRpcMethodMessage is called when entering the rpcMethodMessage production.
	EnterRpcMethodMessage(c *RpcMethodMessageContext)

	// EnterRpcSingleMessage is called when entering the rpcSingleMessage production.
	EnterRpcSingleMessage(c *RpcSingleMessageContext)

	// EnterRpcMessageValue is called when entering the rpcMessageValue production.
	EnterRpcMessageValue(c *RpcMessageValueContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterMessageDefinition is called when entering the messageDefinition production.
	EnterMessageDefinition(c *MessageDefinitionContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterMessageBodyStatement is called when entering the messageBodyStatement production.
	EnterMessageBodyStatement(c *MessageBodyStatementContext)

	// EnterMessageField is called when entering the messageField production.
	EnterMessageField(c *MessageFieldContext)

	// EnterPreprocessorParameters is called when entering the preprocessorParameters production.
	EnterPreprocessorParameters(c *PreprocessorParametersContext)

	// EnterPreprocessorDirective is called when entering the preprocessorDirective production.
	EnterPreprocessorDirective(c *PreprocessorDirectiveContext)

	// EnterAnnotationDirective is called when entering the annotationDirective production.
	EnterAnnotationDirective(c *AnnotationDirectiveContext)

	// EnterOptionalModifier is called when entering the optionalModifier production.
	EnterOptionalModifier(c *OptionalModifierContext)

	// EnterOptionStatement is called when entering the optionStatement production.
	EnterOptionStatement(c *OptionStatementContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterDataTypeText is called when entering the dataTypeText production.
	EnterDataTypeText(c *DataTypeTextContext)

	// ExitSemi is called when exiting the semi production.
	ExitSemi(c *SemiContext)

	// ExitProtoLSD is called when exiting the protoLSD production.
	ExitProtoLSD(c *ProtoLSDContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitServiceDefinition is called when exiting the serviceDefinition production.
	ExitServiceDefinition(c *ServiceDefinitionContext)

	// ExitServiceBody is called when exiting the serviceBody production.
	ExitServiceBody(c *ServiceBodyContext)

	// ExitRpcDefinition is called when exiting the rpcDefinition production.
	ExitRpcDefinition(c *RpcDefinitionContext)

	// ExitRpcParameters is called when exiting the rpcParameters production.
	ExitRpcParameters(c *RpcParametersContext)

	// ExitRpcParameter is called when exiting the rpcParameter production.
	ExitRpcParameter(c *RpcParameterContext)

	// ExitRpcParameterModifier is called when exiting the rpcParameterModifier production.
	ExitRpcParameterModifier(c *RpcParameterModifierContext)

	// ExitRpcReturn is called when exiting the rpcReturn production.
	ExitRpcReturn(c *RpcReturnContext)

	// ExitRpcInput is called when exiting the rpcInput production.
	ExitRpcInput(c *RpcInputContext)

	// ExitRpcMethodMessage is called when exiting the rpcMethodMessage production.
	ExitRpcMethodMessage(c *RpcMethodMessageContext)

	// ExitRpcSingleMessage is called when exiting the rpcSingleMessage production.
	ExitRpcSingleMessage(c *RpcSingleMessageContext)

	// ExitRpcMessageValue is called when exiting the rpcMessageValue production.
	ExitRpcMessageValue(c *RpcMessageValueContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitMessageDefinition is called when exiting the messageDefinition production.
	ExitMessageDefinition(c *MessageDefinitionContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitMessageBodyStatement is called when exiting the messageBodyStatement production.
	ExitMessageBodyStatement(c *MessageBodyStatementContext)

	// ExitMessageField is called when exiting the messageField production.
	ExitMessageField(c *MessageFieldContext)

	// ExitPreprocessorParameters is called when exiting the preprocessorParameters production.
	ExitPreprocessorParameters(c *PreprocessorParametersContext)

	// ExitPreprocessorDirective is called when exiting the preprocessorDirective production.
	ExitPreprocessorDirective(c *PreprocessorDirectiveContext)

	// ExitAnnotationDirective is called when exiting the annotationDirective production.
	ExitAnnotationDirective(c *AnnotationDirectiveContext)

	// ExitOptionalModifier is called when exiting the optionalModifier production.
	ExitOptionalModifier(c *OptionalModifierContext)

	// ExitOptionStatement is called when exiting the optionStatement production.
	ExitOptionStatement(c *OptionStatementContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitDataTypeText is called when exiting the dataTypeText production.
	ExitDataTypeText(c *DataTypeTextContext)
}
