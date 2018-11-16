# gofmt-runner -- Execute gofmt, where ever it may be

`gofmt-runner` executes the real gofmt based on the `GOROOT` of the
default Go compiler.

The intent is to make switching between Go SDKs using
https://godoc.org/golang.org/dl easier, mostly by avoiding dangling
`gofmt` symlinks to a removed SDK.

If you use golang.org/dl, you end up with multiple `goX.Y`
executables. To choose which one is the default Go version to run, the
easiest thing is to symlink one of them to `go`. The easiest way to
choose a `gofmt` version has been to make `gofmt` be a symlink
pointing to something like `/home/MYUSER/sdk/go1.11.2/bin/gofmt`.
That's just annoying, and prone to dangling symlinks as you remove old
Go versions from `~/sdk`.

Instead, make your `gofmt` executable be `gofmt-runner`, and it'll use
your default `go` command to find out which `gofmt` version to
execute. From now on, as you upgrade Go versions, you only need to
update the `go` symlink.
