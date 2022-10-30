---
author: xnzone 
title: 内核系统选择
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 3
tags: ["tutorial", "os"]
---


你需要通过谷歌预先了解的概念：kernel, ELF format, makefile

**目标：创建一个简单内核和一个可以引导的引导扇区**

## 内核

我们的C内核将在屏幕左上角打印一个'X'。直接打开`kernel.c`

我们将注意到一个没有做任何事情的辅助函数。那个函数将强制我们创建一个没有指向`0x0`字节的内核入口点，但是我们知道一个实际标签加载它。在我们这个例子中，就是`main()`函数

{{< highlight bash >}}
i386-elf-gcc -ffreestanding -c kernel.c -o kernel.o
{{< /highlight  >}}

在`kernel_entry.asm`里面有日常工作。阅读它，你将学会在汇编里怎么使用`[extern]`声明。为了生成一个会被`kernel.o`链接的`elf`格式文件，我们需要编译这个文件，而不是生成一个二进制文件。

{{< highlight bash >}}
nasm kernel_entry.asm -f elf -o kernel_entry.o
{{< /highlight  >}}

## 链接

链接器是非常有用的工具，我们现在只需要从中受益即可

为了链接两个对象文件到一个二进制内核，解决标签引用，运行：

{{< highlight bash >}}
i386-elf-ld -o kernel.bin -Ttext 0x1000 kernel_entry.o kernel.o --oformat binary
{{< /highlight  >}}

注意不是把内核放到`0x0`位置，而是`0x1000`位置。引导扇区同样也需要知道地址。

## 引导扇区

这和第10节的内容非常相似。打开`bootsect.asm`并且测试代码。实际上，如果你移除所有用来在屏幕上打印信息的行，它会统计有多少行。

编译它用`nasm bootsect.asm -f bin -o bootsect.bin`

## 所有内容放到一起

什么？对于引导扇区和内核，我们有两个分开的文件？

我们不能把他们连接到一个文件？是的，我们可以，且是非常简单的，仅仅级联他们就可以了：

{{< highlight bash >}}
cat bootsect.bin kernel.bin > os-image.bin
{{< /highlight  >}}

## 运行！

你可以用qemu运行`os-image.bin`

记住，如果你发现硬盘加载错误，你可能需要使用硬盘编号，或者qemu参数(floppy = `0x0`, hdd = `0x80`)。我通常使用 `qemu-system-i386 -fda os-image.bin`

我们可以看到四条信息

- "Started in 16-bit Real Mode"
- "Loading kernel into memory"
- (Top left) "Landed in 32-bit Protected Mode"
- (Top left, overwriting previous message) "X"

恭喜

## Makefile

最后异步，我们用一个Makefile整理编译过程。打开`Makefile`脚本，测试它的内容。如果你不知道Makefile是什么，现在是一个很好的机会用谷歌学习它，这将在今后的日子里为我们节省大量的时间