package cmd

import (
	"github.com/no-src/log"
	"github.com/no-src/log/level"
	"github.com/no-src/nsgo/httputil"
	"github.com/no-src/replacer"
	"github.com/no-src/replacer/internal/logger"
)

// RunWithFlags start the replacer with built-in flags
func RunWithFlags() error {
	return Run(ConfigFile, ConfigUrl, Root, Tag, LogDir, LogLevel, Revert)
}

// Run start the replacer with custom arguments
func Run(conf, configUrl, root, tag, logDir string, logLevel int, revert bool) error {
	if err := logger.InitLogger("replacer", logDir, level.Level(logLevel)); err != nil {
		return err
	}
	if len(configUrl) > 0 {
		client, err := httputil.NewHttpClient(true, "", false)
		if err != nil {
			return err
		}
		if err = client.Download(conf, configUrl, false); err != nil {
			return err
		}
	}
	r, err := replacer.NewReplacer(conf)
	if err != nil {
		log.Error(err, "init replacer config error")
		return err
	}
	if revert {
		err = r.Revert(root, tag)
	} else {
		err = r.Replace(root, tag)
	}
	if err != nil {
		log.Error(err, "execute replace error")
	}
	return err
}
