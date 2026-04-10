// Code generated from iantlr/gorchlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package iantlr // gorchlang
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

type gorchlangParser struct {
	*antlr.BaseParser
}

var GorchlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gorchlangParserInit() {
	staticData := &GorchlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'START'", "'('", "','", "')'", "'{'", "'}'", "'FRAGMENT'", "'REGISTER'",
		"'ON_FINISH'", "'->'", "'['", "']'", "'UNFOLD'", "'GO'", "'WAIT'", "'NO_CHECK_MISS'",
		"'SKIP'", "'|'", "'SWITCH'", "'CASE'", "'=>'", "'LOOP'", "'RETRY'",
		"'TRACE'", "'UNTIL'", "'BREAK'", "'OPERATOR'", "'='", "'@'", "'-'",
		"'s'", "'ms'", "'\\u00B5s'", "'h'", "'m'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "TRUE", "FALSE", "ID", "INT", "DQUOTA_STRING", "LINE_COMMENT",
		"COMMENT", "TERMINATOR", "WS",
	}
	staticData.RuleNames = []string{
		"primary", "startDirective", "fragmentDirective", "registerDirective",
		"onFinishStmt", "exedescStmt", "leafSnippet", "operatorStmt", "serialStmt",
		"serialBracketStmt", "concurrentStmt", "unfoldDirective", "goDirective",
		"waitDirective", "noCheckMissDirective", "skipDirective", "wrapStmt",
		"wrapBracketStmt", "switchDirective", "switchCaseDirective", "loopDirective",
		"retryDirective", "traceDirective", "untilDirective", "breakDirective",
		"registerOperatorDirective", "argStmt", "argsStmt", "ignoreError", "integerLiteral",
		"stringLiteral", "boolLiteral", "durationLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 459, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 71, 8, 0, 10, 0, 12, 0,
		74, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 93, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 112, 8, 3, 10, 3, 12, 3, 115, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 138, 8, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 152,
		8, 6, 1, 7, 3, 7, 155, 8, 7, 1, 7, 1, 7, 3, 7, 159, 8, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 165, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 172, 8, 7,
		1, 7, 3, 7, 175, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 182, 8, 7, 10,
		7, 12, 7, 185, 9, 7, 1, 7, 1, 7, 3, 7, 189, 8, 7, 1, 7, 3, 7, 192, 8, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 201, 8, 7, 10, 7, 12, 7,
		204, 9, 7, 1, 7, 1, 7, 3, 7, 208, 8, 7, 3, 7, 210, 8, 7, 1, 8, 1, 8, 3,
		8, 214, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 219, 8, 8, 4, 8, 221, 8, 8, 11, 8,
		12, 8, 222, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		233, 8, 10, 10, 10, 12, 10, 236, 9, 10, 1, 10, 3, 10, 239, 8, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 3, 13, 256, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 263, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 279, 8, 16,
		10, 16, 12, 16, 282, 9, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 299,
		8, 18, 10, 18, 12, 18, 302, 9, 18, 1, 18, 3, 18, 305, 8, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 3, 25, 363, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 389, 8,
		26, 10, 26, 12, 26, 392, 9, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 5, 26, 402, 8, 26, 10, 26, 12, 26, 405, 9, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 415, 8, 26, 10, 26,
		12, 26, 418, 9, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 5, 26, 428, 8, 26, 10, 26, 12, 26, 431, 9, 26, 1, 26, 1, 26, 3, 26,
		435, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 440, 8, 27, 10, 27, 12, 27, 443,
		9, 27, 1, 28, 1, 28, 1, 29, 3, 29, 448, 8, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 0, 0, 33, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 2, 1, 0, 36, 37, 1, 0, 31, 35,
		496, 0, 72, 1, 0, 0, 0, 2, 75, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 105, 1,
		0, 0, 0, 8, 118, 1, 0, 0, 0, 10, 137, 1, 0, 0, 0, 12, 151, 1, 0, 0, 0,
		14, 209, 1, 0, 0, 0, 16, 213, 1, 0, 0, 0, 18, 224, 1, 0, 0, 0, 20, 228,
		1, 0, 0, 0, 22, 242, 1, 0, 0, 0, 24, 247, 1, 0, 0, 0, 26, 255, 1, 0, 0,
		0, 28, 266, 1, 0, 0, 0, 30, 270, 1, 0, 0, 0, 32, 275, 1, 0, 0, 0, 34, 286,
		1, 0, 0, 0, 36, 290, 1, 0, 0, 0, 38, 308, 1, 0, 0, 0, 40, 313, 1, 0, 0,
		0, 42, 321, 1, 0, 0, 0, 44, 329, 1, 0, 0, 0, 46, 337, 1, 0, 0, 0, 48, 342,
		1, 0, 0, 0, 50, 346, 1, 0, 0, 0, 52, 434, 1, 0, 0, 0, 54, 436, 1, 0, 0,
		0, 56, 444, 1, 0, 0, 0, 58, 447, 1, 0, 0, 0, 60, 451, 1, 0, 0, 0, 62, 453,
		1, 0, 0, 0, 64, 455, 1, 0, 0, 0, 66, 71, 3, 2, 1, 0, 67, 71, 3, 4, 2, 0,
		68, 71, 3, 6, 3, 0, 69, 71, 3, 6, 3, 0, 70, 66, 1, 0, 0, 0, 70, 67, 1,
		0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72,
		70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 1, 1, 0, 0, 0, 74, 72, 1, 0, 0,
		0, 75, 76, 5, 1, 0, 0, 76, 77, 5, 2, 0, 0, 77, 80, 3, 60, 30, 0, 78, 79,
		5, 3, 0, 0, 79, 81, 3, 54, 27, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0,
		0, 81, 82, 1, 0, 0, 0, 82, 83, 5, 4, 0, 0, 83, 92, 5, 5, 0, 0, 84, 85,
		3, 8, 4, 0, 85, 86, 3, 28, 14, 0, 86, 93, 1, 0, 0, 0, 87, 88, 3, 28, 14,
		0, 88, 89, 3, 8, 4, 0, 89, 93, 1, 0, 0, 0, 90, 93, 3, 28, 14, 0, 91, 93,
		3, 8, 4, 0, 92, 84, 1, 0, 0, 0, 92, 87, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 3,
		10, 5, 0, 95, 96, 5, 6, 0, 0, 96, 3, 1, 0, 0, 0, 97, 98, 5, 7, 0, 0, 98,
		99, 5, 2, 0, 0, 99, 100, 3, 60, 30, 0, 100, 101, 5, 4, 0, 0, 101, 102,
		5, 5, 0, 0, 102, 103, 3, 10, 5, 0, 103, 104, 5, 6, 0, 0, 104, 5, 1, 0,
		0, 0, 105, 106, 5, 8, 0, 0, 106, 107, 5, 2, 0, 0, 107, 108, 3, 60, 30,
		0, 108, 109, 5, 4, 0, 0, 109, 113, 5, 5, 0, 0, 110, 112, 3, 50, 25, 0,
		111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117,
		5, 6, 0, 0, 117, 7, 1, 0, 0, 0, 118, 119, 5, 9, 0, 0, 119, 120, 5, 2, 0,
		0, 120, 121, 5, 4, 0, 0, 121, 122, 5, 5, 0, 0, 122, 123, 3, 10, 5, 0, 123,
		124, 5, 6, 0, 0, 124, 9, 1, 0, 0, 0, 125, 138, 3, 16, 8, 0, 126, 138, 3,
		20, 10, 0, 127, 138, 3, 32, 16, 0, 128, 138, 3, 34, 17, 0, 129, 138, 3,
		14, 7, 0, 130, 138, 3, 36, 18, 0, 131, 138, 3, 22, 11, 0, 132, 138, 3,
		40, 20, 0, 133, 138, 3, 42, 21, 0, 134, 138, 3, 44, 22, 0, 135, 138, 3,
		46, 23, 0, 136, 138, 3, 48, 24, 0, 137, 125, 1, 0, 0, 0, 137, 126, 1, 0,
		0, 0, 137, 127, 1, 0, 0, 0, 137, 128, 1, 0, 0, 0, 137, 129, 1, 0, 0, 0,
		137, 130, 1, 0, 0, 0, 137, 131, 1, 0, 0, 0, 137, 132, 1, 0, 0, 0, 137,
		133, 1, 0, 0, 0, 137, 134, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136,
		1, 0, 0, 0, 138, 11, 1, 0, 0, 0, 139, 152, 3, 14, 7, 0, 140, 152, 3, 20,
		10, 0, 141, 152, 3, 18, 9, 0, 142, 152, 3, 22, 11, 0, 143, 152, 3, 34,
		17, 0, 144, 152, 3, 24, 12, 0, 145, 152, 3, 36, 18, 0, 146, 152, 3, 40,
		20, 0, 147, 152, 3, 42, 21, 0, 148, 152, 3, 44, 22, 0, 149, 152, 3, 46,
		23, 0, 150, 152, 3, 48, 24, 0, 151, 139, 1, 0, 0, 0, 151, 140, 1, 0, 0,
		0, 151, 141, 1, 0, 0, 0, 151, 142, 1, 0, 0, 0, 151, 143, 1, 0, 0, 0, 151,
		144, 1, 0, 0, 0, 151, 145, 1, 0, 0, 0, 151, 146, 1, 0, 0, 0, 151, 147,
		1, 0, 0, 0, 151, 148, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 150, 1, 0,
		0, 0, 152, 13, 1, 0, 0, 0, 153, 155, 3, 56, 28, 0, 154, 153, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 210, 5, 38, 0, 0, 157,
		159, 3, 56, 28, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 161, 5, 38, 0, 0, 161, 162, 5, 2, 0, 0, 162, 210, 5, 4,
		0, 0, 163, 165, 3, 56, 28, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 166, 171, 5, 38, 0, 0, 167, 168, 5, 2, 0, 0, 168,
		169, 3, 54, 27, 0, 169, 170, 5, 4, 0, 0, 170, 172, 1, 0, 0, 0, 171, 167,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 210, 1, 0, 0, 0, 173, 175, 3, 56,
		28, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 188, 5, 38, 0, 0, 177, 178, 5, 2, 0, 0, 178, 183, 3, 26, 13, 0, 179,
		180, 5, 3, 0, 0, 180, 182, 3, 26, 13, 0, 181, 179, 1, 0, 0, 0, 182, 185,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0,
		0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 4, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 177, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 210, 1, 0, 0, 0, 190,
		192, 3, 56, 28, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 207, 5, 38, 0, 0, 194, 195, 5, 2, 0, 0, 195, 196, 3, 54,
		27, 0, 196, 197, 5, 3, 0, 0, 197, 202, 3, 26, 13, 0, 198, 199, 5, 3, 0,
		0, 199, 201, 3, 26, 13, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0,
		202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 206, 5, 4, 0, 0, 206, 208, 1, 0, 0, 0, 207, 194,
		1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 210, 1, 0, 0, 0, 209, 154, 1, 0,
		0, 0, 209, 158, 1, 0, 0, 0, 209, 164, 1, 0, 0, 0, 209, 174, 1, 0, 0, 0,
		209, 191, 1, 0, 0, 0, 210, 15, 1, 0, 0, 0, 211, 214, 3, 12, 6, 0, 212,
		214, 3, 30, 15, 0, 213, 211, 1, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214, 220,
		1, 0, 0, 0, 215, 218, 5, 10, 0, 0, 216, 219, 3, 12, 6, 0, 217, 219, 3,
		30, 15, 0, 218, 216, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 221, 1, 0,
		0, 0, 220, 215, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 17, 1, 0, 0, 0, 224, 225, 5, 2, 0, 0, 225, 226,
		3, 16, 8, 0, 226, 227, 5, 4, 0, 0, 227, 19, 1, 0, 0, 0, 228, 229, 5, 11,
		0, 0, 229, 234, 3, 12, 6, 0, 230, 231, 5, 3, 0, 0, 231, 233, 3, 12, 6,
		0, 232, 230, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 239,
		5, 3, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0,
		0, 0, 240, 241, 5, 12, 0, 0, 241, 21, 1, 0, 0, 0, 242, 243, 5, 13, 0, 0,
		243, 244, 5, 2, 0, 0, 244, 245, 3, 60, 30, 0, 245, 246, 5, 4, 0, 0, 246,
		23, 1, 0, 0, 0, 247, 248, 5, 14, 0, 0, 248, 249, 5, 2, 0, 0, 249, 250,
		3, 10, 5, 0, 250, 251, 5, 3, 0, 0, 251, 252, 3, 60, 30, 0, 252, 253, 5,
		4, 0, 0, 253, 25, 1, 0, 0, 0, 254, 256, 3, 56, 28, 0, 255, 254, 1, 0, 0,
		0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 5, 15, 0, 0, 258,
		259, 5, 2, 0, 0, 259, 262, 3, 60, 30, 0, 260, 261, 5, 3, 0, 0, 261, 263,
		3, 54, 27, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1,
		0, 0, 0, 264, 265, 5, 4, 0, 0, 265, 27, 1, 0, 0, 0, 266, 267, 5, 16, 0,
		0, 267, 268, 5, 2, 0, 0, 268, 269, 5, 4, 0, 0, 269, 29, 1, 0, 0, 0, 270,
		271, 5, 17, 0, 0, 271, 272, 5, 2, 0, 0, 272, 273, 3, 14, 7, 0, 273, 274,
		5, 4, 0, 0, 274, 31, 1, 0, 0, 0, 275, 280, 3, 14, 7, 0, 276, 277, 5, 18,
		0, 0, 277, 279, 3, 14, 7, 0, 278, 276, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0,
		280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282,
		280, 1, 0, 0, 0, 283, 284, 5, 18, 0, 0, 284, 285, 3, 12, 6, 0, 285, 33,
		1, 0, 0, 0, 286, 287, 5, 2, 0, 0, 287, 288, 3, 32, 16, 0, 288, 289, 5,
		4, 0, 0, 289, 35, 1, 0, 0, 0, 290, 291, 5, 19, 0, 0, 291, 292, 5, 2, 0,
		0, 292, 293, 3, 14, 7, 0, 293, 294, 5, 4, 0, 0, 294, 295, 5, 5, 0, 0, 295,
		300, 3, 38, 19, 0, 296, 297, 5, 3, 0, 0, 297, 299, 3, 38, 19, 0, 298, 296,
		1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0,
		0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 305, 5, 3, 0, 0,
		304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306,
		307, 5, 6, 0, 0, 307, 37, 1, 0, 0, 0, 308, 309, 5, 20, 0, 0, 309, 310,
		3, 60, 30, 0, 310, 311, 5, 21, 0, 0, 311, 312, 3, 12, 6, 0, 312, 39, 1,
		0, 0, 0, 313, 314, 5, 22, 0, 0, 314, 315, 5, 2, 0, 0, 315, 316, 3, 54,
		27, 0, 316, 317, 5, 4, 0, 0, 317, 318, 5, 5, 0, 0, 318, 319, 3, 10, 5,
		0, 319, 320, 5, 6, 0, 0, 320, 41, 1, 0, 0, 0, 321, 322, 5, 23, 0, 0, 322,
		323, 5, 2, 0, 0, 323, 324, 3, 54, 27, 0, 324, 325, 5, 4, 0, 0, 325, 326,
		5, 5, 0, 0, 326, 327, 3, 10, 5, 0, 327, 328, 5, 6, 0, 0, 328, 43, 1, 0,
		0, 0, 329, 330, 5, 24, 0, 0, 330, 331, 5, 2, 0, 0, 331, 332, 3, 60, 30,
		0, 332, 333, 5, 4, 0, 0, 333, 334, 5, 5, 0, 0, 334, 335, 3, 10, 5, 0, 335,
		336, 5, 6, 0, 0, 336, 45, 1, 0, 0, 0, 337, 338, 5, 25, 0, 0, 338, 339,
		5, 2, 0, 0, 339, 340, 3, 14, 7, 0, 340, 341, 5, 4, 0, 0, 341, 47, 1, 0,
		0, 0, 342, 343, 5, 26, 0, 0, 343, 344, 5, 2, 0, 0, 344, 345, 5, 4, 0, 0,
		345, 49, 1, 0, 0, 0, 346, 347, 5, 27, 0, 0, 347, 362, 5, 2, 0, 0, 348,
		349, 3, 60, 30, 0, 349, 350, 5, 3, 0, 0, 350, 351, 3, 60, 30, 0, 351, 352,
		5, 3, 0, 0, 352, 353, 3, 60, 30, 0, 353, 354, 5, 3, 0, 0, 354, 355, 3,
		58, 29, 0, 355, 363, 1, 0, 0, 0, 356, 357, 3, 60, 30, 0, 357, 358, 5, 3,
		0, 0, 358, 359, 3, 60, 30, 0, 359, 360, 5, 3, 0, 0, 360, 361, 3, 58, 29,
		0, 361, 363, 1, 0, 0, 0, 362, 348, 1, 0, 0, 0, 362, 356, 1, 0, 0, 0, 363,
		364, 1, 0, 0, 0, 364, 365, 5, 4, 0, 0, 365, 51, 1, 0, 0, 0, 366, 367, 5,
		38, 0, 0, 367, 368, 5, 28, 0, 0, 368, 435, 3, 62, 31, 0, 369, 370, 5, 38,
		0, 0, 370, 371, 5, 28, 0, 0, 371, 435, 3, 58, 29, 0, 372, 373, 5, 38, 0,
		0, 373, 374, 5, 28, 0, 0, 374, 435, 3, 60, 30, 0, 375, 376, 5, 38, 0, 0,
		376, 377, 5, 28, 0, 0, 377, 435, 3, 64, 32, 0, 378, 379, 5, 38, 0, 0, 379,
		380, 5, 28, 0, 0, 380, 381, 5, 11, 0, 0, 381, 435, 5, 12, 0, 0, 382, 383,
		5, 38, 0, 0, 383, 384, 5, 28, 0, 0, 384, 385, 5, 11, 0, 0, 385, 390, 3,
		62, 31, 0, 386, 387, 5, 3, 0, 0, 387, 389, 3, 62, 31, 0, 388, 386, 1, 0,
		0, 0, 389, 392, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0,
		391, 393, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 394, 5, 12, 0, 0, 394,
		435, 1, 0, 0, 0, 395, 396, 5, 38, 0, 0, 396, 397, 5, 28, 0, 0, 397, 398,
		5, 11, 0, 0, 398, 403, 3, 58, 29, 0, 399, 400, 5, 3, 0, 0, 400, 402, 3,
		58, 29, 0, 401, 399, 1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401, 1, 0,
		0, 0, 403, 404, 1, 0, 0, 0, 404, 406, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0,
		406, 407, 5, 12, 0, 0, 407, 435, 1, 0, 0, 0, 408, 409, 5, 38, 0, 0, 409,
		410, 5, 28, 0, 0, 410, 411, 5, 11, 0, 0, 411, 416, 3, 60, 30, 0, 412, 413,
		5, 3, 0, 0, 413, 415, 3, 60, 30, 0, 414, 412, 1, 0, 0, 0, 415, 418, 1,
		0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 419, 1, 0, 0,
		0, 418, 416, 1, 0, 0, 0, 419, 420, 5, 12, 0, 0, 420, 435, 1, 0, 0, 0, 421,
		422, 5, 38, 0, 0, 422, 423, 5, 28, 0, 0, 423, 424, 5, 11, 0, 0, 424, 429,
		3, 64, 32, 0, 425, 426, 5, 3, 0, 0, 426, 428, 3, 64, 32, 0, 427, 425, 1,
		0, 0, 0, 428, 431, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0,
		0, 430, 432, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 433, 5, 12, 0, 0, 433,
		435, 1, 0, 0, 0, 434, 366, 1, 0, 0, 0, 434, 369, 1, 0, 0, 0, 434, 372,
		1, 0, 0, 0, 434, 375, 1, 0, 0, 0, 434, 378, 1, 0, 0, 0, 434, 382, 1, 0,
		0, 0, 434, 395, 1, 0, 0, 0, 434, 408, 1, 0, 0, 0, 434, 421, 1, 0, 0, 0,
		435, 53, 1, 0, 0, 0, 436, 441, 3, 52, 26, 0, 437, 438, 5, 3, 0, 0, 438,
		440, 3, 52, 26, 0, 439, 437, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 55, 1, 0, 0, 0, 443, 441, 1, 0,
		0, 0, 444, 445, 5, 29, 0, 0, 445, 57, 1, 0, 0, 0, 446, 448, 5, 30, 0, 0,
		447, 446, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449,
		450, 5, 39, 0, 0, 450, 59, 1, 0, 0, 0, 451, 452, 5, 40, 0, 0, 452, 61,
		1, 0, 0, 0, 453, 454, 7, 0, 0, 0, 454, 63, 1, 0, 0, 0, 455, 456, 5, 39,
		0, 0, 456, 457, 7, 1, 0, 0, 457, 65, 1, 0, 0, 0, 36, 70, 72, 80, 92, 113,
		137, 151, 154, 158, 164, 171, 174, 183, 188, 191, 202, 207, 209, 213, 218,
		222, 234, 238, 255, 262, 280, 300, 304, 362, 390, 403, 416, 429, 434, 441,
		447,
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

// gorchlangParserInit initializes any static state used to implement gorchlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgorchlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GorchlangParserInit() {
	staticData := &GorchlangParserStaticData
	staticData.once.Do(gorchlangParserInit)
}

// NewgorchlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgorchlangParser(input antlr.TokenStream) *gorchlangParser {
	GorchlangParserInit()
	this := new(gorchlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GorchlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gorchlang.g4"

	return this
}

// gorchlangParser tokens.
const (
	gorchlangParserEOF           = antlr.TokenEOF
	gorchlangParserT__0          = 1
	gorchlangParserT__1          = 2
	gorchlangParserT__2          = 3
	gorchlangParserT__3          = 4
	gorchlangParserT__4          = 5
	gorchlangParserT__5          = 6
	gorchlangParserT__6          = 7
	gorchlangParserT__7          = 8
	gorchlangParserT__8          = 9
	gorchlangParserT__9          = 10
	gorchlangParserT__10         = 11
	gorchlangParserT__11         = 12
	gorchlangParserT__12         = 13
	gorchlangParserT__13         = 14
	gorchlangParserT__14         = 15
	gorchlangParserT__15         = 16
	gorchlangParserT__16         = 17
	gorchlangParserT__17         = 18
	gorchlangParserT__18         = 19
	gorchlangParserT__19         = 20
	gorchlangParserT__20         = 21
	gorchlangParserT__21         = 22
	gorchlangParserT__22         = 23
	gorchlangParserT__23         = 24
	gorchlangParserT__24         = 25
	gorchlangParserT__25         = 26
	gorchlangParserT__26         = 27
	gorchlangParserT__27         = 28
	gorchlangParserT__28         = 29
	gorchlangParserT__29         = 30
	gorchlangParserT__30         = 31
	gorchlangParserT__31         = 32
	gorchlangParserT__32         = 33
	gorchlangParserT__33         = 34
	gorchlangParserT__34         = 35
	gorchlangParserTRUE          = 36
	gorchlangParserFALSE         = 37
	gorchlangParserID            = 38
	gorchlangParserINT           = 39
	gorchlangParserDQUOTA_STRING = 40
	gorchlangParserLINE_COMMENT  = 41
	gorchlangParserCOMMENT       = 42
	gorchlangParserTERMINATOR    = 43
	gorchlangParserWS            = 44
)

