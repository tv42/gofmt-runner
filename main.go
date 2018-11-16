// Command gofmt-runner executes the real gofmt based on the GOROOT
// of the default Go compiler.
//
// The intent is to make switching between Go SDKs using
// https://godoc.org/golang.org/dl easier, mostly by avoiding dangling
// gofmt symlinks to a removed SDK.
package main // import "github.com/tv42/gofmt-runner"

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func goroot() (string, error) {
	cmd := exec.Command("go", "env", "GOROOT")
	buf, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("go env: %v", err)
	}
	if bytes.IndexByte(buf, '\n') != len(buf)-1 {
		return "", fmt.Errorf("expected single line from go env: %q", buf)
	}
	buf = buf[:len(buf)-1]
	return string(buf), nil
}

func run() error {
	dir, err := goroot()
	if err != nil {
		return fmt.Errorf("cannot determine GOROOT: %v", err)
	}
	p := filepath.Join(dir, "bin/gofmt")
	if err := unix.Exec(p, os.Args, os.Environ()); err != nil {
		return fmt.Errorf("cannot execute gofmt: %v", err)
	}
	panic("exec didn't work")
}

const prog = "gofmt-runner"

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}
