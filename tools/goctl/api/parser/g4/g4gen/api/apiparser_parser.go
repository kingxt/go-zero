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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 243,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 62, 10, 3, 3, 4, 5, 4, 65, 10, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 73, 10, 4, 3, 5, 3, 5, 5, 5,
	77, 10, 5, 3, 6, 5, 6, 80, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 86, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 92, 10, 7, 3, 7, 6, 7, 95, 10, 7, 13,
	7, 14, 7, 96, 3, 7, 3, 7, 3, 8, 5, 8, 102, 10, 8, 3, 8, 3, 8, 5, 8, 106,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 5, 10, 112, 10, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 118, 10, 10, 3, 10, 6, 10, 121, 10, 10, 13, 10, 14, 10, 122,
	3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 129, 10, 11, 3, 12, 5, 12, 132, 10,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13, 139, 10, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 6, 13, 145, 10, 13, 13, 13, 14, 13, 146, 3, 13, 3, 13, 3,
	14, 3, 14, 5, 14, 153, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 159,
	10, 15, 3, 15, 3, 15, 5, 15, 163, 10, 15, 3, 15, 6, 15, 166, 10, 15, 13,
	15, 14, 15, 167, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 175, 10, 16,
	3, 16, 3, 16, 5, 16, 179, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 184, 10,
	17, 3, 18, 5, 18, 187, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 193,
	10, 18, 3, 18, 5, 18, 196, 10, 18, 3, 19, 5, 19, 199, 10, 19, 3, 19, 5,
	19, 202, 10, 19, 3, 19, 3, 19, 5, 19, 206, 10, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 216, 10, 20, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 24, 5, 24, 233, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 239,
	10, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 3, 3, 2,
	17, 18, 2, 260, 2, 53, 3, 2, 2, 2, 4, 61, 3, 2, 2, 2, 6, 64, 3, 2, 2, 2,
	8, 76, 3, 2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 101, 3,
	2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 111, 3, 2, 2, 2, 20, 128, 3, 2, 2, 2,
	22, 131, 3, 2, 2, 2, 24, 138, 3, 2, 2, 2, 26, 152, 3, 2, 2, 2, 28, 154,
	3, 2, 2, 2, 30, 171, 3, 2, 2, 2, 32, 183, 3, 2, 2, 2, 34, 186, 3, 2, 2,
	2, 36, 198, 3, 2, 2, 2, 38, 215, 3, 2, 2, 2, 40, 217, 3, 2, 2, 2, 42, 220,
	3, 2, 2, 2, 44, 227, 3, 2, 2, 2, 46, 232, 3, 2, 2, 2, 48, 240, 3, 2, 2,
	2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 3, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	56, 62, 5, 6, 4, 2, 57, 62, 5, 8, 5, 2, 58, 62, 5, 18, 10, 2, 59, 62, 5,
	20, 11, 2, 60, 62, 5, 48, 25, 2, 61, 56, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2,
	61, 58, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 5, 3, 2,
	2, 2, 63, 65, 5, 48, 25, 2, 64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65,
	66, 3, 2, 2, 2, 66, 67, 8, 4, 1, 2, 67, 68, 7, 22, 2, 2, 68, 69, 7, 3,
	2, 2, 69, 70, 8, 4, 1, 2, 70, 72, 7, 19, 2, 2, 71, 73, 5, 48, 25, 2, 72,
	71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 7, 3, 2, 2, 2, 74, 77, 5, 10, 6,
	2, 75, 77, 5, 12, 7, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 9,
	3, 2, 2, 2, 78, 80, 5, 48, 25, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2,
	2, 80, 81, 3, 2, 2, 2, 81, 82, 8, 6, 1, 2, 82, 83, 7, 22, 2, 2, 83, 85,
	5, 16, 9, 2, 84, 86, 5, 48, 25, 2, 85, 84, 3, 2, 2, 2, 85, 86, 3, 2, 2,
	2, 86, 11, 3, 2, 2, 2, 87, 88, 8, 7, 1, 2, 88, 89, 7, 22, 2, 2, 89, 91,
	7, 4, 2, 2, 90, 92, 5, 48, 25, 2, 91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 94, 3, 2, 2, 2, 93, 95, 5, 14, 8, 2, 94, 93, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2,
	98, 99, 7, 5, 2, 2, 99, 13, 3, 2, 2, 2, 100, 102, 5, 48, 25, 2, 101, 100,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 105, 5, 16,
	9, 2, 104, 106, 5, 48, 25, 2, 105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 106, 15, 3, 2, 2, 2, 107, 108, 8, 9, 1, 2, 108, 109, 7, 19, 2, 2, 109,
	17, 3, 2, 2, 2, 110, 112, 5, 48, 25, 2, 111, 110, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 8, 10, 1, 2, 114, 115, 7, 22,
	2, 2, 115, 117, 7, 4, 2, 2, 116, 118, 5, 48, 25, 2, 117, 116, 3, 2, 2,
	2, 117, 118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 121, 5, 46, 24, 2,
	120, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122,
	123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 7, 5, 2, 2, 125, 19, 3,
	2, 2, 2, 126, 129, 5, 22, 12, 2, 127, 129, 5, 24, 13, 2, 128, 126, 3, 2,
	2, 2, 128, 127, 3, 2, 2, 2, 129, 21, 3, 2, 2, 2, 130, 132, 5, 48, 25, 2,
	131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	134, 8, 12, 1, 2, 134, 135, 7, 22, 2, 2, 135, 136, 5, 26, 14, 2, 136, 23,
	3, 2, 2, 2, 137, 139, 5, 48, 25, 2, 138, 137, 3, 2, 2, 2, 138, 139, 3,
	2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 8, 13, 1, 2, 141, 142, 7, 22,
	2, 2, 142, 144, 7, 4, 2, 2, 143, 145, 5, 26, 14, 2, 144, 143, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 149, 7, 5, 2, 2, 149, 25, 3, 2, 2, 2, 150, 153, 5,
	30, 16, 2, 151, 153, 5, 28, 15, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2,
	2, 2, 153, 27, 3, 2, 2, 2, 154, 155, 8, 15, 1, 2, 155, 156, 7, 22, 2, 2,
	156, 158, 8, 15, 1, 2, 157, 159, 7, 22, 2, 2, 158, 157, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 7, 6, 2, 2, 161, 163,
	5, 48, 25, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3,
	2, 2, 2, 164, 166, 5, 32, 17, 2, 165, 164, 3, 2, 2, 2, 166, 167, 3, 2,
	2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2,
	169, 170, 7, 7, 2, 2, 170, 29, 3, 2, 2, 2, 171, 172, 8, 16, 1, 2, 172,
	174, 7, 22, 2, 2, 173, 175, 7, 3, 2, 2, 174, 173, 3, 2, 2, 2, 174, 175,
	3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 5, 38, 20, 2, 177, 179, 5,
	48, 25, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 31, 3, 2, 2,
	2, 180, 181, 6, 17, 2, 2, 181, 184, 5, 36, 19, 2, 182, 184, 5, 34, 18,
	2, 183, 180, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 33, 3, 2, 2, 2, 185,
	187, 5, 48, 25, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 189, 7, 22, 2, 2, 189, 190, 5, 38, 20, 2, 190, 192, 8,
	18, 1, 2, 191, 193, 7, 20, 2, 2, 192, 191, 3, 2, 2, 2, 192, 193, 3, 2,
	2, 2, 193, 195, 3, 2, 2, 2, 194, 196, 5, 48, 25, 2, 195, 194, 3, 2, 2,
	2, 195, 196, 3, 2, 2, 2, 196, 35, 3, 2, 2, 2, 197, 199, 5, 48, 25, 2, 198,
	197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201, 3, 2, 2, 2, 200, 202,
	7, 8, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 205, 7, 22, 2, 2, 204, 206, 5, 48, 25, 2, 205, 204, 3, 2, 2,
	2, 205, 206, 3, 2, 2, 2, 206, 37, 3, 2, 2, 2, 207, 208, 8, 20, 1, 2, 208,
	216, 7, 22, 2, 2, 209, 216, 5, 42, 22, 2, 210, 216, 5, 44, 23, 2, 211,
	216, 7, 14, 2, 2, 212, 216, 7, 9, 2, 2, 213, 216, 5, 40, 21, 2, 214, 216,
	5, 28, 15, 2, 215, 207, 3, 2, 2, 2, 215, 209, 3, 2, 2, 2, 215, 210, 3,
	2, 2, 2, 215, 211, 3, 2, 2, 2, 215, 212, 3, 2, 2, 2, 215, 213, 3, 2, 2,
	2, 215, 214, 3, 2, 2, 2, 216, 39, 3, 2, 2, 2, 217, 218, 7, 8, 2, 2, 218,
	219, 7, 22, 2, 2, 219, 41, 3, 2, 2, 2, 220, 221, 8, 22, 1, 2, 221, 222,
	7, 22, 2, 2, 222, 223, 7, 10, 2, 2, 223, 224, 7, 22, 2, 2, 224, 225, 7,
	11, 2, 2, 225, 226, 5, 38, 20, 2, 226, 43, 3, 2, 2, 2, 227, 228, 7, 10,
	2, 2, 228, 229, 7, 11, 2, 2, 229, 230, 5, 38, 20, 2, 230, 45, 3, 2, 2,
	2, 231, 233, 5, 48, 25, 2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2,
	233, 234, 3, 2, 2, 2, 234, 235, 7, 22, 2, 2, 235, 236, 8, 24, 1, 2, 236,
	238, 7, 21, 2, 2, 237, 239, 5, 48, 25, 2, 238, 237, 3, 2, 2, 2, 238, 239,
	3, 2, 2, 2, 239, 47, 3, 2, 2, 2, 240, 241, 9, 2, 2, 2, 241, 49, 3, 2, 2,
	2, 36, 53, 61, 64, 72, 76, 79, 85, 91, 96, 101, 105, 111, 117, 122, 128,
	131, 138, 146, 152, 158, 162, 167, 174, 178, 183, 186, 192, 195, 198, 201,
	205, 215, 232, 238,
}
var literalNames = []string{
	"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'time.Time'", "'['", "']'",
	"'@doc'", "'@handler'", "'interface{}'", "'@server'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "ATDOC", "ATHANDLER", "INTERFACE",
	"ATSERVER", "WS", "COMMENT", "LINE_COMMENT", "STRING", "RAW_STRING", "LINE_VALUE",
	"ID",
}

