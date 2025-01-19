package entity

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LogFile struct {
	DirPath string
	UUID    string
}

func (f *LogFile) prepare(filePath string) (io.Writer, error) {
	err := os.MkdirAll(f.DirPath, 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

func (f *LogFile) getFileName(logType string) string {
	fileName := fmt.Sprintf("%s_%s.txt", f.UUID, logType)
	return filepath.Join(f.DirPath, fileName)
}

func (f *LogFile) GetStdoutWriter() (io.Writer, error) {
	return f.prepare(f.getFileName("stdout"))
}

func (f *LogFile) GetStderrWriter() (io.Writer, error) {
	return f.prepare(f.getFileName("stderr"))
}

func (f *LogFile) GetStdoutReader() (io.Reader, error) {
	return os.Open(f.getFileName("stdout"))
}

func (f *LogFile) GetStderrReader() (io.Reader, error) {
	return os.Open(f.getFileName("stderr"))
}
