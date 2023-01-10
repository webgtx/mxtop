package msg

import (
	"fmt"
	"log"
	"os"

	"github.com/webgtx/mxtop/internal/info"
)

func Log(s string, a ...any) {
	if info.Debug == "ON" {
		log.Printf(s, a...)
	}
}

func ExitLog(err error) {
	if info.Debug == "ON" {
		log.Fatalf("err: %v", err)
	}
	fmt.Printf("err: %v\n", err)
	os.Exit(1)
}
