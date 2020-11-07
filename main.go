package main

import (
	"accessDbLoader/accessDb"
	"accessDbLoader/logger"
	"accessDbLoader/sqlProcessor"
	"fmt"
	"os"
	"time"
)

func main() {
	args := ParseArgs()

	logFile := logger.InitLogFile("dbLoader.log")
	defer logger.CloseLogFile(logFile)

	if _, err := os.Stat(args.ResultDbPath); err == nil {
		logger.Info(fmt.Sprintf("Opening existing file: %s", args.ResultDbPath))
	} else if os.IsNotExist(err) {
		logger.Warning(fmt.Sprintf("Could not find or access file: %s - created a new file", args.ResultDbPath))
		GenerateResultFile(args.ResultDbPath)
	} else {
		panic(err)
	}

	db := accessDb.New(args.ResultDbPath)
	db.Open()
	defer db.Close()

	sqlProcessor.ReadAndLoadSqlFile(&db, args.SqlFilePath)

	// Temporary solution - ExecuteSqlStatement calls need time to finish last inserts before connection is closed
	time.Sleep(2 * time.Second) //TODO: Check records amount (simple validation) instead of sleep
}
