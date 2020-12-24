// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/ApiParser.g4 by ANTLR 4.9. DO NOT EDIT.

package api // ApiParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 264,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	5, 2, 68, 10, 2, 3, 2, 7, 2, 71, 10, 2, 12, 2, 14, 2, 74, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 82, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 5, 5, 90, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7,
	98, 10, 7, 12, 7, 14, 7, 101, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	5, 9, 109, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 117,
	10, 11, 12, 11, 14, 11, 120, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12,
	126, 10, 12, 3, 13, 3, 13, 5, 13, 130, 10, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 5, 14, 136, 10, 14, 3, 14, 3, 14, 7, 14, 140, 10, 14, 12, 14, 14, 14,
	143, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 149, 10, 15, 3, 16, 3,
	16, 5, 16, 153, 10, 16, 3, 16, 5, 16, 156, 10, 16, 3, 17, 5, 17, 159, 10,
	17, 3, 17, 3, 17, 7, 17, 163, 10, 17, 12, 17, 14, 17, 166, 11, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 174, 10, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 7, 21, 187,
	10, 21, 12, 21, 14, 21, 190, 11, 21, 3, 21, 3, 21, 3, 22, 5, 22, 195, 10,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 205,
	10, 24, 12, 24, 14, 24, 208, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25,
	5, 25, 215, 10, 25, 3, 26, 5, 26, 218, 10, 26, 3, 26, 3, 26, 5, 26, 222,
	10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 5, 27, 228, 10, 27, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 5, 28, 235, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 248, 10, 31, 3, 31, 5,
	31, 251, 10, 31, 3, 31, 5, 31, 254, 10, 31, 3, 32, 3, 32, 3, 33, 3, 33,
	5, 33, 260, 10, 33, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 2, 3, 4, 2, 10, 10, 36, 36, 2, 264, 2,
	67, 3, 2, 2, 2, 4, 81, 3, 2, 2, 2, 6, 83, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2,
	10, 91, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 104, 3, 2, 2, 2, 16, 108, 3,
	2, 2, 2, 18, 110, 3, 2, 2, 2, 20, 113, 3, 2, 2, 2, 22, 125, 3, 2, 2, 2,
	24, 127, 3, 2, 2, 2, 26, 133, 3, 2, 2, 2, 28, 146, 3, 2, 2, 2, 30, 152,
	3, 2, 2, 2, 32, 158, 3, 2, 2, 2, 34, 173, 3, 2, 2, 2, 36, 175, 3, 2, 2,
	2, 38, 181, 3, 2, 2, 2, 40, 188, 3, 2, 2, 2, 42, 194, 3, 2, 2, 2, 44, 198,
	3, 2, 2, 2, 46, 200, 3, 2, 2, 2, 48, 211, 3, 2, 2, 2, 50, 217, 3, 2, 2,
	2, 52, 227, 3, 2, 2, 2, 54, 229, 3, 2, 2, 2, 56, 238, 3, 2, 2, 2, 58, 241,
	3, 2, 2, 2, 60, 244, 3, 2, 2, 2, 62, 255, 3, 2, 2, 2, 64, 257, 3, 2, 2,
	2, 66, 68, 5, 6, 4, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 72,
	3, 2, 2, 2, 69, 71, 5, 4, 3, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2,
	72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3,
	2, 2, 2, 75, 76, 7, 2, 2, 3, 76, 3, 3, 2, 2, 2, 77, 82, 5, 8, 5, 2, 78,
	82, 5, 14, 8, 2, 79, 82, 5, 16, 9, 2, 80, 82, 5, 42, 22, 2, 81, 77, 3,
	2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82,
	5, 3, 2, 2, 2, 83, 84, 7, 36, 2, 2, 84, 85, 7, 22, 2, 2, 85, 86, 7, 29,
	2, 2, 86, 7, 3, 2, 2, 2, 87, 90, 5, 10, 6, 2, 88, 90, 5, 12, 7, 2, 89,
	87, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 9, 3, 2, 2, 2, 91, 92, 7, 7, 2,
	2, 92, 93, 7, 30, 2, 2, 93, 11, 3, 2, 2, 2, 94, 95, 7, 7, 2, 2, 95, 99,
	7, 11, 2, 2, 96, 98, 7, 30, 2, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2,
	2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 102, 103, 7, 12, 2, 2, 103, 13, 3, 2, 2, 2, 104, 105, 7,
	33, 2, 2, 105, 15, 3, 2, 2, 2, 106, 109, 5, 18, 10, 2, 107, 109, 5, 20,
	11, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 17, 3, 2, 2, 2,
	110, 111, 7, 6, 2, 2, 111, 112, 5, 22, 12, 2, 112, 19, 3, 2, 2, 2, 113,
	114, 7, 6, 2, 2, 114, 118, 7, 11, 2, 2, 115, 117, 5, 22, 12, 2, 116, 115,
	3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2,
	2, 2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 12, 2, 2,
	122, 21, 3, 2, 2, 2, 123, 126, 5, 24, 13, 2, 124, 126, 5, 26, 14, 2, 125,
	123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 23, 3, 2, 2, 2, 127, 129, 7,
	36, 2, 2, 128, 130, 7, 22, 2, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3, 2,
	2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 5, 34, 18, 2, 132, 25, 3, 2, 2, 2,
	133, 135, 7, 36, 2, 2, 134, 136, 7, 9, 2, 2, 135, 134, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 141, 7, 13, 2, 2, 138, 140,
	5, 28, 15, 2, 139, 138, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3,
	2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 144, 3, 2, 2, 2, 143, 141, 3, 2, 2,
	2, 144, 145, 7, 14, 2, 2, 145, 27, 3, 2, 2, 2, 146, 148, 7, 36, 2, 2, 147,
	149, 5, 30, 16, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 29,
	3, 2, 2, 2, 150, 153, 5, 34, 18, 2, 151, 153, 5, 32, 17, 2, 152, 150, 3,
	2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154, 156, 7, 32, 2,
	2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 31, 3, 2, 2, 2, 157,
	159, 7, 9, 2, 2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160,
	3, 2, 2, 2, 160, 164, 7, 13, 2, 2, 161, 163, 5, 28, 15, 2, 162, 161, 3,
	2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2,
	2, 165, 167, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 168, 7, 14, 2, 2, 168,
	33, 3, 2, 2, 2, 169, 174, 5, 40, 21, 2, 170, 174, 5, 36, 19, 2, 171, 174,
	5, 38, 20, 2, 172, 174, 7, 5, 2, 2, 173, 169, 3, 2, 2, 2, 173, 170, 3,
	2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 35, 3, 2, 2,
	2, 175, 176, 7, 8, 2, 2, 176, 177, 7, 15, 2, 2, 177, 178, 7, 10, 2, 2,
	178, 179, 7, 16, 2, 2, 179, 180, 5, 34, 18, 2, 180, 37, 3, 2, 2, 2, 181,
	182, 7, 15, 2, 2, 182, 183, 7, 16, 2, 2, 183, 184, 5, 34, 18, 2, 184, 39,
	3, 2, 2, 2, 185, 187, 7, 25, 2, 2, 186, 185, 3, 2, 2, 2, 187, 190, 3, 2,
	2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191, 3, 2, 2, 2,
	190, 188, 3, 2, 2, 2, 191, 192, 9, 2, 2, 2, 192, 41, 3, 2, 2, 2, 193, 195,
	5, 44, 23, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3,
	2, 2, 2, 196, 197, 5, 46, 24, 2, 197, 43, 3, 2, 2, 2, 198, 199, 7, 34,
	2, 2, 199, 45, 3, 2, 2, 2, 200, 201, 7, 36, 2, 2, 201, 202, 5, 48, 25,
	2, 202, 206, 7, 13, 2, 2, 203, 205, 5, 50, 26, 2, 204, 203, 3, 2, 2, 2,
	205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	209, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 7, 14, 2, 2, 210, 47,
	3, 2, 2, 2, 211, 214, 7, 36, 2, 2, 212, 213, 7, 23, 2, 2, 213, 215, 7,
	36, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 49, 3, 2, 2,
	2, 216, 218, 5, 52, 27, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2,
	218, 221, 3, 2, 2, 2, 219, 222, 5, 44, 23, 2, 220, 222, 5, 58, 30, 2, 221,
	219, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224,
	5, 60, 31, 2, 224, 51, 3, 2, 2, 2, 225, 228, 5, 54, 28, 2, 226, 228, 5,
	56, 29, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 2, 2, 228, 53, 3, 2, 2,
	2, 229, 230, 7, 3, 2, 2, 230, 231, 7, 11, 2, 2, 231, 232, 7, 36, 2, 2,
	232, 234, 7, 24, 2, 2, 233, 235, 7, 31, 2, 2, 234, 233, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 237, 7, 12, 2, 2, 237, 55,
	3, 2, 2, 2, 238, 239, 7, 3, 2, 2, 239, 240, 7, 31, 2, 2, 240, 57, 3, 2,
	2, 2, 241, 242, 7, 4, 2, 2, 242, 243, 7, 36, 2, 2, 243, 59, 3, 2, 2, 2,
	244, 245, 7, 36, 2, 2, 245, 247, 5, 62, 32, 2, 246, 248, 5, 64, 33, 2,
	247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249,
	251, 7, 36, 2, 2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 253,
	3, 2, 2, 2, 252, 254, 5, 64, 33, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3,
	2, 2, 2, 254, 61, 3, 2, 2, 2, 255, 256, 7, 35, 2, 2, 256, 63, 3, 2, 2,
	2, 257, 259, 7, 11, 2, 2, 258, 260, 7, 36, 2, 2, 259, 258, 3, 2, 2, 2,
	259, 260, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 7, 12, 2, 2, 262,
	65, 3, 2, 2, 2, 31, 67, 72, 81, 89, 99, 108, 118, 125, 129, 135, 141, 148,
	152, 155, 158, 164, 173, 188, 194, 206, 214, 217, 221, 227, 234, 247, 250,
	253, 259,
}
var literalNames = []string{
	"", "'@doc'", "'@handler'", "'interface{}'", "'type'", "'import'", "'map'",
	"'struct'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'",
	"'/'", "'?'", "'&'", "'='", "'-'", "':'", "'*'",
}
var symbolicNames = []string{
	"", "ATDOC", "ATHANDLER", "INTERFACE", "TYPE", "IMPORT", "MAP", "STRUCT",
	"GOTYPE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA",
	"DOT", "SLASH", "QUESTION", "BITAND", "ASSIGN", "SUB", "COLON", "STAR",
	"WS", "COMMENT", "LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH", "STRING_LIT",
	"RAW_STRING", "INFO_BLOCK", "SERVER_META_STRING", "HTTP_PATH", "ID",
}

