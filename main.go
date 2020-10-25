package main

import (
	"access_db_generator/accessDb"
	"access_db_generator/logger"
	"access_db_generator/sqlProcessor"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

	sqlProcessor.ReadAndLoadSqlFile(&db, args.SqlFilePath)

	// Temporary solution - ExecuteSqlStatement calls need time to finish last inserts before connection is closed
	time.Sleep(5 * time.Second) //TODO: Check records amount (simple validation) instead of sleep
}
