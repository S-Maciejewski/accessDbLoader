package sqlProcessor

import (
	"access_db_generator/accessDb"
	"access_db_generator/logger"
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"os"
	"time"
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
	logger.Info(fmt.Sprintf("Executing SQL statements in %s", sqlFilePath))
	start := time.Now()
	file, err := os.Open(sqlFilePath)
	if err != nil {
		logger.Error(err, fmt.Sprintf("Could not open SQL file at %s", sqlFilePath))
		return
	}

	// Initialize progress bar
	stat, _ := file.Stat()
	progressBar := pb.StartNew(int(stat.Size()))

	scanner := bufio.NewScanner(file)
	scanner.Split(SplitStatements)
	statementCount := 0
	for scanner.Scan() {
		stmt := scanner.Text()
		db.ExecuteSqlStatement(scanner.Text())
		// Advance progress bar by the length of current statement
		progressBar.Add(len(stmt))
		statementCount++
	}
	progressBar.Finish()
	logger.Info(fmt.Sprintf("Executed %d SQL statements from %s in %s ", statementCount, sqlFilePath, time.Since(start)))
}
