package ast

import "github.com/praveensanap/glox-interpreter/scanner"

type Expr interface {
	Accept(visitor ExprVisitor) (interface{}, error)
}

type ExprVisitor interface {
	VisitBinaryExpr(BinaryExpr) (interface{}, error)
	VisitGroupingExpr(GroupingExpr) (interface{}, error)
	VisitLiteralExpr(LiteralExpr) (interface{}, error)
	VisitUnaryExpr(UnaryExpr) (interface{}, error)
}

type BinaryExpr struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

func (b BinaryExpr) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitBinaryExpr(b)
}

type GroupingExpr struct {
	Expression Expr
}

func (b GroupingExpr) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitGroupingExpr(b)
}

type LiteralExpr struct {
	Value interface{}
}

func (b LiteralExpr) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitLiteralExpr(b)
}

type UnaryExpr struct {
	Operator scanner.Token
	Right    Expr
}

func (b UnaryExpr) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitUnaryExpr(b)
}
