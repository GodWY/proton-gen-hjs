package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/GodWY/proton-gen-hjs/internal_genhi"
	"github.com/GodWY/proton-gen-hjs/version"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-js: plugins are not supported; use ")
		}
		for _, f := range gen.Files {
			if f.Generate {
				internal_genhi.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = internal_genhi.SupportedFeatures
		return nil
	})
}
