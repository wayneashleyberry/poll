[![Go Reference](https://pkg.go.dev/badge/wayneashleyberry/poll.svg)](https://pkg.go.dev/wayneashleyberry/poll)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/poll)](https://goreportcard.com/report/github.com/wayneashleyberry/poll)
[![golangci-lint](https://github.com/wayneashleyberry/poll/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/wayneashleyberry/poll/actions/workflows/golangci-lint.yml)
[![go](https://github.com/wayneashleyberry/poll/actions/workflows/go.yml/badge.svg)](https://github.com/wayneashleyberry/poll/actions/workflows/go.yml)

> A command-line application to poll other command-line applications.

### Installation

##### You have a working [Go environment](https://golang.org/doc/install).

```sh
go install github.com/wayneashleyberry/poll@latest
```

### Usage

```sh
poll [flags] [command]
```

#### Examples

```sh
poll kubectl get pods
```

#### Flags

##### Polling Interval

Any format supported by the [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration) method will work.

```sh
poll -i 10s date
```

```sh
poll -i 50ms date +%s%3N
```