var ruleNames = []string{
	"api", "body", "syntaxLit", "importSpec", "importLit", "importLitGroup",
	"infoBlock", "typeBlock", "typeLit", "typeGroup", "typeSpec", "typeAlias",
	"typeStruct", "typeField", "filed", "innerStruct", "dataType", "mapType",
	"arrayType", "pointer", "serviceBlock", "serverMeta", "serviceBody", "serviceName",
	"serviceRoute", "routeDoc", "doc", "lineDoc", "routeHandler", "routePath",
	"path", "httpBody",
}

type ApiParser struct {
	*antlr.BaseParser
}

// NewApiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ApiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewApiParser(input antlr.TokenStream) *ApiParser {
	this := new(ApiParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ApiParser.g4"

	return this
}

// ApiParser tokens.
const (
	ApiParserEOF                = antlr.TokenEOF
	ApiParserATDOC              = 1
	ApiParserATHANDLER          = 2
	ApiParserINTERFACE          = 3
	ApiParserTYPE               = 4
	ApiParserIMPORT             = 5
	ApiParserMAP                = 6
	ApiParserSTRUCT             = 7
	ApiParserGOTYPE             = 8
	ApiParserLPAREN             = 9
	ApiParserRPAREN             = 10
	ApiParserLBRACE             = 11
	ApiParserRBRACE             = 12
	ApiParserLBRACK             = 13
	ApiParserRBRACK             = 14
	ApiParserCOMMA              = 15
	ApiParserDOT                = 16
	ApiParserSLASH              = 17
	ApiParserQUESTION           = 18
	ApiParserBITAND             = 19
	ApiParserASSIGN             = 20
	ApiParserSUB                = 21
	ApiParserCOLON              = 22
	ApiParserSTAR               = 23
	ApiParserWS                 = 24
	ApiParserCOMMENT            = 25
	ApiParserLINE_COMMENT       = 26
	ApiParserSYNTAX_VERSION     = 27
	ApiParserIMPORT_PATH        = 28
	ApiParserSTRING_LIT         = 29
	ApiParserRAW_STRING         = 30
	ApiParserINFO_BLOCK         = 31
	ApiParserSERVER_META_STRING = 32
	ApiParserHTTP_PATH          = 33
	ApiParserID                 = 34
)

