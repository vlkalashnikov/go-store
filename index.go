package store

import (
	"errors"
	"os"
)

const (
	localStore  = "local"
	webDavStore = "webdav"
	s3Store     = "s3"
	empty       = "empty"
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
	S3Region   string
	S3Bucket   string
	S3Access   string
	S3Secret   string
}

func New(cfg Config) (StoreIFace, error) {
	switch cfg.StoreType {
	case localStore:
		return NewLocal(cfg)
	case webDavStore:
		return NewWebDav(cfg)
	case s3Store:
		return NewS3(cfg)
	case empty:
		return NewEmpty(cfg)
	default:
		return nil, errors.New("unknown store type")
	}
}

func NewEmpty(cfg Config) (StoreIFace, error) {
	s := new(Empty)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewLocal(cfg Config) (StoreIFace, error) {
	s := new(Local)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewWebDav(cfg Config) (StoreIFace, error) {
	s := new(WebDav)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewS3(cfg Config) (StoreIFace, error) {
	s := new(S3)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}
