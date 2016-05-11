package profile

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) CreateBody(clgName string) (string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateBody")

	for _, fileName := range g.LoaderFileNames() {
		// Load a file and parse its structure.
		raw, err := g.LoaderReadFile(fileName)
		if err != nil {
			return "", maskAny(err)
		}
		fileSet := token.NewFileSet()
		astFile, err := parser.ParseFile(fileSet, fileName, raw, 0)
		if err != nil {
			return "", maskAny(err)
		}
		// Check all function declarations of the current file.
		for _, decl := range astFile.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok {
				fdName := fd.Name.String()
				if fdName != clgName {
					continue
				}

				var tbuf bytes.Buffer
				printer.Fprint(&tbuf, fileSet, fd.Type)
				funcType := tbuf.String()

				var bbuf bytes.Buffer
				printer.Fprint(&bbuf, fileSet, fd.Body)
				funcBody := bbuf.String()

				return fmt.Sprintf("%s %s", funcType, funcBody), nil
			}
		}
	}

	return "", maskAnyf(clgBodyNotFoundError, clgName)
}
