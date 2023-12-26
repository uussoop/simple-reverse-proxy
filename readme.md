# simple reverse proxy Configuration

This README describes the proxy configuration for the application.

## Overview

The application uses a proxy configuration defined in `config.yaml` to route requests. 

## Config File

The `config.yaml` file defines a list of proxy configurations:

```yaml
proxies:
  - name: example
    from:
      scheme: "https://"
      host: "icanhazip.com"
    to: 
      scheme: "https://"
      host: "icanhazip.com"
```

Each proxy has a `name` and defines a `from` and `to` section:

- `from` defines where requests will come from
- `to` defines where requests will be proxied to

## Program Flow

The main.go file reads the config and initializes the proxy router:

```go
func main() {

  config, err := configreader.ReadConfig("config.yaml")
  if err != nil {
    panic(err) 
  }
  
  InitRouter(config)

}
```

The `InitRouter` function takes the config and initializes the router with the proxy rules.

