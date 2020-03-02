package main

import (
	"flag"
	"fmt"
	"gospell"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgospell [flags] # runs on package in current directory\n")
	fmt.Fprintf(os.Stderr, "\tgospell [flags] [directories] # where a '/...' suffix includes all sub-directories\n")
	fmt.Fprintf(os.Stderr, "\tgospell [flags] [files] # all must belong to a single package\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		gospell.DirPerser("./")
	} else {
		var args []string
		var filesRun, dirsRun int
		for _, arg := range flag.Args() {
			if strings.HasSuffix(arg, "/...") && isDir(arg[:len(arg)-len("/...")]) {
				for _, dir := range allDirsInFS(arg) {
					args = append(args, dir)
				}
				dirsRun = 1
			} else if isDir(arg) {
				args = append(args, arg)
				dirsRun = 1
			} else if exists(arg) {
				filesRun = 1
				args = append(args, arg)
			} else {
				fmt.Fprintf(os.Stderr, "warning: %q matched no directories\n", arg)
			}
		}
		if len(args) == 0 || filesRun+dirsRun != 1 {
			usage()
			os.Exit(2)
		} else if filesRun == 1 {
			gospell.FilesPerser(args)
		} else {
			gospell.DirsPerser(args)
		}
	}
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func allDirsInFS(dirPath string) []string {
	i := strings.Index(dirPath, "...")
	dir, _ := path.Split(dirPath[:i])

	var dirs []string
	filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		dirs = append(dirs, path)
		return nil
	})

	return dirs
}