// ApiParser rules.
const (
	ApiParserRULE_api            = 0
	ApiParserRULE_body           = 1
	ApiParserRULE_syntaxLit      = 2
	ApiParserRULE_importSpec     = 3
	ApiParserRULE_importLit      = 4
	ApiParserRULE_importLitGroup = 5
	ApiParserRULE_infoBlock      = 6
	ApiParserRULE_typeBlock      = 7
	ApiParserRULE_typeLit        = 8
	ApiParserRULE_typeGroup      = 9
	ApiParserRULE_typeSpec       = 10
	ApiParserRULE_typeAlias      = 11
	ApiParserRULE_typeStruct     = 12
	ApiParserRULE_typeField      = 13
	ApiParserRULE_filed          = 14
	ApiParserRULE_innerStruct    = 15
	ApiParserRULE_dataType       = 16
	ApiParserRULE_mapType        = 17
	ApiParserRULE_arrayType      = 18
	ApiParserRULE_pointer        = 19
	ApiParserRULE_serviceBlock   = 20
	ApiParserRULE_serverMeta     = 21
	ApiParserRULE_serviceBody    = 22
	ApiParserRULE_serviceName    = 23
	ApiParserRULE_serviceRoute   = 24
	ApiParserRULE_routeDoc       = 25
	ApiParserRULE_doc            = 26
	ApiParserRULE_lineDoc        = 27
	ApiParserRULE_routeHandler   = 28
	ApiParserRULE_routePath      = 29
	ApiParserRULE_path           = 30
	ApiParserRULE_httpBody       = 31
)

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_api
	return p
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) EOF() antlr.TerminalNode {
	return s.GetToken(ApiParserEOF, 0)
}

func (s *ApiContext) SyntaxLit() ISyntaxLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxLitContext)
}

func (s *ApiContext) AllBody() []IBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBodyContext)(nil)).Elem())
	var tst = make([]IBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBodyContext)
		}
	}

	return tst
}

func (s *ApiContext) Body(i int) IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Api() (localctx IApiContext) {
	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserRULE_api)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.SyntaxLit()
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(ApiParserTYPE-4))|(1<<(ApiParserIMPORT-4))|(1<<(ApiParserINFO_BLOCK-4))|(1<<(ApiParserSERVER_META_STRING-4))|(1<<(ApiParserID-4)))) != 0 {
		{
			p.SetState(67)
			p.Body()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(ApiParserEOF)
	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) ImportSpec() IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *BodyContext) InfoBlock() IInfoBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfoBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfoBlockContext)
}

func (s *BodyContext) TypeBlock() ITypeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBlockContext)
}

func (s *BodyContext) ServiceBlock() IServiceBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceBlockContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserRULE_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.ImportSpec()
		}

	case ApiParserINFO_BLOCK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.InfoBlock()
		}

	case ApiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.TypeBlock()
		}

	case ApiParserSERVER_META_STRING, ApiParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.ServiceBlock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISyntaxLitContext is an interface to support dynamic dispatch.
type ISyntaxLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSyntaxToken returns the syntaxToken token.
	GetSyntaxToken() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetSyntaxToken sets the syntaxToken token.
	SetSyntaxToken(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// IsSyntaxLitContext differentiates from other interfaces.
	IsSyntaxLitContext()
}

type SyntaxLitContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	syntaxToken antlr.Token
	version     antlr.Token
}

func NewEmptySyntaxLitContext() *SyntaxLitContext {
	var p = new(SyntaxLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_syntaxLit
	return p
}

func (*SyntaxLitContext) IsSyntaxLitContext() {}

func NewSyntaxLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxLitContext {
	var p = new(SyntaxLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_syntaxLit

	return p
}

func (s *SyntaxLitContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxLitContext) GetSyntaxToken() antlr.Token { return s.syntaxToken }

func (s *SyntaxLitContext) GetVersion() antlr.Token { return s.version }

func (s *SyntaxLitContext) SetSyntaxToken(v antlr.Token) { s.syntaxToken = v }

func (s *SyntaxLitContext) SetVersion(v antlr.Token) { s.version = v }

func (s *SyntaxLitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ApiParserASSIGN, 0)
}

func (s *SyntaxLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *SyntaxLitContext) SYNTAX_VERSION() antlr.TerminalNode {
	return s.GetToken(ApiParserSYNTAX_VERSION, 0)
}

func (s *SyntaxLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSyntaxLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) SyntaxLit() (localctx ISyntaxLitContext) {
	localctx = NewSyntaxLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ApiParserRULE_syntaxLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)

		var _m = p.Match(ApiParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
	}
	{
		p.SetState(82)
		p.Match(ApiParserASSIGN)
	}
	{
		p.SetState(83)

		var _m = p.Match(ApiParserSYNTAX_VERSION)

		localctx.(*SyntaxLitContext).version = _m
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) ImportLit() IImportLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportLitContext)
}

func (s *ImportSpecContext) ImportLitGroup() IImportLitGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportLitGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportLitGroupContext)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ApiParserRULE_importSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.ImportLitGroup()
		}

	}

	return localctx
}

