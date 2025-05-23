// Code generated from /Users/zhangming/gopath/src/github.com/period331/jagep/internal/lang/iantlr/jagelang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // jagelang

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

type jagelangParser struct {
	*antlr.BaseParser
}

var JagelangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jagelangParserInit() {
	staticData := &JagelangParserStaticData
	staticData.LiteralNames = []string{
		"", "'START'", "'('", "','", "')'", "'{'", "'}'", "'FRAGMENT'", "'REGISTER'",
		"'ON_FINISH'", "'->'", "'['", "']'", "'UNFOLD'", "'GO'", "'WAIT'", "'NO_CHECK_MISS'",
		"'SKIP'", "'|'", "'SWITCH'", "'CASE'", "'=>'", "'OPERATOR'", "'='",
		"'@'", "'-'", "'s'", "'ms'", "'\\u00B5s'", "'h'", "'m'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE", "FALSE",
		"ID", "INT", "DQUOTA_STRING", "LINE_COMMENT", "COMMENT", "TERMINATOR",
		"WS",
	}
	staticData.RuleNames = []string{
		"primary", "startDirective", "fragmentDirective", "registerDirective",
		"onFinishStmt", "exedescStmt", "leafSnippet", "operatorStmt", "serialStmt",
		"serialBracketStmt", "concurrentStmt", "unfoldDirective", "goDirective",
		"waitDirective", "noCheckMissDirective", "skipDirective", "wrapStmt",
		"wrapBracketStmt", "switchDirective", "switchCaseDirective", "registerOperatorDirective",
		"argStmt", "argsStmt", "ignoreError", "integerLiteral", "stringLiteral",
		"boolLiteral", "durationLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 403, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 61, 8, 0, 10, 0, 12,
		0, 64, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 71, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 77, 8, 1, 10, 1, 12, 1, 80, 9, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 99, 8, 3, 10, 3, 12, 3, 102, 9, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 120, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 129,
		8, 6, 1, 7, 3, 7, 132, 8, 7, 1, 7, 1, 7, 3, 7, 136, 8, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 142, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 149, 8, 7,
		1, 7, 3, 7, 152, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 159, 8, 7, 10,
		7, 12, 7, 162, 9, 7, 1, 7, 1, 7, 3, 7, 166, 8, 7, 1, 7, 3, 7, 169, 8, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 178, 8, 7, 10, 7, 12, 7,
		181, 9, 7, 1, 7, 1, 7, 3, 7, 185, 8, 7, 3, 7, 187, 8, 7, 1, 8, 1, 8, 3,
		8, 191, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 196, 8, 8, 4, 8, 198, 8, 8, 11, 8,
		12, 8, 199, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		210, 8, 10, 10, 10, 12, 10, 213, 9, 10, 1, 10, 3, 10, 216, 8, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 3, 13, 233, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 240, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 256, 8, 16,
		10, 16, 12, 16, 259, 9, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 276,
		8, 18, 10, 18, 12, 18, 279, 9, 18, 1, 18, 3, 18, 282, 8, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 307, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 333, 8, 21, 10, 21,
		12, 21, 336, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 5, 21, 346, 8, 21, 10, 21, 12, 21, 349, 9, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 359, 8, 21, 10, 21, 12, 21, 362,
		9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 372,
		8, 21, 10, 21, 12, 21, 375, 9, 21, 1, 21, 1, 21, 3, 21, 379, 8, 21, 1,
		22, 1, 22, 1, 22, 5, 22, 384, 8, 22, 10, 22, 12, 22, 387, 9, 22, 1, 23,
		1, 23, 1, 24, 3, 24, 392, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 0, 0, 28, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 0, 2, 1, 0, 31, 32, 1, 0, 26, 30, 433, 0, 62, 1, 0, 0, 0, 2, 65, 1,
		0, 0, 0, 4, 84, 1, 0, 0, 0, 6, 92, 1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10,
		119, 1, 0, 0, 0, 12, 128, 1, 0, 0, 0, 14, 186, 1, 0, 0, 0, 16, 190, 1,
		0, 0, 0, 18, 201, 1, 0, 0, 0, 20, 205, 1, 0, 0, 0, 22, 219, 1, 0, 0, 0,
		24, 224, 1, 0, 0, 0, 26, 232, 1, 0, 0, 0, 28, 243, 1, 0, 0, 0, 30, 247,
		1, 0, 0, 0, 32, 252, 1, 0, 0, 0, 34, 263, 1, 0, 0, 0, 36, 267, 1, 0, 0,
		0, 38, 285, 1, 0, 0, 0, 40, 290, 1, 0, 0, 0, 42, 378, 1, 0, 0, 0, 44, 380,
		1, 0, 0, 0, 46, 388, 1, 0, 0, 0, 48, 391, 1, 0, 0, 0, 50, 395, 1, 0, 0,
		0, 52, 397, 1, 0, 0, 0, 54, 399, 1, 0, 0, 0, 56, 61, 3, 2, 1, 0, 57, 61,
		3, 4, 2, 0, 58, 61, 3, 6, 3, 0, 59, 61, 3, 6, 3, 0, 60, 56, 1, 0, 0, 0,
		60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 64, 1,
		0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 1, 1, 0, 0, 0, 64,
		62, 1, 0, 0, 0, 65, 66, 5, 1, 0, 0, 66, 67, 5, 2, 0, 0, 67, 70, 3, 50,
		25, 0, 68, 69, 5, 3, 0, 0, 69, 71, 3, 44, 22, 0, 70, 68, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 5, 4, 0, 0, 73, 78, 5, 5, 0,
		0, 74, 77, 3, 8, 4, 0, 75, 77, 3, 28, 14, 0, 76, 74, 1, 0, 0, 0, 76, 75,
		1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0,
		79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 3, 10, 5, 0, 82, 83, 5,
		6, 0, 0, 83, 3, 1, 0, 0, 0, 84, 85, 5, 7, 0, 0, 85, 86, 5, 2, 0, 0, 86,
		87, 3, 50, 25, 0, 87, 88, 5, 4, 0, 0, 88, 89, 5, 5, 0, 0, 89, 90, 3, 10,
		5, 0, 90, 91, 5, 6, 0, 0, 91, 5, 1, 0, 0, 0, 92, 93, 5, 8, 0, 0, 93, 94,
		5, 2, 0, 0, 94, 95, 3, 50, 25, 0, 95, 96, 5, 4, 0, 0, 96, 100, 5, 5, 0,
		0, 97, 99, 3, 40, 20, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1,
		0, 0, 0, 103, 104, 5, 6, 0, 0, 104, 7, 1, 0, 0, 0, 105, 106, 5, 9, 0, 0,
		106, 107, 5, 2, 0, 0, 107, 108, 5, 4, 0, 0, 108, 109, 5, 5, 0, 0, 109,
		110, 3, 10, 5, 0, 110, 111, 5, 6, 0, 0, 111, 9, 1, 0, 0, 0, 112, 120, 3,
		16, 8, 0, 113, 120, 3, 20, 10, 0, 114, 120, 3, 32, 16, 0, 115, 120, 3,
		34, 17, 0, 116, 120, 3, 14, 7, 0, 117, 120, 3, 36, 18, 0, 118, 120, 3,
		22, 11, 0, 119, 112, 1, 0, 0, 0, 119, 113, 1, 0, 0, 0, 119, 114, 1, 0,
		0, 0, 119, 115, 1, 0, 0, 0, 119, 116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0,
		119, 118, 1, 0, 0, 0, 120, 11, 1, 0, 0, 0, 121, 129, 3, 14, 7, 0, 122,
		129, 3, 20, 10, 0, 123, 129, 3, 18, 9, 0, 124, 129, 3, 22, 11, 0, 125,
		129, 3, 34, 17, 0, 126, 129, 3, 24, 12, 0, 127, 129, 3, 36, 18, 0, 128,
		121, 1, 0, 0, 0, 128, 122, 1, 0, 0, 0, 128, 123, 1, 0, 0, 0, 128, 124,
		1, 0, 0, 0, 128, 125, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 127, 1, 0,
		0, 0, 129, 13, 1, 0, 0, 0, 130, 132, 3, 46, 23, 0, 131, 130, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 187, 5, 33, 0, 0, 134,
		136, 3, 46, 23, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137,
		1, 0, 0, 0, 137, 138, 5, 33, 0, 0, 138, 139, 5, 2, 0, 0, 139, 187, 5, 4,
		0, 0, 140, 142, 3, 46, 23, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0,
		0, 142, 143, 1, 0, 0, 0, 143, 148, 5, 33, 0, 0, 144, 145, 5, 2, 0, 0, 145,
		146, 3, 44, 22, 0, 146, 147, 5, 4, 0, 0, 147, 149, 1, 0, 0, 0, 148, 144,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 187, 1, 0, 0, 0, 150, 152, 3, 46,
		23, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 165, 5, 33, 0, 0, 154, 155, 5, 2, 0, 0, 155, 160, 3, 26, 13, 0, 156,
		157, 5, 3, 0, 0, 157, 159, 3, 26, 13, 0, 158, 156, 1, 0, 0, 0, 159, 162,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0,
		0, 0, 162, 160, 1, 0, 0, 0, 163, 164, 5, 4, 0, 0, 164, 166, 1, 0, 0, 0,
		165, 154, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 187, 1, 0, 0, 0, 167,
		169, 3, 46, 23, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 184, 5, 33, 0, 0, 171, 172, 5, 2, 0, 0, 172, 173, 3, 44,
		22, 0, 173, 174, 5, 3, 0, 0, 174, 179, 3, 26, 13, 0, 175, 176, 5, 3, 0,
		0, 176, 178, 3, 26, 13, 0, 177, 175, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 183, 5, 4, 0, 0, 183, 185, 1, 0, 0, 0, 184, 171,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 131, 1, 0,
		0, 0, 186, 135, 1, 0, 0, 0, 186, 141, 1, 0, 0, 0, 186, 151, 1, 0, 0, 0,
		186, 168, 1, 0, 0, 0, 187, 15, 1, 0, 0, 0, 188, 191, 3, 12, 6, 0, 189,
		191, 3, 30, 15, 0, 190, 188, 1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191, 197,
		1, 0, 0, 0, 192, 195, 5, 10, 0, 0, 193, 196, 3, 12, 6, 0, 194, 196, 3,
		30, 15, 0, 195, 193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 198, 1, 0,
		0, 0, 197, 192, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 17, 1, 0, 0, 0, 201, 202, 5, 2, 0, 0, 202, 203,
		3, 16, 8, 0, 203, 204, 5, 4, 0, 0, 204, 19, 1, 0, 0, 0, 205, 206, 5, 11,
		0, 0, 206, 211, 3, 12, 6, 0, 207, 208, 5, 3, 0, 0, 208, 210, 3, 12, 6,
		0, 209, 207, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 216,
		5, 3, 0, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0,
		0, 0, 217, 218, 5, 12, 0, 0, 218, 21, 1, 0, 0, 0, 219, 220, 5, 13, 0, 0,
		220, 221, 5, 2, 0, 0, 221, 222, 3, 50, 25, 0, 222, 223, 5, 4, 0, 0, 223,
		23, 1, 0, 0, 0, 224, 225, 5, 14, 0, 0, 225, 226, 5, 2, 0, 0, 226, 227,
		3, 10, 5, 0, 227, 228, 5, 3, 0, 0, 228, 229, 3, 50, 25, 0, 229, 230, 5,
		4, 0, 0, 230, 25, 1, 0, 0, 0, 231, 233, 3, 46, 23, 0, 232, 231, 1, 0, 0,
		0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 5, 15, 0, 0, 235,
		236, 5, 2, 0, 0, 236, 239, 3, 50, 25, 0, 237, 238, 5, 3, 0, 0, 238, 240,
		3, 44, 22, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1,
		0, 0, 0, 241, 242, 5, 4, 0, 0, 242, 27, 1, 0, 0, 0, 243, 244, 5, 16, 0,
		0, 244, 245, 5, 2, 0, 0, 245, 246, 5, 4, 0, 0, 246, 29, 1, 0, 0, 0, 247,
		248, 5, 17, 0, 0, 248, 249, 5, 2, 0, 0, 249, 250, 3, 14, 7, 0, 250, 251,
		5, 4, 0, 0, 251, 31, 1, 0, 0, 0, 252, 257, 3, 14, 7, 0, 253, 254, 5, 18,
		0, 0, 254, 256, 3, 14, 7, 0, 255, 253, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0,
		257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259,
		257, 1, 0, 0, 0, 260, 261, 5, 18, 0, 0, 261, 262, 3, 12, 6, 0, 262, 33,
		1, 0, 0, 0, 263, 264, 5, 2, 0, 0, 264, 265, 3, 32, 16, 0, 265, 266, 5,
		4, 0, 0, 266, 35, 1, 0, 0, 0, 267, 268, 5, 19, 0, 0, 268, 269, 5, 2, 0,
		0, 269, 270, 3, 14, 7, 0, 270, 271, 5, 4, 0, 0, 271, 272, 5, 5, 0, 0, 272,
		277, 3, 38, 19, 0, 273, 274, 5, 3, 0, 0, 274, 276, 3, 38, 19, 0, 275, 273,
		1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0,
		0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280, 282, 5, 3, 0, 0,
		281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		284, 5, 6, 0, 0, 284, 37, 1, 0, 0, 0, 285, 286, 5, 20, 0, 0, 286, 287,
		3, 50, 25, 0, 287, 288, 5, 21, 0, 0, 288, 289, 3, 12, 6, 0, 289, 39, 1,
		0, 0, 0, 290, 291, 5, 22, 0, 0, 291, 306, 5, 2, 0, 0, 292, 293, 3, 50,
		25, 0, 293, 294, 5, 3, 0, 0, 294, 295, 3, 50, 25, 0, 295, 296, 5, 3, 0,
		0, 296, 297, 3, 50, 25, 0, 297, 298, 5, 3, 0, 0, 298, 299, 3, 48, 24, 0,
		299, 307, 1, 0, 0, 0, 300, 301, 3, 50, 25, 0, 301, 302, 5, 3, 0, 0, 302,
		303, 3, 50, 25, 0, 303, 304, 5, 3, 0, 0, 304, 305, 3, 48, 24, 0, 305, 307,
		1, 0, 0, 0, 306, 292, 1, 0, 0, 0, 306, 300, 1, 0, 0, 0, 307, 308, 1, 0,
		0, 0, 308, 309, 5, 4, 0, 0, 309, 41, 1, 0, 0, 0, 310, 311, 5, 33, 0, 0,
		311, 312, 5, 23, 0, 0, 312, 379, 3, 52, 26, 0, 313, 314, 5, 33, 0, 0, 314,
		315, 5, 23, 0, 0, 315, 379, 3, 48, 24, 0, 316, 317, 5, 33, 0, 0, 317, 318,
		5, 23, 0, 0, 318, 379, 3, 50, 25, 0, 319, 320, 5, 33, 0, 0, 320, 321, 5,
		23, 0, 0, 321, 379, 3, 54, 27, 0, 322, 323, 5, 33, 0, 0, 323, 324, 5, 23,
		0, 0, 324, 325, 5, 11, 0, 0, 325, 379, 5, 12, 0, 0, 326, 327, 5, 33, 0,
		0, 327, 328, 5, 23, 0, 0, 328, 329, 5, 11, 0, 0, 329, 334, 3, 52, 26, 0,
		330, 331, 5, 3, 0, 0, 331, 333, 3, 52, 26, 0, 332, 330, 1, 0, 0, 0, 333,
		336, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337,
		1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 338, 5, 12, 0, 0, 338, 379, 1, 0,
		0, 0, 339, 340, 5, 33, 0, 0, 340, 341, 5, 23, 0, 0, 341, 342, 5, 11, 0,
		0, 342, 347, 3, 48, 24, 0, 343, 344, 5, 3, 0, 0, 344, 346, 3, 48, 24, 0,
		345, 343, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347,
		348, 1, 0, 0, 0, 348, 350, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350, 351,
		5, 12, 0, 0, 351, 379, 1, 0, 0, 0, 352, 353, 5, 33, 0, 0, 353, 354, 5,
		23, 0, 0, 354, 355, 5, 11, 0, 0, 355, 360, 3, 50, 25, 0, 356, 357, 5, 3,
		0, 0, 357, 359, 3, 50, 25, 0, 358, 356, 1, 0, 0, 0, 359, 362, 1, 0, 0,
		0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 363, 1, 0, 0, 0, 362,
		360, 1, 0, 0, 0, 363, 364, 5, 12, 0, 0, 364, 379, 1, 0, 0, 0, 365, 366,
		5, 33, 0, 0, 366, 367, 5, 23, 0, 0, 367, 368, 5, 11, 0, 0, 368, 373, 3,
		54, 27, 0, 369, 370, 5, 3, 0, 0, 370, 372, 3, 54, 27, 0, 371, 369, 1, 0,
		0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0,
		374, 376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377, 5, 12, 0, 0, 377,
		379, 1, 0, 0, 0, 378, 310, 1, 0, 0, 0, 378, 313, 1, 0, 0, 0, 378, 316,
		1, 0, 0, 0, 378, 319, 1, 0, 0, 0, 378, 322, 1, 0, 0, 0, 378, 326, 1, 0,
		0, 0, 378, 339, 1, 0, 0, 0, 378, 352, 1, 0, 0, 0, 378, 365, 1, 0, 0, 0,
		379, 43, 1, 0, 0, 0, 380, 385, 3, 42, 21, 0, 381, 382, 5, 3, 0, 0, 382,
		384, 3, 42, 21, 0, 383, 381, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385, 383,
		1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 45, 1, 0, 0, 0, 387, 385, 1, 0,
		0, 0, 388, 389, 5, 24, 0, 0, 389, 47, 1, 0, 0, 0, 390, 392, 5, 25, 0, 0,
		391, 390, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393,
		394, 5, 34, 0, 0, 394, 49, 1, 0, 0, 0, 395, 396, 5, 35, 0, 0, 396, 51,
		1, 0, 0, 0, 397, 398, 7, 0, 0, 0, 398, 53, 1, 0, 0, 0, 399, 400, 5, 34,
		0, 0, 400, 401, 7, 1, 0, 0, 401, 55, 1, 0, 0, 0, 37, 60, 62, 70, 76, 78,
		100, 119, 128, 131, 135, 141, 148, 151, 160, 165, 168, 179, 184, 186, 190,
		195, 199, 211, 215, 232, 239, 257, 277, 281, 306, 334, 347, 360, 373, 378,
		385, 391,
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

// jagelangParserInit initializes any static state used to implement jagelangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewjagelangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JagelangParserInit() {
	staticData := &JagelangParserStaticData
	staticData.once.Do(jagelangParserInit)
}

// NewjagelangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewjagelangParser(input antlr.TokenStream) *jagelangParser {
	JagelangParserInit()
	this := new(jagelangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JagelangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "jagelang.g4"

	return this
}

// jagelangParser tokens.
const (
	jagelangParserEOF           = antlr.TokenEOF
	jagelangParserT__0          = 1
	jagelangParserT__1          = 2
	jagelangParserT__2          = 3
	jagelangParserT__3          = 4
	jagelangParserT__4          = 5
	jagelangParserT__5          = 6
	jagelangParserT__6          = 7
	jagelangParserT__7          = 8
	jagelangParserT__8          = 9
	jagelangParserT__9          = 10
	jagelangParserT__10         = 11
	jagelangParserT__11         = 12
	jagelangParserT__12         = 13
	jagelangParserT__13         = 14
	jagelangParserT__14         = 15
	jagelangParserT__15         = 16
	jagelangParserT__16         = 17
	jagelangParserT__17         = 18
	jagelangParserT__18         = 19
	jagelangParserT__19         = 20
	jagelangParserT__20         = 21
	jagelangParserT__21         = 22
	jagelangParserT__22         = 23
	jagelangParserT__23         = 24
	jagelangParserT__24         = 25
	jagelangParserT__25         = 26
	jagelangParserT__26         = 27
	jagelangParserT__27         = 28
	jagelangParserT__28         = 29
	jagelangParserT__29         = 30
	jagelangParserTRUE          = 31
	jagelangParserFALSE         = 32
	jagelangParserID            = 33
	jagelangParserINT           = 34
	jagelangParserDQUOTA_STRING = 35
	jagelangParserLINE_COMMENT  = 36
	jagelangParserCOMMENT       = 37
	jagelangParserTERMINATOR    = 38
	jagelangParserWS            = 39
)

// jagelangParser rules.
const (
	jagelangParserRULE_primary                   = 0
	jagelangParserRULE_startDirective            = 1
	jagelangParserRULE_fragmentDirective         = 2
	jagelangParserRULE_registerDirective         = 3
	jagelangParserRULE_onFinishStmt              = 4
	jagelangParserRULE_exedescStmt               = 5
	jagelangParserRULE_leafSnippet               = 6
	jagelangParserRULE_operatorStmt              = 7
	jagelangParserRULE_serialStmt                = 8
	jagelangParserRULE_serialBracketStmt         = 9
	jagelangParserRULE_concurrentStmt            = 10
	jagelangParserRULE_unfoldDirective           = 11
	jagelangParserRULE_goDirective               = 12
	jagelangParserRULE_waitDirective             = 13
	jagelangParserRULE_noCheckMissDirective      = 14
	jagelangParserRULE_skipDirective             = 15
	jagelangParserRULE_wrapStmt                  = 16
	jagelangParserRULE_wrapBracketStmt           = 17
	jagelangParserRULE_switchDirective           = 18
	jagelangParserRULE_switchCaseDirective       = 19
	jagelangParserRULE_registerOperatorDirective = 20
	jagelangParserRULE_argStmt                   = 21
	jagelangParserRULE_argsStmt                  = 22
	jagelangParserRULE_ignoreError               = 23
	jagelangParserRULE_integerLiteral            = 24
	jagelangParserRULE_stringLiteral             = 25
	jagelangParserRULE_boolLiteral               = 26
	jagelangParserRULE_durationLiteral           = 27
)

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStartDirective() []IStartDirectiveContext
	StartDirective(i int) IStartDirectiveContext
	AllFragmentDirective() []IFragmentDirectiveContext
	FragmentDirective(i int) IFragmentDirectiveContext
	AllRegisterDirective() []IRegisterDirectiveContext
	RegisterDirective(i int) IRegisterDirectiveContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllStartDirective() []IStartDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IStartDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartDirectiveContext); ok {
			tst[i] = t.(IStartDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) StartDirective(i int) IStartDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartDirectiveContext); ok {
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

	return t.(IStartDirectiveContext)
}

