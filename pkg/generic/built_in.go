package generic

import "path/filepath"

// BuiltInTypes return a slice with built-in generic type names
func BuiltInTypes() []string {
	res := []string{}
	for _, name := range AssetNames() {
		res = append(res, filepath.Dir(name))
	}
	return res
}

// BuiltIn return a generator object by a given built-in type name
func BuiltIn(typeName string) (*Generator, error) {
	src, err := Asset(filepath.Join(typeName, typeName+".go"))
	if err != nil {
		return nil, err
	}
	return FromBytes(src), nil
}
