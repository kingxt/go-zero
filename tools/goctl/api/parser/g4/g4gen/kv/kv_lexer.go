// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/KVLexer.g4 by ANTLR 4.9. DO NOT EDIT.

package kv

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 150,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 41, 10, 5, 13, 5, 14, 5, 42, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 51, 10, 6, 12, 6, 14, 6, 54, 11, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 65, 10, 7, 12,
	7, 14, 7, 68, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 75, 10, 8, 12,
	8, 14, 8, 78, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 84, 10, 9, 12, 9, 14,
	9, 87, 11, 9, 3, 10, 3, 10, 5, 10, 91, 10, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 99, 10, 11, 3, 11, 5, 11, 102, 10, 11, 3, 11,
	3, 11, 3, 11, 6, 11, 107, 10, 11, 13, 11, 14, 11, 108, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 116, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12, 121,
	10, 12, 12, 12, 14, 12, 124, 11, 12, 3, 12, 5, 12, 127, 10, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 7, 14, 133, 10, 14, 12, 14, 14, 14, 136, 11, 14, 3,
	14, 5, 14, 139, 10, 14, 3, 15, 3, 15, 5, 15, 143, 10, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 149, 10, 16, 3, 52, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2,
	31, 2, 3, 2, 18, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 4,
	2, 36, 36, 94, 94, 9, 2, 12, 12, 15, 15, 36, 36, 41, 44, 49, 49, 60, 60,
	98, 98, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 10, 2, 36, 36, 41,
	41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 3, 2, 50,
	53, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 3, 2, 50, 59, 4, 2, 50,
	59, 97, 97, 6, 2, 38, 38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298,
	56321, 3, 2, 55298, 56321, 3, 2, 56322, 57345, 2, 163, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 33, 3, 2, 2,
	2, 5, 35, 3, 2, 2, 2, 7, 37, 3, 2, 2, 2, 9, 40, 3, 2, 2, 2, 11, 46, 3,
	2, 2, 2, 13, 60, 3, 2, 2, 2, 15, 71, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2, 19,
	88, 3, 2, 2, 2, 21, 115, 3, 2, 2, 2, 23, 117, 3, 2, 2, 2, 25, 128, 3, 2,
	2, 2, 27, 130, 3, 2, 2, 2, 29, 142, 3, 2, 2, 2, 31, 148, 3, 2, 2, 2, 33,
	34, 7, 42, 2, 2, 34, 4, 3, 2, 2, 2, 35, 36, 7, 43, 2, 2, 36, 6, 3, 2, 2,
	2, 37, 38, 7, 60, 2, 2, 38, 8, 3, 2, 2, 2, 39, 41, 9, 2, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 44, 3, 2, 2, 2, 44, 45, 8, 5, 2, 2, 45, 10, 3, 2, 2, 2, 46, 47, 7,
	49, 2, 2, 47, 48, 7, 44, 2, 2, 48, 52, 3, 2, 2, 2, 49, 51, 11, 2, 2, 2,
	50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 52, 50, 3,
	2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 7, 44, 2, 2, 56,
	57, 7, 49, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 8, 6, 2, 2, 59, 12, 3, 2,
	2, 2, 60, 61, 7, 49, 2, 2, 61, 62, 7, 49, 2, 2, 62, 66, 3, 2, 2, 2, 63,
	65, 10, 3, 2, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2,
	2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70,
	8, 7, 2, 2, 70, 14, 3, 2, 2, 2, 71, 76, 7, 36, 2, 2, 72, 75, 10, 4, 2,
	2, 73, 75, 5, 21, 11, 2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 78,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2,
	78, 76, 3, 2, 2, 2, 79, 80, 7, 36, 2, 2, 80, 16, 3, 2, 2, 2, 81, 84, 10,
	5, 2, 2, 82, 84, 5, 21, 11, 2, 83, 81, 3, 2, 2, 2, 83, 82, 3, 2, 2, 2,
	84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 18, 3,
	2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 90, 9, 6, 2, 2, 89, 91, 9, 7, 2, 2, 90,
	89, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 5, 27,
	14, 2, 93, 20, 3, 2, 2, 2, 94, 95, 7, 94, 2, 2, 95, 116, 9, 8, 2, 2, 96,
	101, 7, 94, 2, 2, 97, 99, 9, 9, 2, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2,
	2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 9, 10, 2, 2, 101, 98, 3, 2, 2, 2,
	101, 102, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 116, 9, 10, 2, 2, 104,
	106, 7, 94, 2, 2, 105, 107, 7, 119, 2, 2, 106, 105, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 110, 111, 5, 25, 13, 2, 111, 112, 5, 25, 13, 2, 112, 113, 5, 25,
	13, 2, 113, 114, 5, 25, 13, 2, 114, 116, 3, 2, 2, 2, 115, 94, 3, 2, 2,
	2, 115, 96, 3, 2, 2, 2, 115, 104, 3, 2, 2, 2, 116, 22, 3, 2, 2, 2, 117,
	126, 5, 25, 13, 2, 118, 121, 5, 25, 13, 2, 119, 121, 7, 97, 2, 2, 120,
	118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120,
	3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122, 3, 2,
	2, 2, 125, 127, 5, 25, 13, 2, 126, 122, 3, 2, 2, 2, 126, 127, 3, 2, 2,
	2, 127, 24, 3, 2, 2, 2, 128, 129, 9, 11, 2, 2, 129, 26, 3, 2, 2, 2, 130,
	138, 9, 12, 2, 2, 131, 133, 9, 13, 2, 2, 132, 131, 3, 2, 2, 2, 133, 136,
	3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2,
	2, 2, 136, 134, 3, 2, 2, 2, 137, 139, 9, 12, 2, 2, 138, 134, 3, 2, 2, 2,
	138, 139, 3, 2, 2, 2, 139, 28, 3, 2, 2, 2, 140, 143, 5, 31, 16, 2, 141,
	143, 9, 12, 2, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 30,
	3, 2, 2, 2, 144, 149, 9, 14, 2, 2, 145, 149, 10, 15, 2, 2, 146, 147, 9,
	16, 2, 2, 147, 149, 9, 17, 2, 2, 148, 144, 3, 2, 2, 2, 148, 145, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 32, 3, 2, 2, 2, 22, 2, 42, 52, 66, 74,
	76, 83, 85, 90, 98, 101, 108, 115, 120, 122, 126, 134, 138, 142, 148, 3,
	2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "':'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "COLON", "WS", "COMMENT", "LINE_COMMENT", "STRING_LIT",
	"VALUE_LIT",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "COLON", "WS", "COMMENT", "LINE_COMMENT", "STRING_LIT",
	"VALUE_LIT", "ExponentPart", "EscapeSequence", "HexDigits", "HexDigit",
	"Digits", "LetterOrDigit", "Letter",
}

type KVLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewKVLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *KVLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewKVLexer(input antlr.CharStream) *KVLexer {
	l := new(KVLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "KVLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KVLexer tokens.
const (
	KVLexerLPAREN       = 1
	KVLexerRPAREN       = 2
	KVLexerCOLON        = 3
	KVLexerWS           = 4
	KVLexerCOMMENT      = 5
	KVLexerLINE_COMMENT = 6
	KVLexerSTRING_LIT   = 7
	KVLexerVALUE_LIT    = 8
)
