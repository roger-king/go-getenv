# go-getenv

A simple getenv utility to support default values to be set. This is a port from python's `os.getenv("ENV_KEY", "default)`.

# Usage


```
  import "github.com/roger-king/go-getenv/getenv"

  port := getenv.EnvOrDefault("PORT", ":5000")
```
