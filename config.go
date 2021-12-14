package main

import (
	"flag"
)

type Config struct {
	LogNum       string
	Output       string
	ExcludeMerge bool
	Verbose      bool
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Parse() error {
	flag.StringVar(&this.LogNum, "l", "50", "日志条数")
	flag.StringVar(&this.Output, "o", "comment.xlsx", "输出文件")
	flag.BoolVar(&this.ExcludeMerge, "m", false, "包含merge日志")
	flag.BoolVar(&this.Verbose, "v", false, "显示详细输出")
	flag.Parse()

	return nil
}

func (this *Config) Merge() string {
	if this.ExcludeMerge {
		return ""
	} else {
		return "--no-merges"
	}
}
