package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	// 读配置
	ptrConfig := NewConfig()
	if err := ptrConfig.Parse(); err != nil {
		log.Fatalf("读取配置失败:%v", err)
	}

	// 日志文件
	if !ptrConfig.Verbose {
		log.SetOutput(ioutil.Discard)
	}

	// 配置参数输出
	log.Printf("%+v", ptrConfig)

	// 执行svn命令
	var cmd *exec.Cmd
	if ptrConfig.IncludeMerge {
		cmd = exec.Command("git", "log", "--pretty=format:'%s,,,,,%an,,,,,%ce,,,,,%ai,,,,,%H'", "-"+ptrConfig.LogNum)
	} else {
		cmd = exec.Command("git", "log", "--pretty=format:'%s,,,,,%an,,,,,%ce,,,,,%ai,,,,,%H'", "--no-merges", "-"+ptrConfig.LogNum)
	}
	log.Println("cmd:", cmd.String())
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("执行命令失败:%v", err)
	}
	svnData := string(out)
	//svnData := GBKToUTF8(string(out))
	//fmt.Println(svnData)

	// 命令格式化
	sliSvnLog, err := Format(svnData)
	if err != nil {
		log.Fatalf("解析日志失败:%v", err)
	}

	// 输出文件
	//if err = OutputCSV(ptrConfig.Output, sliSvnLog); err != nil {
	//	log.Fatalf("输出日志失败:%v", err)
	//}
	if err = OutputExcel(ptrConfig.Output, sliSvnLog); err != nil {
		log.Fatalf("输出日志失败:%v", err)
	}

	log.Println("执行命令成功,输出文件:", ptrConfig.Output)
}
