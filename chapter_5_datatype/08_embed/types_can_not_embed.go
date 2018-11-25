/*
	声明：
	the following types can't be embedded.
		1. Defined pointer types.
		2. Unnamed non-pointer types.
		3. Pointer types whose base types are either interface or poiner types.

	举例：
	如下情况中：
		type Encoder interface {Encode([]byte) []byte}
		type Person struct {name string; age int}
		type Alias = struct {name string; age int}
		type AliasPtr = *struct {name string; age int}
		type IntPtr *int

	// 可以嵌入结构体的有：These types can be embedded in a struct type.
		Encoder
		Person
		*Person
		Alias
		*Alias
		AliasPtr
		int
		*int

	// These types can't be embedded.
		*Encoder         // base type is an interface type
		*AliasPtr        // base type is a pointer type
		IntPtr           // defined pointer type
		*IntPtr          // base type is a pointer type
		struct {age int} // unnamed non-pointer type
		map[string]int   // unnamed non-pointer type
		[]int64          // unnamed non-pointer type
		func()           // unnamed non-pointer type

*/

