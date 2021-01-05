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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 425,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 7,
	2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 89, 10, 3, 3, 4, 5, 4, 92, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 100, 10, 4, 3, 5, 3, 5, 5, 5, 104, 10, 5, 3, 6, 5, 6, 107,
	10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 113, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 119, 10, 7, 3, 7, 3, 7, 6, 7, 123, 10, 7, 13, 7, 14, 7, 124, 3, 7,
	3, 7, 3, 8, 5, 8, 130, 10, 8, 3, 8, 3, 8, 5, 8, 134, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 5, 10, 140, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 146,
	10, 10, 13, 10, 14, 10, 147, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 154, 10,
	11, 3, 12, 5, 12, 157, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 6, 13, 167, 10, 13, 13, 13, 14, 13, 168, 3, 13, 3, 13, 3,
	14, 3, 14, 5, 14, 175, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 180, 10, 15,
	3, 16, 3, 16, 3, 16, 5, 16, 185, 10, 16, 3, 16, 3, 16, 6, 16, 189, 10,
	16, 13, 16, 14, 16, 190, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 198,
	10, 17, 3, 17, 3, 17, 5, 17, 202, 10, 17, 3, 18, 5, 18, 205, 10, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 210, 10, 18, 3, 18, 3, 18, 6, 18, 214, 10, 18,
	13, 18, 14, 18, 215, 3, 18, 3, 18, 3, 19, 5, 19, 221, 10, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 226, 10, 19, 3, 19, 3, 19, 5, 19, 230, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 5, 20, 236, 10, 20, 3, 21, 7, 21, 239, 10, 21, 12,
	21, 14, 21, 242, 11, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 248, 10, 21,
	3, 21, 5, 21, 251, 10, 21, 3, 22, 7, 22, 254, 10, 22, 12, 22, 14, 22, 257,
	11, 22, 3, 22, 5, 22, 260, 10, 22, 3, 22, 3, 22, 5, 22, 264, 10, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 274, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 5, 27, 293, 10, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 6, 28, 300, 10, 28, 13, 28, 14, 28, 301, 3,
	28, 3, 28, 7, 28, 306, 10, 28, 12, 28, 14, 28, 309, 11, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 5, 29, 316, 10, 29, 3, 29, 7, 29, 319, 10, 29, 12,
	29, 14, 29, 322, 11, 29, 3, 29, 3, 29, 3, 30, 5, 30, 327, 10, 30, 3, 30,
	3, 30, 5, 30, 331, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 338,
	10, 31, 3, 31, 6, 31, 341, 10, 31, 13, 31, 14, 31, 342, 3, 31, 5, 31, 346,
	10, 31, 3, 31, 3, 31, 3, 32, 7, 32, 351, 10, 32, 12, 32, 14, 32, 354, 11,
	32, 3, 32, 3, 32, 3, 32, 5, 32, 359, 10, 32, 3, 33, 7, 33, 362, 10, 33,
	12, 33, 14, 33, 365, 11, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 371, 10,
	33, 3, 33, 5, 33, 374, 10, 33, 3, 33, 5, 33, 377, 10, 33, 3, 33, 5, 33,
	380, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 7, 35, 388, 10,
	35, 12, 35, 14, 35, 391, 11, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 397,
	10, 35, 3, 36, 3, 36, 3, 37, 3, 37, 5, 37, 403, 10, 37, 6, 37, 405, 10,
	37, 13, 37, 14, 37, 406, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 413, 10, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 419, 10, 38, 6, 38, 421, 10, 38, 13,
	38, 14, 38, 422, 3, 38, 2, 2, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 2, 3, 3, 2, 20, 21, 2, 460, 2, 79, 3, 2,
	2, 2, 4, 88, 3, 2, 2, 2, 6, 91, 3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 106,
	3, 2, 2, 2, 12, 114, 3, 2, 2, 2, 14, 129, 3, 2, 2, 2, 16, 135, 3, 2, 2,
	2, 18, 139, 3, 2, 2, 2, 20, 153, 3, 2, 2, 2, 22, 156, 3, 2, 2, 2, 24, 162,
	3, 2, 2, 2, 26, 174, 3, 2, 2, 2, 28, 179, 3, 2, 2, 2, 30, 181, 3, 2, 2,
	2, 32, 194, 3, 2, 2, 2, 34, 204, 3, 2, 2, 2, 36, 220, 3, 2, 2, 2, 38, 235,
	3, 2, 2, 2, 40, 240, 3, 2, 2, 2, 42, 255, 3, 2, 2, 2, 44, 273, 3, 2, 2,
	2, 46, 275, 3, 2, 2, 2, 48, 279, 3, 2, 2, 2, 50, 287, 3, 2, 2, 2, 52, 292,
	3, 2, 2, 2, 54, 296, 3, 2, 2, 2, 56, 310, 3, 2, 2, 2, 58, 326, 3, 2, 2,
	2, 60, 334, 3, 2, 2, 2, 62, 352, 3, 2, 2, 2, 64, 363, 3, 2, 2, 2, 66, 381,
	3, 2, 2, 2, 68, 389, 3, 2, 2, 2, 70, 398, 3, 2, 2, 2, 72, 404, 3, 2, 2,
	2, 74, 420, 3, 2, 2, 2, 76, 78, 5, 4, 3, 2, 77, 76, 3, 2, 2, 2, 78, 81,
	3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 3, 3, 2, 2, 2,
	81, 79, 3, 2, 2, 2, 82, 89, 5, 6, 4, 2, 83, 89, 5, 8, 5, 2, 84, 89, 5,
	18, 10, 2, 85, 89, 5, 20, 11, 2, 86, 89, 5, 52, 27, 2, 87, 89, 5, 70, 36,
	2, 88, 82, 3, 2, 2, 2, 88, 83, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88, 85,
	3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 5, 3, 2, 2, 2,
	90, 92, 5, 70, 36, 2, 91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 3,
	2, 2, 2, 93, 94, 8, 4, 1, 2, 94, 95, 7, 25, 2, 2, 95, 96, 7, 3, 2, 2, 96,
	97, 8, 4, 1, 2, 97, 99, 7, 22, 2, 2, 98, 100, 5, 70, 36, 2, 99, 98, 3,
	2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 7, 3, 2, 2, 2, 101, 104, 5, 10, 6, 2,
	102, 104, 5, 12, 7, 2, 103, 101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104,
	9, 3, 2, 2, 2, 105, 107, 5, 70, 36, 2, 106, 105, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 8, 6, 1, 2, 109, 110, 7, 25,
	2, 2, 110, 112, 5, 16, 9, 2, 111, 113, 5, 70, 36, 2, 112, 111, 3, 2, 2,
	2, 112, 113, 3, 2, 2, 2, 113, 11, 3, 2, 2, 2, 114, 115, 8, 7, 1, 2, 115,
	116, 7, 25, 2, 2, 116, 118, 7, 4, 2, 2, 117, 119, 5, 70, 36, 2, 118, 117,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 123, 5, 14,
	8, 2, 121, 123, 5, 70, 36, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 127, 7, 5, 2, 2, 127, 13, 3, 2, 2, 2, 128, 130, 5,
	70, 36, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3, 2,
	2, 2, 131, 133, 5, 16, 9, 2, 132, 134, 5, 70, 36, 2, 133, 132, 3, 2, 2,
	2, 133, 134, 3, 2, 2, 2, 134, 15, 3, 2, 2, 2, 135, 136, 8, 9, 1, 2, 136,
	137, 7, 22, 2, 2, 137, 17, 3, 2, 2, 2, 138, 140, 5, 70, 36, 2, 139, 138,
	3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 8, 10,
	1, 2, 142, 143, 7, 25, 2, 2, 143, 145, 7, 4, 2, 2, 144, 146, 5, 68, 35,
	2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 7, 5, 2, 2, 150, 19, 3,
	2, 2, 2, 151, 154, 5, 22, 12, 2, 152, 154, 5, 24, 13, 2, 153, 151, 3, 2,
	2, 2, 153, 152, 3, 2, 2, 2, 154, 21, 3, 2, 2, 2, 155, 157, 5, 70, 36, 2,
	156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158,
	159, 8, 12, 1, 2, 159, 160, 7, 25, 2, 2, 160, 161, 5, 26, 14, 2, 161, 23,
	3, 2, 2, 2, 162, 163, 8, 13, 1, 2, 163, 164, 7, 25, 2, 2, 164, 166, 7,
	4, 2, 2, 165, 167, 5, 28, 15, 2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2,
	2, 2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2,
	170, 171, 7, 5, 2, 2, 171, 25, 3, 2, 2, 2, 172, 175, 5, 30, 16, 2, 173,
	175, 5, 32, 17, 2, 174, 172, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 27,
	3, 2, 2, 2, 176, 180, 5, 34, 18, 2, 177, 180, 5, 36, 19, 2, 178, 180, 5,
	70, 36, 2, 179, 176, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3, 2,
	2, 2, 180, 29, 3, 2, 2, 2, 181, 182, 8, 16, 1, 2, 182, 184, 7, 25, 2, 2,
	183, 185, 7, 25, 2, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 188, 7, 6, 2, 2, 187, 189, 5, 38, 20, 2, 188, 187,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2,
	2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 7, 7, 2, 2, 193, 31, 3, 2, 2, 2,
	194, 195, 8, 17, 1, 2, 195, 197, 7, 25, 2, 2, 196, 198, 7, 3, 2, 2, 197,
	196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201,
	5, 44, 23, 2, 200, 202, 5, 70, 36, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3,
	2, 2, 2, 202, 33, 3, 2, 2, 2, 203, 205, 5, 70, 36, 2, 204, 203, 3, 2, 2,
	2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 8, 18, 1, 2, 207,
	209, 7, 25, 2, 2, 208, 210, 7, 25, 2, 2, 209, 208, 3, 2, 2, 2, 209, 210,
	3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 7, 6, 2, 2, 212, 214, 5, 38,
	20, 2, 213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2,
	215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 7, 7, 2, 2, 218,
	35, 3, 2, 2, 2, 219, 221, 5, 70, 36, 2, 220, 219, 3, 2, 2, 2, 220, 221,
	3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 8, 19, 1, 2, 223, 225, 7, 25,
	2, 2, 224, 226, 7, 3, 2, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2,
	226, 227, 3, 2, 2, 2, 227, 229, 5, 44, 23, 2, 228, 230, 5, 70, 36, 2, 229,
	228, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 37, 3, 2, 2, 2, 231, 232, 6,
	20, 2, 2, 232, 236, 5, 40, 21, 2, 233, 236, 5, 42, 22, 2, 234, 236, 5,
	70, 36, 2, 235, 231, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 234, 3, 2,
	2, 2, 236, 39, 3, 2, 2, 2, 237, 239, 5, 70, 36, 2, 238, 237, 3, 2, 2, 2,
	239, 242, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241,
	243, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 243, 244, 8, 21, 1, 2, 244, 245,
	7, 25, 2, 2, 245, 247, 5, 44, 23, 2, 246, 248, 7, 23, 2, 2, 247, 246, 3,
	2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249, 251, 5, 70, 36,
	2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 41, 3, 2, 2, 2, 252,
	254, 5, 70, 36, 2, 253, 252, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253,
	3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2,
	2, 2, 258, 260, 7, 8, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2,
	260, 261, 3, 2, 2, 2, 261, 263, 7, 25, 2, 2, 262, 264, 5, 70, 36, 2, 263,
	262, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 43, 3, 2, 2, 2, 265, 266, 8,
	23, 1, 2, 266, 274, 7, 25, 2, 2, 267, 274, 5, 48, 25, 2, 268, 274, 5, 50,
	26, 2, 269, 274, 7, 17, 2, 2, 270, 274, 7, 9, 2, 2, 271, 274, 5, 46, 24,
	2, 272, 274, 5, 30, 16, 2, 273, 265, 3, 2, 2, 2, 273, 267, 3, 2, 2, 2,
	273, 268, 3, 2, 2, 2, 273, 269, 3, 2, 2, 2, 273, 270, 3, 2, 2, 2, 273,
	271, 3, 2, 2, 2, 273, 272, 3, 2, 2, 2, 274, 45, 3, 2, 2, 2, 275, 276, 7,
	8, 2, 2, 276, 277, 8, 24, 1, 2, 277, 278, 7, 25, 2, 2, 278, 47, 3, 2, 2,
	2, 279, 280, 8, 25, 1, 2, 280, 281, 7, 25, 2, 2, 281, 282, 7, 10, 2, 2,
	282, 283, 8, 25, 1, 2, 283, 284, 7, 25, 2, 2, 284, 285, 7, 11, 2, 2, 285,
	286, 5, 44, 23, 2, 286, 49, 3, 2, 2, 2, 287, 288, 7, 10, 2, 2, 288, 289,
	7, 11, 2, 2, 289, 290, 5, 44, 23, 2, 290, 51, 3, 2, 2, 2, 291, 293, 5,
	54, 28, 2, 292, 291, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294, 3, 2,
	2, 2, 294, 295, 5, 56, 29, 2, 295, 53, 3, 2, 2, 2, 296, 297, 7, 18, 2,
	2, 297, 299, 7, 4, 2, 2, 298, 300, 5, 68, 35, 2, 299, 298, 3, 2, 2, 2,
	300, 301, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302,
	303, 3, 2, 2, 2, 303, 307, 7, 5, 2, 2, 304, 306, 5, 70, 36, 2, 305, 304,
	3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2,
	2, 2, 308, 55, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311, 8, 29, 1, 2,
	311, 312, 7, 25, 2, 2, 312, 313, 5, 72, 37, 2, 313, 315, 7, 6, 2, 2, 314,
	316, 5, 70, 36, 2, 315, 314, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 320,
	3, 2, 2, 2, 317, 319, 5, 58, 30, 2, 318, 317, 3, 2, 2, 2, 319, 322, 3,
	2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 323, 3, 2, 2,
	2, 322, 320, 3, 2, 2, 2, 323, 324, 7, 7, 2, 2, 324, 57, 3, 2, 2, 2, 325,
	327, 5, 60, 31, 2, 326, 325, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 330,
	3, 2, 2, 2, 328, 331, 5, 54, 28, 2, 329, 331, 5, 62, 32, 2, 330, 328, 3,
	2, 2, 2, 330, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 333, 5, 64, 33,
	2, 333, 59, 3, 2, 2, 2, 334, 335, 7, 15, 2, 2, 335, 345, 7, 4, 2, 2, 336,
	338, 5, 70, 36, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 340,
	3, 2, 2, 2, 339, 341, 5, 68, 35, 2, 340, 339, 3, 2, 2, 2, 341, 342, 3,
	2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 346, 3, 2, 2,
	2, 344, 346, 7, 22, 2, 2, 345, 337, 3, 2, 2, 2, 345, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 348, 7, 5, 2, 2, 348, 61, 3, 2, 2, 2, 349, 351, 5,
	70, 36, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2,
	2, 2, 352, 353, 3, 2, 2, 2, 353, 355, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2,
	355, 356, 7, 16, 2, 2, 356, 358, 7, 25, 2, 2, 357, 359, 5, 70, 36, 2, 358,
	357, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 63, 3, 2, 2, 2, 360, 362, 5,
	70, 36, 2, 361, 360, 3, 2, 2, 2, 362, 365, 3, 2, 2, 2, 363, 361, 3, 2,
	2, 2, 363, 364, 3, 2, 2, 2, 364, 366, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2,
	366, 367, 8, 33, 1, 2, 367, 368, 7, 25, 2, 2, 368, 370, 5, 74, 38, 2, 369,
	371, 5, 66, 34, 2, 370, 369, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 373,
	3, 2, 2, 2, 372, 374, 7, 25, 2, 2, 373, 372, 3, 2, 2, 2, 373, 374, 3, 2,
	2, 2, 374, 376, 3, 2, 2, 2, 375, 377, 5, 66, 34, 2, 376, 375, 3, 2, 2,
	2, 376, 377, 3, 2, 2, 2, 377, 379, 3, 2, 2, 2, 378, 380, 5, 70, 36, 2,
	379, 378, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 65, 3, 2, 2, 2, 381, 382,
	7, 4, 2, 2, 382, 383, 8, 34, 1, 2, 383, 384, 7, 25, 2, 2, 384, 385, 7,
	5, 2, 2, 385, 67, 3, 2, 2, 2, 386, 388, 5, 70, 36, 2, 387, 386, 3, 2, 2,
	2, 388, 391, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390,
	392, 3, 2, 2, 2, 391, 389, 3, 2, 2, 2, 392, 393, 7, 25, 2, 2, 393, 394,
	8, 35, 1, 2, 394, 396, 7, 24, 2, 2, 395, 397, 5, 70, 36, 2, 396, 395, 3,
	2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 69, 3, 2, 2, 2, 398, 399, 9, 2, 2,
	2, 399, 71, 3, 2, 2, 2, 400, 402, 7, 25, 2, 2, 401, 403, 7, 12, 2, 2, 402,
	401, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 405, 3, 2, 2, 2, 404, 400,
	3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2,
	2, 2, 407, 73, 3, 2, 2, 2, 408, 409, 7, 13, 2, 2, 409, 412, 7, 25, 2, 2,
	410, 411, 7, 12, 2, 2, 411, 413, 7, 25, 2, 2, 412, 410, 3, 2, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 421, 3, 2, 2, 2, 414, 415, 7, 14, 2, 2, 415, 418,
	7, 25, 2, 2, 416, 417, 7, 12, 2, 2, 417, 419, 7, 25, 2, 2, 418, 416, 3,
	2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 421, 3, 2, 2, 2, 420, 408, 3, 2, 2,
	2, 420, 414, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422,
	423, 3, 2, 2, 2, 423, 75, 3, 2, 2, 2, 64, 79, 88, 91, 99, 103, 106, 112,
	118, 122, 124, 129, 133, 139, 147, 153, 156, 168, 174, 179, 184, 190, 197,
	201, 204, 209, 215, 220, 225, 229, 235, 240, 247, 250, 255, 259, 263, 273,
	292, 301, 307, 315, 320, 326, 330, 337, 342, 345, 352, 358, 363, 370, 373,
	376, 379, 389, 396, 402, 406, 412, 418, 420, 422,
}
var literalNames = []string{
	"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'time.Time'", "'['", "']'",
	"'-'", "'/'", "'/:'", "'@doc'", "'@handler'", "'interface{}'", "'@server'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "ATDOC", "ATHANDLER",
	"INTERFACE", "ATSERVER", "WS", "COMMENT", "LINE_COMMENT", "STRING", "RAW_STRING",
	"LINE_VALUE", "ID",
}

