package eval

import (
	"fmt"
	"github.com/praveensanap/glox-interpreter/ast"
	"github.com/praveensanap/glox-interpreter/scanner"
	"strings"
)

// post-order traversal on the AST
type Interpreter struct {
}

func (i Interpreter) Interpret(expr ast.Expr) (interface{}, error) {
	return i.evaluate(expr)
}

func (i Interpreter) InterpretAndPrint(expr ast.Expr) {
	result, err := i.evaluate(expr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stringify(result))
	}
}

func stringify(obj interface{}) string {
	if obj == nil {
		return "nil"
	}
	if v, ok := obj.(float64); ok {
		s := fmt.Sprintf("%f", v)
		if strings.HasSuffix(s, ".0") {
			return strings.TrimSuffix(s, ".0")
		}
	}
	return fmt.Sprintf("%v", obj)
}

func (i Interpreter) VisitBinaryExpr(expr ast.BinaryExpr) (interface{}, error) {
	left, err := i.evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.GetTokenType() {
	case scanner.MINUS:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case scanner.SLASH:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		if right.(float64) == 0 {
			return nil, &RuntimeError{expr.Operator, "division by zero"}
		}
		return left.(float64) / right.(float64), nil
	case scanner.STAR:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil
	case scanner.PLUS:
		// add two numbers
		if left, ok := left.(float64); ok {
			if right, ok := right.(float64); ok {
				return left + right, nil
			} else {
				return nil, &RuntimeError{expr.Operator, "for + both operands must be numbers"}
			}
		}
		// concat two string
		if left, ok := left.(string); ok {
			if right, ok := right.(string); ok {
				return left + right, nil
			} else {
				return nil, &RuntimeError{expr.Operator, "for + both operands must be strings"}
			}
		}
		return nil, &RuntimeError{expr.Operator, "for + both operands must be numbers or strings"}
		// handle error
	case scanner.GREATER:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil
	case scanner.GREATER_EQUAL:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) >= right.(float64), nil
	case scanner.LESS:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil
	case scanner.LESS_EQUAL:
		err := checkNumberOperands(expr.Operator, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) <= right.(float64), nil
	case scanner.BANG_EQUAL:
		return !isEqual(left, right), nil
	case scanner.EQUAL_EQUAL:
		return isEqual(left, right), nil
	}
	return nil, &RuntimeError{expr.Operator, "unknown binary operator"}
}

// no implicit type conversion
func isEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i Interpreter) VisitGroupingExpr(expr ast.GroupingExpr) (interface{}, error) {
	return i.evaluate(expr.Expression)
}

func (i Interpreter) evaluate(expression ast.Expr) (interface{}, error) {
	return expression.Accept(i)
}

func (i Interpreter) VisitLiteralExpr(expr ast.LiteralExpr) (interface{}, error) {
	return expr.Value, nil
}

func (i Interpreter) VisitUnaryExpr(expr ast.UnaryExpr) (interface{}, error) {
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.GetTokenType() {
	case scanner.BANG:
		return !isTruthy(right), nil
	case scanner.MINUS:
		err := checkNumberOperand(expr.Operator, right)
		if err != nil {
			return nil, err
		}
		return -right.(float64), nil
	}
	return nil, &RuntimeError{expr.Operator, "unknown unary operator"}
}

// like Ruby. nil and false are falsey
func isTruthy(value interface{}) bool {
	if value == nil {
		return false
	}
	if value, ok := value.(bool); ok {
		return value
	}
	return true
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func checkNumberOperands(operator scanner.Token, left interface{}, right interface{}) error {
	if _, ok := left.(float64); ok {
		if _, ok := right.(float64); ok {
			return nil
		}
	}
	return &RuntimeError{operator, operator.GetLexeme() + "operands must be numbers"}

}

func checkNumberOperand(operator scanner.Token, left interface{}) error {
	if _, ok := left.(float64); ok {
		return nil
	}
	return &RuntimeError{operator, operator.GetLexeme() + "operand must be number"}

}

type RuntimeError struct {
	Token   scanner.Token
	Message string
}

func (r RuntimeError) Error() string {
	return fmt.Sprintf("%s\n[line %d]", r.Message, r.Token.GetLine())
}
