package store

import "os"

type Empty struct {
}

func (l *Empty) init(cfg EmptyConfig) error {
	return nil
}

func (l *Empty) IsExist(filePath string) bool {
	return false
}

func (l *Empty) CreateFile(path string, file []byte) error {
	return nil
}

func (l *Empty) RemoveFile(path string) error {
	return nil
}

func (l *Empty) GetFile(path string) ([]byte, error) {
	return nil, nil
}

func (l *Empty) GetFilePartially(path string, offset, length int64) ([]byte, error) {
	return nil, nil
}

func (l *Empty) Stat(path string) (os.FileInfo, error) {
	return nil, nil
}

func (l *Empty) ClearDir(dir string) error {
	return nil
}

func (l *Empty) MkdirAll(path string) error {
	return nil
}

func (l *Empty) CreateJsonFile(path string, data interface{}) error {
	return nil
}

func (l *Empty) GetJsonFile(path string, file interface{}) error {
	return nil
}
