package accessDb

import (
	"io"
	"os"
	"reflect"
	"testing"
)

// Function for test database reproduction
func reproduceTestDatabase(path string) {
	in, _ := os.Open("../test_assets/test.accdb")
	defer in.Close()
	out, _ := os.Create(path)
	defer out.Close()
	_, _ = io.Copy(out, in)
	out.Close()
}

func removeTestDatabase(path string) {
	_ = os.Remove(path)
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantAdb AccessDb
	}{
		{"Non-empty path", "./db.accdb", AccessDb{"./db.accdb", nil}},
		{"Empty path", "", AccessDb{"", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdb := New(tt.path); !reflect.DeepEqual(gotAdb, tt.wantAdb) {
				t.Errorf("New() = %v, want %v", gotAdb, tt.wantAdb)
			}
		})
	}
}

// This tests also adb.GetConnectionCount()
func TestAccessDb_Open(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"Open and check connection number", "db.accdb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reproduceTestDatabase(tt.path)
			db := New(tt.path)
			db.Open()
			if count := db.GetConnectionCount(); count != 1 {
				t.Error("No connections opened")
			}
			removeTestDatabase(tt.path)
		})
	}
}

func TestAccessDb_Close(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"Open, close and check connection number", "db.accdb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reproduceTestDatabase(tt.path)
			db := New(tt.path)
			db.Open()
			if count := db.GetConnectionCount(); count != 1 {
				t.Error("No connections opened - db.Open failed")
			}
			db.Close()
			if count := db.GetConnectionCount(); count != 0 {
				t.Error("Connection failed to close")
			}
			removeTestDatabase(tt.path)
		})
	}
}

func TestAccessDb_ExecuteSqlStatement(t *testing.T) {
	tests := []struct {
		name       string
		statements []string
	}{
		{"Create simple table, insert one row",
			[]string{"create table TEST(ID integer, VAL text);",
				"insert into TEST values (1, 'test text');"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbPath := "test.accdb"
			reproduceTestDatabase(dbPath)
			db := New(dbPath)
			db.Open()
			if count := db.GetConnectionCount(); count != 1 {
				t.Error("No connections opened - db.Open failed")
			}
			db.BeginTransaction()
			for _, statement := range tt.statements {
				db.ExecuteSqlStatement(statement)
			}
			db.CommitTransaction()
			rows := db.GetQueryResult("select VAL from TEST;")
			var rowResult string
			for rows.Next() {
				err := rows.Scan(&rowResult)
				if err != nil {
					t.Error(err)
				}
				println(rowResult)
			}
		})
	}
}

//
//func TestAccessDb_BeginTransaction(t *testing.T) {
//	type fields struct {
//		path string
//		db   *sql.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test_assets cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			adb := &AccessDb{
//				path: tt.fields.path,
//				db:   tt.fields.db,
//			}
//		})
//	}
//}
//

//
//func TestAccessDb_CommitTransaction(t *testing.T) {
//	type fields struct {
//		path string
//		db   *sql.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test_assets cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			adb := &AccessDb{
//				path: tt.fields.path,
//				db:   tt.fields.db,
//			}
//		})
//	}
//}
//

//
//func TestAccessDb_RefreshTransaction(t *testing.T) {
//	type fields struct {
//		path string
//		db   *sql.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test_assets cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			adb := &AccessDb{
//				path: tt.fields.path,
//				db:   tt.fields.db,
//			}
//		})
//	}
//}
//
//func Test_synchronizeQueryExecution(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test_assets cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
