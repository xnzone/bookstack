---
author: xnzone 
title: 进入32位模式
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 711
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：interrupts, pipelining

**目标：进入32位保护模式，并且测试之前课程的代码**

为了进入32位模式：

1. 关闭中断
2. 加载GDT
3. 把CPU的控制寄存器`cr0`置位
4. 通过精心设计的跳转刷新CPU管道
5. 更新所有段寄存器
6. 更新栈
7. 调用一个周知的label，该label包含第一个有用的32位代码

我们将把流程浓缩到`32bit-switch.asm`文件中，打开它并且查阅代码

在进入32位模式后，我们将调用`BEGIN_PM`,该命令是我们实际有用代码(例如内核代码等)的入口。你可以在`32bit-main.asm`中阅读到这段代码。编译并且运行最后一个文件，你会在屏幕上看到两条消息

恭喜！我们下一步将写一个简单的内核