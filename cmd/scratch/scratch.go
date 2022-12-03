package main

import (
	"fmt"
	"embed"
	"github.com/jaredpar/advent2022/util"
)

//go:embed input.txt
var f embed.FS

func main() {
	lines := util.MustReadLines(f, "input.txt")
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}



