package scanner

import (
	"fmt"
	"github.com/praveensanap/glox-interpreter/errors"
)

type scanner struct {
	source  string
	tokens  []*Token
	start   int
	current int
	line    int
}

func New(source string) *scanner {
	return &scanner{
		source: source,
		tokens: []*Token{},
		line:   1,
	}
}

func (s *scanner) ScanTokens() {
	for !s.isEnd() {
		s.start = s.current
		s.scanToken()
	}
	t := NewToken(EOF, "", nil, s.line)
	s.tokens = append(s.tokens, &t)
	fmt.Println(s.tokens)
}

func (s *scanner) isEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) advance() string {
	c := s.source[s.current]
	s.current += 1
	return string(c)

}

func (s *scanner) scanToken() {
	c := s.advance()

	switch c {
	case "(":
		s.addToken(LEFT_PAREN)
		break
	case ")":
		s.addToken(RIGHT_PAREN)
		break
	case "{":
		s.addToken(LEFT_BRACE)
		break
	case "}":
		s.addToken(RIGHT_BRACE)
		break
	case ",":
		s.addToken(COMMA)
		break
	case ".":
		s.addToken(DOT)
		break
	case "-":
		s.addToken(MINUS)
		break
	case "+":
		s.addToken(PLUS)
		break
	case ";":
		s.addToken(SEMICOLON)
		break
	case "*":
		s.addToken(STAR)
		break
	default:
		errors.Error(s.line, "Unexpected character.")
		break
	}

}

func (s *scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}
func (s *scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	t := NewToken(tokenType, text, nil, 0)
	s.tokens = append(s.tokens, &t)
}
