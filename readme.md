[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/poll)](https://goreportcard.com/report/github.com/wayneashleyberry/poll)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/poll?status.svg)](https://godoc.org/github.com/wayneashleyberry/poll)
[![wercker status](https://app.wercker.com/status/ef0020b87e80a613dcd4b275105ee2d4/s/master "wercker status")](https://app.wercker.com/project/byKey/ef0020b87e80a613dcd4b275105ee2d4)

> A command-line application to poll other command-line applications.

### Installation

##### You have a working [Go environment](https://golang.org/doc/install).

```sh
go get github.com/wayneashleyberry/poll
```

### Usage

```
poll [flags] [command]
```

#### Examples

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
