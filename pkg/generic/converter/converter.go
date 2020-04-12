package converter

// T0 is a generic variable type placeholder of a source value. It will not appear in the generated code
type T0 string

// T1 is a generic variable type placeholder of a destination value. It will not appear in the generated code
type T1 interface{}

// T0Slice is a slice of T0 elements to convert
type T0Slice []T0

// ToT1Slice converts this slice to a slice of T1 elements using provided conversion function
func (s T0Slice) ToT1Slice(convert func(T0) T1) []T1 {
	res := make([]T1, len(s))
	for i, item := range s {
		res[i] = convert(item)
	}
	return res
}