var ruleNames = []string{
	"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock", "importBlockValue",
	"importValue", "infoSpec", "typeSpec", "typeLit", "typeBlock", "typeBody",
	"typeStruct", "typeAlias", "field", "normalField", "anonymousFiled", "dataType",
	"pointerType", "mapType", "arrayType", "kvLit", "commentSpec",
}

type ApiParserParser struct {
	*antlr.BaseParser
}

// NewApiParserParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ApiParserParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewApiParserParser(input antlr.TokenStream) *ApiParserParser {
	this := new(ApiParserParser)
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

// ApiParserParser tokens.
const (
	ApiParserParserEOF          = antlr.TokenEOF
	ApiParserParserT__0         = 1
	ApiParserParserT__1         = 2
	ApiParserParserT__2         = 3
	ApiParserParserT__3         = 4
	ApiParserParserT__4         = 5
	ApiParserParserT__5         = 6
	ApiParserParserT__6         = 7
	ApiParserParserT__7         = 8
	ApiParserParserT__8         = 9
	ApiParserParserATDOC        = 10
	ApiParserParserATHANDLER    = 11
	ApiParserParserINTERFACE    = 12
	ApiParserParserATSERVER     = 13
	ApiParserParserWS           = 14
	ApiParserParserCOMMENT      = 15
	ApiParserParserLINE_COMMENT = 16
	ApiParserParserSTRING       = 17
	ApiParserParserRAW_STRING   = 18
	ApiParserParserLINE_VALUE   = 19
	ApiParserParserID           = 20
)

