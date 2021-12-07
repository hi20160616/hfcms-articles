//go:build windows

package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows/registry"
)

type ProjectName string

type Config struct {
	ProjectName ProjectName
	RootPath    string
	Raw         []byte
	Debug       bool
	Verbose     bool // if true, prompt enter to exit.
	LogName     string
	Err         error
}

func NewConfig(projectName ProjectName) *Config {
	return setRootPath(&Config{ProjectName: projectName}).load()
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
		cfg.RootPath = filepath.Join("..\\", ".\\")
	}
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
	return c
}

func setRootPath(cfg *Config) *Config {
	cfg.RootPath, cfg.Err = readKey()
	if cfg.Err != nil {
		return cfg
	}
	cfg.RootPath = filepath.Dir(cfg.RootPath)
	if strings.Contains(os.Args[0], ".test") {
		return rootPath4Test(cfg)
	}
	return cfg
}

func readKey() (string, error) {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `Directory\shell\SpamKiller\command`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue("")
	if err != nil {
		return "", err
	}
	s = strings.Split(s, " ")[0]
	return strings.ReplaceAll(s, "\"", ""), nil
}
