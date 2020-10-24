package main

import (
	"access_db_generator/accessDb"
	"access_db_generator/logger"
	"access_db_generator/sqlReader"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := ParseArgs()

	logFile := logger.InitLogFile("dbLoader.log")
	defer logger.CloseLogFile(logFile)

	data, err := Asset("resource/blank_db.accdb")
	if err != nil {
		logger.Error(err, "Could not read resources via bindata.go - corrupted build")
		os.Exit(1)
	}

	err = ioutil.WriteFile(args.ResultDbPath, data, 0666)
	if err != nil {
		logger.Error(err, "Could not generate Access database file")
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("Created an empty Access database file: %s", args.ResultDbPath))

	db := accessDb.New(args.ResultDbPath)
	db.Open()
	defer db.Close()

	sqlReader.ReadAndLoadSqlFile(&db, args.SqlFilePath)
}
