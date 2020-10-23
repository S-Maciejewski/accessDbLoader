package sqlReader

import (
	"access_db_generator/accessDb"
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func SplitStatements(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	// Scan until semicolon, marking end of statement.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ';' {
			return i + width, data[start : i+1], nil
		}
	}
	// Return final statement if at EOF
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func ReadAndLoadSqlFile(db *accessDb.AccessDb, sqlFilePath string) {
	file, err := os.Open(sqlFilePath)
	if err != nil {
		panic(err)
	}

	stat, _ := file.Stat()
	fmt.Println("file size:", stat.Size()) // TODO: Progress bar

	scanner := bufio.NewScanner(file)
	scanner.Split(SplitStatements)

	for scanner.Scan() {
		fmt.Println("---")
		fmt.Println(scanner.Text())
	}
}
