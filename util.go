package main

import (
	"flag"
	"io"
	"os"
)

type Args struct {
	ResultDbPath string
	SqlFilePath  string
}

func ParseArgs() (args Args) {
	flag.PrintDefaults()
	resultDbPath := flag.String("db-path", "./result.accdb", "Specify output database file path")
	sqlFilePath := flag.String("sql-path", "", "Specify SQL file to execute path")
	flag.Parse()
	args = Args{ResultDbPath: *resultDbPath, SqlFilePath: *sqlFilePath}
	return
}

func CopyFile(src, dst string) error {
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