var ruleNames = []string{
	"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock", "importBlockValue",
	"importValue", "infoSpec", "typeSpec", "typeLit", "typeBlock", "typeLitBody",
	"typeBlockBody", "typeStruct", "typeAlias", "typeBlockStruct", "typeBlockAlias",
	"field", "normalField", "anonymousFiled", "dataType", "pointerType", "mapType",
	"arrayType", "serviceSpec", "atServer", "serviceApi", "serviceRoute", "atDoc",
	"atHandler", "route", "body", "kvLit", "commentSpec", "serviceName", "path",
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
	ApiParserParserT__9         = 10
	ApiParserParserT__10        = 11
	ApiParserParserT__11        = 12
	ApiParserParserATDOC        = 13
	ApiParserParserATHANDLER    = 14
	ApiParserParserINTERFACE    = 15
	ApiParserParserATSERVER     = 16
	ApiParserParserWS           = 17
	ApiParserParserCOMMENT      = 18
	ApiParserParserLINE_COMMENT = 19
	ApiParserParserSTRING       = 20
	ApiParserParserRAW_STRING   = 21
	ApiParserParserLINE_VALUE   = 22
	ApiParserParserID           = 23
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
	ApiParserParserRULE_serviceSpec      = 25
	ApiParserParserRULE_atServer         = 26
	ApiParserParserRULE_serviceApi       = 27
	ApiParserParserRULE_serviceRoute     = 28
	ApiParserParserRULE_atDoc            = 29
	ApiParserParserRULE_atHandler        = 30
	ApiParserParserRULE_route            = 31
	ApiParserParserRULE_body             = 32
	ApiParserParserRULE_kvLit            = 33
	ApiParserParserRULE_commentSpec      = 34
	ApiParserParserRULE_serviceName      = 35
	ApiParserParserRULE_path             = 36
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserATSERVER)|(1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0 {
		{
			p.SetState(74)
			p.Spec()
		}

		p.SetState(79)
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

func (s *SpecContext) ServiceSpec() IServiceSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceSpecContext)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.ServiceSpec()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(88)

			var _x = p.CommentSpec()

			localctx.(*SyntaxLitContext).doc = _x
		}

	}
	match(p, "syntax")
	{
		p.SetState(92)

		var _m = p.Match(ApiParserParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
	}
	{
		p.SetState(93)

		var _m = p.Match(ApiParserParserT__0)

		localctx.(*SyntaxLitContext).assign = _m
	}
	checkVersion(p)
	{
		p.SetState(95)

		var _m = p.Match(ApiParserParserSTRING)

		localctx.(*SyntaxLitContext).version = _m
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(96)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
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
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(103)

			var _x = p.CommentSpec()

			localctx.(*ImportLitContext).doc = _x
		}

	}
	match(p, "import")
	{
		p.SetState(107)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportLitContext).importToken = _m
	}
	{
		p.SetState(108)
		p.ImportValue()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(109)

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

func (s *ImportBlockContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *ImportBlockContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

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
		p.SetState(113)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportBlockContext).importToken = _m
	}
	{
		p.SetState(114)
		p.Match(ApiParserParserT__1)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(115)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockContext).comment = _x
		}

	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserSTRING))) != 0) {
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(118)
				p.ImportBlockValue()
			}

		case 2:
			{
				p.SetState(119)
				p.CommentSpec()
			}

		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
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
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(126)

			var _x = p.CommentSpec()

			localctx.(*ImportBlockValueContext).doc = _x
		}

	}
	{
		p.SetState(129)
		p.ImportValue()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(130)

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
		p.SetState(134)
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

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// IsInfoSpecContext differentiates from other interfaces.
	IsInfoSpecContext()
}

type InfoSpecContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	doc       ICommentSpecContext
	infoToken antlr.Token
	lp        antlr.Token
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

func (s *InfoSpecContext) SetDoc(v ICommentSpecContext) { s.doc = v }

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

func (s *InfoSpecContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(136)

			var _x = p.CommentSpec()

			localctx.(*InfoSpecContext).doc = _x
		}

	}
	match(p, "info")
	{
		p.SetState(140)

		var _m = p.Match(ApiParserParserID)

		localctx.(*InfoSpecContext).infoToken = _m
	}
	{
		p.SetState(141)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*InfoSpecContext).lp = _m
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(142)
			p.KvLit()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(153)

			var _x = p.CommentSpec()

			localctx.(*TypeLitContext).doc = _x
		}

	}
	match(p, "type")
	{
		p.SetState(157)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeLitContext).typeToken = _m
	}
	{
		p.SetState(158)
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
		p.SetState(161)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockContext).typeToken = _m
	}
	{
		p.SetState(162)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*TypeBlockContext).lp = _m
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(163)
			p.TypeBlockBody()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)

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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.TypeStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
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

func (s *TypeBlockBodyContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.TypeBlockStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.TypeBlockAlias()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.CommentSpec()
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

	// IsTypeStructContext differentiates from other interfaces.
	IsTypeStructContext()
}

type TypeStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
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
		p.SetState(180)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeStructContext).structName = _m
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(181)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeStructContext).structToken = _m
		}

	}
	{
		p.SetState(184)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeStructContext).lbrace = _m
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(185)
				p.Field()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	{
		p.SetState(190)

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
		p.SetState(193)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(194)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeAliasContext).assign = _m
		}

	}
	{
		p.SetState(197)
		p.DataType()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(198)

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

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

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

func (s *TypeBlockStructContext) SetDoc(v ICommentSpecContext) { s.doc = v }

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

func (s *TypeBlockStructContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(201)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockStructContext).doc = _x
		}

	}
	checkKeyword(p)
	{
		p.SetState(205)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockStructContext).structName = _m
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(206)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeBlockStructContext).structToken = _m
		}

	}
	{
		p.SetState(209)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeBlockStructContext).lbrace = _m
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(210)
				p.Field()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	{
		p.SetState(215)

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
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(217)

			var _x = p.CommentSpec()

			localctx.(*TypeBlockAliasContext).doc = _x
		}

	}
	checkKeyword(p)
	{
		p.SetState(221)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockAliasContext).alias = _m
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(222)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeBlockAliasContext).assign = _m
		}

	}
	{
		p.SetState(225)
		p.DataType()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(226)

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

func (s *FieldContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(229)

		if !(isNormal(p)) {
			panic(antlr.NewFailedPredicateException(p, "isNormal(p)", ""))
		}
		{
			p.SetState(230)
			p.NormalField()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.AnonymousFiled()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.CommentSpec()
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(235)

			var _x = p.CommentSpec()

			localctx.(*NormalFieldContext).doc = _x
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	checkKeyword(p)
	{
		p.SetState(242)

		var _m = p.Match(ApiParserParserID)

		localctx.(*NormalFieldContext).fieldName = _m
	}
	{
		p.SetState(243)
		p.DataType()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)

			var _m = p.Match(ApiParserParserRAW_STRING)

			localctx.(*NormalFieldContext).tag = _m
		}

	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(247)

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
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(250)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).doc = _x
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__5 {
		{
			p.SetState(256)

			var _m = p.Match(ApiParserParserT__5)

			localctx.(*AnonymousFiledContext).star = _m
		}

	}
	{
		p.SetState(259)
		p.Match(ApiParserParserID)
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(260)

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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		isInterface(p)
		{
			p.SetState(264)
			p.Match(ApiParserParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.MapType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(266)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(267)

			var _m = p.Match(ApiParserParserINTERFACE)

			localctx.(*DataTypeContext).inter = _m
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(268)

			var _m = p.Match(ApiParserParserT__6)

			localctx.(*DataTypeContext).time = _m
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(269)
			p.PointerType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(270)
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
		p.SetState(273)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*PointerTypeContext).star = _m
	}
	checkKeyword(p)
	{
		p.SetState(275)
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
		p.SetState(278)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).mapToken = _m
	}
	{
		p.SetState(279)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*MapTypeContext).lbrack = _m
	}
	checkKey(p)
	{
		p.SetState(281)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).key = _m
	}
	{
		p.SetState(282)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*MapTypeContext).rbrack = _m
	}
	{
		p.SetState(283)

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
		p.SetState(285)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*ArrayTypeContext).lbrack = _m
	}
	{
		p.SetState(286)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*ArrayTypeContext).rbrack = _m
	}
	{
		p.SetState(287)
		p.DataType()
	}

	return localctx
}

