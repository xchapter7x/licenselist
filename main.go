package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ryanuber/go-license"
)

type visitor struct {
	prefix string
}

func (s *visitor) visit(visitpath string, f os.FileInfo, err error) error {
	if strings.HasPrefix(path.Base(visitpath), s.prefix) {

		l, err := license.NewFromFile(visitpath)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("License: %s \t found in Lib: %s\n", l.Type, visitpath)
	}
	return nil
}

func main() {

	flag.Parse()
	rootDir := flag.Arg(0)
	filenamePrefix := flag.Arg(1)
	v := &visitor{
		prefix: filenamePrefix,
	}
	err := filepath.Walk(rootDir, v.visit)
	fmt.Println("any problems?:", err)
}
