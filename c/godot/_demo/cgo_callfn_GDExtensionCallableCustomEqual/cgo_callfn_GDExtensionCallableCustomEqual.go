package main

import (
	"unsafe"

	"github.com/goplus/llgo/c/godot"
)

// 一个示例的比较函数
func exampleEqualFunc(a, b unsafe.Pointer) godot.GDExtensionBool {
	if a == b {
		return godot.GDExtensionBool(1)
	}
	return godot.GDExtensionBool(0)
}
func main() {
	// 定义两个示例指针
	var dataA int
	var dataB int

	// 将指针转换为unsafe.Pointer
	ptrA := unsafe.Pointer(&dataA)
	ptrB := unsafe.Pointer(&dataB)
	// 调用CgoCallfnGDExtensionCallableCustomEqual
	result := godot.CgoCallfnGDExtensionCallableCustomEqual(exampleEqualFunc, ptrA, ptrB)
	println(result)
}
