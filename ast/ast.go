package ast;

import "github.com/praveensanap/glox-interpreter/scanner";

type Expr interface {
	Accept(visitor  ExprVisitor) interface{}
}

type ExprVisitor interface {
	VisitBinaryExpr (BinaryExpr) interface{}
	VisitGroupingExpr (GroupingExpr) interface{}
	VisitLiteralExpr (LiteralExpr) interface{}
	VisitUnaryExpr (UnaryExpr) interface{}
}

type BinaryExpr struct {
	Left Expr
	Operator scanner.Token
	Right Expr
}

func (b BinaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(b)
}

type GroupingExpr struct {
	Expression Expr
}

func (b GroupingExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(b)
}

type LiteralExpr struct {
	Value interface{}
}

func (b LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(b)
}

type UnaryExpr struct {
	Operator scanner.Token
	Right Expr
}

func (b UnaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(b)
}

