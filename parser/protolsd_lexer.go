// Code generated from ProtoLSD.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ProtoLSDLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ProtoLSDLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protolsdlexerLexerInit() {
	staticData := &ProtoLSDLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "IMPORT", "PRIVATE", "SERVICE",
		"RPC", "ENUM", "MESSAGE", "RETURNS", "OPTION", "OPTIONAL", "REQUIRED",
		"REPEATED", "PACKAGE", "LINE_COMMENT", "BLOCK_COMMENT", "IDENTIFIER",
		"PREPROCESSOR_NAME", "STRING", "NUMBER", "MAP", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "SEMI", "COMMA", "EQUAL",
		"DOT", "HASH", "QUESTION", "ARRAY", "OBJECT", "BANG", "PLUS", "ANNOTATION",
		"WS", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 309, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 204, 8, 17, 10,
		17, 12, 17, 207, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18,
		215, 8, 18, 10, 18, 12, 18, 218, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 5, 19, 227, 8, 19, 10, 19, 12, 19, 230, 9, 19, 1, 20,
		1, 20, 5, 20, 234, 8, 20, 10, 20, 12, 20, 237, 9, 20, 1, 21, 1, 21, 1,
		21, 5, 21, 242, 8, 21, 10, 21, 12, 21, 245, 9, 21, 1, 21, 1, 21, 1, 22,
		4, 22, 250, 8, 22, 11, 22, 12, 22, 251, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		41, 4, 41, 301, 8, 41, 11, 41, 12, 41, 302, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 216, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 0, 1, 0, 8,
		2, 0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2,
		0, 34, 34, 92, 92, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 8, 0, 34,
		34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 315,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 87,
		1, 0, 0, 0, 3, 89, 1, 0, 0, 0, 5, 93, 1, 0, 0, 0, 7, 98, 1, 0, 0, 0, 9,
		103, 1, 0, 0, 0, 11, 109, 1, 0, 0, 0, 13, 116, 1, 0, 0, 0, 15, 124, 1,
		0, 0, 0, 17, 132, 1, 0, 0, 0, 19, 136, 1, 0, 0, 0, 21, 141, 1, 0, 0, 0,
		23, 149, 1, 0, 0, 0, 25, 157, 1, 0, 0, 0, 27, 164, 1, 0, 0, 0, 29, 173,
		1, 0, 0, 0, 31, 182, 1, 0, 0, 0, 33, 191, 1, 0, 0, 0, 35, 199, 1, 0, 0,
		0, 37, 210, 1, 0, 0, 0, 39, 224, 1, 0, 0, 0, 41, 231, 1, 0, 0, 0, 43, 238,
		1, 0, 0, 0, 45, 249, 1, 0, 0, 0, 47, 253, 1, 0, 0, 0, 49, 263, 1, 0, 0,
		0, 51, 265, 1, 0, 0, 0, 53, 267, 1, 0, 0, 0, 55, 269, 1, 0, 0, 0, 57, 271,
		1, 0, 0, 0, 59, 273, 1, 0, 0, 0, 61, 275, 1, 0, 0, 0, 63, 277, 1, 0, 0,
		0, 65, 279, 1, 0, 0, 0, 67, 281, 1, 0, 0, 0, 69, 283, 1, 0, 0, 0, 71, 285,
		1, 0, 0, 0, 73, 287, 1, 0, 0, 0, 75, 290, 1, 0, 0, 0, 77, 293, 1, 0, 0,
		0, 79, 295, 1, 0, 0, 0, 81, 297, 1, 0, 0, 0, 83, 300, 1, 0, 0, 0, 85, 306,
		1, 0, 0, 0, 87, 88, 5, 42, 0, 0, 88, 2, 1, 0, 0, 0, 89, 90, 5, 105, 0,
		0, 90, 91, 5, 110, 0, 0, 91, 92, 5, 116, 0, 0, 92, 4, 1, 0, 0, 0, 93, 94,
		5, 117, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 97, 5, 116,
		0, 0, 97, 6, 1, 0, 0, 0, 98, 99, 5, 108, 0, 0, 99, 100, 5, 111, 0, 0, 100,
		101, 5, 110, 0, 0, 101, 102, 5, 103, 0, 0, 102, 8, 1, 0, 0, 0, 103, 104,
		5, 117, 0, 0, 104, 105, 5, 108, 0, 0, 105, 106, 5, 111, 0, 0, 106, 107,
		5, 110, 0, 0, 107, 108, 5, 103, 0, 0, 108, 10, 1, 0, 0, 0, 109, 110, 5,
		105, 0, 0, 110, 111, 5, 109, 0, 0, 111, 112, 5, 112, 0, 0, 112, 113, 5,
		111, 0, 0, 113, 114, 5, 114, 0, 0, 114, 115, 5, 116, 0, 0, 115, 12, 1,
		0, 0, 0, 116, 117, 5, 112, 0, 0, 117, 118, 5, 114, 0, 0, 118, 119, 5, 105,
		0, 0, 119, 120, 5, 118, 0, 0, 120, 121, 5, 97, 0, 0, 121, 122, 5, 116,
		0, 0, 122, 123, 5, 101, 0, 0, 123, 14, 1, 0, 0, 0, 124, 125, 5, 115, 0,
		0, 125, 126, 5, 101, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128, 5, 118, 0,
		0, 128, 129, 5, 105, 0, 0, 129, 130, 5, 99, 0, 0, 130, 131, 5, 101, 0,
		0, 131, 16, 1, 0, 0, 0, 132, 133, 5, 114, 0, 0, 133, 134, 5, 112, 0, 0,
		134, 135, 5, 99, 0, 0, 135, 18, 1, 0, 0, 0, 136, 137, 5, 101, 0, 0, 137,
		138, 5, 110, 0, 0, 138, 139, 5, 117, 0, 0, 139, 140, 5, 109, 0, 0, 140,
		20, 1, 0, 0, 0, 141, 142, 5, 109, 0, 0, 142, 143, 5, 101, 0, 0, 143, 144,
		5, 115, 0, 0, 144, 145, 5, 115, 0, 0, 145, 146, 5, 97, 0, 0, 146, 147,
		5, 103, 0, 0, 147, 148, 5, 101, 0, 0, 148, 22, 1, 0, 0, 0, 149, 150, 5,
		114, 0, 0, 150, 151, 5, 101, 0, 0, 151, 152, 5, 116, 0, 0, 152, 153, 5,
		117, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 110, 0, 0, 155, 156, 5,
		115, 0, 0, 156, 24, 1, 0, 0, 0, 157, 158, 5, 111, 0, 0, 158, 159, 5, 112,
		0, 0, 159, 160, 5, 116, 0, 0, 160, 161, 5, 105, 0, 0, 161, 162, 5, 111,
		0, 0, 162, 163, 5, 110, 0, 0, 163, 26, 1, 0, 0, 0, 164, 165, 5, 111, 0,
		0, 165, 166, 5, 112, 0, 0, 166, 167, 5, 116, 0, 0, 167, 168, 5, 105, 0,
		0, 168, 169, 5, 111, 0, 0, 169, 170, 5, 110, 0, 0, 170, 171, 5, 97, 0,
		0, 171, 172, 5, 108, 0, 0, 172, 28, 1, 0, 0, 0, 173, 174, 5, 114, 0, 0,
		174, 175, 5, 101, 0, 0, 175, 176, 5, 113, 0, 0, 176, 177, 5, 117, 0, 0,
		177, 178, 5, 105, 0, 0, 178, 179, 5, 114, 0, 0, 179, 180, 5, 101, 0, 0,
		180, 181, 5, 100, 0, 0, 181, 30, 1, 0, 0, 0, 182, 183, 5, 114, 0, 0, 183,
		184, 5, 101, 0, 0, 184, 185, 5, 112, 0, 0, 185, 186, 5, 101, 0, 0, 186,
		187, 5, 97, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 101, 0, 0, 189,
		190, 5, 100, 0, 0, 190, 32, 1, 0, 0, 0, 191, 192, 5, 112, 0, 0, 192, 193,
		5, 97, 0, 0, 193, 194, 5, 99, 0, 0, 194, 195, 5, 107, 0, 0, 195, 196, 5,
		97, 0, 0, 196, 197, 5, 103, 0, 0, 197, 198, 5, 101, 0, 0, 198, 34, 1, 0,
		0, 0, 199, 200, 5, 47, 0, 0, 200, 201, 5, 47, 0, 0, 201, 205, 1, 0, 0,
		0, 202, 204, 8, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205,
		203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 1, 0, 0, 0, 207, 205,
		1, 0, 0, 0, 208, 209, 6, 17, 0, 0, 209, 36, 1, 0, 0, 0, 210, 211, 5, 47,
		0, 0, 211, 212, 5, 42, 0, 0, 212, 216, 1, 0, 0, 0, 213, 215, 9, 0, 0, 0,
		214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220,
		5, 42, 0, 0, 220, 221, 5, 47, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 6,
		18, 0, 0, 223, 38, 1, 0, 0, 0, 224, 228, 7, 1, 0, 0, 225, 227, 7, 2, 0,
		0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 40, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 235, 7,
		1, 0, 0, 232, 234, 7, 3, 0, 0, 233, 232, 1, 0, 0, 0, 234, 237, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 42, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 238, 243, 5, 34, 0, 0, 239, 242, 3, 85, 42, 0, 240, 242,
		8, 4, 0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 246, 1, 0, 0, 0,
		245, 243, 1, 0, 0, 0, 246, 247, 5, 34, 0, 0, 247, 44, 1, 0, 0, 0, 248,
		250, 7, 5, 0, 0, 249, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 46, 1, 0, 0, 0, 253, 254, 5, 109,
		0, 0, 254, 255, 5, 97, 0, 0, 255, 256, 5, 112, 0, 0, 256, 257, 5, 60, 0,
		0, 257, 258, 1, 0, 0, 0, 258, 259, 3, 39, 19, 0, 259, 260, 5, 44, 0, 0,
		260, 261, 3, 39, 19, 0, 261, 262, 5, 62, 0, 0, 262, 48, 1, 0, 0, 0, 263,
		264, 5, 40, 0, 0, 264, 50, 1, 0, 0, 0, 265, 266, 5, 41, 0, 0, 266, 52,
		1, 0, 0, 0, 267, 268, 5, 123, 0, 0, 268, 54, 1, 0, 0, 0, 269, 270, 5, 125,
		0, 0, 270, 56, 1, 0, 0, 0, 271, 272, 5, 91, 0, 0, 272, 58, 1, 0, 0, 0,
		273, 274, 5, 93, 0, 0, 274, 60, 1, 0, 0, 0, 275, 276, 5, 59, 0, 0, 276,
		62, 1, 0, 0, 0, 277, 278, 5, 44, 0, 0, 278, 64, 1, 0, 0, 0, 279, 280, 5,
		61, 0, 0, 280, 66, 1, 0, 0, 0, 281, 282, 5, 46, 0, 0, 282, 68, 1, 0, 0,
		0, 283, 284, 5, 35, 0, 0, 284, 70, 1, 0, 0, 0, 285, 286, 5, 63, 0, 0, 286,
		72, 1, 0, 0, 0, 287, 288, 5, 91, 0, 0, 288, 289, 5, 93, 0, 0, 289, 74,
		1, 0, 0, 0, 290, 291, 5, 123, 0, 0, 291, 292, 5, 125, 0, 0, 292, 76, 1,
		0, 0, 0, 293, 294, 5, 33, 0, 0, 294, 78, 1, 0, 0, 0, 295, 296, 5, 43, 0,
		0, 296, 80, 1, 0, 0, 0, 297, 298, 5, 64, 0, 0, 298, 82, 1, 0, 0, 0, 299,
		301, 7, 6, 0, 0, 300, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 300,
		1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 6, 41,
		0, 0, 305, 84, 1, 0, 0, 0, 306, 307, 5, 92, 0, 0, 307, 308, 7, 7, 0, 0,
		308, 86, 1, 0, 0, 0, 9, 0, 205, 216, 228, 235, 241, 243, 251, 302, 1, 6,
		0, 0,
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

