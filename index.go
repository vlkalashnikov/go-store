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

type StoreConfigIFace interface {
	S3Config | WebDavConfig | EmptyConfig | LocalConfig
}

type StoreIFace interface {
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
	StoreType string
	EmptyConfig
	LocalConfig
	WebDavConfig
	S3Config
}

type S3Config struct {
	S3Region      string
	S3Bucket      string
	S3AccessKeyID string
	S3AccessKey   string
	S3Token       string
}

type WebDavConfig struct {
	WebDavHost string
	WebDavUser string
	WebDavPass string
}

type EmptyConfig struct{}
type LocalConfig struct{}

func New(cfg Config) (StoreIFace, error) {
	switch cfg.StoreType {
	case localStore:
		return NewLocal(cfg.LocalConfig)
	case webDavStore:
		return NewWebDav(cfg.WebDavConfig)
	case s3Store:
		return NewS3(cfg.S3Config)
	case empty:
		return NewEmpty(cfg.EmptyConfig)
	default:
		return nil, errors.New("unknown store type")
	}
}

func NewEmpty(cfg EmptyConfig) (StoreIFace, error) {
	s := new(Empty)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewLocal(cfg LocalConfig) (StoreIFace, error) {
	s := new(Local)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewWebDav(cfg WebDavConfig) (StoreIFace, error) {
	s := new(WebDav)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func NewS3(cfg S3Config) (StoreIFace, error) {
	s := new(S3)
	if err := s.init(cfg); err != nil {
		return nil, err
	}
	return s, nil
}
