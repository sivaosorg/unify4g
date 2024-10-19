package unify4go

import "reflect"

// IsPrimitive checks whether the given value is a primitive type in Go.
//
// Primitive types include:
// - Signed integers: int, int8, int16, int32, int64
// - Unsigned integers: uint, uint8, uint16, uint32, uint64, uintptr
// - Floating-point numbers: float32, float64
// - Complex numbers: complex64, complex128
// - Boolean: bool
// - String: string
//
// The function first checks if the input value is `nil`, returning `false` if so. Otherwise, it uses reflection to determine
// the type of the value and compares it against known primitive types.
//
// Parameters:
// - `value`: An interface{} that can hold any Go value. The function checks the type of this value.
//
// Returns:
// - `true` if the value is of a primitive type.
// - `false` if the value is `nil` or not a primitive type.
//
// Example:
//
// primitive := 42
//
//	if IsPrimitive(primitive) {
//	    fmt.Println("The value is a primitive type.")
//	} else {
//
//	    fmt.Println("The value is not a primitive type.")
//	}
func IsPrimitive(value interface{}) bool {
	if value == nil {
		return false
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Bool, reflect.String:
		return true
	default:
		return false
	}
}
