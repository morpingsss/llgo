package godot

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// 定义GDExtensionVariantType类型和枚举值
type GDExtensionVariantType int

const (
	//  atomic types
	GDEXTENSION_VARIANT_TYPE_NIL GDExtensionVariantType = iota
	GDEXTENSION_VARIANT_TYPE_BOOL
	GDEXTENSION_VARIANT_TYPE_INT
	GDEXTENSION_VARIANT_TYPE_FLOAT
	GDEXTENSION_VARIANT_TYPE_STRING

	// math types
	GDEXTENSION_VARIANT_TYPE_VECTOR2
	GDEXTENSION_VARIANT_TYPE_VECTOR2I
	GDEXTENSION_VARIANT_TYPE_RECT2
	GDEXTENSION_VARIANT_TYPE_RECT2I
	GDEXTENSION_VARIANT_TYPE_VECTOR3
	GDEXTENSION_VARIANT_TYPE_VECTOR3I
	GDEXTENSION_VARIANT_TYPE_TRANSFORM2D
	GDEXTENSION_VARIANT_TYPE_VECTOR4
	GDEXTENSION_VARIANT_TYPE_VECTOR4I
	GDEXTENSION_VARIANT_TYPE_PLANE
	GDEXTENSION_VARIANT_TYPE_QUATERNION
	GDEXTENSION_VARIANT_TYPE_AABB
	GDEXTENSION_VARIANT_TYPE_BASIS
	GDEXTENSION_VARIANT_TYPE_TRANSFORM3D
	GDEXTENSION_VARIANT_TYPE_PROJECTION

	// misc types
	GDEXTENSION_VARIANT_TYPE_COLOR
	GDEXTENSION_VARIANT_TYPE_STRING_NAME
	GDEXTENSION_VARIANT_TYPE_NODE_PATH
	GDEXTENSION_VARIANT_TYPE_RID
	GDEXTENSION_VARIANT_TYPE_OBJECT
	GDEXTENSION_VARIANT_TYPE_CALLABLE
	GDEXTENSION_VARIANT_TYPE_SIGNAL
	GDEXTENSION_VARIANT_TYPE_DICTIONARY
	GDEXTENSION_VARIANT_TYPE_ARRAY

	// typed arrays
	GDEXTENSION_VARIANT_TYPE_PACKED_BYTE_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_INT32_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_INT64_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_FLOAT32_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_FLOAT64_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_STRING_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_VECTOR2_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_VECTOR3_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_COLOR_ARRAY

	GDEXTENSION_VARIANT_TYPE_VARIANT_MAX
)

type GDExtensionClassMethodArgumentMetadata int

const (
	GDEXTENSION_METHOD_ARGUMENT_METADATA_NONE GDExtensionClassMethodArgumentMetadata = iota
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT8
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT16
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT32
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT64
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT8
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT16
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT32
	GDEXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT64
	GDEXTENSION_METHOD_ARGUMENT_METADATA_REAL_IS_FLOAT
	GDEXTENSION_METHOD_ARGUMENT_METADATA_REAL_IS_DOUBLE
)

type GDExtensionCallError c.Int

const (
	GDEXTENSION_CALL_OK GDExtensionCallError = iota
	GDEXTENSION_CALL_ERROR_INVALID_METHOD
	GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT
	GDEXTENSION_CALL_ERROR_TOO_MANY_ARGUMENTS
	GDEXTENSION_CALL_ERROR_TOO_FEW_ARGUMENTS
	GDEXTENSION_CALL_ERROR_INSTANCE_IS_NULL
	GDEXTENSION_CALL_ERROR_METHOD_NOT_CONST
)

type GDExtensionConstVariantPtr c.Pointer
type GDExtensionVariantPtr c.Pointer
type GDExtensionClassLibraryPtr c.Pointer
type GDExtensionConstStringNamePtr c.Pointer
type GDExtensionStringNamePtr c.Pointer
type GDExtensionClassInstancePtr c.Pointer
type GDExtensionConstTypePtr c.Pointer
type GDExtensionTypePtr c.Pointer
type GDExtensionStringPtr c.Pointer

type GDExtensionBool uint8
type GDExtensionInt int64

type GDExtensionClassMethodInfo struct {
	Name                 GDExtensionStringNamePtr
	MethodUserdata       c.Pointer
	CallFunc             GDExtensionClassMethodCall
	PtrcallFunc          GDExtensionClassMethodPtrCall
	MethodFlags          uint32
	HasReturnValue       GDExtensionBool
	ReturnValueInfo      *GDExtensionPropertyInfo
	ReturnValueMetadata  GDExtensionClassMethodArgumentMetadata
	ArgumentCount        uint32
	ArgumentsInfo        *GDExtensionPropertyInfo
	ArgumentsMetadata    *GDExtensionClassMethodArgumentMetadata
	DefaultArgumentCount uint32
	DefaultArguments     *GDExtensionVariantPtr
}

type GDExtensionPropertyInfo struct {
	Type       GDExtensionVariantType
	Name       GDExtensionStringNamePtr
	ClassName  GDExtensionStringNamePtr
	Hint       uint32 // Bitfield of `PropertyHint` (defined in `extension_api.json`).
	HintString GDExtensionStringPtr
	Usage      uint32 // Bitfield of `PropertyUsageFlags` (defined in `extension_api.json`).
}

type GDExtensionInterfaceClassdbRegisterExtensionClassMethod func(pLibrary GDExtensionClassLibraryPtr, pClassName GDExtensionConstStringNamePtr, pMethodInfo *GDExtensionClassMethodInfo)

type GDExtensionClassMethodCall func(methodUserdata c.Pointer, pInstance GDExtensionClassInstancePtr, pArgs *GDExtensionConstVariantPtr, pArgumentCount GDExtensionInt, rReturn GDExtensionVariantPtr, rError *GDExtensionCallError)

type GDExtensionClassMethodPtrCall func(methodUserdata c.Pointer, pInstance GDExtensionClassInstancePtr, pArgs *GDExtensionConstTypePtr, rRet GDExtensionTypePtr)

type GDExtensionCallableCustomCall func(callableUserdata c.Pointer, pArgs *GDExtensionConstVariantPtr, pArgumentCount GDExtensionInt, rReturn GDExtensionVariantPtr, rError *GDExtensionCallError)

type GDExtensionCallableCustomEqual func(callableUserdataA, callableUserdataB c.Pointer) GDExtensionBool
