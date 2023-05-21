package main

import (
	python3 "github.com/go-python/cpy3"
)

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()

	code := `
from nameparser import HumanName

print(HumanName("Dr. Juan Q. Xavier de la Vega III (Doc Vega)").as_dict())
`
	python3.PyRun_SimpleString(code)
}