// IImportLitContext is an interface to support dynamic dispatch.
type IImportLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportPath returns the importPath token.
	GetImportPath() antlr.Token

	// SetImportPath sets the importPath token.
	SetImportPath(antlr.Token)

	// IsImportLitContext differentiates from other interfaces.
	IsImportLitContext()
}

type ImportLitContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	importPath antlr.Token
}

func NewEmptyImportLitContext() *ImportLitContext {
	var p = new(ImportLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_importLit
	return p
}

func (*ImportLitContext) IsImportLitContext() {}

func NewImportLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLitContext {
	var p = new(ImportLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_importLit

	return p
}

func (s *ImportLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportLitContext) GetImportPath() antlr.Token { return s.importPath }

func (s *ImportLitContext) SetImportPath(v antlr.Token) { s.importPath = v }

func (s *ImportLitContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ApiParserIMPORT, 0)
}

func (s *ImportLitContext) IMPORT_PATH() antlr.TerminalNode {
	return s.GetToken(ApiParserIMPORT_PATH, 0)
}

func (s *ImportLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ImportLit() (localctx IImportLitContext) {
	localctx = NewImportLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ApiParserRULE_importLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(ApiParserIMPORT)
	}
	{
		p.SetState(90)

		var _m = p.Match(ApiParserIMPORT_PATH)

		localctx.(*ImportLitContext).importPath = _m
	}

	return localctx
}

// IImportLitGroupContext is an interface to support dynamic dispatch.
type IImportLitGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportPath returns the importPath token.
	GetImportPath() antlr.Token

	// SetImportPath sets the importPath token.
	SetImportPath(antlr.Token)

	// IsImportLitGroupContext differentiates from other interfaces.
	IsImportLitGroupContext()
}

type ImportLitGroupContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	importPath antlr.Token
}

func NewEmptyImportLitGroupContext() *ImportLitGroupContext {
	var p = new(ImportLitGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_importLitGroup
	return p
}

func (*ImportLitGroupContext) IsImportLitGroupContext() {}

func NewImportLitGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLitGroupContext {
	var p = new(ImportLitGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_importLitGroup

	return p
}

func (s *ImportLitGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportLitGroupContext) GetImportPath() antlr.Token { return s.importPath }

func (s *ImportLitGroupContext) SetImportPath(v antlr.Token) { s.importPath = v }

func (s *ImportLitGroupContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ApiParserIMPORT, 0)
}

func (s *ImportLitGroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *ImportLitGroupContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *ImportLitGroupContext) AllIMPORT_PATH() []antlr.TerminalNode {
	return s.GetTokens(ApiParserIMPORT_PATH)
}

func (s *ImportLitGroupContext) IMPORT_PATH(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserIMPORT_PATH, i)
}

func (s *ImportLitGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportLitGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportLitGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportLitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ImportLitGroup() (localctx IImportLitGroupContext) {
	localctx = NewImportLitGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ApiParserRULE_importLitGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(ApiParserIMPORT)
	}
	{
		p.SetState(93)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserIMPORT_PATH {
		{
			p.SetState(94)

			var _m = p.Match(ApiParserIMPORT_PATH)

			localctx.(*ImportLitGroupContext).importPath = _m
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}

// IInfoBlockContext is an interface to support dynamic dispatch.
type IInfoBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfoBlockContext differentiates from other interfaces.
	IsInfoBlockContext()
}

type InfoBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfoBlockContext() *InfoBlockContext {
	var p = new(InfoBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_infoBlock
	return p
}

func (*InfoBlockContext) IsInfoBlockContext() {}

func NewInfoBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoBlockContext {
	var p = new(InfoBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_infoBlock

	return p
}

func (s *InfoBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InfoBlockContext) INFO_BLOCK() antlr.TerminalNode {
	return s.GetToken(ApiParserINFO_BLOCK, 0)
}

func (s *InfoBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfoBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfoBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitInfoBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) InfoBlock() (localctx IInfoBlockContext) {
	localctx = NewInfoBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApiParserRULE_infoBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(ApiParserINFO_BLOCK)
	}

	return localctx
}

// ITypeBlockContext is an interface to support dynamic dispatch.
type ITypeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBlockContext differentiates from other interfaces.
	IsTypeBlockContext()
}

type TypeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBlockContext() *TypeBlockContext {
	var p = new(TypeBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeBlock
	return p
}

func (*TypeBlockContext) IsTypeBlockContext() {}

func NewTypeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockContext {
	var p = new(TypeBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeBlock

	return p
}

func (s *TypeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockContext) TypeLit() ITypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeBlockContext) TypeGroup() ITypeGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeGroupContext)
}

func (s *TypeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeBlock() (localctx ITypeBlockContext) {
	localctx = NewTypeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApiParserRULE_typeBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.TypeGroup()
		}

	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ApiParserTYPE, 0)
}

func (s *TypeLitContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ApiParserRULE_typeLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(ApiParserTYPE)
	}
	{
		p.SetState(109)
		p.TypeSpec()
	}

	return localctx
}

// ITypeGroupContext is an interface to support dynamic dispatch.
type ITypeGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeGroupContext differentiates from other interfaces.
	IsTypeGroupContext()
}

type TypeGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeGroupContext() *TypeGroupContext {
	var p = new(TypeGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeGroup
	return p
}

func (*TypeGroupContext) IsTypeGroupContext() {}

func NewTypeGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeGroupContext {
	var p = new(TypeGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeGroup

	return p
}

func (s *TypeGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeGroupContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ApiParserTYPE, 0)
}

func (s *TypeGroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *TypeGroupContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *TypeGroupContext) AllTypeSpec() []ITypeSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem())
	var tst = make([]ITypeSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecContext)
		}
	}

	return tst
}

func (s *TypeGroupContext) TypeSpec(i int) ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeGroup() (localctx ITypeGroupContext) {
	localctx = NewTypeGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ApiParserRULE_typeGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ApiParserTYPE)
	}
	{
		p.SetState(112)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(113)
			p.TypeSpec()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) TypeAlias() ITypeAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAliasContext)
}

func (s *TypeSpecContext) TypeStruct() ITypeStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ApiParserRULE_typeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.TypeAlias()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.TypeStruct()
		}

	}

	return localctx
}

