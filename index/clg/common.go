package clg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"reflect"

	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/spec"
)

// TODO
func (i *clgIndex) getCLGInputExamples(methodValue reflect.Value) ([]interface{}, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGInputExamples")

	return nil, nil
}

// TODO
func (i *clgIndex) getCLGInputTypes(methodValue reflect.Value) ([]reflect.Kind, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGInputTypes")

	return nil, nil
}

func (i *clgIndex) getCLGLookupTable(clgNames []string, packageInfos []packageInfo) (map[string]string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGLookupTable")

	newLookupTable := map[string]string{}

	for _, clgName := range clgNames {
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

// TODO
func (i *clgIndex) getCLGMethodHash(clgName, clgBody string) (string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGMethodHash")

	return "", nil
}

func (i *clgIndex) getCLGNames(clgCollection spec.CLGCollection) ([]string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGNames")

	args, err := clgCollection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	newCLGNames, err := collection.ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCLGNames, nil
}

func (i *clgIndex) getCLGNameQueue(clgNames []string) chan string {
	newCLGNameQueue := make(chan string, len(clgNames))

	for _, clgName := range clgNames {
		newCLGNameQueue <- clgName
	}

	return newCLGNameQueue
}

// TODO
func (i *clgIndex) getCLGOutputExamples(methodValue reflect.Value) ([]interface{}, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGOutputExamples")

	return nil, nil
}

// TODO
func (i *clgIndex) getCLGOutputTypes(methodValue reflect.Value) ([]reflect.Kind, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGOutputTypes")

	return nil, nil
}

type packageInfo struct {
	AstFile      *ast.File
	TokenFileSet *token.FileSet
}

func (i *clgIndex) getCLGPackageInfos() ([]packageInfo, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGPackageInfos")

	var newPackageInfos []packageInfo

	for _, fileName := range LoaderFileNames() {
		raw, err := LoaderReadFile(fileName)
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

// TODO
func (i *clgIndex) getCLGRightSideNeighbours(clgCollection spec.CLGCollection, clgName string, methodValue reflect.Value, canceler <-chan struct{}) ([]string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGRightSideNeighbours")

	// Fill a queue.
	// TODO create method for this getCLGMethodQueue
	args, err := clgCollection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	clgNames, err := collection.ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}

	//     find right side neighbours for given clg name
	//         if no profile for checked neighbour
	//             push neighbour name back to channel

	return nil, nil
}

func (i *clgIndex) getWorkerFunc(clgCollection spec.CLGCollection, clgNameQueue chan string, lookupTable map[string]string) func(canceler <-chan struct{}) error {
	workerFunc := func(canceler <-chan struct{}) error {
		for {
			select {
			case <-canceler:
				return maskAny(workerCanceledError)
			case clgName := <-clgNameQueue:
				// Try to fetch the CLG profile in advance.
				currentCLGProfile, err := i.GetCLGProfileByName(clgName)
				if IsCLGProfileNotFound(err) {
					// In case the CLG profile cannot be found, we are going ahead to create
					// one.
				} else if err != nil {
					return maskAny(err)
				}

				clgBody, ok := lookupTable[clgName]
				if !ok {
					return maskAnyf(clgBodyNotFoundError, clgName)
				}

				newCLGProfile, err := i.CreateCLGProfile(clgCollection, clgName, clgBody, canceler)
				if err != nil {
					return maskAny(err)
				}

				if currentCLGProfile != nil && currentCLGProfile.Equals(newCLGProfile) {
					// The CLG profile has not changed. Thus nothing to do here.
					continue
				}

				err = i.StoreCLGProfile(newCLGProfile)
				if err != nil {
					return maskAny(err)
				}
			}
		}
	}

	return workerFunc
}

func (i *clgIndex) isMethodValue(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}

	if v.Kind() != reflect.Func {
		return false
	}

	return true
}

// TODO
func (i *clgIndex) isRightSideCLGNeighbour(clgCollection spec.CLGCollection, left, right spec.CLGProfile) (bool, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call isRightSideNeighbour")

	// run clg chain
	// if error
	//     return false

	return false, nil
}

// maybeReturnAndLogErrors returns the very first error that may be given by
// errors. All upcoming errors may be given by the provided error channel will
// be logged using the configured logger with log level E and verbosity 4.
func (i *clgIndex) maybeReturnAndLogErrors(errors chan error) error {
	var executeErr error

	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		if executeErr == nil {
			// Only return the first error.
			executeErr = err
		} else {
			// Log all errors but the first one
			i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}

	if executeErr != nil {
		return maskAny(executeErr)
	}

	return nil
}
