---
author: xnzone 
title: 磁盘
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 7
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：hard disk, cylinder, head, sector, carry bit

**目标：让引导扇区从磁盘加载数据，以便于启动内核**

我们操作系统不能满足内置引导扇区只有512字节的标准，所以我们需要从磁盘读取代码，以便于能够运行内核

万分感谢的是，我们不需要处理打开和关闭磁盘旋转，我们仅仅需要调用一些BIOS的协程，就像我们打印字母到屏幕上一样。为了达到这点，我们把`al`设置成`0x02`(其他寄存器需要磁盘，磁头和扇区)并且使用中断`int 0x13`

具体中断内容可以参考[a detailed int 13h guide here](http://stanislavs.org/helppc/int_13-2.html)

这节课程，我们将第一次使用进位，当前寄存器溢出时，需要额外的一位保存进位

{{< highlight asm >}}
mov ax, 0xFFFF
add ax, 1 ; ax = 0x0000 and carry = 1
{{< /highlight  >}}

金文不能直接获取，但是可以被其他操作，例如`jc`（如果进位设置了，跳转）用来作为控制结构

BIOS也会把从扇区读到的数据设置给`al`,所以总是把它和期望的数值进行比较

## Code(编码)

打开和实验`boot_sect_disk.asm`，进行一次完整从磁盘读取数据的过程 `boot_sect_main.asm`位磁盘读取和调用`disk_load`准备了参数。注意怎么把从不属于引导扇区的额外数据写入，因为外部有512位标记

引导扇区实际上是硬盘0上的0号柱头0磁头的1号扇区

主程序用一些样本数据填充，然后让引导扇区读取

**注意：如果运行错误且你的代码看上去没问题，确保qemu是从正确的驱动启动的，直接设置从`dl`启动**

BIOS在调用引导扇区前，把驱动数据写入到`dl`。然而，当我用qemu从hdd启动的时候发现了一些问题

有两个快速解决方案：

1. 尝试使用`-fda`参数。 `qemu -fda boot_sect_main.bin`将把`dl`设置成`0x00`,这似乎是有用的
2. 明确使用`-boot`参数，例如`qemu boot_sect_main.bin -boot c` 会自动把`dl`设置成`0x80` 并且让引导扇区读取数据