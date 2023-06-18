package controller

import "io"

type IController interface {
	Handle(file io.Reader, lengthBatch int) error
}
