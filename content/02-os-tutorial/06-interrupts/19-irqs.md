---
author: xnzone 
title: IRQ寄存器
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 2
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：IRQs, PIC, polling

**目标：完成中断实现和CPU定时器**

当CPU启动的时候，PIC把IRQs 0-7 映射到 INT 0x8-0xF，把 IRQs 8-15 映射到 INT 0x70-0x77。这和我们上一节编程的ISRs有冲突。因为，我们编程的ISR 0-31，把IRQs重映射到ISRs 32-47是一个标准做法。

PICs通过I/O端口(参考15课)交互。主PIC命令在0x20，数据在0x21，然而从PIC命令在0xA0，数据在0xA1。

重映射PICs的代码是非常奇怪的，其中包含一些掩码，如果你好奇的话，请参考[这篇文章](http://www.osdev.org/wiki/PIC)。否则，仅仅看`cpu/isr.c`，在设置为IRS的IDT门后，有一段新代码。毕竟，我们是为IRQs添加IDT门

现在跳转到汇编`interrupt.asm`。首先任务是为我们在C代码里使用的IRQ符号添加全局定义。查看`global`描述的最后就可以了。

然后，添加IRQ处理器。和`interrupt.asm`一样，在最后添加的。注意他们是怎么跳到一个新的常规根的：`irq_common_stub`（下一步介绍)

和ISR一样，创建一个`irq_common_stub`。它位于`interrupt.asm`文件的顶部，也定义了一个新的`[extern irq_handler]`

现在回到C代码，在`isr.c`文件里写`irq_handler()`的代码。它把EOIs发送给PICs，然后调用合适的处理器，处理器是保存在一个名为`interrupt_handlers`数组里的，这个数据也被定义在文件的顶部。一个新的结构体定义在`isr.h`里。我们也使用一个简单的函数去注册中断处理器。

这有很多工作，但是我们现在可以定义我们第一个IRQ处理器了

`kernel.c`没有任何变化，所以也没有新的东西需要运行和讲解的。开始移步下一节，检查你新闪闪的IRQs吧。