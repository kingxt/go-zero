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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 303,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 5, 2, 74, 10, 2, 3, 2, 7, 2, 77,
	10, 2, 12, 2, 14, 2, 80, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 88, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 96, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 104, 10, 7, 12, 7, 14, 7, 107, 11,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 115, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 123, 10, 11, 12, 11, 14, 11, 126, 11,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 132, 10, 12, 3, 13, 3, 13, 5, 13,
	136, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 142, 10, 14, 3, 14, 3,
	14, 7, 14, 146, 10, 14, 12, 14, 14, 14, 149, 11, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 5, 15, 155, 10, 15, 3, 16, 3, 16, 5, 16, 159, 10, 16, 3, 16, 5,
	16, 162, 10, 16, 3, 17, 5, 17, 165, 10, 17, 3, 17, 3, 17, 7, 17, 169, 10,
	17, 12, 17, 14, 17, 172, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 5, 18, 180, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 21, 7, 21, 193, 10, 21, 12, 21, 14, 21, 196, 11,
	21, 3, 21, 3, 21, 3, 22, 5, 22, 201, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 7, 23, 208, 10, 23, 12, 23, 14, 23, 211, 11, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 5, 24, 218, 10, 24, 3, 25, 3, 25, 3, 25, 5, 25, 223,
	10, 25, 6, 25, 225, 10, 25, 13, 25, 14, 25, 226, 3, 26, 3, 26, 3, 26, 3,
	26, 7, 26, 233, 10, 26, 12, 26, 14, 26, 236, 11, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 5, 27, 243, 10, 27, 3, 28, 5, 28, 246, 10, 28, 3, 28, 3,
	28, 5, 28, 250, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 256, 10, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 263, 10, 30, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 5, 33, 276,
	10, 33, 3, 33, 5, 33, 279, 10, 33, 3, 34, 3, 34, 5, 34, 283, 10, 34, 3,
	34, 3, 34, 3, 34, 5, 34, 288, 10, 34, 6, 34, 290, 10, 34, 13, 34, 14, 34,
	291, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 2, 4, 4, 2, 12, 12, 36, 36, 3, 2, 22, 24, 2, 305, 2, 73, 3, 2, 2, 2,
	4, 87, 3, 2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 95, 3, 2, 2, 2, 10, 97, 3, 2,
	2, 2, 12, 100, 3, 2, 2, 2, 14, 110, 3, 2, 2, 2, 16, 114, 3, 2, 2, 2, 18,
	116, 3, 2, 2, 2, 20, 119, 3, 2, 2, 2, 22, 131, 3, 2, 2, 2, 24, 133, 3,
	2, 2, 2, 26, 139, 3, 2, 2, 2, 28, 152, 3, 2, 2, 2, 30, 158, 3, 2, 2, 2,
	32, 164, 3, 2, 2, 2, 34, 179, 3, 2, 2, 2, 36, 181, 3, 2, 2, 2, 38, 187,
	3, 2, 2, 2, 40, 194, 3, 2, 2, 2, 42, 200, 3, 2, 2, 2, 44, 204, 3, 2, 2,
	2, 46, 214, 3, 2, 2, 2, 48, 224, 3, 2, 2, 2, 50, 228, 3, 2, 2, 2, 52, 239,
	3, 2, 2, 2, 54, 245, 3, 2, 2, 2, 56, 255, 3, 2, 2, 2, 58, 257, 3, 2, 2,
	2, 60, 266, 3, 2, 2, 2, 62, 269, 3, 2, 2, 2, 64, 272, 3, 2, 2, 2, 66, 289,
	3, 2, 2, 2, 68, 293, 3, 2, 2, 2, 70, 297, 3, 2, 2, 2, 72, 74, 5, 6, 4,
	2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 78, 3, 2, 2, 2, 75, 77,
	5, 4, 3, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2,
	78, 79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 82, 7,
	2, 2, 3, 82, 3, 3, 2, 2, 2, 83, 88, 5, 8, 5, 2, 84, 88, 5, 14, 8, 2, 85,
	88, 5, 16, 9, 2, 86, 88, 5, 42, 22, 2, 87, 83, 3, 2, 2, 2, 87, 84, 3, 2,
	2, 2, 87, 85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 5, 3, 2, 2, 2, 89, 90,
	7, 36, 2, 2, 90, 91, 7, 24, 2, 2, 91, 92, 7, 31, 2, 2, 92, 7, 3, 2, 2,
	2, 93, 96, 5, 10, 6, 2, 94, 96, 5, 12, 7, 2, 95, 93, 3, 2, 2, 2, 95, 94,
	3, 2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 98, 7, 8, 2, 2, 98, 99, 7, 32, 2, 2,
	99, 11, 3, 2, 2, 2, 100, 101, 7, 8, 2, 2, 101, 105, 7, 13, 2, 2, 102, 104,
	7, 32, 2, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2,
	2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2,
	108, 109, 7, 14, 2, 2, 109, 13, 3, 2, 2, 2, 110, 111, 7, 35, 2, 2, 111,
	15, 3, 2, 2, 2, 112, 115, 5, 18, 10, 2, 113, 115, 5, 20, 11, 2, 114, 112,
	3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117, 7, 7,
	2, 2, 117, 118, 5, 22, 12, 2, 118, 19, 3, 2, 2, 2, 119, 120, 7, 7, 2, 2,
	120, 124, 7, 13, 2, 2, 121, 123, 5, 22, 12, 2, 122, 121, 3, 2, 2, 2, 123,
	126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127,
	3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7, 14, 2, 2, 128, 21, 3, 2,
	2, 2, 129, 132, 5, 24, 13, 2, 130, 132, 5, 26, 14, 2, 131, 129, 3, 2, 2,
	2, 131, 130, 3, 2, 2, 2, 132, 23, 3, 2, 2, 2, 133, 135, 7, 36, 2, 2, 134,
	136, 7, 24, 2, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137,
	3, 2, 2, 2, 137, 138, 5, 34, 18, 2, 138, 25, 3, 2, 2, 2, 139, 141, 7, 36,
	2, 2, 140, 142, 7, 11, 2, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 147, 7, 15, 2, 2, 144, 146, 5, 28, 15, 2, 145,
	144, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148,
	3, 2, 2, 2, 148, 150, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151, 7, 16,
	2, 2, 151, 27, 3, 2, 2, 2, 152, 154, 7, 36, 2, 2, 153, 155, 5, 30, 16,
	2, 154, 153, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 29, 3, 2, 2, 2, 156,
	159, 5, 34, 18, 2, 157, 159, 5, 32, 17, 2, 158, 156, 3, 2, 2, 2, 158, 157,
	3, 2, 2, 2, 159, 161, 3, 2, 2, 2, 160, 162, 7, 34, 2, 2, 161, 160, 3, 2,
	2, 2, 161, 162, 3, 2, 2, 2, 162, 31, 3, 2, 2, 2, 163, 165, 7, 11, 2, 2,
	164, 163, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	170, 7, 15, 2, 2, 167, 169, 5, 28, 15, 2, 168, 167, 3, 2, 2, 2, 169, 172,
	3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2,
	2, 2, 172, 170, 3, 2, 2, 2, 173, 174, 7, 16, 2, 2, 174, 33, 3, 2, 2, 2,
	175, 180, 5, 40, 21, 2, 176, 180, 5, 36, 19, 2, 177, 180, 5, 38, 20, 2,
	178, 180, 7, 6, 2, 2, 179, 175, 3, 2, 2, 2, 179, 176, 3, 2, 2, 2, 179,
	177, 3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 35, 3, 2, 2, 2, 181, 182, 7,
	10, 2, 2, 182, 183, 7, 17, 2, 2, 183, 184, 7, 12, 2, 2, 184, 185, 7, 18,
	2, 2, 185, 186, 5, 34, 18, 2, 186, 37, 3, 2, 2, 2, 187, 188, 7, 17, 2,
	2, 188, 189, 7, 18, 2, 2, 189, 190, 5, 34, 18, 2, 190, 39, 3, 2, 2, 2,
	191, 193, 7, 27, 2, 2, 192, 191, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2, 194,
	192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196, 194,
	3, 2, 2, 2, 197, 198, 9, 2, 2, 2, 198, 41, 3, 2, 2, 2, 199, 201, 5, 44,
	23, 2, 200, 199, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2,
	202, 203, 5, 50, 26, 2, 203, 43, 3, 2, 2, 2, 204, 205, 7, 3, 2, 2, 205,
	209, 7, 13, 2, 2, 206, 208, 5, 46, 24, 2, 207, 206, 3, 2, 2, 2, 208, 211,
	3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 212, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 212, 213, 7, 14, 2, 2, 213, 45, 3, 2, 2, 2,
	214, 215, 7, 36, 2, 2, 215, 217, 7, 26, 2, 2, 216, 218, 5, 48, 25, 2, 217,
	216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 47, 3, 2, 2, 2, 219, 222, 7,
	36, 2, 2, 220, 221, 7, 21, 2, 2, 221, 223, 7, 36, 2, 2, 222, 220, 3, 2,
	2, 2, 222, 223, 3, 2, 2, 2, 223, 225, 3, 2, 2, 2, 224, 219, 3, 2, 2, 2,
	225, 226, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227,
	49, 3, 2, 2, 2, 228, 229, 7, 36, 2, 2, 229, 230, 5, 52, 27, 2, 230, 234,
	7, 15, 2, 2, 231, 233, 5, 54, 28, 2, 232, 231, 3, 2, 2, 2, 233, 236, 3,
	2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 237, 3, 2, 2,
	2, 236, 234, 3, 2, 2, 2, 237, 238, 7, 16, 2, 2, 238, 51, 3, 2, 2, 2, 239,
	242, 7, 36, 2, 2, 240, 241, 7, 25, 2, 2, 241, 243, 7, 36, 2, 2, 242, 240,
	3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 53, 3, 2, 2, 2, 244, 246, 5, 56,
	29, 2, 245, 244, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2,
	247, 250, 5, 44, 23, 2, 248, 250, 5, 62, 32, 2, 249, 247, 3, 2, 2, 2, 249,
	248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 5, 64, 33, 2, 252, 55,
	3, 2, 2, 2, 253, 256, 5, 58, 30, 2, 254, 256, 5, 60, 31, 2, 255, 253, 3,
	2, 2, 2, 255, 254, 3, 2, 2, 2, 256, 57, 3, 2, 2, 2, 257, 258, 7, 4, 2,
	2, 258, 259, 7, 13, 2, 2, 259, 260, 7, 36, 2, 2, 260, 262, 7, 26, 2, 2,
	261, 263, 7, 33, 2, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263,
	264, 3, 2, 2, 2, 264, 265, 7, 14, 2, 2, 265, 59, 3, 2, 2, 2, 266, 267,
	7, 4, 2, 2, 267, 268, 7, 33, 2, 2, 268, 61, 3, 2, 2, 2, 269, 270, 7, 5,
	2, 2, 270, 271, 7, 36, 2, 2, 271, 63, 3, 2, 2, 2, 272, 273, 7, 36, 2, 2,
	273, 275, 5, 66, 34, 2, 274, 276, 5, 68, 35, 2, 275, 274, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 278, 3, 2, 2, 2, 277, 279, 5, 70, 36, 2, 278, 277,
	3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 65, 3, 2, 2, 2, 280, 282, 7, 21,
	2, 2, 281, 283, 7, 26, 2, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2,
	283, 284, 3, 2, 2, 2, 284, 287, 7, 36, 2, 2, 285, 286, 9, 3, 2, 2, 286,
	288, 7, 36, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 290,
	3, 2, 2, 2, 289, 280, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 289, 3, 2,
	2, 2, 291, 292, 3, 2, 2, 2, 292, 67, 3, 2, 2, 2, 293, 294, 7, 13, 2, 2,
	294, 295, 7, 36, 2, 2, 295, 296, 7, 14, 2, 2, 296, 69, 3, 2, 2, 2, 297,
	298, 7, 36, 2, 2, 298, 299, 7, 13, 2, 2, 299, 300, 7, 36, 2, 2, 300, 301,
	7, 14, 2, 2, 301, 71, 3, 2, 2, 2, 36, 73, 78, 87, 95, 105, 114, 124, 131,
	135, 141, 147, 154, 158, 161, 164, 170, 179, 194, 200, 209, 217, 222, 226,
	234, 242, 245, 249, 255, 262, 275, 278, 282, 287, 291,
}
var literalNames = []string{
	"", "'@server'", "'@doc'", "'@handler'", "'interface{}'", "'type'", "'import'",
	"'info'", "'map'", "'struct'", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"','", "'.'", "'/'", "'?'", "'&'", "'='", "'-'", "':'", "'*'",
}
var symbolicNames = []string{
	"", "ATSERVER", "ATDOC", "ATHANDLER", "INTERFACE", "TYPE", "IMPORT", "INFO",
	"MAP", "STRUCT", "GOTYPE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
	"RBRACK", "COMMA", "DOT", "SLASH", "QUESTION", "BITAND", "ASSIGN", "SUB",
	"COLON", "STAR", "WS", "COMMENT", "LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH",
	"STRING_LIT", "RAW_STRING", "INFO_BLOCK", "ID",
}