// ApiParserParser rules.
const (
	ApiParserParserRULE_api              = 0
	ApiParserParserRULE_spec             = 1
	ApiParserParserRULE_syntaxLit        = 2
	ApiParserParserRULE_importSpec       = 3
	ApiParserParserRULE_importLit        = 4
	ApiParserParserRULE_importBlock      = 5
	ApiParserParserRULE_importBlockValue = 6
	ApiParserParserRULE_importValue      = 7
	ApiParserParserRULE_infoSpec         = 8
	ApiParserParserRULE_typeSpec         = 9
	ApiParserParserRULE_typeLit          = 10
	ApiParserParserRULE_typeBlock        = 11
	ApiParserParserRULE_typeBody         = 12
	ApiParserParserRULE_typeStruct       = 13
	ApiParserParserRULE_typeAlias        = 14
	ApiParserParserRULE_field            = 15
	ApiParserParserRULE_normalField      = 16
	ApiParserParserRULE_anonymousFiled   = 17
	ApiParserParserRULE_dataType         = 18
	ApiParserParserRULE_pointerType      = 19
	ApiParserParserRULE_mapType          = 20
	ApiParserParserRULE_arrayType        = 21
	ApiParserParserRULE_kvLit            = 22
	ApiParserParserRULE_commentSpec      = 23
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
	p.RuleIndex = ApiParserParserRULE_api
	return p
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) AllSpec() []ISpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpecContext)(nil)).Elem())
	var tst = make([]ISpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpecContext)
		}
	}

	return tst
}

func (s *ApiContext) Spec(i int) ISpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
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

func (p *ApiParserParser) Api() (localctx IApiContext) {
	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserParserRULE_api)
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0 {
		{
			p.SetState(48)
			p.Spec()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) SyntaxLit() ISyntaxLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxLitContext)
}

func (s *SpecContext) ImportSpec() IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *SpecContext) InfoSpec() IInfoSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfoSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfoSpecContext)
}

func (s *SpecContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpecContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserParserRULE_spec)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.CommentSpec()
		}

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

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetSyntaxToken sets the syntaxToken token.
	SetSyntaxToken(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsSyntaxLitContext differentiates from other interfaces.
	IsSyntaxLitContext()
}

type SyntaxLitContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	doc         ICommentSpecContext
	syntaxToken antlr.Token
	assign      antlr.Token
	version     antlr.Token
	comment     ICommentSpecContext
}

func NewEmptySyntaxLitContext() *SyntaxLitContext {
	var p = new(SyntaxLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_syntaxLit
	return p
}

func (*SyntaxLitContext) IsSyntaxLitContext() {}

func NewSyntaxLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxLitContext {
	var p = new(SyntaxLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_syntaxLit

	return p
}

func (s *SyntaxLitContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxLitContext) GetSyntaxToken() antlr.Token { return s.syntaxToken }

func (s *SyntaxLitContext) GetAssign() antlr.Token { return s.assign }

func (s *SyntaxLitContext) GetVersion() antlr.Token { return s.version }

func (s *SyntaxLitContext) SetSyntaxToken(v antlr.Token) { s.syntaxToken = v }

func (s *SyntaxLitContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *SyntaxLitContext) SetVersion(v antlr.Token) { s.version = v }

func (s *SyntaxLitContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *SyntaxLitContext) GetComment() ICommentSpecContext { return s.comment }

func (s *SyntaxLitContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *SyntaxLitContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *SyntaxLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *SyntaxLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *SyntaxLitContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *SyntaxLitContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) SyntaxLit() (localctx ISyntaxLitContext) {
	localctx = NewSyntaxLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ApiParserParserRULE_syntaxLit)
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
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(61)

			var _x = p.CommentSpec()

			localctx.(*SyntaxLitContext).doc = _x
		}

	}
	match(p, "syntax")
	{
		p.SetState(65)

		var _m = p.Match(ApiParserParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
	}
	{
		p.SetState(66)

		var _m = p.Match(ApiParserParserT__0)

		localctx.(*SyntaxLitContext).assign = _m
	}
	checkVersion(p)
	{
		p.SetState(68)

		var _m = p.Match(ApiParserParserSTRING)

		localctx.(*SyntaxLitContext).version = _m
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(69)

			var _x = p.CommentSpec()

			localctx.(*SyntaxLitContext).comment = _x
		}

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
	p.RuleIndex = ApiParserParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importSpec

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

func (s *ImportSpecContext) ImportBlock() IImportBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportBlockContext)
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

