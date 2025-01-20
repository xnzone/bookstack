---
author: xnzone 
title: 内存分配
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 723
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：malloc

**目标：实现一个内存分配器**

我们将在`lib/mem.c`里面加一个内核内存分配器。它是由指向空闲内存，并且不断增长的指针实现的

`kmalloc()`函数能够用来请求一个对齐的页，也返回真实的物理地址以便后续使用

我们将修改`kernel.c`并保留所有的shell代码，让我们尝试新的`kmalloc()`，并且检查我们第一页起始位置0x10000（像`mem.c`的硬件编码一样），然后`kmalloc()`产生一个新的地址，该地址从上一个地址对齐4096字节或0x1000

注意我们添加了一个新的`string.c:hex_to_ascii()`，用于打印十六进制数字

另一个修改是把`types.c`重命名为`type.c`，以保持语言一致性

其他文件与上一节相比，都没有改变