package ast

import (
	"fmt"
)

type Printer struct {
}

func (p Printer) Print(expr Expr) {
	expr.accept(p)
}

func (p Printer) VisitBinaryExpr(expr BinaryExpr) {
	p.paranthesize(expr.Operator.GetLexeme(), expr.Left, expr.Right)
}

func (p Printer) VisitGroupingExpr(expr GroupingExpr) {
	p.paranthesize("group", expr.Expression)
}

func (p Printer) VisitLiteralExpr(expr LiteralExpr) {
	if expr.Value == nil {
		fmt.Print("nil")
	} else {
		fmt.Print(fmt.Sprintf("%+v", expr.Value))
	}
}

func (p Printer) VisitUnaryExpr(expr UnaryExpr) {
	p.paranthesize(expr.Operator.GetLexeme(), expr.Right)
}

func (p Printer) paranthesize(name string, expression ...Expr) {
	fmt.Print("(")
	fmt.Print(name)

	for _, expr := range expression {
		fmt.Print(" ")
		expr.accept(p)
	}
	fmt.Print(")")
}
