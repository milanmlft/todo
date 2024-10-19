package database

import (
	"encoding/json"
	"os"
)

func InitialiseDB(path string) error {
	encoding, err := json.Marshal([]string{})
	if err != nil {
		return err
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
