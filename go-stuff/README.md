# Go-stuff

This directory contains all go-related code, including vendored code
we will use.  Either direct copies or git submodules are fine. Make
sure to edit the `.bettercodehub.yml` in the root of the repo if you
shuffle things around.

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

## Hummingbird 

### Overview
Hummingbird will in the long run be distributed across a vast
collection of nodes in the Skycoin network. It will leverage standard
tcp connection to talk to mobile devices, and route traffic using the
skywire protocol over the self-healing mesh network.

### Build 'n Run
- `cd src/github.com/milvum/hummingbird/cmd`
- `go install ./...`
- Your files should be in `bin` (which should be available in your path by direnv)

Check out the config-aux directory for some configuration
snippets. Run skywire-enabled stuff by doing:
`bin/skywire-node config-aux/bride-config-1.json`.

## Demo instructions:
Basically what is written above :-).