func (s *PrimaryContext) AllFragmentDirective() []IFragmentDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFragmentDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IFragmentDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFragmentDirectiveContext); ok {
			tst[i] = t.(IFragmentDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) FragmentDirective(i int) IFragmentDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFragmentDirectiveContext); ok {
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

	return t.(IFragmentDirectiveContext)
}

func (s *PrimaryContext) AllRegisterDirective() []IRegisterDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IRegisterDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterDirectiveContext); ok {
			tst[i] = t.(IRegisterDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) RegisterDirective(i int) IRegisterDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterDirectiveContext); ok {
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

	return t.(IRegisterDirectiveContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *jagelangParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, jagelangParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&386) != 0 {
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(56)
				p.StartDirective()
			}

		case 2:
			{
				p.SetState(57)
				p.FragmentDirective()
			}

		case 3:
			{
				p.SetState(58)
				p.RegisterDirective()
			}

		case 4:
			{
				p.SetState(59)
				p.RegisterDirective()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(64)
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

// IStartDirectiveContext is an interface to support dynamic dispatch.
type IStartDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IStringLiteralContext

	// SetName sets the name rule contexts.
	SetName(IStringLiteralContext)

	// Getter signatures
	ExedescStmt() IExedescStmtContext
	StringLiteral() IStringLiteralContext
	ArgsStmt() IArgsStmtContext
	AllOnFinishStmt() []IOnFinishStmtContext
	OnFinishStmt(i int) IOnFinishStmtContext
	AllNoCheckMissDirective() []INoCheckMissDirectiveContext
	NoCheckMissDirective(i int) INoCheckMissDirectiveContext

	// IsStartDirectiveContext differentiates from other interfaces.
	IsStartDirectiveContext()
}

type StartDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyStartDirectiveContext() *StartDirectiveContext {
	var p = new(StartDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_startDirective
	return p
}

func InitEmptyStartDirectiveContext(p *StartDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_startDirective
}

func (*StartDirectiveContext) IsStartDirectiveContext() {}

func NewStartDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartDirectiveContext {
	var p = new(StartDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_startDirective

	return p
}

func (s *StartDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *StartDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *StartDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *StartDirectiveContext) ExedescStmt() IExedescStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExedescStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExedescStmtContext)
}

func (s *StartDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *StartDirectiveContext) ArgsStmt() IArgsStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsStmtContext)
}