// ITypeAliasContext is an interface to support dynamic dispatch.
type ITypeAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// IsTypeAliasContext differentiates from other interfaces.
	IsTypeAliasContext()
}

type TypeAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
}

func NewEmptyTypeAliasContext() *TypeAliasContext {
	var p = new(TypeAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeAlias
	return p
}

func (*TypeAliasContext) IsTypeAliasContext() {}

func NewTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasContext {
	var p = new(TypeAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeAlias

	return p
}

func (s *TypeAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeAliasContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *TypeAliasContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ApiParserASSIGN, 0)
}

func (s *TypeAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeAlias() (localctx ITypeAliasContext) {
	localctx = NewTypeAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ApiParserRULE_typeAlias)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserASSIGN {
		{
			p.SetState(126)
			p.Match(ApiParserASSIGN)
		}

	}
	{
		p.SetState(129)
		p.DataType()
	}

	return localctx
}

// ITypeStructContext is an interface to support dynamic dispatch.
type ITypeStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsTypeStructContext differentiates from other interfaces.
	IsTypeStructContext()
}

type TypeStructContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyTypeStructContext() *TypeStructContext {
	var p = new(TypeStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeStruct
	return p
}

func (*TypeStructContext) IsTypeStructContext() {}

func NewTypeStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStructContext {
	var p = new(TypeStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeStruct

	return p
}

func (s *TypeStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStructContext) GetName() antlr.Token { return s.name }

func (s *TypeStructContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeStructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserLBRACE, 0)
}

func (s *TypeStructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserRBRACE, 0)
}

func (s *TypeStructContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *TypeStructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ApiParserSTRUCT, 0)
}

func (s *TypeStructContext) AllTypeField() []ITypeFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem())
	var tst = make([]ITypeFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFieldContext)
		}
	}

	return tst
}

func (s *TypeStructContext) TypeField(i int) ITypeFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFieldContext)
}

func (s *TypeStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeStruct() (localctx ITypeStructContext) {
	localctx = NewTypeStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ApiParserRULE_typeStruct)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeStructContext).name = _m
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRUCT {
		{
			p.SetState(132)
			p.Match(ApiParserSTRUCT)
		}

	}
	{
		p.SetState(135)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(136)
			p.TypeField()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(ApiParserRBRACE)
	}

	return localctx
}

// ITypeFieldContext is an interface to support dynamic dispatch.
type ITypeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsTypeFieldContext differentiates from other interfaces.
	IsTypeFieldContext()
}

type TypeFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyTypeFieldContext() *TypeFieldContext {
	var p = new(TypeFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_typeField
	return p
}

func (*TypeFieldContext) IsTypeFieldContext() {}

func NewTypeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFieldContext {
	var p = new(TypeFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_typeField

	return p
}

func (s *TypeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFieldContext) GetName() antlr.Token { return s.name }

func (s *TypeFieldContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *TypeFieldContext) Filed() IFiledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFiledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFiledContext)
}

func (s *TypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) TypeField() (localctx ITypeFieldContext) {
	localctx = NewTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ApiParserRULE_typeField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeFieldContext).name = _m
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(145)
			p.Filed()
		}

	}

	return localctx
}

// IFiledContext is an interface to support dynamic dispatch.
type IFiledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTag returns the tag token.
	GetTag() antlr.Token

	// SetTag sets the tag token.
	SetTag(antlr.Token)

	// IsFiledContext differentiates from other interfaces.
	IsFiledContext()
}

type FiledContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tag    antlr.Token
}

func NewEmptyFiledContext() *FiledContext {
	var p = new(FiledContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_filed
	return p
}

func (*FiledContext) IsFiledContext() {}

func NewFiledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiledContext {
	var p = new(FiledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_filed

	return p
}

func (s *FiledContext) GetParser() antlr.Parser { return s.parser }

func (s *FiledContext) GetTag() antlr.Token { return s.tag }

func (s *FiledContext) SetTag(v antlr.Token) { s.tag = v }

func (s *FiledContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *FiledContext) InnerStruct() IInnerStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInnerStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInnerStructContext)
}

func (s *FiledContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserRAW_STRING, 0)
}

func (s *FiledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitFiled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Filed() (localctx IFiledContext) {
	localctx = NewFiledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ApiParserRULE_filed)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserINTERFACE, ApiParserMAP, ApiParserGOTYPE, ApiParserLBRACK, ApiParserSTAR, ApiParserID:
		{
			p.SetState(148)
			p.DataType()
		}

	case ApiParserSTRUCT, ApiParserLBRACE:
		{
			p.SetState(149)
			p.InnerStruct()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserRAW_STRING {
		{
			p.SetState(152)

			var _m = p.Match(ApiParserRAW_STRING)

			localctx.(*FiledContext).tag = _m
		}

	}

	return localctx
}

// IInnerStructContext is an interface to support dynamic dispatch.
type IInnerStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerStructContext differentiates from other interfaces.
	IsInnerStructContext()
}

type InnerStructContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerStructContext() *InnerStructContext {
	var p = new(InnerStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_innerStruct
	return p
}

func (*InnerStructContext) IsInnerStructContext() {}

func NewInnerStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerStructContext {
	var p = new(InnerStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_innerStruct

	return p
}

func (s *InnerStructContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerStructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserLBRACE, 0)
}

func (s *InnerStructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserRBRACE, 0)
}

func (s *InnerStructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ApiParserSTRUCT, 0)
}

func (s *InnerStructContext) AllTypeField() []ITypeFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem())
	var tst = make([]ITypeFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFieldContext)
		}
	}

	return tst
}

func (s *InnerStructContext) TypeField(i int) ITypeFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFieldContext)
}

