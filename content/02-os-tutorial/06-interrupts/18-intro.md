---
author: xnzone 
title: 中断介绍
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 1
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：C types and structs, include guards, type attributes: packed, extern, volatile, exceptions

**目标：设置中断描述表，用于处理CPU中断**

这节课的灵感很大程度来自于[JamesM's tutorial](https://web.archive.org/web/20160412174753/http://www.jamesmolloy.co.uk/tutorial_html/index.html)

## 数据类型(Data types)

首先，在`cpu/types.h`中定义一些特殊的数据类型，这些数据结构帮助我们从字符和整数中解耦原始字节。已经放在了`cpu/`目录下，从现在开始，这个文件夹将会放一些机器依赖的代码。是的，引导代码，特别是x86，仍然在`boot/`目录下，但是我们现在依然让它在那里

一些已经存在的文件已经被修改成使用新的`u8`, `u16`和`u32`数据类型了

从现在开始，我们C头文件还将包含保护措施

## 中断(Interrupts)

中断是内核需要处理的一个主要事情。在将来的课程中，我们尽快实现能够接受键盘输入

中断的另外的例子有：除数为0，溢出，非法操作符，页错误等

中断由一个向量表处理，这个入口和GDT(第9节)类似。然而，汇编语言中没有IDT，所以我们需要用C语言实现。

`cpu/idt.h`定义了一个idt入口是怎么保存在`idt_gate`的（有256个中断需要处理，null或者CPU panic)。实际上，BIOS会加载idt结构，`idt_register`类似于GDT寄存器，仅仅有一个内存地址和大小。

最终，我们定义了一些变量用于从汇编代码中访问这些数据结构。

`cpu/idt.c` 仅仅是填充了每个结构的处理，正如你看到的那样，设置结构数据和调用`ldt`汇编命令是一个问题。

## ISRs

每次CPU检测一个中断的时候，中断服务程序都会运行，这是非常严重的。

我们将写足够的代码去处理他们，打印一个错误信息，销毁CPU

在`cpu/isr.h`上，我们手动定义了32个。右`extern`声明，因为他们将会在`cpu/interrupt.asm`中，由汇编语言实现。

在跳到汇编代码之前，检查`cpu/isr.c`。正如你看到的一样，我们马上定义了一个函数去安装所有isrs，然后加载IDT，一系列错误信息和高级的处理打印信息。你可以用`isr_handler`打印你想要的任何东西

现在来到低一层，每个`idt_gate`把低层和高层联系到一起了。打开`cpu/interrupt.asm`。在这里，我们定义了一个常用的低级ISR代码，这个代码基本用于保存状态和调用C个代码的。然而实际ISR汇编函数是指向`cpu/isr.h`的

注意，`registers_t`结构是怎么表示我们放置在`interrupt.asm`的所有寄存器的

这是非常基本的。现在我们需要从我们的Makefile引用`cpu/interrupt.asm`, 让内核安装ISRs，然后加载其中一个。注意，怎么让CPU在一些中断之后不销毁，尽管这是一个很好的练习。