func (s *StartDirectiveContext) AllOnFinishStmt() []IOnFinishStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOnFinishStmtContext); ok {
			len++
		}
	}

	tst := make([]IOnFinishStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOnFinishStmtContext); ok {
			tst[i] = t.(IOnFinishStmtContext)
			i++
		}
	}

	return tst
}

func (s *StartDirectiveContext) OnFinishStmt(i int) IOnFinishStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnFinishStmtContext); ok {
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

	return t.(IOnFinishStmtContext)
}

func (s *StartDirectiveContext) AllNoCheckMissDirective() []INoCheckMissDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INoCheckMissDirectiveContext); ok {
			len++
		}
	}

	tst := make([]INoCheckMissDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INoCheckMissDirectiveContext); ok {
			tst[i] = t.(INoCheckMissDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *StartDirectiveContext) NoCheckMissDirective(i int) INoCheckMissDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoCheckMissDirectiveContext); ok {
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

	return t.(INoCheckMissDirectiveContext)
}

func (s *StartDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterStartDirective(s)
	}
}

func (s *StartDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitStartDirective(s)
	}
}

func (p *jagelangParser) StartDirective() (localctx IStartDirectiveContext) {
	localctx = NewStartDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, jagelangParserRULE_startDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(jagelangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(67)

		var _x = p.StringLiteral()

		localctx.(*StartDirectiveContext).name = _x
	}

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__2 {
		{
			p.SetState(68)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.ArgsStmt()
		}

	}
	{
		p.SetState(72)
		p.Match(jagelangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(jagelangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == jagelangParserT__8 || _la == jagelangParserT__15 {
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case jagelangParserT__8:
			{
				p.SetState(74)
				p.OnFinishStmt()
			}

		case jagelangParserT__15:
			{
				p.SetState(75)
				p.NoCheckMissDirective()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.ExedescStmt()
	}
	{
		p.SetState(82)
		p.Match(jagelangParserT__5)
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

// IFragmentDirectiveContext is an interface to support dynamic dispatch.
type IFragmentDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IStringLiteralContext

	// SetName sets the name rule contexts.
	SetName(IStringLiteralContext)

	// Getter signatures
	ExedescStmt() IExedescStmtContext
	StringLiteral() IStringLiteralContext

	// IsFragmentDirectiveContext differentiates from other interfaces.
	IsFragmentDirectiveContext()
}

type FragmentDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyFragmentDirectiveContext() *FragmentDirectiveContext {
	var p = new(FragmentDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_fragmentDirective
	return p
}

func InitEmptyFragmentDirectiveContext(p *FragmentDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_fragmentDirective
}

func (*FragmentDirectiveContext) IsFragmentDirectiveContext() {}

func NewFragmentDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDirectiveContext {
	var p = new(FragmentDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_fragmentDirective

	return p
}

func (s *FragmentDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *FragmentDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *FragmentDirectiveContext) ExedescStmt() IExedescStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExedescStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExedescStmtContext)
}

func (s *FragmentDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *FragmentDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterFragmentDirective(s)
	}
}

func (s *FragmentDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitFragmentDirective(s)
	}
}

func (p *jagelangParser) FragmentDirective() (localctx IFragmentDirectiveContext) {
	localctx = NewFragmentDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, jagelangParserRULE_fragmentDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(jagelangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(86)

		var _x = p.StringLiteral()

		localctx.(*FragmentDirectiveContext).name = _x
	}

	{
		p.SetState(87)
		p.Match(jagelangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(jagelangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.ExedescStmt()
	}
	{
		p.SetState(90)
		p.Match(jagelangParserT__5)
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

// IRegisterDirectiveContext is an interface to support dynamic dispatch.
type IRegisterDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPkg returns the pkg rule contexts.
	GetPkg() IStringLiteralContext

	// SetPkg sets the pkg rule contexts.
	SetPkg(IStringLiteralContext)

	// Getter signatures
	StringLiteral() IStringLiteralContext
	AllRegisterOperatorDirective() []IRegisterOperatorDirectiveContext
	RegisterOperatorDirective(i int) IRegisterOperatorDirectiveContext

	// IsRegisterDirectiveContext differentiates from other interfaces.
	IsRegisterDirectiveContext()
}

type RegisterDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	pkg    IStringLiteralContext
}

func NewEmptyRegisterDirectiveContext() *RegisterDirectiveContext {
	var p = new(RegisterDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_registerDirective
	return p
}

func InitEmptyRegisterDirectiveContext(p *RegisterDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_registerDirective
}

func (*RegisterDirectiveContext) IsRegisterDirectiveContext() {}

func NewRegisterDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterDirectiveContext {
	var p = new(RegisterDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_registerDirective

	return p
}

func (s *RegisterDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterDirectiveContext) GetPkg() IStringLiteralContext { return s.pkg }

func (s *RegisterDirectiveContext) SetPkg(v IStringLiteralContext) { s.pkg = v }

func (s *RegisterDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RegisterDirectiveContext) AllRegisterOperatorDirective() []IRegisterOperatorDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterOperatorDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IRegisterOperatorDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterOperatorDirectiveContext); ok {
			tst[i] = t.(IRegisterOperatorDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *RegisterDirectiveContext) RegisterOperatorDirective(i int) IRegisterOperatorDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterOperatorDirectiveContext); ok {
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

	return t.(IRegisterOperatorDirectiveContext)
}

func (s *RegisterDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterRegisterDirective(s)
	}
}

func (s *RegisterDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitRegisterDirective(s)
	}
}

func (p *jagelangParser) RegisterDirective() (localctx IRegisterDirectiveContext) {
	localctx = NewRegisterDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, jagelangParserRULE_registerDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(jagelangParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(94)

		var _x = p.StringLiteral()

		localctx.(*RegisterDirectiveContext).pkg = _x
	}

	{
		p.SetState(95)
		p.Match(jagelangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(jagelangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == jagelangParserT__21 {
		{
			p.SetState(97)
			p.RegisterOperatorDirective()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(jagelangParserT__5)
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

// IOnFinishStmtContext is an interface to support dynamic dispatch.
type IOnFinishStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExedescStmt() IExedescStmtContext

	// IsOnFinishStmtContext differentiates from other interfaces.
	IsOnFinishStmtContext()
}

type OnFinishStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnFinishStmtContext() *OnFinishStmtContext {
	var p = new(OnFinishStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_onFinishStmt
	return p
}

func InitEmptyOnFinishStmtContext(p *OnFinishStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_onFinishStmt
}

func (*OnFinishStmtContext) IsOnFinishStmtContext() {}

func NewOnFinishStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnFinishStmtContext {
	var p = new(OnFinishStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_onFinishStmt

	return p
}

func (s *OnFinishStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OnFinishStmtContext) ExedescStmt() IExedescStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExedescStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExedescStmtContext)
}

func (s *OnFinishStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnFinishStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnFinishStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterOnFinishStmt(s)
	}
}

func (s *OnFinishStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitOnFinishStmt(s)
	}
}

func (p *jagelangParser) OnFinishStmt() (localctx IOnFinishStmtContext) {
	localctx = NewOnFinishStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, jagelangParserRULE_onFinishStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(jagelangParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(jagelangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(jagelangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.ExedescStmt()
	}
	{
		p.SetState(110)
		p.Match(jagelangParserT__5)
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

// IExedescStmtContext is an interface to support dynamic dispatch.
type IExedescStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SerialStmt() ISerialStmtContext
	ConcurrentStmt() IConcurrentStmtContext
	WrapStmt() IWrapStmtContext
	WrapBracketStmt() IWrapBracketStmtContext
	OperatorStmt() IOperatorStmtContext
	SwitchDirective() ISwitchDirectiveContext
	UnfoldDirective() IUnfoldDirectiveContext

	// IsExedescStmtContext differentiates from other interfaces.
	IsExedescStmtContext()
}

type ExedescStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExedescStmtContext() *ExedescStmtContext {
	var p = new(ExedescStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_exedescStmt
	return p
}

func InitEmptyExedescStmtContext(p *ExedescStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_exedescStmt
}

func (*ExedescStmtContext) IsExedescStmtContext() {}

func NewExedescStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExedescStmtContext {
	var p = new(ExedescStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_exedescStmt

	return p
}

func (s *ExedescStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExedescStmtContext) SerialStmt() ISerialStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISerialStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISerialStmtContext)
}

func (s *ExedescStmtContext) ConcurrentStmt() IConcurrentStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcurrentStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcurrentStmtContext)
}

func (s *ExedescStmtContext) WrapStmt() IWrapStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWrapStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWrapStmtContext)
}

func (s *ExedescStmtContext) WrapBracketStmt() IWrapBracketStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWrapBracketStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWrapBracketStmtContext)
}

func (s *ExedescStmtContext) OperatorStmt() IOperatorStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorStmtContext)
}

func (s *ExedescStmtContext) SwitchDirective() ISwitchDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDirectiveContext)
}

