package ast

//
//import "github.com/praveensanap/glox-interpreter/scanner"
//
//type Expr interface {
//	accept(visitor ExprVisitor)
//}
//
//type ExprVisitor interface {
//	VisitBinaryExpr(BinaryExpr) Expr
//	VisitGroupingExpr(GroupingExpr) Expr
//	VisitLiteralExpr(LiteralExpr) Expr
//	VisitUnaryExpr(UnaryExpr) Expr
//}
//
//type BinaryExpr struct {
//	left     Expr
//	operator scanner.Token
//	right    Expr
//}
//
//func (b BinaryExpr) accept(visitor ExprVisitor) {
//	visitor.VisitBinaryExpr(b)
//}
//
//type GroupingExpr struct {
//	expression Expr
//}
//
//func (b GroupingExpr) accept(visitor ExprVisitor) {
//	visitor.VisitGroupingExpr(b)
//}
//
//type LiteralExpr struct {
//	value interface{}
//}
//
//func (b LiteralExpr) accept(visitor ExprVisitor) {
//	visitor.VisitLiteralExpr(b)
//}
//
//type UnaryExpr struct {
//	operator scanner.Token
//	right    Expr
//}
//
//func (b UnaryExpr) accept(visitor ExprVisitor) {
//	visitor.VisitUnaryExpr(b)
//}
