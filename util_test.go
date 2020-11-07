package main

import (
	"os"
	"testing"
)

func TestGenerateResultFile(t *testing.T) {
	type args struct {
		dbFilePath string
	}
	tests := []struct {
		name     string
		fileName string
	}{
		{"File name only (no path)", "db.accdb"},
		{"Same directory file", "db.accdb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateResultFile(tt.fileName)
			if _, err := os.Stat(tt.fileName); err == nil {
			} else {
				t.Error(err)
			}
		})
		err := os.Remove(tt.fileName)
		if err != nil {
			panic(err)
		}
	}
}
