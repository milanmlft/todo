package database

import (
	"encoding/json"
	"os"
)

func InitialiseDB(path string) error {
	b, err := json.Marshal("[]")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
