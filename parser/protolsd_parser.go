// Code generated from ProtoLSD.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ProtoLSD

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ProtoLSDParser struct {
	*antlr.BaseParser
}

var ProtoLSDParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protolsdParserInit() {
	staticData := &ProtoLSDParserStaticData
	staticData.LiteralNames = []string{
		"", "'*'", "'int'", "'uint'", "'long'", "'ulong'", "'import'", "'private'",
		"'service'", "'rpc'", "'enum'", "'message'", "'returns'", "'option'",
		"'optional'", "'required'", "'repeated'", "'package'", "", "", "", "",
		"", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "','",
		"'='", "'.'", "'#'", "'?'", "'[]'", "'{}'", "'!'", "'+'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "IMPORT", "PRIVATE", "SERVICE", "RPC", "ENUM",
		"MESSAGE", "RETURNS", "OPTION", "OPTIONAL", "REQUIRED", "REPEATED",
		"PACKAGE", "LINE_COMMENT", "BLOCK_COMMENT", "IDENTIFIER", "PREPROCESSOR_NAME",
		"STRING", "NUMBER", "MAP", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET",
		"RBRACKET", "SEMI", "COMMA", "EQUAL", "DOT", "HASH", "QUESTION", "ARRAY",
		"OBJECT", "BANG", "PLUS", "ANNOTATION", "WS",
	}
	staticData.RuleNames = []string{
		"semi", "protoLSD", "statement", "importStatement", "serviceDefinition",
		"serviceBody", "rpcDefinition", "rpcParameters", "rpcParameter", "rpcParameterModifier",
		"rpcReturn", "rpcInput", "rpcMethodMessage", "rpcSingleMessage", "rpcMessageValue",
		"enumDefinition", "enumBody", "enumField", "messageDefinition", "messageBody",
		"messageBodyStatement", "messageField", "preprocessorParameters", "preprocessorDirective",
		"annotationDirective", "optionalModifier", "optionStatement", "packageStatement",
		"dataType", "dataTypeText",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 280, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 3, 0, 62, 8, 0,
		1, 1, 5, 1, 65, 8, 1, 10, 1, 12, 1, 68, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 77, 8, 2, 1, 3, 3, 3, 80, 8, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 85, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 5,
		5, 96, 8, 5, 10, 5, 12, 5, 99, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 104, 8, 6,
		1, 6, 3, 6, 107, 8, 6, 1, 6, 3, 6, 110, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 116, 8, 7, 10, 7, 12, 7, 119, 9, 7, 3, 7, 121, 8, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 3, 8, 127, 8, 8, 1, 8, 3, 8, 130, 8, 8, 1, 8, 3, 8, 133, 8, 8,
		1, 8, 3, 8, 136, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 141, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 151, 8, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 158, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14,
		163, 8, 14, 10, 14, 12, 14, 166, 9, 14, 1, 15, 5, 15, 169, 8, 15, 10, 15,
		12, 15, 172, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 184, 8, 16, 10, 16, 12, 16, 187, 9, 16, 1, 17,
		1, 17, 1, 17, 3, 17, 192, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 5, 19, 201, 8, 19, 10, 19, 12, 19, 204, 9, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 3, 20, 210, 8, 20, 1, 21, 3, 21, 213, 8, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 220, 8, 21, 3, 21, 222, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 5, 22, 228, 8, 22, 10, 22, 12, 22, 231, 9, 22, 3, 22, 233,
		8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 240, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 5, 27, 257, 8, 27, 10, 27, 12, 27, 260, 9, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 3, 28, 266, 8, 28, 1, 28, 3, 28, 269, 8, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 278, 8, 29, 1, 29,
		0, 0, 30, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 0, 2, 2, 0, 31, 31,
		38, 38, 1, 0, 14, 15, 294, 0, 61, 1, 0, 0, 0, 2, 66, 1, 0, 0, 0, 4, 76,
		1, 0, 0, 0, 6, 79, 1, 0, 0, 0, 8, 88, 1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12,
		100, 1, 0, 0, 0, 14, 111, 1, 0, 0, 0, 16, 135, 1, 0, 0, 0, 18, 140, 1,
		0, 0, 0, 20, 142, 1, 0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 150, 1, 0, 0, 0,
		26, 157, 1, 0, 0, 0, 28, 159, 1, 0, 0, 0, 30, 170, 1, 0, 0, 0, 32, 179,
		1, 0, 0, 0, 34, 188, 1, 0, 0, 0, 36, 193, 1, 0, 0, 0, 38, 202, 1, 0, 0,
		0, 40, 209, 1, 0, 0, 0, 42, 212, 1, 0, 0, 0, 44, 223, 1, 0, 0, 0, 46, 236,
		1, 0, 0, 0, 48, 241, 1, 0, 0, 0, 50, 244, 1, 0, 0, 0, 52, 246, 1, 0, 0,
		0, 54, 252, 1, 0, 0, 0, 56, 263, 1, 0, 0, 0, 58, 277, 1, 0, 0, 0, 60, 62,
		5, 31, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 1, 1, 0, 0, 0,
		63, 65, 3, 4, 2, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 3, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69,
		77, 3, 6, 3, 0, 70, 77, 3, 8, 4, 0, 71, 77, 3, 30, 15, 0, 72, 77, 3, 36,
		18, 0, 73, 77, 3, 46, 23, 0, 74, 77, 3, 52, 26, 0, 75, 77, 3, 54, 27, 0,
		76, 69, 1, 0, 0, 0, 76, 70, 1, 0, 0, 0, 76, 71, 1, 0, 0, 0, 76, 72, 1,
		0, 0, 0, 76, 73, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77,
		5, 1, 0, 0, 0, 78, 80, 5, 7, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0,
		0, 80, 81, 1, 0, 0, 0, 81, 82, 5, 6, 0, 0, 82, 84, 5, 22, 0, 0, 83, 85,
		5, 39, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 87, 3, 0, 0, 0, 87, 7, 1, 0, 0, 0, 88, 89, 5, 8, 0, 0, 89, 90, 5, 20,
		0, 0, 90, 91, 5, 27, 0, 0, 91, 92, 3, 10, 5, 0, 92, 93, 5, 28, 0, 0, 93,
		9, 1, 0, 0, 0, 94, 96, 3, 12, 6, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0,
		0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 11, 1, 0, 0, 0, 99, 97,
		1, 0, 0, 0, 100, 101, 5, 9, 0, 0, 101, 103, 5, 20, 0, 0, 102, 104, 3, 22,
		11, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0,
		105, 107, 3, 20, 10, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107,
		109, 1, 0, 0, 0, 108, 110, 7, 0, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110,
		1, 0, 0, 0, 110, 13, 1, 0, 0, 0, 111, 120, 5, 25, 0, 0, 112, 117, 3, 16,
		8, 0, 113, 114, 5, 32, 0, 0, 114, 116, 3, 16, 8, 0, 115, 113, 1, 0, 0,
		0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 112, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 5, 26, 0, 0, 123, 15, 1, 0,
		0, 0, 124, 126, 3, 56, 28, 0, 125, 127, 5, 40, 0, 0, 126, 125, 1, 0, 0,
		0, 126, 127, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128, 130, 3, 18, 9, 0, 129,
		128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 136, 1, 0, 0, 0, 131, 133,
		3, 56, 28, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1,
		0, 0, 0, 134, 136, 5, 20, 0, 0, 135, 124, 1, 0, 0, 0, 135, 132, 1, 0, 0,
		0, 136, 17, 1, 0, 0, 0, 137, 141, 5, 20, 0, 0, 138, 139, 5, 1, 0, 0, 139,
		141, 5, 23, 0, 0, 140, 137, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 19,
		1, 0, 0, 0, 142, 143, 5, 12, 0, 0, 143, 144, 3, 24, 12, 0, 144, 21, 1,
		0, 0, 0, 145, 146, 3, 24, 12, 0, 146, 23, 1, 0, 0, 0, 147, 151, 5, 38,
		0, 0, 148, 151, 3, 26, 13, 0, 149, 151, 3, 14, 7, 0, 150, 147, 1, 0, 0,
		0, 150, 148, 1, 0, 0, 0, 150, 149, 1, 0, 0, 0, 151, 25, 1, 0, 0, 0, 152,
		153, 5, 25, 0, 0, 153, 154, 3, 28, 14, 0, 154, 155, 5, 26, 0, 0, 155, 158,
		1, 0, 0, 0, 156, 158, 3, 28, 14, 0, 157, 152, 1, 0, 0, 0, 157, 156, 1,
		0, 0, 0, 158, 27, 1, 0, 0, 0, 159, 164, 5, 20, 0, 0, 160, 161, 5, 34, 0,
		0, 161, 163, 5, 20, 0, 0, 162, 160, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164,
		162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 29, 1, 0, 0, 0, 166, 164, 1,
		0, 0, 0, 167, 169, 3, 48, 24, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0,
		0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0, 0,
		172, 170, 1, 0, 0, 0, 173, 174, 5, 10, 0, 0, 174, 175, 5, 20, 0, 0, 175,
		176, 5, 27, 0, 0, 176, 177, 3, 32, 16, 0, 177, 178, 5, 28, 0, 0, 178, 31,
		1, 0, 0, 0, 179, 185, 3, 34, 17, 0, 180, 181, 3, 0, 0, 0, 181, 182, 3,
		34, 17, 0, 182, 184, 1, 0, 0, 0, 183, 180, 1, 0, 0, 0, 184, 187, 1, 0,
		0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 33, 1, 0, 0, 0,
		187, 185, 1, 0, 0, 0, 188, 191, 5, 20, 0, 0, 189, 190, 5, 33, 0, 0, 190,
		192, 5, 23, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 35,
		1, 0, 0, 0, 193, 194, 5, 11, 0, 0, 194, 195, 5, 20, 0, 0, 195, 196, 5,
		27, 0, 0, 196, 197, 3, 38, 19, 0, 197, 198, 5, 28, 0, 0, 198, 37, 1, 0,
		0, 0, 199, 201, 3, 40, 20, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 39, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 206, 3, 42, 21, 0, 206, 207, 3, 0, 0, 0, 207, 210,
		1, 0, 0, 0, 208, 210, 3, 36, 18, 0, 209, 205, 1, 0, 0, 0, 209, 208, 1,
		0, 0, 0, 210, 41, 1, 0, 0, 0, 211, 213, 3, 50, 25, 0, 212, 211, 1, 0, 0,
		0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 3, 56, 28, 0,
		215, 221, 5, 20, 0, 0, 216, 217, 5, 33, 0, 0, 217, 219, 5, 23, 0, 0, 218,
		220, 5, 39, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222,
		1, 0, 0, 0, 221, 216, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 43, 1, 0,
		0, 0, 223, 232, 5, 25, 0, 0, 224, 229, 3, 28, 14, 0, 225, 226, 5, 32, 0,
		0, 226, 228, 3, 28, 14, 0, 227, 225, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0,
		229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231,
		229, 1, 0, 0, 0, 232, 224, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 235, 5, 26, 0, 0, 235, 45, 1, 0, 0, 0, 236, 237, 5, 35,
		0, 0, 237, 239, 5, 21, 0, 0, 238, 240, 3, 44, 22, 0, 239, 238, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 47, 1, 0, 0, 0, 241, 242, 5, 41, 0, 0, 242,
		243, 5, 20, 0, 0, 243, 49, 1, 0, 0, 0, 244, 245, 7, 1, 0, 0, 245, 51, 1,
		0, 0, 0, 246, 247, 5, 13, 0, 0, 247, 248, 5, 20, 0, 0, 248, 249, 5, 33,
		0, 0, 249, 250, 5, 22, 0, 0, 250, 251, 3, 0, 0, 0, 251, 53, 1, 0, 0, 0,
		252, 253, 5, 17, 0, 0, 253, 258, 5, 20, 0, 0, 254, 255, 5, 34, 0, 0, 255,
		257, 5, 20, 0, 0, 256, 254, 1, 0, 0, 0, 257, 260, 1, 0, 0, 0, 258, 256,
		1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261, 1, 0, 0, 0, 260, 258, 1, 0,
		0, 0, 261, 262, 3, 0, 0, 0, 262, 55, 1, 0, 0, 0, 263, 265, 3, 58, 29, 0,
		264, 266, 5, 37, 0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266,
		268, 1, 0, 0, 0, 267, 269, 5, 36, 0, 0, 268, 267, 1, 0, 0, 0, 268, 269,
		1, 0, 0, 0, 269, 57, 1, 0, 0, 0, 270, 278, 5, 2, 0, 0, 271, 278, 5, 3,
		0, 0, 272, 278, 5, 4, 0, 0, 273, 278, 5, 5, 0, 0, 274, 278, 5, 24, 0, 0,
		275, 278, 5, 20, 0, 0, 276, 278, 3, 28, 14, 0, 277, 270, 1, 0, 0, 0, 277,
		271, 1, 0, 0, 0, 277, 272, 1, 0, 0, 0, 277, 273, 1, 0, 0, 0, 277, 274,
		1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278, 59, 1, 0,
		0, 0, 34, 61, 66, 76, 79, 84, 97, 103, 106, 109, 117, 120, 126, 129, 132,
		135, 140, 150, 157, 164, 170, 185, 191, 202, 209, 212, 219, 221, 229, 232,
		239, 258, 265, 268, 277,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ProtoLSDParserInit initializes any static state used to implement ProtoLSDParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtoLSDParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProtoLSDParserInit() {
	staticData := &ProtoLSDParserStaticData
	staticData.once.Do(protolsdParserInit)
}

// NewProtoLSDParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtoLSDParser(input antlr.TokenStream) *ProtoLSDParser {
	ProtoLSDParserInit()
	this := new(ProtoLSDParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ProtoLSDParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ProtoLSD.g4"

	return this
}

// ProtoLSDParser tokens.
const (
	ProtoLSDParserEOF               = antlr.TokenEOF
	ProtoLSDParserT__0              = 1
	ProtoLSDParserT__1              = 2
	ProtoLSDParserT__2              = 3
	ProtoLSDParserT__3              = 4
	ProtoLSDParserT__4              = 5
	ProtoLSDParserIMPORT            = 6
	ProtoLSDParserPRIVATE           = 7
	ProtoLSDParserSERVICE           = 8
	ProtoLSDParserRPC               = 9
	ProtoLSDParserENUM              = 10
	ProtoLSDParserMESSAGE           = 11
	ProtoLSDParserRETURNS           = 12
	ProtoLSDParserOPTION            = 13
	ProtoLSDParserOPTIONAL          = 14
	ProtoLSDParserREQUIRED          = 15
	ProtoLSDParserREPEATED          = 16
	ProtoLSDParserPACKAGE           = 17
	ProtoLSDParserLINE_COMMENT      = 18
	ProtoLSDParserBLOCK_COMMENT     = 19
	ProtoLSDParserIDENTIFIER        = 20
	ProtoLSDParserPREPROCESSOR_NAME = 21
	ProtoLSDParserSTRING            = 22
	ProtoLSDParserNUMBER            = 23
	ProtoLSDParserMAP               = 24
	ProtoLSDParserLPAREN            = 25
	ProtoLSDParserRPAREN            = 26
	ProtoLSDParserLBRACE            = 27
	ProtoLSDParserRBRACE            = 28
	ProtoLSDParserLBRACKET          = 29
	ProtoLSDParserRBRACKET          = 30
	ProtoLSDParserSEMI              = 31
	ProtoLSDParserCOMMA             = 32
	ProtoLSDParserEQUAL             = 33
	ProtoLSDParserDOT               = 34
	ProtoLSDParserHASH              = 35
	ProtoLSDParserQUESTION          = 36
	ProtoLSDParserARRAY             = 37
	ProtoLSDParserOBJECT            = 38
	ProtoLSDParserBANG              = 39
	ProtoLSDParserPLUS              = 40
	ProtoLSDParserANNOTATION        = 41
	ProtoLSDParserWS                = 42
)

// ProtoLSDParser rules.
const (
	ProtoLSDParserRULE_semi                   = 0
	ProtoLSDParserRULE_protoLSD               = 1
	ProtoLSDParserRULE_statement              = 2
	ProtoLSDParserRULE_importStatement        = 3
	ProtoLSDParserRULE_serviceDefinition      = 4
	ProtoLSDParserRULE_serviceBody            = 5
	ProtoLSDParserRULE_rpcDefinition          = 6
	ProtoLSDParserRULE_rpcParameters          = 7
	ProtoLSDParserRULE_rpcParameter           = 8
	ProtoLSDParserRULE_rpcParameterModifier   = 9
	ProtoLSDParserRULE_rpcReturn              = 10
	ProtoLSDParserRULE_rpcInput               = 11
	ProtoLSDParserRULE_rpcMethodMessage       = 12
	ProtoLSDParserRULE_rpcSingleMessage       = 13
	ProtoLSDParserRULE_rpcMessageValue        = 14
	ProtoLSDParserRULE_enumDefinition         = 15
	ProtoLSDParserRULE_enumBody               = 16
	ProtoLSDParserRULE_enumField              = 17
	ProtoLSDParserRULE_messageDefinition      = 18
	ProtoLSDParserRULE_messageBody            = 19
	ProtoLSDParserRULE_messageBodyStatement   = 20
	ProtoLSDParserRULE_messageField           = 21
	ProtoLSDParserRULE_preprocessorParameters = 22
	ProtoLSDParserRULE_preprocessorDirective  = 23
	ProtoLSDParserRULE_annotationDirective    = 24
	ProtoLSDParserRULE_optionalModifier       = 25
	ProtoLSDParserRULE_optionStatement        = 26
	ProtoLSDParserRULE_packageStatement       = 27
	ProtoLSDParserRULE_dataType               = 28
	ProtoLSDParserRULE_dataTypeText           = 29
)

// ISemiContext is an interface to support dynamic dispatch.
type ISemiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsSemiContext differentiates from other interfaces.
	IsSemiContext()
}

type SemiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySemiContext() *SemiContext {
	var p = new(SemiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_semi
	return p
}

func InitEmptySemiContext(p *SemiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_semi
}

func (*SemiContext) IsSemiContext() {}

func NewSemiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SemiContext {
	var p = new(SemiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_semi

	return p
}

func (s *SemiContext) GetParser() antlr.Parser { return s.parser }

func (s *SemiContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserSEMI, 0)
}

func (s *SemiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SemiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SemiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterSemi(s)
	}
}

func (s *SemiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitSemi(s)
	}
}

