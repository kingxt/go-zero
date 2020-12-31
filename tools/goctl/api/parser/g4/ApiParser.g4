grammar ApiParser;

import ApiLexer;

api:            spec*;
spec:           syntaxLit
                |importSpec
                |infoSpec
                |typeSpec
                |commentSpec;

// syntax
syntaxLit:      doc=commentSpec? {match(p,"syntax")}syntaxToken=ID assign='=' {checkVersion(p)}version=STRING comment=commentSpec?;

// import
importSpec:     importLit|importBlock;
importLit:      doc=commentSpec? {match(p,"import")}importToken=ID importValue comment=commentSpec? ;
importBlock:    {match(p,"import")}importToken=ID '(' comment=commentSpec? importBlockValue+ ')';
importBlockValue:   doc=commentSpec? importValue comment=commentSpec?;
importValue:    {checkImportValue(p)}STRING;

// info
infoSpec:       doc=commentSpec? {match(p,"info")}infoToken=ID lp='(' comment=commentSpec? kvLit+ rp=')';

// type
typeSpec:       typeLit
                |typeBlock;

// eg: type (...)
typeLit:        doc=commentSpec?{match(p,"type")}typeToken=ID  typeBody;
typeBlock:      doc=commentSpec? {match(p,"type")}typeToken=ID lp='(' typeBody+ rp=')';
typeBody:       typeAlias|typeStruct;
typeStruct:     {checkFieldName(p)}structName=ID {match(p,"struct")}structToken=ID? lbrace='{'  comment=commentSpec? field+ rbrace='}';
typeAlias:      {checkFieldName(p)}alias=ID assign='='? dataType comment=commentSpec?;
field:          ({isAnonymous(p)}? anonymousFiled) | normalField;
normalField:    doc=commentSpec? fieldName=ID dataType {checkTag(p)}tag=RAW_STRING? comment=commentSpec?;
anonymousFiled: doc=commentSpec? star='*'? ID comment=commentSpec?;
dataType:       {isInterface(p)}ID
                |mapType
                |arrayType
                |inter='interface{}'
                |time='time.Time'
                |pointerType
                |typeStruct
                ;
pointerType:        star='*' ID;
mapType:            {match(p,"map")}mapToken=ID lbrack='[' key=ID rbrack=']' value=dataType;
arrayType:          lbrack='[' rbrack=']' dataType;

// kv
kvLit:          doc=commentSpec? key=ID {checkKeyValue(p)}value=LINE_VALUE comment=commentSpec?;

// comment
commentSpec:        COMMENT|LINE_COMMENT;