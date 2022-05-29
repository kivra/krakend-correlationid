# CorrelationID Middleware

This package contains custom middleware to add correlation IDs to the request
headers of incoming requests.

## Installation

To install `correlationid` from GitHub:

    go get -u github.com/kivra/krakend-correlationid@<commit hash>

Then add `correlationid` to the KrakenD [`router_engine`](https://github.com/devopsfaith/krakend-ce/blob/master/router_engine.go):

```go
func NewEngine(cfg config.ServiceConfig, opt luragin.EngineOptions) *gin.Engine {
	engine := luragin.NewEngine(cfg, opt)
	engine.Use(correlationid.HandlerFunc(cfg.ExtraConfig))
  ...
```

You can also use the `correlationid` middleware with any other `gin` application:

```go
engine := gin.New()
correlationid.Header = "My-Correlation-Id-Header" // optional
engine.Use(correlationid.Middleware)
```

## Usage

The `correlationid` middleware can be enabled using the following global
`extra_config`:

```json
"kivra/correlation-id": {
  "enabled": true
}
```

By default, the correlation ID is added to the `X-Correlation-Id` header. To
use a different header instead, use:

```json
"kivra/correlation-id": {
  "enabled": true,
  "header": "My-Correlation-Id-Header"
}
```
