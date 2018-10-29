package interfaces

import "io"

type IStorageRepository interface {
	CreateFile(string, io.Reader) (string, error)
}
