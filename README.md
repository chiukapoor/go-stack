# go-stack

[![Build Status](https://github.com/chiukapoor/go-stack/actions/workflows/docker-image.yml/badge.svg)](https://github.com/chiukapoor/go-stack/actions)
[![License](https://img.shields.io/github/license/chiukapoor/go-stack)](https://github.com/chiukapoor/go-stack/blob/main/LICENSE)
[![Docker Repository](https://img.shields.io/badge/docker-repo-blue)](https://hub.docker.com/r/csociety/go-stack)

## Description

This repository contains a simple implementation of a thread-safe stack in Go. It includes concurrent push and pop operations, unit tests, and a Dockerfile for containerization.

## Features

- Concurrency support
- Dockerized
- Kubernetes-ready
- CI/CD with GitHub Actions

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker

### Running the Application

#### Locally

```bash
go run main.go
```

#### Docker

```bash
docker run -d csociety/go-stack 
```

#### Kubernetes

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: go-stack
  labels:
    app: go-stack
spec:
  containers:
  - name: go-stack
    image: csociety/go-stack:latest
```