var ruleNames = []string{
	"api", "body", "syntaxLit", "importSpec", "importLit", "importLitGroup",
	"infoBlock", "typeBlock", "typeLit", "typeGroup", "typeSpec", "typeAlias",
	"typeStruct", "typeField", "filed", "innerStruct", "dataType", "mapType",
	"arrayType", "pointer", "serviceBlock", "serverMeta", "annotation", "annotationKeyValue",
	"serviceBody", "serviceName", "serviceRoute", "routeDoc", "doc", "lineDoc",
	"routeHandler", "routePath", "path", "request", "reply",
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
	ApiParserEOF            = antlr.TokenEOF
	ApiParserATSERVER       = 1
	ApiParserATDOC          = 2
	ApiParserATHANDLER      = 3
	ApiParserINTERFACE      = 4
	ApiParserTYPE           = 5
	ApiParserIMPORT         = 6
	ApiParserINFO           = 7
	ApiParserMAP            = 8
	ApiParserSTRUCT         = 9
	ApiParserGOTYPE         = 10
	ApiParserLPAREN         = 11
	ApiParserRPAREN         = 12
	ApiParserLBRACE         = 13
	ApiParserRBRACE         = 14
	ApiParserLBRACK         = 15
	ApiParserRBRACK         = 16
	ApiParserCOMMA          = 17
	ApiParserDOT            = 18
	ApiParserSLASH          = 19
	ApiParserQUESTION       = 20
	ApiParserBITAND         = 21
	ApiParserASSIGN         = 22
	ApiParserSUB            = 23
	ApiParserCOLON          = 24
	ApiParserSTAR           = 25
	ApiParserWS             = 26
	ApiParserCOMMENT        = 27
	ApiParserLINE_COMMENT   = 28
	ApiParserSYNTAX_VERSION = 29
	ApiParserIMPORT_PATH    = 30
	ApiParserSTRING_LIT     = 31
	ApiParserRAW_STRING     = 32
	ApiParserINFO_BLOCK     = 33
	ApiParserID             = 34
)