func (p *ApiParserParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ApiParserParserRULE_importSpec)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.ImportBlock()
		}

	}

	return localctx
}

// IImportLitContext is an interface to support dynamic dispatch.
type IImportLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportToken returns the importToken token.
	GetImportToken() antlr.Token

	// SetImportToken sets the importToken token.
	SetImportToken(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsImportLitContext differentiates from other interfaces.
	IsImportLitContext()
}

type ImportLitContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	doc         ICommentSpecContext
	importToken antlr.Token
	comment     ICommentSpecContext
}

func NewEmptyImportLitContext() *ImportLitContext {
	var p = new(ImportLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_importLit
	return p
}

func (*ImportLitContext) IsImportLitContext() {}

func NewImportLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLitContext {
	var p = new(ImportLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importLit

	return p
}

func (s *ImportLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportLitContext) GetImportToken() antlr.Token { return s.importToken }

func (s *ImportLitContext) SetImportToken(v antlr.Token) { s.importToken = v }

func (s *ImportLitContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *ImportLitContext) GetComment() ICommentSpecContext { return s.comment }

func (s *ImportLitContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *ImportLitContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *ImportLitContext) ImportValue() IImportValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportValueContext)
}

func (s *ImportLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ImportLitContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *ImportLitContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) ImportLit() (localctx IImportLitContext) {
	localctx = NewImportLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ApiParserParserRULE_importLit)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(76)

			var _x = p.CommentSpec()

			localctx.(*ImportLitContext).doc = _x
		}

	}
	match(p, "import")
	{
		p.SetState(80)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportLitContext).importToken = _m
	}
	{
		p.SetState(81)
		p.ImportValue()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(82)

			var _x = p.CommentSpec()

			localctx.(*ImportLitContext).comment = _x
		}

	}

	return localctx
}

// IImportBlockContext is an interface to support dynamic dispatch.
type IImportBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportToken returns the importToken token.
	GetImportToken() antlr.Token

	// SetImportToken sets the importToken token.
	SetImportToken(antlr.Token)

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsImportBlockContext differentiates from other interfaces.
	IsImportBlockContext()
}

type ImportBlockContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	importToken antlr.Token
	comment     ICommentSpecContext
}

func NewEmptyImportBlockContext() *ImportBlockContext {
	var p = new(ImportBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlock
	return p
}

func (*ImportBlockContext) IsImportBlockContext() {}

func NewImportBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockContext {
	var p = new(ImportBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importBlock

	return p
}

func (s *ImportBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportBlockContext) GetImportToken() antlr.Token { return s.importToken }

func (s *ImportBlockContext) SetImportToken(v antlr.Token) { s.importToken = v }

func (s *ImportBlockContext) GetComment() ICommentSpecContext { return s.comment }

func (s *ImportBlockContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *ImportBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ImportBlockContext) AllImportBlockValue() []IImportBlockValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportBlockValueContext)(nil)).Elem())
	var tst = make([]IImportBlockValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportBlockValueContext)
		}
	}

	return tst
}

func (s *ImportBlockContext) ImportBlockValue(i int) IImportBlockValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportBlockValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportBlockValueContext)
}

func (s *ImportBlockContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *ImportBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportBlock() (localctx IImportBlockContext) {
	localctx = NewImportBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ApiParserParserRULE_importBlock)
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
	match(p, "import")
	{
		p.SetState(86)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportBlockContext).importToken = _m
	}
	{
		p.SetState(87)
		p.Match(ApiParserParserT__1)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(88)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockContext).comment = _x
		}

	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserSTRING))) != 0) {
		{
			p.SetState(91)
			p.ImportBlockValue()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(ApiParserParserT__2)
	}

	return localctx
}

// IImportBlockValueContext is an interface to support dynamic dispatch.
type IImportBlockValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsImportBlockValueContext differentiates from other interfaces.
	IsImportBlockValueContext()
}

type ImportBlockValueContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	doc     ICommentSpecContext
	comment ICommentSpecContext
}

func NewEmptyImportBlockValueContext() *ImportBlockValueContext {
	var p = new(ImportBlockValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlockValue
	return p
}

func (*ImportBlockValueContext) IsImportBlockValueContext() {}

func NewImportBlockValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockValueContext {
	var p = new(ImportBlockValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importBlockValue

	return p
}

func (s *ImportBlockValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportBlockValueContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *ImportBlockValueContext) GetComment() ICommentSpecContext { return s.comment }

func (s *ImportBlockValueContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *ImportBlockValueContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *ImportBlockValueContext) ImportValue() IImportValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportValueContext)
}

func (s *ImportBlockValueContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *ImportBlockValueContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *ImportBlockValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportBlockValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportBlockValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportBlockValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportBlockValue() (localctx IImportBlockValueContext) {
	localctx = NewImportBlockValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApiParserParserRULE_importBlockValue)
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(98)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockValueContext).doc = _x
		}

	}
	{
		p.SetState(101)
		p.ImportValue()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(102)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockValueContext).comment = _x
		}

	}

	return localctx
}

