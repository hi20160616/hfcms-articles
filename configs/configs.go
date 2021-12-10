package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type ProjectName string

type Config struct {
	ProjectName ProjectName
	RootPath    string
	Raw         []byte
	Debug       bool
	Verbose     bool // if true, prompt enter to exit.
	LogName     string
	API         struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	Database struct {
		Driver string `json:"Driver"`
		Source string `json:"Source"`
	} `json:"database"`
	Web struct {
		Addr string `json:"Addr"`
		Tmpl string `json:"Tmpl"`
	} `json:"web"`
	Err error
}

type api struct {
	Addr, Timeout string
}

func NewConfig(projectName ProjectName) *Config {
	return setRootPath(&Config{ProjectName: projectName}).load()
}

func setRootPath(cfg *Config) *Config {
	cfg.RootPath, cfg.Err = os.Getwd()
	if cfg.Err != nil {
		return cfg
	}
	if strings.Contains(os.Args[0], ".test") {
		return rootPath4Test(cfg)
	}
	return cfg
}

func rootPath4Test(cfg *Config) *Config {
	cfg.RootPath, cfg.Err = os.Getwd()
	if cfg.Err != nil {
		return cfg
	}
	ps := strings.Split(cfg.RootPath, string(cfg.ProjectName))
	n := 0
	if len(ps) > 1 {
		n = strings.Count(ps[1], string(os.PathSeparator))
	}
	for i := 0; i < n; i++ {
		cfg.RootPath = filepath.Join("../", "./")
	}
	cfg.RootPath = filepath.FromSlash(cfg.RootPath)
	return cfg
}

func (c *Config) load() *Config {
	if c.Err != nil {
		return c
	}
	cfgFile := filepath.Join(c.RootPath, "configs", "configs.json")
	c.Raw, c.Err = os.ReadFile(cfgFile)
	if c.Err != nil {
		if errors.Is(c.Err, os.ErrNotExist) {
			c.Err = errors.WithMessage(c.Err, "ReadFile error: no configs.json")
		}
		return c
	}
	cfgTemp := &Config{}
	if c.Err = json.Unmarshal(c.Raw, cfgTemp); c.Err != nil {
		c.Err = errors.WithMessage(c.Err, "Unmarshal configs.json error")
		return c
	}
	c.Debug = cfgTemp.Debug
	c.Verbose = cfgTemp.Verbose
	c.LogName = cfgTemp.LogName
	c.ProjectName = cfgTemp.ProjectName
	c.API = cfgTemp.API
	c.Database = cfgTemp.Database
	c.Web = cfgTemp.Web

	return c
}
