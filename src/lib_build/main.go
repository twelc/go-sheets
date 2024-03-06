package main

import "C"

import (
	"fmt"

	lib "github.com/twelc/go-sheets/lib"
)

func main() {}

//export GetFiltered
func GetFiltered(cred string, sheet string, name string, range_ string, querry string, min string, max string) [][]string {
	conf := lib.GetConfig(cred, sheet, name)
	return lib.GetFiltered(querry, min, max, range_, conf)
}

//export GetAll
func GetAll(cred string, sheet string, name string, range_ string) [][]string {
	conf := lib.GetConfig(cred, sheet, name)
	return lib.GetAll(conf, range_)
}

//export AppendData
func AppendData(cred *C.char, sheet *C.char, sheet_name *C.char, obj_name *C.char, distrct *C.char, value *C.char, ind C.int) {
	conf := lib.GetConfig(C.GoString(cred), C.GoString(sheet), C.GoString(sheet_name))

	lib.AppendData(C.GoString(obj_name), C.GoString(distrct), C.GoString(value), int(ind), conf)
}

//export SaveLine
func SaveLine(cred *C.char, sheet *C.char, sheet_name *C.char, obj_name *C.char, distrct *C.char, value *C.char, ind C.int) {
	conf := lib.GetConfig(C.GoString(cred), C.GoString(sheet), C.GoString(sheet_name))
	lib.SetLine([]interface{}{C.GoString(obj_name), C.GoString(distrct), C.GoString(value)}, fmt.Sprintf("A%v:c%v", int(ind), int(ind)), conf)
}