func (s *ExedescStmtContext) UnfoldDirective() IUnfoldDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnfoldDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnfoldDirectiveContext)
}

func (s *ExedescStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExedescStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExedescStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterExedescStmt(s)
	}
}

func (s *ExedescStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitExedescStmt(s)
	}
}

func (p *jagelangParser) ExedescStmt() (localctx IExedescStmtContext) {
	localctx = NewExedescStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, jagelangParserRULE_exedescStmt)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.SerialStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.ConcurrentStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.WrapStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)
			p.WrapBracketStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)
			p.OperatorStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(117)
			p.SwitchDirective()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(118)
			p.UnfoldDirective()
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

// ILeafSnippetContext is an interface to support dynamic dispatch.
type ILeafSnippetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperatorStmt() IOperatorStmtContext
	ConcurrentStmt() IConcurrentStmtContext
	SerialBracketStmt() ISerialBracketStmtContext
	UnfoldDirective() IUnfoldDirectiveContext
	WrapBracketStmt() IWrapBracketStmtContext
	GoDirective() IGoDirectiveContext
	SwitchDirective() ISwitchDirectiveContext

	// IsLeafSnippetContext differentiates from other interfaces.
	IsLeafSnippetContext()
}

type LeafSnippetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeafSnippetContext() *LeafSnippetContext {
	var p = new(LeafSnippetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_leafSnippet
	return p
}

func InitEmptyLeafSnippetContext(p *LeafSnippetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_leafSnippet
}

func (*LeafSnippetContext) IsLeafSnippetContext() {}

func NewLeafSnippetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeafSnippetContext {
	var p = new(LeafSnippetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_leafSnippet

	return p
}

func (s *LeafSnippetContext) GetParser() antlr.Parser { return s.parser }

func (s *LeafSnippetContext) OperatorStmt() IOperatorStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorStmtContext)
}

func (s *LeafSnippetContext) ConcurrentStmt() IConcurrentStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcurrentStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcurrentStmtContext)
}

func (s *LeafSnippetContext) SerialBracketStmt() ISerialBracketStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISerialBracketStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISerialBracketStmtContext)
}

func (s *LeafSnippetContext) UnfoldDirective() IUnfoldDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnfoldDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnfoldDirectiveContext)
}

func (s *LeafSnippetContext) WrapBracketStmt() IWrapBracketStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWrapBracketStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWrapBracketStmtContext)
}

func (s *LeafSnippetContext) GoDirective() IGoDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoDirectiveContext)
}

func (s *LeafSnippetContext) SwitchDirective() ISwitchDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDirectiveContext)
}

func (s *LeafSnippetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeafSnippetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeafSnippetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterLeafSnippet(s)
	}
}

func (s *LeafSnippetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitLeafSnippet(s)
	}
}

func (p *jagelangParser) LeafSnippet() (localctx ILeafSnippetContext) {
	localctx = NewLeafSnippetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, jagelangParserRULE_leafSnippet)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.OperatorStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.ConcurrentStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.SerialBracketStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.UnfoldDirective()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.WrapBracketStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.GoDirective()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(127)
			p.SwitchDirective()
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

// IOperatorStmtContext is an interface to support dynamic dispatch.
type IOperatorStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperatorName returns the operatorName token.
	GetOperatorName() antlr.Token

	// SetOperatorName sets the operatorName token.
	SetOperatorName(antlr.Token)

	// Getter signatures
	IgnoreError() IIgnoreErrorContext
	ID() antlr.TerminalNode
	ArgsStmt() IArgsStmtContext
	AllWaitDirective() []IWaitDirectiveContext
	WaitDirective(i int) IWaitDirectiveContext

	// IsOperatorStmtContext differentiates from other interfaces.
	IsOperatorStmtContext()
}

type OperatorStmtContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	operatorName antlr.Token
}

func NewEmptyOperatorStmtContext() *OperatorStmtContext {
	var p = new(OperatorStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_operatorStmt
	return p
}

func InitEmptyOperatorStmtContext(p *OperatorStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_operatorStmt
}

func (*OperatorStmtContext) IsOperatorStmtContext() {}

func NewOperatorStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStmtContext {
	var p = new(OperatorStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_operatorStmt

	return p
}

func (s *OperatorStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorStmtContext) GetOperatorName() antlr.Token { return s.operatorName }

func (s *OperatorStmtContext) SetOperatorName(v antlr.Token) { s.operatorName = v }

func (s *OperatorStmtContext) IgnoreError() IIgnoreErrorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIgnoreErrorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIgnoreErrorContext)
}

func (s *OperatorStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(jagelangParserID, 0)
}

func (s *OperatorStmtContext) ArgsStmt() IArgsStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsStmtContext)
}

func (s *OperatorStmtContext) AllWaitDirective() []IWaitDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWaitDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IWaitDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWaitDirectiveContext); ok {
			tst[i] = t.(IWaitDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *OperatorStmtContext) WaitDirective(i int) IWaitDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaitDirectiveContext); ok {
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

	return t.(IWaitDirectiveContext)
}

func (s *OperatorStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterOperatorStmt(s)
	}
}

func (s *OperatorStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitOperatorStmt(s)
	}
}

