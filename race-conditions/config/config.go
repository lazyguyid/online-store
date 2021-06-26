package config

import (
	"online-store/core"
	"os"
	"strconv"
)

type config struct {
	vars map[string]interface{}
}

// Load func
func Load() core.Config {
	cfg := new(config)
	cfg.vars = make(map[string]interface{})

	cfg.vars["PORT"] = os.Getenv("PORT")
	cfg.vars["DEBUG"], _ = strconv.ParseBool(os.Getenv("DEBUG"))
	cfg.vars["TIMEZONE"] = os.Getenv("TIMEZONE")
	cfg.vars["DATABASE_URL"] = os.Getenv("DATABASE_URL")

	return cfg
}

func (cfg *config) Get(k string) interface{} {
	return cfg.vars[k]
}

func (cfg *config) Set(k string, v interface{}) {
	cfg.vars[k] = v
}
