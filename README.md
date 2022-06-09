# validate-json
[![Build Docker Image](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml/badge.svg)](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml)

This action checks the validity of json files. The action is run using a Docker container and is written in Go.

## Usage
See [action.yml](action.yml)

Complete sample workflow:
```yaml
name: Validate json files

on:
  pull_request:
    branches: [ main ]

jobs:
  validate-json:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v2
      with:
        fetch-depth: 2

    - name: Validate json files
      uses: RaaLabs/validate-json@0.0.6
      with:
        directory: "."

```