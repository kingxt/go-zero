parser grammar MetaParser;

options {
    tokenVocab=MetaLexer;
}

kv:        '(' kvLit* ')' EOF;
kvLit:     key=VALUE_LIT COLON value=VALUE_LIT?;