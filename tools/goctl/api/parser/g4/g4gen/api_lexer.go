// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/ApiLexer.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 517,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 199,
	10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 26, 6, 26, 232, 10, 26, 13, 26, 14, 26, 233, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 7, 27, 242, 10, 27, 12, 27, 14, 27, 245, 11, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 256, 10,
	28, 12, 28, 14, 28, 259, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 7, 29, 267, 10, 29, 12, 29, 14, 29, 270, 11, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 5, 30, 276, 10, 30, 3, 30, 3, 30, 3, 30, 7, 30, 281, 10, 30, 12,
	30, 14, 30, 284, 11, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 7, 31, 296, 10, 31, 12, 31, 14, 31, 299, 11, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 306, 10, 32, 12, 32, 14, 32, 309,
	11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 7, 33, 315, 10, 33, 12, 33, 14, 33,
	318, 11, 33, 3, 34, 3, 34, 5, 34, 322, 10, 34, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 5, 35, 330, 10, 35, 3, 35, 5, 35, 333, 10, 35, 3, 35,
	3, 35, 3, 35, 6, 35, 338, 10, 35, 13, 35, 14, 35, 339, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 5, 35, 347, 10, 35, 3, 36, 3, 36, 3, 36, 7, 36, 352,
	10, 36, 12, 36, 14, 36, 355, 11, 36, 3, 36, 5, 36, 358, 10, 36, 3, 37,
	3, 37, 3, 38, 3, 38, 7, 38, 364, 10, 38, 12, 38, 14, 38, 367, 11, 38, 3,
	38, 5, 38, 370, 10, 38, 3, 39, 3, 39, 5, 39, 374, 10, 39, 3, 40, 3, 40,
	3, 40, 3, 40, 5, 40, 380, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55,
	3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 243, 2, 61, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29,
	16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47,
	25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65,
	34, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85,
	2, 87, 2, 89, 2, 91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105,
	2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 3, 2, 19, 5,
	2, 11, 12, 14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 3, 2, 51, 59, 3, 2, 50,
	59, 4, 2, 36, 36, 94, 94, 6, 2, 12, 12, 15, 15, 94, 94, 98, 98, 4, 2, 71,
	71, 103, 103, 4, 2, 45, 45, 47, 47, 10, 2, 36, 36, 41, 41, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 3, 2, 50, 53, 3, 2, 50, 57,
	5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 50, 59, 97, 97, 6, 2, 38, 38, 67,
	92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321, 3,
	2, 56322, 57345, 2, 533, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2,
	2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2,
	2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3,
	2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 121, 3, 2, 2, 2, 5, 129,
	3, 2, 2, 2, 7, 134, 3, 2, 2, 2, 9, 143, 3, 2, 2, 2, 11, 155, 3, 2, 2, 2,
	13, 160, 3, 2, 2, 2, 15, 167, 3, 2, 2, 2, 17, 171, 3, 2, 2, 2, 19, 198,
	3, 2, 2, 2, 21, 200, 3, 2, 2, 2, 23, 202, 3, 2, 2, 2, 25, 204, 3, 2, 2,
	2, 27, 206, 3, 2, 2, 2, 29, 208, 3, 2, 2, 2, 31, 210, 3, 2, 2, 2, 33, 212,
	3, 2, 2, 2, 35, 214, 3, 2, 2, 2, 37, 216, 3, 2, 2, 2, 39, 218, 3, 2, 2,
	2, 41, 220, 3, 2, 2, 2, 43, 222, 3, 2, 2, 2, 45, 224, 3, 2, 2, 2, 47, 226,
	3, 2, 2, 2, 49, 228, 3, 2, 2, 2, 51, 231, 3, 2, 2, 2, 53, 237, 3, 2, 2,
	2, 55, 251, 3, 2, 2, 2, 57, 262, 3, 2, 2, 2, 59, 273, 3, 2, 2, 2, 61, 292,
	3, 2, 2, 2, 63, 302, 3, 2, 2, 2, 65, 312, 3, 2, 2, 2, 67, 319, 3, 2, 2,
	2, 69, 346, 3, 2, 2, 2, 71, 348, 3, 2, 2, 2, 73, 359, 3, 2, 2, 2, 75, 361,
	3, 2, 2, 2, 77, 373, 3, 2, 2, 2, 79, 379, 3, 2, 2, 2, 81, 381, 3, 2, 2,
	2, 83, 386, 3, 2, 2, 2, 85, 392, 3, 2, 2, 2, 87, 399, 3, 2, 2, 2, 89, 406,
	3, 2, 2, 2, 91, 413, 3, 2, 2, 2, 93, 418, 3, 2, 2, 2, 95, 424, 3, 2, 2,
	2, 97, 430, 3, 2, 2, 2, 99, 436, 3, 2, 2, 2, 101, 444, 3, 2, 2, 2, 103,
	452, 3, 2, 2, 2, 105, 462, 3, 2, 2, 2, 107, 473, 3, 2, 2, 2, 109, 480,
	3, 2, 2, 2, 111, 484, 3, 2, 2, 2, 113, 489, 3, 2, 2, 2, 115, 497, 3, 2,
	2, 2, 117, 502, 3, 2, 2, 2, 119, 507, 3, 2, 2, 2, 121, 122, 7, 66, 2, 2,
	122, 123, 7, 117, 2, 2, 123, 124, 7, 103, 2, 2, 124, 125, 7, 116, 2, 2,
	125, 126, 7, 120, 2, 2, 126, 127, 7, 103, 2, 2, 127, 128, 7, 116, 2, 2,
	128, 4, 3, 2, 2, 2, 129, 130, 7, 66, 2, 2, 130, 131, 7, 102, 2, 2, 131,
	132, 7, 113, 2, 2, 132, 133, 7, 101, 2, 2, 133, 6, 3, 2, 2, 2, 134, 135,
	7, 66, 2, 2, 135, 136, 7, 106, 2, 2, 136, 137, 7, 99, 2, 2, 137, 138, 7,
	112, 2, 2, 138, 139, 7, 102, 2, 2, 139, 140, 7, 110, 2, 2, 140, 141, 7,
	103, 2, 2, 141, 142, 7, 116, 2, 2, 142, 8, 3, 2, 2, 2, 143, 144, 7, 107,
	2, 2, 144, 145, 7, 112, 2, 2, 145, 146, 7, 118, 2, 2, 146, 147, 7, 103,
	2, 2, 147, 148, 7, 116, 2, 2, 148, 149, 7, 104, 2, 2, 149, 150, 7, 99,
	2, 2, 150, 151, 7, 101, 2, 2, 151, 152, 7, 103, 2, 2, 152, 153, 7, 125,
	2, 2, 153, 154, 7, 127, 2, 2, 154, 10, 3, 2, 2, 2, 155, 156, 7, 118, 2,
	2, 156, 157, 7, 123, 2, 2, 157, 158, 7, 114, 2, 2, 158, 159, 7, 103, 2,
	2, 159, 12, 3, 2, 2, 2, 160, 161, 7, 107, 2, 2, 161, 162, 7, 111, 2, 2,
	162, 163, 7, 114, 2, 2, 163, 164, 7, 113, 2, 2, 164, 165, 7, 116, 2, 2,
	165, 166, 7, 118, 2, 2, 166, 14, 3, 2, 2, 2, 167, 168, 7, 111, 2, 2, 168,
	169, 7, 99, 2, 2, 169, 170, 7, 114, 2, 2, 170, 16, 3, 2, 2, 2, 171, 172,
	7, 117, 2, 2, 172, 173, 7, 118, 2, 2, 173, 174, 7, 116, 2, 2, 174, 175,
	7, 119, 2, 2, 175, 176, 7, 101, 2, 2, 176, 177, 7, 118, 2, 2, 177, 18,
	3, 2, 2, 2, 178, 199, 5, 81, 41, 2, 179, 199, 5, 83, 42, 2, 180, 199, 5,
	85, 43, 2, 181, 199, 5, 87, 44, 2, 182, 199, 5, 89, 45, 2, 183, 199, 5,
	91, 46, 2, 184, 199, 5, 93, 47, 2, 185, 199, 5, 95, 48, 2, 186, 199, 5,
	97, 49, 2, 187, 199, 5, 99, 50, 2, 188, 199, 5, 101, 51, 2, 189, 199, 5,
	103, 52, 2, 190, 199, 5, 105, 53, 2, 191, 199, 5, 107, 54, 2, 192, 199,
	5, 109, 55, 2, 193, 199, 5, 111, 56, 2, 194, 199, 5, 113, 57, 2, 195, 199,
	5, 115, 58, 2, 196, 199, 5, 117, 59, 2, 197, 199, 5, 119, 60, 2, 198, 178,
	3, 2, 2, 2, 198, 179, 3, 2, 2, 2, 198, 180, 3, 2, 2, 2, 198, 181, 3, 2,
	2, 2, 198, 182, 3, 2, 2, 2, 198, 183, 3, 2, 2, 2, 198, 184, 3, 2, 2, 2,
	198, 185, 3, 2, 2, 2, 198, 186, 3, 2, 2, 2, 198, 187, 3, 2, 2, 2, 198,
	188, 3, 2, 2, 2, 198, 189, 3, 2, 2, 2, 198, 190, 3, 2, 2, 2, 198, 191,
	3, 2, 2, 2, 198, 192, 3, 2, 2, 2, 198, 193, 3, 2, 2, 2, 198, 194, 3, 2,
	2, 2, 198, 195, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 197, 3, 2, 2, 2,
	199, 20, 3, 2, 2, 2, 200, 201, 7, 42, 2, 2, 201, 22, 3, 2, 2, 2, 202, 203,
	7, 43, 2, 2, 203, 24, 3, 2, 2, 2, 204, 205, 7, 125, 2, 2, 205, 26, 3, 2,
	2, 2, 206, 207, 7, 127, 2, 2, 207, 28, 3, 2, 2, 2, 208, 209, 7, 93, 2,
	2, 209, 30, 3, 2, 2, 2, 210, 211, 7, 95, 2, 2, 211, 32, 3, 2, 2, 2, 212,
	213, 7, 46, 2, 2, 213, 34, 3, 2, 2, 2, 214, 215, 7, 48, 2, 2, 215, 36,
	3, 2, 2, 2, 216, 217, 7, 49, 2, 2, 217, 38, 3, 2, 2, 2, 218, 219, 7, 65,
	2, 2, 219, 40, 3, 2, 2, 2, 220, 221, 7, 40, 2, 2, 221, 42, 3, 2, 2, 2,
	222, 223, 7, 63, 2, 2, 223, 44, 3, 2, 2, 2, 224, 225, 7, 47, 2, 2, 225,
	46, 3, 2, 2, 2, 226, 227, 7, 60, 2, 2, 227, 48, 3, 2, 2, 2, 228, 229, 7,
	44, 2, 2, 229, 50, 3, 2, 2, 2, 230, 232, 9, 2, 2, 2, 231, 230, 3, 2, 2,
	2, 232, 233, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 236, 8, 26, 2, 2, 236, 52, 3, 2, 2, 2, 237, 238,
	7, 49, 2, 2, 238, 239, 7, 44, 2, 2, 239, 243, 3, 2, 2, 2, 240, 242, 11,
	2, 2, 2, 241, 240, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 244, 3, 2, 2,
	2, 243, 241, 3, 2, 2, 2, 244, 246, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246,
	247, 7, 44, 2, 2, 247, 248, 7, 49, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250,
	8, 27, 2, 2, 250, 54, 3, 2, 2, 2, 251, 252, 7, 49, 2, 2, 252, 253, 7, 49,
	2, 2, 253, 257, 3, 2, 2, 2, 254, 256, 10, 3, 2, 2, 255, 254, 3, 2, 2, 2,
	256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258,
	260, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 261, 8, 28, 2, 2, 261, 56,
	3, 2, 2, 2, 262, 263, 7, 36, 2, 2, 263, 264, 7, 120, 2, 2, 264, 268, 9,
	4, 2, 2, 265, 267, 9, 5, 2, 2, 266, 265, 3, 2, 2, 2, 267, 270, 3, 2, 2,
	2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2, 270,
	268, 3, 2, 2, 2, 271, 272, 7, 36, 2, 2, 272, 58, 3, 2, 2, 2, 273, 275,
	7, 36, 2, 2, 274, 276, 7, 49, 2, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3,
	2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 282, 5, 65, 33, 2, 278, 279, 7, 49,
	2, 2, 279, 281, 5, 65, 33, 2, 280, 278, 3, 2, 2, 2, 281, 284, 3, 2, 2,
	2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 285, 3, 2, 2, 2, 284,
	282, 3, 2, 2, 2, 285, 286, 7, 48, 2, 2, 286, 287, 7, 99, 2, 2, 287, 288,
	7, 114, 2, 2, 288, 289, 7, 107, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 7,
	36, 2, 2, 291, 60, 3, 2, 2, 2, 292, 297, 7, 36, 2, 2, 293, 296, 10, 6,
	2, 2, 294, 296, 5, 69, 35, 2, 295, 293, 3, 2, 2, 2, 295, 294, 3, 2, 2,
	2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298,
	300, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 36, 2, 2, 301, 62,
	3, 2, 2, 2, 302, 307, 7, 98, 2, 2, 303, 306, 10, 7, 2, 2, 304, 306, 5,
	69, 35, 2, 305, 303, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2, 306, 309, 3, 2,
	2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 310, 3, 2, 2, 2,
	309, 307, 3, 2, 2, 2, 310, 311, 7, 98, 2, 2, 311, 64, 3, 2, 2, 2, 312,
	316, 5, 79, 40, 2, 313, 315, 5, 77, 39, 2, 314, 313, 3, 2, 2, 2, 315, 318,
	3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 66, 3, 2,
	2, 2, 318, 316, 3, 2, 2, 2, 319, 321, 9, 8, 2, 2, 320, 322, 9, 9, 2, 2,
	321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323,
	324, 5, 75, 38, 2, 324, 68, 3, 2, 2, 2, 325, 326, 7, 94, 2, 2, 326, 347,
	9, 10, 2, 2, 327, 332, 7, 94, 2, 2, 328, 330, 9, 11, 2, 2, 329, 328, 3,
	2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 333, 9, 12, 2,
	2, 332, 329, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334,
	347, 9, 12, 2, 2, 335, 337, 7, 94, 2, 2, 336, 338, 7, 119, 2, 2, 337, 336,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2,
	2, 2, 340, 341, 3, 2, 2, 2, 341, 342, 5, 73, 37, 2, 342, 343, 5, 73, 37,
	2, 343, 344, 5, 73, 37, 2, 344, 345, 5, 73, 37, 2, 345, 347, 3, 2, 2, 2,
	346, 325, 3, 2, 2, 2, 346, 327, 3, 2, 2, 2, 346, 335, 3, 2, 2, 2, 347,
	70, 3, 2, 2, 2, 348, 357, 5, 73, 37, 2, 349, 352, 5, 73, 37, 2, 350, 352,
	7, 97, 2, 2, 351, 349, 3, 2, 2, 2, 351, 350, 3, 2, 2, 2, 352, 355, 3, 2,
	2, 2, 353, 351, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 3, 2, 2, 2,
	355, 353, 3, 2, 2, 2, 356, 358, 5, 73, 37, 2, 357, 353, 3, 2, 2, 2, 357,
	358, 3, 2, 2, 2, 358, 72, 3, 2, 2, 2, 359, 360, 9, 13, 2, 2, 360, 74, 3,
	2, 2, 2, 361, 369, 9, 5, 2, 2, 362, 364, 9, 14, 2, 2, 363, 362, 3, 2, 2,
	2, 364, 367, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366,
	368, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 368, 370, 9, 5, 2, 2, 369, 365,
	3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 76, 3, 2, 2, 2, 371, 374, 5, 79,
	40, 2, 372, 374, 9, 5, 2, 2, 373, 371, 3, 2, 2, 2, 373, 372, 3, 2, 2, 2,
	374, 78, 3, 2, 2, 2, 375, 380, 9, 15, 2, 2, 376, 380, 10, 16, 2, 2, 377,
	378, 9, 17, 2, 2, 378, 380, 9, 18, 2, 2, 379, 375, 3, 2, 2, 2, 379, 376,
	3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 380, 80, 3, 2, 2, 2, 381, 382, 7, 100,
	2, 2, 382, 383, 7, 113, 2, 2, 383, 384, 7, 113, 2, 2, 384, 385, 7, 110,
	2, 2, 385, 82, 3, 2, 2, 2, 386, 387, 7, 119, 2, 2, 387, 388, 7, 107, 2,
	2, 388, 389, 7, 112, 2, 2, 389, 390, 7, 118, 2, 2, 390, 391, 7, 58, 2,
	2, 391, 84, 3, 2, 2, 2, 392, 393, 7, 119, 2, 2, 393, 394, 7, 107, 2, 2,
	394, 395, 7, 112, 2, 2, 395, 396, 7, 118, 2, 2, 396, 397, 7, 51, 2, 2,
	397, 398, 7, 56, 2, 2, 398, 86, 3, 2, 2, 2, 399, 400, 7, 119, 2, 2, 400,
	401, 7, 107, 2, 2, 401, 402, 7, 112, 2, 2, 402, 403, 7, 118, 2, 2, 403,
	404, 7, 53, 2, 2, 404, 405, 7, 52, 2, 2, 405, 88, 3, 2, 2, 2, 406, 407,
	7, 119, 2, 2, 407, 408, 7, 107, 2, 2, 408, 409, 7, 112, 2, 2, 409, 410,
	7, 118, 2, 2, 410, 411, 7, 56, 2, 2, 411, 412, 7, 54, 2, 2, 412, 90, 3,
	2, 2, 2, 413, 414, 7, 107, 2, 2, 414, 415, 7, 112, 2, 2, 415, 416, 7, 118,
	2, 2, 416, 417, 7, 58, 2, 2, 417, 92, 3, 2, 2, 2, 418, 419, 7, 107, 2,
	2, 419, 420, 7, 112, 2, 2, 420, 421, 7, 118, 2, 2, 421, 422, 7, 51, 2,
	2, 422, 423, 7, 56, 2, 2, 423, 94, 3, 2, 2, 2, 424, 425, 7, 107, 2, 2,
	425, 426, 7, 112, 2, 2, 426, 427, 7, 118, 2, 2, 427, 428, 7, 53, 2, 2,
	428, 429, 7, 52, 2, 2, 429, 96, 3, 2, 2, 2, 430, 431, 7, 107, 2, 2, 431,
	432, 7, 112, 2, 2, 432, 433, 7, 118, 2, 2, 433, 434, 7, 56, 2, 2, 434,
	435, 7, 54, 2, 2, 435, 98, 3, 2, 2, 2, 436, 437, 7, 104, 2, 2, 437, 438,
	7, 110, 2, 2, 438, 439, 7, 113, 2, 2, 439, 440, 7, 99, 2, 2, 440, 441,
	7, 118, 2, 2, 441, 442, 7, 53, 2, 2, 442, 443, 7, 52, 2, 2, 443, 100, 3,
	2, 2, 2, 444, 445, 7, 104, 2, 2, 445, 446, 7, 110, 2, 2, 446, 447, 7, 113,
	2, 2, 447, 448, 7, 99, 2, 2, 448, 449, 7, 118, 2, 2, 449, 450, 7, 56, 2,
	2, 450, 451, 7, 54, 2, 2, 451, 102, 3, 2, 2, 2, 452, 453, 7, 101, 2, 2,
	453, 454, 7, 113, 2, 2, 454, 455, 7, 111, 2, 2, 455, 456, 7, 114, 2, 2,
	456, 457, 7, 110, 2, 2, 457, 458, 7, 103, 2, 2, 458, 459, 7, 122, 2, 2,
	459, 460, 7, 56, 2, 2, 460, 461, 7, 54, 2, 2, 461, 104, 3, 2, 2, 2, 462,
	463, 7, 101, 2, 2, 463, 464, 7, 113, 2, 2, 464, 465, 7, 111, 2, 2, 465,
	466, 7, 114, 2, 2, 466, 467, 7, 110, 2, 2, 467, 468, 7, 103, 2, 2, 468,
	469, 7, 122, 2, 2, 469, 470, 7, 51, 2, 2, 470, 471, 7, 52, 2, 2, 471, 472,
	7, 58, 2, 2, 472, 106, 3, 2, 2, 2, 473, 474, 7, 117, 2, 2, 474, 475, 7,
	118, 2, 2, 475, 476, 7, 116, 2, 2, 476, 477, 7, 107, 2, 2, 477, 478, 7,
	112, 2, 2, 478, 479, 7, 105, 2, 2, 479, 108, 3, 2, 2, 2, 480, 481, 7, 107,
	2, 2, 481, 482, 7, 112, 2, 2, 482, 483, 7, 118, 2, 2, 483, 110, 3, 2, 2,
	2, 484, 485, 7, 119, 2, 2, 485, 486, 7, 107, 2, 2, 486, 487, 7, 112, 2,
	2, 487, 488, 7, 118, 2, 2, 488, 112, 3, 2, 2, 2, 489, 490, 7, 119, 2, 2,
	490, 491, 7, 107, 2, 2, 491, 492, 7, 112, 2, 2, 492, 493, 7, 118, 2, 2,
	493, 494, 7, 114, 2, 2, 494, 495, 7, 118, 2, 2, 495, 496, 7, 116, 2, 2,
	496, 114, 3, 2, 2, 2, 497, 498, 7, 100, 2, 2, 498, 499, 7, 123, 2, 2, 499,
	500, 7, 118, 2, 2, 500, 501, 7, 103, 2, 2, 501, 116, 3, 2, 2, 2, 502, 503,
	7, 116, 2, 2, 503, 504, 7, 119, 2, 2, 504, 505, 7, 112, 2, 2, 505, 506,
	7, 103, 2, 2, 506, 118, 3, 2, 2, 2, 507, 508, 7, 118, 2, 2, 508, 509, 7,
	107, 2, 2, 509, 510, 7, 111, 2, 2, 510, 511, 7, 103, 2, 2, 511, 512, 7,
	48, 2, 2, 512, 513, 7, 86, 2, 2, 513, 514, 7, 107, 2, 2, 514, 515, 7, 111,
	2, 2, 515, 516, 7, 103, 2, 2, 516, 120, 3, 2, 2, 2, 27, 2, 198, 233, 243,
	257, 268, 275, 282, 295, 297, 305, 307, 316, 321, 329, 332, 339, 346, 351,
	353, 357, 365, 369, 373, 379, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@server'", "'@doc'", "'@handler'", "'interface{}'", "'type'", "'import'",
	"'map'", "'struct'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
	"'.'", "'/'", "'?'", "'&'", "'='", "'-'", "':'", "'*'",
}

