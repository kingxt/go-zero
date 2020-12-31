package api

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/anonymous"
)

const (
	versionRegex     = `(?m)"v[1-9][0-9]*"`
	importValueRegex = `(?m)"(/?[a-zA-Z][a-zA-Z0-9_]*)+.api"`
	tagRegex         = `(?m)\x60[a-z]+:".+"\x60`
)

func match(p *ApiParserParser, text string) {
	v := getCurrentTokenText(p)
	if v != text {
		notifyErrorListeners(p, expecting(text, v))
	}
}

func checkVersion(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if !matchRegex(v, versionRegex) {
		notifyErrorListeners(p, mismatched("version", v))
	}
}

func checkImportValue(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if !matchRegex(v, importValueRegex) {
		notifyErrorListeners(p, mismatched("import value", v))
	}
}

func checkKeyValue(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if !strings.HasPrefix(v, ":") {
		notifyErrorListeners(p, mismatched(":", v))
	}

	v = strings.TrimPrefix(v, ":")
	v = strings.TrimFunc(v, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	setCurrentTokenText(p, v)
}

func checkFieldName(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if IsGolangKeyWord(v) {
		notifyErrorListeners(p, fmt.Sprintf("expecting ID, found golang keyword: '%s'", v))
	}
}

func IsGolangKeyWord(text string) bool {
	switch text {
	case "var", "const", "package", "import", "func", "return",
		"defer", "go", "select", "interface", "struct", "break", "case",
		"continue", "for", "fallthrough", "else", "if", "switch", "goto",
		"default", "chan", "type", "map", "range":
		return true
	default:
		return false
	}
}

func isAnonymous(p *ApiParserParser) bool {
	text := p.GetTokenStream().GetAllText()
	defaultParser := NewBaseParser()
	v, err := defaultParser.Accept(func(ap *anonymous.AnonymousParserParser) interface{} {
		iAnonymousFiledContext := ap.AnonymousFiled()
		if iAnonymousFiledContext != nil {
			ctx := iAnonymousFiledContext.(*anonymous.AnonymousFiledContext)
			currentText := ctx.GetText()
			if IsGolangKeyWord(ctx.ID().GetText()) {
				panic(fmt.Errorf("expecting ID, found golang keyword: '%s'", ctx.ID().GetText()))
			}
			return currentText
		}
		return ""
	}, text)
	if err != nil {
		notifyErrorListeners(p, err.Error())
	}
	str := v.(string)
	replace := func(s string) string {
		v := strings.ReplaceAll(s, " ", "")
		v = strings.ReplaceAll(v, "\r", "")
		v = strings.ReplaceAll(v, "\n", "")
		v = strings.ReplaceAll(v, "\t", "")
		return v
	}
	return replace(text) == replace(str)
}

func checkTag(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if v == "" || v == "<EOF>" {
		return
	}
	if !matchRegex(v, tagRegex) {
		notifyErrorListeners(p, mismatched("key-value tag", v))
	}
}

func isInterface(p *ApiParserParser) {
	v := getCurrentTokenText(p)
	if IsGolangKeyWord(v) {
		notifyErrorListeners(p, fmt.Sprintf("expecting ID, found golang keyword: '%s'", v))
	}
}

func getCurrentTokenText(p *ApiParserParser) string {
	token := p.GetCurrentToken()
	if token == nil {
		return ""
	}

	return token.GetText()
}

func setCurrentTokenText(p *ApiParserParser, text string) {
	token := p.GetCurrentToken()
	if token == nil {
		return
	}

	token.SetText(text)
}

func notifyErrorListeners(p *ApiParserParser, msg string) {
	p.NotifyErrorListeners(msg, nil, nil)
}

func matchRegex(text, str string) bool {
	re := regexp.MustCompile(str)
	v := re.FindString(text)
	text = strings.TrimFunc(text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	return v == text
}

func expecting(expecting, found string) string {
	return fmt.Sprintf(`expecting '%s', found input '%s'`, expecting, found)
}

func mismatched(expecting, found string) string {
	return fmt.Sprintf(`mismatched '%s', found input '%s'`, expecting, found)
}

type BaseParser struct {
	*anonymous.AnonymousParserParser
	antlr.DefaultErrorListener
}

func NewBaseParser() *BaseParser {
	return &BaseParser{}
}

func (p *BaseParser) Accept(fn func(p *anonymous.AnonymousParserParser) interface{}, content string) (v interface{}, err error) {
	defer func() {
		p := recover()
		if p != nil {
			switch e := p.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("%+v", p)
			}
		}
	}()

	inputStream := antlr.NewInputStream(content)
	lexer := anonymous.NewAnonymousParserLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	anonymousParser := anonymous.NewAnonymousParserParser(tokens)
	anonymousParser.RemoveErrorListeners()
	anonymousParser.AddErrorListener(p)
	v = fn(anonymousParser)
	return
}

func (p *BaseParser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	str := fmt.Sprintf(`line %d:%d  %s`, line, column, msg)
	panic(str)
}