// gorchlangParser rules.
const (
	gorchlangParserRULE_primary                   = 0
	gorchlangParserRULE_startDirective            = 1
	gorchlangParserRULE_fragmentDirective         = 2
	gorchlangParserRULE_registerDirective         = 3
	gorchlangParserRULE_onFinishStmt              = 4
	gorchlangParserRULE_exedescStmt               = 5
	gorchlangParserRULE_leafSnippet               = 6
	gorchlangParserRULE_operatorStmt              = 7
	gorchlangParserRULE_serialStmt                = 8
	gorchlangParserRULE_serialBracketStmt         = 9
	gorchlangParserRULE_concurrentStmt            = 10
	gorchlangParserRULE_unfoldDirective           = 11
	gorchlangParserRULE_goDirective               = 12
	gorchlangParserRULE_waitDirective             = 13
	gorchlangParserRULE_noCheckMissDirective      = 14
	gorchlangParserRULE_skipDirective             = 15
	gorchlangParserRULE_wrapStmt                  = 16
	gorchlangParserRULE_wrapBracketStmt           = 17
	gorchlangParserRULE_switchDirective           = 18
	gorchlangParserRULE_switchCaseDirective       = 19
	gorchlangParserRULE_loopDirective             = 20
	gorchlangParserRULE_retryDirective            = 21
	gorchlangParserRULE_traceDirective            = 22
	gorchlangParserRULE_untilDirective            = 23
	gorchlangParserRULE_breakDirective            = 24
	gorchlangParserRULE_registerOperatorDirective = 25
	gorchlangParserRULE_argStmt                   = 26
	gorchlangParserRULE_argsStmt                  = 27
	gorchlangParserRULE_ignoreError               = 28
	gorchlangParserRULE_integerLiteral            = 29
	gorchlangParserRULE_stringLiteral             = 30
	gorchlangParserRULE_boolLiteral               = 31
	gorchlangParserRULE_durationLiteral           = 32
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
	p.RuleIndex = gorchlangParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_primary

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *gorchlangParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gorchlangParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&386) != 0 {
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(66)
				p.StartDirective()
			}

		case 2:
			{
				p.SetState(67)
				p.FragmentDirective()
			}

		case 3:
			{
				p.SetState(68)
				p.RegisterDirective()
			}

		case 4:
			{
				p.SetState(69)
				p.RegisterDirective()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(74)
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
	NoCheckMissDirective() INoCheckMissDirectiveContext
	OnFinishStmt() IOnFinishStmtContext

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
	p.RuleIndex = gorchlangParserRULE_startDirective
	return p
}

