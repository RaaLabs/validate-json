# validate-json

This action checks the validity of json files.

## Usage
See [action.yml](action.yml)

Sample workflow:
```yaml
steps:
  - name: Validate json files
    uses: RaaLabs/validate-json@0.0.2
    with:
      directory: "."
```