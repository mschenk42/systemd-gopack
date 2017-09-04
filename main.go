package main

import (
	"github.com/mschenk42/gopack"
	"github.com/mschenk42/systemd-runpack/systemd"
)

func main() {
	systemd.Run(gopack.LoadProperties())
}