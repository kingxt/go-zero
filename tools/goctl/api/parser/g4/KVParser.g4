parser grammar KVParser;

options {
    tokenVocab=KVLexer;
}

kv:        '(' kvLit* ')' EOF;
kvLit:     key=VALUE_LIT COLON value=(VALUE_LIT|STRING_LIT)?;