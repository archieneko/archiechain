package main

import (
	_ "embed"

	"github.com/archieneko/archiechain/command/root"
	"github.com/archieneko/archiechain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
