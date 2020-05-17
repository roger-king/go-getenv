# go-getenv

A simple getenv utility to support default values to be set. This is a port from python's `os.getenv("ENV_KEY", "default)`. This utility function will do a few things:

1. Return the passed default value (with a logging message).
2. Return the environment variable set over the default value
3. Panic if the Environment variable does not exist and no default value set.

# Usage

```
  import "github.com/roger-king/go-getenv/getenv"

  port := getenv.EnvOrDefault("PORT", ":5000")
```
