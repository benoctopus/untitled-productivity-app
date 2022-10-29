package util_test

import (
	"fmt"
	"testing"
	"util"
)

func TestIsNillStruct(t *testing.T) {
	happyCases := []interface{}{
		struct {
			A *int
			B *string
		}{},
		struct {
		}{},
	}

	sadCases := []interface{}{
		struct {
			A int
			B string
		}{},
		struct {
			A *int
			B string
		}{
			A: new(int),
		},
	}

	errorCases := []interface{}{
		"string",
		1,
		1.0,
		[]int{1, 2, 3},
	}

	for _, c := range happyCases {
		t.Run("Should return true when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if v, err := util.IsNillStruct(c); err != nil {
				t.Errorf("Should not have returned an error")
			} else if !v {
				t.Errorf("Should have returned true")
			}
		})
	}

	for _, c := range sadCases {
		t.Run("Should return false when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if v, err := util.IsNillStruct(c); err != nil {
				t.Errorf("Should not have returned an error")
			} else if v {
				t.Errorf("Should have returned false")
			}
		})
	}

	for _, c := range errorCases {
		t.Run("Should return an error when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if _, err := util.IsNillStruct(c); err == nil {
				t.Errorf("Should have returned an error")
			} else if err != util.ErrNotStruct {
				t.Errorf("Should have returned ErrNotStruct")
			}
		})
	}
}

func TestIsStruct(t *testing.T) {
	happyCases := []interface{}{
		struct {
		}{},
		struct {
			A int
			B string
		}{},
	}

	sadCases := []interface{}{
		"string",
		1,
		1.0,
		[]int{1, 2, 3},
	}

	for _, c := range happyCases {
		t.Run("Should return true when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if !util.IsStruct(c) {
				t.Errorf("Should have returned true")
			}
		})
	}

	for _, c := range sadCases {
		t.Run("Should return false when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if util.IsStruct(c) {
				t.Errorf("Should have returned false")
			}
		})
	}
}

func TestGetStructValues(t *testing.T) {
	// TODO: verify output is expected.

	happyCases := []struct {
		in  interface{}
		out []interface{}
	}{
		{
			in: struct {
				A int
				B string
			}{
				A: 1,
				B: "2",
			},
			out: []interface{}{1, "2"},
		},
	}

	sadCases := []interface{}{
		"string",
		1,
		1.0,
		[]int{1, 2, 3},
	}

	for _, c := range happyCases {
		t.Run("Should return true when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if v, err := util.GetStructValues(c.in); err != nil {
				t.Errorf("Should not have returned an error")
			} else if len(v) != len(c.out) {
				t.Errorf("Should have returned %v", c.out)
			}
		})
	}

	for _, c := range sadCases {
		t.Run("Should return false when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if _, err := util.GetStructValues(c); err == nil {
				t.Errorf("Should have returned an error")
			} else if err != util.ErrNotStruct {
				t.Errorf("Should have returned ErrNotStruct")
			}
		})
	}
}

func TestGetStructEntries(t *testing.T) {
	// TODO: verify output is expected.

	happyCases := []struct {
		in  interface{}
		out []util.StructKeyVal[any]
	}{
		{
			in: struct {
				A int
				B string
			}{
				A: 1,
				B: "2",
			},
			out: []util.StructKeyVal[any]{
				{
					Key: "A",
					Val: 1,
				},
				{
					Key: "B",
					Val: "2",
				},
			},
		},
	}

	sadCases := []interface{}{
		"string",
		1,
		1.0,
		[]int{1, 2, 3},
	}

	for _, c := range happyCases {
		t.Run("Should return true when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if v, err := util.GetStructEntries(c.in); err != nil {
				t.Errorf("Should not have returned an error")
			} else if len(v) != len(c.out) {
				t.Errorf("Should have returned %v", c.out)
			}
		})
	}

	for _, c := range sadCases {
		t.Run("Should return false when passed "+fmt.Sprintf("%v", c), func(t *testing.T) {
			t.Parallel()

			if _, err := util.GetStructEntries(c); err == nil {
				t.Errorf("Should have returned an error")
			} else if err != util.ErrNotStruct {
				t.Errorf("Should have returned ErrNotStruct")
			}
		})
	}
}
