# validate-json
[![Build Docker Image](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml/badge.svg)](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml)

This action checks the validity of json files. The action is run using a Docker container and is written in Go.

## Usage
See [action.yml](action.yml)

Sample workflow:
```yaml
steps:
  - name: Validate json files
    uses: RaaLabs/validate-json@0.0.4
    with:
      directory: "."
```