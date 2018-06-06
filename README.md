# Packages for Go

[![GoDoc](https://godoc.org/github.com/lwithers/pkg?status.svg)](https://godoc.org/github.com/lwithers/pkg)

This repository contains small, general purpose utility packages for Go. Its
structure and content are inspired by https://github.com/pkg . As a guideline
for inclusion in this repository, packages should only depend upon the
standard library, golang.org/x packages, and other packages in this repository.

## Overview

- **stdinprompt** — prompts users that we are awaiting data on stdin.
- **writefile** — idiomatic Unix file writing (tmpfile, rename into place once
	synced, dereference symlinks).
- **byteio** — byte-oriented I/O adapters, and simple binary read/write
	functions. Useful if you are reading and writing lots of small values.
- **versionsort** — simple version sort akin to C's versionsort(3).
