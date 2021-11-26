package main

import (
	"flag"
)

type Config struct {
	LogNum       string
	Output       string
	IncludeMerge bool
	Verbose      bool
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Parse() error {
	flag.StringVar(&this.LogNum, "l", "10", "日志条数")
	flag.StringVar(&this.Output, "o", "comment.xlsx", "输出文件")
	flag.BoolVar(&this.IncludeMerge, "m", false, "包含merge日志")
	flag.BoolVar(&this.Verbose, "v", false, "显示详细输出")
	flag.Parse()

	return nil
}

func (this *Config) Merge() string {
	if this.IncludeMerge {
		return ""
	} else {
		return "--no-merges"
	}
}
