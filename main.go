package main

import (
	"access_db_generator/accessDb"
	"io"
	"os"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func main() {
	resultFilePath := "result.accdb"
	err := copy("blank_db.accdb", resultFilePath)
	panicErr(err)

	db := accessDb.New(resultFilePath)
	db.Open()
	db.Close()

	//queryRes, err := db.Query("select MSysObjects.name from MSysObjects")
}
