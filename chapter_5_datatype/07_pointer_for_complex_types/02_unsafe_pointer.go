package main

// http://go101.org/article/unsafe.html

func main() {

	// Fact 1: Unsafe Pointers Are Pointers And Uintptr Values Are Intergers
	// Fact 2: Unused Values May Be Collected At Any Time
	// Fact 3: We Can Use A runtime.KeepAlive Function Call To Mark A Value As Still In Using (Reachable) Currently
	// Fact 4: *unsafe.Pointer Is A General Safe Pointer Type

	// Pattern 1: Convert *T1 To Unsafe Poniter, Then Convert The Unsafe Pointer Value To *T2
	// Pattern 2: Convert Unsafe Pointer To Uintptr, Then Use The Uintptr Value
	// Pattern 3: Convert Unsafe Pointer To Uintptr, Do Arithmetic Operations With The Uintptr Value, Then Convert Back
	// Pattern 4: Convert Unsafe Pointer To uintptr When Calling syscall.Syscall
	// Pattern 5: Convert The uintptr Result Of reflect.Value.Pointer Or reflect.Value.UnsafeAddr Method Call To Unsafe Pointer
	// Pattern 6: Convert A reflect.SliceHeader.Data Or reflect.StringHeader.Data Field To Unsafe Pointer, And The Inverse.
	// 具体参考 slice_pointer.go

}
