package main

import (
	"fmt"

	"github.com/webgtx/mxtop/internal/info"
)

func main() {
	fmt.Println("ver    -", info.Version)
	fmt.Println("commit -", info.CommitHash)
	fmt.Println("debug  -", info.Debug)
}
