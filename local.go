package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Local struct {
}

func (l *Local) init(cfg Config) error {
	return nil
}

func (l *Local) IsExist(filePath string) bool {
	info, err := os.Stat(filePath)
	return err == nil && info.Size() > 0
}

func (l *Local) CreateFile(path string, file []byte) error {
	return os.WriteFile(path, file, perm)
}

func (l *Local) GetFile(path string) ([]byte, error) {
	if !l.IsExist(path) {
		return nil, nil
	}
	return os.ReadFile(path)
}

func (l *Local) RemoveFile(path string) error {
	return os.Remove(path)
}

// State can return default value
func (l *Local) State(path string) (os.FileInfo, error) {
	if !l.IsExist(path) {
		return nil, nil
	}
	return os.Stat(path)
}

func (l *Local) ClearDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Local) MkdirAll(path string) error {
	return os.MkdirAll(path, perm)
}

func (l *Local) CreateJsonFile(path string, data interface{}) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return l.CreateFile(path, content)
}

func (l *Local) GetJsonFile(path string, file interface{}) error {
	content, err := l.GetFile(path)
	if err != nil {
		return err
	}
	if content == nil {
		return nil
	}
	return json.Unmarshal(content, file)
}
