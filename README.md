# Apple Silicon Check

A CLI tool that checks all your apps for their Apple Silicon support.

## Installation

### Homebrew Tap

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

#### run

```shell
./ascheck -h
```