// IServiceSpecContext is an interface to support dynamic dispatch.
type IServiceSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceSpecContext differentiates from other interfaces.
	IsServiceSpecContext()
}

type ServiceSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceSpecContext() *ServiceSpecContext {
	var p = new(ServiceSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceSpec
	return p
}

func (*ServiceSpecContext) IsServiceSpecContext() {}

func NewServiceSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceSpecContext {
	var p = new(ServiceSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceSpec

	return p
}

func (s *ServiceSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceSpecContext) ServiceApi() IServiceApiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceApiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceApiContext)
}

func (s *ServiceSpecContext) AtServer() IAtServerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtServerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}

func (s *ServiceSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceSpec() (localctx IServiceSpecContext) {
	localctx = NewServiceSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ApiParserParserRULE_serviceSpec)
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
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATSERVER {
		{
			p.SetState(289)
			p.AtServer()
		}

	}
	{
		p.SetState(292)
		p.ServiceApi()
	}

	return localctx
}

// IAtServerContext is an interface to support dynamic dispatch.
type IAtServerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// IsAtServerContext differentiates from other interfaces.
	IsAtServerContext()
}

type AtServerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtServerContext() *AtServerContext {
	var p = new(AtServerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_atServer
	return p
}

func (*AtServerContext) IsAtServerContext() {}

func NewAtServerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtServerContext {
	var p = new(AtServerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atServer

	return p
}

func (s *AtServerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtServerContext) GetLp() antlr.Token { return s.lp }

func (s *AtServerContext) GetRp() antlr.Token { return s.rp }

func (s *AtServerContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtServerContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtServerContext) ATSERVER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATSERVER, 0)
}