// IImportValueContext is an interface to support dynamic dispatch.
type IImportValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportValueContext differentiates from other interfaces.
	IsImportValueContext()
}

type ImportValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportValueContext() *ImportValueContext {
	var p = new(ImportValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_importValue
	return p
}

func (*ImportValueContext) IsImportValueContext() {}

func NewImportValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportValueContext {
	var p = new(ImportValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importValue

	return p
}

func (s *ImportValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *ImportValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportValue() (localctx IImportValueContext) {
	localctx = NewImportValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApiParserParserRULE_importValue)

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
	checkImportValue(p)
	{
		p.SetState(106)
		p.Match(ApiParserParserSTRING)
	}

	return localctx
}

// IInfoSpecContext is an interface to support dynamic dispatch.
type IInfoSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInfoToken returns the infoToken token.
	GetInfoToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetInfoToken sets the infoToken token.
	SetInfoToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsInfoSpecContext differentiates from other interfaces.
	IsInfoSpecContext()
}

type InfoSpecContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	doc       ICommentSpecContext
	infoToken antlr.Token
	lp        antlr.Token
	comment   ICommentSpecContext
	rp        antlr.Token
}

func NewEmptyInfoSpecContext() *InfoSpecContext {
	var p = new(InfoSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_infoSpec
	return p
}

func (*InfoSpecContext) IsInfoSpecContext() {}

func NewInfoSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoSpecContext {
	var p = new(InfoSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_infoSpec

	return p
}

func (s *InfoSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *InfoSpecContext) GetInfoToken() antlr.Token { return s.infoToken }

func (s *InfoSpecContext) GetLp() antlr.Token { return s.lp }

func (s *InfoSpecContext) GetRp() antlr.Token { return s.rp }

func (s *InfoSpecContext) SetInfoToken(v antlr.Token) { s.infoToken = v }

func (s *InfoSpecContext) SetLp(v antlr.Token) { s.lp = v }

func (s *InfoSpecContext) SetRp(v antlr.Token) { s.rp = v }

func (s *InfoSpecContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *InfoSpecContext) GetComment() ICommentSpecContext { return s.comment }

func (s *InfoSpecContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *InfoSpecContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *InfoSpecContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *InfoSpecContext) AllKvLit() []IKvLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvLitContext)(nil)).Elem())
	var tst = make([]IKvLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvLitContext)
		}
	}

	return tst
}

func (s *InfoSpecContext) KvLit(i int) IKvLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *InfoSpecContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *InfoSpecContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *InfoSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfoSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfoSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitInfoSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) InfoSpec() (localctx IInfoSpecContext) {
	localctx = NewInfoSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ApiParserParserRULE_infoSpec)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(108)

			var _x = p.CommentSpec()

			localctx.(*InfoSpecContext).doc = _x
		}

	}
	match(p, "info")
	{
		p.SetState(112)

		var _m = p.Match(ApiParserParserID)

		localctx.(*InfoSpecContext).infoToken = _m
	}
	{
		p.SetState(113)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*InfoSpecContext).lp = _m
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)

			var _x = p.CommentSpec()

			localctx.(*InfoSpecContext).comment = _x
		}

	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(117)
			p.KvLit()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*InfoSpecContext).rp = _m
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
	p.RuleIndex = ApiParserParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) TypeLit() ITypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeSpecContext) TypeBlock() ITypeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBlockContext)
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

func (p *ApiParserParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ApiParserParserRULE_typeSpec)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.TypeBlock()
		}

	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	doc       ICommentSpecContext
	typeToken antlr.Token
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) GetTypeToken() antlr.Token { return s.typeToken }

func (s *TypeLitContext) SetTypeToken(v antlr.Token) { s.typeToken = v }

func (s *TypeLitContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *TypeLitContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *TypeLitContext) TypeBody() ITypeBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBodyContext)
}

func (s *TypeLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeLitContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ApiParserParserRULE_typeLit)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(128)

			var _x = p.CommentSpec()

			localctx.(*TypeLitContext).doc = _x
		}

	}
	match(p, "type")
	{
		p.SetState(132)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeLitContext).typeToken = _m
	}
	{
		p.SetState(133)
		p.TypeBody()
	}

	return localctx
}

// ITypeBlockContext is an interface to support dynamic dispatch.
type ITypeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// IsTypeBlockContext differentiates from other interfaces.
	IsTypeBlockContext()
}

type TypeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	doc       ICommentSpecContext
	typeToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyTypeBlockContext() *TypeBlockContext {
	var p = new(TypeBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlock
	return p
}

func (*TypeBlockContext) IsTypeBlockContext() {}

func NewTypeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockContext {
	var p = new(TypeBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlock

	return p
}

func (s *TypeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockContext) GetTypeToken() antlr.Token { return s.typeToken }

func (s *TypeBlockContext) GetLp() antlr.Token { return s.lp }

func (s *TypeBlockContext) GetRp() antlr.Token { return s.rp }

func (s *TypeBlockContext) SetTypeToken(v antlr.Token) { s.typeToken = v }

func (s *TypeBlockContext) SetLp(v antlr.Token) { s.lp = v }

func (s *TypeBlockContext) SetRp(v antlr.Token) { s.rp = v }

func (s *TypeBlockContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *TypeBlockContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *TypeBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeBlockContext) AllTypeBody() []ITypeBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeBodyContext)(nil)).Elem())
	var tst = make([]ITypeBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeBodyContext)
		}
	}

	return tst
}

func (s *TypeBlockContext) TypeBody(i int) ITypeBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeBodyContext)
}

func (s *TypeBlockContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) TypeBlock() (localctx ITypeBlockContext) {
	localctx = NewTypeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ApiParserParserRULE_typeBlock)
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
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(135)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockContext).doc = _x
		}

	}
	match(p, "type")
	{
		p.SetState(139)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockContext).typeToken = _m
	}
	{
		p.SetState(140)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*TypeBlockContext).lp = _m
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(141)
			p.TypeBody()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*TypeBlockContext).rp = _m
	}

	return localctx
}

// ITypeBodyContext is an interface to support dynamic dispatch.
type ITypeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBodyContext differentiates from other interfaces.
	IsTypeBodyContext()
}

type TypeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBodyContext() *TypeBodyContext {
	var p = new(TypeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBody
	return p
}

func (*TypeBodyContext) IsTypeBodyContext() {}

func NewTypeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBodyContext {
	var p = new(TypeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBody

	return p
}

func (s *TypeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBodyContext) TypeAlias() ITypeAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAliasContext)
}

func (s *TypeBodyContext) TypeStruct() ITypeStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
}

func (s *TypeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBody() (localctx ITypeBodyContext) {
	localctx = NewTypeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ApiParserParserRULE_typeBody)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.TypeAlias()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.TypeStruct()
		}

	}

	return localctx
}

// ITypeStructContext is an interface to support dynamic dispatch.
type ITypeStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// GetStructToken returns the structToken token.
	GetStructToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// SetStructToken sets the structToken token.
	SetStructToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsTypeStructContext differentiates from other interfaces.
	IsTypeStructContext()
}

type TypeStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	comment     ICommentSpecContext
	rbrace      antlr.Token
}

func NewEmptyTypeStructContext() *TypeStructContext {
	var p = new(TypeStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeStruct
	return p
}

func (*TypeStructContext) IsTypeStructContext() {}

func NewTypeStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStructContext {
	var p = new(TypeStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeStruct

	return p
}

func (s *TypeStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeStructContext) GetComment() ICommentSpecContext { return s.comment }

func (s *TypeStructContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *TypeStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeStructContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *TypeStructContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeStructContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) TypeStruct() (localctx ITypeStructContext) {
	localctx = NewTypeStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ApiParserParserRULE_typeStruct)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	checkFieldName(p)
	{
		p.SetState(153)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeStructContext).structName = _m
	}
	match(p, "struct")
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(155)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeStructContext).structToken = _m
		}

	}
	{
		p.SetState(158)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeStructContext).lbrace = _m
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(159)

			var _x = p.CommentSpec()

			localctx.(*TypeStructContext).comment = _x
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(162)
				p.Field()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	{
		p.SetState(167)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeStructContext).rbrace = _m
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

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsTypeAliasContext differentiates from other interfaces.
	IsTypeAliasContext()
}

type TypeAliasContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	alias   antlr.Token
	assign  antlr.Token
	comment ICommentSpecContext
}

func NewEmptyTypeAliasContext() *TypeAliasContext {
	var p = new(TypeAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeAlias
	return p
}

func (*TypeAliasContext) IsTypeAliasContext() {}

func NewTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasContext {
	var p = new(TypeAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeAlias

	return p
}

func (s *TypeAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeAliasContext) GetAssign() antlr.Token { return s.assign }

func (s *TypeAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeAliasContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *TypeAliasContext) GetComment() ICommentSpecContext { return s.comment }

func (s *TypeAliasContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *TypeAliasContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeAliasContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

func (p *ApiParserParser) TypeAlias() (localctx ITypeAliasContext) {
	localctx = NewTypeAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ApiParserParserRULE_typeAlias)
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
	checkFieldName(p)
	{
		p.SetState(170)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(171)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeAliasContext).assign = _m
		}

	}
	{
		p.SetState(174)
		p.DataType()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)

			var _x = p.CommentSpec()

			localctx.(*TypeAliasContext).comment = _x
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AnonymousFiled() IAnonymousFiledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnonymousFiledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnonymousFiledContext)
}

func (s *FieldContext) NormalField() INormalFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INormalFieldContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ApiParserParserRULE_field)

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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(178)

		if !(isAnonymous(p)) {
			panic(antlr.NewFailedPredicateException(p, "isAnonymous(p)", ""))
		}
		{
			p.SetState(179)
			p.AnonymousFiled()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.NormalField()
		}

	}

	return localctx
}

// INormalFieldContext is an interface to support dynamic dispatch.
type INormalFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFieldName returns the fieldName token.
	GetFieldName() antlr.Token

	// GetTag returns the tag token.
	GetTag() antlr.Token

	// SetFieldName sets the fieldName token.
	SetFieldName(antlr.Token)

	// SetTag sets the tag token.
	SetTag(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsNormalFieldContext differentiates from other interfaces.
	IsNormalFieldContext()
}

type NormalFieldContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	doc       ICommentSpecContext
	fieldName antlr.Token
	tag       antlr.Token
	comment   ICommentSpecContext
}

