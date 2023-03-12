package scanner

import (
	"fmt"
	"github.com/praveensanap/glox-interpreter/errors"
	"strconv"
)

type scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func New(source string) *scanner {
	return &scanner{
		source: source,
		tokens: []Token{},
		line:   1,
	}
}

// Called once to scan all tokens in scanner.source
// Adds EOF token at the end

func (s *scanner) ScanTokens() []Token {
	// iterate through all characters in the source
	for !s.isEnd() {
		// start is advanced here. implies scanSingleTokenAndAdvance should process the whole lexeme
		s.start = s.current

		// process a token at a time and advance.
		s.scanSingleTokenAndAdvance()
	}
	t := NewToken(EOF, "", nil, s.line)
	s.tokens = append(s.tokens, t)
	for _, t := range s.tokens {
		fmt.Println(fmt.Sprintf("[line %d] [lexeme %s] [literal %v] [token %v]", t.GetLine(), t.GetLexeme(), t.GetLiteral(), t.GetTokenType()))
	}
	return s.tokens
}

// Job is to identify a token in the grammar
// Token identifies:
// 1. single character token (){},;+-*
// 2. two(one) character token  !, !=, =, ==, <, <=, >, >=
// 3. block or line comments.
// 4. ignores white space. advances line number for new line token.
// 5. string literals that are enclosed in double quotes.
// 6. number literals
// 7. keyword
// 8. identifier
// Errors on invalid tokens.
// each case advances the pointer to the start of next token.
func (s *scanner) scanSingleTokenAndAdvance() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
		break
	case ')':
		s.addToken(RIGHT_PAREN)
		break
	case '{':
		s.addToken(LEFT_BRACE)
		break
	case '}':
		s.addToken(RIGHT_BRACE)
		break
	case ',':
		s.addToken(COMMA)
		break
	case '.':
		s.addToken(DOT)
		break
	case '-':
		s.addToken(MINUS)
		break
	case '+':
		s.addToken(PLUS)
		break
	case ';':
		s.addToken(SEMICOLON)
		break
	case '*':
		s.addToken(STAR)
		break
		// two char matches
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	// comment?
	case '/':
		if s.match('/') {
			s.lineComment()
		} else if s.match('*') {
			s.blockComment()
		} else {
			s.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
		break
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.keywordOrIdentifier()
		} else {
			errors.Error(s.line, "Unexpected character.")
		}
		break
	}
}

// init tokens
func (s *scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}
func (s *scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	t := NewToken(tokenType, text, literal, s.line)
	s.tokens = append(s.tokens, t)
}

// source navigation. checks if at the end
func (s *scanner) isEnd() bool {
	return s.current >= len(s.source)
}

// source navigation. return the current character and move one step
func (s *scanner) advance() uint8 {
	c := s.source[s.current]
	s.current += 1
	return c
}

// advance if current matched with expected
func (s *scanner) match(expected uint8) bool {
	if s.isEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

// advance if current string matches with expected.
func (s *scanner) matchString(expected string) bool {
	if s.isEnd() {
		return false
	}

	if s.current+len(expected) > len(s.source) {
		return false
	}

	if s.source[s.current:s.current+len(expected)] != expected {
		return false
	}
	s.current += len(expected)
	return true
}

func (s *scanner) peek() uint8 {
	if s.isEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *scanner) peekNext() uint8 {
	if s.isEnd() && s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

// assumes "//" and find \n to terminate line comment
func (s *scanner) lineComment() {
	for s.peek() != '\n' && !s.isEnd() {
		s.advance()
	}
}

func (s *scanner) blockComment() {
	for {
		if s.isEnd() {
			errors.Error(s.line, "Unterminated block comment")
			return
		} else if s.matchString("/*") {
			s.blockComment()
		} else if s.peek() != '*' {
			if s.peek() == '\n' {
				s.line++
			}
			s.advance()
		} else if s.matchString("*/") {
			return
		}
	}
}

// given '"' finds the rest of the string
func (s *scanner) string() {
	for {
		if s.peek() != '"' && !s.isEnd() {
			if s.peek() == '\n' {
				s.line++
			}
			s.advance()
		} else {
			break
		}
	}
	if s.isEnd() {
		errors.Error(s.line, "Unterminated string")
		return
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(STRING, value)
}

// given a numeric char finds the whole number
// e.g. 1234, 12.34
func (s *scanner) number() {
	for {
		if isDigit(s.peek()) {
			s.advance()
		} else {
			break
		}
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// consume '.'
		s.advance()
	}
	for {
		if isDigit(s.peek()) {
			s.advance()
		} else {
			break
		}
	}
	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		errors.Error(s.line, "failed to parse into a number")
		return
	}

	s.addTokenWithLiteral(NUMBER, value)

}

func (s *scanner) keywordOrIdentifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if t, ok := KEYWORDS[text]; ok {
		s.addToken(t)
	} else {
		s.addToken(IDENTIFIER)
	}

}

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c > 'Z') || (c == '_')
}

func isAlphaNumeric(c uint8) bool {
	return isAlpha(c) || isDigit(c)
}
