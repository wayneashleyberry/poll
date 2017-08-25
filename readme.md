[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/poll)](https://goreportcard.com/report/github.com/wayneashleyberry/poll)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/poll?status.svg)](https://godoc.org/github.com/wayneashleyberry/poll)

> A command-line application to poll other command-line applications.

### Installation

```sh
go get github.com/wayneashleyberry/poll
```

### Usage

```
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
