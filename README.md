# 说明

## 功能

将git日志导出成excel

## 使用方法

- 自己将程序放到bin目录或者复制到程序执行的目录
- 运行`gitlog.exe`
- 参数
  - `-h`,显示帮助
  - `-l`,导出日志数量,默认`10`
  - `-m`,导出包含merge,默认`false`
  - `-o`,输出文件,默认`comment.xlsx`
  - `-v`,显示详细输出,,默认`false`

## 例子

```bash
# 生成文件comment.xlsx,包含最近10条log日志
gitlog.exe

# 最近20条日志
gitlog.exe -l 20

# 输出到E:/abc.xlsx
gitlog.exe -l 20 -m -o /e/abc.xlsx

# 显示详细输出
gitlog.exe -l 20 -m -o /e/abc.xlsx -v
```
