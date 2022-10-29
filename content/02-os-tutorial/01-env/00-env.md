---
author: xnzone 
title: 开发环境
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 2
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念:linux,mac,terminal,compiler,emulator,nasm,qemu

**目标:安装用来运行教程所需要的软件**

我使用Mac工作，尽管Linux更适合，因为Linux已经为你安装了所有的标准软件

在一台Mac上，[安装Homebrew](http://brew.sh/)，然后`brew install qemu nasm`

不要使用Xcode开发者工具`nasm`，如果你已经安装了，在某些情况下，是不能工作的。总是使用`/usr/local/bin/nasm`

在一些操作系统上，qemu被分割了多个二进制文件，你可能需要调用`qemu-system-x86_64 binfile`