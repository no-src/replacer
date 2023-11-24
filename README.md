# replacer

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/replacer/go.yml?branch=main)](https://github.com/no-src/replacer/actions)
[![License](https://img.shields.io/github/license/no-src/replacer)](https://github.com/no-src/replacer/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/replacer.svg)](https://pkg.go.dev/github.com/no-src/replacer)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/replacer)](https://goreportcard.com/report/github.com/no-src/replacer)
[![codecov](https://codecov.io/gh/no-src/replacer/graph/badge.svg?token=aJZPSpxpK7)](https://codecov.io/gh/no-src/replacer)
[![Release](https://img.shields.io/github/v/release/no-src/replacer)](https://github.com/no-src/replacer/releases)

The replacer is a configuration-based file replace tool.

## Installation

The first need [Go](https://go.dev/doc/install) installed (**version 1.21+ is required**), then you can use the below
command to install `replacer`.

```bash
go install github.com/no-src/replacer/...@latest
```

### Run In Docker

You can use the [build-docker.sh](/scripts/build-docker.sh) script to build the docker image and you should clone this
repository and `cd` to the root path of the repository first.

```bash
$ ./scripts/build-docker.sh
```

Or pull the docker image directly from [DockerHub](https://hub.docker.com/r/nosrc/replacer) with the command below.

```bash
$ docker pull nosrc/replacer
```

## Quick Start

Create a config file named `replacer.yaml`, content is as follows

```yaml
name: the configuration of replacer
version: v0.0.1
items:
  - name: replace domain
    paths:
      - "*/test.yaml"
    rules:
      - old:
          - 127.0.0.1:8888
          - 127.0.0.1:9999
        new:
          test: test-github.com
          uat: uat-github.com
          product: product-github.com
```

Now running the command below start to replace the files.

```bash
$ replacer -root="./testdata/testfile" -tag=test
```