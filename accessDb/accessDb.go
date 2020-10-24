package accessDb

import (
	"access_db_generator/logger"
	"database/sql"
	"fmt"
	"github.com/bennof/accessDBwE"
)

type AccessDb struct {
	path string
	db   *sql.DB
}

func New(path string) (adb AccessDb) {
	adb = AccessDb{path, nil}
	return adb
}

func (adb *AccessDb) Open() {
	db, err := accessdbwe.Open("adodb", fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s", adb.path))
	if err != nil {
		logger.Error(err, fmt.Sprintf("Could not open Access database connection for %s", adb.path))
	}
	adb.db = db
}

func (adb *AccessDb) Close() {
	if adb.db != nil {
		err := adb.db.Close()
		if err != nil {
			logger.Error(err, "Could not close Access database connection")
		}
	} else {
		logger.Warning("There is no open Access database connection to close")
	}
}

func (adb *AccessDb) GetConnectionCount() int {
	if adb.db != nil {
		return adb.db.Stats().OpenConnections
	}
	return 0
}

func (adb *AccessDb) ExecuteSqlStatement(stmt string) {
	if adb.db != nil {
		_, err := adb.db.Query(stmt)
		if err != nil {
			logger.ExecutionError(err, fmt.Sprintf("Could not execute the statement query: %s", stmt))
		}
	}
}
