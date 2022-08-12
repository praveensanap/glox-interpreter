package ast

import (
	"github.com/praveensanap/glox-interpreter/scanner"
	"testing"
)

func TestBasic(t *testing.T) {
	expression := BinaryExpr{
		left: UnaryExpr{
			scanner.NewToken(scanner.MINUS, "-", nil, 1),
			LiteralExpr{value: 123},
		},
		operator: scanner.NewToken(scanner.SLASH, "*", nil, 1),
		right: GroupingExpr{
			expression: LiteralExpr{value: 45.67},
		},
	}

	p := Printer{}
	p.print(expression)
}