func (p *ProtoLSDParser) Semi() (localctx ISemiContext) {
	localctx = NewSemiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProtoLSDParserRULE_semi)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserSEMI {
		{
			p.SetState(60)
			p.Match(ProtoLSDParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProtoLSDContext is an interface to support dynamic dispatch.
type IProtoLSDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProtoLSDContext differentiates from other interfaces.
	IsProtoLSDContext()
}

type ProtoLSDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoLSDContext() *ProtoLSDContext {
	var p = new(ProtoLSDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_protoLSD
	return p
}

func InitEmptyProtoLSDContext(p *ProtoLSDContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_protoLSD
}

func (*ProtoLSDContext) IsProtoLSDContext() {}

func NewProtoLSDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoLSDContext {
	var p = new(ProtoLSDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_protoLSD

	return p
}

func (s *ProtoLSDContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoLSDContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoLSDContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProtoLSDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoLSDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoLSDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterProtoLSD(s)
	}
}

func (s *ProtoLSDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitProtoLSD(s)
	}
}

func (p *ProtoLSDParser) ProtoLSD() (localctx IProtoLSDContext) {
	localctx = NewProtoLSDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProtoLSDParserRULE_protoLSD)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2233383136704) != 0 {
		{
			p.SetState(63)
			p.Statement()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportStatement() IImportStatementContext
	ServiceDefinition() IServiceDefinitionContext
	EnumDefinition() IEnumDefinitionContext
	MessageDefinition() IMessageDefinitionContext
	PreprocessorDirective() IPreprocessorDirectiveContext
	OptionStatement() IOptionStatementContext
	PackageStatement() IPackageStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportStatement() IImportStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *StatementContext) ServiceDefinition() IServiceDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceDefinitionContext)
}

