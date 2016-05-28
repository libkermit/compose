# Libkermit
[![GoDoc](https://godoc.org/github.com/libkermit/compose?status.png)](https://godoc.org/github.com/libkermit/compose)
[![Build Status](https://travis-ci.org/libkermit/compose.svg?branch=master)](https://travis-ci.org/libkermit/compose)
[![Go Report Card](https://goreportcard.com/badge/github.com/libkermit/compose)](https://goreportcard.com/report/github.com/libkermit/compose)
[![License](https://img.shields.io/github/license/libkermit/compose.svg)]()


> When green is all there is to be<br/>
> It could make you wonder why<br/>
> But why wonder why wonder<br/>
> I am green, and it'll do fine<br/>
> It's beautiful,<br/>
> and I think it's what I want to be.<br/>
> -- Kermit the Frog

When [Docker](https://github.com/docker/docker) meets with
integration/acceptance tests to make you see everything in
green. **Libkermit** is a Go(lang) library that aims to ease the
writing of integration tests (any non unit tests actually) with the
helps of Docker and it's ecosystem (mainly
[libcompose](https://github.com/docker/libcompose)).

The goals are :

- Easy docker manipulation, from managing a simple container to boot
  up a whole stack.
    - create, delete, pause, … containers
    - check for a certain state containers (inspect them)
    - support *compose files* to allow starting a whole stack
- Testing suite and functions, in a simple fashion.
- Works seamlessly with the Go(lang) `testing` framework.
- Try to not force any testing framework but also tries to integrate
  with them ([go-check](https://github.com/go-check/check),
  [testify](https://github.com/stretchr/testify), …).

**Note: This is experimental and not even implemented yet. You are on your own right now**


## Package `compose`

This package holds functions and structs to ease docker uses.

```go
package yours

import (
    "testing"

    "github.com/libkermit/docker/compose"
)

func TestItMyFriend(t *testing.T) {
    project, err := compose.CreateProject("simple", "./assets/simple.yml")
    if err != nil {
        t.Fatal(err)
    }
    err = project.Start()
	if err != nil {
		t.Fatal(err)
	}

    // Do your stuff

    err = project.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
```

### Package `compose/testing`

This package map the `compose` package but takes a `*testing.T` struct
on all methods. The idea is to write even less. Let's write the same
example as above.


```go
package yours

import (
    "testing"

    docker "github.com/libkermit/docker/compose/testing"
)

func TestItMyFriend(t *testing.T) {
    project := compose.CreateProject(t, "simple", "./assets/simple.yml")
    project.Start(t)

    // Do your stuff

    project.Stop(t)
}
```


## Other packages to come

- `suite` : functions and structs to setup tests suites.


