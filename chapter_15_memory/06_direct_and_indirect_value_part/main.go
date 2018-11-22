package main

// http://go101.org/article/value-part.html
func main() {
	// direct value and underlaying value(which pointer points to.)
	/*
		 types below may contain underlying parts :
			slice types,
			map types,
			channel types,
			function types,
			interface types,
			string types.

		builtin types definition:

		Internal Definition Of String Types
				// string types
				type _string struct {
					elements *byte // underlying bytes
					len      int   // number of bytes
				}

				// slice types
				type _slice struct {
					elements unsafe.Pointer // underlying elements
					len      int            // number of elements
					cap      int            // capacity
				}


				// general interface types
				type _interface struct {
					dynamicTypeInfo *struct {
						dynamicType *_type       // the dynamic type
						methods     []*_function // implemented methods
					}
					dynamicValue unsafe.Pointer // the dynamic value
				}
	*/

	// Return Pointers Of Local Variables Is Safe In Go

	//

}
