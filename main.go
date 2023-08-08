/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package main

import (
	"embed"

	"github.com/yanosea/mindnum/cmd"
)

//go:embed resources/*
var embedded embed.FS

func main() {
	cmd.Execute(embedded)
}
