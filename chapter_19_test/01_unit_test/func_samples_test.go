package main

import "testing"

func Test_add(t *testing.T) {
	t.Parallel()

	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{"add1", args{1, 6}, 7},
		{"add2", args{3, 3}, 6},
		{"add3", args{1, 6}, 7},
		{"add4", args{3, 3}, 6},
		{"add5", args{1, 6}, 7},
		{"add6", args{3, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addSlice(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"slice1", args{[]int{1, 3, 4}}, 8},
		{"slice2", args{[]int{3, 3, 4}}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSlice(tt.args.a...); got != tt.want {
				t.Errorf("addSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjoin(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"adjoin1", args{"kw", "args"}, "kwargs"},
		{"adjoin2", args{"1", ""}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjoin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("adjoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
