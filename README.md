# validate-json
[![Build Docker Image](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml/badge.svg)](https://github.com/RaaLabs/validate-json/actions/workflows/build.yml)

This action validates json files. The action scans the repository for files with the .json extension, and tries to unmarshall them. If there are unmarshalling errors, the action terminates and your workflow will break. It will log all errors encountered before exiting.

The action is run using a Docker container and is written in Go.

> **Warning**
> We are switching from Docker Hub to GitHub container registry to store the images for this action (Docker sunsets free plans for teams). Use version >=1.0.0 going forward as the previous versions will stop working once Docker removes the images from Docker hub (pull rate limits apply from 14.April 2023, removal on 14.May 2023)


## Inputs
- `directory`: (string, optional) The path on which to scan .json files, defaults to `.`, in which case it scans the whole repository.


## Usage
Complete sample workflow. For mor information see the [action.yml](action.yml), and the GitHub help documentation here https://docs.github.com/en/actions/using-workflows#creating-a-workflow-file
```yaml
name: Validate json files

on:
  pull_request:
    branches: [ main ]

jobs:
  validate-json:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3
      with:
        fetch-depth: 2

    - name: Validate json files
      uses: RaaLabs/validate-json@v1.0.1
      with:
        directory: "."

```
