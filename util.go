package main

import (
	"accessDbLoader/logger"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Args struct {
	ResultDbPath string
	SqlFilePath  string
}

func ParseArgs() (args Args) {
	flag.PrintDefaults()
	resultDbPath := flag.String("db-path", "./result.accdb",
		"Specify output database file path or a path where new database file should be created")
	sqlFilePath := flag.String("sql-path", "", "Specify SQL file to execute path")
	flag.Parse()
	args = Args{ResultDbPath: *resultDbPath, SqlFilePath: *sqlFilePath}
	return
}

func GenerateResultFile(dbFilePath string) {
	data, err := Asset("resource/blank_db.accdb")
	if err != nil {
		logger.Error(err, "Could not read resources via bindata.go - corrupted build")
		os.Exit(1)
	}

	err = ioutil.WriteFile(dbFilePath, data, 0666)
	if err != nil {
		logger.Error(err, "Could not generate Access database file")
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("Created an empty Access database file: %s", dbFilePath))
}