func (s *InnerStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitInnerStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) InnerStruct() (localctx IInnerStructContext) {
	localctx = NewInnerStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ApiParserRULE_innerStruct)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRUCT {
		{
			p.SetState(155)
			p.Match(ApiParserSTRUCT)
		}

	}
	{
		p.SetState(158)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(159)
			p.TypeField()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(165)
		p.Match(ApiParserRBRACE)
	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *DataTypeContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *DataTypeContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *DataTypeContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ApiParserINTERFACE, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ApiParserRULE_dataType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserGOTYPE, ApiParserSTAR, ApiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Pointer()
		}

	case ApiParserMAP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.MapType()
		}

	case ApiParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.ArrayType()
		}

	case ApiParserINTERFACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(170)
			p.Match(ApiParserINTERFACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IDataTypeContext

	// SetValue sets the value rule contexts.
	SetValue(IDataTypeContext)

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  IDataTypeContext
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) GetKey() antlr.Token { return s.key }

func (s *MapTypeContext) SetKey(v antlr.Token) { s.key = v }

func (s *MapTypeContext) GetValue() IDataTypeContext { return s.value }

func (s *MapTypeContext) SetValue(v IDataTypeContext) { s.value = v }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(ApiParserMAP, 0)
}

func (s *MapTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ApiParserLBRACK, 0)
}

func (s *MapTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ApiParserRBRACK, 0)
}

func (s *MapTypeContext) GOTYPE() antlr.TerminalNode {
	return s.GetToken(ApiParserGOTYPE, 0)
}

func (s *MapTypeContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ApiParserRULE_mapType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(ApiParserMAP)
	}
	{
		p.SetState(174)
		p.Match(ApiParserLBRACK)
	}
	{
		p.SetState(175)

		var _m = p.Match(ApiParserGOTYPE)

		localctx.(*MapTypeContext).key = _m
	}
	{
		p.SetState(176)
		p.Match(ApiParserRBRACK)
	}
	{
		p.SetState(177)

		var _x = p.DataType()

		localctx.(*MapTypeContext).value = _x
	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLit returns the lit rule contexts.
	GetLit() IDataTypeContext

	// SetLit sets the lit rule contexts.
	SetLit(IDataTypeContext)

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lit    IDataTypeContext
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) GetLit() IDataTypeContext { return s.lit }

func (s *ArrayTypeContext) SetLit(v IDataTypeContext) { s.lit = v }

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ApiParserLBRACK, 0)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ApiParserRBRACK, 0)
}

func (s *ArrayTypeContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ApiParserRULE_arrayType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(ApiParserLBRACK)
	}
	{
		p.SetState(180)
		p.Match(ApiParserRBRACK)
	}
	{
		p.SetState(181)

		var _x = p.DataType()

		localctx.(*ArrayTypeContext).lit = _x
	}

	return localctx
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_pointer
	return p
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerContext) GOTYPE() antlr.TerminalNode {
	return s.GetToken(ApiParserGOTYPE, 0)
}

func (s *PointerContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *PointerContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(ApiParserSTAR)
}

func (s *PointerContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserSTAR, i)
}

func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPointer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Pointer() (localctx IPointerContext) {
	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ApiParserRULE_pointer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserSTAR {
		{
			p.SetState(183)
			p.Match(ApiParserSTAR)
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ApiParserGOTYPE || _la == ApiParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IServiceBlockContext is an interface to support dynamic dispatch.
type IServiceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceBlockContext differentiates from other interfaces.
	IsServiceBlockContext()
}

type ServiceBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceBlockContext() *ServiceBlockContext {
	var p = new(ServiceBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_serviceBlock
	return p
}

func (*ServiceBlockContext) IsServiceBlockContext() {}

func NewServiceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBlockContext {
	var p = new(ServiceBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_serviceBlock

	return p
}

func (s *ServiceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBlockContext) ServiceBody() IServiceBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceBodyContext)
}

func (s *ServiceBlockContext) ServerMeta() IServerMetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServerMetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServerMetaContext)
}

func (s *ServiceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ServiceBlock() (localctx IServiceBlockContext) {
	localctx = NewServiceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ApiParserRULE_serviceBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSERVER_META_STRING {
		{
			p.SetState(191)
			p.ServerMeta()
		}

	}
	{
		p.SetState(194)
		p.ServiceBody()
	}

	return localctx
}

// IServerMetaContext is an interface to support dynamic dispatch.
type IServerMetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServerMetaContext differentiates from other interfaces.
	IsServerMetaContext()
}

type ServerMetaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerMetaContext() *ServerMetaContext {
	var p = new(ServerMetaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_serverMeta
	return p
}

func (*ServerMetaContext) IsServerMetaContext() {}

func NewServerMetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerMetaContext {
	var p = new(ServerMetaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_serverMeta

	return p
}

func (s *ServerMetaContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerMetaContext) SERVER_META_STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserSERVER_META_STRING, 0)
}

func (s *ServerMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerMetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerMetaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServerMeta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ServerMeta() (localctx IServerMetaContext) {
	localctx = NewServerMetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ApiParserRULE_serverMeta)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ApiParserSERVER_META_STRING)
	}

	return localctx
}

// IServiceBodyContext is an interface to support dynamic dispatch.
type IServiceBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetServiceToken returns the serviceToken token.
	GetServiceToken() antlr.Token

	// SetServiceToken sets the serviceToken token.
	SetServiceToken(antlr.Token)

	// GetRoutes returns the routes rule contexts.
	GetRoutes() IServiceRouteContext

	// SetRoutes sets the routes rule contexts.
	SetRoutes(IServiceRouteContext)

	// IsServiceBodyContext differentiates from other interfaces.
	IsServiceBodyContext()
}

type ServiceBodyContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	serviceToken antlr.Token
	routes       IServiceRouteContext
}

func NewEmptyServiceBodyContext() *ServiceBodyContext {
	var p = new(ServiceBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_serviceBody
	return p
}

func (*ServiceBodyContext) IsServiceBodyContext() {}

func NewServiceBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceBodyContext {
	var p = new(ServiceBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_serviceBody

	return p
}

func (s *ServiceBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceBodyContext) GetServiceToken() antlr.Token { return s.serviceToken }

func (s *ServiceBodyContext) SetServiceToken(v antlr.Token) { s.serviceToken = v }

func (s *ServiceBodyContext) GetRoutes() IServiceRouteContext { return s.routes }

func (s *ServiceBodyContext) SetRoutes(v IServiceRouteContext) { s.routes = v }

func (s *ServiceBodyContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserLBRACE, 0)
}

func (s *ServiceBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ApiParserRBRACE, 0)
}

