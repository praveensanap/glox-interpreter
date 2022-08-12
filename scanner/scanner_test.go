package scanner

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScanTokens(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []Token
	}{
		{
			name:   "basic",
			source: "{{}}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.source)
			s.ScanTokens()
			if len(s.tokens) != len(tt.want) {
				t.Fatal(fmt.Sprintf("len mismatch: \n%+v got: \n%+v", len(s.tokens), len(tt.want)))
			}
			for i := 0; i < len(s.tokens); i++ {
				if !reflect.DeepEqual(s.tokens[i], tt.want[i]) {
					t.Fatal(fmt.Sprintf("item mismatch %d: \n%+v got: \n%+v", i, s.tokens[i], tt.want[i]))
				}
			}

		})
	}
}
