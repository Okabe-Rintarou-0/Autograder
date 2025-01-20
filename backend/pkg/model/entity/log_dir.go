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
	return d.prepare(d.getFileName(logType))
}

func (d *LogDir) GetReader(logType string) (io.ReadCloser, error) {
	return os.Open(d.getFileName(logType))
}
