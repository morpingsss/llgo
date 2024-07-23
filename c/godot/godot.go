package godot

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname CgoCallFnGDExtensionCallableCustomCall C.cgo_callfn_GDExtensionCallableCustomCall
func CgoCallFnGDExtensionCallableCustomCall(fn GDExtensionCallableCustomCall, callable_userdata c.Pointer,
	p_args *GDExtensionConstVariantPtr, p_argument_count GDExtensionInt,
	r_return GDExtensionVariantPtr, r_error *GDExtensionCallError)

//go:linkname CgoCallfnGDExtensionCallableCustomEqual C.cgo_callfn_GDExtensionCallableCustomEqual
func CgoCallfnGDExtensionCallableCustomEqual(fn GDExtensionCallableCustomEqual, callable_userdata_a, callable_userdata_b c.Pointer) GDExtensionBool
