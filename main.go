package main

import (
	"access_db_generator/accessDb"
	"access_db_generator/logger"
	"access_db_generator/sqlReader"
	"fmt"
)

func main() {
	args := ParseArgs()

	logFile := logger.InitLogFile("dbLoader.log")
	defer logger.CloseLogFile(logFile)

	err := CopyFile("blank_db.accdb", args.ResultDbPath)
	if err != nil {
		logger.Error(err, "Could not create result database file")
		panic(err)
	}
	logger.Info(fmt.Sprintf("Created an empty Access database file: %s", args.ResultDbPath))

	db := accessDb.New(args.ResultDbPath)
	db.Open()
	defer db.Close()

	sqlReader.ReadAndLoadSqlFile(&db, args.SqlFilePath)
}
