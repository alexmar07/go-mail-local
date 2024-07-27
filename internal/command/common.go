package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	setupFile = ".setup.%s.json"
	dirSetup  = ".go-mail-local/config"
)

type Options map[string]interface{}

func existSetupFile(profile string) bool {

	path := getSetupFilePath(profile)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func createSetupFile(profile string) error {

	path := getSetupFilePath(profile)

	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	if _, err := os.Create(path); err != nil {
		return err
	}

	return nil
}

func getSetupFilePath(profile string) string {
	return getSetupDir() + "/" + getSetupFilename(profile)
}

func getSetup(profile string) (map[string]interface{}, error) {

	if !existSetupFile(profile) {
		return nil, errors.New("setup file not found; run go-mail-local [-profile=] setup")
	}

	buff, _ := os.ReadFile(getSetupFilePath(profile))

	var result map[string]interface{}

	if err := json.Unmarshal(buff, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func getSetupFilename(profile string) string {
	return fmt.Sprintf(setupFile, profile)
}

func getSetupDir() string {

	homeDir, _ := os.UserHomeDir()

	return homeDir + "/" + dirSetup
}
