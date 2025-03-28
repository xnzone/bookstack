---
author: xnzone 
title: 32位全局描述表
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 710
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：GDT

**目标：编码GDT**

还记得第6节的segmentation吗？偏移是朝地址的反方向

在32位模式，分段工作方式不一样。现在，在GDT中，偏移量编程了段描述符(SD)。描述符定义了32位基地址，长度为20位和其他标志位，比如只读，权限等。总而言之，数据结构被划分了，所以打开os-dev.pdf文件，查阅34页的图片或者维基百科关于GDT的页面

早期编写GDT的方式是定义两个段，一个用于代码，另一个用于数据。这样可能会被覆盖，也就意味着没有保护，但是却足够引导，我们将在高级语言中修复这个问题

出于好奇，为了确保程序员在管理地址的时候不会出错，GDT的入口必须是`0x00`

其次，CPU不能直接加载GDT地址，但是需要一个名为"GDT描述符"的元结构，其长度为16b，地址为32b的GDT。通过`lgdt`操作被加载

让我们直接跳到GDT的汇编代码。为了再次理解所有段标志位，请查阅os-dev.pdf文档。这节可得理论是十分复杂的

在下一节课程，我们切换到32位保护模式，并从这些课程中测试我们的代码