# go-store
os vs webdav


##### Интерфейс для работы с файлами
Переменная **STORE_TYPE** определяет с каким хранилищем работает сервис - webdav либо локальная директория
```go
type StoreIFace interface {
	Init(cfg Config) error
	IsExist(filePath string) bool
	CreateFile(path string, file []byte) error
	GetFile(path string) ([]byte, error)
	CreateJsonFile(path string, data interface{}) error
	GetJsonFile(path string, file interface{}) error
	MkdirAll(path string) error
}
```