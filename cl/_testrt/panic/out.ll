; ModuleID = 'main'
source_filename = "main"

%"github.com/goplus/llgo/internal/runtime.String" = type { ptr, i64 }
%"github.com/goplus/llgo/internal/runtime.iface" = type { ptr, ptr }

@"main.init$guard" = global ptr null
@0 = private unnamed_addr constant [14 x i8] c"panic message\00", align 1

define void @main.init() {
_llgo_0:
  %0 = load i1, ptr @"main.init$guard", align 1
  br i1 %0, label %_llgo_2, label %_llgo_1

_llgo_1:                                          ; preds = %_llgo_0
  store i1 true, ptr @"main.init$guard", align 1
  br label %_llgo_2

_llgo_2:                                          ; preds = %_llgo_1, %_llgo_0
  ret void
}

define void @main() {
_llgo_0:
  call void @"github.com/goplus/llgo/internal/runtime.init"()
  call void @main.init()
  %0 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr @0, i64 13)
  %1 = call %"github.com/goplus/llgo/internal/runtime.iface" @"github.com/goplus/llgo/internal/runtime.MakeAnyString"(%"github.com/goplus/llgo/internal/runtime.String" %0)
  call void @"github.com/goplus/llgo/internal/runtime.TracePanic"(%"github.com/goplus/llgo/internal/runtime.iface" %1)
  unreachable
}

declare void @"github.com/goplus/llgo/internal/runtime.init"()

declare %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr, i64)

declare %"github.com/goplus/llgo/internal/runtime.iface" @"github.com/goplus/llgo/internal/runtime.MakeAnyString"(%"github.com/goplus/llgo/internal/runtime.String")

declare void @"github.com/goplus/llgo/internal/runtime.TracePanic"(%"github.com/goplus/llgo/internal/runtime.iface")