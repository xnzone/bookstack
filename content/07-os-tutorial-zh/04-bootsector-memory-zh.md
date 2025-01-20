---
author: xnzone 
title: 引导扇区内存
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 704
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：memory offsets, pointers

**目标：学习操作系统内存组成**

请打开[this document](http://www.cs.bham.ac.uk/~exr/lectures/opsys/10_11/lectures/os-dev.pdf)第14页, 并且查阅内存结构图

![内存结构图](https://gitcode.net/xnzone/solar/-/raw/master/2021/03/05/20210305173850.png)

这节课的唯一目的是学习引导扇区保存在哪里

我很直白的告诉你，BIOS放在`0x7C00`,但是一个错误的案例将更能看清这个事实

我们想在屏幕上打印一个`X`,我们将采取4种不同的策略，然后看看哪些可以工作，为什么能工作

打开文件`boot_sect_memory.asm`

首先我们将`X`定义成一个数据，并且使用一个label

```armasm
the_secret:
    db "X"
```

然后我们将尝试4种不同的方式去读取`the_secret`

1. `mov al, the_secret`
2. `mov al, [the_secret]`
3. `mov al, the_secret + 0x7C00`
4. `mov al, 2d + 0x7C00` `2d`是`X`字节在二进制中实际所在的位置

看代码并且阅读注释

编译并运行代码，你应该看到一个类似`1[2¢3X4X`的字符串，是1和2产生的随机垃圾二进制内容

如果你加上或者移除介绍，记住计算新的`X`的新的字节偏移值，并且代替`0x2d`

除非你已经完全理解了引导扇区偏移和内存地址，否则不要继续到下一节

## 全局偏移

现在，在每个地方都偏移`0x7C00`属实不方便,汇编可以对每个内存位置定义一个全局偏移(global offset)，其命令为`org`

```armasm
[org 0x7C00]
```

1. 现在去打开`boot_sect_memory_org.asm`，你将看到用引导扇区打印数据的典型方法
2. 编译并运行代码，你将看到`org`是如何印象每个之前的解决方案

阅读注释，以获取有关使用和不使用`org`的完整解释