func (s *AtServerContext) AllKvLit() []IKvLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvLitContext)(nil)).Elem())
	var tst = make([]IKvLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvLitContext)
		}
	}

	return tst
}

func (s *AtServerContext) KvLit(i int) IKvLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtServerContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *AtServerContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *AtServerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtServerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtServerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtServer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtServer() (localctx IAtServerContext) {
	localctx = NewAtServerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ApiParserParserRULE_atServer)
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
	{
		p.SetState(294)
		p.Match(ApiParserParserATSERVER)
	}
	{
		p.SetState(295)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*AtServerContext).lp = _m
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
		{
			p.SetState(296)
			p.KvLit()
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(301)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*AtServerContext).rp = _m
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(302)
				p.CommentSpec()
			}

		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IServiceApiContext is an interface to support dynamic dispatch.
type IServiceApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetServiceToken returns the serviceToken token.
	GetServiceToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetServiceToken sets the serviceToken token.
	SetServiceToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsServiceApiContext differentiates from other interfaces.
	IsServiceApiContext()
}

type ServiceApiContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	serviceToken antlr.Token
	lbrace       antlr.Token
	comment      ICommentSpecContext
	rbrace       antlr.Token
}

func NewEmptyServiceApiContext() *ServiceApiContext {
	var p = new(ServiceApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceApi
	return p
}

func (*ServiceApiContext) IsServiceApiContext() {}

func NewServiceApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceApiContext {
	var p = new(ServiceApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceApi

	return p
}

func (s *ServiceApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceApiContext) GetServiceToken() antlr.Token { return s.serviceToken }

func (s *ServiceApiContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *ServiceApiContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *ServiceApiContext) SetServiceToken(v antlr.Token) { s.serviceToken = v }

func (s *ServiceApiContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *ServiceApiContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *ServiceApiContext) GetComment() ICommentSpecContext { return s.comment }

func (s *ServiceApiContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *ServiceApiContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceApiContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ServiceApiContext) AllServiceRoute() []IServiceRouteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceRouteContext)(nil)).Elem())
	var tst = make([]IServiceRouteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceRouteContext)
		}
	}

	return tst
}

func (s *ServiceApiContext) ServiceRoute(i int) IServiceRouteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceRouteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceRouteContext)
}

func (s *ServiceApiContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *ServiceApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceApi() (localctx IServiceApiContext) {
	localctx = NewServiceApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ApiParserParserRULE_serviceApi)
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
	match(p, "service")
	{
		p.SetState(309)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ServiceApiContext).serviceToken = _m
	}
	{
		p.SetState(310)
		p.ServiceName()
	}
	{
		p.SetState(311)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*ServiceApiContext).lbrace = _m
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(312)

			var _x = p.CommentSpec()

			localctx.(*ServiceApiContext).comment = _x
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserATDOC)|(1<<ApiParserParserATHANDLER)|(1<<ApiParserParserATSERVER)|(1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT))) != 0 {
		{
			p.SetState(315)
			p.ServiceRoute()
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(321)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*ServiceApiContext).rbrace = _m
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
	p.RuleIndex = ApiParserParserRULE_serviceRoute
	return p
}

func (*ServiceRouteContext) IsServiceRouteContext() {}

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	var p = new(ServiceRouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceRoute

	return p
}