func InitEmptyStartDirectiveContext(p *StartDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_startDirective
}

func (*StartDirectiveContext) IsStartDirectiveContext() {}

func NewStartDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartDirectiveContext {
	var p = new(StartDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_startDirective

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

func (s *StartDirectiveContext) NoCheckMissDirective() INoCheckMissDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoCheckMissDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoCheckMissDirectiveContext)
}

func (s *StartDirectiveContext) OnFinishStmt() IOnFinishStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnFinishStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnFinishStmtContext)
}

func (s *StartDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterStartDirective(s)
	}
}

func (s *StartDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitStartDirective(s)
	}
}

func (p *gorchlangParser) StartDirective() (localctx IStartDirectiveContext) {
	localctx = NewStartDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gorchlangParserRULE_startDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(gorchlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(77)

		var _x = p.StringLiteral()

		localctx.(*StartDirectiveContext).name = _x
	}

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__2 {
		{
			p.SetState(78)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.ArgsStmt()
		}

	}
	{
		p.SetState(82)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(84)
			p.OnFinishStmt()
		}
		{
			p.SetState(85)
			p.NoCheckMissDirective()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(87)
			p.NoCheckMissDirective()
		}
		{
			p.SetState(88)
			p.OnFinishStmt()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(90)
			p.NoCheckMissDirective()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(91)
			p.OnFinishStmt()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(94)
		p.ExedescStmt()
	}
	{
		p.SetState(95)
		p.Match(gorchlangParserT__5)
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
	p.RuleIndex = gorchlangParserRULE_fragmentDirective
	return p
}

func InitEmptyFragmentDirectiveContext(p *FragmentDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_fragmentDirective
}

func (*FragmentDirectiveContext) IsFragmentDirectiveContext() {}

func NewFragmentDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDirectiveContext {
	var p = new(FragmentDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_fragmentDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterFragmentDirective(s)
	}
}

func (s *FragmentDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitFragmentDirective(s)
	}
}

func (p *gorchlangParser) FragmentDirective() (localctx IFragmentDirectiveContext) {
	localctx = NewFragmentDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gorchlangParserRULE_fragmentDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(gorchlangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(99)

		var _x = p.StringLiteral()

		localctx.(*FragmentDirectiveContext).name = _x
	}

	{
		p.SetState(100)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.ExedescStmt()
	}
	{
		p.SetState(103)
		p.Match(gorchlangParserT__5)
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
	p.RuleIndex = gorchlangParserRULE_registerDirective
	return p
}

func InitEmptyRegisterDirectiveContext(p *RegisterDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_registerDirective
}

func (*RegisterDirectiveContext) IsRegisterDirectiveContext() {}

func NewRegisterDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterDirectiveContext {
	var p = new(RegisterDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_registerDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterRegisterDirective(s)
	}
}

func (s *RegisterDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitRegisterDirective(s)
	}
}

func (p *gorchlangParser) RegisterDirective() (localctx IRegisterDirectiveContext) {
	localctx = NewRegisterDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gorchlangParserRULE_registerDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(gorchlangParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(107)

		var _x = p.StringLiteral()

		localctx.(*RegisterDirectiveContext).pkg = _x
	}

	{
		p.SetState(108)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gorchlangParserT__26 {
		{
			p.SetState(110)
			p.RegisterOperatorDirective()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Match(gorchlangParserT__5)
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
	p.RuleIndex = gorchlangParserRULE_onFinishStmt
	return p
}

func InitEmptyOnFinishStmtContext(p *OnFinishStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_onFinishStmt
}

func (*OnFinishStmtContext) IsOnFinishStmtContext() {}

func NewOnFinishStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnFinishStmtContext {
	var p = new(OnFinishStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_onFinishStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterOnFinishStmt(s)
	}
}

func (s *OnFinishStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitOnFinishStmt(s)
	}
}

func (p *gorchlangParser) OnFinishStmt() (localctx IOnFinishStmtContext) {
	localctx = NewOnFinishStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gorchlangParserRULE_onFinishStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(gorchlangParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.ExedescStmt()
	}
	{
		p.SetState(123)
		p.Match(gorchlangParserT__5)
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
	LoopDirective() ILoopDirectiveContext
	RetryDirective() IRetryDirectiveContext
	TraceDirective() ITraceDirectiveContext
	UntilDirective() IUntilDirectiveContext
	BreakDirective() IBreakDirectiveContext

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
	p.RuleIndex = gorchlangParserRULE_exedescStmt
	return p
}

func InitEmptyExedescStmtContext(p *ExedescStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_exedescStmt
}

func (*ExedescStmtContext) IsExedescStmtContext() {}

func NewExedescStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExedescStmtContext {
	var p = new(ExedescStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_exedescStmt

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

func (s *ExedescStmtContext) LoopDirective() ILoopDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopDirectiveContext)
}

func (s *ExedescStmtContext) RetryDirective() IRetryDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetryDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetryDirectiveContext)
}

func (s *ExedescStmtContext) TraceDirective() ITraceDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITraceDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITraceDirectiveContext)
}

