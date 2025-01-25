package utils

import (
	"os"
	"path/filepath"
)

func GetAllFileNames(dir string, extension string) ([]string, error) {
	var fileNames []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() || filepath.Ext(info.Name()) == extension {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileNames, nil
}
