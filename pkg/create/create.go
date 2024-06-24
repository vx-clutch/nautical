package create

import "os"

func Create(name string) {
	os.Mkdir(name, 0666)
	os.Chdir(name)
	os.Mkdir("src", 0666)
	os.WriteFile("src/main.nm", []byte("fn main() => int {\n\treturn 0\n}"), 0666)
}