func (s *StatementContext) EnumDefinition() IEnumDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *StatementContext) MessageDefinition() IMessageDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefinitionContext)
}

func (s *StatementContext) PreprocessorDirective() IPreprocessorDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreprocessorDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreprocessorDirectiveContext)
}

func (s *StatementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *StatementContext) PackageStatement() IPackageStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ProtoLSDParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProtoLSDParserRULE_statement)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProtoLSDParserIMPORT, ProtoLSDParserPRIVATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.ImportStatement()
		}

	case ProtoLSDParserSERVICE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.ServiceDefinition()
		}

	case ProtoLSDParserENUM, ProtoLSDParserANNOTATION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.EnumDefinition()
		}

	case ProtoLSDParserMESSAGE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.MessageDefinition()
		}

	case ProtoLSDParserHASH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.PreprocessorDirective()
		}

	case ProtoLSDParserOPTION:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.OptionStatement()
		}

	case ProtoLSDParserPACKAGE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(75)
			p.PackageStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Semi() ISemiContext
	PRIVATE() antlr.TerminalNode
	BANG() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIMPORT, 0)
}

func (s *ImportStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserSTRING, 0)
}

func (s *ImportStatementContext) Semi() ISemiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiContext)
}

func (s *ImportStatementContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserPRIVATE, 0)
}

