package scanner

type Scanner interface {

}


type Token interface {
	GetTokenType() TokenType
	GetLexeme() string
	GetLiteral() interface{}
	GetLine() int
}