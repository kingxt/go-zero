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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 637,
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
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 170, 10, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 266, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 31, 6, 31, 299, 10, 31, 13, 31, 14, 31, 300, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 309, 10, 32, 12, 32, 14,
	32, 312, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 7, 33, 323, 10, 33, 12, 33, 14, 33, 326, 11, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 7, 34, 334, 10, 34, 12, 34, 14, 34, 337, 11, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 5, 35, 343, 10, 35, 3, 35, 3, 35, 3, 35, 7,
	35, 348, 10, 35, 12, 35, 14, 35, 351, 11, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 363, 10, 36, 12, 36, 14,
	36, 366, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 7, 37, 373, 10, 37,
	12, 37, 14, 37, 376, 11, 37, 3, 37, 3, 37, 3, 38, 3, 38, 7, 38, 382, 10,
	38, 12, 38, 14, 38, 385, 11, 38, 3, 39, 3, 39, 5, 39, 389, 10, 39, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 397, 10, 40, 3, 40, 5, 40, 400,
	10, 40, 3, 40, 3, 40, 3, 40, 6, 40, 405, 10, 40, 13, 40, 14, 40, 406, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 414, 10, 40, 3, 41, 3, 41, 3, 41,
	7, 41, 419, 10, 41, 12, 41, 14, 41, 422, 11, 41, 3, 41, 5, 41, 425, 10,
	41, 3, 42, 3, 42, 3, 43, 3, 43, 7, 43, 431, 10, 43, 12, 43, 14, 43, 434,
	11, 43, 3, 43, 5, 43, 437, 10, 43, 3, 44, 3, 44, 5, 44, 441, 10, 44, 3,
	45, 3, 45, 3, 45, 3, 45, 5, 45, 447, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62,
	3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3,
	63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65,
	3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3,
	67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69,
	3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3,
	71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72,
	3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3,
	73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 310, 2, 75, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23,
	45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32,
	63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 2, 79, 2, 81,
	2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 101,
	2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119,
	2, 121, 2, 123, 2, 125, 2, 127, 2, 129, 2, 131, 2, 133, 2, 135, 2, 137,
	2, 139, 2, 141, 2, 143, 2, 145, 2, 147, 2, 3, 2, 19, 5, 2, 11, 12, 14,
	15, 34, 34, 4, 2, 12, 12, 15, 15, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 36,
	36, 94, 94, 6, 2, 12, 12, 15, 15, 94, 94, 98, 98, 4, 2, 71, 71, 103, 103,
	4, 2, 45, 45, 47, 47, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104,
	112, 112, 116, 116, 118, 118, 3, 2, 50, 53, 3, 2, 50, 57, 5, 2, 50, 59,
	67, 72, 99, 104, 4, 2, 50, 59, 97, 97, 6, 2, 38, 38, 67, 92, 97, 97, 99,
	124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321, 3, 2, 56322, 57345,
	2, 652, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63,
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2,
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 3, 169, 3, 2, 2,
	2, 5, 171, 3, 2, 2, 2, 7, 178, 3, 2, 2, 2, 9, 183, 3, 2, 2, 2, 11, 187,
	3, 2, 2, 2, 13, 194, 3, 2, 2, 2, 15, 206, 3, 2, 2, 2, 17, 211, 3, 2, 2,
	2, 19, 219, 3, 2, 2, 2, 21, 227, 3, 2, 2, 2, 23, 234, 3, 2, 2, 2, 25, 242,
	3, 2, 2, 2, 27, 247, 3, 2, 2, 2, 29, 265, 3, 2, 2, 2, 31, 267, 3, 2, 2,
	2, 33, 269, 3, 2, 2, 2, 35, 271, 3, 2, 2, 2, 37, 273, 3, 2, 2, 2, 39, 275,
	3, 2, 2, 2, 41, 277, 3, 2, 2, 2, 43, 279, 3, 2, 2, 2, 45, 281, 3, 2, 2,
	2, 47, 283, 3, 2, 2, 2, 49, 285, 3, 2, 2, 2, 51, 287, 3, 2, 2, 2, 53, 289,
	3, 2, 2, 2, 55, 291, 3, 2, 2, 2, 57, 293, 3, 2, 2, 2, 59, 295, 3, 2, 2,
	2, 61, 298, 3, 2, 2, 2, 63, 304, 3, 2, 2, 2, 65, 318, 3, 2, 2, 2, 67, 329,
	3, 2, 2, 2, 69, 340, 3, 2, 2, 2, 71, 359, 3, 2, 2, 2, 73, 369, 3, 2, 2,
	2, 75, 379, 3, 2, 2, 2, 77, 386, 3, 2, 2, 2, 79, 413, 3, 2, 2, 2, 81, 415,
	3, 2, 2, 2, 83, 426, 3, 2, 2, 2, 85, 428, 3, 2, 2, 2, 87, 440, 3, 2, 2,
	2, 89, 446, 3, 2, 2, 2, 91, 448, 3, 2, 2, 2, 93, 453, 3, 2, 2, 2, 95, 459,
	3, 2, 2, 2, 97, 466, 3, 2, 2, 2, 99, 473, 3, 2, 2, 2, 101, 480, 3, 2, 2,
	2, 103, 485, 3, 2, 2, 2, 105, 491, 3, 2, 2, 2, 107, 497, 3, 2, 2, 2, 109,
	503, 3, 2, 2, 2, 111, 511, 3, 2, 2, 2, 113, 519, 3, 2, 2, 2, 115, 529,
	3, 2, 2, 2, 117, 540, 3, 2, 2, 2, 119, 547, 3, 2, 2, 2, 121, 551, 3, 2,
	2, 2, 123, 556, 3, 2, 2, 2, 125, 564, 3, 2, 2, 2, 127, 569, 3, 2, 2, 2,
	129, 574, 3, 2, 2, 2, 131, 584, 3, 2, 2, 2, 133, 588, 3, 2, 2, 2, 135,
	593, 3, 2, 2, 2, 137, 598, 3, 2, 2, 2, 139, 602, 3, 2, 2, 2, 141, 608,
	3, 2, 2, 2, 143, 615, 3, 2, 2, 2, 145, 623, 3, 2, 2, 2, 147, 631, 3, 2,
	2, 2, 149, 170, 5, 91, 46, 2, 150, 170, 5, 93, 47, 2, 151, 170, 5, 95,
	48, 2, 152, 170, 5, 97, 49, 2, 153, 170, 5, 99, 50, 2, 154, 170, 5, 101,
	51, 2, 155, 170, 5, 103, 52, 2, 156, 170, 5, 105, 53, 2, 157, 170, 5, 107,
	54, 2, 158, 170, 5, 109, 55, 2, 159, 170, 5, 111, 56, 2, 160, 170, 5, 113,
	57, 2, 161, 170, 5, 115, 58, 2, 162, 170, 5, 117, 59, 2, 163, 170, 5, 119,
	60, 2, 164, 170, 5, 121, 61, 2, 165, 170, 5, 123, 62, 2, 166, 170, 5, 125,
	63, 2, 167, 170, 5, 127, 64, 2, 168, 170, 5, 129, 65, 2, 169, 149, 3, 2,
	2, 2, 169, 150, 3, 2, 2, 2, 169, 151, 3, 2, 2, 2, 169, 152, 3, 2, 2, 2,
	169, 153, 3, 2, 2, 2, 169, 154, 3, 2, 2, 2, 169, 155, 3, 2, 2, 2, 169,
	156, 3, 2, 2, 2, 169, 157, 3, 2, 2, 2, 169, 158, 3, 2, 2, 2, 169, 159,
	3, 2, 2, 2, 169, 160, 3, 2, 2, 2, 169, 161, 3, 2, 2, 2, 169, 162, 3, 2,
	2, 2, 169, 163, 3, 2, 2, 2, 169, 164, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2,
	169, 166, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170,
	4, 3, 2, 2, 2, 171, 172, 7, 117, 2, 2, 172, 173, 7, 123, 2, 2, 173, 174,
	7, 112, 2, 2, 174, 175, 7, 118, 2, 2, 175, 176, 7, 99, 2, 2, 176, 177,
	7, 122, 2, 2, 177, 6, 3, 2, 2, 2, 178, 179, 7, 107, 2, 2, 179, 180, 7,
	112, 2, 2, 180, 181, 7, 104, 2, 2, 181, 182, 7, 113, 2, 2, 182, 8, 3, 2,
	2, 2, 183, 184, 7, 111, 2, 2, 184, 185, 7, 99, 2, 2, 185, 186, 7, 114,
	2, 2, 186, 10, 3, 2, 2, 2, 187, 188, 7, 117, 2, 2, 188, 189, 7, 118, 2,
	2, 189, 190, 7, 116, 2, 2, 190, 191, 7, 119, 2, 2, 191, 192, 7, 101, 2,
	2, 192, 193, 7, 118, 2, 2, 193, 12, 3, 2, 2, 2, 194, 195, 7, 107, 2, 2,
	195, 196, 7, 112, 2, 2, 196, 197, 7, 118, 2, 2, 197, 198, 7, 103, 2, 2,
	198, 199, 7, 116, 2, 2, 199, 200, 7, 104, 2, 2, 200, 201, 7, 99, 2, 2,
	201, 202, 7, 101, 2, 2, 202, 203, 7, 103, 2, 2, 203, 204, 7, 125, 2, 2,
	204, 205, 7, 127, 2, 2, 205, 14, 3, 2, 2, 2, 206, 207, 7, 118, 2, 2, 207,
	208, 7, 123, 2, 2, 208, 209, 7, 114, 2, 2, 209, 210, 7, 103, 2, 2, 210,
	16, 3, 2, 2, 2, 211, 212, 7, 117, 2, 2, 212, 213, 7, 103, 2, 2, 213, 214,
	7, 116, 2, 2, 214, 215, 7, 120, 2, 2, 215, 216, 7, 107, 2, 2, 216, 217,
	7, 101, 2, 2, 217, 218, 7, 103, 2, 2, 218, 18, 3, 2, 2, 2, 219, 220, 7,
	116, 2, 2, 220, 221, 7, 103, 2, 2, 221, 222, 7, 118, 2, 2, 222, 223, 7,
	119, 2, 2, 223, 224, 7, 116, 2, 2, 224, 225, 7, 112, 2, 2, 225, 226, 7,
	117, 2, 2, 226, 20, 3, 2, 2, 2, 227, 228, 7, 107, 2, 2, 228, 229, 7, 111,
	2, 2, 229, 230, 7, 114, 2, 2, 230, 231, 7, 113, 2, 2, 231, 232, 7, 116,
	2, 2, 232, 233, 7, 118, 2, 2, 233, 22, 3, 2, 2, 2, 234, 235, 7, 66, 2,
	2, 235, 236, 7, 117, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7, 116, 2,
	2, 238, 239, 7, 120, 2, 2, 239, 240, 7, 103, 2, 2, 240, 241, 7, 116, 2,
	2, 241, 24, 3, 2, 2, 2, 242, 243, 7, 66, 2, 2, 243, 244, 7, 102, 2, 2,
	244, 245, 7, 113, 2, 2, 245, 246, 7, 101, 2, 2, 246, 26, 3, 2, 2, 2, 247,
	248, 7, 66, 2, 2, 248, 249, 7, 106, 2, 2, 249, 250, 7, 99, 2, 2, 250, 251,
	7, 112, 2, 2, 251, 252, 7, 102, 2, 2, 252, 253, 7, 110, 2, 2, 253, 254,
	7, 103, 2, 2, 254, 255, 7, 116, 2, 2, 255, 28, 3, 2, 2, 2, 256, 266, 5,
	131, 66, 2, 257, 266, 5, 133, 67, 2, 258, 266, 5, 135, 68, 2, 259, 266,
	5, 137, 69, 2, 260, 266, 5, 139, 70, 2, 261, 266, 5, 141, 71, 2, 262, 266,
	5, 143, 72, 2, 263, 266, 5, 145, 73, 2, 264, 266, 5, 147, 74, 2, 265, 256,
	3, 2, 2, 2, 265, 257, 3, 2, 2, 2, 265, 258, 3, 2, 2, 2, 265, 259, 3, 2,
	2, 2, 265, 260, 3, 2, 2, 2, 265, 261, 3, 2, 2, 2, 265, 262, 3, 2, 2, 2,
	265, 263, 3, 2, 2, 2, 265, 264, 3, 2, 2, 2, 266, 30, 3, 2, 2, 2, 267, 268,
	7, 42, 2, 2, 268, 32, 3, 2, 2, 2, 269, 270, 7, 43, 2, 2, 270, 34, 3, 2,
	2, 2, 271, 272, 7, 125, 2, 2, 272, 36, 3, 2, 2, 2, 273, 274, 7, 127, 2,
	2, 274, 38, 3, 2, 2, 2, 275, 276, 7, 93, 2, 2, 276, 40, 3, 2, 2, 2, 277,
	278, 7, 95, 2, 2, 278, 42, 3, 2, 2, 2, 279, 280, 7, 46, 2, 2, 280, 44,
	3, 2, 2, 2, 281, 282, 7, 48, 2, 2, 282, 46, 3, 2, 2, 2, 283, 284, 7, 49,
	2, 2, 284, 48, 3, 2, 2, 2, 285, 286, 7, 65, 2, 2, 286, 50, 3, 2, 2, 2,
	287, 288, 7, 40, 2, 2, 288, 52, 3, 2, 2, 2, 289, 290, 7, 63, 2, 2, 290,
	54, 3, 2, 2, 2, 291, 292, 7, 47, 2, 2, 292, 56, 3, 2, 2, 2, 293, 294, 7,
	60, 2, 2, 294, 58, 3, 2, 2, 2, 295, 296, 7, 44, 2, 2, 296, 60, 3, 2, 2,
	2, 297, 299, 9, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300,
	298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303,
	8, 31, 2, 2, 303, 62, 3, 2, 2, 2, 304, 305, 7, 49, 2, 2, 305, 306, 7, 44,
	2, 2, 306, 310, 3, 2, 2, 2, 307, 309, 11, 2, 2, 2, 308, 307, 3, 2, 2, 2,
	309, 312, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 311,
	313, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 313, 314, 7, 44, 2, 2, 314, 315,
	7, 49, 2, 2, 315, 316, 3, 2, 2, 2, 316, 317, 8, 32, 2, 2, 317, 64, 3, 2,
	2, 2, 318, 319, 7, 49, 2, 2, 319, 320, 7, 49, 2, 2, 320, 324, 3, 2, 2,
	2, 321, 323, 10, 3, 2, 2, 322, 321, 3, 2, 2, 2, 323, 326, 3, 2, 2, 2, 324,
	322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 327, 3, 2, 2, 2, 326, 324,
	3, 2, 2, 2, 327, 328, 8, 33, 2, 2, 328, 66, 3, 2, 2, 2, 329, 330, 7, 36,
	2, 2, 330, 331, 7, 120, 2, 2, 331, 335, 9, 4, 2, 2, 332, 334, 9, 5, 2,
	2, 333, 332, 3, 2, 2, 2, 334, 337, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 335,
	336, 3, 2, 2, 2, 336, 338, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 338, 339,
	7, 36, 2, 2, 339, 68, 3, 2, 2, 2, 340, 342, 7, 36, 2, 2, 341, 343, 7, 49,
	2, 2, 342, 341, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2,
	344, 349, 5, 75, 38, 2, 345, 346, 7, 49, 2, 2, 346, 348, 5, 75, 38, 2,
	347, 345, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349,
	350, 3, 2, 2, 2, 350, 352, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352, 353,
	7, 48, 2, 2, 353, 354, 7, 99, 2, 2, 354, 355, 7, 114, 2, 2, 355, 356, 7,
	107, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 7, 36, 2, 2, 358, 70, 3, 2,
	2, 2, 359, 364, 7, 36, 2, 2, 360, 363, 10, 6, 2, 2, 361, 363, 5, 79, 40,
	2, 362, 360, 3, 2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364,
	362, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 367, 3, 2, 2, 2, 366, 364,
	3, 2, 2, 2, 367, 368, 7, 36, 2, 2, 368, 72, 3, 2, 2, 2, 369, 374, 7, 98,
	2, 2, 370, 373, 10, 7, 2, 2, 371, 373, 5, 79, 40, 2, 372, 370, 3, 2, 2,
	2, 372, 371, 3, 2, 2, 2, 373, 376, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374,
	375, 3, 2, 2, 2, 375, 377, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 377, 378,
	7, 98, 2, 2, 378, 74, 3, 2, 2, 2, 379, 383, 5, 89, 45, 2, 380, 382, 5,
	87, 44, 2, 381, 380, 3, 2, 2, 2, 382, 385, 3, 2, 2, 2, 383, 381, 3, 2,
	2, 2, 383, 384, 3, 2, 2, 2, 384, 76, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2,
	386, 388, 9, 8, 2, 2, 387, 389, 9, 9, 2, 2, 388, 387, 3, 2, 2, 2, 388,
	389, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 5, 85, 43, 2, 391, 78,
	3, 2, 2, 2, 392, 393, 7, 94, 2, 2, 393, 414, 9, 10, 2, 2, 394, 399, 7,
	94, 2, 2, 395, 397, 9, 11, 2, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2,
	2, 2, 397, 398, 3, 2, 2, 2, 398, 400, 9, 12, 2, 2, 399, 396, 3, 2, 2, 2,
	399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 414, 9, 12, 2, 2, 402,
	404, 7, 94, 2, 2, 403, 405, 7, 119, 2, 2, 404, 403, 3, 2, 2, 2, 405, 406,
	3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 408, 3, 2,
	2, 2, 408, 409, 5, 83, 42, 2, 409, 410, 5, 83, 42, 2, 410, 411, 5, 83,
	42, 2, 411, 412, 5, 83, 42, 2, 412, 414, 3, 2, 2, 2, 413, 392, 3, 2, 2,
	2, 413, 394, 3, 2, 2, 2, 413, 402, 3, 2, 2, 2, 414, 80, 3, 2, 2, 2, 415,
	424, 5, 83, 42, 2, 416, 419, 5, 83, 42, 2, 417, 419, 7, 97, 2, 2, 418,
	416, 3, 2, 2, 2, 418, 417, 3, 2, 2, 2, 419, 422, 3, 2, 2, 2, 420, 418,
	3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 423, 3, 2, 2, 2, 422, 420, 3, 2,
	2, 2, 423, 425, 5, 83, 42, 2, 424, 420, 3, 2, 2, 2, 424, 425, 3, 2, 2,
	2, 425, 82, 3, 2, 2, 2, 426, 427, 9, 13, 2, 2, 427, 84, 3, 2, 2, 2, 428,
	436, 9, 5, 2, 2, 429, 431, 9, 14, 2, 2, 430, 429, 3, 2, 2, 2, 431, 434,
	3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 435, 3, 2,
	2, 2, 434, 432, 3, 2, 2, 2, 435, 437, 9, 5, 2, 2, 436, 432, 3, 2, 2, 2,
	436, 437, 3, 2, 2, 2, 437, 86, 3, 2, 2, 2, 438, 441, 5, 89, 45, 2, 439,
	441, 9, 5, 2, 2, 440, 438, 3, 2, 2, 2, 440, 439, 3, 2, 2, 2, 441, 88, 3,
	2, 2, 2, 442, 447, 9, 15, 2, 2, 443, 447, 10, 16, 2, 2, 444, 445, 9, 17,
	2, 2, 445, 447, 9, 18, 2, 2, 446, 442, 3, 2, 2, 2, 446, 443, 3, 2, 2, 2,
	446, 444, 3, 2, 2, 2, 447, 90, 3, 2, 2, 2, 448, 449, 7, 100, 2, 2, 449,
	450, 7, 113, 2, 2, 450, 451, 7, 113, 2, 2, 451, 452, 7, 110, 2, 2, 452,
	92, 3, 2, 2, 2, 453, 454, 7, 119, 2, 2, 454, 455, 7, 107, 2, 2, 455, 456,
	7, 112, 2, 2, 456, 457, 7, 118, 2, 2, 457, 458, 7, 58, 2, 2, 458, 94, 3,
	2, 2, 2, 459, 460, 7, 119, 2, 2, 460, 461, 7, 107, 2, 2, 461, 462, 7, 112,
	2, 2, 462, 463, 7, 118, 2, 2, 463, 464, 7, 51, 2, 2, 464, 465, 7, 56, 2,
	2, 465, 96, 3, 2, 2, 2, 466, 467, 7, 119, 2, 2, 467, 468, 7, 107, 2, 2,
	468, 469, 7, 112, 2, 2, 469, 470, 7, 118, 2, 2, 470, 471, 7, 53, 2, 2,
	471, 472, 7, 52, 2, 2, 472, 98, 3, 2, 2, 2, 473, 474, 7, 119, 2, 2, 474,
	475, 7, 107, 2, 2, 475, 476, 7, 112, 2, 2, 476, 477, 7, 118, 2, 2, 477,
	478, 7, 56, 2, 2, 478, 479, 7, 54, 2, 2, 479, 100, 3, 2, 2, 2, 480, 481,
	7, 107, 2, 2, 481, 482, 7, 112, 2, 2, 482, 483, 7, 118, 2, 2, 483, 484,
	7, 58, 2, 2, 484, 102, 3, 2, 2, 2, 485, 486, 7, 107, 2, 2, 486, 487, 7,
	112, 2, 2, 487, 488, 7, 118, 2, 2, 488, 489, 7, 51, 2, 2, 489, 490, 7,
	56, 2, 2, 490, 104, 3, 2, 2, 2, 491, 492, 7, 107, 2, 2, 492, 493, 7, 112,
	2, 2, 493, 494, 7, 118, 2, 2, 494, 495, 7, 53, 2, 2, 495, 496, 7, 52, 2,
	2, 496, 106, 3, 2, 2, 2, 497, 498, 7, 107, 2, 2, 498, 499, 7, 112, 2, 2,
	499, 500, 7, 118, 2, 2, 500, 501, 7, 56, 2, 2, 501, 502, 7, 54, 2, 2, 502,
	108, 3, 2, 2, 2, 503, 504, 7, 104, 2, 2, 504, 505, 7, 110, 2, 2, 505, 506,
	7, 113, 2, 2, 506, 507, 7, 99, 2, 2, 507, 508, 7, 118, 2, 2, 508, 509,
	7, 53, 2, 2, 509, 510, 7, 52, 2, 2, 510, 110, 3, 2, 2, 2, 511, 512, 7,
	104, 2, 2, 512, 513, 7, 110, 2, 2, 513, 514, 7, 113, 2, 2, 514, 515, 7,
	99, 2, 2, 515, 516, 7, 118, 2, 2, 516, 517, 7, 56, 2, 2, 517, 518, 7, 54,
	2, 2, 518, 112, 3, 2, 2, 2, 519, 520, 7, 101, 2, 2, 520, 521, 7, 113, 2,
	2, 521, 522, 7, 111, 2, 2, 522, 523, 7, 114, 2, 2, 523, 524, 7, 110, 2,
	2, 524, 525, 7, 103, 2, 2, 525, 526, 7, 122, 2, 2, 526, 527, 7, 56, 2,
	2, 527, 528, 7, 54, 2, 2, 528, 114, 3, 2, 2, 2, 529, 530, 7, 101, 2, 2,
	530, 531, 7, 113, 2, 2, 531, 532, 7, 111, 2, 2, 532, 533, 7, 114, 2, 2,
	533, 534, 7, 110, 2, 2, 534, 535, 7, 103, 2, 2, 535, 536, 7, 122, 2, 2,
	536, 537, 7, 51, 2, 2, 537, 538, 7, 52, 2, 2, 538, 539, 7, 58, 2, 2, 539,
	116, 3, 2, 2, 2, 540, 541, 7, 117, 2, 2, 541, 542, 7, 118, 2, 2, 542, 543,
	7, 116, 2, 2, 543, 544, 7, 107, 2, 2, 544, 545, 7, 112, 2, 2, 545, 546,
	7, 105, 2, 2, 546, 118, 3, 2, 2, 2, 547, 548, 7, 107, 2, 2, 548, 549, 7,
	112, 2, 2, 549, 550, 7, 118, 2, 2, 550, 120, 3, 2, 2, 2, 551, 552, 7, 119,
	2, 2, 552, 553, 7, 107, 2, 2, 553, 554, 7, 112, 2, 2, 554, 555, 7, 118,
	2, 2, 555, 122, 3, 2, 2, 2, 556, 557, 7, 119, 2, 2, 557, 558, 7, 107, 2,
	2, 558, 559, 7, 112, 2, 2, 559, 560, 7, 118, 2, 2, 560, 561, 7, 114, 2,
	2, 561, 562, 7, 118, 2, 2, 562, 563, 7, 116, 2, 2, 563, 124, 3, 2, 2, 2,
	564, 565, 7, 100, 2, 2, 565, 566, 7, 123, 2, 2, 566, 567, 7, 118, 2, 2,
	567, 568, 7, 103, 2, 2, 568, 126, 3, 2, 2, 2, 569, 570, 7, 116, 2, 2, 570,
	571, 7, 119, 2, 2, 571, 572, 7, 112, 2, 2, 572, 573, 7, 103, 2, 2, 573,
	128, 3, 2, 2, 2, 574, 575, 7, 118, 2, 2, 575, 576, 7, 107, 2, 2, 576, 577,
	7, 111, 2, 2, 577, 578, 7, 103, 2, 2, 578, 579, 7, 48, 2, 2, 579, 580,
	7, 86, 2, 2, 580, 581, 7, 107, 2, 2, 581, 582, 7, 111, 2, 2, 582, 583,
	7, 103, 2, 2, 583, 130, 3, 2, 2, 2, 584, 585, 7, 105, 2, 2, 585, 586, 7,
	103, 2, 2, 586, 587, 7, 118, 2, 2, 587, 132, 3, 2, 2, 2, 588, 589, 7, 106,
	2, 2, 589, 590, 7, 103, 2, 2, 590, 591, 7, 99, 2, 2, 591, 592, 7, 102,
	2, 2, 592, 134, 3, 2, 2, 2, 593, 594, 7, 114, 2, 2, 594, 595, 7, 113, 2,
	2, 595, 596, 7, 117, 2, 2, 596, 597, 7, 118, 2, 2, 597, 136, 3, 2, 2, 2,
	598, 599, 7, 114, 2, 2, 599, 600, 7, 119, 2, 2, 600, 601, 7, 118, 2, 2,
	601, 138, 3, 2, 2, 2, 602, 603, 7, 114, 2, 2, 603, 604, 7, 99, 2, 2, 604,
	605, 7, 118, 2, 2, 605, 606, 7, 101, 2, 2, 606, 607, 7, 106, 2, 2, 607,
	140, 3, 2, 2, 2, 608, 609, 7, 102, 2, 2, 609, 610, 7, 103, 2, 2, 610, 611,
	7, 110, 2, 2, 611, 612, 7, 103, 2, 2, 612, 613, 7, 118, 2, 2, 613, 614,
	7, 103, 2, 2, 614, 142, 3, 2, 2, 2, 615, 616, 7, 101, 2, 2, 616, 617, 7,
	113, 2, 2, 617, 618, 7, 112, 2, 2, 618, 619, 7, 112, 2, 2, 619, 620, 7,
	103, 2, 2, 620, 621, 7, 101, 2, 2, 621, 622, 7, 118, 2, 2, 622, 144, 3,
	2, 2, 2, 623, 624, 7, 113, 2, 2, 624, 625, 7, 114, 2, 2, 625, 626, 7, 118,
	2, 2, 626, 627, 7, 107, 2, 2, 627, 628, 7, 113, 2, 2, 628, 629, 7, 112,
	2, 2, 629, 630, 7, 117, 2, 2, 630, 146, 3, 2, 2, 2, 631, 632, 7, 118, 2,
	2, 632, 633, 7, 116, 2, 2, 633, 634, 7, 99, 2, 2, 634, 635, 7, 101, 2,
	2, 635, 636, 7, 103, 2, 2, 636, 148, 3, 2, 2, 2, 28, 2, 169, 265, 300,
	310, 324, 335, 342, 349, 362, 364, 372, 374, 383, 388, 396, 399, 406, 413,
	418, 420, 424, 432, 436, 440, 446, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'syntax'", "'info'", "'map'", "'struct'", "'interface{}'", "'type'",
	"'service'", "'returns'", "'import'", "'@server'", "'@doc'", "'@handler'",
	"", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'", "'/'", "'?'",
	"'&'", "'='", "'-'", "':'", "'*'",
}

