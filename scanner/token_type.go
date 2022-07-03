package scanner

type TokenType int

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
	return &token{}
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
