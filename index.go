package store

import (
	"errors"
	"os"
)

const (
	localStore  = "local"
	webDavStore = "webdav"
	perm        = 0644
)

type StoreIFace interface {
	init(cfg Config) error
	IsExist(filePath string) bool
	CreateFile(path string, file []byte) error
	GetFile(path string) ([]byte, error)
	CreateJsonFile(path string, data interface{}) error
	ClearDir(path string) error
	GetJsonFile(path string, file interface{}) error
	State(path string) (os.FileInfo, error)
	MkdirAll(path string) error
}

type Config struct {
	StoreType  string
	WebDavHost string
	WebDavUser string
	WebDavPass string
}

func New(cfg Config) (StoreIFace, error) {
	var s StoreIFace

	switch cfg.StoreType {
	case localStore:
		s = new(Local)
	case webDavStore:
		s = new(WebDav)
	default:
		return nil, errors.New("unknown store type")
	}
	return s, s.init(cfg)
}
