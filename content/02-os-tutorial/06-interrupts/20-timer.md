---
author: xnzone 
title: 定时器
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 3
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：CPU timer, keyboard interrupts, scancode

**目标：实现我们第一个IRQ处理器：CPU定时器和键盘**

现在所有的一切都准备测试我们的硬件中断

## 定时器(Timer)

定时器是非常容易配置的。首先我们在`cpu/timer.h`声明一个`init_timer()`并且在`cpu/timer.c`中实现它。可能处理时钟频率和发送字节给合适的端口是比较麻烦。

现在我们修复`kernel/util.c`的`int_to_ascii()`，用正确的顺序来打印数字。因此，我们需要实现`reverse()`和`strlen()`

最后，回到`kernel/kernel.c`，需要做两件事。再次开启中断(非常重用！)，然后初始化定时器中断。

运行`make run`你可以看到时钟

## 键盘(Keyboard)

键盘就更简单了，但是有一个缺点。PIC发送给我们的不是按下键的ASCII码，而是对于按下或弹起事件的扫描码。所以我们需要转换他们

查阅`drivers/keyboard.c`，这里有两个函数：回调和在中断回调中配置的初始化函数。一个新的`keyboard.h`用来定义这些函数

`keyboard.c`也有一个长表，用来扫描码到ASCII码键的转换。开始，我们只需要实现一个简单的US键盘转换。你可以查阅更多关于[扫描码](http://www.win.tue.nl/~aeb/linux/kbd/scancodes-1.html)信息

我不知道你现在感受如何，我反正是非常激动了。我们非常接近构建一个简单的shell了，在下一个章节，我们将在键盘输入上稍微扩展一下