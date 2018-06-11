# gorogue
[![GoDoc](https://godoc.org/github.com/sbrow/gorogue?status.svg)](https://godoc.org/github.com/sbrow/gorogue) [![Build Status](https://travis-ci.org/sbrow/gorogue.svg?branch=master)](https://travis-ci.org/sbrow/gorogue) [![Coverage Status](https://coveralls.io/repos/github/sbrow/gorogue/badge.svg?branch=master)](https://coveralls.io/github/sbrow/gorogue?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sbrow/gorogue)](https://goreportcard.com/report/github.com/sbrow/gorogue)

Package gorogue is a flexible roguelike engine written in Go. Gorogue aims to be
small, versatile, and modular.


### Game Modes

One of this project's goals is to support a wide variety of game modes. However,
emphasis is first placed on the stability and thorough documentation of the
existing modes.

Currently, there are two game modes: online and local. Both support one Action
per Actor per tick.

Planned modes include:

    - "Realtime"	 : 30 ticks per second, Actors that don't act in time are skipped.
    - "Squad Based"	 : Allows each player to control more than one Actor."
    - "Action Points": Characters spend AP to perform actions, AP refreshes each tick.

## Installation
```bash
$ go get -u github.com/sbrow/gorogue
```
