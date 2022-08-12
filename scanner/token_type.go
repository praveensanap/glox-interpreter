package scanner

type TokenType int

var KEYWORDS = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one to twp char tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	//literals
	IDENTIFIER
	STRING
	NUMBER

	//keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type token struct {
	type_   TokenType
	lexeme  string
	literal interface{}
	line    int
}

func NewToken(type_ TokenType, lexeme string, literal interface{}, line int) Token {
	return &token{
		type_:   type_,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t token) GetTokenType() TokenType {
	return t.type_
}

func (t token) GetLexeme() string {
	return t.lexeme
}

func (t token) GetLiteral() interface{} {
	return t.literal
}

func (t token) GetLine() int {
	return t.line
}