func (s *ServiceRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceRouteContext) Route() IRouteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRouteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *ServiceRouteContext) AtServer() IAtServerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtServerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}

func (s *ServiceRouteContext) AtHandler() IAtHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtHandlerContext)
}

func (s *ServiceRouteContext) AtDoc() IAtDocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtDocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtDocContext)
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

func (p *ApiParserParser) ServiceRoute() (localctx IServiceRouteContext) {
	localctx = NewServiceRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ApiParserParserRULE_serviceRoute)
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
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATDOC {
		{
			p.SetState(323)
			p.AtDoc()
		}

	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserATSERVER:
		{
			p.SetState(326)
			p.AtServer()
		}

	case ApiParserParserATHANDLER, ApiParserParserCOMMENT, ApiParserParserLINE_COMMENT:
		{
			p.SetState(327)
			p.AtHandler()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(330)
		p.Route()
	}

	return localctx
}

// IAtDocContext is an interface to support dynamic dispatch.
type IAtDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsAtDocContext differentiates from other interfaces.
	IsAtDocContext()
}

type AtDocContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lp      antlr.Token
	comment ICommentSpecContext
	rp      antlr.Token
}

func NewEmptyAtDocContext() *AtDocContext {
	var p = new(AtDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_atDoc
	return p
}

func (*AtDocContext) IsAtDocContext() {}

func NewAtDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDocContext {
	var p = new(AtDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atDoc

	return p
}

func (s *AtDocContext) GetParser() antlr.Parser { return s.parser }

func (s *AtDocContext) GetLp() antlr.Token { return s.lp }

func (s *AtDocContext) GetRp() antlr.Token { return s.rp }

func (s *AtDocContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtDocContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtDocContext) GetComment() ICommentSpecContext { return s.comment }

func (s *AtDocContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *AtDocContext) ATDOC() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATDOC, 0)
}

func (s *AtDocContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *AtDocContext) AllKvLit() []IKvLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvLitContext)(nil)).Elem())
	var tst = make([]IKvLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvLitContext)
		}
	}

	return tst
}

func (s *AtDocContext) KvLit(i int) IKvLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtDocContext) CommentSpec() ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *AtDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtDocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtDoc() (localctx IAtDocContext) {
	localctx = NewAtDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ApiParserParserRULE_atDoc)
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
		p.SetState(332)
		p.Match(ApiParserParserATDOC)
	}
	{
		p.SetState(333)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*AtDocContext).lp = _m
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserCOMMENT, ApiParserParserLINE_COMMENT, ApiParserParserID:
		p.SetState(335)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(334)

				var _x = p.CommentSpec()

				localctx.(*AtDocContext).comment = _x
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ApiParserParserCOMMENT)|(1<<ApiParserParserLINE_COMMENT)|(1<<ApiParserParserID))) != 0) {
			{
				p.SetState(337)
				p.KvLit()
			}

			p.SetState(340)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ApiParserParserSTRING:
		{
			p.SetState(342)
			p.Match(ApiParserParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(345)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*AtDocContext).rp = _m
	}

	return localctx
}

// IAtHandlerContext is an interface to support dynamic dispatch.
type IAtHandlerContext interface {
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

	// IsAtHandlerContext differentiates from other interfaces.
	IsAtHandlerContext()
}

type AtHandlerContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	doc     ICommentSpecContext
	comment ICommentSpecContext
}

func NewEmptyAtHandlerContext() *AtHandlerContext {
	var p = new(AtHandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_atHandler
	return p
}

func (*AtHandlerContext) IsAtHandlerContext() {}

func NewAtHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtHandlerContext {
	var p = new(AtHandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atHandler

	return p
}

func (s *AtHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtHandlerContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *AtHandlerContext) GetComment() ICommentSpecContext { return s.comment }

func (s *AtHandlerContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *AtHandlerContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *AtHandlerContext) ATHANDLER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATHANDLER, 0)
}

func (s *AtHandlerContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *AtHandlerContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *AtHandlerContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *AtHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtHandler() (localctx IAtHandlerContext) {
	localctx = NewAtHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ApiParserParserRULE_atHandler)
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
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(347)

			var _x = p.CommentSpec()

			localctx.(*AtHandlerContext).doc = _x
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(353)
		p.Match(ApiParserParserATHANDLER)
	}
	{
		p.SetState(354)
		p.Match(ApiParserParserID)
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(355)

			var _x = p.CommentSpec()

			localctx.(*AtHandlerContext).comment = _x
		}

	}

	return localctx
}

// IRouteContext is an interface to support dynamic dispatch.
type IRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHttpMethod returns the httpMethod token.
	GetHttpMethod() antlr.Token

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// SetHttpMethod sets the httpMethod token.
	SetHttpMethod(antlr.Token)

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// GetDoc returns the doc rule contexts.
	GetDoc() ICommentSpecContext

	// GetRequest returns the request rule contexts.
	GetRequest() IBodyContext

	// GetResponse returns the response rule contexts.
	GetResponse() IBodyContext

	// GetComment returns the comment rule contexts.
	GetComment() ICommentSpecContext

	// SetDoc sets the doc rule contexts.
	SetDoc(ICommentSpecContext)

	// SetRequest sets the request rule contexts.
	SetRequest(IBodyContext)

	// SetResponse sets the response rule contexts.
	SetResponse(IBodyContext)

	// SetComment sets the comment rule contexts.
	SetComment(ICommentSpecContext)

	// IsRouteContext differentiates from other interfaces.
	IsRouteContext()
}

type RouteContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	doc         ICommentSpecContext
	httpMethod  antlr.Token
	request     IBodyContext
	returnToken antlr.Token
	response    IBodyContext
	comment     ICommentSpecContext
}

func NewEmptyRouteContext() *RouteContext {
	var p = new(RouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_route
	return p
}

func (*RouteContext) IsRouteContext() {}

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	var p = new(RouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_route

	return p
}

func (s *RouteContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteContext) GetHttpMethod() antlr.Token { return s.httpMethod }

func (s *RouteContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *RouteContext) SetHttpMethod(v antlr.Token) { s.httpMethod = v }

func (s *RouteContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *RouteContext) GetDoc() ICommentSpecContext { return s.doc }

func (s *RouteContext) GetRequest() IBodyContext { return s.request }

func (s *RouteContext) GetResponse() IBodyContext { return s.response }

func (s *RouteContext) GetComment() ICommentSpecContext { return s.comment }

func (s *RouteContext) SetDoc(v ICommentSpecContext) { s.doc = v }

func (s *RouteContext) SetRequest(v IBodyContext) { s.request = v }

func (s *RouteContext) SetResponse(v IBodyContext) { s.response = v }

func (s *RouteContext) SetComment(v ICommentSpecContext) { s.comment = v }

func (s *RouteContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *RouteContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *RouteContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *RouteContext) AllCommentSpec() []ICommentSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem())
	var tst = make([]ICommentSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentSpecContext)
		}
	}

	return tst
}

func (s *RouteContext) CommentSpec(i int) ICommentSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentSpecContext)
}

func (s *RouteContext) AllBody() []IBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBodyContext)(nil)).Elem())
	var tst = make([]IBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBodyContext)
		}
	}

	return tst
}

func (s *RouteContext) Body(i int) IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Route() (localctx IRouteContext) {
	localctx = NewRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ApiParserParserRULE_route)
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
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(358)

			var _x = p.CommentSpec()

			localctx.(*RouteContext).doc = _x
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	checkHttpMethod(p)
	{
		p.SetState(365)

		var _m = p.Match(ApiParserParserID)

		localctx.(*RouteContext).httpMethod = _m
	}
	{
		p.SetState(366)
		p.Path()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(367)

			var _x = p.Body()

			localctx.(*RouteContext).request = _x
		}

	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(370)

			var _m = p.Match(ApiParserParserID)

			localctx.(*RouteContext).returnToken = _m
		}

	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__1 {
		{
			p.SetState(373)

			var _x = p.Body()

			localctx.(*RouteContext).response = _x
		}

	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(376)

			var _x = p.CommentSpec()

			localctx.(*RouteContext).comment = _x
		}

	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) GetLp() antlr.Token { return s.lp }

func (s *BodyContext) GetRp() antlr.Token { return s.rp }

func (s *BodyContext) SetLp(v antlr.Token) { s.lp = v }

func (s *BodyContext) SetRp(v antlr.Token) { s.rp = v }

func (s *BodyContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
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

func (p *ApiParserParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ApiParserParserRULE_body)

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
		p.SetState(379)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*BodyContext).lp = _m
	}
	checkKeyword(p)
	{
		p.SetState(381)
		p.Match(ApiParserParserID)
	}
	{
		p.SetState(382)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*BodyContext).rp = _m
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
	p.EnterRule(localctx, 66, ApiParserParserRULE_kvLit)
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
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserCOMMENT || _la == ApiParserParserLINE_COMMENT {
		{
			p.SetState(384)

			var _x = p.CommentSpec()

			localctx.(*KvLitContext).doc = _x
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(390)

		var _m = p.Match(ApiParserParserID)

		localctx.(*KvLitContext).key = _m
	}
	checkKeyValue(p)
	{
		p.SetState(392)

		var _m = p.Match(ApiParserParserLINE_VALUE)

		localctx.(*KvLitContext).value = _m
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(393)

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
	p.EnterRule(localctx, 68, ApiParserParserRULE_commentSpec)
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
		p.SetState(396)
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
	p.RuleIndex = ApiParserParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
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

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ApiParserParserRULE_serviceName)
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
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(398)
			p.Match(ApiParserParserID)
		}
		p.SetState(400)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__9 {
			{
				p.SetState(399)
				p.Match(ApiParserParserT__9)
			}

		}

		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ApiParserParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *PathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
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

func (p *ApiParserParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ApiParserParserRULE_path)
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
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserT__10 || _la == ApiParserParserT__11 {
		p.SetState(418)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ApiParserParserT__10:
			{
				p.SetState(406)
				p.Match(ApiParserParserT__10)
			}

			{
				p.SetState(407)
				p.Match(ApiParserParserID)
			}
			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ApiParserParserT__9 {
				{
					p.SetState(408)
					p.Match(ApiParserParserT__9)
				}
				{
					p.SetState(409)
					p.Match(ApiParserParserID)
				}

			}

		case ApiParserParserT__11:
			{
				p.SetState(412)
				p.Match(ApiParserParserT__11)
			}

			{
				p.SetState(413)
				p.Match(ApiParserParserID)
			}
			p.SetState(416)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ApiParserParserT__9 {
				{
					p.SetState(414)
					p.Match(ApiParserParserT__9)
				}
				{
					p.SetState(415)
					p.Match(ApiParserParserID)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