// ProtoLSDLexerInit initializes any static state used to implement ProtoLSDLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewProtoLSDLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProtoLSDLexerInit() {
	staticData := &ProtoLSDLexerLexerStaticData
	staticData.once.Do(protolsdlexerLexerInit)
}

// NewProtoLSDLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewProtoLSDLexer(input antlr.CharStream) *ProtoLSDLexer {
	ProtoLSDLexerInit()
	l := new(ProtoLSDLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ProtoLSDLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ProtoLSD.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ProtoLSDLexer tokens.
const (
	ProtoLSDLexerT__0              = 1
	ProtoLSDLexerT__1              = 2
	ProtoLSDLexerT__2              = 3
	ProtoLSDLexerT__3              = 4
	ProtoLSDLexerT__4              = 5
	ProtoLSDLexerIMPORT            = 6
	ProtoLSDLexerPRIVATE           = 7
	ProtoLSDLexerSERVICE           = 8
	ProtoLSDLexerRPC               = 9
	ProtoLSDLexerENUM              = 10
	ProtoLSDLexerMESSAGE           = 11
	ProtoLSDLexerRETURNS           = 12
	ProtoLSDLexerOPTION            = 13
	ProtoLSDLexerOPTIONAL          = 14
	ProtoLSDLexerREQUIRED          = 15
	ProtoLSDLexerREPEATED          = 16
	ProtoLSDLexerPACKAGE           = 17
	ProtoLSDLexerLINE_COMMENT      = 18
	ProtoLSDLexerBLOCK_COMMENT     = 19
	ProtoLSDLexerIDENTIFIER        = 20
	ProtoLSDLexerPREPROCESSOR_NAME = 21
	ProtoLSDLexerSTRING            = 22
	ProtoLSDLexerNUMBER            = 23
	ProtoLSDLexerMAP               = 24
	ProtoLSDLexerLPAREN            = 25
	ProtoLSDLexerRPAREN            = 26
	ProtoLSDLexerLBRACE            = 27
	ProtoLSDLexerRBRACE            = 28
	ProtoLSDLexerLBRACKET          = 29
	ProtoLSDLexerRBRACKET          = 30
	ProtoLSDLexerSEMI              = 31
	ProtoLSDLexerCOMMA             = 32
	ProtoLSDLexerEQUAL             = 33
	ProtoLSDLexerDOT               = 34
	ProtoLSDLexerHASH              = 35
	ProtoLSDLexerQUESTION          = 36
	ProtoLSDLexerARRAY             = 37
	ProtoLSDLexerOBJECT            = 38
	ProtoLSDLexerBANG              = 39
	ProtoLSDLexerPLUS              = 40
	ProtoLSDLexerANNOTATION        = 41
	ProtoLSDLexerWS                = 42
)