func (s *ImportStatementContext) BANG() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserBANG, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *ProtoLSDParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProtoLSDParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserPRIVATE {
		{
			p.SetState(78)
			p.Match(ProtoLSDParserPRIVATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(81)
		p.Match(ProtoLSDParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(ProtoLSDParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserBANG {
		{
			p.SetState(83)
			p.Match(ProtoLSDParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(86)
		p.Semi()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceDefinitionContext is an interface to support dynamic dispatch.
type IServiceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	ServiceBody() IServiceBodyContext
	RBRACE() antlr.TerminalNode

	// IsServiceDefinitionContext differentiates from other interfaces.
	IsServiceDefinitionContext()
}

type ServiceDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefinitionContext() *ServiceDefinitionContext {
	var p = new(ServiceDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_serviceDefinition
	return p
}

func InitEmptyServiceDefinitionContext(p *ServiceDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_serviceDefinition
}

func (*ServiceDefinitionContext) IsServiceDefinitionContext() {}

func NewServiceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefinitionContext {
	var p = new(ServiceDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_serviceDefinition

	return p
}

func (s *ServiceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefinitionContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserSERVICE, 0)
}

func (s *ServiceDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *ServiceDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLBRACE, 0)
}

func (s *ServiceDefinitionContext) ServiceBody() IServiceBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceBodyContext)
}

func (s *ServiceDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRBRACE, 0)
}

func (s *ServiceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterServiceDefinition(s)
	}
}

func (s *ServiceDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitServiceDefinition(s)
	}
}

func (p *ProtoLSDParser) ServiceDefinition() (localctx IServiceDefinitionContext) {
	localctx = NewServiceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProtoLSDParserRULE_serviceDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ProtoLSDParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(ProtoLSDParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.ServiceBody()
	}
	{
		p.SetState(92)
		p.Match(ProtoLSDParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceBodyContext is an interface to support dynamic dispatch.
type IServiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRpcDefinition() []IRpcDefinitionContext
	RpcDefinition(i int) IRpcDefinitionContext

	// IsServiceBodyContext differentiates from other interfaces.
	IsServiceBodyContext()
}

type ServiceBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceBodyContext() *ServiceBodyContext {
	var p = new(ServiceBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_serviceBody
	return p
}

func InitEmptyServiceBodyContext(p *ServiceBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_serviceBody
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_serviceBody

	return p
}

func (s *ServiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBodyContext) AllRpcDefinition() []IRpcDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRpcDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IRpcDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRpcDefinitionContext); ok {
			tst[i] = t.(IRpcDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ServiceBodyContext) RpcDefinition(i int) IRpcDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcDefinitionContext)
}

func (s *ServiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterServiceBody(s)
	}
}

func (s *ServiceBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitServiceBody(s)
	}
}

func (p *ProtoLSDParser) ServiceBody() (localctx IServiceBodyContext) {
	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProtoLSDParserRULE_serviceBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProtoLSDParserRPC {
		{
			p.SetState(94)
			p.RpcDefinition()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcDefinitionContext is an interface to support dynamic dispatch.
type IRpcDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RpcInput() IRpcInputContext
	RpcReturn() IRpcReturnContext
	SEMI() antlr.TerminalNode
	OBJECT() antlr.TerminalNode

	// IsRpcDefinitionContext differentiates from other interfaces.
	IsRpcDefinitionContext()
}

type RpcDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcDefinitionContext() *RpcDefinitionContext {
	var p = new(RpcDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcDefinition
	return p
}

func InitEmptyRpcDefinitionContext(p *RpcDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcDefinition
}

func (*RpcDefinitionContext) IsRpcDefinitionContext() {}

func NewRpcDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcDefinitionContext {
	var p = new(RpcDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcDefinition

	return p
}

func (s *RpcDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcDefinitionContext) RPC() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRPC, 0)
}

func (s *RpcDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *RpcDefinitionContext) RpcInput() IRpcInputContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcInputContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcInputContext)
}

func (s *RpcDefinitionContext) RpcReturn() IRpcReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcReturnContext)
}

func (s *RpcDefinitionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserSEMI, 0)
}

func (s *RpcDefinitionContext) OBJECT() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserOBJECT, 0)
}

func (s *RpcDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcDefinition(s)
	}
}

func (s *RpcDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcDefinition(s)
	}
}

func (p *ProtoLSDParser) RpcDefinition() (localctx IRpcDefinitionContext) {
	localctx = NewRpcDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProtoLSDParserRULE_rpcDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(ProtoLSDParserRPC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(102)
			p.RpcInput()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserRETURNS {
		{
			p.SetState(105)
			p.RpcReturn()
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserSEMI || _la == ProtoLSDParserOBJECT {
		{
			p.SetState(108)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ProtoLSDParserSEMI || _la == ProtoLSDParserOBJECT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcParametersContext is an interface to support dynamic dispatch.
type IRpcParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllRpcParameter() []IRpcParameterContext
	RpcParameter(i int) IRpcParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRpcParametersContext differentiates from other interfaces.
	IsRpcParametersContext()
}

type RpcParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcParametersContext() *RpcParametersContext {
	var p = new(RpcParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameters
	return p
}

func InitEmptyRpcParametersContext(p *RpcParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameters
}

func (*RpcParametersContext) IsRpcParametersContext() {}

func NewRpcParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParametersContext {
	var p = new(RpcParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcParameters

	return p
}

func (s *RpcParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLPAREN, 0)
}

func (s *RpcParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRPAREN, 0)
}

func (s *RpcParametersContext) AllRpcParameter() []IRpcParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRpcParameterContext); ok {
			len++
		}
	}

	tst := make([]IRpcParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRpcParameterContext); ok {
			tst[i] = t.(IRpcParameterContext)
			i++
		}
	}

	return tst
}

func (s *RpcParametersContext) RpcParameter(i int) IRpcParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcParameterContext)
}

func (s *RpcParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserCOMMA)
}

func (s *RpcParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserCOMMA, i)
}

func (s *RpcParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcParameters(s)
	}
}

func (s *RpcParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcParameters(s)
	}
}

