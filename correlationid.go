package correlationid

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-contrib/uuid"
	"github.com/luraproject/lura/v2/config"
)

const Namespace = "kivra/correlation-id"

var Header = "X-Correlation-Id"

type Config struct {
	Enabled bool   `json:"enabled"`
	Header  string `json:"header"`
}

func ConfigGetter(extraCfg config.ExtraConfig) (*Config, bool) {
	cfg := new(Config)

	tmp, ok := extraCfg[Namespace]
	if !ok {
		return cfg, false
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(tmp); err != nil {
		return cfg, false
	}
	if err := json.NewDecoder(buf).Decode(cfg); err != nil {
		return cfg, false
	}

	if cfg.Header == "" {
		cfg.Header = Header
	} else {
		Header = cfg.Header
	}

	return cfg, true
}

func Middleware(c *gin.Context) {
	if c.Request.Header.Get(Header) == "" {
		c.Request.Header.Set(Header, strings.ToUpper(uuid.NewV4().String()))
	}
	c.Next()
}

func HandlerFunc(extraCfg config.ExtraConfig) gin.HandlerFunc {
	cfg, ok := ConfigGetter(extraCfg)
	if !ok || !cfg.Enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	return Middleware
}
