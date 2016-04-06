package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/moqada/pjsval"
)

const (
	version = "0.0.1"
)

var (
	fp  = kingpin.Arg("file", "Path of JSON Schema").File()
	op  = kingpin.Flag("output", "Path of Go struct file").Short('o').String()
	pkg = kingpin.Flag("package", "Package name for Go validator file").Default("main").Short('p').String()
)

func init() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(version)
	kingpin.Parse()
}

func exec() error {
	var err error
	var out io.Writer
	in := *fp
	defer func() {
		(in).Close()
	}()
	if in == nil {
		info, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("File does not exists: %s", err)
		}
		if info.Size() == 0 {
			kingpin.Usage()
			return err
		}
		in = os.Stdin
	}
	if *op == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*op)
		if err != nil {
			return err
		}
	}
	if err := pjsval.Generate(in, out, *pkg); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := exec(); err != nil {
		kingpin.Errorf("%s", err)
		os.Exit(1)
	}
}
