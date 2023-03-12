package ast

import (
	"github.com/praveensanap/glox-interpreter/scanner"
	"testing"
)

func TestPrinter_Print(t *testing.T) {
	type args struct {
		expr Expr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				BinaryExpr{
					Left: UnaryExpr{
						scanner.NewToken(scanner.MINUS, "-", nil, 1),
						LiteralExpr{Value: 123},
					},
					Operator: scanner.NewToken(scanner.SLASH, "*", nil, 1),
					Right: GroupingExpr{
						Expression: LiteralExpr{Value: 45.67},
					},
				},
			},
			want: "(* (- 123) (group 45.67))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Printer{}
			if got := p.Print(tt.args.expr); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
