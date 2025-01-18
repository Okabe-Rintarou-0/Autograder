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

func (f *LogFile) prepare(fileName string) (io.Writer, error) {
	err := os.MkdirAll(f.DirPath, 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(f.DirPath, fileName))
}

func (f *LogFile) GetStdoutWriter() (io.Writer, error) {
	return f.prepare(fmt.Sprintf("%s_stdout.txt", f.UUID))
}

func (f *LogFile) GetStderrWriter() (io.Writer, error) {
	return f.prepare(fmt.Sprintf("%s_stderr.txt", f.UUID))
}
