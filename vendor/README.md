# Vendor

This directory contains all vendored code we will use.  Either direct
copies or git submodules are fine.

*TODO* Make sure BetterCodeHub does not punish us for anything in this
directory.

## Skywire

As we will be using skywire, it makes sense to have our own (stable)
version of skywire that we can work with, and if needed, hack
on. Skywire is built using go (v1.9 or higher).

Requirements:
- [Go 1.9 or newer](https://golang.org/doc/install)
- [direnv](https://direnv.net/)

Steps to build:
- Allow direnv settings with `direnv allow`
- `cd src/github.com/skycoin/skywire/cmd`
- `go install ./...`
- Your files should be in `bin` (which should be available in your path by direnv)
