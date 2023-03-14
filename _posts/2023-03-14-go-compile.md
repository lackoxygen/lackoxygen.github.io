---
layout: post
title: go编译
subtitle: go在不同开发环境下打包可执行文件
tags: [go]
comments: true
---

**Mac下编译Linux, Windows平台的64位可执行程序**

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
```

**Linux下编译Mac, Windows平台的64位可执行程序**

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
```

**Windows下编译Mac, Linux平台的64位可执行程序**

```
SET CGO_ENABLED=0SET GOOS=darwin3 SET GOARCH=amd64 go build test.go

SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build test.go
```

**参数说明**

- GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows

- GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm