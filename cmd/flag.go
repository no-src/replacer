package cmd

import (
	"flag"

	"github.com/no-src/log/level"
)

var (
	ConfigFile   string
	ConfigUrl    string
	Root         string
	Tag          string
	LogDir       string
	LogLevel     int
	Revert       bool
	PrintVersion bool
	PrintAbout   bool
)

// InitFlags init built-in flags
func InitFlags() {
	flag.StringVar(&ConfigFile, "conf", "./replacer.yaml", "the config file of replacer")
	flag.StringVar(&ConfigUrl, "conf_url", "", "the remote url of replacer config file")
	flag.StringVar(&Root, "root", "./", "workspace root path")
	flag.StringVar(&Tag, "tag", "default", "tag name")
	flag.StringVar(&LogDir, "log_dir", "./logs", "log directory")
	flag.IntVar(&LogLevel, "log_level", int(level.DebugLevel), "log level")
	flag.BoolVar(&Revert, "revert", false, "revert the replace operations")
	flag.BoolVar(&PrintVersion, "v", false, "print the version info")
	flag.BoolVar(&PrintAbout, "about", false, "print the about info")
	flag.Parse()
}
