package iparser

import (
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type InputStream struct {
	File string
	*antlr.InputStream
}

func NewIoStream(fileName string) (is *InputStream, err error) {
	var fd *os.File
	fd, err = os.Open(fileName)
	if err != nil {
		return
	}

	defer func() {
		err = fd.Close()
	}()

	innerIs := antlr.NewIoStream(fd)

	return &InputStream{
		File:        fileName,
		InputStream: innerIs,
	}, nil
}
