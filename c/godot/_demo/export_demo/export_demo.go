package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/godot"
)

func testCallfnGDExtensionInterfaceClassdbRegisterExtensionClassMethod(
	fn godot.GDExtensionInterfaceClassdbRegisterExtensionClassMethod,
	pLibrary godot.GDExtensionClassLibraryPtr,
	pClassName godot.GDExtensionConstStringNamePtr,
	pMethodInfo *godot.GDExtensionClassMethodInfo,
) {
	fn(pLibrary, pClassName, pMethodInfo)
	if pMethodInfo != nil {
		methodName := (*c.Char)(pMethodInfo.Name)
		println(c.GoString(methodName))

		paramsCount := godot.GDExtensionInt(4)
		var returnValue c.Int
		pMethodInfo.CallFunc(nil, nil, nil, paramsCount, godot.GDExtensionVariantPtr(&returnValue), nil)
		println(returnValue)
	}
}

func GDExtensionInterfaceClassdbRegisterExtensionClassMethodImpl(
	pLibrary godot.GDExtensionClassLibraryPtr,
	pClassName godot.GDExtensionConstStringNamePtr,
	pMethodInfo *godot.GDExtensionClassMethodInfo,
) {
	// 假设这个实现函数只做一些简单的初始化
	if pMethodInfo != nil {
		pMethodInfo.Name = godot.GDExtensionStringNamePtr(c.Pointer(c.Str("example_method")))
		pMethodInfo.CallFunc = func(
			methodUserdata c.Pointer,
			pInstance godot.GDExtensionClassInstancePtr,
			pArgs *godot.GDExtensionConstVariantPtr,
			pArgumentCount godot.GDExtensionInt,
			rReturn godot.GDExtensionVariantPtr,
			rError *godot.GDExtensionCallError,
		) {
			// 示例函数实现，设置返回值为42
			if rReturn != nil {
				*(*int)(rReturn) = 42
			}
		}
	}
}

func main() {
	var fn godot.GDExtensionInterfaceClassdbRegisterExtensionClassMethod = GDExtensionInterfaceClassdbRegisterExtensionClassMethodImpl
	var library godot.GDExtensionClassLibraryPtr
	var className godot.GDExtensionConstStringNamePtr
	var methodInfo godot.GDExtensionClassMethodInfo

	testCallfnGDExtensionInterfaceClassdbRegisterExtensionClassMethod(fn, library, className, &methodInfo)
}
