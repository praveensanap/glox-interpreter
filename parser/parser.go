package parser

import (
	"fmt"
	"github.com/praveensanap/glox-interpreter/ast"
	"github.com/praveensanap/glox-interpreter/errors"
	"github.com/praveensanap/glox-interpreter/scanner"
)

// []Token -> SyntaxTree

type Parse interface {
	Parse() (ast.Expr, error)
}

type parser struct {
	tokens  []scanner.Token
	current int
}

func NewParser(tokens []scanner.Token) Parse {
	p := parser{
		tokens:  tokens,
		current: 0,
	}
	return &p
}

func (p *parser) Parse() (ast.Expr, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *parser) expression() (ast.Expr, error) {
	return p.equality()
}

func (p *parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *parser) comparison() (ast.Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.LESS, scanner.LESS_EQUAL, scanner.GREATER, scanner.GREATER_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *parser) term() (ast.Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.MINUS, scanner.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *parser) factor() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.SLASH, scanner.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *parser) unary() (ast.Expr, error) {
	if p.match(scanner.BANG, scanner.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return ast.UnaryExpr{
			Operator: operator,
			Right:    right,
		}, nil
	}
	return p.primary()
}

func (p *parser) primary() (ast.Expr, error) {
	if p.match(scanner.FALSE) {
		return ast.LiteralExpr{Value: false}, nil
	}
	if p.match(scanner.TRUE) {
		return ast.LiteralExpr{Value: true}, nil
	}
	if p.match(scanner.NIL) {
		return ast.LiteralExpr{Value: nil}, nil
	}
	if p.match(scanner.NUMBER, scanner.STRING) {
		return ast.LiteralExpr{Value: p.previous().GetLiteral()}, nil
	}
	if p.match(scanner.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(scanner.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return ast.GroupingExpr{Expression: expr}, nil
	}
	return nil, errors.New("Expected expression.")
}

func (p *parser) match(tokens ...scanner.TokenType) bool {
	for _, token := range tokens {
		if p.check(token) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) isAtEnd() bool {
	return p.peek().GetTokenType() == scanner.EOF
}

func (p *parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *parser) check(tokenType scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().GetTokenType() == tokenType
}

func (p *parser) consume(tokenType scanner.TokenType, errMsg string) (scanner.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	p.error(p.peek(), errMsg)
	return nil, errors.New("parse err")
}

func (p *parser) error(token scanner.Token, errMsg string) {
	if token.GetTokenType() == scanner.EOF {
		errors.Error(token.GetLine(), fmt.Sprintf(" at end %s ", errMsg))
	} else {
		errors.Error(token.GetLine(), fmt.Sprintf(" at  %s '%s'", token.GetLexeme(), errMsg))
	}
}

func (p *parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().GetTokenType() == scanner.SEMICOLON {
			return
		}
		switch p.peek().GetTokenType() {
		case scanner.CLASS:
		case scanner.FUN:
		case scanner.VAR:
		case scanner.FOR:
		case scanner.IF:
		case scanner.WHILE:
		case scanner.PRINT:
		case scanner.RETURN:
			return
		}
		p.advance()
	}

}
