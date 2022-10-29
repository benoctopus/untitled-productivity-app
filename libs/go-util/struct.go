package util

import (
	"fmt"
	"reflect"
)

type StructKeyVal[v any] struct {
	Key string
	Val v
}

// ErrNotStruct is returned when a function expecting a struct is passed something else
var ErrNotStruct = fmt.Errorf("argument is not a struct")

// IsNillStruct tests a struct to see if all of it's values are nil
// returns an error if `s` is not a struct
func IsNillStruct(s interface{}, exclude ...string) (bool, error) {
	ents, err := GetStructEntries(s)
	if err != nil {
		return false, err
	}

	for _, v := range ents {
		for _, k := range exclude {
			if k == v.Key {
				continue
			}
		}

		if v.Val != nil {
			return false, nil
		}
	}

	return true, nil
}

// GetStructValues returns a slice of `s`'s values
// returns an error if `s` is not a struct
func GetStructValues(s interface{}) ([]interface{}, error) {
	if !IsStruct(s) {
		return nil, ErrNotStruct
	}

	v := reflect.ValueOf(s)
	vs := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		vs[i] = v.Field(i).Interface()
	}

	return vs, nil
}

// GetStructValues returns a slice of key-value pairs of `s`'s fields
// returns an error if `s` is not a struct
func GetStructEntries(s interface{}) ([]StructKeyVal[any], error) {
	if !IsStruct(s) {
		return nil, ErrNotStruct
	}

	v := reflect.ValueOf(s)
	vs := make([]StructKeyVal[any], v.NumField())

	for i := 0; i < v.NumField(); i++ {
		vs[i] = StructKeyVal[any]{Key: v.Type().Field(i).Name, Val: v.Field(i).Interface()}
	}

	return vs, nil
}

// IsStruct returns true if `s` is a struct
func IsStruct(s interface{}) bool {
	return reflect.ValueOf(s).Kind() == reflect.Struct
}
