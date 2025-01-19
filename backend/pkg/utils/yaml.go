package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func UnmarshalYamlFile[T any](filePath string, out *T) error {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("[utils.UnmarshalYamlFile] call os.Open error %+v", err)
		return err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		logrus.Errorf("[utils.UnmarshalYamlFile] call io.ReadAll error %+v", err)
		return err
	}
	return yaml.Unmarshal(bytes, out)
}