func (s *ServiceBodyContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *ServiceBodyContext) AllServiceRoute() []IServiceRouteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceRouteContext)(nil)).Elem())
	var tst = make([]IServiceRouteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceRouteContext)
		}
	}

	return tst
}

func (s *ServiceBodyContext) ServiceRoute(i int) IServiceRouteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceRouteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceRouteContext)
}

func (s *ServiceBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ServiceBody() (localctx IServiceBodyContext) {
	localctx = NewServiceBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ApiParserRULE_serviceBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)

		var _m = p.Match(ApiParserID)

		localctx.(*ServiceBodyContext).serviceToken = _m
	}
	{
		p.SetState(199)
		p.ServiceName()
	}
	{
		p.SetState(200)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(ApiParserATDOC-1))|(1<<(ApiParserATHANDLER-1))|(1<<(ApiParserSERVER_META_STRING-1)))) != 0 {
		{
			p.SetState(201)

			var _x = p.ServiceRoute()

			localctx.(*ServiceBodyContext).routes = _x
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(ApiParserRBRACE)
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserID, i)
}

func (s *ServiceNameContext) SUB() antlr.TerminalNode {
	return s.GetToken(ApiParserSUB, 0)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ApiParserRULE_serviceName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(ApiParserID)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSUB {
		{
			p.SetState(210)
			p.Match(ApiParserSUB)
		}
		{
			p.SetState(211)
			p.Match(ApiParserID)
		}

	}

	return localctx
}

// IServiceRouteContext is an interface to support dynamic dispatch.
type IServiceRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceRouteContext differentiates from other interfaces.
	IsServiceRouteContext()
}

type ServiceRouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceRouteContext() *ServiceRouteContext {
	var p = new(ServiceRouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_serviceRoute
	return p
}

func (*ServiceRouteContext) IsServiceRouteContext() {}

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	var p = new(ServiceRouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_serviceRoute

	return p
}

func (s *ServiceRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceRouteContext) RoutePath() IRoutePathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoutePathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRoutePathContext)
}

func (s *ServiceRouteContext) ServerMeta() IServerMetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServerMetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServerMetaContext)
}

func (s *ServiceRouteContext) RouteHandler() IRouteHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRouteHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRouteHandlerContext)
}

func (s *ServiceRouteContext) RouteDoc() IRouteDocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRouteDocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRouteDocContext)
}

func (s *ServiceRouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceRouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) ServiceRoute() (localctx IServiceRouteContext) {
	localctx = NewServiceRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ApiParserRULE_serviceRoute)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserATDOC {
		{
			p.SetState(214)
			p.RouteDoc()
		}

	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserSERVER_META_STRING:
		{
			p.SetState(217)
			p.ServerMeta()
		}

	case ApiParserATHANDLER:
		{
			p.SetState(218)
			p.RouteHandler()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(221)
		p.RoutePath()
	}

	return localctx
}

// IRouteDocContext is an interface to support dynamic dispatch.
type IRouteDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRouteDocContext differentiates from other interfaces.
	IsRouteDocContext()
}

type RouteDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteDocContext() *RouteDocContext {
	var p = new(RouteDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_routeDoc
	return p
}

func (*RouteDocContext) IsRouteDocContext() {}

func NewRouteDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteDocContext {
	var p = new(RouteDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_routeDoc

	return p
}

func (s *RouteDocContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteDocContext) Doc() IDocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocContext)
}

func (s *RouteDocContext) LineDoc() ILineDocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineDocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineDocContext)
}

func (s *RouteDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteDocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRouteDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) RouteDoc() (localctx IRouteDocContext) {
	localctx = NewRouteDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ApiParserRULE_routeDoc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Doc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.LineDoc()
		}

	}

	return localctx
}

// IDocContext is an interface to support dynamic dispatch.
type IDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSummaryToken returns the summaryToken token.
	GetSummaryToken() antlr.Token

	// SetSummaryToken sets the summaryToken token.
	SetSummaryToken(antlr.Token)

	// IsDocContext differentiates from other interfaces.
	IsDocContext()
}

type DocContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	summaryToken antlr.Token
}

