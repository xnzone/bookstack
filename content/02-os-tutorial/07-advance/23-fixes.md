---
author: xnzone 
title: 性能优化
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 3
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：freestanding, uint32_t, size_t

**目标：修复代码中各种各样的问题**

OSDev wiki有一节描述了[JameM's教程存在的一些问题](http://wiki.osdev.org/James_Molloy's_Tutorial_Known_Bugs)。因为我们在18-22节课(interrupts through malloc)是参考他的教程的，我们在继续之前，必须修复这些问题

## 1.错误标志(Wrong CFLAGS)

当编译`.o`文件的时候，我们添加了`-ffreestanding`，这个操作包括`kernel_entry.o` 从而影响了`kernel.bin`和`os-image.bin`

之前，我们通过使用`-nostdlib`禁用了libgcc(不是libc)，我们在链接的时候，没有重新启用它。由于这个很棘手，我们将删除`-nostdlib`

`-nostdinc`也会传递给gcc，但是我们在第三步需要，所以现在就先删掉它

## 2.内核main函数(kernel.c `main()` function)

修改`kernel/kernel.c`，把`main()`改成`kernel_main()`。因为gcc把"main"当作一个特殊的关键字，我们不想弄混

修改`boot/kernel_entry.asm`直接指向新的名称

为了修复`i386-elf-ld: warning: cannot find entry symbol _start; defaulting to 0000000000001000` 告警信息，添加一个`global_start;`, 然后在`boot/kernel_entry.asm`定义一个label`_start:`

## 3.重构数据类型(Reinvented datatypes)

定义非标准数据类型比如`u32`是一个糟糕的主意，从C99开始，引入了一个标准修复长度数据类型如`uint32_t`

我们需要包含`<stdint.h>`，它甚至可以在`-ffreestanding`(但是需要stdlibs)工作，使用这些数据类型，而不是自定义的，删掉`type.h`里面的自定义的数据类型

删除`__asm__`和`__volatile__`旁边的下划线，因为他们不需要了

## 4.不合适对齐`kmalloc`(Improperly aligned `kmalloc`)

首先，因为`kmalloc`使用一个长度参数，我们将使用正确的数据类型`size_t`而不是`u32int_t`。`size_t`应该被用来计算物品个数的所有参数，并且是非负数。包含在`<stddef.h>`

我们后面会修复`kmalloc`，让它作为一个合适的内存管理单元和对齐的数据类型。但是现在，它总是返回一个新的页对齐内存块

## 5.缺失的函数(Missing functions)

我们在接下来的课程中，实现缺失的`mem*`函数

## 6.中断处理(Interrupt handlers)

`cli`是多余的，因为我们已经建立了IDT入口，在处理程序中使用`idt_gate_t`标志，就可以启用中断

`sti`也是多余的，当`iret`从栈中加载标志位的时候，里面有个标志位可以知道是否启用了中断。换句话说，在中断之前，中断处理程序自动保存中断是否启用

在`cpu/isr.h`中，`struct registers_t`有一些问题。首先，`esp`被重命名为了`useless`。这个值是有用的，因为它必须处理当前栈的上下文信息，而不是被中断。然而把`useresp`重命名为`esp`

osdev wiki建议在`cpu/interrupt.asm`的`call_isr_handler`之前添加`cld`

最后， `cpu/interrupt.asm`中一些重要的问题。通用根在堆栈上创建结构寄存器实例，然后调用C处理程序。但是这破坏了ABI，因为堆栈属于被调用的函数，并且可以根据需要修改。需要将其作为指针进行结构传递

为了实现这个，编辑`cpu/isr.h`和`cpu/isr.c`， 把`registers_t r`修改成`registers_t *t`， 然后使用`->`访问成员而不是`.`。最后，在`cpu/interrupt.asm`中，在调用`isr_handler`和``irq_handler`之前都添加一个`push esp`。 记住后面也要`pop eax`用于清理指针

所有当前回调，定时器和键盘，也要使用指针修改`registers_t`