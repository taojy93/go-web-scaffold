// main.go
package main

import (
	"go-web-scaffold/cmd"
	"go-web-scaffold/internal/logging"
)

func main() {

	cmd.Execute()

	// 作用是确保日志缓冲区中的日志条目在程序退出之前被刷新到日志输出目标（例如，文件、控制台等）。这是因为 zap 采用了高效的异步日志写入方式，日志可能会暂时存储在内存中的缓冲区里，如果不进行同步，程序退出时可能会丢失一些尚未写入的日志条目。
	logging.SyncLogger()
}
