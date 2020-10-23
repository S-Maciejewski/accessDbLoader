package main

import (
	"access_db_generator/accessDb"
	"access_db_generator/sqlReader"
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
	sqlScriptPath := "script.sql"

	err := copyFile("blank_db.accdb", resultFilePath)
	panicErr(err)

	db := accessDb.New(resultFilePath)
	db.Open()
	defer db.Close()

	sqlReader.ReadAndLoadSqlFile(&db, sqlScriptPath)

	//db2 := accessDb.New(resultFilePath)
	//db2.Open()
	//fmt.Println(db2.GetConnectionCount())
	//db.ExecuteSqlStatement("CREATE TABLE TEST(ID INTEGER, NAME CHAR(50));")

	//queryRes, err := db.Query("select MSysObjects.name from MSysObjects")
}
