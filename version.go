package app

import "fmt"

const (
	VersionName = "app"

	VersionNumber = "v0.0.1"

	Website = "https://zlab.dev"

	// http://patorjk.com/software/taag/#p=display&h=0&f=Small%20Slant&t=RPC
	banner = `
  ____   __    ___    ___
 /_  /  / /   / _ |  / _ )
  / /_ / /__ / __ | / _  |
 /___//____//_/ |_|/____/  %s %s

High performance, App framework
Support by %s
_____________________________________________
%s
____________________________________O/_______
                                    O\
`
)

func Banner(message string) {
	fmt.Printf(banner, VersionName, VersionNumber, Website, message)
}
