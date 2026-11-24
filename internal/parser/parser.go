package parser

import (
	"io/fs"
	"log"
	"os"
)

func parse_scene(name string) ([]byte, error) {
	fsys := os.DirFS(".")
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
