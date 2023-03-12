package ast

import (
	"bytes"
	"fmt"
)

type Printer struct {
}

func (p Printer) Print(expr Expr) string {
	result := expr.Accept(p)
	return result.(string)
}

func (p Printer) VisitBinaryExpr(expr BinaryExpr) interface{} {
	return p.paranthesize(expr.Operator.GetLexeme(), expr.Left, expr.Right)
}

func (p Printer) VisitGroupingExpr(expr GroupingExpr) interface{} {
	return p.paranthesize("group", expr.Expression)
}

func (p Printer) VisitLiteralExpr(expr LiteralExpr) interface{} {
	var buffer bytes.Buffer
	if expr.Value == nil {
		buffer.WriteString("nil")
	} else {
		buffer.WriteString(fmt.Sprintf("%+v", expr.Value))
	}
	return buffer.String()
}

func (p Printer) VisitUnaryExpr(expr UnaryExpr) interface{} {
	return p.paranthesize(expr.Operator.GetLexeme(), expr.Right)
}

func (p Printer) paranthesize(name string, expression ...Expr) string {
	var buffer bytes.Buffer

	buffer.WriteString("(")
	buffer.WriteString(name)

	for _, expr := range expression {
		buffer.WriteString(" ")
		a := expr.Accept(p)
		buffer.WriteString(a.(string))
	}
	buffer.WriteString(")")
	return buffer.String()
}
