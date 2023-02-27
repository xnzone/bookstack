---
author: xnzone 
title: 分段
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 6
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：segmentation

**目标：学会如何在16位实模式下分段**

如果你已经了解分段，你可以跳过这部分内容

我们将使用第3节的`[org]`进行分段。分段意味着你可以给你的引用数据指定一个偏移

完成这些需要使用特殊寄存器：`cs`,`ds`,`ss`和`es`来编码，保存数据，堆栈等待

注意：它们被CPU隐式使用，所以一旦你为某个寄存器（比如`ds`）设置了值，你的所有内存地址都会偏移`ds`,你可以阅读更多[相关内容](http://wiki.osdev.org/Segmentation)

更进一步，为了计算真实地址，我们不是仅仅加入两个地址，而是覆盖他们：`segment << 4 + address`.例如，如果`ds`是`0x4d`,[0x20]`的真实地址指向`0x4d0 + 0x20 = 0x4f0`

理论知识足够了。看一下代码吧

提示：我们不要在这些寄存器`mov`文字，我们必须使用一个之前的普通专门寄存器