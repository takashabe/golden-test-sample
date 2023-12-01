package main

import (
	"bytes"
	"io"

	"github.com/xuri/excelize/v2"
)

func main() {
	reader, err := genXlsx()
	if err != nil {
		panic(err)
	}

	// do something...
	_ = reader
}

func genXlsx() (io.Reader, error) {
	f := excelize.NewFile()
	if err := f.SetCellValue("Sheet1", "A1", "Hello world!"); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}