func (p *jagelangParser) OperatorStmt() (localctx IOperatorStmtContext) {
	localctx = NewOperatorStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, jagelangParserRULE_operatorStmt)
	var _la int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__23 {
			{
				p.SetState(130)
				p.IgnoreError()
			}

		}

		{
			p.SetState(133)

			var _m = p.Match(jagelangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__23 {
			{
				p.SetState(134)
				p.IgnoreError()
			}

		}

		{
			p.SetState(137)

			var _m = p.Match(jagelangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(138)
			p.Match(jagelangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(jagelangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__23 {
			{
				p.SetState(140)
				p.IgnoreError()
			}

		}

		{
			p.SetState(143)

			var _m = p.Match(jagelangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__1 {
			{
				p.SetState(144)
				p.Match(jagelangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(145)
				p.ArgsStmt()
			}
			{
				p.SetState(146)
				p.Match(jagelangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__23 {
			{
				p.SetState(150)
				p.IgnoreError()
			}

		}

		{
			p.SetState(153)

			var _m = p.Match(jagelangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__1 {
			{
				p.SetState(154)
				p.Match(jagelangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(155)
				p.WaitDirective()
			}
			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == jagelangParserT__2 {
				{
					p.SetState(156)
					p.Match(jagelangParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(157)
					p.WaitDirective()
				}

				p.SetState(162)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(163)
				p.Match(jagelangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__23 {
			{
				p.SetState(167)
				p.IgnoreError()
			}

		}

		{
			p.SetState(170)

			var _m = p.Match(jagelangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == jagelangParserT__1 {
			{
				p.SetState(171)
				p.Match(jagelangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(172)
				p.ArgsStmt()
			}
			{
				p.SetState(173)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(174)
				p.WaitDirective()
			}
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == jagelangParserT__2 {
				{
					p.SetState(175)
					p.Match(jagelangParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(176)
					p.WaitDirective()
				}

				p.SetState(181)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(182)
				p.Match(jagelangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// ISerialStmtContext is an interface to support dynamic dispatch.
type ISerialStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLeafSnippet() []ILeafSnippetContext
	LeafSnippet(i int) ILeafSnippetContext
	AllSkipDirective() []ISkipDirectiveContext
	SkipDirective(i int) ISkipDirectiveContext

	// IsSerialStmtContext differentiates from other interfaces.
	IsSerialStmtContext()
}

type SerialStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySerialStmtContext() *SerialStmtContext {
	var p = new(SerialStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_serialStmt
	return p
}

func InitEmptySerialStmtContext(p *SerialStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_serialStmt
}

func (*SerialStmtContext) IsSerialStmtContext() {}

func NewSerialStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SerialStmtContext {
	var p = new(SerialStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_serialStmt

	return p
}

func (s *SerialStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SerialStmtContext) AllLeafSnippet() []ILeafSnippetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILeafSnippetContext); ok {
			len++
		}
	}

	tst := make([]ILeafSnippetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILeafSnippetContext); ok {
			tst[i] = t.(ILeafSnippetContext)
			i++
		}
	}

	return tst
}

func (s *SerialStmtContext) LeafSnippet(i int) ILeafSnippetContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeafSnippetContext); ok {
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

	return t.(ILeafSnippetContext)
}

func (s *SerialStmtContext) AllSkipDirective() []ISkipDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISkipDirectiveContext); ok {
			len++
		}
	}

	tst := make([]ISkipDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISkipDirectiveContext); ok {
			tst[i] = t.(ISkipDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *SerialStmtContext) SkipDirective(i int) ISkipDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISkipDirectiveContext); ok {
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

	return t.(ISkipDirectiveContext)
}

func (s *SerialStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SerialStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SerialStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterSerialStmt(s)
	}
}

func (s *SerialStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitSerialStmt(s)
	}
}

func (p *jagelangParser) SerialStmt() (localctx ISerialStmtContext) {
	localctx = NewSerialStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, jagelangParserRULE_serialStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case jagelangParserT__1, jagelangParserT__10, jagelangParserT__12, jagelangParserT__13, jagelangParserT__18, jagelangParserT__23, jagelangParserID:
		{
			p.SetState(188)
			p.LeafSnippet()
		}

	case jagelangParserT__16:
		{
			p.SetState(189)
			p.SkipDirective()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == jagelangParserT__9 {
		{
			p.SetState(192)
			p.Match(jagelangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case jagelangParserT__1, jagelangParserT__10, jagelangParserT__12, jagelangParserT__13, jagelangParserT__18, jagelangParserT__23, jagelangParserID:
			{
				p.SetState(193)
				p.LeafSnippet()
			}

		case jagelangParserT__16:
			{
				p.SetState(194)
				p.SkipDirective()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(199)
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

// ISerialBracketStmtContext is an interface to support dynamic dispatch.
type ISerialBracketStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SerialStmt() ISerialStmtContext

	// IsSerialBracketStmtContext differentiates from other interfaces.
	IsSerialBracketStmtContext()
}

type SerialBracketStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySerialBracketStmtContext() *SerialBracketStmtContext {
	var p = new(SerialBracketStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_serialBracketStmt
	return p
}

func InitEmptySerialBracketStmtContext(p *SerialBracketStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_serialBracketStmt
}

func (*SerialBracketStmtContext) IsSerialBracketStmtContext() {}

func NewSerialBracketStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SerialBracketStmtContext {
	var p = new(SerialBracketStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_serialBracketStmt

	return p
}

func (s *SerialBracketStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SerialBracketStmtContext) SerialStmt() ISerialStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISerialStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISerialStmtContext)
}

func (s *SerialBracketStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SerialBracketStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SerialBracketStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterSerialBracketStmt(s)
	}
}

func (s *SerialBracketStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitSerialBracketStmt(s)
	}
}

func (p *jagelangParser) SerialBracketStmt() (localctx ISerialBracketStmtContext) {
	localctx = NewSerialBracketStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, jagelangParserRULE_serialBracketStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.SerialStmt()
	}
	{
		p.SetState(203)
		p.Match(jagelangParserT__3)
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

// IConcurrentStmtContext is an interface to support dynamic dispatch.
type IConcurrentStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLeafSnippet() []ILeafSnippetContext
	LeafSnippet(i int) ILeafSnippetContext

	// IsConcurrentStmtContext differentiates from other interfaces.
	IsConcurrentStmtContext()
}

type ConcurrentStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcurrentStmtContext() *ConcurrentStmtContext {
	var p = new(ConcurrentStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_concurrentStmt
	return p
}

func InitEmptyConcurrentStmtContext(p *ConcurrentStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_concurrentStmt
}

func (*ConcurrentStmtContext) IsConcurrentStmtContext() {}

func NewConcurrentStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcurrentStmtContext {
	var p = new(ConcurrentStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_concurrentStmt

	return p
}

func (s *ConcurrentStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcurrentStmtContext) AllLeafSnippet() []ILeafSnippetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILeafSnippetContext); ok {
			len++
		}
	}

	tst := make([]ILeafSnippetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILeafSnippetContext); ok {
			tst[i] = t.(ILeafSnippetContext)
			i++
		}
	}

	return tst
}

func (s *ConcurrentStmtContext) LeafSnippet(i int) ILeafSnippetContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeafSnippetContext); ok {
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

	return t.(ILeafSnippetContext)
}

func (s *ConcurrentStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcurrentStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcurrentStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterConcurrentStmt(s)
	}
}

func (s *ConcurrentStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitConcurrentStmt(s)
	}
}

func (p *jagelangParser) ConcurrentStmt() (localctx IConcurrentStmtContext) {
	localctx = NewConcurrentStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, jagelangParserRULE_concurrentStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(jagelangParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.LeafSnippet()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(207)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(208)
				p.LeafSnippet()
			}

		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__2 {
		{
			p.SetState(214)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(217)
		p.Match(jagelangParserT__11)
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

// IUnfoldDirectiveContext is an interface to support dynamic dispatch.
type IUnfoldDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IStringLiteralContext

	// SetName sets the name rule contexts.
	SetName(IStringLiteralContext)

	// Getter signatures
	StringLiteral() IStringLiteralContext

	// IsUnfoldDirectiveContext differentiates from other interfaces.
	IsUnfoldDirectiveContext()
}

type UnfoldDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyUnfoldDirectiveContext() *UnfoldDirectiveContext {
	var p = new(UnfoldDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_unfoldDirective
	return p
}

func InitEmptyUnfoldDirectiveContext(p *UnfoldDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_unfoldDirective
}

func (*UnfoldDirectiveContext) IsUnfoldDirectiveContext() {}

func NewUnfoldDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnfoldDirectiveContext {
	var p = new(UnfoldDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_unfoldDirective

	return p
}

func (s *UnfoldDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *UnfoldDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *UnfoldDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *UnfoldDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *UnfoldDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnfoldDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnfoldDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterUnfoldDirective(s)
	}
}

func (s *UnfoldDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitUnfoldDirective(s)
	}
}

func (p *jagelangParser) UnfoldDirective() (localctx IUnfoldDirectiveContext) {
	localctx = NewUnfoldDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, jagelangParserRULE_unfoldDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(jagelangParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(221)

		var _x = p.StringLiteral()

		localctx.(*UnfoldDirectiveContext).name = _x
	}

	{
		p.SetState(222)
		p.Match(jagelangParserT__3)
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

// IGoDirectiveContext is an interface to support dynamic dispatch.
type IGoDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IStringLiteralContext

	// SetName sets the name rule contexts.
	SetName(IStringLiteralContext)

	// Getter signatures
	ExedescStmt() IExedescStmtContext
	StringLiteral() IStringLiteralContext

	// IsGoDirectiveContext differentiates from other interfaces.
	IsGoDirectiveContext()
}

type GoDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyGoDirectiveContext() *GoDirectiveContext {
	var p = new(GoDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_goDirective
	return p
}

func InitEmptyGoDirectiveContext(p *GoDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_goDirective
}

func (*GoDirectiveContext) IsGoDirectiveContext() {}

func NewGoDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoDirectiveContext {
	var p = new(GoDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_goDirective

	return p
}

func (s *GoDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *GoDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *GoDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *GoDirectiveContext) ExedescStmt() IExedescStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExedescStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExedescStmtContext)
}

func (s *GoDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *GoDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterGoDirective(s)
	}
}

func (s *GoDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitGoDirective(s)
	}
}

func (p *jagelangParser) GoDirective() (localctx IGoDirectiveContext) {
	localctx = NewGoDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, jagelangParserRULE_goDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(jagelangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.ExedescStmt()
	}
	{
		p.SetState(227)
		p.Match(jagelangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(228)

		var _x = p.StringLiteral()

		localctx.(*GoDirectiveContext).name = _x
	}

	{
		p.SetState(229)
		p.Match(jagelangParserT__3)
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

// IWaitDirectiveContext is an interface to support dynamic dispatch.
type IWaitDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IStringLiteralContext

	// SetName sets the name rule contexts.
	SetName(IStringLiteralContext)

	// Getter signatures
	IgnoreError() IIgnoreErrorContext
	StringLiteral() IStringLiteralContext
	ArgsStmt() IArgsStmtContext

	// IsWaitDirectiveContext differentiates from other interfaces.
	IsWaitDirectiveContext()
}

type WaitDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyWaitDirectiveContext() *WaitDirectiveContext {
	var p = new(WaitDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_waitDirective
	return p
}

func InitEmptyWaitDirectiveContext(p *WaitDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_waitDirective
}

func (*WaitDirectiveContext) IsWaitDirectiveContext() {}

func NewWaitDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaitDirectiveContext {
	var p = new(WaitDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_waitDirective

	return p
}

func (s *WaitDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *WaitDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *WaitDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *WaitDirectiveContext) IgnoreError() IIgnoreErrorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIgnoreErrorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIgnoreErrorContext)
}

func (s *WaitDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *WaitDirectiveContext) ArgsStmt() IArgsStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsStmtContext)
}

func (s *WaitDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaitDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WaitDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterWaitDirective(s)
	}
}

func (s *WaitDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitWaitDirective(s)
	}
}

func (p *jagelangParser) WaitDirective() (localctx IWaitDirectiveContext) {
	localctx = NewWaitDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, jagelangParserRULE_waitDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__23 {
		{
			p.SetState(231)
			p.IgnoreError()
		}

	}
	{
		p.SetState(234)
		p.Match(jagelangParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(236)

		var _x = p.StringLiteral()

		localctx.(*WaitDirectiveContext).name = _x
	}

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__2 {
		{
			p.SetState(237)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.ArgsStmt()
		}

	}
	{
		p.SetState(241)
		p.Match(jagelangParserT__3)
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

// INoCheckMissDirectiveContext is an interface to support dynamic dispatch.
type INoCheckMissDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNoCheckMissDirectiveContext differentiates from other interfaces.
	IsNoCheckMissDirectiveContext()
}

type NoCheckMissDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoCheckMissDirectiveContext() *NoCheckMissDirectiveContext {
	var p = new(NoCheckMissDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_noCheckMissDirective
	return p
}

func InitEmptyNoCheckMissDirectiveContext(p *NoCheckMissDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_noCheckMissDirective
}

func (*NoCheckMissDirectiveContext) IsNoCheckMissDirectiveContext() {}

func NewNoCheckMissDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoCheckMissDirectiveContext {
	var p = new(NoCheckMissDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_noCheckMissDirective

	return p
}

func (s *NoCheckMissDirectiveContext) GetParser() antlr.Parser { return s.parser }
func (s *NoCheckMissDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoCheckMissDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoCheckMissDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterNoCheckMissDirective(s)
	}
}

func (s *NoCheckMissDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitNoCheckMissDirective(s)
	}
}

func (p *jagelangParser) NoCheckMissDirective() (localctx INoCheckMissDirectiveContext) {
	localctx = NewNoCheckMissDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, jagelangParserRULE_noCheckMissDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(jagelangParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(jagelangParserT__3)
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

// ISkipDirectiveContext is an interface to support dynamic dispatch.
type ISkipDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperatorStmt() IOperatorStmtContext

	// IsSkipDirectiveContext differentiates from other interfaces.
	IsSkipDirectiveContext()
}

type SkipDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipDirectiveContext() *SkipDirectiveContext {
	var p = new(SkipDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_skipDirective
	return p
}

func InitEmptySkipDirectiveContext(p *SkipDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_skipDirective
}

func (*SkipDirectiveContext) IsSkipDirectiveContext() {}

func NewSkipDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipDirectiveContext {
	var p = new(SkipDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_skipDirective

	return p
}

func (s *SkipDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipDirectiveContext) OperatorStmt() IOperatorStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorStmtContext)
}

func (s *SkipDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterSkipDirective(s)
	}
}

func (s *SkipDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitSkipDirective(s)
	}
}

func (p *jagelangParser) SkipDirective() (localctx ISkipDirectiveContext) {
	localctx = NewSkipDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, jagelangParserRULE_skipDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(jagelangParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.OperatorStmt()
	}
	{
		p.SetState(250)
		p.Match(jagelangParserT__3)
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

// IWrapStmtContext is an interface to support dynamic dispatch.
type IWrapStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperatorStmt() []IOperatorStmtContext
	OperatorStmt(i int) IOperatorStmtContext
	LeafSnippet() ILeafSnippetContext

	// IsWrapStmtContext differentiates from other interfaces.
	IsWrapStmtContext()
}

type WrapStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWrapStmtContext() *WrapStmtContext {
	var p = new(WrapStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_wrapStmt
	return p
}

func InitEmptyWrapStmtContext(p *WrapStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_wrapStmt
}

func (*WrapStmtContext) IsWrapStmtContext() {}

func NewWrapStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WrapStmtContext {
	var p = new(WrapStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_wrapStmt

	return p
}

func (s *WrapStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WrapStmtContext) AllOperatorStmt() []IOperatorStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperatorStmtContext); ok {
			len++
		}
	}

	tst := make([]IOperatorStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperatorStmtContext); ok {
			tst[i] = t.(IOperatorStmtContext)
			i++
		}
	}

	return tst
}

func (s *WrapStmtContext) OperatorStmt(i int) IOperatorStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStmtContext); ok {
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

	return t.(IOperatorStmtContext)
}

func (s *WrapStmtContext) LeafSnippet() ILeafSnippetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeafSnippetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeafSnippetContext)
}

func (s *WrapStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WrapStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WrapStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterWrapStmt(s)
	}
}

func (s *WrapStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitWrapStmt(s)
	}
}

func (p *jagelangParser) WrapStmt() (localctx IWrapStmtContext) {
	localctx = NewWrapStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, jagelangParserRULE_wrapStmt)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.OperatorStmt()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(253)
				p.Match(jagelangParserT__17)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(254)
				p.OperatorStmt()
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(jagelangParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.LeafSnippet()
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

// IWrapBracketStmtContext is an interface to support dynamic dispatch.
type IWrapBracketStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WrapStmt() IWrapStmtContext

	// IsWrapBracketStmtContext differentiates from other interfaces.
	IsWrapBracketStmtContext()
}

type WrapBracketStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWrapBracketStmtContext() *WrapBracketStmtContext {
	var p = new(WrapBracketStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_wrapBracketStmt
	return p
}

func InitEmptyWrapBracketStmtContext(p *WrapBracketStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_wrapBracketStmt
}

func (*WrapBracketStmtContext) IsWrapBracketStmtContext() {}

func NewWrapBracketStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WrapBracketStmtContext {
	var p = new(WrapBracketStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_wrapBracketStmt

	return p
}

func (s *WrapBracketStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WrapBracketStmtContext) WrapStmt() IWrapStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWrapStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWrapStmtContext)
}

func (s *WrapBracketStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WrapBracketStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WrapBracketStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterWrapBracketStmt(s)
	}
}

func (s *WrapBracketStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitWrapBracketStmt(s)
	}
}

func (p *jagelangParser) WrapBracketStmt() (localctx IWrapBracketStmtContext) {
	localctx = NewWrapBracketStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, jagelangParserRULE_wrapBracketStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.WrapStmt()
	}
	{
		p.SetState(265)
		p.Match(jagelangParserT__3)
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

// ISwitchDirectiveContext is an interface to support dynamic dispatch.
type ISwitchDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperatorStmt() IOperatorStmtContext
	AllSwitchCaseDirective() []ISwitchCaseDirectiveContext
	SwitchCaseDirective(i int) ISwitchCaseDirectiveContext

	// IsSwitchDirectiveContext differentiates from other interfaces.
	IsSwitchDirectiveContext()
}

type SwitchDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchDirectiveContext() *SwitchDirectiveContext {
	var p = new(SwitchDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_switchDirective
	return p
}

func InitEmptySwitchDirectiveContext(p *SwitchDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_switchDirective
}

func (*SwitchDirectiveContext) IsSwitchDirectiveContext() {}

func NewSwitchDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDirectiveContext {
	var p = new(SwitchDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_switchDirective

	return p
}

func (s *SwitchDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDirectiveContext) OperatorStmt() IOperatorStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorStmtContext)
}

func (s *SwitchDirectiveContext) AllSwitchCaseDirective() []ISwitchCaseDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseDirectiveContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseDirectiveContext); ok {
			tst[i] = t.(ISwitchCaseDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *SwitchDirectiveContext) SwitchCaseDirective(i int) ISwitchCaseDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseDirectiveContext); ok {
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

	return t.(ISwitchCaseDirectiveContext)
}

func (s *SwitchDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterSwitchDirective(s)
	}
}

func (s *SwitchDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitSwitchDirective(s)
	}
}

func (p *jagelangParser) SwitchDirective() (localctx ISwitchDirectiveContext) {
	localctx = NewSwitchDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, jagelangParserRULE_switchDirective)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(jagelangParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.OperatorStmt()
	}
	{
		p.SetState(270)
		p.Match(jagelangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(jagelangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.SwitchCaseDirective()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(273)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(274)
				p.SwitchCaseDirective()
			}

		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__2 {
		{
			p.SetState(280)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(283)
		p.Match(jagelangParserT__5)
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

// ISwitchCaseDirectiveContext is an interface to support dynamic dispatch.
type ISwitchCaseDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCaseName returns the caseName rule contexts.
	GetCaseName() IStringLiteralContext

	// SetCaseName sets the caseName rule contexts.
	SetCaseName(IStringLiteralContext)

	// Getter signatures
	LeafSnippet() ILeafSnippetContext
	StringLiteral() IStringLiteralContext

	// IsSwitchCaseDirectiveContext differentiates from other interfaces.
	IsSwitchCaseDirectiveContext()
}

type SwitchCaseDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	caseName IStringLiteralContext
}

func NewEmptySwitchCaseDirectiveContext() *SwitchCaseDirectiveContext {
	var p = new(SwitchCaseDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_switchCaseDirective
	return p
}

func InitEmptySwitchCaseDirectiveContext(p *SwitchCaseDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_switchCaseDirective
}

func (*SwitchCaseDirectiveContext) IsSwitchCaseDirectiveContext() {}

func NewSwitchCaseDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseDirectiveContext {
	var p = new(SwitchCaseDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_switchCaseDirective

	return p
}

func (s *SwitchCaseDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseDirectiveContext) GetCaseName() IStringLiteralContext { return s.caseName }

func (s *SwitchCaseDirectiveContext) SetCaseName(v IStringLiteralContext) { s.caseName = v }

func (s *SwitchCaseDirectiveContext) LeafSnippet() ILeafSnippetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeafSnippetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeafSnippetContext)
}

func (s *SwitchCaseDirectiveContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *SwitchCaseDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterSwitchCaseDirective(s)
	}
}

func (s *SwitchCaseDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitSwitchCaseDirective(s)
	}
}

func (p *jagelangParser) SwitchCaseDirective() (localctx ISwitchCaseDirectiveContext) {
	localctx = NewSwitchCaseDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, jagelangParserRULE_switchCaseDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(jagelangParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(286)

		var _x = p.StringLiteral()

		localctx.(*SwitchCaseDirectiveContext).caseName = _x
	}

	{
		p.SetState(287)
		p.Match(jagelangParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.LeafSnippet()
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

// IRegisterOperatorDirectiveContext is an interface to support dynamic dispatch.
type IRegisterOperatorDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPkgPath returns the pkgPath rule contexts.
	GetPkgPath() IStringLiteralContext

	// GetStructName returns the structName rule contexts.
	GetStructName() IStringLiteralContext

	// GetOperatorName returns the operatorName rule contexts.
	GetOperatorName() IStringLiteralContext

	// GetSeq returns the seq rule contexts.
	GetSeq() IIntegerLiteralContext

	// SetPkgPath sets the pkgPath rule contexts.
	SetPkgPath(IStringLiteralContext)

	// SetStructName sets the structName rule contexts.
	SetStructName(IStringLiteralContext)

	// SetOperatorName sets the operatorName rule contexts.
	SetOperatorName(IStringLiteralContext)

	// SetSeq sets the seq rule contexts.
	SetSeq(IIntegerLiteralContext)

	// Getter signatures
	AllStringLiteral() []IStringLiteralContext
	StringLiteral(i int) IStringLiteralContext
	IntegerLiteral() IIntegerLiteralContext

	// IsRegisterOperatorDirectiveContext differentiates from other interfaces.
	IsRegisterOperatorDirectiveContext()
}

type RegisterOperatorDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	pkgPath      IStringLiteralContext
	structName   IStringLiteralContext
	operatorName IStringLiteralContext
	seq          IIntegerLiteralContext
}

func NewEmptyRegisterOperatorDirectiveContext() *RegisterOperatorDirectiveContext {
	var p = new(RegisterOperatorDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_registerOperatorDirective
	return p
}

func InitEmptyRegisterOperatorDirectiveContext(p *RegisterOperatorDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_registerOperatorDirective
}

func (*RegisterOperatorDirectiveContext) IsRegisterOperatorDirectiveContext() {}

func NewRegisterOperatorDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterOperatorDirectiveContext {
	var p = new(RegisterOperatorDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_registerOperatorDirective

	return p
}

func (s *RegisterOperatorDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterOperatorDirectiveContext) GetPkgPath() IStringLiteralContext { return s.pkgPath }

func (s *RegisterOperatorDirectiveContext) GetStructName() IStringLiteralContext { return s.structName }

func (s *RegisterOperatorDirectiveContext) GetOperatorName() IStringLiteralContext {
	return s.operatorName
}

func (s *RegisterOperatorDirectiveContext) GetSeq() IIntegerLiteralContext { return s.seq }

func (s *RegisterOperatorDirectiveContext) SetPkgPath(v IStringLiteralContext) { s.pkgPath = v }

func (s *RegisterOperatorDirectiveContext) SetStructName(v IStringLiteralContext) { s.structName = v }

func (s *RegisterOperatorDirectiveContext) SetOperatorName(v IStringLiteralContext) {
	s.operatorName = v
}

func (s *RegisterOperatorDirectiveContext) SetSeq(v IIntegerLiteralContext) { s.seq = v }

func (s *RegisterOperatorDirectiveContext) AllStringLiteral() []IStringLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringLiteralContext); ok {
			len++
		}
	}

	tst := make([]IStringLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringLiteralContext); ok {
			tst[i] = t.(IStringLiteralContext)
			i++
		}
	}

	return tst
}

func (s *RegisterOperatorDirectiveContext) StringLiteral(i int) IStringLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
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

	return t.(IStringLiteralContext)
}

func (s *RegisterOperatorDirectiveContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *RegisterOperatorDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterOperatorDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterOperatorDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterRegisterOperatorDirective(s)
	}
}

func (s *RegisterOperatorDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitRegisterOperatorDirective(s)
	}
}

func (p *jagelangParser) RegisterOperatorDirective() (localctx IRegisterOperatorDirectiveContext) {
	localctx = NewRegisterOperatorDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, jagelangParserRULE_registerOperatorDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(jagelangParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(jagelangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(292)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).pkgPath = _x
		}

		{
			p.SetState(293)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(294)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).structName = _x
		}

		{
			p.SetState(295)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(296)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).operatorName = _x
		}

		{
			p.SetState(297)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(298)

			var _x = p.IntegerLiteral()

			localctx.(*RegisterOperatorDirectiveContext).seq = _x
		}

	case 2:
		{
			p.SetState(300)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).pkgPath = _x
		}

		{
			p.SetState(301)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(302)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).structName = _x
		}

		{
			p.SetState(303)
			p.Match(jagelangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(304)

			var _x = p.IntegerLiteral()

			localctx.(*RegisterOperatorDirectiveContext).seq = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(308)
		p.Match(jagelangParserT__3)
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

// IArgStmtContext is an interface to support dynamic dispatch.
type IArgStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// Getter signatures
	AllBoolLiteral() []IBoolLiteralContext
	BoolLiteral(i int) IBoolLiteralContext
	ID() antlr.TerminalNode
	AllIntegerLiteral() []IIntegerLiteralContext
	IntegerLiteral(i int) IIntegerLiteralContext
	AllStringLiteral() []IStringLiteralContext
	StringLiteral(i int) IStringLiteralContext
	AllDurationLiteral() []IDurationLiteralContext
	DurationLiteral(i int) IDurationLiteralContext

	// IsArgStmtContext differentiates from other interfaces.
	IsArgStmtContext()
}

type ArgStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyArgStmtContext() *ArgStmtContext {
	var p = new(ArgStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_argStmt
	return p
}

func InitEmptyArgStmtContext(p *ArgStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_argStmt
}

func (*ArgStmtContext) IsArgStmtContext() {}

func NewArgStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgStmtContext {
	var p = new(ArgStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_argStmt

	return p
}

func (s *ArgStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgStmtContext) GetKey() antlr.Token { return s.key }

func (s *ArgStmtContext) SetKey(v antlr.Token) { s.key = v }

func (s *ArgStmtContext) AllBoolLiteral() []IBoolLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolLiteralContext); ok {
			len++
		}
	}

	tst := make([]IBoolLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolLiteralContext); ok {
			tst[i] = t.(IBoolLiteralContext)
			i++
		}
	}

	return tst
}

