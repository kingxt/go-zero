// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/MetaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package meta // MetaParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 22, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 7, 2, 9, 10, 2, 12, 2, 14, 2, 12, 11,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 2, 2, 4,
	2, 4, 2, 2, 2, 21, 2, 6, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6, 10, 7, 3, 2,
	2, 7, 9, 5, 4, 3, 2, 8, 7, 3, 2, 2, 2, 9, 12, 3, 2, 2, 2, 10, 8, 3, 2,
	2, 2, 10, 11, 3, 2, 2, 2, 11, 13, 3, 2, 2, 2, 12, 10, 3, 2, 2, 2, 13, 14,
	7, 4, 2, 2, 14, 15, 7, 2, 2, 3, 15, 3, 3, 2, 2, 2, 16, 17, 7, 9, 2, 2,
	17, 19, 7, 5, 2, 2, 18, 20, 7, 9, 2, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3,
	2, 2, 2, 20, 5, 3, 2, 2, 2, 4, 10, 19,
}
var literalNames = []string{
	"", "'('", "')'", "':'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "COLON", "WS", "COMMENT", "LINE_COMMENT", "VALUE_LIT",
}

var ruleNames = []string{
	"kv", "kvLit",
}

type MetaParser struct {
	*antlr.BaseParser
}

// NewMetaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MetaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMetaParser(input antlr.TokenStream) *MetaParser {
	this := new(MetaParser)
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
	this.GrammarFileName = "MetaParser.g4"

	return this
}

// MetaParser tokens.
const (
	MetaParserEOF          = antlr.TokenEOF
	MetaParserLPAREN       = 1
	MetaParserRPAREN       = 2
	MetaParserCOLON        = 3
	MetaParserWS           = 4
	MetaParserCOMMENT      = 5
	MetaParserLINE_COMMENT = 6
	MetaParserVALUE_LIT    = 7
)

// MetaParser rules.
const (
	MetaParserRULE_kv    = 0
	MetaParserRULE_kvLit = 1
)

// IKvContext is an interface to support dynamic dispatch.
type IKvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKvContext differentiates from other interfaces.
	IsKvContext()
}

type KvContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKvContext() *KvContext {
	var p = new(KvContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_kv
	return p
}

func (*KvContext) IsKvContext() {}

func NewKvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvContext {
	var p = new(KvContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_kv

	return p
}

func (s *KvContext) GetParser() antlr.Parser { return s.parser }

func (s *KvContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MetaParserLPAREN, 0)
}

func (s *KvContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MetaParserRPAREN, 0)
}

func (s *KvContext) EOF() antlr.TerminalNode {
	return s.GetToken(MetaParserEOF, 0)
}

func (s *KvContext) AllKvLit() []IKvLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvLitContext)(nil)).Elem())
	var tst = make([]IKvLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvLitContext)
		}
	}

	return tst
}

func (s *KvContext) KvLit(i int) IKvLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *KvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MetaParserVisitor:
		return t.VisitKv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MetaParser) Kv() (localctx IKvContext) {
	localctx = NewKvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MetaParserRULE_kv)
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
		p.SetState(4)
		p.Match(MetaParserLPAREN)
	}
	p.SetState(8)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MetaParserVALUE_LIT {
		{
			p.SetState(5)
			p.KvLit()
		}

		p.SetState(10)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(11)
		p.Match(MetaParserRPAREN)
	}
	{
		p.SetState(12)
		p.Match(MetaParserEOF)
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

	// IsKvLitContext differentiates from other interfaces.
	IsKvLitContext()
}

type KvLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyKvLitContext() *KvLitContext {
	var p = new(KvLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_kvLit
	return p
}

func (*KvLitContext) IsKvLitContext() {}

func NewKvLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvLitContext {
	var p = new(KvLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_kvLit

	return p
}

func (s *KvLitContext) GetParser() antlr.Parser { return s.parser }

func (s *KvLitContext) GetKey() antlr.Token { return s.key }

func (s *KvLitContext) GetValue() antlr.Token { return s.value }

func (s *KvLitContext) SetKey(v antlr.Token) { s.key = v }

func (s *KvLitContext) SetValue(v antlr.Token) { s.value = v }

func (s *KvLitContext) COLON() antlr.TerminalNode {
	return s.GetToken(MetaParserCOLON, 0)
}

func (s *KvLitContext) AllVALUE_LIT() []antlr.TerminalNode {
	return s.GetTokens(MetaParserVALUE_LIT)
}

func (s *KvLitContext) VALUE_LIT(i int) antlr.TerminalNode {
	return s.GetToken(MetaParserVALUE_LIT, i)
}

func (s *KvLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MetaParserVisitor:
		return t.VisitKvLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MetaParser) KvLit() (localctx IKvLitContext) {
	localctx = NewKvLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MetaParserRULE_kvLit)

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
		p.SetState(14)

		var _m = p.Match(MetaParserVALUE_LIT)

		localctx.(*KvLitContext).key = _m
	}
	{
		p.SetState(15)
		p.Match(MetaParserCOLON)
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(16)

			var _m = p.Match(MetaParserVALUE_LIT)

			localctx.(*KvLitContext).value = _m
		}

	}

	return localctx
}
