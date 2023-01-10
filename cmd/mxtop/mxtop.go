package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/webgtx/mxtop/internal/conf"
	"github.com/webgtx/mxtop/internal/info"
	"github.com/webgtx/mxtop/internal/msg"
)

var (
	versionFlag     bool
	fetchFlag       bool
	interactiveFlag bool
)

func init() {
	flag.Usage = func() {
		fmt.Println()
		fmt.Println("   mxtop - top for your server park")
		fmt.Println("  ╭──────────────────────────────────────────────────╮")
		fmt.Println("  │ flag             │                   description │")
		fmt.Println("  │──────────────────────────────────────────────────│")
		fmt.Println("  │ -i --interactive │ run mxtop in interactive mode │")
		fmt.Println("  │ -f --fetch       │ fetch info from server        │")
		fmt.Println("  │ -v --version     │ get info about version        │")
		fmt.Println("  │ -h --help        │ get help message              │")
		fmt.Println("  ╰──────────────────────────────────────────────────╯")
		fmt.Println()
	}

	flag.BoolVar(&versionFlag, "v", false, "get info about version")
	flag.BoolVar(&versionFlag, "version", false, "get info about version")
	flag.BoolVar(&fetchFlag, "f", false, "fetch info from server")
	flag.BoolVar(&fetchFlag, "fetch", false, "fetch info from server")
	flag.BoolVar(&interactiveFlag, "i", false, "run mxtop in interactive mode")
	flag.BoolVar(&interactiveFlag, "interactive", false, "run mxtop in interactive mode")

	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Println("   __  __   ___                         mxtop")
		fmt.Println("  |  |/  `.'   `.                      ╭───────────────────────────────────────────╮")
		fmt.Println("  |   .-.  .-.   '                     │ Authors                                   │")
		fmt.Println("  |  |  |  |  |  | ____     _____      │  ├── webgtx (https://github.com/webgtx)   │")
		fmt.Println("  |  |  |  |  |  |`.   \\  .'    /      │  └── ssleert (https://github.com/ssleert) │")
		fmt.Println("  |  |  |  |  |  |  `.  `'    .'       │ License - BSD 3-Clause                    │")
		fmt.Println("  |  |  |  |  |  |    '.    .'         │ Language - Go                             │")
		fmt.Println("  |__|  |__|  |__|    .'     `.        │ Repo - https://github.com/webgtx/mxtop    │")
		fmt.Printf("                    .'  .'`.   `.      │ Commit - %s                          │\n", info.CommitHash)
		fmt.Printf("                  .'   /    `.   `.    │ Version - %s                           │\n", info.Version)
		fmt.Println("                 '----'       '----'   ╰───────────────────────────────────────────╯")
	} else if fetchFlag {
		err := conf.Parse()
		if err != nil {
			msg.ExitLog(err)
		}
		fmt.Println(conf.AnsibleDir)
	} else if interactiveFlag {
		msg.ExitLog(errors.New("not implemented"))
	}
}
