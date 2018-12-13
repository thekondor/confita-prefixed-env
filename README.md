# Confita's Prefixed Environment Backend

This code provides with a backend implementation of reading environment variables for [Confita library](https://github.com/heetch/confita). In contrast to the original Confita's one, this implementation allows to prefix/vendor environment variables to mitigate naming mess and make application's configuration more clear.

# Example

In a setup similar to the following one:

`example-config.yml`:

   ```yml
   ---
   server:
     host: 127.0.0.1
     port: 8080
   ```
   
`config.go`:

   ```go
   import (
   ...
       prefixed_env "github.com/thekondor/confita-prefixed-env"
   ...
   )
   ...
   type Config struct {
      Server struct {
          Host string `config:"server_host"`
          Port uint32 `config:"server_port"`
      }
   }
   ...
   loader := confita.NewLoader(
       ...
       prefixed_env.NewDefaultBackend("MY_APP")
       ...
   )
   ```

the values of `Config.Host` and `Config.Port` could be overwritten by setting `MY_APP_SERVER_HOST` and `MY_APP_SERVER_PORT` environment variables instead of `SERVER_HOST` and `SERVER_PORT` respectively.

# Usage

`prefixed-env` backend is created using the following functions:
- `NewDefaultBackend(prefix_name)`
- `NewBackend(prefix_name, prefix_delimiter)`

where `prefix_name` is a name of a vendored prefix to be used in a variable lookup; `prefix_delimiter` is a separator between `prefix_name` and environment's variable name (default one is underscore `_`).

Example of the usafe could be found in [backend_test.go](backend_test.go).

# License

The library is released under the MIT license. See LICENSE file.
