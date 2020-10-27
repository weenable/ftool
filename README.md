# ftool

一款基于命令行的文件管理小工具，目前包含以下功能
- 复制文件
- 移动文件
- 重命名文件
- 压缩文件

```
Usage:
  ftool [command]

Available Commands:
  compress    Compress files in folder
  copy        Copy files from source folder to target folder
  help        Help about any command
  move        Move files from source folder to target folder
  rename      Rename files in batch, auto-increment number will be added as suffix

Flags:
  -e, --ext string      Return files that have the specified extension
  -h, --help            help for ftool
  -l, --log             Print logs
  -p, --prefix string   Return files that have the specified prefix in the file name
  -r, --recursive       Scan files in sub-directories recursively
  -x, --regex string    Return files that match the regex pattern in the file name
  -s, --suffix string   Return files that have the specified suffix in the file name
  -v, --version         version for ftool

Use "ftool [command] --help" for more information about a command.
```