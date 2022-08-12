package ast

import (
	"fmt"
)

type Printer struct {
}

func (p Printer) print(expr Expr) {
	expr.accept(p)
}

func (p Printer) VisitBinaryExpr(expr BinaryExpr) {
	p.paranthesize(expr.operator.GetLexeme(), expr.left, expr.right)
}

func (p Printer) VisitGroupingExpr(expr GroupingExpr) {
	p.paranthesize("group", expr.expression)
}

func (p Printer) VisitLiteralExpr(expr LiteralExpr) {
	if expr.value == nil {
		fmt.Print("nil")
	} else {
		fmt.Print(fmt.Sprintf("%+v", expr.value))
	}
}

func (p Printer) VisitUnaryExpr(expr UnaryExpr) {
	p.paranthesize(expr.operator.GetLexeme(), expr.right)
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
