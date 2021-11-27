package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type GitLog struct {
	version  string
	author   string
	email    string
	datetime string
	comment  string
	phx      string
}

func NewGitLog() *GitLog {
	return &GitLog{}
}

func (this *GitLog) Load(data string, ptrCommitFormat *regexp.Regexp) error {
	log.Println("[加载]", data)
	data = strings.ReplaceAll(data, "\r\n", " ")
	data = strings.ReplaceAll(data, "\n", " ")
	sliData := strings.Split(data, ",,,,,")
	l := len(sliData)
	i := 0
	if l > i {
		this.comment = sliData[i]
		if ptrCommitFormat != nil {
			sliStr := ptrCommitFormat.FindStringSubmatch(this.comment)
			if len(sliStr) > 2 {
				this.phx = sliStr[1]
				this.comment = sliStr[2]
			}
		}
	}
	i++
	if l > i {
		this.author = sliData[i]
	}
	i++
	if l > i {
		this.email = sliData[i]
	}
	i++
	if l > i {
		this.datetime = sliData[i]
	}
	i++
	if l > i {
		this.version = sliData[i]
	}
	i++

	return nil
}

func (this *GitLog) ToCsv() string {
	return fmt.Sprintf("%v,%v,%v,%v\r\n", this.comment, this.author, this.datetime, this.version)
}

func (this *GitLog) Test() {
	log.Printf("%+v", *this)
}

func Format(gitData string) ([]*GitLog, error) {
	log.Println("完整数据:", gitData)

	// 找到分隔符
	sep := "'\r\n'"
	if strings.Index(gitData, sep) < 0 {
		sep = "'\n'"
	}

	// 去掉首尾的`'`
	if len(gitData) > 2 {
		gitData = gitData[1 : len(gitData)-1]
	}

	// 数据解析
	ptrCommitFormat := regexp.MustCompile("(?i)^(PHX-[0-9]+)(.*)")
	sliGitData := strings.Split(gitData, sep)
	sliGitLog := make([]*GitLog, 0)
	for _, data := range sliGitData {
		data = strings.TrimSpace(data)
		data = strings.ReplaceAll(data, "\r\n\r\n", "\r\n")
		data = strings.ReplaceAll(data, "\n\n", "\n")
		ptrGitLog := NewGitLog()
		if err := ptrGitLog.Load(data, ptrCommitFormat); err != nil {
			return sliGitLog, err
		}
		sliGitLog = append(sliGitLog, ptrGitLog)
	}

	return sliGitLog, nil
}
