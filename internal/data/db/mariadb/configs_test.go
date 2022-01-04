package mariadb

import (
	"fmt"
	"testing"

	"github.com/hi20160616/hfcms-articles/configs"
)

func TestLoad(t *testing.T) {
	cfg := configs.NewConfig("hfcms-articles")
	// cfg := &configs.Config{ProjectName: "hfcms-articles"}
	fmt.Println(cfg)
}