func NewEmptyNormalFieldContext() *NormalFieldContext {
	var p = new(NormalFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_normalField
	return p
}

func (*NormalFieldContext) IsNormalFieldContext() {}

func NewNormalFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalFieldContext {
	var p = new(NormalFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_normalField

	return p
}

func (s *NormalFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalFieldContext) GetFieldName() antlr.Token { return s.fieldName }

func (s *NormalFieldContext) GetTag() antlr.Token { return s.tag }

func (s *NormalFieldContext) SetFieldName(v antlr.Token) { s.fieldName = v }

func (s *NormalFieldContext) SetTag(v antlr.Token) { s.tag = v }

func (s *NormalFieldContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *NormalFieldContext) GetComment() ICommentSpecContext { return s.comment }

func (s *NormalFieldContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *NormalFieldContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *NormalFieldContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *NormalFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *NormalFieldContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *NormalFieldContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *NormalFieldContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserRAW_STRING, 0)
}

func (s *NormalFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitNormalField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) NormalField() (localctx INormalFieldContext) {
	localctx = NewNormalFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ApiParserParserRULE_normalField)
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
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(183)

			var _x = p.CommentSpec()

			localctx.(*NormalFieldContext).doc = _x
		}

	}
	{
		p.SetState(186)

		var _m = p.Match(ApiParserParserID)

		localctx.(*NormalFieldContext).fieldName = _m
	}
	{
		p.SetState(187)
		p.DataType()
	}
	checkTag(p)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(189)

			var _m = p.Match(ApiParserParserRAW_STRING)

			localctx.(*NormalFieldContext).tag = _m
		}

	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(192)

			var _x = p.CommentSpec()

			localctx.(*NormalFieldContext).comment = _x
		}

	}

	return localctx
}

// IAnonymousFiledContext is an interface to support dynamic dispatch.
type IAnonymousFiledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsAnonymousFiledContext differentiates from other interfaces.
	IsAnonymousFiledContext()
}

type AnonymousFiledContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	doc     ICommentSpecContext
	star    antlr.Token
	comment ICommentSpecContext
}

func NewEmptyAnonymousFiledContext() *AnonymousFiledContext {
	var p = new(AnonymousFiledContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_anonymousFiled
	return p
}

func (*AnonymousFiledContext) IsAnonymousFiledContext() {}

func NewAnonymousFiledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFiledContext {
	var p = new(AnonymousFiledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_anonymousFiled

	return p
}

func (s *AnonymousFiledContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonymousFiledContext) GetStar() antlr.Token { return s.star }

func (s *AnonymousFiledContext) SetStar(v antlr.Token) { s.star = v }

func (s *AnonymousFiledContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *AnonymousFiledContext) GetComment() ICommentSpecContext { return s.comment }

func (s *AnonymousFiledContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *AnonymousFiledContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *AnonymousFiledContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *AnonymousFiledContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *AnonymousFiledContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *AnonymousFiledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonymousFiledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonymousFiledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAnonymousFiled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AnonymousFiled() (localctx IAnonymousFiledContext) {
	localctx = NewAnonymousFiledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ApiParserParserRULE_anonymousFiled)
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
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(195)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).doc = _x
		}

	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__5 {
		{
			p.SetState(198)

			var _m = p.Match(ApiParserParserT__5)

			localctx.(*AnonymousFiledContext).star = _m
		}

	}
	{
		p.SetState(201)
		p.Match(ApiParserParserID)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(202)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).comment = _x
		}

	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInter returns the inter token.
	GetInter() antlr.Token

	// GetTime returns the time token.
	GetTime() antlr.Token

	// SetInter sets the inter token.
	SetInter(antlr.Token)

	// SetTime sets the time token.
	SetTime(antlr.Token)

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inter  antlr.Token
	time   antlr.Token
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) GetInter() antlr.Token { return s.inter }

func (s *DataTypeContext) GetTime() antlr.Token { return s.time }

func (s *DataTypeContext) SetInter(v antlr.Token) { s.inter = v }

func (s *DataTypeContext) SetTime(v antlr.Token) { s.time = v }

func (s *DataTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
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
	return s.GetToken(ApiParserParserINTERFACE, 0)
}

func (s *DataTypeContext) PointerType() IPointerTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *DataTypeContext) TypeStruct() ITypeStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
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

func (p *ApiParserParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ApiParserParserRULE_dataType)

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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		isInterface(p)
		{
			p.SetState(206)
			p.Match(ApiParserParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.MapType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)

			var _m = p.Match(ApiParserParserINTERFACE)

			localctx.(*DataTypeContext).inter = _m
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(210)

			var _m = p.Match(ApiParserParserT__6)

			localctx.(*DataTypeContext).time = _m
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(211)
			p.PointerType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(212)
			p.TypeStruct()
		}

	}

	return localctx
}

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_pointerType
	return p
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) GetStar() antlr.Token { return s.star }

func (s *PointerTypeContext) SetStar(v antlr.Token) { s.star = v }

