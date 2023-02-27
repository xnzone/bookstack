---
author: xnzone 
title: 滚动显示
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 3
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：scroll

**目标：当文本到达底部的时候，滚动屏幕**

对于这个小节，打开`drivers/screen.c` 注意`print_char`的尾部，有新的代码片段(84行)，这段代码是用来检查如果当前偏移超过了屏幕尺寸，就滚动文本。

实际滚动操作是由一个新的函数`memory_copy`处理的。它是`memcpy`的一个精简版本，但是为了避免命名冲突，我们取个不一样的名字，至少对现在来说是这样的。打开`kernel/util.c`可以看一下具体的实现

为了帮助可视化滚动，我们也要实现一个函数把整数转换为文本的函数`int_to_ascii`。同样的，这也是标准`itoa`的一个精简版的实现。注意整数有多位数，它们输出是相反的。这是有意的。在将来的课程中，我们将扩展我们的辅助函数，但是不是现在的重点。

最终打开`kernel/kernel.c`。每一行显示它的行号。你可以在14行设置断点，去证实它。然后接下来的`kprint`强制内核向下滚动

这节课也是os-dev.pdf的结尾。从现在开始，我们将参考[OSDev wiki](http://wiki.osdev.org/Meaty_Skeleton)和其他源码。非常感谢Prof. Blundell提供这么好的文档