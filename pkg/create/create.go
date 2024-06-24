package create

import (
	"os"
	"path/filepath"
)

func Create(name string) {
	os.MkdirAll(filepath.Join(name, "src"), 0666)
	os.WriteFile(filepath.Join(name, "src", name+".nm"), []byte("fn main() => int {\n\treturn 0\n}"), 0666)
}
