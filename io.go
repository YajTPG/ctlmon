package main

import (
	"os"
	"strings"
)

// const CacheDir = "/tmp/"
// const CacheFile = CacheDir + ".ctlmon-cache.tmp"

func MakeDir(path string) error {
	return os.Mkdir(path, 0755)
}

func LoadTempFile(path string) []string {
	data, err := ReadTempFile(path)
	if err != nil {
		logger.Error(err)
	}
	return strings.Split(string(data), ";")
}

func AppendTempFile(path string, arr []string) error {
	_, err := ReadTempFile(path)
	if err != nil {
		logger.Error(err)
	}
	err = os.WriteFile(path, []byte(strings.Join(arr, ";")), 0755	)
	return err
}

func ReadTempFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		if err == os.ErrNotExist {
			_, err = os.Create(path)
		}
	}
	return file, err
}