func NewEmptyDocContext() *DocContext {
	var p = new(DocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_doc
	return p
}

func (*DocContext) IsDocContext() {}

func NewDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocContext {
	var p = new(DocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_doc

	return p
}

func (s *DocContext) GetParser() antlr.Parser { return s.parser }

func (s *DocContext) GetSummaryToken() antlr.Token { return s.summaryToken }

func (s *DocContext) SetSummaryToken(v antlr.Token) { s.summaryToken = v }

func (s *DocContext) ATDOC() antlr.TerminalNode {
	return s.GetToken(ApiParserATDOC, 0)
}

func (s *DocContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *DocContext) COLON() antlr.TerminalNode {
	return s.GetToken(ApiParserCOLON, 0)
}

func (s *DocContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *DocContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *DocContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ApiParserSTRING_LIT, 0)
}

func (s *DocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Doc() (localctx IDocContext) {
	localctx = NewDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ApiParserRULE_doc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(ApiParserATDOC)
	}
	{
		p.SetState(228)
		p.Match(ApiParserLPAREN)
	}
	{
		p.SetState(229)

		var _m = p.Match(ApiParserID)

		localctx.(*DocContext).summaryToken = _m
	}
	{
		p.SetState(230)
		p.Match(ApiParserCOLON)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRING_LIT {
		{
			p.SetState(231)
			p.Match(ApiParserSTRING_LIT)
		}

	}
	{
		p.SetState(234)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}

// ILineDocContext is an interface to support dynamic dispatch.
type ILineDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineDocContext differentiates from other interfaces.
	IsLineDocContext()
}

type LineDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineDocContext() *LineDocContext {
	var p = new(LineDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_lineDoc
	return p
}

func (*LineDocContext) IsLineDocContext() {}

func NewLineDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineDocContext {
	var p = new(LineDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_lineDoc

	return p
}

func (s *LineDocContext) GetParser() antlr.Parser { return s.parser }

func (s *LineDocContext) ATDOC() antlr.TerminalNode {
	return s.GetToken(ApiParserATDOC, 0)
}

func (s *LineDocContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ApiParserSTRING_LIT, 0)
}

func (s *LineDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineDocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitLineDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) LineDoc() (localctx ILineDocContext) {
	localctx = NewLineDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ApiParserRULE_lineDoc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(ApiParserATDOC)
	}
	{
		p.SetState(237)
		p.Match(ApiParserSTRING_LIT)
	}

	return localctx
}

// IRouteHandlerContext is an interface to support dynamic dispatch.
type IRouteHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRouteHandlerContext differentiates from other interfaces.
	IsRouteHandlerContext()
}

type RouteHandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteHandlerContext() *RouteHandlerContext {
	var p = new(RouteHandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_routeHandler
	return p
}

func (*RouteHandlerContext) IsRouteHandlerContext() {}

func NewRouteHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteHandlerContext {
	var p = new(RouteHandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_routeHandler

	return p
}

func (s *RouteHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteHandlerContext) ATHANDLER() antlr.TerminalNode {
	return s.GetToken(ApiParserATHANDLER, 0)
}

func (s *RouteHandlerContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *RouteHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRouteHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) RouteHandler() (localctx IRouteHandlerContext) {
	localctx = NewRouteHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ApiParserRULE_routeHandler)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(ApiParserATHANDLER)
	}
	{
		p.SetState(240)
		p.Match(ApiParserID)
	}

	return localctx
}

// IRoutePathContext is an interface to support dynamic dispatch.
type IRoutePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHttpMethodToken returns the httpMethodToken token.
	GetHttpMethodToken() antlr.Token

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// SetHttpMethodToken sets the httpMethodToken token.
	SetHttpMethodToken(antlr.Token)

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// GetReq returns the req rule contexts.
	GetReq() IHttpBodyContext

	// GetReply returns the reply rule contexts.
	GetReply() IHttpBodyContext

	// SetReq sets the req rule contexts.
	SetReq(IHttpBodyContext)

	// SetReply sets the reply rule contexts.
	SetReply(IHttpBodyContext)

	// IsRoutePathContext differentiates from other interfaces.
	IsRoutePathContext()
}

type RoutePathContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	httpMethodToken antlr.Token
	req             IHttpBodyContext
	returnToken     antlr.Token
	reply           IHttpBodyContext
}

func NewEmptyRoutePathContext() *RoutePathContext {
	var p = new(RoutePathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_routePath
	return p
}

func (*RoutePathContext) IsRoutePathContext() {}

func NewRoutePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutePathContext {
	var p = new(RoutePathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_routePath

	return p
}

func (s *RoutePathContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutePathContext) GetHttpMethodToken() antlr.Token { return s.httpMethodToken }

func (s *RoutePathContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *RoutePathContext) SetHttpMethodToken(v antlr.Token) { s.httpMethodToken = v }

func (s *RoutePathContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *RoutePathContext) GetReq() IHttpBodyContext { return s.req }

func (s *RoutePathContext) GetReply() IHttpBodyContext { return s.reply }

func (s *RoutePathContext) SetReq(v IHttpBodyContext) { s.req = v }

func (s *RoutePathContext) SetReply(v IHttpBodyContext) { s.reply = v }

func (s *RoutePathContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *RoutePathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserID)
}

func (s *RoutePathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserID, i)
}

func (s *RoutePathContext) AllHttpBody() []IHttpBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHttpBodyContext)(nil)).Elem())
	var tst = make([]IHttpBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHttpBodyContext)
		}
	}

	return tst
}

func (s *RoutePathContext) HttpBody(i int) IHttpBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHttpBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHttpBodyContext)
}

func (s *RoutePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutePathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRoutePath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) RoutePath() (localctx IRoutePathContext) {
	localctx = NewRoutePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ApiParserRULE_routePath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)

		var _m = p.Match(ApiParserID)

		localctx.(*RoutePathContext).httpMethodToken = _m
	}
	{
		p.SetState(243)
		p.Path()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)

			var _x = p.HttpBody()

			localctx.(*RoutePathContext).req = _x
		}

	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserID {
		{
			p.SetState(247)

			var _m = p.Match(ApiParserID)

			localctx.(*RoutePathContext).returnToken = _m
		}

	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserLPAREN {
		{
			p.SetState(250)

			var _x = p.HttpBody()

			localctx.(*RoutePathContext).reply = _x
		}

	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) HTTP_PATH() antlr.TerminalNode {
	return s.GetToken(ApiParserHTTP_PATH, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ApiParserRULE_path)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(ApiParserHTTP_PATH)
	}

	return localctx
}

// IHttpBodyContext is an interface to support dynamic dispatch.
type IHttpBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetObj returns the obj token.
	GetObj() antlr.Token

	// SetObj sets the obj token.
	SetObj(antlr.Token)

	// IsHttpBodyContext differentiates from other interfaces.
	IsHttpBodyContext()
}

type HttpBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	obj    antlr.Token
}

func NewEmptyHttpBodyContext() *HttpBodyContext {
	var p = new(HttpBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_httpBody
	return p
}

func (*HttpBodyContext) IsHttpBodyContext() {}

func NewHttpBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpBodyContext {
	var p = new(HttpBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_httpBody

	return p
}

func (s *HttpBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpBodyContext) GetObj() antlr.Token { return s.obj }

func (s *HttpBodyContext) SetObj(v antlr.Token) { s.obj = v }

func (s *HttpBodyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *HttpBodyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *HttpBodyContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *HttpBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitHttpBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) HttpBody() (localctx IHttpBodyContext) {
	localctx = NewHttpBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ApiParserRULE_httpBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserID {
		{
			p.SetState(256)

			var _m = p.Match(ApiParserID)

			localctx.(*HttpBodyContext).obj = _m
		}

	}
	{
		p.SetState(259)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}
