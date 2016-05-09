package profile

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/spec"
)

func namesFromCollection(collection spec.CLGCollection) ([]string, error) {
	args, err := collection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	newNames, err := collection.ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNames, nil
}

type packageInfo struct {
	AstFile      *ast.File
	TokenFileSet *token.FileSet
}

func getPackageInfos() ([]packageInfo, error) {
	var newPackageInfos []packageInfo

	for _, fileName := range collection.LoaderFileNames() {
		raw, err := collection.LoaderReadFile(fileName)
		if err != nil {
			return nil, maskAny(err)
		}
		fileSet := token.NewFileSet()
		astFile, err := parser.ParseFile(fileSet, fileName, raw, 0)
		if err != nil {
			return nil, maskAny(err)
		}

		newPackageInfo := packageInfo{
			AstFile:      astFile,
			TokenFileSet: fileSet,
		}

		newPackageInfos = append(newPackageInfos, newPackageInfo)
	}

	return newPackageInfos, nil
}

func createLookupTable() (map[string]string, error) {
	// Fetch CLG package information and create a lookup table for CLG names and
	// their corresponding method bodies.
	packageInfos, err := getPackageInfos()
	if err != nil {
		return nil, maskAny(err)
	}

	// Create the lookup table.
	newLookupTable := map[string]string{}

	for _, clgName := range lt.CLGNames {
		for _, packageInfo := range packageInfos {
			for _, decl := range packageInfo.AstFile.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					fdName := fd.Name.String()
					if fdName != clgName {
						continue
					}

					var tbuf bytes.Buffer
					printer.Fprint(&tbuf, packageInfo.TokenFileSet, fd.Type)
					funcType := tbuf.String()

					var bbuf bytes.Buffer
					printer.Fprint(&bbuf, packageInfo.TokenFileSet, fd.Body)
					funcBody := bbuf.String()

					newLookupTable[fdName] = fmt.Sprintf("%s %s", funcType, funcBody)
				}
			}
		}
	}

	return newLookupTable, nil
}