func (s *ExedescStmtContext) UntilDirective() IUntilDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntilDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntilDirectiveContext)
}

func (s *ExedescStmtContext) BreakDirective() IBreakDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakDirectiveContext)
}

func (s *ExedescStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExedescStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExedescStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterExedescStmt(s)
	}
}

func (s *ExedescStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitExedescStmt(s)
	}
}

func (p *gorchlangParser) ExedescStmt() (localctx IExedescStmtContext) {
	localctx = NewExedescStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gorchlangParserRULE_exedescStmt)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.SerialStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.ConcurrentStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.WrapStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(128)
			p.WrapBracketStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(129)
			p.OperatorStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(130)
			p.SwitchDirective()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(131)
			p.UnfoldDirective()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(132)
			p.LoopDirective()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(133)
			p.RetryDirective()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(134)
			p.TraceDirective()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(135)
			p.UntilDirective()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(136)
			p.BreakDirective()
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
	LoopDirective() ILoopDirectiveContext
	RetryDirective() IRetryDirectiveContext
	TraceDirective() ITraceDirectiveContext
	UntilDirective() IUntilDirectiveContext
	BreakDirective() IBreakDirectiveContext

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
	p.RuleIndex = gorchlangParserRULE_leafSnippet
	return p
}

func InitEmptyLeafSnippetContext(p *LeafSnippetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_leafSnippet
}

func (*LeafSnippetContext) IsLeafSnippetContext() {}

func NewLeafSnippetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeafSnippetContext {
	var p = new(LeafSnippetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_leafSnippet

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

func (s *LeafSnippetContext) LoopDirective() ILoopDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopDirectiveContext)
}

func (s *LeafSnippetContext) RetryDirective() IRetryDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetryDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetryDirectiveContext)
}

func (s *LeafSnippetContext) TraceDirective() ITraceDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITraceDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITraceDirectiveContext)
}

func (s *LeafSnippetContext) UntilDirective() IUntilDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntilDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntilDirectiveContext)
}

func (s *LeafSnippetContext) BreakDirective() IBreakDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakDirectiveContext)
}

func (s *LeafSnippetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeafSnippetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeafSnippetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterLeafSnippet(s)
	}
}

func (s *LeafSnippetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitLeafSnippet(s)
	}
}

func (p *gorchlangParser) LeafSnippet() (localctx ILeafSnippetContext) {
	localctx = NewLeafSnippetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gorchlangParserRULE_leafSnippet)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.OperatorStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.ConcurrentStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.SerialBracketStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.UnfoldDirective()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(143)
			p.WrapBracketStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(144)
			p.GoDirective()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(145)
			p.SwitchDirective()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(146)
			p.LoopDirective()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(147)
			p.RetryDirective()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(148)
			p.TraceDirective()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(149)
			p.UntilDirective()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(150)
			p.BreakDirective()
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
	p.RuleIndex = gorchlangParserRULE_operatorStmt
	return p
}

func InitEmptyOperatorStmtContext(p *OperatorStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_operatorStmt
}

func (*OperatorStmtContext) IsOperatorStmtContext() {}

func NewOperatorStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStmtContext {
	var p = new(OperatorStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_operatorStmt

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
	return s.GetToken(gorchlangParserID, 0)
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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterOperatorStmt(s)
	}
}

func (s *OperatorStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitOperatorStmt(s)
	}
}

