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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 282,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 7, 2, 58, 10, 2, 12, 2, 14, 2, 61, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 68, 10, 3, 3, 4, 5, 4, 71, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 79, 10, 4, 3, 5, 3, 5, 5, 5, 83, 10, 5, 3, 6, 5, 6, 86, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 92, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 98, 10, 7, 3, 7, 6, 7, 101, 10, 7, 13, 7, 14, 7, 102, 3, 7, 3, 7, 3,
	8, 5, 8, 108, 10, 8, 3, 8, 3, 8, 5, 8, 112, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	10, 5, 10, 118, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 124, 10, 10,
	3, 10, 6, 10, 127, 10, 10, 13, 10, 14, 10, 128, 3, 10, 3, 10, 3, 11, 3,
	11, 5, 11, 135, 10, 11, 3, 12, 5, 12, 138, 10, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 6, 13, 148, 10, 13, 13, 13, 14, 13,
	149, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 156, 10, 14, 3, 15, 3, 15, 5, 15,
	160, 10, 15, 3, 16, 3, 16, 3, 16, 5, 16, 165, 10, 16, 3, 16, 3, 16, 5,
	16, 169, 10, 16, 3, 16, 6, 16, 172, 10, 16, 13, 16, 14, 16, 173, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 181, 10, 17, 3, 17, 3, 17, 5, 17, 185,
	10, 17, 3, 18, 5, 18, 188, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 193, 10,
	18, 3, 18, 3, 18, 5, 18, 197, 10, 18, 3, 18, 6, 18, 200, 10, 18, 13, 18,
	14, 18, 201, 3, 18, 3, 18, 3, 19, 5, 19, 207, 10, 19, 3, 19, 3, 19, 3,
	19, 5, 19, 212, 10, 19, 3, 19, 3, 19, 5, 19, 216, 10, 19, 3, 20, 3, 20,
	3, 20, 5, 20, 221, 10, 20, 3, 21, 5, 21, 224, 10, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 5, 21, 230, 10, 21, 3, 21, 5, 21, 233, 10, 21, 3, 22, 5, 22,
	236, 10, 22, 3, 22, 5, 22, 239, 10, 22, 3, 22, 3, 22, 5, 22, 243, 10, 22,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 253, 10,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 5, 27, 272, 10, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 5, 27, 278, 10, 27, 3, 28, 3, 28, 3, 28, 2, 2,
	29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 2, 3, 3, 2, 17, 18, 2, 303, 2, 59,
	3, 2, 2, 2, 4, 67, 3, 2, 2, 2, 6, 70, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10,
	85, 3, 2, 2, 2, 12, 93, 3, 2, 2, 2, 14, 107, 3, 2, 2, 2, 16, 113, 3, 2,
	2, 2, 18, 117, 3, 2, 2, 2, 20, 134, 3, 2, 2, 2, 22, 137, 3, 2, 2, 2, 24,
	143, 3, 2, 2, 2, 26, 155, 3, 2, 2, 2, 28, 159, 3, 2, 2, 2, 30, 161, 3,
	2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 187, 3, 2, 2, 2, 36, 206, 3, 2, 2, 2,
	38, 220, 3, 2, 2, 2, 40, 223, 3, 2, 2, 2, 42, 235, 3, 2, 2, 2, 44, 252,
	3, 2, 2, 2, 46, 254, 3, 2, 2, 2, 48, 258, 3, 2, 2, 2, 50, 266, 3, 2, 2,
	2, 52, 271, 3, 2, 2, 2, 54, 279, 3, 2, 2, 2, 56, 58, 5, 4, 3, 2, 57, 56,
	3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 3, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 68, 5, 6, 4, 2, 63, 68, 5, 8,
	5, 2, 64, 68, 5, 18, 10, 2, 65, 68, 5, 20, 11, 2, 66, 68, 5, 54, 28, 2,
	67, 62, 3, 2, 2, 2, 67, 63, 3, 2, 2, 2, 67, 64, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 5, 3, 2, 2, 2, 69, 71, 5, 54, 28, 2, 70,
	69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 8, 4, 1,
	2, 73, 74, 7, 22, 2, 2, 74, 75, 7, 3, 2, 2, 75, 76, 8, 4, 1, 2, 76, 78,
	7, 19, 2, 2, 77, 79, 5, 54, 28, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2,
	2, 79, 7, 3, 2, 2, 2, 80, 83, 5, 10, 6, 2, 81, 83, 5, 12, 7, 2, 82, 80,
	3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 9, 3, 2, 2, 2, 84, 86, 5, 54, 28, 2,
	85, 84, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 8,
	6, 1, 2, 88, 89, 7, 22, 2, 2, 89, 91, 5, 16, 9, 2, 90, 92, 5, 54, 28, 2,
	91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 11, 3, 2, 2, 2, 93, 94, 8,
	7, 1, 2, 94, 95, 7, 22, 2, 2, 95, 97, 7, 4, 2, 2, 96, 98, 5, 54, 28, 2,
	97, 96, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 101, 5,
	14, 8, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2,
	2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 7, 5, 2, 2, 105,
	13, 3, 2, 2, 2, 106, 108, 5, 54, 28, 2, 107, 106, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 5, 16, 9, 2, 110, 112, 5, 54,
	28, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 15, 3, 2, 2, 2,
	113, 114, 8, 9, 1, 2, 114, 115, 7, 19, 2, 2, 115, 17, 3, 2, 2, 2, 116,
	118, 5, 54, 28, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 120, 8, 10, 1, 2, 120, 121, 7, 22, 2, 2, 121, 123, 7,
	4, 2, 2, 122, 124, 5, 54, 28, 2, 123, 122, 3, 2, 2, 2, 123, 124, 3, 2,
	2, 2, 124, 126, 3, 2, 2, 2, 125, 127, 5, 52, 27, 2, 126, 125, 3, 2, 2,
	2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	130, 3, 2, 2, 2, 130, 131, 7, 5, 2, 2, 131, 19, 3, 2, 2, 2, 132, 135, 5,
	22, 12, 2, 133, 135, 5, 24, 13, 2, 134, 132, 3, 2, 2, 2, 134, 133, 3, 2,
	2, 2, 135, 21, 3, 2, 2, 2, 136, 138, 5, 54, 28, 2, 137, 136, 3, 2, 2, 2,
	137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 8, 12, 1, 2, 140,
	141, 7, 22, 2, 2, 141, 142, 5, 26, 14, 2, 142, 23, 3, 2, 2, 2, 143, 144,
	8, 13, 1, 2, 144, 145, 7, 22, 2, 2, 145, 147, 7, 4, 2, 2, 146, 148, 5,
	28, 15, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 147, 3, 2,
	2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 7, 5, 2, 2,
	152, 25, 3, 2, 2, 2, 153, 156, 5, 30, 16, 2, 154, 156, 5, 32, 17, 2, 155,
	153, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 27, 3, 2, 2, 2, 157, 160, 5,
	34, 18, 2, 158, 160, 5, 36, 19, 2, 159, 157, 3, 2, 2, 2, 159, 158, 3, 2,
	2, 2, 160, 29, 3, 2, 2, 2, 161, 162, 8, 16, 1, 2, 162, 164, 7, 22, 2, 2,
	163, 165, 7, 22, 2, 2, 164, 163, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165,
	166, 3, 2, 2, 2, 166, 168, 7, 6, 2, 2, 167, 169, 5, 54, 28, 2, 168, 167,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 172, 5, 38,
	20, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2,
	173, 174, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 7, 2, 2, 176,
	31, 3, 2, 2, 2, 177, 178, 8, 17, 1, 2, 178, 180, 7, 22, 2, 2, 179, 181,
	7, 3, 2, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2,
	2, 2, 182, 184, 5, 44, 23, 2, 183, 185, 5, 54, 28, 2, 184, 183, 3, 2, 2,
	2, 184, 185, 3, 2, 2, 2, 185, 33, 3, 2, 2, 2, 186, 188, 5, 54, 28, 2, 187,
	186, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190,
	8, 18, 1, 2, 190, 192, 7, 22, 2, 2, 191, 193, 7, 22, 2, 2, 192, 191, 3,
	2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 7, 6, 2,
	2, 195, 197, 5, 54, 28, 2, 196, 195, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2,
	197, 199, 3, 2, 2, 2, 198, 200, 5, 38, 20, 2, 199, 198, 3, 2, 2, 2, 200,
	201, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 204, 7, 7, 2, 2, 204, 35, 3, 2, 2, 2, 205, 207, 5, 54,
	28, 2, 206, 205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2,
	208, 209, 8, 19, 1, 2, 209, 211, 7, 22, 2, 2, 210, 212, 7, 3, 2, 2, 211,
	210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215,
	5, 44, 23, 2, 214, 216, 5, 54, 28, 2, 215, 214, 3, 2, 2, 2, 215, 216, 3,
	2, 2, 2, 216, 37, 3, 2, 2, 2, 217, 218, 6, 20, 2, 2, 218, 221, 5, 40, 21,
	2, 219, 221, 5, 42, 22, 2, 220, 217, 3, 2, 2, 2, 220, 219, 3, 2, 2, 2,
	221, 39, 3, 2, 2, 2, 222, 224, 5, 54, 28, 2, 223, 222, 3, 2, 2, 2, 223,
	224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 226, 8, 21, 1, 2, 226, 227,
	7, 22, 2, 2, 227, 229, 5, 44, 23, 2, 228, 230, 7, 20, 2, 2, 229, 228, 3,
	2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232, 3, 2, 2, 2, 231, 233, 5, 54, 28,
	2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 41, 3, 2, 2, 2, 234,
	236, 5, 54, 28, 2, 235, 234, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 238,
	3, 2, 2, 2, 237, 239, 7, 8, 2, 2, 238, 237, 3, 2, 2, 2, 238, 239, 3, 2,
	2, 2, 239, 240, 3, 2, 2, 2, 240, 242, 7, 22, 2, 2, 241, 243, 5, 54, 28,
	2, 242, 241, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 43, 3, 2, 2, 2, 244,
	245, 8, 23, 1, 2, 245, 253, 7, 22, 2, 2, 246, 253, 5, 48, 25, 2, 247, 253,
	5, 50, 26, 2, 248, 253, 7, 14, 2, 2, 249, 253, 7, 9, 2, 2, 250, 253, 5,
	46, 24, 2, 251, 253, 5, 30, 16, 2, 252, 244, 3, 2, 2, 2, 252, 246, 3, 2,
	2, 2, 252, 247, 3, 2, 2, 2, 252, 248, 3, 2, 2, 2, 252, 249, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 251, 3, 2, 2, 2, 253, 45, 3, 2, 2, 2, 254, 255,
	7, 8, 2, 2, 255, 256, 8, 24, 1, 2, 256, 257, 7, 22, 2, 2, 257, 47, 3, 2,
	2, 2, 258, 259, 8, 25, 1, 2, 259, 260, 7, 22, 2, 2, 260, 261, 7, 10, 2,
	2, 261, 262, 8, 25, 1, 2, 262, 263, 7, 22, 2, 2, 263, 264, 7, 11, 2, 2,
	264, 265, 5, 44, 23, 2, 265, 49, 3, 2, 2, 2, 266, 267, 7, 10, 2, 2, 267,
	268, 7, 11, 2, 2, 268, 269, 5, 44, 23, 2, 269, 51, 3, 2, 2, 2, 270, 272,
	5, 54, 28, 2, 271, 270, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 273, 3,
	2, 2, 2, 273, 274, 7, 22, 2, 2, 274, 275, 8, 27, 1, 2, 275, 277, 7, 21,
	2, 2, 276, 278, 5, 54, 28, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2,
	2, 278, 53, 3, 2, 2, 2, 279, 280, 9, 2, 2, 2, 280, 55, 3, 2, 2, 2, 43,
	59, 67, 70, 78, 82, 85, 91, 97, 102, 107, 111, 117, 123, 128, 134, 137,
	149, 155, 159, 164, 168, 173, 180, 184, 187, 192, 196, 201, 206, 211, 215,
	220, 223, 229, 232, 235, 238, 242, 252, 271, 277,
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
	"importValue", "infoSpec", "typeSpec", "typeLit", "typeBlock", "typeLitBody",
	"typeBlockBody", "typeStruct", "typeAlias", "typeBlockStruct", "typeBlockAlias",
	"field", "normalField", "anonymousFiled", "dataType", "pointerType", "mapType",
	"arrayType", "kvLit", "commentSpec",
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
	ApiParserParserRULE_typeLitBody      = 12
	ApiParserParserRULE_typeBlockBody    = 13
	ApiParserParserRULE_typeStruct       = 14
	ApiParserParserRULE_typeAlias        = 15
	ApiParserParserRULE_typeBlockStruct  = 16
	ApiParserParserRULE_typeBlockAlias   = 17
	ApiParserParserRULE_field            = 18
	ApiParserParserRULE_normalField      = 19
	ApiParserParserRULE_anonymousFiled   = 20
	ApiParserParserRULE_dataType         = 21
	ApiParserParserRULE_pointerType      = 22
	ApiParserParserRULE_mapType          = 23
	ApiParserParserRULE_arrayType        = 24
	ApiParserParserRULE_kvLit            = 25
	ApiParserParserRULE_commentSpec      = 26
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0 {
		{
			p.SetState(54)
			p.Spec()
		}

		p.SetState(59)
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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(64)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(67)

			var _x = p.CommentSpec()

			localctx.(*SyntaxLitContext).doc = _x
		}

	}
	match(p, "syntax")
	{
		p.SetState(71)

		var _m = p.Match(ApiParserParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
	}
	{
		p.SetState(72)

		var _m = p.Match(ApiParserParserT__0)

		localctx.(*SyntaxLitContext).assign = _m
	}
	checkVersion(p)
	{
		p.SetState(74)

		var _m = p.Match(ApiParserParserSTRING)

		localctx.(*SyntaxLitContext).version = _m
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(75)

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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(82)

			var _x = p.CommentSpec()

			localctx.(*ImportLitContext).doc = _x
		}

	}
	match(p, "import")
	{
		p.SetState(86)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportLitContext).importToken = _m
	}
	{
		p.SetState(87)
		p.ImportValue()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(88)

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
		p.SetState(92)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportBlockContext).importToken = _m
	}
	{
		p.SetState(93)
		p.Match(ApiParserParserT__1)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(94)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockContext).comment = _x
		}

	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserSTRING))) != 0) {
		{
			p.SetState(97)
			p.ImportBlockValue()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(104)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockValueContext).doc = _x
		}

	}
	{
		p.SetState(107)
		p.ImportValue()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(108)

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
		p.SetState(112)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(114)

			var _x = p.CommentSpec()

			localctx.(*InfoSpecContext).doc = _x
		}

	}
	match(p, "info")
	{
		p.SetState(118)

		var _m = p.Match(ApiParserParserID)

		localctx.(*InfoSpecContext).infoToken = _m
	}
	{
		p.SetState(119)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*InfoSpecContext).lp = _m
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(120)

			var _x = p.CommentSpec()

			localctx.(*InfoSpecContext).comment = _x
		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(123)
			p.KvLit()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(128)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
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

