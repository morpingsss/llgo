package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"github.com/goplus/llgo/chore/_xtool/llcppsymg/common"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/cjson"
	"github.com/goplus/llgo/c/clang"
)

type Context struct {
	namespaceName string
	className     string
	astInfo       []common.ASTInformation
	filename      *c.Char
}

func newContext() *Context {
	return &Context{
		astInfo: make([]common.ASTInformation, 0),
	}
}

func (c *Context) setNamespaceName(name string) {
	c.namespaceName = name
}

func (c *Context) setClassName(name string) {
	c.className = name
}

func (c *Context) setFilename(filename *c.Char) {
	c.filename = filename
}

var context = newContext()

func collectFuncInfo(cursor clang.Cursor) common.ASTInformation {

	info := common.ASTInformation{
		Namespace: context.namespaceName,
		Class:     context.className,
	}

	cursorStr := cursor.String()
	symbol := cursor.Mangling()

	info.Name = c.GoString(cursorStr.CStr())

	info.Symbol = c.GoString(symbol.CStr())
	if len(info.Symbol) >= 1 {
		if info.Symbol[0] == '_' {
			info.Symbol = info.Symbol[1:]
		}
	}

	defer symbol.Dispose()
	defer cursorStr.Dispose()

	if context.namespaceName != "" {
		info.Namespace = context.namespaceName
	}
	if context.className != "" {
		info.Class = context.className
	}

	typeStr := cursor.ResultType().String()
	defer typeStr.Dispose()
	info.ReturnType = c.GoString(typeStr.CStr())

	info.Parameters = make([]common.Parameter, cursor.NumArguments())
	for i := 0; i < int(cursor.NumArguments()); i++ {
		argCurSor := cursor.Argument(c.Uint(i))
		argType := argCurSor.Type().String()
		argName := argCurSor.String()
		info.Parameters[i] = common.Parameter{
			Name: c.GoString(argName.CStr()),
			Type: c.GoString(argType.CStr()),
		}

		argType.Dispose()
		argName.Dispose()
	}

	return info
}

func visit(cursor, parent clang.Cursor, clientData c.Pointer) clang.ChildVisitResult {
	if cursor.Kind == clang.Namespace {
		nameStr := cursor.String()
		context.setNamespaceName(c.GoString(nameStr.CStr()))
		clang.VisitChildren(cursor, visit, nil)
		context.setNamespaceName("")
	} else if cursor.Kind == clang.ClassDecl {
		nameStr := cursor.String()
		context.setClassName(c.GoString(nameStr.CStr()))
		clang.VisitChildren(cursor, visit, nil)
		context.setClassName("")
	} else if cursor.Kind == clang.CXXMethod || cursor.Kind == clang.FunctionDecl || cursor.Kind == clang.Constructor || cursor.Kind == clang.Destructor {
		loc := cursor.Location()
		var file clang.File
		var line, column c.Uint

		loc.SpellingLocation(&file, &line, &column, nil)
		filename := file.FileName()

		if c.Strcmp(filename.CStr(), context.filename) == 0 {
			info := collectFuncInfo(cursor)
			info.Location = c.GoString(filename.CStr()) + ":" + strconv.Itoa(int(line)) + ":" + strconv.Itoa(int(column))
			context.astInfo = append(context.astInfo, info)
		}

		defer filename.Dispose()

	}

	return clang.ChildVisit_Continue
}

func parse(filename *c.Char) []common.ASTInformation {
	index := clang.CreateIndex(0, 0)
	args := make([]*c.Char, 3)
	args[0] = c.Str("-x")
	args[1] = c.Str("c++")
	args[1] = c.Str("c++")
	args[2] = c.Str("-std=c++11")
	unit := index.ParseTranslationUnit(
		filename,
		unsafe.SliceData(args), 3,
		nil, 0,
		clang.TranslationUnit_None,
	)

	if unit == nil {
		println("Unable to parse translation unit. Quitting.")
		c.Exit(1)
	}

	cursor := unit.Cursor()
	context.setFilename(filename)

	clang.VisitChildren(cursor, visit, nil)

	unit.Dispose()
	index.Dispose()

	return context.astInfo
}
func printJson(infos []common.ASTInformation) {
	root := cjson.Array()

	for _, info := range infos {
		item := cjson.Object()
		item.SetItem(c.Str("namespace"), cjson.String(c.AllocaCStr(info.Namespace)))
		item.SetItem(c.Str("class"), cjson.String(c.AllocaCStr(info.Class)))
		item.SetItem(c.Str("name"), cjson.String(c.AllocaCStr(info.Name)))
		item.SetItem(c.Str("returnType"), cjson.String(c.AllocaCStr(info.ReturnType)))
		item.SetItem(c.Str("location"), cjson.String(c.AllocaCStr(info.Location)))
		item.SetItem(c.Str("symbol"), cjson.String(c.AllocaCStr(info.Symbol)))

		params := cjson.Array()
		for _, param := range info.Parameters {
			paramObj := cjson.Object()
			paramObj.SetItem(c.Str("name"), cjson.String(c.AllocaCStr(param.Name)))
			paramObj.SetItem(c.Str("type"), cjson.String(c.AllocaCStr(param.Type)))
			params.AddItem(paramObj)
		}
		item.SetItem(c.Str("parameters"), params)

		root.AddItem(item)
	}
	c.Printf(c.Str("%s\n"), root.Print())
}

func main() {
	if c.Argc < 2 {
		fmt.Fprintln(os.Stderr, "Usage: <C++ header file>\n")
		return
	} else {
		// todo(zzy): receive files
		printJson(parse(c.Index(c.Argv, 1)))
	}
}
