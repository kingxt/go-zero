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
typeLit:        doc=commentSpec?{match(p,"type")}typeToken=ID  typeLitBody;
typeBlock:      {match(p,"type")}typeToken=ID lp='(' typeBlockBody+ rp=')';
typeLitBody:    typeStruct|typeAlias;
typeBlockBody: typeBlockStruct|typeBlockAlias;
typeStruct:     {checkKeyword(p)}structName=ID structToken=ID? lbrace='{'  comment=commentSpec? field+ rbrace='}';
typeAlias:      {checkKeyword(p)}alias=ID assign='='? dataType comment=commentSpec?;
typeBlockStruct:doc=commentSpec?  {checkKeyword(p)}structName=ID structToken=ID? lbrace='{'  comment=commentSpec? field+ rbrace='}';
typeBlockAlias: doc=commentSpec? {checkKeyword(p)}alias=ID assign='='? dataType comment=commentSpec?;
field:          {isNormal(p)}? normalField|anonymousFiled ;
normalField:    doc=commentSpec? {checkKeyword(p)}fieldName=ID dataType tag=RAW_STRING? comment=commentSpec?;
anonymousFiled: doc=commentSpec? star='*'? ID comment=commentSpec?;
dataType:       {isInterface(p)}ID
                |mapType
                |arrayType
                |inter='interface{}'
                |time='time.Time'
                |pointerType
                |typeStruct
                ;
pointerType:        star='*' {checkKeyword(p)}ID;
mapType:            {match(p,"map")}mapToken=ID lbrack='[' {checkKey(p)}key=ID rbrack=']' value=dataType;
arrayType:          lbrack='[' rbrack=']' dataType;

// kv
kvLit:          doc=commentSpec? key=ID {checkKeyValue(p)}value=LINE_VALUE comment=commentSpec?;

// comment
commentSpec:        COMMENT|LINE_COMMENT;