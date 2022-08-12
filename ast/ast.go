package ast

import "github.com/praveensanap/glox-interpreter/scanner"

type Expr interface {
	accept(visitor ExprVisitor)
}

type ExprVisitor interface {
	VisitBinaryExpr(BinaryExpr)
	VisitGroupingExpr(GroupingExpr)
	VisitLiteralExpr(LiteralExpr)
	VisitUnaryExpr(UnaryExpr)
}

type BinaryExpr struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

func (b BinaryExpr) accept(visitor ExprVisitor) {
	visitor.VisitBinaryExpr(b)
}

type GroupingExpr struct {
	Expression Expr
}

func (b GroupingExpr) accept(visitor ExprVisitor) {
	visitor.VisitGroupingExpr(b)
}

type LiteralExpr struct {
	Value interface{}
}

func (b LiteralExpr) accept(visitor ExprVisitor) {
	visitor.VisitLiteralExpr(b)
}

type UnaryExpr struct {
	Operator scanner.Token
	Right    Expr
}

func (b UnaryExpr) accept(visitor ExprVisitor) {
	visitor.VisitUnaryExpr(b)
}
