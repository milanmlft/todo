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
	err = os.WriteFile(path, encoding, 0644)
	return err
}
