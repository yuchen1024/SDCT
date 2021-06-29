package utils

import (
	"io/ioutil"
	"os"
)

// WriteToFile write the data to file. if the file doesn't exist,
// it will create a new file. if the file exists, it will cover
// the data.
func WriteToFile(data []byte, path string) {
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		panic(err)
	}
}

// Delete tries to delete a file.
func Delete(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	if err := os.Remove(path); err != nil {
		panic(err)
	}
}

// Read returns byte of a file.n
func Read(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	return raw
}

// Write writes bytes to certain file.
func Write(path string, data []byte) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		Delete(path)
	}

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		panic(err)
	}
}
