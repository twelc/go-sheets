package main

import "C"

import (
	lib "github.com/twelc/go-sheets/lib"
)

func main() {
	AppendData(
		"/mnt/E83C5EC13C5E8A88/projects/go-table-editor/src/config/credentials.json",
		"141maOrpeeFsydVAWP-kIaziMCHn_fI8nQv0mFB78TVk",
		"history",
		"Test_name",
		"Test_district",
		"Test_value",
		0,
	)
}

//export GetFiltered
func GetFiltered(cred string, sheet string, name string, range_ string, querry string, min string, max string) [][]string {
	conf := lib.GetConfig(cred, sheet, name, range_)
	return lib.GetFiltered(querry, min, max, conf)
}

//export GetAll
func GetAll(cred string, sheet string, name string, range_ string) [][]string {
	conf := lib.GetConfig(cred, sheet, name, range_)
	return lib.GetAll(conf)
}

//export AppendData
func AppendData(cred string, sheet string, sheet_name string, obj_name string, distrct string, value string, ind int) {
	conf := lib.GetConfig(cred, sheet, sheet_name, "A:A")
	lib.AppendData(obj_name, distrct, value, ind, conf)
}
