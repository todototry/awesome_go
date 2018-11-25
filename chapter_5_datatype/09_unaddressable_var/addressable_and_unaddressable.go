package main

/*
Values that can't be taken addresses
Following values can't be taken addresses:
		bytes in strings
		map elements
		dynamic values of interface values (exposed by type assertions)
		constant values
		literal values
		package-level functions
		methods (used as function values)
		intermediate values
		function calls
		explicit value conversions
		all sorts of operations, except pointer dereference operations, but including:
		channel receive operations
		sub-string operations
		sub-slice operations
		addition, subtraction, multiplication, and division, etc.
Please note, there is a syntax sugar, &T{}, in Go. It is a short form of tmp := T{}; (&tmp), so &T{} is legal doesn't mean the composite literal T{} is addressable.


BTW, following values can be taken addresses (except the just mentioned composite literals):
		variables
		fields of addressable structs
		elements of addressable arrays
		elements of any slices (whether the slices are addressable or not)
		pointer dereference operations
*/
func main() {
	// http://go101.org/article/summaries.html#not-addressable

}
