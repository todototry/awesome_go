package main

/*
Types which don't support comparisons
Following types don't support comparisons:
		map
		slice
		function
		struct types containing uncomparable fields
		array types with uncomparable elements
		Types which don't support comparisons can't be used as the key types of map types.

Please note,
	although map, slice and function types don't support comparisons, their values can be compared to the bare nil identifier.
comparing two interface values with the same uncomparable dynamic types will panic at run time.
*/

func main() {
	// http://127.0.0.1:55555/article/value-conversions-assignments-and-comparisons.html

}
