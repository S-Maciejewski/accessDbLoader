package main

import (
	"access_db_generator/accessDb"
	"access_db_generator/logger"
	"access_db_generator/sqlReader"
	"fmt"
	"io"
	"os"
)

const (
	resultFilePath = "result.accdb"
	sqlScriptPath  = "script.sql"
)

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
	logFile := logger.InitLogFile("dbLoader.log")
	defer logger.CloseLogFile(logFile)

	err := copyFile("blank_db.accdb", resultFilePath)
	if err != nil {
		logger.Error(err, "Could not create result database file")
		panic(err)
	}
	logger.Info(fmt.Sprintf("Created an empty Access database file: %s", resultFilePath))

	db := accessDb.New(resultFilePath)
	db.Open()
	defer db.Close()

	sqlReader.ReadAndLoadSqlFile(&db, sqlScriptPath)
}
