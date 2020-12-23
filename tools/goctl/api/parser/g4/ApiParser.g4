parser grammar ApiParser;

options {
    tokenVocab=ApiLexer;
}

api:        syntaxLit? body* EOF;

body:       importSpec
            |infoBlock
            |typeBlock
            |serviceBlock
            ;

syntaxLit:      syntaxToken=ID ASSIGN version=SYNTAX_VERSION;
importSpec:     importLit|importLitGroup;
importLit:      importToken=ID importPath=IMPORT_PATH;
importLitGroup:     importToken=ID '(' (importPath=IMPORT_PATH)* ')';

infoBlock: infoToken=ID '(' kvLit* ')';

typeBlock:      typeLit|typeGroup;
typeLit:        typeToken=ID typeSpec;
typeGroup:      typeToken=ID '(' typeSpec* ')';
typeSpec:       typeAlias|typeStruct;
typeAlias:      alias=ID '='? dataType;
typeStruct:     name=ID structToken=ID? '{' typeField* '}';
typeField:       name=ID filed?;
filed:      (dataType|innerStruct) tag=RAW_STRING?;
innerStruct:        structToken=ID? '{' typeField* '}';
dataType:       pointer
                |mapType
                |arrayType
                |INTERFACE
                ;
mapType:        mapToken=ID '[' key=ID ']' value=dataType;
arrayType:      '['']'lit=dataType;
pointer:        STAR* ID;

serviceBlock:       serverMeta? serviceBody;
serverMeta:     ATSERVER '(' annotation* ')';
annotation: key=ID COLON value=annotationKeyValue?;
annotationKeyValue:        (ID ('/' ID)?)+;
serviceBody:        serviceToken=ID serviceName '{' routes=serviceRoute* '}';
serviceName:        ID ('-' ID)?;
serviceRoute:       routeDoc? (serverMeta|routeHandler) routePath ;
routeDoc:       doc|lineDoc;
doc:        ATDOC '(' summaryToken=ID COLON STRING_LIT? ')';
lineDoc:        ATDOC STRING_LIT;
routeHandler:       ATHANDLER ID;
routePath:      httpMethodToken=ID path request? reply?;
path:      ('/' ':'? ID (('?'|'&'|'=') ID)?)+;
request:       '(' ID ')';
reply:      returnToken=ID '(' obj=ID ')';
kvLit:      key=ID COLON value=STRING_LIT?;