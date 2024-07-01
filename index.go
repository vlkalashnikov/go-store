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
	init(Config) error
	IsExist(string) bool
	CreateFile(string, []byte) error
	GetFile(path string) ([]byte, error)
	RemoveFile(path string) error
	CreateJsonFile(string, interface{}) error
	ClearDir(string) error
	GetJsonFile(string, interface{}) error
	State(string) (os.FileInfo, error)
	MkdirAll(string) error
}

type Config struct {
	StoreType     string
	WebDavHost    string
	WebDavUser    string
	WebDavPass    string
	S3Region      string
	S3Bucket      string
	S3AccessKeyID string
	S3AccessKey   string
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
