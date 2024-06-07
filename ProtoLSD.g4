grammar ProtoLSD;

/* Lexer rules */

// Keywords
IMPORT: 'import';
PRIVATE: 'private';
SERVICE: 'service';
RPC: 'rpc';
ENUM: 'enum';
MESSAGE: 'message';
RETURNS: 'returns';
OPTION: 'option';
OPTIONAL: 'optional';
REQUIRED: 'required';
REPEATED: 'repeated';
PACKAGE: 'package';

// Preprocessor Directives

// Comments
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Identifiers
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
PREPROCESSOR_NAME: [a-zA-Z_][a-zA-Z0-9_-]*;

// Literals
STRING: '"' (ESC_SEQ | ~('\\' | '"'))* '"';
NUMBER: [0-9]+;
MAP: 'map<' IDENTIFIER ',' IDENTIFIER '>';

// Operators and delimiters
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
SEMI: ';';
COMMA: ',';
EQUAL: '=';
DOT: '.';
HASH: '#';
QUESTION: '?';
ARRAY: '[]';
OBJECT: '{}';
BANG: '!';
PLUS: '+';
ANNOTATION: '@';

// Whitespace
WS: [ \t\r\n]+ -> skip;

/* Parser rules */

semi: SEMI?;

protoLSD: statement*;

statement
    : importStatement
    | serviceDefinition
    | enumDefinition
    | messageDefinition
    | preprocessorDirective
    | optionStatement
    | packageStatement
    ;

importStatement
    : (PRIVATE? IMPORT STRING BANG? semi)
    ;

serviceDefinition
    : SERVICE IDENTIFIER LBRACE serviceBody RBRACE
    ;

serviceBody
    : rpcDefinition*
    ;

rpcDefinition
    : RPC IDENTIFIER rpcInput? rpcReturn? (SEMI | OBJECT)?
    ;

rpcParameters
    : LPAREN (rpcParameter (COMMA rpcParameter)*)? RPAREN
    ;

rpcParameter
    : dataType PLUS? rpcParameterModifier?
    | dataType? IDENTIFIER
    ;

rpcParameterModifier
    : IDENTIFIER
    | '*' NUMBER
    ;

rpcReturn
    : RETURNS rpcMethodMessage
    ;

rpcInput
    : rpcMethodMessage
    ;

rpcMethodMessage
    : OBJECT
    | rpcSingleMessage
    | rpcParameters
    ;

rpcSingleMessage
    : LPAREN rpcMessageValue RPAREN
    | rpcMessageValue
    ;

rpcMessageValue
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

enumDefinition
    : annotationDirective* ENUM IDENTIFIER LBRACE enumBody RBRACE
    ;

enumBody
    : enumField (semi enumField)*
    ;

enumField
    : IDENTIFIER (EQUAL NUMBER)?
    ;

messageDefinition
    : MESSAGE IDENTIFIER LBRACE messageBody RBRACE
    ;

messageBody
    : messageBodyStatement*
    ;

messageBodyStatement
    : messageField semi
    | messageDefinition
    ;

messageField
    : optionalModifier? dataType IDENTIFIER (EQUAL NUMBER)?
    ;

preprocessorParameters
    : LPAREN (IDENTIFIER (COMMA IDENTIFIER)*)? RPAREN
    ;

preprocessorDirective
    : HASH PREPROCESSOR_NAME preprocessorParameters?
    ;

annotationDirective
    : ANNOTATION IDENTIFIER
    ;

optionalModifier
    : OPTIONAL
    | REQUIRED
    ;

optionStatement
    : OPTION IDENTIFIER EQUAL STRING semi
    ;

packageStatement
    : PACKAGE IDENTIFIER (DOT IDENTIFIER)* semi
    ;

dataType
    : dataTypeText ARRAY? QUESTION?
    ;

dataTypeText
    : 'int' 
    | 'uint'
    | 'long'
    | 'ulong'
    | MAP
    | IDENTIFIER
    | rpcMessageValue
    ;

/* Fragments for Lexer rules */

fragment ESC_SEQ: '\\' ('b' | 't' | 'n' | 'f' | 'r' | '"' | '\'' | '\\');
