grammar ApiParser;

import ApiLexer;

api:            spec*;
spec:           syntaxLit
                |importSpec
                |infoSpec
                |commentSpec;

// syntax
syntaxLit:      doc=commentSpec? {match(p,"syntax")}syntaxToken=ID assign='=' {checkVersion(p)}version=STRING comment=commentSpec?;

// import
importSpec:     importLit|importBlock;
importLit:      doc=commentSpec? {match(p,"import")}importToken=ID importValue comment=commentSpec? ;
importBlock:    {match(p,"import")}importToken=ID '(' importBlockValue+ ')';
importBlockValue:   doc=commentSpec? importValue comment=commentSpec?;
importValue:    {checkImportValue(p)}STRING;

// info
infoSpec:       doc=commentSpec? {match(p,"info")}infoToken=ID lp='(' comment=commentSpec? kvLit+ rp=')';

// kv
kvLit:          doc=commentSpec? key=ID {checkKeyValue(p)}value=LINE_VALUE comment=commentSpec?;
// comment
commentSpec:        COMMENT|LINE_COMMENT;