func (s *TypeLitContext) TypeLitBody() ITypeLitBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitBodyContext)
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
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(134)

			var _x = p.CommentSpec()

			localctx.(*TypeLitContext).doc = _x
		}

	}
	match(p, "type")
	{
		p.SetState(138)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeLitContext).typeToken = _m
	}
	{
		p.SetState(139)
		p.TypeLitBody()
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

	// IsTypeBlockContext differentiates from other interfaces.
	IsTypeBlockContext()
}

type TypeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
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

func (s *TypeBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeBlockContext) AllTypeBlockBody() []ITypeBlockBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeBlockBodyContext)(nil)).Elem())
	var tst = make([]ITypeBlockBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeBlockBodyContext)
		}
	}

	return tst
}

func (s *TypeBlockContext) TypeBlockBody(i int) ITypeBlockBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBlockBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeBlockBodyContext)
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
	match(p, "type")
	{
		p.SetState(142)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockContext).typeToken = _m
	}
	{
		p.SetState(143)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*TypeBlockContext).lp = _m
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(144)
			p.TypeBlockBody()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*TypeBlockContext).rp = _m
	}

	return localctx
}

// ITypeLitBodyContext is an interface to support dynamic dispatch.
type ITypeLitBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLitBodyContext differentiates from other interfaces.
	IsTypeLitBodyContext()
}

type TypeLitBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitBodyContext() *TypeLitBodyContext {
	var p = new(TypeLitBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLitBody
	return p
}

func (*TypeLitBodyContext) IsTypeLitBodyContext() {}

func NewTypeLitBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitBodyContext {
	var p = new(TypeLitBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeLitBody

	return p
}

func (s *TypeLitBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitBodyContext) TypeStruct() ITypeStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
}

func (s *TypeLitBodyContext) TypeAlias() ITypeAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAliasContext)
}

func (s *TypeLitBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeLitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeLitBody() (localctx ITypeLitBodyContext) {
	localctx = NewTypeLitBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ApiParserParserRULE_typeLitBody)

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

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.TypeStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.TypeAlias()
		}

	}

	return localctx
}

// ITypeBlockBodyContext is an interface to support dynamic dispatch.
type ITypeBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBlockBodyContext differentiates from other interfaces.
	IsTypeBlockBodyContext()
}

type TypeBlockBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBlockBodyContext() *TypeBlockBodyContext {
	var p = new(TypeBlockBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockBody
	return p
}

func (*TypeBlockBodyContext) IsTypeBlockBodyContext() {}

func NewTypeBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockBodyContext {
	var p = new(TypeBlockBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockBody

	return p
}

func (s *TypeBlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockBodyContext) TypeBlockStruct() ITypeBlockStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBlockStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBlockStructContext)
}

