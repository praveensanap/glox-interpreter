package ast

import (
	"bytes"
	"fmt"
)

type Printer struct {
	Errors []error
}

func (p Printer) Print(expr Expr) string {
	result, err := expr.Accept(p)

	if err != nil {
		p.Errors = append(p.Errors, err)
		return ""
	}
	return result.(string)
}

func (p Printer) VisitBinaryExpr(expr BinaryExpr) (interface{}, error) {
	return p.paranthesize(expr.Operator.GetLexeme(), expr.Left, expr.Right)
}

func (p Printer) VisitGroupingExpr(expr GroupingExpr) (interface{}, error) {
	return p.paranthesize("group", expr.Expression)
}

func (p Printer) VisitLiteralExpr(expr LiteralExpr) (interface{}, error) {
	var buffer bytes.Buffer
	if expr.Value == nil {
		buffer.WriteString("nil")
	} else {
		buffer.WriteString(fmt.Sprintf("%+v", expr.Value))
	}
	return buffer.String(), nil
}

func (p Printer) VisitUnaryExpr(expr UnaryExpr) (interface{}, error) {
	return p.paranthesize(expr.Operator.GetLexeme(), expr.Right)
}

func (p Printer) paranthesize(name string, expression ...Expr) (string, error) {
	var buffer bytes.Buffer

	buffer.WriteString("(")
	buffer.WriteString(name)

	for _, expr := range expression {
		buffer.WriteString(" ")
		a, err := expr.Accept(p)
		if err != nil {
			p.Errors = append(p.Errors, err)
			//return "" continue
		}
		buffer.WriteString(a.(string))
	}
	buffer.WriteString(")")
	return buffer.String(), nil
}
