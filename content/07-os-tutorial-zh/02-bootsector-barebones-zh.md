---
author: xnzone 
title: 引导扇区基础架构
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 702
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念:汇编，BIOS

**目标：创建一个文件，BIOS将该文件解释为可引导磁盘**

这是非常激动的，我们将创建我们自己的引导扇区

## 理论知识

当计算机启动时，BIOS并不知道怎么加载操作系统，所以需要引导扇区来完成这个任务。因此，引导扇区必须放在一个已知的，标准的位置。这个位置时磁盘的第一个扇区(0柱面，0磁头，0扇区）并且占用512个字节

为了保证磁盘是可引导的，BIOS需要检查511字节和512字节是否为`0xAA55`

这是一个最简单的引导扇区

```text
e9 fd ff 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
[ 29 more lines with sixteen zero-bytes each ]
00 00 00 00 00 00 00 00 00 00 00 00 00 00 55 aa
```

基本上都是0，最后16位以`0xAA55`(需要注意大小端，x86是小端编码)，前3个字节表示无限跳转

## 有史以来最简单的引导扇区

你既可以把上面512个字节写入到二进制编辑器，也可以使用下面的汇编语言

```asm
;无限循环(e9 fd ff)
loop:
    jmp loop

;填充510个0
times 510-($-$$) db 0
;魔术数
dw 0xaa55
```

编译：`nasm -f bin boot_sec_sample.asm -o boot_sec_sample.bin`

> 警告:如果有error，阅读[环境配置](/02-os-tutorial/01-env/00-env)

我知道你非常期待想要尝试运行它，我也一样，所以淦 `qemu boot_sec_sample.bin`

> 在某些操作系统，你必须运行`qemu-system-x86_64 boot-sec_sample.bin`,如果报SDL错误，尝试使用`--nographic`或者`--curses`变量

你将看见一个窗口，上面除了 "Booting from Hard Disk..."， 什么也没有。你上一次看到一个无限循环是如此激动的时刻是什么时候？;-)