func (s *TypeBlockBodyContext) TypeBlockAlias() ITypeBlockAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBlockAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBlockAliasContext)
}

func (s *TypeBlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockBody() (localctx ITypeBlockBodyContext) {
	localctx = NewTypeBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ApiParserParserRULE_typeBlockBody)

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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.TypeBlockStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.TypeBlockAlias()
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
	p.EnterRule(localctx, 28, ApiParserParserRULE_typeStruct)
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
	checkKeyword(p)
	{
		p.SetState(160)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeStructContext).structName = _m
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(161)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeStructContext).structToken = _m
		}

	}
	{
		p.SetState(164)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeStructContext).lbrace = _m
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)

			var _x = p.CommentSpec()

			localctx.(*TypeStructContext).comment = _x
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(168)
				p.Field()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	{
		p.SetState(173)

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
	p.EnterRule(localctx, 30, ApiParserParserRULE_typeAlias)
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
	checkKeyword(p)
	{
		p.SetState(176)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(177)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeAliasContext).assign = _m
		}

	}
	{
		p.SetState(180)
		p.DataType()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(181)

			var _x = p.CommentSpec()

			localctx.(*TypeAliasContext).comment = _x
		}

	}

	return localctx
}

// ITypeBlockStructContext is an interface to support dynamic dispatch.
type ITypeBlockStructContext interface {
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

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsTypeBlockStructContext differentiates from other interfaces.
	IsTypeBlockStructContext()
}

type TypeBlockStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	doc         ICommentSpecContext
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	comment     ICommentSpecContext
	rbrace      antlr.Token
}

func NewEmptyTypeBlockStructContext() *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct
	return p
}

func (*TypeBlockStructContext) IsTypeBlockStructContext() {}

func NewTypeBlockStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct

	return p
}

func (s *TypeBlockStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeBlockStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeBlockStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeBlockStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeBlockStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeBlockStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeBlockStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeBlockStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeBlockStructContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *TypeBlockStructContext) GetComment() ICommentSpecContext { return s.comment }

func (s *TypeBlockStructContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *TypeBlockStructContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *TypeBlockStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeBlockStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeBlockStructContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *TypeBlockStructContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeBlockStructContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *TypeBlockStructContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *TypeBlockStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockStruct() (localctx ITypeBlockStructContext) {
	localctx = NewTypeBlockStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ApiParserParserRULE_typeBlockStruct)
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
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(184)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockStructContext).doc = _x
		}

	}
	checkKeyword(p)
	{
		p.SetState(188)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockStructContext).structName = _m
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(189)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeBlockStructContext).structToken = _m
		}

	}
	{
		p.SetState(192)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeBlockStructContext).lbrace = _m
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(193)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockStructContext).comment = _x
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(196)
				p.Field()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	{
		p.SetState(201)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeBlockStructContext).rbrace = _m
	}

	return localctx
}

// ITypeBlockAliasContext is an interface to support dynamic dispatch.
type ITypeBlockAliasContext interface {
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

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsTypeBlockAliasContext differentiates from other interfaces.
	IsTypeBlockAliasContext()
}

type TypeBlockAliasContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	doc     ICommentSpecContext
	alias   antlr.Token
	assign  antlr.Token
	comment ICommentSpecContext
}

func NewEmptyTypeBlockAliasContext() *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias
	return p
}

func (*TypeBlockAliasContext) IsTypeBlockAliasContext() {}

func NewTypeBlockAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias

	return p
}

func (s *TypeBlockAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeBlockAliasContext) GetAssign() antlr.Token { return s.assign }

