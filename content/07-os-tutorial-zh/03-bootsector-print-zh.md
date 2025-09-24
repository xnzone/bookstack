---
author: xnzone 
title: 引导扇区输出
date: 2021-09-10 10:23:32
image: https://s2.loli.net/2025/09/24/J8Y6zBqZmSn9NkW.jpg
cover: false
weight: 703
tags: ["tutorial", "os"]
---


你需要通过谷歌预先了解的概念：interrupts, CPU registers

**目标：让我们上一个安静的引导扇区打印一些文本**

我们将在无限循环引导扇区做一点改进，让其在屏幕上输出一些文本。为此，我们将要使用中断

在这个例子中，我们将"Hello"单词逐个字母写入`al`寄存器(`ax`的低位部分),字节`0x0e`写入`ah`(`ax`高位部分),并且使用`0x10`中断，这是一个视频服务的普通中断

`ah`上的`0x0e`告知视频中断器， 我们实际想做的事情是"使用tty模式，在`al`中写入内容"

我们只需要设置一次tty模式，然而在现实情况中，我们不能保证`ah`的内容是常数。当我们进程在休眠的时候，一些进程可能在CPU上运行,没有正确清理并留下垃圾数据在`ah`上

在这个例子中，我们不需要关心这个，因为我们在CPU上只运行一件事

我们的引导扇区像这样:

```armasm
mov ah, 0x0e ; tty模式
mov al, 'H'
int 0x10
mov al, 'e'
int 0x10
mov al, 'l'
int 0x10
int 0x10 ; 'l' 依然在al,记住了吗? 
mov al, 'o'
int 0x10

jmp $ ; 跳转到当前的地址 = 无限循环

;填充和魔术数 
times 510 - ($-$$) db 0
dw 0xaa55
```

你可以用`xxd file.bin`检查二进制数据

不管怎样，你知道命令: `nasm -f bin boot_sect_print.asm -o boot_sect_print.bin` `qemu boot_sect_print.bin`

你的引导扇区将输出"Hello"，并且进入无限循环