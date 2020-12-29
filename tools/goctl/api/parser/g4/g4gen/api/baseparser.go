package api

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const (
	versionRegex     = `(?m)"v[1-9][0-9]*"`
	importValueRegex = `(?m)"(/?[a-zA-Z][a-zA-Z0-9_]*)+.api"`
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
	return v == text
}

func expecting(expecting, found string) string {
	return fmt.Sprintf(`expecting '%s', found input '%s'`, expecting, found)
}

func mismatched(expecting, found string) string {
	return fmt.Sprintf(`mismatched '%s', found input '%s'`, expecting, found)
}
