# Parmesan

API wrapper for v3 of the Bittrex API

# Motivation
v1.1 of the Bittrex API will no longer be supported after 30/9/2020, and I want a go wrapper for their API.

You can see the Bittrex API documentation [here](https://bittrex.github.io/api/v3)

:warning: This repo is in development :warning:

Most features are missing at the time of writing, use what exists at your own risk. There may be breaking changes.

# Installation

`go get github.com/elliotcubit/parmesan`

# Usage
```
import (
  "github.com/elliotcubit/parmesan"
)
```

Code examples will live in the `examples` directory.

# Features

:x: `/account`

:x: `/addresses`

:x: `/balances`

:x: `/conditional-orders`

:x: `/currencies`

:x: `/deposits`

:heavy_check_mark: `/markets`

:x: `/orders`

:x: `/ping`

:x: `/subaccounts`

:x: `/transfers`

:x: `/withdrawals`

:x: `websocket streams`

# Contribute

If you'd like to contribute to this project, please be sure to `go fmt` your code before submitting a pull request.
