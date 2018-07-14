package generic

import "path/filepath"

func BuiltInTypes() []string {
	res := []string{}
	for _, name := range AssetNames() {
		res = append(res, filepath.Dir(name))
	}
	return res
}

func BuiltIn(typeName string, pkgName string) (Processor, error) {
	src, err := Asset(filepath.Join(typeName, typeName+".go"))
	if err != nil {
		return nil, err
	}
	return FromBytes(src, pkgName), nil
}
