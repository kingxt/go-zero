// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/AnonymousParser.g4 by ANTLR 4.9. DO NOT EDIT.

package anonymous // AnonymousParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 19, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 5, 2, 8, 10, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2,
	3, 2, 5, 2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 2, 2, 4, 2, 4, 2, 3, 3, 2, 9,
	10, 2, 19, 2, 7, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6, 8, 5, 4, 3, 2, 7, 6,
	3, 2, 2, 2, 7, 8, 3, 2, 2, 2, 8, 10, 3, 2, 2, 2, 9, 11, 7, 3, 2, 2, 10,
	9, 3, 2, 2, 2, 10, 11, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 14, 7, 14, 2,
	2, 13, 15, 5, 4, 3, 2, 14, 13, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 3, 3,
	2, 2, 2, 16, 17, 9, 2, 2, 2, 17, 5, 3, 2, 2, 2, 5, 7, 10, 14,
}
var literalNames = []string{
	"", "'*'", "'@doc'", "'@handler'", "'interface{}'", "'@server'",
}
var symbolicNames = []string{
	"", "", "ATDOC", "ATHANDLER", "INTERFACE", "ATSERVER", "WS", "COMMENT",
	"LINE_COMMENT", "STRING", "RAW_STRING", "LINE_VALUE", "ID",
}

var ruleNames = []string{
	"anonymousFiled", "commentSpec",
}

type AnonymousParserParser struct {
	*antlr.BaseParser
}

// NewAnonymousParserParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *AnonymousParserParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAnonymousParserParser(input antlr.TokenStream) *AnonymousParserParser {
	this := new(AnonymousParserParser)
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
	this.GrammarFileName = "AnonymousParser.g4"

	return this
}

// AnonymousParserParser tokens.
const (
	AnonymousParserParserEOF          = antlr.TokenEOF
	AnonymousParserParserT__0         = 1
	AnonymousParserParserATDOC        = 2
	AnonymousParserParserATHANDLER    = 3
	AnonymousParserParserINTERFACE    = 4
	AnonymousParserParserATSERVER     = 5
	AnonymousParserParserWS           = 6
	AnonymousParserParserCOMMENT      = 7
	AnonymousParserParserLINE_COMMENT = 8
	AnonymousParserParserSTRING       = 9
	AnonymousParserParserRAW_STRING   = 10
	AnonymousParserParserLINE_VALUE   = 11
	AnonymousParserParserID           = 12
)

// AnonymousParserParser rules.
const (
	AnonymousParserParserRULE_anonymousFiled = 0
	AnonymousParserParserRULE_commentSpec    = 1
)

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
	p.RuleIndex = AnonymousParserParserRULE_anonymousFiled
	return p
}

func (*AnonymousFiledContext) IsAnonymousFiledContext() {}

func NewAnonymousFiledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFiledContext {
	var p = new(AnonymousFiledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnonymousParserParserRULE_anonymousFiled

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
	return s.GetToken(AnonymousParserParserID, 0)
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
	case AnonymousParserVisitor:
		return t.VisitAnonymousFiled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnonymousParserParser) AnonymousFiled() (localctx IAnonymousFiledContext) {
	localctx = NewAnonymousFiledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnonymousParserParserRULE_anonymousFiled)
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
	p.SetState(5)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AnonymousParserParserCOMMENT || _la == AnonymousParserParserLINE_COMMENT {
		{
			p.SetState(4)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).doc = _x
		}

	}
	p.SetState(8)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AnonymousParserParserT__0 {
		{
			p.SetState(7)

			var _m = p.Match(AnonymousParserParserT__0)

			localctx.(*AnonymousFiledContext).star = _m
		}

	}
	{
		p.SetState(10)
		p.Match(AnonymousParserParserID)
	}
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AnonymousParserParserCOMMENT || _la == AnonymousParserParserLINE_COMMENT {
		{
			p.SetState(11)

			var _x = p.CommentSpec()

			localctx.(*AnonymousFiledContext).comment = _x
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
	p.RuleIndex = AnonymousParserParserRULE_commentSpec
	return p
}

func (*CommentSpecContext) IsCommentSpecContext() {}

func NewCommentSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentSpecContext {
	var p = new(CommentSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnonymousParserParserRULE_commentSpec

	return p
}

func (s *CommentSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentSpecContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(AnonymousParserParserCOMMENT, 0)
}

func (s *CommentSpecContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(AnonymousParserParserLINE_COMMENT, 0)
}

func (s *CommentSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnonymousParserVisitor:
		return t.VisitCommentSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnonymousParserParser) CommentSpec() (localctx ICommentSpecContext) {
	localctx = NewCommentSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnonymousParserParserRULE_commentSpec)
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
		p.SetState(14)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AnonymousParserParserCOMMENT || _la == AnonymousParserParserLINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
