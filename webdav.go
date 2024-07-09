package store

import (
	"encoding/json"
	"os"

	"github.com/studio-b12/gowebdav"
)

type WebDav struct {
	client *gowebdav.Client
}

func (w *WebDav) init(cfg WebDavConfig) error {
	w.client = gowebdav.NewClient(cfg.WebDavHost, cfg.WebDavUser, cfg.WebDavPass)
	return nil
}

func (w *WebDav) IsExist(filePath string) bool {
	info, err := w.client.Stat(filePath)
	return err == nil && info.Size() > 0
}

func (w *WebDav) CreateFile(path string, file []byte) error {
	return w.client.Write(path, file, perm)
}

func (w *WebDav) GetFile(path string) ([]byte, error) {
	if !w.IsExist(path) {
		return nil, nil
	}
	return w.client.Read(path)
}

func (w *WebDav) RemoveFile(path string) error {
	return w.client.Remove(path)
}

// State can return default value
func (w *WebDav) State(path string) (os.FileInfo, error) {
	if !w.IsExist(path) {
		return nil, nil
	}
	return w.client.Stat(path)
}

func (w *WebDav) ClearDir(path string) error {
	files, _ := w.client.ReadDir(path)
	for _, file := range files {
		if err := w.client.Remove(path + "/" + file.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (w *WebDav) MkdirAll(path string) error {
	return w.client.MkdirAll(path, perm)
}

func (w *WebDav) CreateJsonFile(path string, data interface{}) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return w.CreateFile(path, content)
}

func (w *WebDav) GetJsonFile(path string, file interface{}) error {
	content, err := w.GetFile(path)
	if err != nil {
		return err
	}
	if content == nil {
		return nil
	}
	return json.Unmarshal(content, file)
}
