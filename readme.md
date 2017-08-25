[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/poll)](https://goreportcard.com/report/github.com/wayneashleyberry/poll)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/poll?status.svg)](https://godoc.org/github.com/wayneashleyberry/poll)

> A stupid command-line application to poll other command-line applications.

```sh
go get github.com/wayneashleyberry/poll
poll kubectl get pods
```

The default polling interval is `1s`, you can change this with a flag. Any
format supported by the
`[time.ParseDuration](https://golang.org/pkg/time/#ParseDuration)` method will
work.

```sh
poll -i 10s date
```
