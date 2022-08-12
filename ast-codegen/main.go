package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	output := "ast/ast.go"
	source := generateCode()
	err := ioutil.WriteFile(output, []byte(source), 0o644)
	if err != nil {
		panic(err)
	}
}

func generateCode() string {
	return defineAst("Expr", []string{
		"Binary : Expr left, scanner.Token operator, Expr right",
		"Grouping: Expr expression",
		"Literal : interface{} value",
		"Unary: scanner.Token operator, Expr right",
	})
}

type Ttype struct {
	baseName  string
	className string
	fields    []Tfields
}

type Tfields struct {
	name  string
	ttype string
}

func defineAst(baseName string, rules []string) string {
	var buffer bytes.Buffer

	ttypes := []Ttype{}
	for _, r := range rules {
		expr := strings.Split(r, ":")
		className := strings.TrimSpace(expr[0])
		fields := strings.Split(strings.TrimSpace(expr[1]), ",")
		tfields := []Tfields{}
		for _, f := range fields {
			sp := strings.Split(strings.TrimSpace(f), " ")
			tfields = append(tfields, Tfields{
				name:  strings.TrimSpace(sp[1]),
				ttype: strings.TrimSpace(sp[0]),
			})

		}

		ttypes = append(ttypes, Ttype{
			baseName:  baseName,
			className: className,
			fields:    tfields,
		})
	}

	buffer.WriteString("package ast;\n")
	buffer.WriteString("import \"github.com/praveensanap/glox-interpreter/scanner\";\n")

	buffer.WriteString(fmt.Sprintf("type %s interface {\n", baseName))
	buffer.WriteString(fmt.Sprintf("\taccept(visitor  %sVisitor)\n", baseName))
	buffer.WriteString("}\n\n")

	buffer.WriteString(fmt.Sprintf("type %sVisitor interface {\n", baseName))
	for _, f := range ttypes {
		buffer.WriteString(fmt.Sprintf("\tVisit%s%s (%s%s) %s\n", f.className, baseName, f.className, baseName, baseName))
	}
	buffer.WriteString("}\n\n")

	for _, f := range ttypes {
		buffer.WriteString(fmt.Sprintf("type %s%s struct {\n", f.className, baseName))
		for _, ff := range f.fields {
			buffer.WriteString(fmt.Sprintf("\t%s %s\n", ff.name, ff.ttype))
		}
		buffer.WriteString("}\n\n")

		buffer.WriteString(fmt.Sprintf("func (b %s%s) accept(visitor %sVisitor) {\n", f.className, baseName, baseName))
		buffer.WriteString(fmt.Sprintf("\tvisitor.Visit%s%s(b)\n", f.className, f.baseName))
		buffer.WriteString("}\n\n")
	}

	return buffer.String()
}