func (p *ProtoLSDParser) RpcParameters() (localctx IRpcParametersContext) {
	localctx = NewRpcParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProtoLSDParserRULE_rpcParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ProtoLSDParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17825852) != 0 {
		{
			p.SetState(112)
			p.RpcParameter()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ProtoLSDParserCOMMA {
			{
				p.SetState(113)
				p.Match(ProtoLSDParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.RpcParameter()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(122)
		p.Match(ProtoLSDParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcParameterContext is an interface to support dynamic dispatch.
type IRpcParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataType() IDataTypeContext
	PLUS() antlr.TerminalNode
	RpcParameterModifier() IRpcParameterModifierContext
	IDENTIFIER() antlr.TerminalNode

	// IsRpcParameterContext differentiates from other interfaces.
	IsRpcParameterContext()
}

type RpcParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcParameterContext() *RpcParameterContext {
	var p = new(RpcParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameter
	return p
}

func InitEmptyRpcParameterContext(p *RpcParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameter
}

func (*RpcParameterContext) IsRpcParameterContext() {}

func NewRpcParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParameterContext {
	var p = new(RpcParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcParameter

	return p
}

func (s *RpcParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParameterContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *RpcParameterContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserPLUS, 0)
}

func (s *RpcParameterContext) RpcParameterModifier() IRpcParameterModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcParameterModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcParameterModifierContext)
}

func (s *RpcParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *RpcParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcParameter(s)
	}
}

func (s *RpcParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcParameter(s)
	}
}

func (p *ProtoLSDParser) RpcParameter() (localctx IRpcParameterContext) {
	localctx = NewRpcParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProtoLSDParserRULE_rpcParameter)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.DataType()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ProtoLSDParserPLUS {
			{
				p.SetState(125)
				p.Match(ProtoLSDParserPLUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ProtoLSDParserT__0 || _la == ProtoLSDParserIDENTIFIER {
			{
				p.SetState(128)
				p.RpcParameterModifier()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(132)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(131)
				p.DataType()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(134)
			p.Match(ProtoLSDParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcParameterModifierContext is an interface to support dynamic dispatch.
type IRpcParameterModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsRpcParameterModifierContext differentiates from other interfaces.
	IsRpcParameterModifierContext()
}

type RpcParameterModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcParameterModifierContext() *RpcParameterModifierContext {
	var p = new(RpcParameterModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameterModifier
	return p
}

func InitEmptyRpcParameterModifierContext(p *RpcParameterModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcParameterModifier
}

func (*RpcParameterModifierContext) IsRpcParameterModifierContext() {}

func NewRpcParameterModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParameterModifierContext {
	var p = new(RpcParameterModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcParameterModifier

	return p
}

func (s *RpcParameterModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParameterModifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *RpcParameterModifierContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserNUMBER, 0)
}

func (s *RpcParameterModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParameterModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParameterModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcParameterModifier(s)
	}
}

func (s *RpcParameterModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcParameterModifier(s)
	}
}

func (p *ProtoLSDParser) RpcParameterModifier() (localctx IRpcParameterModifierContext) {
	localctx = NewRpcParameterModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProtoLSDParserRULE_rpcParameterModifier)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProtoLSDParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(ProtoLSDParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProtoLSDParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(ProtoLSDParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(ProtoLSDParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcReturnContext is an interface to support dynamic dispatch.
type IRpcReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURNS() antlr.TerminalNode
	RpcMethodMessage() IRpcMethodMessageContext

	// IsRpcReturnContext differentiates from other interfaces.
	IsRpcReturnContext()
}

type RpcReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcReturnContext() *RpcReturnContext {
	var p = new(RpcReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcReturn
	return p
}

func InitEmptyRpcReturnContext(p *RpcReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcReturn
}

func (*RpcReturnContext) IsRpcReturnContext() {}

func NewRpcReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcReturnContext {
	var p = new(RpcReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcReturn

	return p
}

func (s *RpcReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcReturnContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRETURNS, 0)
}

func (s *RpcReturnContext) RpcMethodMessage() IRpcMethodMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcMethodMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcMethodMessageContext)
}

func (s *RpcReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcReturn(s)
	}
}

func (s *RpcReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcReturn(s)
	}
}

func (p *ProtoLSDParser) RpcReturn() (localctx IRpcReturnContext) {
	localctx = NewRpcReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProtoLSDParserRULE_rpcReturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(ProtoLSDParserRETURNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.RpcMethodMessage()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcInputContext is an interface to support dynamic dispatch.
type IRpcInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RpcMethodMessage() IRpcMethodMessageContext

	// IsRpcInputContext differentiates from other interfaces.
	IsRpcInputContext()
}

type RpcInputContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcInputContext() *RpcInputContext {
	var p = new(RpcInputContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcInput
	return p
}

func InitEmptyRpcInputContext(p *RpcInputContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcInput
}

func (*RpcInputContext) IsRpcInputContext() {}

func NewRpcInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcInputContext {
	var p = new(RpcInputContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcInput

	return p
}

func (s *RpcInputContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcInputContext) RpcMethodMessage() IRpcMethodMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcMethodMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcMethodMessageContext)
}

func (s *RpcInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcInput(s)
	}
}

func (s *RpcInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcInput(s)
	}
}

func (p *ProtoLSDParser) RpcInput() (localctx IRpcInputContext) {
	localctx = NewRpcInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProtoLSDParserRULE_rpcInput)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.RpcMethodMessage()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcMethodMessageContext is an interface to support dynamic dispatch.
type IRpcMethodMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBJECT() antlr.TerminalNode
	RpcSingleMessage() IRpcSingleMessageContext
	RpcParameters() IRpcParametersContext

	// IsRpcMethodMessageContext differentiates from other interfaces.
	IsRpcMethodMessageContext()
}

type RpcMethodMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcMethodMessageContext() *RpcMethodMessageContext {
	var p = new(RpcMethodMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcMethodMessage
	return p
}

func InitEmptyRpcMethodMessageContext(p *RpcMethodMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcMethodMessage
}

func (*RpcMethodMessageContext) IsRpcMethodMessageContext() {}

func NewRpcMethodMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcMethodMessageContext {
	var p = new(RpcMethodMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcMethodMessage

	return p
}

func (s *RpcMethodMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcMethodMessageContext) OBJECT() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserOBJECT, 0)
}

func (s *RpcMethodMessageContext) RpcSingleMessage() IRpcSingleMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcSingleMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcSingleMessageContext)
}

func (s *RpcMethodMessageContext) RpcParameters() IRpcParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcParametersContext)
}

func (s *RpcMethodMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcMethodMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcMethodMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcMethodMessage(s)
	}
}

func (s *RpcMethodMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcMethodMessage(s)
	}
}

func (p *ProtoLSDParser) RpcMethodMessage() (localctx IRpcMethodMessageContext) {
	localctx = NewRpcMethodMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ProtoLSDParserRULE_rpcMethodMessage)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Match(ProtoLSDParserOBJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.RpcSingleMessage()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.RpcParameters()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcSingleMessageContext is an interface to support dynamic dispatch.
type IRpcSingleMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RpcMessageValue() IRpcMessageValueContext
	RPAREN() antlr.TerminalNode

	// IsRpcSingleMessageContext differentiates from other interfaces.
	IsRpcSingleMessageContext()
}

type RpcSingleMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcSingleMessageContext() *RpcSingleMessageContext {
	var p = new(RpcSingleMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcSingleMessage
	return p
}

func InitEmptyRpcSingleMessageContext(p *RpcSingleMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcSingleMessage
}

func (*RpcSingleMessageContext) IsRpcSingleMessageContext() {}

func NewRpcSingleMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcSingleMessageContext {
	var p = new(RpcSingleMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcSingleMessage

	return p
}

func (s *RpcSingleMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcSingleMessageContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLPAREN, 0)
}

func (s *RpcSingleMessageContext) RpcMessageValue() IRpcMessageValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcMessageValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcMessageValueContext)
}

func (s *RpcSingleMessageContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRPAREN, 0)
}

func (s *RpcSingleMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcSingleMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcSingleMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcSingleMessage(s)
	}
}

func (s *RpcSingleMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcSingleMessage(s)
	}
}

func (p *ProtoLSDParser) RpcSingleMessage() (localctx IRpcSingleMessageContext) {
	localctx = NewRpcSingleMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ProtoLSDParserRULE_rpcSingleMessage)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProtoLSDParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(ProtoLSDParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.RpcMessageValue()
		}
		{
			p.SetState(154)
			p.Match(ProtoLSDParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProtoLSDParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.RpcMessageValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcMessageValueContext is an interface to support dynamic dispatch.
type IRpcMessageValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsRpcMessageValueContext differentiates from other interfaces.
	IsRpcMessageValueContext()
}

type RpcMessageValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcMessageValueContext() *RpcMessageValueContext {
	var p = new(RpcMessageValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcMessageValue
	return p
}

func InitEmptyRpcMessageValueContext(p *RpcMessageValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_rpcMessageValue
}

func (*RpcMessageValueContext) IsRpcMessageValueContext() {}

func NewRpcMessageValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcMessageValueContext {
	var p = new(RpcMessageValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_rpcMessageValue

	return p
}

func (s *RpcMessageValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcMessageValueContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserIDENTIFIER)
}

func (s *RpcMessageValueContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, i)
}

func (s *RpcMessageValueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserDOT)
}

func (s *RpcMessageValueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserDOT, i)
}

func (s *RpcMessageValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcMessageValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcMessageValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterRpcMessageValue(s)
	}
}

func (s *RpcMessageValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitRpcMessageValue(s)
	}
}

func (p *ProtoLSDParser) RpcMessageValue() (localctx IRpcMessageValueContext) {
	localctx = NewRpcMessageValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ProtoLSDParserRULE_rpcMessageValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProtoLSDParserDOT {
		{
			p.SetState(160)
			p.Match(ProtoLSDParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(ProtoLSDParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumDefinitionContext is an interface to support dynamic dispatch.
type IEnumDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	EnumBody() IEnumBodyContext
	RBRACE() antlr.TerminalNode
	AllAnnotationDirective() []IAnnotationDirectiveContext
	AnnotationDirective(i int) IAnnotationDirectiveContext

	// IsEnumDefinitionContext differentiates from other interfaces.
	IsEnumDefinitionContext()
}

type EnumDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefinitionContext() *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumDefinition
	return p
}

func InitEmptyEnumDefinitionContext(p *EnumDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumDefinition
}

func (*EnumDefinitionContext) IsEnumDefinitionContext() {}

func NewEnumDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_enumDefinition

	return p
}

func (s *EnumDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserENUM, 0)
}

func (s *EnumDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *EnumDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLBRACE, 0)
}

func (s *EnumDefinitionContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRBRACE, 0)
}

func (s *EnumDefinitionContext) AllAnnotationDirective() []IAnnotationDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationDirectiveContext); ok {
			tst[i] = t.(IAnnotationDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *EnumDefinitionContext) AnnotationDirective(i int) IAnnotationDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationDirectiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationDirectiveContext)
}

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitEnumDefinition(s)
	}
}

func (p *ProtoLSDParser) EnumDefinition() (localctx IEnumDefinitionContext) {
	localctx = NewEnumDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ProtoLSDParserRULE_enumDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProtoLSDParserANNOTATION {
		{
			p.SetState(167)
			p.AnnotationDirective()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.Match(ProtoLSDParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(ProtoLSDParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.EnumBody()
	}
	{
		p.SetState(177)
		p.Match(ProtoLSDParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEnumField() []IEnumFieldContext
	EnumField(i int) IEnumFieldContext
	AllSemi() []ISemiContext
	Semi(i int) ISemiContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllEnumField() []IEnumFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumFieldContext); ok {
			len++
		}
	}

	tst := make([]IEnumFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumFieldContext); ok {
			tst[i] = t.(IEnumFieldContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumField(i int) IEnumFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumBodyContext) AllSemi() []ISemiContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISemiContext); ok {
			len++
		}
	}

	tst := make([]ISemiContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISemiContext); ok {
			tst[i] = t.(ISemiContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) Semi(i int) ISemiContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *ProtoLSDParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ProtoLSDParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.EnumField()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProtoLSDParserIDENTIFIER || _la == ProtoLSDParserSEMI {
		{
			p.SetState(180)
			p.Semi()
		}
		{
			p.SetState(181)
			p.EnumField()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *EnumFieldContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserEQUAL, 0)
}

func (s *EnumFieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserNUMBER, 0)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *ProtoLSDParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ProtoLSDParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserEQUAL {
		{
			p.SetState(189)
			p.Match(ProtoLSDParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(ProtoLSDParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageDefinitionContext is an interface to support dynamic dispatch.
type IMessageDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MESSAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	MessageBody() IMessageBodyContext
	RBRACE() antlr.TerminalNode

	// IsMessageDefinitionContext differentiates from other interfaces.
	IsMessageDefinitionContext()
}

type MessageDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefinitionContext() *MessageDefinitionContext {
	var p = new(MessageDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageDefinition
	return p
}

func InitEmptyMessageDefinitionContext(p *MessageDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageDefinition
}

func (*MessageDefinitionContext) IsMessageDefinitionContext() {}

func NewMessageDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefinitionContext {
	var p = new(MessageDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_messageDefinition

	return p
}

func (s *MessageDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefinitionContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserMESSAGE, 0)
}

func (s *MessageDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *MessageDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLBRACE, 0)
}

func (s *MessageDefinitionContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRBRACE, 0)
}

func (s *MessageDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterMessageDefinition(s)
	}
}

func (s *MessageDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitMessageDefinition(s)
	}
}

func (p *ProtoLSDParser) MessageDefinition() (localctx IMessageDefinitionContext) {
	localctx = NewMessageDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ProtoLSDParserRULE_messageDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(ProtoLSDParserMESSAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(ProtoLSDParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.MessageBody()
	}
	{
		p.SetState(197)
		p.Match(ProtoLSDParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMessageBodyStatement() []IMessageBodyStatementContext
	MessageBodyStatement(i int) IMessageBodyStatementContext

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageBody
	return p
}

func InitEmptyMessageBodyContext(p *MessageBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageBody
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) AllMessageBodyStatement() []IMessageBodyStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageBodyStatementContext); ok {
			len++
		}
	}

	tst := make([]IMessageBodyStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageBodyStatementContext); ok {
			tst[i] = t.(IMessageBodyStatementContext)
			i++
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageBodyStatement(i int) IMessageBodyStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyStatementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *ProtoLSDParser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ProtoLSDParserRULE_messageBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17877052) != 0 {
		{
			p.SetState(199)
			p.MessageBodyStatement()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageBodyStatementContext is an interface to support dynamic dispatch.
type IMessageBodyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageField() IMessageFieldContext
	Semi() ISemiContext
	MessageDefinition() IMessageDefinitionContext

	// IsMessageBodyStatementContext differentiates from other interfaces.
	IsMessageBodyStatementContext()
}

type MessageBodyStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyStatementContext() *MessageBodyStatementContext {
	var p = new(MessageBodyStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageBodyStatement
	return p
}

func InitEmptyMessageBodyStatementContext(p *MessageBodyStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageBodyStatement
}

func (*MessageBodyStatementContext) IsMessageBodyStatementContext() {}

func NewMessageBodyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyStatementContext {
	var p = new(MessageBodyStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_messageBodyStatement

	return p
}

func (s *MessageBodyStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyStatementContext) MessageField() IMessageFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageFieldContext)
}

func (s *MessageBodyStatementContext) Semi() ISemiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiContext)
}

func (s *MessageBodyStatementContext) MessageDefinition() IMessageDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefinitionContext)
}

func (s *MessageBodyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterMessageBodyStatement(s)
	}
}

func (s *MessageBodyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitMessageBodyStatement(s)
	}
}

