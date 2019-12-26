package converter

type T0 string
type T1 interface{}

type T0Slice []T0

func (s T0Slice) ToT1Slice(convert func(T0) T1) []T1 {
	res := make([]T1, len(s))
	for i, item := range s {
		res[i] = convert(item)
	}
	return res
}
