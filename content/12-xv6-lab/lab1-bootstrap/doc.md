---
author: xnzone 
title: 实验手册 
date: 2021-09-10
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1211
tags: ["xv6", "os", "bootstrap"]
---

## PC Bootstrap
第一个练习的目的是介绍x86会变语言和PC bootstrap程序，使用QEMU和QEMU/GDB进行调试。这部分你不用写任何代码，但是不管怎样，你最好过一遍，并且回答文章后面的问题

### x86汇编

如果你对x86汇编语言不熟悉，通过这个课程，你将很快熟悉它。[PC Assembly Language Book](https://pdos.csail.mit.edu/6.828/2018/readings/pcasm-book.pdf)是一个非常好的入门书籍，这本书混杂了最新和旧版本的信息

*警告*：不幸的是，这本书是用NASM汇编语言来编写的，然而，我们将使用GNU汇编。NASM使用所谓的*Intel*语法，然而GNU使用的是*AT&T*语法。两者的差异非常大，所幸的是,使用这个工具[Brenman's Guide to Inline Assembly](http://www.delorie.com/djgpp/doc/brennan/brennan_att_inline_djgpp.html)，能够快速的转换

>进行Exercise 1

当然，x86汇编语言编程的参考是Intel的白皮书，你可以在[6.828的参考页](https://pdos.csail.mit.edu/6.828/2018/reference.html)找到两个版本：一个是老的[80386 Programemer's Reference Manual](https://pdos.csail.mit.edu/6.828/2018/readings/i386/toc.htm), 这个版本简短，也比较简单，但是描述了所有x86处理器的特征，6.828课程也是使用这个作为参考；另一个是最新版本的[IA-32 Intel Architecture Soft Developer's Manuals](https://software.intel.com/content/www/us/en/develop/articles/intel-sdm.html)，最新版包含了最新处理器的所有特征，但是这个课程用不上，如果你感兴趣，可以阅读。还有一个关于[AMD的手册](https://developer.amd.com/resources/developer-guides-manuals/)相对来说更加友好，但是仅针对AMD的处理器

### 仿真x86
不是在真正的物理机上开发一个操作系统，而是使用模拟器模拟一个完整的PC。适用于模拟器的代码当然也可以在实际物理机器上跑。使用模拟器可以简化调试。例如，你可以设置在模拟的x86中设置断点，但是这个在实际的x86系统中却很难做到

在6.828课程中，我们将使用[QEMU](http://www.qemu.org/),QEMU可以配合[GDB](http://www.gnu.org/software/gdb/)一起使用，进行调试。

在`lab`目录输入`make`,可以看到下面的输出
```text
+ as kern/entry.S
+ cc kern/entrypgdir.c
+ cc kern/init.c
+ cc kern/console.c
+ cc kern/monitor.c
+ cc kern/printf.c
+ cc kern/kdebug.c
+ cc lib/printfmt.c
+ cc lib/readline.c
+ cc lib/string.c
+ ld obj/kern/kernel
ld: warning: section `.bss' type changed to PROGBITS
+ as boot/boot.S
+ cc -Os boot/main.c
+ ld boot/boot
boot block is 390 bytes (max 510)
+ mk obj/kern/kernel.img
```

如果你有类似于"undefined reference to `__udivdi3`"这种错误，你可能没有32位gcc编译包，如果你运行在Ubuntu或Debian，尝试安装`gcc-multilib`包。使用我的Dockerfile，不会出现这个问题

现在，你准备运行QEMU，装载`obj/kern/kernel.img`文件，这个文件包含引导加载程序(obj/boot/boot)和内核(obj/kernel)

运行`make qemu`(有界面)或者`make qemu-nox`(无界面)。将会启动QEMU并且加载硬盘，成功进入系统。具体显示如下

```text
6828 decimal is XXX octal!
entering test_backtrace 5
entering test_backtrace 4
entering test_backtrace 3
entering test_backtrace 2
entering test_backtrace 1
entering test_backtrace 0
leaving test_backtrace 0
leaving test_backtrace 1
leaving test_backtrace 2
leaving test_backtrace 3
leaving test_backtrace 4
leaving test_backtrace 5
Welcome to the JOS kernel monitor!
Type 'help' for a list of commands.
K> 
```

### PC物理地址空间

一个计算机的物理内存地址通常是下面的结构

```text
+------------------+  <- 0xFFFFFFFF (4GB)
|      32-bit      |
|  memory mapped   |
|     devices      |
|                  |
/\/\/\/\/\/\/\/\/\/\

/\/\/\/\/\/\/\/\/\/\
|                  |
|      Unused      |
|                  |
+------------------+  <- depends on amount of RAM
|                  |
|                  |
| Extended Memory  |
|                  |
|                  |
+------------------+  <- 0x00100000 (1MB)
|     BIOS ROM     |
+------------------+  <- 0x000F0000 (960KB)
|  16-bit devices, |
|  expansion ROMs  |
+------------------+  <- 0x000C0000 (768KB)
|   VGA Display    |
+------------------+  <- 0x000A0000 (640KB)
|                  |
|    Low Memory    |
|                  |
+------------------+  <- 0x00000000
```

第一代PC是基于16位Intel 8088处理器，仅仅只有1MB的物理内存。因此早期PC的物理地址空间是从`0x00000000`到`0x000FFFFF`,而不是`0xFFFFFFFF`。其中640KB的区域被标记为只有早期计算机能够使用的RAM(random-access memory),实际上，非常早期的PC能够配置16KB,32KB,或者64KB大小的RAM

从0x000A0000到0x000FFFFF的384KB是由硬件保留用作特殊用途的，比如视频显示缓冲区和保存在非易失内存的固件。最重要的保留区域是BIOS(Basic Input/Output System)，BIOS占用从0x000F0000到0x000FFFFF的64KB，这个区域也被称为ROM(read-only memory)，但是现在PC把BIOS保存在可更新的闪存中。BIOS负责系统的初始化，例如激活显卡，检查内存空间大小。完成初始化后，BIOS会从合适的位置(例如软盘，硬盘，CD-ROM或网络中)加载操作系统，然后把机器的控制权交给操作系统

Intel的80286和80386处理器最终打破了1MB的障碍，这两款处理器可以支持16MB和4GB物理地址空间，尽管如此，PC架构还是保留了低1MB物理地址空间的原始布局，以确保与现有软件的向后兼容性。因此现代PC在物理内存中存在一个"洞", 从0x000A0000到0x00100000,把RAM分成"低"或者"传统内存"(最初640KB)和"扩展内存"(剩下部分)。此外，PC的32位物理地址空间RAM的顶部，现在被保留为由BIOS使用的32为PCI设备

最近x86处理器能够支持超过4GB的物理RAM，所以RAM可以延伸到超过0xFFFFFFFF。这样，为了为这些32位设备预留空间映射，BIOS必须在32位可寻址区域顶部的系统RAM留出第二个洞。由于设计限制，JOS只会使用前256MB的物理内存，所以现在假设所有的PC只有一个32位物理地址空间。但是处理复杂的物理地址空间和硬件组织的其他方面是操作系统开发的重要挑战之一

### BIOS
这部分实验，你将使用QEMU的调试工具来探索IA-32兼容的电脑是怎么启动的

打开两个两个terminal窗口，进入到实验目录，输入`make qemu-nox-gdb`。这会启动QEMU，但是在QEMU会在第一条指令前停止，等待GDB的连接。在第二个terminal，输入`make gdb`，你可以看到gdb的输出

实验提供了一个`.gdbinit`文件，用来启动GDB的16位代码调试，并将其链接到正在监听的QEMU(如果没有起作用，你必须添加一个`add-auto-load-safe-path`在你的`.gdbinit`，确保`gdb`程序是按照上述的操作连到QEMU)

QEMU输出
```text
***
*** Now run 'make gdb'.
***
qemu-system-i386 -nographic -drive file=obj/kern/kernel.img,index=0,media=disk,format=raw -serial mon:stdio -gdb tcp::25000 -D qemu.log  -S
```

GDB输出
```text
warning: A handler for the OS ABI "GNU/Linux" is not built into this configuration
of GDB.  Attempting to continue with the default i8086 settings.

The target architecture is assumed to be i8086
[f000:fff0]    0xffff0:	ljmp   $0xf000,$0xe05b
0x0000fff0 in ?? ()
+ symbol-file obj/kern/kernel
```

为什么QEMU会如此启动？这与Intel设计的8088处理器有关。因为PC里的BIOS是"硬连接"到物理地址的0x000F0000~0x000FFFFF的，这个设计保证了PC在开机或重启的时候，BIOS总是可以拿到控制权，这一点是非常重要的，因为开机的时候，机器的RAM没有处理器能够执行的其他软件。QEMU仿真器有自己的BIOS，位于处理器仿真的物理地址空间。当处理器重置的时候，仿真的处理器进入到实模式，设置`CS`为`0xF00`， `IP`为`0xFFF0`，以便执行从CS:IP段地址开始。段地址0xF000:0xFFF0是怎么转换成物理地址的呢？

为了解答上面的问题，我们需要知道实模式地址。在实模式中，地址转换是根据公示$物理地址 = 16 * 段 + 偏移值$ 进行计算的。所以，当P设置`CS`为`0xF00`， `IP`为`0xFFF0`,物理地址为$ 16 * 0xF000 + 0xFFF0 = 0xFFFF0`

`0xFFFF0`是BIOS结束前(`0x100000`)的16个字节,因此，我们不应该对BIOS所做的第一件事就是`jmp`到BIOS较早的位置感到惊讶；毕竟在16字节内，能完成多少？

>进行练习2

当BIOS运行之后，它会设置中断描述符表，然后初始化各种各样的设备，例如VGA显示

初始化PC总线和所有BIOS知道的重要设备之后，它会搜索可引导设备例如软盘，硬盘或者CD-ROM。最终，当它找到可引导磁盘之后，BIOS从磁盘读取引导扇区(boot loader)并把控制权交给它

## 引导加载程序
PC的软盘和硬盘被分成512字节的区域成为扇区(*sectors*)。一个扇区是磁盘最小转换单位：每次读或写操作必须是一个或多个扇区。如果磁盘是可引导的，第一个扇区被称为引导扇区(*boot sector*),因为这个地方是引导加载程序代码所在的位置。当BIOS找到了一个可引导的软盘或硬盘，它会加载512字节的引导扇区到物理地址0x7c00～0x7dff的内存处，然后使用`jmp`指令设置CS:IP为`0000:7c000`,把控制权交给引导扇区。像BIOS加载地址一样，这些地址相当随意-但是这个地址已经固定并且标准化了

在PC发展过程中，从CD-ROM中启动的能力出现的更晚，因此PC架构师借此机会稍微重新考虑了启动过程。结果，现代BIOS从CD-ROM启动有一点复杂，但是也更强大。CD-ROM使用2048字节作为一个扇区，而不是512字节，在移交控制权之前，BIOS能够从磁盘加载更大的启动镜像到内存(不仅仅是一个扇区)。对于更多的细节可以参考["El Torito" Bootable CD-ROM Format Specification](https://pdos.csail.mit.edu/6.828/2018/readings/boot-cdrom.pdf)

然而，对于6.828，我们使用传统的硬件驱动引导机制，这也意味着我们的引导扇区必须是512字节。引导扇区由汇编语言源码文件`boot/boot.S`和一个C源码文件`boot/main.c`组成。仔细查阅这些源码文件，确保你理解它是怎么运行的，引导扇区有两个主要的功能：
1. 引导扇区把处理器从16位实模式切换到32为保护模式(*protected mod*),因为只有在保护模式下，软件才能使用超过1MB的物理地址空间。保护模式在[PC Assembly Language](https://pdos.csail.mit.edu/6.828/2018/readings/pcasm-book.pdf)的1.2.7和1.2.8部分有简短的介绍，Intel架构手册中有关于保护模式的详细介绍。此时，你只需要理解段地址转换(段地址:偏移值)在保护模式下转换时不同的，转换之后是32位而不是16位
2. 引导扇区通过x86特殊I/O指令访问IDE磁盘设备寄存器，以此来从磁盘读取内核。如果你想对这里的特殊I/O指令了解更深，查阅[6.828参考页](https://pdos.csail.mit.edu/6.828/2018/reference.html)的"IDE hard drive controller"部分内容。

理解了引导扇区的源码后，查看`obj/boot/boot.asm`文件，这个文件是引导扇区的反汇编代码，是由GNU的`makefile`在编译引导扇区代码之后创建的。这个反汇编文件可以很容易理解所有引导扇区代码在物理内存中的位置，也让GDB跟踪下一步代码发生了什么变得更加容易。同样的，`obj/kern/kernel.asm`也包含了JOS内核的反汇编

你可以在GDB用`b`命令设置地址断点。例如`b *0x7c00`，在0x7C00设置一个断点，一旦运行到断点，你可以使用`c`和`si`命令继续执行，`c`是断点执行，`si`是单步执行

>进行练习3

### 加载内核
我们将进入到引导扇区的C语言代码部分，去了解细节，在`boot/main.c`。但是在这之前，是时候停下来，阅读一些C程序的基础知识了。

>进行练习4

为了理解`boot/main.c`，你需要了解什么是ELF二进制文件。当你编译和链接C程序(例如JOS内核)时，编译器把每个C源码(*.c*)编译成包含汇编语言的目标文件(*.o*),链接器然后把所有编译的目标文件组合到一个二进制镜像，例如`obj/kern/kernel`,这个二进制文件就是ELF格式，标准名称为"Executable and Linkable Format"

关于这个格式的详细信息可以参考[ELF specification](https://pdos.csail.mit.edu/6.828/2018/readings/elf.pdf),但是你不必深入到这个格式的每个细节。尽管整个格式是非常强大和复杂的，但是大部分复杂的部分都支持动态链接库动态加载，[维基百科](http://en.wikipedia.org/wiki/Executable_and_Linkable_Format)有简短的介绍

对于6.828，你可以把ELF可执行文件看成包含头部和一些程序部分，为了加载到指定地址的内存，每个部分都是连续的代码块和数据块。启动扇区不会修改代码或数据，它会加载到内存然后开始执行

一个ELF二进制文件以一个定长ELF头部开始，然后是可变长度的程序头部，程序头部列出了每个程序会加载多少扇区。C定义了这些ELF头部在`inc/elf.h`，我们感兴趣的代码部分是：
- `.text`: 程序可执行的指令
- `.rodata`: 只读数据，比如ASCII字符
- `.data`: 数据部分包含程序初始化数据，例如全局变量

当链接器计算程序的内存结构，它会为没有初始化的全局变量保留空间，这个部分被称为`.bss`，紧接者`.data`之后。C会把没有初始化的全局变量初始化成0。因此ELF二进制文件中的`.bss`没有内容。然而，链接器仅仅记录了`.bss`部分的地址和大小。加载器或程序本身必须把0分配给`.bss`部分

测试所有部分的名称，大小和链接地址，可以用指令`objdump -h obj/kern/kernel`
```text
Sections:
Idx Name          Size      VMA       LMA       File off  Algn
  0 .text         00001871  f0100000  00100000  00001000  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
  1 .rodata       00000714  f0101880  00101880  00002880  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  2 .stab         000038d1  f0101f94  00101f94  00002f94  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  3 .stabstr      000018bb  f0105865  00105865  00006865  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  4 .data         0000a300  f0108000  00108000  00009000  2**12
                  CONTENTS, ALLOC, LOAD, DATA
  5 .bss          00000648  f0112300  00112300  00013300  2**5
                  CONTENTS, ALLOC, LOAD, DATA
  6 .comment      00000035  00000000  00000000  00013948  2**0
                  CONTENTS, READONLY
```

可以看到不仅仅是上面列的那些内容，但是其他的都不重要。其他的大部分都是保存调试信息的，实际运行过程中是不会加载到内存的

特别注意`.text`部分的"VMA"(或*link address*)和"LMA"(或*load address*)，LMA是加载的地址，VMA是链接的地址。链接器以各种方式对链接地址进行二进制编码，例如当代码需要全局变量地址是，结果是如果从一个没有链接的地址执行，二进制通常不能工作。值得一提的是，可以生成不包含任何绝对地址的代码，这就是通常说的动态链接库，但是6.828不使用

通常，链接和加载地址是相同的，例如可以查看`.text`部分的引导扇区内容`objdump -h obj/boot/boot.out`。可以看到VMA和LMA都是0x7c00，说明引导扇区从这个地方开始加载程序
```text
Sections:
Idx Name          Size      VMA       LMA       File off  Algn
  0 .text         00000186  00007c00  00007c00  00000074  2**2
                  CONTENTS, ALLOC, LOAD, CODE
  1 .eh_frame     000000a8  00007d88  00007d88  000001fc  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  2 .stab         00000720  00000000  00000000  000002a4  2**2
                  CONTENTS, READONLY, DEBUGGING
  3 .stabstr      0000088f  00000000  00000000  000009c4  2**0
                  CONTENTS, READONLY, DEBUGGING
  4 .comment      00000035  00000000  00000000  00001253  2**0
                  CONTENTS, READONLY
```

引导扇区使用ELF程序头去决定怎么加载片段，程序头指定了要加载的内容和目标地址。你可以检查程序头，通过`objdump -x obj/kern/kernel`

```text
obj/kern/kernel:     file format elf32-i386
obj/kern/kernel
architecture: i386, flags 0x00000112:
EXEC_P, HAS_SYMS, D_PAGED
start address 0x0010000c

Program Header:
    LOAD off    0x00001000 vaddr 0xf0100000 paddr 0x00100000 align 2**12
         filesz 0x00007120 memsz 0x00007120 flags r-x
    LOAD off    0x00009000 vaddr 0xf0108000 paddr 0x00108000 align 2**12
         filesz 0x0000a948 memsz 0x0000a948 flags rw-
   STACK off    0x00000000 vaddr 0x00000000 paddr 0x00000000 align 2**4
         filesz 0x00000000 memsz 0x00000000 flags rwx

Sections:
Idx Name          Size      VMA       LMA       File off  Algn
  0 .text         00001871  f0100000  00100000  00001000  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
  1 .rodata       00000714  f0101880  00101880  00002880  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  2 .stab         000038d1  f0101f94  00101f94  00002f94  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  3 .stabstr      000018bb  f0105865  00105865  00006865  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  4 .data         0000a300  f0108000  00108000  00009000  2**12
                  CONTENTS, ALLOC, LOAD, DATA
  5 .bss          00000648  f0112300  00112300  00013300  2**5
                  CONTENTS, ALLOC, LOAD, DATA
  6 .comment      00000035  00000000  00000000  00013948  2**0
                  CONTENTS, READONLY
```

结果中的程序头部包含了`ELF`的相关信息，需要被加载进内存的用"LOAD"标记。其他信息，如`vaddr`是虚拟地址，`paddr`是物理地址,`memsz`和`filesz`是加载区域。在`boot/main.c`中的代码。

>进行练习5

回到`boot/main.c`程序，每个程序段的`ph->p_pa`字段包含段目标的物理地址，这种情况下，也是实际的物理地址

BIOS把引导扇区加载到0x7c00的内存地址位置，因此这是引导扇区的加载地址。这也是引导扇区执行的起始位置，所以也是链接地址。链接地址是通过`boot/Makefrag`文件中通过`-Ttext`来设置的，所以，链接器在生成代码过程中将产生正确的内存地址

现在回头来看看内核加载地址和链接地址。不像启动引导，这两个地址是不同的：内核告诉引导加载程序加载低地址(1MB)的内存,但是可能从一个高地址执行，将会在下一个部分深挖

除了段信息，ELF头部也有一个重要的部分，叫做`e_entry`，这个部分保留了程序入口的链接地址。可以通过`objdump -f obj/kern/kernel`查看入口地址
```text
obj/kern/kernel:     file format elf32-i386
architecture: i386, flags 0x00000112:
EXEC_P, HAS_SYMS, D_PAGED
start address 0x0010000c
```

现在看来最小的`boot/main.c`中加载ELF，就是把内核从硬盘加载到内存，然后跳到内核入口

>进行练习6

## 内核
尝试了解更多关于最小JOS内核的一些细节。跟引导加载程序一样，内核也有设置事件的汇编代码和能够执行的C语言代码

### 使用虚拟内存
使用虚拟内存是为了解决位置依赖问题。

上面通过`objdump`指令查询了内核的加载地址和链接地址。内核的链接比引导加载程序的链接要复杂，所以链接和加载地址在`kern/kernel.ld`的最顶部

操作系统链接和运行在高的虚拟地址，例如0xf0100000，以便于将处理器虚拟地址空间的低位部分留给用户使用。具体原因，下一个lab会更加清晰

很多机器都没有0xf0100000的物理地址，所以不能直接把内核保存在这里。相反，我们需要用处理器内存管理硬件把虚拟地址0xf0100000(内核链接和运行的地址)映射到物理地址0x00100000(内核加载程序加载的物理地址)。如此，尽管内核虚拟地址足够用户程序使用，它也会被加载到计算机RAM(在ROM上面)的1MB物理地址内。这个方法，PC需要至少几MB的物理地址，但是这可以适用于1990年以后生产的任何计算机

实际上，在下一个lab中，将会把256MB映射到PC的物理内存地址，你现在明白为啥JOS可以仅使用前256MB的物理内存

现在，仅仅只映射前4MB的物理内存，这也足够运行了。我们使用手写的，静态初始化的页表目录和页表来实现，具体代码可以参考`kern/entrypgdir.c`。现在不必理解这个工作的细节。`kern/entry.S`设置`CR0_PG`为1，内存引用被当作物理内存(严格来说，应该是线性地址，但是`boot/boot.S`会把线性地址映射到物理地址)。一旦`CR0_PG`被置位成功，内存引用就是虚拟地址。`entry_pgdir`转换虚拟地址到物理地址，同样也把物理地址转换成虚拟地址。任何不在这个范围内(0f0000000~0xf0400000)的虚拟地址都会导致硬件异常，由于没有处理这些异常，就会导致QEMU退出

>进行练习7

### 格式化打印输出
大多数人都会使用`printf()`来输出内容，有时候甚至会优先考虑C语言，但是在OS内核中，必须实现所有I/O

阅读`kern/printf.c` `lib/printfmt.c` 和 `kern/console.c` 确保你理解了其中的关系，你后面会理解为什么`printfmt.c`是位于`lib`目录

`kern/printf.c`里面主要是`print`的接口，一共有三个`putch` `vcprintf` `cprintf` 其中`putch`最终调用的是`kern/console.c`里面实现的`cputchar`打印一个字符；`vcprintf`调用的是在`lib/printfmt.c`中实现的`vprintfmt`，用来格式化字符串，最终也会调用`putch`

>进行练习8

### 栈
这个实验的最后一个练习，这个练习会更详细讲解C语言在x86中使用栈，并且在程序中写一个新的监控函数来打印栈的`backtrace`：即函数调用栈信息

>进行练习9

x86栈指针(`esp`寄存器)指向当前正在使用栈的最低内存位置，保留区域中低于这个位置的都是可以使用的。入栈操作会先把栈指针变小，然后把这个值写入栈指针指向的位置。出栈动作是先从栈中读取数值，然后再把栈指针增加。在32位模式中，栈只能保存32位的数值，`esp`寄存器也总是被拆分成4个。不同的x86指令，例如`call`，是硬连接(hard-wired)使用栈指针寄存器的

相比之下，`ebp`(基指针)寄存器，主要通过软件约定与栈相关联。在进入C函数入口前，`prologue`代码通常通过入栈操作保存之前函数的基指针，然后在函数使用期间，拷贝当前`esp`的值到`ebp`。程序中所有的函数都遵循这个规则，在程序运行中的任何一个给定的时刻，都是可以通过b保存的`ebp`的调用链来获取调用栈的数据。这个能力非常有用，例如当程序`assert`失败或者`panic`的时候，调用栈可以跟踪哪个地方有问题

>进行练习10

后面都是一些废话了，具体就是实现`mon_backtrace`达到不同的功能

>进行练习11

>进行练习12