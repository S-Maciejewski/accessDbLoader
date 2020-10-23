package accessDb

import (
	"access_db_generator/log"
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
		log.Error(err, fmt.Sprintf("Could not open Access database connection for %s", adb.path))
	}
	adb.db = db
}

func (adb *AccessDb) Close() {
	if adb.db != nil {
		err := adb.db.Close()
		if err != nil {
			log.Error(err, "Could not close Access database connection")
		}
	} else {
		log.WarningMessage("There is no open Access database connection to close")
	}
}
