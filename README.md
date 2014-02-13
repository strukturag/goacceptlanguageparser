goacceptlanguageparser
======================

[Go](http://golang.org) library for parsing HTTP request header "Accept-Language".

## Installation

```bash
$ go get github.com/strukturag/goacceptlanguageparser
```

## Run Python example

```bash
$ python python/acceptlanguageparser.py
```

## Run unit tests

```bash
$ go test -v
```

## Run benchmarks

```bash
$ go test -cpuprofile=/tmp/cpu.out -bench .
```

## View cpu profile

```bash
$ go tool pprof --text goacceptlanguageparser.test /tmp/cpu.out
```

Replace ``--text`` with ``--callgrind`` or to create a callgrind format file.

## Contributing

1. "Fork".
2. Make a feature branch.
3. Make changes.
4. Do your commits (run ``go fmt`` before commit).
5. Send "pull request".


## License

`goacceptlanguageparser` uses a BSD-style license, see our `LICENSE` file.