// ApiParser rules.
const (
	ApiParserRULE_api                = 0
	ApiParserRULE_body               = 1
	ApiParserRULE_syntaxLit          = 2
	ApiParserRULE_importSpec         = 3
	ApiParserRULE_importLit          = 4
	ApiParserRULE_importLitGroup     = 5
	ApiParserRULE_infoBlock          = 6
	ApiParserRULE_typeBlock          = 7
	ApiParserRULE_typeLit            = 8
	ApiParserRULE_typeGroup          = 9
	ApiParserRULE_typeSpec           = 10
	ApiParserRULE_typeAlias          = 11
	ApiParserRULE_typeStruct         = 12
	ApiParserRULE_typeField          = 13
	ApiParserRULE_filed              = 14
	ApiParserRULE_innerStruct        = 15
	ApiParserRULE_dataType           = 16
	ApiParserRULE_mapType            = 17
	ApiParserRULE_arrayType          = 18
	ApiParserRULE_pointer            = 19
	ApiParserRULE_serviceBlock       = 20
	ApiParserRULE_serverMeta         = 21
	ApiParserRULE_annotation         = 22
	ApiParserRULE_annotationKeyValue = 23
	ApiParserRULE_serviceBody        = 24
	ApiParserRULE_serviceName        = 25
	ApiParserRULE_serviceRoute       = 26
	ApiParserRULE_routeDoc           = 27
	ApiParserRULE_doc                = 28
	ApiParserRULE_lineDoc            = 29
	ApiParserRULE_routeHandler       = 30
	ApiParserRULE_routePath          = 31
	ApiParserRULE_path               = 32
	ApiParserRULE_request            = 33
	ApiParserRULE_reply              = 34
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(70)
			p.SyntaxLit()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserATSERVER)|(1<<ApiParserTYPE)|(1<<ApiParserIMPORT))) != 0) || _la == ApiParserINFO_BLOCK || _la == ApiParserID {
		{
			p.SetState(73)
			p.Body()
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.ImportSpec()
		}

	case ApiParserINFO_BLOCK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.InfoBlock()
		}

	case ApiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.TypeBlock()
		}

	case ApiParserATSERVER, ApiParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(84)
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
		p.SetState(87)

		var _m = p.Match(ApiParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
	}
	{
		p.SetState(88)
		p.Match(ApiParserASSIGN)
	}
	{
		p.SetState(89)

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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
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
		p.SetState(95)
		p.Match(ApiParserIMPORT)
	}
	{
		p.SetState(96)

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
		p.SetState(98)
		p.Match(ApiParserIMPORT)
	}
	{
		p.SetState(99)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserIMPORT_PATH {
		{
			p.SetState(100)

			var _m = p.Match(ApiParserIMPORT_PATH)

			localctx.(*ImportLitGroupContext).importPath = _m
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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
		p.SetState(108)
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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
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
		p.SetState(114)
		p.Match(ApiParserTYPE)
	}
	{
		p.SetState(115)
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
		p.SetState(117)
		p.Match(ApiParserTYPE)
	}
	{
		p.SetState(118)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(119)
			p.TypeSpec()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.TypeAlias()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
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
		p.SetState(131)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserASSIGN {
		{
			p.SetState(132)
			p.Match(ApiParserASSIGN)
		}

	}
	{
		p.SetState(135)
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
		p.SetState(137)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeStructContext).name = _m
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRUCT {
		{
			p.SetState(138)
			p.Match(ApiParserSTRUCT)
		}

	}
	{
		p.SetState(141)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(142)
			p.TypeField()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
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
		p.SetState(150)

		var _m = p.Match(ApiParserID)

		localctx.(*TypeFieldContext).name = _m
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserINTERFACE, ApiParserMAP, ApiParserGOTYPE, ApiParserLBRACK, ApiParserSTAR, ApiParserID:
		{
			p.SetState(154)
			p.DataType()
		}

	case ApiParserSTRUCT, ApiParserLBRACE:
		{
			p.SetState(155)
			p.InnerStruct()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserRAW_STRING {
		{
			p.SetState(158)

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
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRUCT {
		{
			p.SetState(161)
			p.Match(ApiParserSTRUCT)
		}

	}
	{
		p.SetState(164)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(165)
			p.TypeField()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserGOTYPE, ApiParserSTAR, ApiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Pointer()
		}

	case ApiParserMAP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.MapType()
		}

	case ApiParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.ArrayType()
		}

	case ApiParserINTERFACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
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
		p.SetState(179)
		p.Match(ApiParserMAP)
	}
	{
		p.SetState(180)
		p.Match(ApiParserLBRACK)
	}
	{
		p.SetState(181)

		var _m = p.Match(ApiParserGOTYPE)

		localctx.(*MapTypeContext).key = _m
	}
	{
		p.SetState(182)
		p.Match(ApiParserRBRACK)
	}
	{
		p.SetState(183)

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
		p.SetState(185)
		p.Match(ApiParserLBRACK)
	}
	{
		p.SetState(186)
		p.Match(ApiParserRBRACK)
	}
	{
		p.SetState(187)

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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserSTAR {
		{
			p.SetState(189)
			p.Match(ApiParserSTAR)
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
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
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserATSERVER {
		{
			p.SetState(197)
			p.ServerMeta()
		}

	}
	{
		p.SetState(200)
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

func (s *ServerMetaContext) ATSERVER() antlr.TerminalNode {
	return s.GetToken(ApiParserATSERVER, 0)
}

func (s *ServerMetaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *ServerMetaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *ServerMetaContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ServerMetaContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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
		p.SetState(202)
		p.Match(ApiParserATSERVER)
	}
	{
		p.SetState(203)
		p.Match(ApiParserLPAREN)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserID {
		{
			p.SetState(204)
			p.Annotation()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(210)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IAnnotationKeyValueContext

	// SetValue sets the value rule contexts.
	SetValue(IAnnotationKeyValueContext)

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  IAnnotationKeyValueContext
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) GetKey() antlr.Token { return s.key }

func (s *AnnotationContext) SetKey(v antlr.Token) { s.key = v }

func (s *AnnotationContext) GetValue() IAnnotationKeyValueContext { return s.value }

func (s *AnnotationContext) SetValue(v IAnnotationKeyValueContext) { s.value = v }

func (s *AnnotationContext) COLON() antlr.TerminalNode {
	return s.GetToken(ApiParserCOLON, 0)
}

func (s *AnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *AnnotationContext) AnnotationKeyValue() IAnnotationKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationKeyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationKeyValueContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ApiParserRULE_annotation)

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
		p.SetState(212)

		var _m = p.Match(ApiParserID)

		localctx.(*AnnotationContext).key = _m
	}
	{
		p.SetState(213)
		p.Match(ApiParserCOLON)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(214)

			var _x = p.AnnotationKeyValue()

			localctx.(*AnnotationContext).value = _x
		}

	}

	return localctx
}

// IAnnotationKeyValueContext is an interface to support dynamic dispatch.
type IAnnotationKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationKeyValueContext differentiates from other interfaces.
	IsAnnotationKeyValueContext()
}

type AnnotationKeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationKeyValueContext() *AnnotationKeyValueContext {
	var p = new(AnnotationKeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_annotationKeyValue
	return p
}

func (*AnnotationKeyValueContext) IsAnnotationKeyValueContext() {}

func NewAnnotationKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationKeyValueContext {
	var p = new(AnnotationKeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_annotationKeyValue

	return p
}

func (s *AnnotationKeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationKeyValueContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserID)
}

func (s *AnnotationKeyValueContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserID, i)
}

func (s *AnnotationKeyValueContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(ApiParserSLASH)
}

func (s *AnnotationKeyValueContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserSLASH, i)
}

func (s *AnnotationKeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationKeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationKeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAnnotationKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) AnnotationKeyValue() (localctx IAnnotationKeyValueContext) {
	localctx = NewAnnotationKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ApiParserRULE_annotationKeyValue)
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
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(217)
				p.Match(ApiParserID)
			}
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ApiParserSLASH {
				{
					p.SetState(218)
					p.Match(ApiParserSLASH)
				}
				{
					p.SetState(219)
					p.Match(ApiParserID)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, ApiParserRULE_serviceBody)
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
		p.SetState(226)

		var _m = p.Match(ApiParserID)

		localctx.(*ServiceBodyContext).serviceToken = _m
	}
	{
		p.SetState(227)
		p.ServiceName()
	}
	{
		p.SetState(228)
		p.Match(ApiParserLBRACE)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserATSERVER)|(1<<ApiParserATDOC)|(1<<ApiParserATHANDLER))) != 0 {
		{
			p.SetState(229)

			var _x = p.ServiceRoute()

			localctx.(*ServiceBodyContext).routes = _x
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 50, ApiParserRULE_serviceName)
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
		p.SetState(237)
		p.Match(ApiParserID)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSUB {
		{
			p.SetState(238)
			p.Match(ApiParserSUB)
		}
		{
			p.SetState(239)
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
	p.EnterRule(localctx, 52, ApiParserRULE_serviceRoute)
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
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserATDOC {
		{
			p.SetState(242)
			p.RouteDoc()
		}

	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserATSERVER:
		{
			p.SetState(245)
			p.ServerMeta()
		}

	case ApiParserATHANDLER:
		{
			p.SetState(246)
			p.RouteHandler()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(249)
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
	p.EnterRule(localctx, 54, ApiParserRULE_routeDoc)

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

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Doc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
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
	p.EnterRule(localctx, 56, ApiParserRULE_doc)
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
		p.Match(ApiParserATDOC)
	}
	{
		p.SetState(256)
		p.Match(ApiParserLPAREN)
	}
	{
		p.SetState(257)

		var _m = p.Match(ApiParserID)

		localctx.(*DocContext).summaryToken = _m
	}
	{
		p.SetState(258)
		p.Match(ApiParserCOLON)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserSTRING_LIT {
		{
			p.SetState(259)
			p.Match(ApiParserSTRING_LIT)
		}

	}
	{
		p.SetState(262)
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
	p.EnterRule(localctx, 58, ApiParserRULE_lineDoc)

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
		p.Match(ApiParserATDOC)
	}
	{
		p.SetState(265)
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
	p.EnterRule(localctx, 60, ApiParserRULE_routeHandler)

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
		p.SetState(267)
		p.Match(ApiParserATHANDLER)
	}
	{
		p.SetState(268)
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

	// SetHttpMethodToken sets the httpMethodToken token.
	SetHttpMethodToken(antlr.Token)

	// IsRoutePathContext differentiates from other interfaces.
	IsRoutePathContext()
}

type RoutePathContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	httpMethodToken antlr.Token
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

func (s *RoutePathContext) SetHttpMethodToken(v antlr.Token) { s.httpMethodToken = v }

func (s *RoutePathContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *RoutePathContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *RoutePathContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RoutePathContext) Reply() IReplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplyContext)
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
	p.EnterRule(localctx, 62, ApiParserRULE_routePath)
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
		p.SetState(270)

		var _m = p.Match(ApiParserID)

		localctx.(*RoutePathContext).httpMethodToken = _m
	}
	{
		p.SetState(271)
		p.Path()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserLPAREN {
		{
			p.SetState(272)
			p.Request()
		}

	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserID {
		{
			p.SetState(275)
			p.Reply()
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

func (s *PathContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(ApiParserSLASH)
}

func (s *PathContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserSLASH, i)
}

func (s *PathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserID)
}

func (s *PathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserID, i)
}

func (s *PathContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ApiParserCOLON)
}

func (s *PathContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserCOLON, i)
}

func (s *PathContext) AllQUESTION() []antlr.TerminalNode {
	return s.GetTokens(ApiParserQUESTION)
}

func (s *PathContext) QUESTION(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserQUESTION, i)
}

func (s *PathContext) AllBITAND() []antlr.TerminalNode {
	return s.GetTokens(ApiParserBITAND)
}

func (s *PathContext) BITAND(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserBITAND, i)
}

func (s *PathContext) AllASSIGN() []antlr.TerminalNode {
	return s.GetTokens(ApiParserASSIGN)
}

func (s *PathContext) ASSIGN(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserASSIGN, i)
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
	p.EnterRule(localctx, 64, ApiParserRULE_path)
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
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserSLASH {
		{
			p.SetState(278)
			p.Match(ApiParserSLASH)
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserCOLON {
			{
				p.SetState(279)
				p.Match(ApiParserCOLON)
			}

		}
		{
			p.SetState(282)
			p.Match(ApiParserID)
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserQUESTION)|(1<<ApiParserBITAND)|(1<<ApiParserASSIGN))) != 0 {
			{
				p.SetState(283)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserQUESTION)|(1<<ApiParserBITAND)|(1<<ApiParserASSIGN))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(284)
				p.Match(ApiParserID)
			}

		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *RequestContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserID, 0)
}

