package main

import (
	"database/sql"
	"fmt"
	"github.com/bennof/accessDBwE"
	"io"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func copy(src, dst string) error {
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

func getAccessDatabase(fileName string) (db *sql.DB, err error)  {
	db, err = accessdbwe.Open("adodb", fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s;", fileName))
	return
}

func main() {
	err := copy("blank_db.accdb", "result.accdb")
	checkErr(err)
	db, err := getAccessDatabase("result.accdb")
	checkErr(err)


	db.Close()
	// use db like any other sql.db handle
}
