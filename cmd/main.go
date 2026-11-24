package main

import (
	"log"

	p "github.com/LinerGit/cub3d-go/internal/parser"
)

const winWidth, winHeight = 960, 600
const texWidth, texHeight = 64, 64

type texture struct {
	pixels      []byte
	w, h, pitch int
}

type rgba struct {
	r, g, b, a byte
}

func main() {
	data, err := p.ParseScene("scene.cub")
	if err != nil {
		log.Fatal(err)
	}
	println(string(data))
}