func (s *PointerTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPointerType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ApiParserParserRULE_pointerType)

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
		p.SetState(215)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*PointerTypeContext).star = _m
	}
	{
		p.SetState(216)
		p.Match(ApiParserParserID)
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMapToken returns the mapToken token.
	GetMapToken() antlr.Token

	// GetLbrack returns the lbrack token.
	GetLbrack() antlr.Token

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetRbrack returns the rbrack token.
	GetRbrack() antlr.Token

	// SetMapToken sets the mapToken token.
	SetMapToken(antlr.Token)

	// SetLbrack sets the lbrack token.
	SetLbrack(antlr.Token)

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetRbrack sets the rbrack token.
	SetRbrack(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IDataTypeContext

	// SetValue sets the value rule contexts.
	SetValue(IDataTypeContext)

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	mapToken antlr.Token
	lbrack   antlr.Token
	key      antlr.Token
	rbrack   antlr.Token
	value    IDataTypeContext
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) GetMapToken() antlr.Token { return s.mapToken }

func (s *MapTypeContext) GetLbrack() antlr.Token { return s.lbrack }

func (s *MapTypeContext) GetKey() antlr.Token { return s.key }

func (s *MapTypeContext) GetRbrack() antlr.Token { return s.rbrack }

func (s *MapTypeContext) SetMapToken(v antlr.Token) { s.mapToken = v }

func (s *MapTypeContext) SetLbrack(v antlr.Token) { s.lbrack = v }

func (s *MapTypeContext) SetKey(v antlr.Token) { s.key = v }

func (s *MapTypeContext) SetRbrack(v antlr.Token) { s.rbrack = v }

func (s *MapTypeContext) GetValue() IDataTypeContext { return s.value }

func (s *MapTypeContext) SetValue(v IDataTypeContext) { s.value = v }

func (s *MapTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *MapTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
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

func (p *ApiParserParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ApiParserParserRULE_mapType)

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
	match(p, "map")
	{
		p.SetState(219)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).mapToken = _m
	}
	{
		p.SetState(220)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*MapTypeContext).lbrack = _m
	}
	{
		p.SetState(221)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).key = _m
	}
	{
		p.SetState(222)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*MapTypeContext).rbrack = _m
	}
	{
		p.SetState(223)

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

	// GetLbrack returns the lbrack token.
	GetLbrack() antlr.Token

	// GetRbrack returns the rbrack token.
	GetRbrack() antlr.Token

	// SetLbrack sets the lbrack token.
	SetLbrack(antlr.Token)

	// SetRbrack sets the rbrack token.
	SetRbrack(antlr.Token)

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lbrack antlr.Token
	rbrack antlr.Token
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) GetLbrack() antlr.Token { return s.lbrack }

func (s *ArrayTypeContext) GetRbrack() antlr.Token { return s.rbrack }

func (s *ArrayTypeContext) SetLbrack(v antlr.Token) { s.lbrack = v }

func (s *ArrayTypeContext) SetRbrack(v antlr.Token) { s.rbrack = v }

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

func (p *ApiParserParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ApiParserParserRULE_arrayType)

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
		p.SetState(225)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*ArrayTypeContext).lbrack = _m
	}
	{
		p.SetState(226)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*ArrayTypeContext).rbrack = _m
	}
	{
		p.SetState(227)
		p.DataType()
	}

	return localctx
}

// IKvLitContext is an interface to support dynamic dispatch.
type IKvLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsKvLitContext differentiates from other interfaces.
	IsKvLitContext()
}

type KvLitContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	doc     ICommentSpecContext
	key     antlr.Token
	value   antlr.Token
	comment ICommentSpecContext
}

func NewEmptyKvLitContext() *KvLitContext {
	var p = new(KvLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_kvLit
	return p
}

func (*KvLitContext) IsKvLitContext() {}

func NewKvLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvLitContext {
	var p = new(KvLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_kvLit

	return p
}

func (s *KvLitContext) GetParser() antlr.Parser { return s.parser }

func (s *KvLitContext) GetKey() antlr.Token { return s.key }

func (s *KvLitContext) GetValue() antlr.Token { return s.value }

func (s *KvLitContext) SetKey(v antlr.Token) { s.key = v }

func (s *KvLitContext) SetValue(v antlr.Token) { s.value = v }

func (s *KvLitContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *KvLitContext) GetComment() ICommentSpecContext { return s.comment }

func (s *KvLitContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *KvLitContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *KvLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *KvLitContext) LINE_VALUE() antlr.TerminalNode {
	return s.GetToken(ApiParserParserLINE_VALUE, 0)
}

func (s *KvLitContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *KvLitContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *KvLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitKvLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) KvLit() (localctx IKvLitContext) {
	localctx = NewKvLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ApiParserParserRULE_kvLit)
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
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(229)

			var _x = p.CommentSpec()

			localctx.(*KvLitContext).doc = _x
		}

	}
	{
		p.SetState(232)

		var _m = p.Match(ApiParserParserID)

		localctx.(*KvLitContext).key = _m
	}
	checkKeyValue(p)
	{
		p.SetState(234)

		var _m = p.Match(ApiParserParserLINE_VALUE)

		localctx.(*KvLitContext).value = _m
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(235)

			var _x = p.CommentSpec()

			localctx.(*KvLitContext).comment = _x
		}

	}

	return localctx
}

// ICommentSpecContext is an interface to support dynamic dispatch.
type ICommentSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentSpecContext differentiates from other interfaces.
	IsCommentSpecContext()
}

type CommentSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentSpecContext() *CommentSpecContext {
	var p = new(CommentSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_commentSpec
	return p
}

func (*CommentSpecContext) IsCommentSpecContext() {}

func NewCommentSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentSpecContext {
	var p = new(CommentSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_commentSpec

	return p
}

func (s *CommentSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentSpecContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ApiParserParserCOMMENT, 0)
}

func (s *CommentSpecContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(ApiParserParserLINE_COMMENT, 0)
}

func (s *CommentSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitCommentSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) CommentSpec() (localctx ICommentSpecContext) {
	localctx = NewCommentSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ApiParserParserRULE_commentSpec)
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
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *ApiParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ApiParserParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return isAnonymous(p)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
