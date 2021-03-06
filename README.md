# Apple Silicon Check

[![Actions Status](https://github.com/harryzcy/ascheck/workflows/CI/badge.svg)](https://github.com/harryzcy/ascheck/actions)
[![codecov](https://codecov.io/gh/harryzcy/ascheck/branch/main/graph/badge.svg)](https://codecov.io/gh/harryzcy/ascheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/harryzcy/ascheck)](https://goreportcard.com/report/github.com/harryzcy/ascheck)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)

A CLI tool that bulk-checks your apps for the Apple Silicon support.

---

## Table of Contents

- [Installation](#installation)
  - [Homebrew tap](#homebrew-tap)
  - [go install](#go-install)
  - [Compile from source](#compile-from-source)
- [Example Usage](#example-usage)
  - [Show help](#show-help)
  - [Run](#run)
  - [Output](#output)

---

## Installation

### Homebrew tap

```shell
brew tap harryzcy/ascheck
brew install ascheck
```

### go install

```shell
go install github.com/harryzcy/ascheck
```

### Compile from source

#### clone

```shell
git clone https://github.com/harryzheng/ascheck
cd ascheck
```

#### get the dependencies

```shell
go mod tidy
```

#### build

```shell
go build -o ascheck .
```

## Example Usage

### Show help

```shell
ascheck -h
```

### Run

```shell
ascheck
```

### Output

The output will show:

```shell
NAME        CURRENT ARCHITECTURES  ARM SUPPORT
------------------------------------------------
App Store   Intel 64               Supported
Automator   Intel 64               Supported
...
```

- NAME: name of the app
- CURRENT ARCHITECTURES: the architecture of the currently installed version
- ARM SUPPORT: the arm support information on [Does it Arm](https://github.com/ThatGuySam/doesitarm)
