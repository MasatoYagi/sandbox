package main

import (
	"bufio"
	"io"
)

func main() {
	results, resultWriter := io.Pipe()
	go func() {
		defer func() {
			if err := resultWriter.Close(); err != nil {
				return
			}
		}()

		var writer io.Writer
		switch fileCharset {
		case "ShiftJIS":
			writer = bufio.NewWriter(transform.NewWriter(resultWriter, japanese.ShiftJIS.NewEncoder()))
		case "UTF-8":
			writer = bufio.NewWriter(resultWriter)
		}

		if withHeader != nil && *withHeader {
			header := MeterValueCSV{}
			if err := header.WriteHeader(writer); err != nil {
				_ = resultWriter.CloseWithError(err)
				return
			}
		}

		for _, d := range rs {
			if err := d.WriteLine(writer); err != nil {
				_ = resultWriter.CloseWithError(err)
				return
			}
		}
	}()
}
