package main

import (
	"fmt"
	"os"
)

func main() {
	read("./test.txt")
	write("./aaa.txt")
}

func read(path string) error {
	sf, err := os.Open(path)
	if err != nil {
		return err
	}
	fmt.Print(sf)
	defer sf.Close()
	return nil
}

func write(path string) error {
	df, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() error {
		if err := df.Close(); err != nil {
			return err
		}
		return nil
	}()
	return nil
}
