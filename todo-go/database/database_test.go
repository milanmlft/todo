package database

import (
	"os"
	"testing"
)

func TestInitialiseDB(t *testing.T) {
	path, err := os.CreateTemp("", "tmp.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(path.Name())

	InitialiseDB(path.Name())
	result, err := os.ReadFile(path.Name())
	if err != nil {
		t.Fatalf("Failed to read from %s; %v", path.Name(), err)
	}
	resultString := string(result)
	want := "[]"
	if resultString != want {
		t.Fatalf("InitialiseDB() wrote %s; want %s", resultString, want)
	}
}

func TestInitialiseDBErrors(t *testing.T) {
	err := InitialiseDB("~/fail.json")
	if err == nil {
		t.Fatalf("InitialiseDB(\"~/fail.json\") gave nil error; want non-nill error")
	}
}
