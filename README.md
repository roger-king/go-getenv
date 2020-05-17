# go-getenv

A simple getenv utility to support default values to be set. This is a port from python's `os.getenv("ENV_KEY", "default)`. This utility function will do a few things:

1. Return the passed default value (with a logging message).
2. Return the environment variable set over the default value
3. Panic if the Environment variable does not exist and no default value set.

# Installation

Installation is done using the go get command:

```
  go get -u github.com/roger-king/go-getenv
```

# Usage

Using the utility function is simple. We call the function and pass in as the first arg the `Environment Variable Key` and the `getenv.String("some-string")`.

```
  import "github.com/roger-king/go-getenv"

  port := getenv.EnvOrDefault("PORT", getenv.String(":5000"))
```

## Why `getenv.String("some-string")`

I decided to create a `String()` function because I want to support the ability to pass `nil` as the second or default argument.
