package main

import (
	"fmt"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/godot"
)

// 示例函数
func fn(callableUserdata c.Pointer, pArgs *godot.GDExtensionConstVariantPtr, pArgumentCount godot.GDExtensionInt, rReturn godot.GDExtensionVariantPtr, rError *godot.GDExtensionCallError) {
	fmt.Printf("dummy_function called with %d arguments.\n", int(pArgumentCount))
}

func main() {
	// 示例变量
	var callableUserdata c.Pointer
	var pArgs godot.GDExtensionConstVariantPtr
	var rReturn godot.GDExtensionVariantPtr

	// 设置错误类型为无错误
	rError := godot.GDEXTENSION_CALL_OK

	// 调用 cgoCallFnGDExtensionCallableCustomCall
	godot.CgoCallFnGDExtensionCallableCustomCall(
		godot.GDExtensionCallableCustomCall(fn),
		callableUserdata,
		&pArgs,
		0,
		rReturn,
		&rError,
	)
}
