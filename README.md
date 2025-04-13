# pathlist
Snag words from URL paths to build quick wordlists for fuzzing

## Installation:

Using go:

```sh
go install github.com/ayushkr12/pathlist@latest
```

## Usage:

```sh
$ echo "https://example.com/api/example_admin_pannel/api/admin" | pathlist
api
example_admin_pannel
admin
```
