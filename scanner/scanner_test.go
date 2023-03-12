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
		{
			name:   "string",
			source: "{\"string\"}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(STRING, "\"string\"", "string", 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
		{
			name:   "line comment",
			source: "{}\n//abd\n{}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(LEFT_BRACE, "{", nil, 3),
				NewToken(RIGHT_BRACE, "}", nil, 3),
				NewToken(EOF, "", nil, 3),
			},
		},

		{
			name:   "single line block comment",
			source: "{}/*abd*/{}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
		{
			name:   "single line nested block comment",
			source: "{}/*a /*b*/d*/{}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
		{
			name:   "multi line nested block comment",
			source: "{}/*a\n/*b\n*/\nd*/{}",
			want: []Token{
				NewToken(LEFT_BRACE, "{", nil, 1),
				NewToken(RIGHT_BRACE, "}", nil, 1),
				NewToken(LEFT_BRACE, "{", nil, 4),
				NewToken(RIGHT_BRACE, "}", nil, 4),
				NewToken(EOF, "", nil, 4),
			},
		},
		{
			name:   "string and number",
			source: "\"x\"=1.2",
			want: []Token{
				NewToken(STRING, "\"x\"", "x", 1),
				NewToken(EQUAL, "=", nil, 1),
				NewToken(NUMBER, "1.2", 1.2, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
		{
			name:   "keywordOrIdentifier",
			source: "x=1.2",
			want: []Token{
				NewToken(IDENTIFIER, "x", nil, 1),
				NewToken(EQUAL, "=", nil, 1),
				NewToken(NUMBER, "1.2", 1.2, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
		{
			name:   "keyword",
			source: "this.x = 1.2",
			want: []Token{
				NewToken(THIS, "this", nil, 1),
				NewToken(DOT, ".", nil, 1),
				NewToken(IDENTIFIER, "x", nil, 1),
				NewToken(EQUAL, "=", nil, 1),
				NewToken(NUMBER, "1.2", 1.2, 1),
				NewToken(EOF, "", nil, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.source)
			s.ScanTokens()
			for i := 0; i < len(s.tokens); i++ {
				if i >= len(tt.want) {
					t.Fatal(fmt.Sprintf("len mismatch: \n%+v got: \n%+v", len(s.tokens), len(tt.want)))
				}
				if !reflect.DeepEqual(s.tokens[i], tt.want[i]) {
					t.Fatal(fmt.Sprintf("item mismatch %d: \n%+v got: \n%+v", i, tt.want[i], s.tokens[i]))
				}
			}

		})
	}
}