func (p *ProtoLSDParser) MessageBodyStatement() (localctx IMessageBodyStatementContext) {
	localctx = NewMessageBodyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ProtoLSDParserRULE_messageBodyStatement)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProtoLSDParserT__1, ProtoLSDParserT__2, ProtoLSDParserT__3, ProtoLSDParserT__4, ProtoLSDParserOPTIONAL, ProtoLSDParserREQUIRED, ProtoLSDParserIDENTIFIER, ProtoLSDParserMAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.MessageField()
		}
		{
			p.SetState(206)
			p.Semi()
		}

	case ProtoLSDParserMESSAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.MessageDefinition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageFieldContext is an interface to support dynamic dispatch.
type IMessageFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataType() IDataTypeContext
	IDENTIFIER() antlr.TerminalNode
	OptionalModifier() IOptionalModifierContext
	EQUAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BANG() antlr.TerminalNode

	// IsMessageFieldContext differentiates from other interfaces.
	IsMessageFieldContext()
}

type MessageFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageFieldContext() *MessageFieldContext {
	var p = new(MessageFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageField
	return p
}

func InitEmptyMessageFieldContext(p *MessageFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_messageField
}

func (*MessageFieldContext) IsMessageFieldContext() {}

func NewMessageFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageFieldContext {
	var p = new(MessageFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_messageField

	return p
}

func (s *MessageFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageFieldContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *MessageFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *MessageFieldContext) OptionalModifier() IOptionalModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionalModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionalModifierContext)
}

func (s *MessageFieldContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserEQUAL, 0)
}

func (s *MessageFieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserNUMBER, 0)
}

func (s *MessageFieldContext) BANG() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserBANG, 0)
}

func (s *MessageFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterMessageField(s)
	}
}

func (s *MessageFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitMessageField(s)
	}
}

func (p *ProtoLSDParser) MessageField() (localctx IMessageFieldContext) {
	localctx = NewMessageFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ProtoLSDParserRULE_messageField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserOPTIONAL || _la == ProtoLSDParserREQUIRED {
		{
			p.SetState(211)
			p.OptionalModifier()
		}

	}
	{
		p.SetState(214)
		p.DataType()
	}
	{
		p.SetState(215)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserEQUAL {
		{
			p.SetState(216)
			p.Match(ProtoLSDParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(ProtoLSDParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ProtoLSDParserBANG {
			{
				p.SetState(218)
				p.Match(ProtoLSDParserBANG)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPreprocessorParametersContext is an interface to support dynamic dispatch.
type IPreprocessorParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllRpcMessageValue() []IRpcMessageValueContext
	RpcMessageValue(i int) IRpcMessageValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPreprocessorParametersContext differentiates from other interfaces.
	IsPreprocessorParametersContext()
}

type PreprocessorParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreprocessorParametersContext() *PreprocessorParametersContext {
	var p = new(PreprocessorParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_preprocessorParameters
	return p
}

func InitEmptyPreprocessorParametersContext(p *PreprocessorParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_preprocessorParameters
}

func (*PreprocessorParametersContext) IsPreprocessorParametersContext() {}

func NewPreprocessorParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreprocessorParametersContext {
	var p = new(PreprocessorParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_preprocessorParameters

	return p
}

func (s *PreprocessorParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *PreprocessorParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserLPAREN, 0)
}

func (s *PreprocessorParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserRPAREN, 0)
}

func (s *PreprocessorParametersContext) AllRpcMessageValue() []IRpcMessageValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRpcMessageValueContext); ok {
			len++
		}
	}

	tst := make([]IRpcMessageValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRpcMessageValueContext); ok {
			tst[i] = t.(IRpcMessageValueContext)
			i++
		}
	}

	return tst
}

func (s *PreprocessorParametersContext) RpcMessageValue(i int) IRpcMessageValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcMessageValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcMessageValueContext)
}

func (s *PreprocessorParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserCOMMA)
}

func (s *PreprocessorParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserCOMMA, i)
}

func (s *PreprocessorParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreprocessorParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterPreprocessorParameters(s)
	}
}

func (s *PreprocessorParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitPreprocessorParameters(s)
	}
}

func (p *ProtoLSDParser) PreprocessorParameters() (localctx IPreprocessorParametersContext) {
	localctx = NewPreprocessorParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ProtoLSDParserRULE_preprocessorParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(ProtoLSDParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserIDENTIFIER {
		{
			p.SetState(224)
			p.RpcMessageValue()
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ProtoLSDParserCOMMA {
			{
				p.SetState(225)
				p.Match(ProtoLSDParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(226)
				p.RpcMessageValue()
			}

			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(234)
		p.Match(ProtoLSDParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPreprocessorDirectiveContext is an interface to support dynamic dispatch.
type IPreprocessorDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	PREPROCESSOR_NAME() antlr.TerminalNode
	PreprocessorParameters() IPreprocessorParametersContext

	// IsPreprocessorDirectiveContext differentiates from other interfaces.
	IsPreprocessorDirectiveContext()
}

type PreprocessorDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreprocessorDirectiveContext() *PreprocessorDirectiveContext {
	var p = new(PreprocessorDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_preprocessorDirective
	return p
}

func InitEmptyPreprocessorDirectiveContext(p *PreprocessorDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_preprocessorDirective
}

func (*PreprocessorDirectiveContext) IsPreprocessorDirectiveContext() {}

func NewPreprocessorDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreprocessorDirectiveContext {
	var p = new(PreprocessorDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_preprocessorDirective

	return p
}

func (s *PreprocessorDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PreprocessorDirectiveContext) HASH() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserHASH, 0)
}

func (s *PreprocessorDirectiveContext) PREPROCESSOR_NAME() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserPREPROCESSOR_NAME, 0)
}

func (s *PreprocessorDirectiveContext) PreprocessorParameters() IPreprocessorParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreprocessorParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreprocessorParametersContext)
}

func (s *PreprocessorDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreprocessorDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterPreprocessorDirective(s)
	}
}

func (s *PreprocessorDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitPreprocessorDirective(s)
	}
}

