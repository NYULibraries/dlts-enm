TODO: Copy/move this to the main README.md

See https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md section
"Package graph rules: required and ignored"

```bash
dep ensure
cd vendor/github.com/xo/xo
go build
cd ../../../..
mv vendor/github.com/xo/xo/xo bin/xo
```