func (s *RequestContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRequest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ApiParserRULE_request)

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
		p.SetState(291)
		p.Match(ApiParserLPAREN)
	}
	{
		p.SetState(292)
		p.Match(ApiParserID)
	}
	{
		p.SetState(293)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}

// IReplyContext is an interface to support dynamic dispatch.
type IReplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// GetObj returns the obj token.
	GetObj() antlr.Token

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// SetObj sets the obj token.
	SetObj(antlr.Token)

	// IsReplyContext differentiates from other interfaces.
	IsReplyContext()
}

type ReplyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	returnToken antlr.Token
	obj         antlr.Token
}

func NewEmptyReplyContext() *ReplyContext {
	var p = new(ReplyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserRULE_reply
	return p
}

func (*ReplyContext) IsReplyContext() {}

func NewReplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplyContext {
	var p = new(ReplyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserRULE_reply

	return p
}

func (s *ReplyContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplyContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *ReplyContext) GetObj() antlr.Token { return s.obj }

func (s *ReplyContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *ReplyContext) SetObj(v antlr.Token) { s.obj = v }

func (s *ReplyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserLPAREN, 0)
}

func (s *ReplyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ApiParserRPAREN, 0)
}

func (s *ReplyContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserID)
}

func (s *ReplyContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserID, i)
}

func (s *ReplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitReply(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParser) Reply() (localctx IReplyContext) {
	localctx = NewReplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ApiParserRULE_reply)

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
		p.SetState(295)

		var _m = p.Match(ApiParserID)

		localctx.(*ReplyContext).returnToken = _m
	}
	{
		p.SetState(296)
		p.Match(ApiParserLPAREN)
	}
	{
		p.SetState(297)

		var _m = p.Match(ApiParserID)

		localctx.(*ReplyContext).obj = _m
	}
	{
		p.SetState(298)
		p.Match(ApiParserRPAREN)
	}

	return localctx
}