func (p *ProtoLSDParser) PreprocessorDirective() (localctx IPreprocessorDirectiveContext) {
	localctx = NewPreprocessorDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ProtoLSDParserRULE_preprocessorDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(ProtoLSDParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(ProtoLSDParserPREPROCESSOR_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserLPAREN {
		{
			p.SetState(238)
			p.PreprocessorParameters()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationDirectiveContext is an interface to support dynamic dispatch.
type IAnnotationDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ANNOTATION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsAnnotationDirectiveContext differentiates from other interfaces.
	IsAnnotationDirectiveContext()
}

type AnnotationDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationDirectiveContext() *AnnotationDirectiveContext {
	var p = new(AnnotationDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_annotationDirective
	return p
}

func InitEmptyAnnotationDirectiveContext(p *AnnotationDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_annotationDirective
}

func (*AnnotationDirectiveContext) IsAnnotationDirectiveContext() {}

func NewAnnotationDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationDirectiveContext {
	var p = new(AnnotationDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_annotationDirective

	return p
}

func (s *AnnotationDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationDirectiveContext) ANNOTATION() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserANNOTATION, 0)
}

func (s *AnnotationDirectiveContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *AnnotationDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterAnnotationDirective(s)
	}
}

func (s *AnnotationDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitAnnotationDirective(s)
	}
}

func (p *ProtoLSDParser) AnnotationDirective() (localctx IAnnotationDirectiveContext) {
	localctx = NewAnnotationDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ProtoLSDParserRULE_annotationDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(ProtoLSDParserANNOTATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionalModifierContext is an interface to support dynamic dispatch.
type IOptionalModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTIONAL() antlr.TerminalNode
	REQUIRED() antlr.TerminalNode

	// IsOptionalModifierContext differentiates from other interfaces.
	IsOptionalModifierContext()
}

type OptionalModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalModifierContext() *OptionalModifierContext {
	var p = new(OptionalModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_optionalModifier
	return p
}

func InitEmptyOptionalModifierContext(p *OptionalModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_optionalModifier
}

func (*OptionalModifierContext) IsOptionalModifierContext() {}

func NewOptionalModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalModifierContext {
	var p = new(OptionalModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_optionalModifier

	return p
}

func (s *OptionalModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalModifierContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserOPTIONAL, 0)
}

func (s *OptionalModifierContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserREQUIRED, 0)
}

func (s *OptionalModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterOptionalModifier(s)
	}
}

func (s *OptionalModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitOptionalModifier(s)
	}
}

func (p *ProtoLSDParser) OptionalModifier() (localctx IOptionalModifierContext) {
	localctx = NewOptionalModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ProtoLSDParserRULE_optionalModifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProtoLSDParserOPTIONAL || _la == ProtoLSDParserREQUIRED) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Semi() ISemiContext

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserOPTION, 0)
}

func (s *OptionStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *OptionStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserEQUAL, 0)
}

func (s *OptionStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserSTRING, 0)
}

func (s *OptionStatementContext) Semi() ISemiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiContext)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterOptionStatement(s)
	}
}

func (s *OptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitOptionStatement(s)
	}
}

func (p *ProtoLSDParser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ProtoLSDParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(ProtoLSDParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(ProtoLSDParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(ProtoLSDParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Semi()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	Semi() ISemiContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserPACKAGE, 0)
}

func (s *PackageStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserIDENTIFIER)
}

func (s *PackageStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, i)
}

func (s *PackageStatementContext) Semi() ISemiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiContext)
}

func (s *PackageStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ProtoLSDParserDOT)
}

func (s *PackageStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserDOT, i)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *ProtoLSDParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ProtoLSDParserRULE_packageStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(ProtoLSDParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(ProtoLSDParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProtoLSDParserDOT {
		{
			p.SetState(254)
			p.Match(ProtoLSDParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(ProtoLSDParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)
		p.Semi()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataTypeText() IDataTypeTextContext
	ARRAY() antlr.TerminalNode
	QUESTION() antlr.TerminalNode

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_dataType
	return p
}

func InitEmptyDataTypeContext(p *DataTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_dataType
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) DataTypeText() IDataTypeTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeTextContext)
}

func (s *DataTypeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserARRAY, 0)
}

func (s *DataTypeContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserQUESTION, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (p *ProtoLSDParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ProtoLSDParserRULE_dataType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.DataTypeText()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserARRAY {
		{
			p.SetState(264)
			p.Match(ProtoLSDParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ProtoLSDParserQUESTION {
		{
			p.SetState(267)
			p.Match(ProtoLSDParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataTypeTextContext is an interface to support dynamic dispatch.
type IDataTypeTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RpcMessageValue() IRpcMessageValueContext

	// IsDataTypeTextContext differentiates from other interfaces.
	IsDataTypeTextContext()
}

type DataTypeTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeTextContext() *DataTypeTextContext {
	var p = new(DataTypeTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_dataTypeText
	return p
}

func InitEmptyDataTypeTextContext(p *DataTypeTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProtoLSDParserRULE_dataTypeText
}

func (*DataTypeTextContext) IsDataTypeTextContext() {}

func NewDataTypeTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeTextContext {
	var p = new(DataTypeTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtoLSDParserRULE_dataTypeText

	return p
}

func (s *DataTypeTextContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeTextContext) MAP() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserMAP, 0)
}

func (s *DataTypeTextContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtoLSDParserIDENTIFIER, 0)
}

func (s *DataTypeTextContext) RpcMessageValue() IRpcMessageValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcMessageValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcMessageValueContext)
}

func (s *DataTypeTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.EnterDataTypeText(s)
	}
}

func (s *DataTypeTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtoLSDListener); ok {
		listenerT.ExitDataTypeText(s)
	}
}

func (p *ProtoLSDParser) DataTypeText() (localctx IDataTypeTextContext) {
	localctx = NewDataTypeTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ProtoLSDParserRULE_dataTypeText)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Match(ProtoLSDParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(ProtoLSDParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.Match(ProtoLSDParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(273)
			p.Match(ProtoLSDParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(274)
			p.Match(ProtoLSDParserMAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(275)
			p.Match(ProtoLSDParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(276)
			p.RpcMessageValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