func (s *ArgStmtContext) BoolLiteral(i int) IBoolLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLiteralContext); ok {
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

	return t.(IBoolLiteralContext)
}

func (s *ArgStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(jagelangParserID, 0)
}

func (s *ArgStmtContext) AllIntegerLiteral() []IIntegerLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			len++
		}
	}

	tst := make([]IIntegerLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntegerLiteralContext); ok {
			tst[i] = t.(IIntegerLiteralContext)
			i++
		}
	}

	return tst
}

func (s *ArgStmtContext) IntegerLiteral(i int) IIntegerLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
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

	return t.(IIntegerLiteralContext)
}

func (s *ArgStmtContext) AllStringLiteral() []IStringLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringLiteralContext); ok {
			len++
		}
	}

	tst := make([]IStringLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringLiteralContext); ok {
			tst[i] = t.(IStringLiteralContext)
			i++
		}
	}

	return tst
}

func (s *ArgStmtContext) StringLiteral(i int) IStringLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
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

	return t.(IStringLiteralContext)
}

func (s *ArgStmtContext) AllDurationLiteral() []IDurationLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationLiteralContext); ok {
			len++
		}
	}

	tst := make([]IDurationLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationLiteralContext); ok {
			tst[i] = t.(IDurationLiteralContext)
			i++
		}
	}

	return tst
}

