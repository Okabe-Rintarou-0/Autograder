package entity

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LogDir struct {
	DirPath string
	UUID    string
}

func (d *LogDir) prepare(filePath string) (io.WriteCloser, error) {
	err := os.MkdirAll(d.DirPath, 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

func (d *LogDir) getFileName(logType string) string {
	fileName := fmt.Sprintf("%s_%s.txt", d.UUID, logType)
	return filepath.Join(d.DirPath, fileName)
}

func (d *LogDir) GetWriter(logType string) (io.WriteCloser, error) {
	filePath := d.getFileName(logType)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if os.IsNotExist(err) {
		return d.prepare(filePath)
	}
	return file, nil
}

func (d *LogDir) GetReader(logType string) (io.ReadCloser, error) {
	return os.Open(d.getFileName(logType))
}
