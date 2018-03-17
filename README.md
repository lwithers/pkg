# Packages for Go

This repository contains small, general purpose utility packages for Go. Its
structure and content are inspired by https://github.com/pkg . As a guideline
for inclusion in this repository, packages should only depend upon the
standard library, golang.org/x packages, and other packages in this repository.

## Overview

- **stdinprompt** — prompts users that we are awaiting data on stdin.
- **writefile** — idiomatic Unix file writing (tmpfile, rename into place once
	synced, dereference symlinks).
