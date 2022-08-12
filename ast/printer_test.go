package ast

import (
	"github.com/praveensanap/glox-interpreter/scanner"
	"testing"
)

func TestBasic(t *testing.T) {
	expression := BinaryExpr{
		Left: UnaryExpr{
			scanner.NewToken(scanner.MINUS, "-", nil, 1),
			LiteralExpr{Value: 123},
		},
		Operator: scanner.NewToken(scanner.SLASH, "*", nil, 1),
		Right: GroupingExpr{
			Expression: LiteralExpr{Value: 45.67},
		},
	}

	p := Printer{}
	p.Print(expression)
}