var lexerSymbolicNames = []string{
	"", "ATSERVER", "ATDOC", "ATHANDLER", "INTERFACE", "TYPE", "IMPORT", "MAP",
	"STRUCT", "GOTYPE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"COMMA", "DOT", "SLASH", "QUESTION", "BITAND", "ASSIGN", "SUB", "COLON",
	"STAR", "WS", "COMMENT", "LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH",
	"STRING_LIT", "RAW_STRING", "ID",
}

var lexerRuleNames = []string{
	"ATSERVER", "ATDOC", "ATHANDLER", "INTERFACE", "TYPE", "IMPORT", "MAP",
	"STRUCT", "GOTYPE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"COMMA", "DOT", "SLASH", "QUESTION", "BITAND", "ASSIGN", "SUB", "COLON",
	"STAR", "WS", "COMMENT", "LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH",
	"STRING_LIT", "RAW_STRING", "ID", "ExponentPart", "EscapeSequence", "HexDigits",
	"HexDigit", "Digits", "LetterOrDigit", "Letter", "BOOL", "UINT8", "UINT16",
	"UINT32", "UINT64", "INT8", "INT16", "INT32", "INT64", "FLOAT32", "FLOAT64",
	"COMPLEX64", "COMPLEX128", "STRING", "INT", "UINT", "UINTPTR", "BYTE",
	"RUNE", "TIME",
}

type ApiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewApiLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ApiLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewApiLexer(input antlr.CharStream) *ApiLexer {
	l := new(ApiLexer)
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
	l.GrammarFileName = "ApiLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ApiLexer tokens.
const (
	ApiLexerATSERVER       = 1
	ApiLexerATDOC          = 2
	ApiLexerATHANDLER      = 3
	ApiLexerINTERFACE      = 4
	ApiLexerTYPE           = 5
	ApiLexerIMPORT         = 6
	ApiLexerMAP            = 7
	ApiLexerSTRUCT         = 8
	ApiLexerGOTYPE         = 9
	ApiLexerLPAREN         = 10
	ApiLexerRPAREN         = 11
	ApiLexerLBRACE         = 12
	ApiLexerRBRACE         = 13
	ApiLexerLBRACK         = 14
	ApiLexerRBRACK         = 15
	ApiLexerCOMMA          = 16
	ApiLexerDOT            = 17
	ApiLexerSLASH          = 18
	ApiLexerQUESTION       = 19
	ApiLexerBITAND         = 20
	ApiLexerASSIGN         = 21
	ApiLexerSUB            = 22
	ApiLexerCOLON          = 23
	ApiLexerSTAR           = 24
	ApiLexerWS             = 25
	ApiLexerCOMMENT        = 26
	ApiLexerLINE_COMMENT   = 27
	ApiLexerSYNTAX_VERSION = 28
	ApiLexerIMPORT_PATH    = 29
	ApiLexerSTRING_LIT     = 30
	ApiLexerRAW_STRING     = 31
	ApiLexerID             = 32
)
