package main

import (
	"archive/zip"
	"flag"
	"io"
	"os"
)

func main() {
	target := flag.String("t", "", "target file")
	dest := flag.String("d", "", "destication")
	flag.Parse()

	file, err := os.Open(*target)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipFile, err := os.Create(*dest)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileWriter, err := zipWriter.Create(file.Name())
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}
}
