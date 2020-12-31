grammar AnonymousParser;

import ApiLexer;

anonymousFiled: doc=commentSpec? star='*'? ID comment=commentSpec?;

// comment
commentSpec:        COMMENT|LINE_COMMENT;