func (s *ArgStmtContext) DurationLiteral(i int) IDurationLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationLiteralContext); ok {
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

	return t.(IDurationLiteralContext)
}

func (s *ArgStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterArgStmt(s)
	}
}

func (s *ArgStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitArgStmt(s)
	}
}

func (p *jagelangParser) ArgStmt() (localctx IArgStmtContext) {
	localctx = NewArgStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, jagelangParserRULE_argStmt)
	var _la int

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(311)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.BoolLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(314)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(317)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.StringLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(319)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(320)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.DurationLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(322)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(323)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(jagelangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(jagelangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(326)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(327)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(jagelangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.BoolLiteral()
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == jagelangParserT__2 {
			{
				p.SetState(330)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(331)
				p.BoolLiteral()
			}

			p.SetState(336)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(337)
			p.Match(jagelangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(339)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(340)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(jagelangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.IntegerLiteral()
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == jagelangParserT__2 {
			{
				p.SetState(343)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(344)
				p.IntegerLiteral()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(350)
			p.Match(jagelangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(352)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(353)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(jagelangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.StringLiteral()
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == jagelangParserT__2 {
			{
				p.SetState(356)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(357)
				p.StringLiteral()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(363)
			p.Match(jagelangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(365)

			var _m = p.Match(jagelangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(366)
			p.Match(jagelangParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(jagelangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.DurationLiteral()
		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == jagelangParserT__2 {
			{
				p.SetState(369)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(370)
				p.DurationLiteral()
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(376)
			p.Match(jagelangParserT__11)
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

// IArgsStmtContext is an interface to support dynamic dispatch.
type IArgsStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgStmt() []IArgStmtContext
	ArgStmt(i int) IArgStmtContext

	// IsArgsStmtContext differentiates from other interfaces.
	IsArgsStmtContext()
}

type ArgsStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsStmtContext() *ArgsStmtContext {
	var p = new(ArgsStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_argsStmt
	return p
}

func InitEmptyArgsStmtContext(p *ArgsStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_argsStmt
}

func (*ArgsStmtContext) IsArgsStmtContext() {}

func NewArgsStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsStmtContext {
	var p = new(ArgsStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_argsStmt

	return p
}

func (s *ArgsStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsStmtContext) AllArgStmt() []IArgStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgStmtContext); ok {
			len++
		}
	}

	tst := make([]IArgStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgStmtContext); ok {
			tst[i] = t.(IArgStmtContext)
			i++
		}
	}

	return tst
}

func (s *ArgsStmtContext) ArgStmt(i int) IArgStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgStmtContext); ok {
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

	return t.(IArgStmtContext)
}

func (s *ArgsStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterArgsStmt(s)
	}
}

func (s *ArgsStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitArgsStmt(s)
	}
}

func (p *jagelangParser) ArgsStmt() (localctx IArgsStmtContext) {
	localctx = NewArgsStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, jagelangParserRULE_argsStmt)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.ArgStmt()
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.Match(jagelangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(382)
				p.ArgStmt()
			}

		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
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

// IIgnoreErrorContext is an interface to support dynamic dispatch.
type IIgnoreErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIgnoreErrorContext differentiates from other interfaces.
	IsIgnoreErrorContext()
}

type IgnoreErrorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIgnoreErrorContext() *IgnoreErrorContext {
	var p = new(IgnoreErrorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_ignoreError
	return p
}

func InitEmptyIgnoreErrorContext(p *IgnoreErrorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_ignoreError
}

func (*IgnoreErrorContext) IsIgnoreErrorContext() {}

func NewIgnoreErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoreErrorContext {
	var p = new(IgnoreErrorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_ignoreError

	return p
}

func (s *IgnoreErrorContext) GetParser() antlr.Parser { return s.parser }
func (s *IgnoreErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgnoreErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgnoreErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterIgnoreError(s)
	}
}

func (s *IgnoreErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitIgnoreError(s)
	}
}

func (p *jagelangParser) IgnoreError() (localctx IIgnoreErrorContext) {
	localctx = NewIgnoreErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, jagelangParserRULE_ignoreError)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(jagelangParserT__23)
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

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_integerLiteral
	return p
}

func InitEmptyIntegerLiteralContext(p *IntegerLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_integerLiteral
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(jagelangParserINT, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *jagelangParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, jagelangParserRULE_integerLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == jagelangParserT__24 {
		{
			p.SetState(390)
			p.Match(jagelangParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(393)
		p.Match(jagelangParserINT)
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

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DQUOTA_STRING() antlr.TerminalNode

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(jagelangParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *jagelangParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, jagelangParserRULE_stringLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(jagelangParserDQUOTA_STRING)
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

// IBoolLiteralContext is an interface to support dynamic dispatch.
type IBoolLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolLiteralContext differentiates from other interfaces.
	IsBoolLiteralContext()
}

type BoolLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLiteralContext() *BoolLiteralContext {
	var p = new(BoolLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_boolLiteral
	return p
}

func InitEmptyBoolLiteralContext(p *BoolLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_boolLiteral
}

func (*BoolLiteralContext) IsBoolLiteralContext() {}

func NewBoolLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_boolLiteral

	return p
}

func (s *BoolLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(jagelangParserTRUE, 0)
}

func (s *BoolLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(jagelangParserFALSE, 0)
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (p *jagelangParser) BoolLiteral() (localctx IBoolLiteralContext) {
	localctx = NewBoolLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, jagelangParserRULE_boolLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		_la = p.GetTokenStream().LA(1)

		if !(_la == jagelangParserTRUE || _la == jagelangParserFALSE) {
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

// IDurationLiteralContext is an interface to support dynamic dispatch.
type IDurationLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsDurationLiteralContext differentiates from other interfaces.
	IsDurationLiteralContext()
}

type DurationLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationLiteralContext() *DurationLiteralContext {
	var p = new(DurationLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_durationLiteral
	return p
}

func InitEmptyDurationLiteralContext(p *DurationLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = jagelangParserRULE_durationLiteral
}

func (*DurationLiteralContext) IsDurationLiteralContext() {}

func NewDurationLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationLiteralContext {
	var p = new(DurationLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = jagelangParserRULE_durationLiteral

	return p
}

func (s *DurationLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(jagelangParserINT, 0)
}

func (s *DurationLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.EnterDurationLiteral(s)
	}
}

func (s *DurationLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jagelangListener); ok {
		listenerT.ExitDurationLiteral(s)
	}
}

func (p *jagelangParser) DurationLiteral() (localctx IDurationLiteralContext) {
	localctx = NewDurationLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, jagelangParserRULE_durationLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(jagelangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2080374784) != 0) {
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
