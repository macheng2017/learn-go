package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("hello", "world", "print hello world")
	flag.Parse()
	fmt.Printf("s %v", *s)
}

//猜测用法
//$ go run flagString/main.go flag- hello
//s world
//正确用法
//1.先编译    go build main.go
//2.运行编译后代码 ./main.exe -hello haha
//3.结果是命令后的value值,usage则是提示内容 s haha

// 通过部署方式反推这部分代码的作用
// /data/supergroup.mixin.one/supergroup.mixin.one -service http -dir /data/supergroup.mixin.one
// /data/supergroup.mixin.one/supergroup.mixin.one -service message -dir /data/supergroup.mixin.one
// 启动message的时候通过传入不同的参数启动了两个服务一个是api service 一个是 message service

//[Mixin 大群部署完全教程 – Exin 团队博客](https://blog.exin.one/2019/05/25/mixin-super-group/)
//4.13 部署服务
//为了管理方便，为两个 Go 进程部署 Service 服务。因为重构后的代码，编译不依赖配置文件，所以 Service 服务脚本，需要指定 -dir 参数。
//
//$ cd /etc/systemd/system
//$ touch test.api.service
//$ touch test.message.service
//Bash
//test.api.service 文件内容如下：
//
//[Unit]
//Description=Group API Daemon
//After=network.target
//
//[Service]
//# 需要把 test 替换为运行 service 的 linux 用户，比如 ubuntu
//User=test
//Type=simple
//# 新版无需依赖 $GOPATH
//ExecStart=/data/supergroup.mixin.one/supergroup.mixin.one -service http -dir /data/supergroup.mixin.one
//Restart=on-failure
//LimitNOFILE=65536
//
//[Install]
//WantedBy=multi-user.target
//Bash
//test.message.service 文件内容如下：
//
//[Unit]
//Description=Group Message Daemon
//After=network.target
//
//[Service]
//# 需要把 test 替换为运行 service 的 linux 用户，比如 ubuntu
//User=test
//Type=simple
//# 新版无需依赖 $GOPATH
//ExecStart=/data/supergroup.mixin.one/supergroup.mixin.one -service message -dir /data/supergroup.mixin.one
//Restart=on-failure
//LimitNOFILE=65536
//
//[Install]
//WantedBy=multi-user.target
//Bash
//接下来重启服务：
//
//$ systemctl restart test.api.service
//$ systemctl restart test.message.service
