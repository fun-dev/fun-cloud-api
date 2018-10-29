package interfaces

import "io"

type IStorageService interface {
	CreateFile(string, io.Reader) error
}