func (p *gorchlangParser) OperatorStmt() (localctx IOperatorStmtContext) {
	localctx = NewOperatorStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gorchlangParserRULE_operatorStmt)
	var _la int

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__28 {
			{
				p.SetState(153)
				p.IgnoreError()
			}

		}

		{
			p.SetState(156)

			var _m = p.Match(gorchlangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__28 {
			{
				p.SetState(157)
				p.IgnoreError()
			}

		}

		{
			p.SetState(160)

			var _m = p.Match(gorchlangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(161)
			p.Match(gorchlangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(gorchlangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__28 {
			{
				p.SetState(163)
				p.IgnoreError()
			}

		}

		{
			p.SetState(166)

			var _m = p.Match(gorchlangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__1 {
			{
				p.SetState(167)
				p.Match(gorchlangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(168)
				p.ArgsStmt()
			}
			{
				p.SetState(169)
				p.Match(gorchlangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__28 {
			{
				p.SetState(173)
				p.IgnoreError()
			}

		}

		{
			p.SetState(176)

			var _m = p.Match(gorchlangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__1 {
			{
				p.SetState(177)
				p.Match(gorchlangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(178)
				p.WaitDirective()
			}
			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == gorchlangParserT__2 {
				{
					p.SetState(179)
					p.Match(gorchlangParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(180)
					p.WaitDirective()
				}

				p.SetState(185)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(186)
				p.Match(gorchlangParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__28 {
			{
				p.SetState(190)
				p.IgnoreError()
			}

		}

		{
			p.SetState(193)

			var _m = p.Match(gorchlangParserID)

			localctx.(*OperatorStmtContext).operatorName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gorchlangParserT__1 {
			{
				p.SetState(194)
				p.Match(gorchlangParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(195)
				p.ArgsStmt()
			}
			{
				p.SetState(196)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(197)
				p.WaitDirective()
			}
			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == gorchlangParserT__2 {
				{
					p.SetState(198)
					p.Match(gorchlangParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(199)
					p.WaitDirective()
				}

				p.SetState(204)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(205)
				p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_serialStmt
	return p
}

func InitEmptySerialStmtContext(p *SerialStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_serialStmt
}

func (*SerialStmtContext) IsSerialStmtContext() {}

func NewSerialStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SerialStmtContext {
	var p = new(SerialStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_serialStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterSerialStmt(s)
	}
}

func (s *SerialStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitSerialStmt(s)
	}
}

func (p *gorchlangParser) SerialStmt() (localctx ISerialStmtContext) {
	localctx = NewSerialStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gorchlangParserRULE_serialStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gorchlangParserT__1, gorchlangParserT__10, gorchlangParserT__12, gorchlangParserT__13, gorchlangParserT__18, gorchlangParserT__21, gorchlangParserT__22, gorchlangParserT__23, gorchlangParserT__24, gorchlangParserT__25, gorchlangParserT__28, gorchlangParserID:
		{
			p.SetState(211)
			p.LeafSnippet()
		}

	case gorchlangParserT__16:
		{
			p.SetState(212)
			p.SkipDirective()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gorchlangParserT__9 {
		{
			p.SetState(215)
			p.Match(gorchlangParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case gorchlangParserT__1, gorchlangParserT__10, gorchlangParserT__12, gorchlangParserT__13, gorchlangParserT__18, gorchlangParserT__21, gorchlangParserT__22, gorchlangParserT__23, gorchlangParserT__24, gorchlangParserT__25, gorchlangParserT__28, gorchlangParserID:
			{
				p.SetState(216)
				p.LeafSnippet()
			}

		case gorchlangParserT__16:
			{
				p.SetState(217)
				p.SkipDirective()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(222)
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
	p.RuleIndex = gorchlangParserRULE_serialBracketStmt
	return p
}

func InitEmptySerialBracketStmtContext(p *SerialBracketStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_serialBracketStmt
}

func (*SerialBracketStmtContext) IsSerialBracketStmtContext() {}

func NewSerialBracketStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SerialBracketStmtContext {
	var p = new(SerialBracketStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_serialBracketStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterSerialBracketStmt(s)
	}
}

func (s *SerialBracketStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitSerialBracketStmt(s)
	}
}

func (p *gorchlangParser) SerialBracketStmt() (localctx ISerialBracketStmtContext) {
	localctx = NewSerialBracketStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gorchlangParserRULE_serialBracketStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.SerialStmt()
	}
	{
		p.SetState(226)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_concurrentStmt
	return p
}

func InitEmptyConcurrentStmtContext(p *ConcurrentStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_concurrentStmt
}

func (*ConcurrentStmtContext) IsConcurrentStmtContext() {}

func NewConcurrentStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcurrentStmtContext {
	var p = new(ConcurrentStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_concurrentStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterConcurrentStmt(s)
	}
}

func (s *ConcurrentStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitConcurrentStmt(s)
	}
}

func (p *gorchlangParser) ConcurrentStmt() (localctx IConcurrentStmtContext) {
	localctx = NewConcurrentStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gorchlangParserRULE_concurrentStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(gorchlangParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.LeafSnippet()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.LeafSnippet()
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__2 {
		{
			p.SetState(237)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(240)
		p.Match(gorchlangParserT__11)
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
	p.RuleIndex = gorchlangParserRULE_unfoldDirective
	return p
}

func InitEmptyUnfoldDirectiveContext(p *UnfoldDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_unfoldDirective
}

func (*UnfoldDirectiveContext) IsUnfoldDirectiveContext() {}

func NewUnfoldDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnfoldDirectiveContext {
	var p = new(UnfoldDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_unfoldDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterUnfoldDirective(s)
	}
}

func (s *UnfoldDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitUnfoldDirective(s)
	}
}

func (p *gorchlangParser) UnfoldDirective() (localctx IUnfoldDirectiveContext) {
	localctx = NewUnfoldDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gorchlangParserRULE_unfoldDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(gorchlangParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(244)

		var _x = p.StringLiteral()

		localctx.(*UnfoldDirectiveContext).name = _x
	}

	{
		p.SetState(245)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_goDirective
	return p
}

func InitEmptyGoDirectiveContext(p *GoDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_goDirective
}

func (*GoDirectiveContext) IsGoDirectiveContext() {}

func NewGoDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoDirectiveContext {
	var p = new(GoDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_goDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterGoDirective(s)
	}
}

func (s *GoDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitGoDirective(s)
	}
}

func (p *gorchlangParser) GoDirective() (localctx IGoDirectiveContext) {
	localctx = NewGoDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gorchlangParserRULE_goDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(gorchlangParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.ExedescStmt()
	}
	{
		p.SetState(250)
		p.Match(gorchlangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(251)

		var _x = p.StringLiteral()

		localctx.(*GoDirectiveContext).name = _x
	}

	{
		p.SetState(252)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_waitDirective
	return p
}

func InitEmptyWaitDirectiveContext(p *WaitDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_waitDirective
}

func (*WaitDirectiveContext) IsWaitDirectiveContext() {}

func NewWaitDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaitDirectiveContext {
	var p = new(WaitDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_waitDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterWaitDirective(s)
	}
}

func (s *WaitDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitWaitDirective(s)
	}
}

func (p *gorchlangParser) WaitDirective() (localctx IWaitDirectiveContext) {
	localctx = NewWaitDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gorchlangParserRULE_waitDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__28 {
		{
			p.SetState(254)
			p.IgnoreError()
		}

	}
	{
		p.SetState(257)
		p.Match(gorchlangParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(259)

		var _x = p.StringLiteral()

		localctx.(*WaitDirectiveContext).name = _x
	}

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__2 {
		{
			p.SetState(260)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.ArgsStmt()
		}

	}
	{
		p.SetState(264)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_noCheckMissDirective
	return p
}

func InitEmptyNoCheckMissDirectiveContext(p *NoCheckMissDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_noCheckMissDirective
}

func (*NoCheckMissDirectiveContext) IsNoCheckMissDirectiveContext() {}

func NewNoCheckMissDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoCheckMissDirectiveContext {
	var p = new(NoCheckMissDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_noCheckMissDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterNoCheckMissDirective(s)
	}
}

func (s *NoCheckMissDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitNoCheckMissDirective(s)
	}
}

func (p *gorchlangParser) NoCheckMissDirective() (localctx INoCheckMissDirectiveContext) {
	localctx = NewNoCheckMissDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gorchlangParserRULE_noCheckMissDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(gorchlangParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_skipDirective
	return p
}

func InitEmptySkipDirectiveContext(p *SkipDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_skipDirective
}

func (*SkipDirectiveContext) IsSkipDirectiveContext() {}

func NewSkipDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipDirectiveContext {
	var p = new(SkipDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_skipDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterSkipDirective(s)
	}
}

func (s *SkipDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitSkipDirective(s)
	}
}

func (p *gorchlangParser) SkipDirective() (localctx ISkipDirectiveContext) {
	localctx = NewSkipDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gorchlangParserRULE_skipDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(gorchlangParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.OperatorStmt()
	}
	{
		p.SetState(273)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_wrapStmt
	return p
}

func InitEmptyWrapStmtContext(p *WrapStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_wrapStmt
}

func (*WrapStmtContext) IsWrapStmtContext() {}

func NewWrapStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WrapStmtContext {
	var p = new(WrapStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_wrapStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterWrapStmt(s)
	}
}

func (s *WrapStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitWrapStmt(s)
	}
}

func (p *gorchlangParser) WrapStmt() (localctx IWrapStmtContext) {
	localctx = NewWrapStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gorchlangParserRULE_wrapStmt)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.OperatorStmt()
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(276)
				p.Match(gorchlangParserT__17)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(277)
				p.OperatorStmt()
			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Match(gorchlangParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
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
	p.RuleIndex = gorchlangParserRULE_wrapBracketStmt
	return p
}

func InitEmptyWrapBracketStmtContext(p *WrapBracketStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_wrapBracketStmt
}

func (*WrapBracketStmtContext) IsWrapBracketStmtContext() {}

func NewWrapBracketStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WrapBracketStmtContext {
	var p = new(WrapBracketStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_wrapBracketStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterWrapBracketStmt(s)
	}
}

func (s *WrapBracketStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitWrapBracketStmt(s)
	}
}

func (p *gorchlangParser) WrapBracketStmt() (localctx IWrapBracketStmtContext) {
	localctx = NewWrapBracketStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gorchlangParserRULE_wrapBracketStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.WrapStmt()
	}
	{
		p.SetState(288)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_switchDirective
	return p
}

func InitEmptySwitchDirectiveContext(p *SwitchDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_switchDirective
}

func (*SwitchDirectiveContext) IsSwitchDirectiveContext() {}

func NewSwitchDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDirectiveContext {
	var p = new(SwitchDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_switchDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterSwitchDirective(s)
	}
}

func (s *SwitchDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitSwitchDirective(s)
	}
}

func (p *gorchlangParser) SwitchDirective() (localctx ISwitchDirectiveContext) {
	localctx = NewSwitchDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gorchlangParserRULE_switchDirective)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(gorchlangParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.OperatorStmt()
	}
	{
		p.SetState(293)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.SwitchCaseDirective()
	}
	p.SetState(300)
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
				p.SetState(296)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(297)
				p.SwitchCaseDirective()
			}

		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__2 {
		{
			p.SetState(303)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(306)
		p.Match(gorchlangParserT__5)
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
	p.RuleIndex = gorchlangParserRULE_switchCaseDirective
	return p
}

func InitEmptySwitchCaseDirectiveContext(p *SwitchCaseDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_switchCaseDirective
}

func (*SwitchCaseDirectiveContext) IsSwitchCaseDirectiveContext() {}

func NewSwitchCaseDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseDirectiveContext {
	var p = new(SwitchCaseDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_switchCaseDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterSwitchCaseDirective(s)
	}
}

func (s *SwitchCaseDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitSwitchCaseDirective(s)
	}
}

func (p *gorchlangParser) SwitchCaseDirective() (localctx ISwitchCaseDirectiveContext) {
	localctx = NewSwitchCaseDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gorchlangParserRULE_switchCaseDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(gorchlangParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(309)

		var _x = p.StringLiteral()

		localctx.(*SwitchCaseDirectiveContext).caseName = _x
	}

	{
		p.SetState(310)
		p.Match(gorchlangParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
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

// ILoopDirectiveContext is an interface to support dynamic dispatch.
type ILoopDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArgsStmt() IArgsStmtContext
	ExedescStmt() IExedescStmtContext

	// IsLoopDirectiveContext differentiates from other interfaces.
	IsLoopDirectiveContext()
}

type LoopDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopDirectiveContext() *LoopDirectiveContext {
	var p = new(LoopDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_loopDirective
	return p
}

func InitEmptyLoopDirectiveContext(p *LoopDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_loopDirective
}

func (*LoopDirectiveContext) IsLoopDirectiveContext() {}

func NewLoopDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopDirectiveContext {
	var p = new(LoopDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_loopDirective

	return p
}

func (s *LoopDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopDirectiveContext) ArgsStmt() IArgsStmtContext {
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

func (s *LoopDirectiveContext) ExedescStmt() IExedescStmtContext {
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

func (s *LoopDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterLoopDirective(s)
	}
}

func (s *LoopDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitLoopDirective(s)
	}
}

func (p *gorchlangParser) LoopDirective() (localctx ILoopDirectiveContext) {
	localctx = NewLoopDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gorchlangParserRULE_loopDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(gorchlangParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.ArgsStmt()
	}
	{
		p.SetState(316)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.ExedescStmt()
	}
	{
		p.SetState(319)
		p.Match(gorchlangParserT__5)
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

// IRetryDirectiveContext is an interface to support dynamic dispatch.
type IRetryDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArgsStmt() IArgsStmtContext
	ExedescStmt() IExedescStmtContext

	// IsRetryDirectiveContext differentiates from other interfaces.
	IsRetryDirectiveContext()
}

type RetryDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetryDirectiveContext() *RetryDirectiveContext {
	var p = new(RetryDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_retryDirective
	return p
}

func InitEmptyRetryDirectiveContext(p *RetryDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_retryDirective
}

func (*RetryDirectiveContext) IsRetryDirectiveContext() {}

func NewRetryDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetryDirectiveContext {
	var p = new(RetryDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_retryDirective

	return p
}

func (s *RetryDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *RetryDirectiveContext) ArgsStmt() IArgsStmtContext {
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

func (s *RetryDirectiveContext) ExedescStmt() IExedescStmtContext {
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

func (s *RetryDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetryDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetryDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterRetryDirective(s)
	}
}

func (s *RetryDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitRetryDirective(s)
	}
}

func (p *gorchlangParser) RetryDirective() (localctx IRetryDirectiveContext) {
	localctx = NewRetryDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gorchlangParserRULE_retryDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(gorchlangParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.ArgsStmt()
	}
	{
		p.SetState(324)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.ExedescStmt()
	}
	{
		p.SetState(327)
		p.Match(gorchlangParserT__5)
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

// ITraceDirectiveContext is an interface to support dynamic dispatch.
type ITraceDirectiveContext interface {
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

	// IsTraceDirectiveContext differentiates from other interfaces.
	IsTraceDirectiveContext()
}

type TraceDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IStringLiteralContext
}

func NewEmptyTraceDirectiveContext() *TraceDirectiveContext {
	var p = new(TraceDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_traceDirective
	return p
}

func InitEmptyTraceDirectiveContext(p *TraceDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_traceDirective
}

func (*TraceDirectiveContext) IsTraceDirectiveContext() {}

func NewTraceDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TraceDirectiveContext {
	var p = new(TraceDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_traceDirective

	return p
}

func (s *TraceDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *TraceDirectiveContext) GetName() IStringLiteralContext { return s.name }

func (s *TraceDirectiveContext) SetName(v IStringLiteralContext) { s.name = v }

func (s *TraceDirectiveContext) ExedescStmt() IExedescStmtContext {
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

func (s *TraceDirectiveContext) StringLiteral() IStringLiteralContext {
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

func (s *TraceDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TraceDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TraceDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterTraceDirective(s)
	}
}

func (s *TraceDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitTraceDirective(s)
	}
}

func (p *gorchlangParser) TraceDirective() (localctx ITraceDirectiveContext) {
	localctx = NewTraceDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gorchlangParserRULE_traceDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(gorchlangParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(331)

		var _x = p.StringLiteral()

		localctx.(*TraceDirectiveContext).name = _x
	}

	{
		p.SetState(332)
		p.Match(gorchlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Match(gorchlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.ExedescStmt()
	}
	{
		p.SetState(335)
		p.Match(gorchlangParserT__5)
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

// IUntilDirectiveContext is an interface to support dynamic dispatch.
type IUntilDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperatorStmt() IOperatorStmtContext

	// IsUntilDirectiveContext differentiates from other interfaces.
	IsUntilDirectiveContext()
}

type UntilDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntilDirectiveContext() *UntilDirectiveContext {
	var p = new(UntilDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_untilDirective
	return p
}

func InitEmptyUntilDirectiveContext(p *UntilDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_untilDirective
}

func (*UntilDirectiveContext) IsUntilDirectiveContext() {}

func NewUntilDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UntilDirectiveContext {
	var p = new(UntilDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_untilDirective

	return p
}

func (s *UntilDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *UntilDirectiveContext) OperatorStmt() IOperatorStmtContext {
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

func (s *UntilDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntilDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UntilDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterUntilDirective(s)
	}
}

func (s *UntilDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitUntilDirective(s)
	}
}

func (p *gorchlangParser) UntilDirective() (localctx IUntilDirectiveContext) {
	localctx = NewUntilDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gorchlangParserRULE_untilDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(gorchlangParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.OperatorStmt()
	}
	{
		p.SetState(340)
		p.Match(gorchlangParserT__3)
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

// IBreakDirectiveContext is an interface to support dynamic dispatch.
type IBreakDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakDirectiveContext differentiates from other interfaces.
	IsBreakDirectiveContext()
}

type BreakDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakDirectiveContext() *BreakDirectiveContext {
	var p = new(BreakDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_breakDirective
	return p
}

func InitEmptyBreakDirectiveContext(p *BreakDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_breakDirective
}

func (*BreakDirectiveContext) IsBreakDirectiveContext() {}

func NewBreakDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakDirectiveContext {
	var p = new(BreakDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_breakDirective

	return p
}

func (s *BreakDirectiveContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterBreakDirective(s)
	}
}

func (s *BreakDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitBreakDirective(s)
	}
}

func (p *gorchlangParser) BreakDirective() (localctx IBreakDirectiveContext) {
	localctx = NewBreakDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gorchlangParserRULE_breakDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(gorchlangParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_registerOperatorDirective
	return p
}

func InitEmptyRegisterOperatorDirectiveContext(p *RegisterOperatorDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_registerOperatorDirective
}

func (*RegisterOperatorDirectiveContext) IsRegisterOperatorDirectiveContext() {}

func NewRegisterOperatorDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterOperatorDirectiveContext {
	var p = new(RegisterOperatorDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_registerOperatorDirective

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterRegisterOperatorDirective(s)
	}
}

func (s *RegisterOperatorDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitRegisterOperatorDirective(s)
	}
}

func (p *gorchlangParser) RegisterOperatorDirective() (localctx IRegisterOperatorDirectiveContext) {
	localctx = NewRegisterOperatorDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gorchlangParserRULE_registerOperatorDirective)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(gorchlangParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Match(gorchlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(348)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).pkgPath = _x
		}

		{
			p.SetState(349)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(350)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).structName = _x
		}

		{
			p.SetState(351)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(352)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).operatorName = _x
		}

		{
			p.SetState(353)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(354)

			var _x = p.IntegerLiteral()

			localctx.(*RegisterOperatorDirectiveContext).seq = _x
		}

	case 2:
		{
			p.SetState(356)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).pkgPath = _x
		}

		{
			p.SetState(357)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(358)

			var _x = p.StringLiteral()

			localctx.(*RegisterOperatorDirectiveContext).structName = _x
		}

		{
			p.SetState(359)
			p.Match(gorchlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(360)

			var _x = p.IntegerLiteral()

			localctx.(*RegisterOperatorDirectiveContext).seq = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(364)
		p.Match(gorchlangParserT__3)
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
	p.RuleIndex = gorchlangParserRULE_argStmt
	return p
}

func InitEmptyArgStmtContext(p *ArgStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_argStmt
}

func (*ArgStmtContext) IsArgStmtContext() {}

func NewArgStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgStmtContext {
	var p = new(ArgStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_argStmt

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
	return s.GetToken(gorchlangParserID, 0)
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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterArgStmt(s)
	}
}

func (s *ArgStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitArgStmt(s)
	}
}

func (p *gorchlangParser) ArgStmt() (localctx IArgStmtContext) {
	localctx = NewArgStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gorchlangParserRULE_argStmt)
	var _la int

	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(367)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.BoolLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(370)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(372)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(373)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.StringLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(376)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.DurationLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(378)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(379)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(gorchlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(gorchlangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(382)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(383)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.Match(gorchlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.BoolLiteral()
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gorchlangParserT__2 {
			{
				p.SetState(386)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(387)
				p.BoolLiteral()
			}

			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(393)
			p.Match(gorchlangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(395)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(396)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(gorchlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.IntegerLiteral()
		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gorchlangParserT__2 {
			{
				p.SetState(399)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(400)
				p.IntegerLiteral()
			}

			p.SetState(405)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(406)
			p.Match(gorchlangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(408)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(409)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(gorchlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.StringLiteral()
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gorchlangParserT__2 {
			{
				p.SetState(412)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(413)
				p.StringLiteral()
			}

			p.SetState(418)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(419)
			p.Match(gorchlangParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(421)

			var _m = p.Match(gorchlangParserID)

			localctx.(*ArgStmtContext).key = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(422)
			p.Match(gorchlangParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(gorchlangParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.DurationLiteral()
		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gorchlangParserT__2 {
			{
				p.SetState(425)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(426)
				p.DurationLiteral()
			}

			p.SetState(431)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(432)
			p.Match(gorchlangParserT__11)
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
	p.RuleIndex = gorchlangParserRULE_argsStmt
	return p
}

func InitEmptyArgsStmtContext(p *ArgsStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_argsStmt
}

func (*ArgsStmtContext) IsArgsStmtContext() {}

func NewArgsStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsStmtContext {
	var p = new(ArgsStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_argsStmt

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterArgsStmt(s)
	}
}

func (s *ArgsStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitArgsStmt(s)
	}
}

func (p *gorchlangParser) ArgsStmt() (localctx IArgsStmtContext) {
	localctx = NewArgsStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gorchlangParserRULE_argsStmt)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.ArgStmt()
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(437)
				p.Match(gorchlangParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(438)
				p.ArgStmt()
			}

		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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
	p.RuleIndex = gorchlangParserRULE_ignoreError
	return p
}

func InitEmptyIgnoreErrorContext(p *IgnoreErrorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_ignoreError
}

func (*IgnoreErrorContext) IsIgnoreErrorContext() {}

func NewIgnoreErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoreErrorContext {
	var p = new(IgnoreErrorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_ignoreError

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
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterIgnoreError(s)
	}
}

func (s *IgnoreErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitIgnoreError(s)
	}
}

func (p *gorchlangParser) IgnoreError() (localctx IIgnoreErrorContext) {
	localctx = NewIgnoreErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gorchlangParserRULE_ignoreError)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(gorchlangParserT__28)
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
	p.RuleIndex = gorchlangParserRULE_integerLiteral
	return p
}

func InitEmptyIntegerLiteralContext(p *IntegerLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_integerLiteral
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(gorchlangParserINT, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *gorchlangParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gorchlangParserRULE_integerLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gorchlangParserT__29 {
		{
			p.SetState(446)
			p.Match(gorchlangParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(449)
		p.Match(gorchlangParserINT)
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
	p.RuleIndex = gorchlangParserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(gorchlangParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *gorchlangParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gorchlangParserRULE_stringLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(gorchlangParserDQUOTA_STRING)
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
	p.RuleIndex = gorchlangParserRULE_boolLiteral
	return p
}

func InitEmptyBoolLiteralContext(p *BoolLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_boolLiteral
}

func (*BoolLiteralContext) IsBoolLiteralContext() {}

func NewBoolLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_boolLiteral

	return p
}

func (s *BoolLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gorchlangParserTRUE, 0)
}

func (s *BoolLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gorchlangParserFALSE, 0)
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (p *gorchlangParser) BoolLiteral() (localctx IBoolLiteralContext) {
	localctx = NewBoolLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gorchlangParserRULE_boolLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gorchlangParserTRUE || _la == gorchlangParserFALSE) {
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
	p.RuleIndex = gorchlangParserRULE_durationLiteral
	return p
}

func InitEmptyDurationLiteralContext(p *DurationLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gorchlangParserRULE_durationLiteral
}

func (*DurationLiteralContext) IsDurationLiteralContext() {}

func NewDurationLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationLiteralContext {
	var p = new(DurationLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gorchlangParserRULE_durationLiteral

	return p
}

func (s *DurationLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(gorchlangParserINT, 0)
}

func (s *DurationLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.EnterDurationLiteral(s)
	}
}

func (s *DurationLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gorchlangListener); ok {
		listenerT.ExitDurationLiteral(s)
	}
}

func (p *gorchlangParser) DurationLiteral() (localctx IDurationLiteralContext) {
	localctx = NewDurationLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gorchlangParserRULE_durationLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(gorchlangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66571993088) != 0) {
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
