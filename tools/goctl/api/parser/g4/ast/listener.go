package ast

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type (
	ErrorListener struct {
		*antlr.DefaultErrorListener
		callback ErrCallback
	}

	ErrCallback func(err error)
)

func NewErrorListener(callback ErrCallback) *ErrorListener {
	return &ErrorListener{
		callback: callback,
	}
}

func (listener *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	lineHeader := "line " + strconv.Itoa(line) + ":" + strconv.Itoa(column)
	if listener.callback != nil {
		listener.callback(fmt.Errorf("%s, %s", lineHeader, msg))
		return
	}
	errString := fmt.Sprintf(lineHeader + ", " + msg)
	panic(errors.New(errString))
}