var lexerSymbolicNames = []string{
	"", "GOTYPE", "SYNTAX", "INFO", "MAP", "STRUCT", "INTERFACE", "TYPE", "SERVICE",
	"RETURNS", "IMPORT", "ATSERVER", "ATDOC", "ATHANDLER", "HTTPMETHOD", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA", "DOT", "SLASH",
	"QUESTION", "BITAND", "ASSIGN", "SUB", "COLON", "STAR", "WS", "COMMENT",
	"LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH", "STRING_LIT", "RAW_STRING",
	"ID",
}

var lexerRuleNames = []string{
	"GOTYPE", "SYNTAX", "INFO", "MAP", "STRUCT", "INTERFACE", "TYPE", "SERVICE",
	"RETURNS", "IMPORT", "ATSERVER", "ATDOC", "ATHANDLER", "HTTPMETHOD", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA", "DOT", "SLASH",
	"QUESTION", "BITAND", "ASSIGN", "SUB", "COLON", "STAR", "WS", "COMMENT",
	"LINE_COMMENT", "SYNTAX_VERSION", "IMPORT_PATH", "STRING_LIT", "RAW_STRING",
	"ID", "ExponentPart", "EscapeSequence", "HexDigits", "HexDigit", "Digits",
	"LetterOrDigit", "Letter", "BOOL", "UINT8", "UINT16", "UINT32", "UINT64",
	"INT8", "INT16", "INT32", "INT64", "FLOAT32", "FLOAT64", "COMPLEX64", "COMPLEX128",
	"STRING", "INT", "UINT", "UINTPTR", "BYTE", "RUNE", "TIME", "GET", "HEAD",
	"POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE",
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
	ApiLexerGOTYPE         = 1
	ApiLexerSYNTAX         = 2
	ApiLexerINFO           = 3
	ApiLexerMAP            = 4
	ApiLexerSTRUCT         = 5
	ApiLexerINTERFACE      = 6
	ApiLexerTYPE           = 7
	ApiLexerSERVICE        = 8
	ApiLexerRETURNS        = 9
	ApiLexerIMPORT         = 10
	ApiLexerATSERVER       = 11
	ApiLexerATDOC          = 12
	ApiLexerATHANDLER      = 13
	ApiLexerHTTPMETHOD     = 14
	ApiLexerLPAREN         = 15
	ApiLexerRPAREN         = 16
	ApiLexerLBRACE         = 17
	ApiLexerRBRACE         = 18
	ApiLexerLBRACK         = 19
	ApiLexerRBRACK         = 20
	ApiLexerCOMMA          = 21
	ApiLexerDOT            = 22
	ApiLexerSLASH          = 23
	ApiLexerQUESTION       = 24
	ApiLexerBITAND         = 25
	ApiLexerASSIGN         = 26
	ApiLexerSUB            = 27
	ApiLexerCOLON          = 28
	ApiLexerSTAR           = 29
	ApiLexerWS             = 30
	ApiLexerCOMMENT        = 31
	ApiLexerLINE_COMMENT   = 32
	ApiLexerSYNTAX_VERSION = 33
	ApiLexerIMPORT_PATH    = 34
	ApiLexerSTRING_LIT     = 35
	ApiLexerRAW_STRING     = 36
	ApiLexerID             = 37
)