func (s *TypeBlockAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeBlockAliasContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *TypeBlockAliasContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *TypeBlockAliasContext) GetComment() ICommentSpecContext { return s.comment }

func (s *TypeBlockAliasContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *TypeBlockAliasContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *TypeBlockAliasContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TypeBlockAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeBlockAliasContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *TypeBlockAliasContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *TypeBlockAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockAlias() (localctx ITypeBlockAliasContext) {
	localctx = NewTypeBlockAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ApiParserParserRULE_typeBlockAlias)
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
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(203)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockAliasContext).doc = _x
		}

	}
	checkKeyword(p)
	{
		p.SetState(207)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockAliasContext).alias = _m
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(208)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeBlockAliasContext).assign = _m
		}

	}
	{
		p.SetState(211)
		p.DataType()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(212)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockAliasContext).comment = _x
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

func (s *FieldContext) NormalField() INormalFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INormalFieldContext)
}

func (s *FieldContext) AnonymousFiled() IAnonymousFiledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnonymousFiledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnonymousFiledContext)
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
	p.EnterRule(localctx, 36, ApiParserParserRULE_field)

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

	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(215)

		if !(isNormal(p)) {
			panic(antlr.NewFailedPredicateException(p, "isNormal(p)", ""))
		}
		{
			p.SetState(216)
			p.NormalField()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.AnonymousFiled()
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
	p.EnterRule(localctx, 38, ApiParserParserRULE_normalField)
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
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(220)

			var _x = p.CommentSpec()

			localctx.(*NormalFieldContext).doc = _x
		}

	}
	checkKeyword(p)
	{
		p.SetState(224)

		var _m = p.Match(ApiParserParserID)

		localctx.(*NormalFieldContext).fieldName = _m
	}
	{
		p.SetState(225)
		p.DataType()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(226)

			var _m = p.Match(ApiParserParserRAW_STRING)

			localctx.(*NormalFieldContext).tag = _m
		}

	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(229)

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
	p.EnterRule(localctx, 40, ApiParserParserRULE_anonymousFiled)
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
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(232)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).doc = _x
		}

	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__5 {
		{
			p.SetState(235)

			var _m = p.Match(ApiParserParserT__5)

			localctx.(*AnonymousFiledContext).star = _m
		}

	}
	{
		p.SetState(238)
		p.Match(ApiParserParserID)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(239)

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
	p.EnterRule(localctx, 42, ApiParserParserRULE_dataType)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		isInterface(p)
		{
			p.SetState(243)
			p.Match(ApiParserParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.MapType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)

			var _m = p.Match(ApiParserParserINTERFACE)

			localctx.(*DataTypeContext).inter = _m
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(247)

			var _m = p.Match(ApiParserParserT__6)

			localctx.(*DataTypeContext).time = _m
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(248)
			p.PointerType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(249)
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
	p.EnterRule(localctx, 44, ApiParserParserRULE_pointerType)

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
		p.SetState(252)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*PointerTypeContext).star = _m
	}
	checkKeyword(p)
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 46, ApiParserParserRULE_mapType)

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
		p.SetState(257)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).mapToken = _m
	}
	{
		p.SetState(258)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*MapTypeContext).lbrack = _m
	}
	checkKey(p)
	{
		p.SetState(260)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).key = _m
	}
	{
		p.SetState(261)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*MapTypeContext).rbrack = _m
	}
	{
		p.SetState(262)

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
	p.EnterRule(localctx, 48, ApiParserParserRULE_arrayType)

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
		p.SetState(264)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*ArrayTypeContext).lbrack = _m
	}
	{
		p.SetState(265)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*ArrayTypeContext).rbrack = _m
	}
	{
		p.SetState(266)
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
	p.EnterRule(localctx, 50, ApiParserParserRULE_kvLit)
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
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(268)

			var _x = p.CommentSpec()

			localctx.(*KvLitContext).doc = _x
		}

	}
	{
		p.SetState(271)

		var _m = p.Match(ApiParserParserID)

		localctx.(*KvLitContext).key = _m
	}
	checkKeyValue(p)
	{
		p.SetState(273)

		var _m = p.Match(ApiParserParserLINE_VALUE)

		localctx.(*KvLitContext).value = _m
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(274)

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
	p.EnterRule(localctx, 52, ApiParserParserRULE_commentSpec)
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
		p.SetState(277)
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
	case 18:
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
		return isNormal(p)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
