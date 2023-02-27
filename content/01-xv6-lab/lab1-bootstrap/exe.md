---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: /covers/xv6.png
cover: false 
weight: 2
tags: ["xv6", "os", "bootstrap"]
---

## Exercise 1
>Exercise 1. Familiarize yourself with the assembly language materials available on the 6.828 reference page. You don't have to read them now, but you'll almost certainly want to refer to some of this material when reading and writing x86 assembly.

>We do recommend reading the section "The Syntax" in Brennan's Guide to Inline Assembly. It gives a good (and quite brief) description of the AT&T assembly syntax we'll be using with the GNU assembler in JOS.

练习1的目的是为了熟悉汇编语言，但是不必马上就去学习编写，可以在阅读代码和编写代码的时候去查阅。但是强烈要求看一下Brennan's的关于"The Syntax"这部分的内容

## Exercise 2
>Exercise 2.Use GDB's si (Step Instruction) command to trace into the ROM BIOS for a fewmore instructions, and try to guess what it might be doing. You might want to look at PhilStorrs I/O Ports Description, as well as other materials on the 6.828 reference materials page.No need to figure out all the details - just the general idea of what the BIOS is doing first

练习2的目的是为了使用GDB调试程序，可以尝试猜测接下来会干什么。你可能想要阅读[Phil Storrs PC Hardware book](http://web.archive.org/web/20040404164813/members.iweb.net.au/~pstorr/pcbook/book2/book2.htm)或者6.828参考材料中的其他的材料。没有必要理解所有的细节，仅仅只需要了解BIOS最开始做什么

## Exercise 3
>Exercise 3. Take a look at the lab tools guide, especially the section on GDB commands. Even if you're familiar with GDB, this includes some esoteric GDB commands that are useful for OS work.

>Set a breakpoint at address 0x7c00, which is where the boot sector will be loaded. Continue execution until that breakpoint. Trace through the code in boot/boot.S, using the source code and the disassembly file obj/boot/boot.asm to keep track of where you are. Also use the x/i command in GDB to disassemble sequences of instructions in the boot loader, and compare the original boot loader source code with both the disassembly in obj/boot/boot.asm and GDB.

>Trace into bootmain() in boot/main.c, and then into readsect(). Identify the exact assembly instructions that correspond to each of the statements in readsect(). Trace through the rest of readsect() and back out into bootmain(), and identify the begin and end of the for loop that reads the remaining sectors of the kernel from the disk. Find out what code will run when the loop is finished, set a breakpoint there, and continue to that breakpoint. Then step through the remainder of the boot loader.

练习3的目的主要是为了熟悉使用GDB，同时观察引导扇区和内核的具体运作方式。要回答四个问题

1.At what point does the processor start executing 32-bit code? What exactly causes the switch from 16- to 32-bitmode?
{{< highlight asm >}}
ljmp    $PROT_MODE_CSEG, $protcseg
{{< /highlight  >}}
从这句开始，执行完就从16位实模式切换到32位保护模式了

2.What is the last instruction of the boot loader executed, and what is the first instruction of the kernel it justloaded?
{{< highlight c >}}
((void (*)(void)) (ELFHDR->e_entry))();
{{< /highlight  >}}

bootloader最后一行代码是这一条，整个`bootmain`函数的作用就是从硬盘读取内核，然后跳到`entry`入口，执行内核。镜像文件是按照`elf`格式存在硬盘上的，`ELFHDR`是指向`0x10000`，整个内核程序块是从`0x10000`（物理地址）开始运行的。`entry`的虚拟地址，在之前打印过，是`0xf010000c`,转化成物理地址为`0x10000c`。所以内核加载的第一条指令是`movw   $0x1234,0x472` 

3.Where is the first instruction of the kernel?
{{< highlight asm >}}
movw   $0x1234,0x472
{{< /highlight  >}}

4.How does the boot loader decide how many sectors it must read in order to fetch the entire kernel from disk?Where does it find this information?

`elf`格式文件有写，具体结构可以参考[ELF文件结构](https://baike.baidu.com/item/ELF/7120560)，读取操作看`main.c`的`readsect`函数。最终概括成以下这句话

一个扇区大小为512字节。读一个扇区的流程大致为通过outb指令访问I/O地址:0x1f2~-0x1f7来发出读扇区命令，通过in指令了解硬盘是否空闲且就绪，如果空闲且就绪，则通过inb指令读取硬盘扇区数据都内存中。

ELF文件具体格式如下

{{< highlight text >}}
I/O地址功能
0x1f0读数据，当0x1f7不为忙状态时，可以读。
0x1f2要读写的扇区数，每次读写前，需要指出要读写几个扇区。
0x1f3如果是LBA模式，就是LBA参数的0-7位
0x1f4如果是LBA模式，就是LBA参数的8-15位
0x1f5如果是LBA模式，就是LBA参数的16-23位
0x1f6第0~3位：如果是LBA模式就是24-27位第4位：为0主盘；为1从盘
第6位：为1=LBA模式；0= CHS模式第7位和第5位必须为1
0x1f7状态和命令寄存器。操作时先给命令，再读取内容；如果不是忙状态就从0x1f0端口读数据
{{< /highlight  >}}


## Exercise 4
>Exercise 4. Read about programming with pointers in C. The best reference for the C language is The C Programming Language by Brian Kernighan and Dennis Ritchie (known as 'K&R'). We recommend that students purchase this book (here is an Amazon Link) or find one of MIT's 7 copies.

>Read 5.1 (Pointers and Addresses) through 5.5 (Character Pointers and Functions) in K&R. Then download the code for pointers.c, run it, and make sure you understand where all of the printed values come from. In particular, make sure you understand where the pointer addresses in printed lines 1 and 6 come from, how all the values in printed lines 2 through 4 get there, and why the values printed in line 5 are seemingly corrupted.

>There are other references on pointers in C (e.g., A tutorial by Ted Jensen that cites K&R heavily), though not as strongly recommended.

>Warning: Unless you are already thoroughly versed in C, do not skip or even skim this reading exercise. If you do not really understand pointers in C, you will suffer untold pain and misery in subsequent labs, and then eventually come to understand them the hard way. Trust us; you don't want to find out what "the hard way" is.

练习4的目的是为了熟悉C语言代码规则，主要是看K&R的前5章内容，尤其是第五章关于指针部分的内容。可以自行查看

## Exercise 5
>Exercise 5. Trace through the first few instructions of the boot loader again and identify the first instruction that would "break" or otherwise do the wrong thing if you were to get the boot loader's link address wrong. Then change the link address in boot/Makefrag to something wrong, run make clean, recompile the lab with make, and trace into the boot loader again to see what happens. Don't forget to change the link address back and make clean again afterward!

练习5的目的是为了验证`boot/Makefrag`设置链接地址的正确性。需要`make clean`, 然后修改`0x7c00`的数值，尝试实验就可以。其实可以发现并不能启动

## Exercise 6
>Exercise 6.We can examine memory using GDB's x command. The GDB manual has fulldetails, but for now, it is enough to know that the command x/Nx ADDR prints N words ofmemory at ADDR. (Note that both 'x's in the command are lowercase.) Warning: The size of aword is not a universal standard. In GNU assembly, a word is two bytes (the 'w' in xorw, whichstands for word, means 2 bytes).

>Reset the machine (exit QEMU/GDB and start them again). Examine the 8 words of memoryat 0x00100000 at the point the BIOS enters the boot loader, and then again at the point the bootloader enters the kernel. Why are they different? What is there at the second breakpoint? (You do not really need to use QEMU to answer this question. Just think.)

练习6主要是使用GDB去查看启动时，加载过程，可以在`0x7c00`和`0x0010000c`地址处加断点，使用`x/Nx ADDR`打印加载的地址，主要是看`0x00100000`的内容。可以看到`0x7c00`的时候，全都是空，此时没有数据，运行到`0x0010000c`处时，已经有数据了，说明已经加载了ELF的内容进内存了

GDB主要的命令是
{{< highlight gdb >}}
b *0x7c00 # 0x7c00处设置断点
b *0x00100000 # 0x00100000处设置断点

c # 继续执行代码到0x7c00处
x /8w 0x0010000 # 查看0x0010000内存处的内容
x /8w 0xf010000 # 查看 0xf010000 内存处的内容

c # 继续执行代码到0x00100000处
x /8w 0x0010000 # 查看0x0010000内存处的内容
x /8w 0xf010000 # 查看 0xf010000 内存处的内容
{{< /highlight  >}}


## Exercise 7
>Exercise 7. Use QEMU and GDB to trace into the JOS kernel and stop at the movl %eax, %cr0. Examine memory at 0x00100000 and at 0xf0100000. Now, single step over that instruction using the stepi GDB command. Again, examine memory at 0x00100000 and at 0xf0100000. Make sure you understand what just happened.

>What is the first instruction after the new mapping is established that would fail to work properly if the mapping weren't in place? Comment out the movl %eax, %cr0 in kern/entry.S, trace into it, and see if you were right.

练习7其实跟练习6是类似的，是需要单步执行，去查看内核加载到底发生了什么。查看`obj/kern/kernel.asm`找到`movl %eax, %cr0`所在的地址，打一个断点，然后查看0x00100000和0xf0100000地址的内容，然后单步运行，在启用保护模式之后再查看两个地址的内容，就可以知道地址映射有区别

{{< highlight gdb >}}
b *0x10000c # 在0x10000c设置断点,在入口处打断点
si # 单步执行
x /5w 0x0010000 # 查看内存内容
x /5w 0xf010000
{{< /highlight  >}}

## Exercise 8
>Exercise 8. We have omitted a small fragment of code - the code necessary to print octal numbers using patterns of the form "%o". Find and fill in this code fragment.

练习8的目的是为了熟悉`printfmt.c`的代码。查阅`kern/printf.c` `lib/printfmt.c`和`kern/console.c`的代码， 然后填充关于"%o"输出八进制的代码

具体代码如下：

{{< highlight c >}}
case 'o':
// Replace this with your code.
	num = getuint(&ap, lflag);
	base = 8;
	goto number;	
	break;
{{< /highlight  >}}

1.Explain the interface between printf.c and console.c. Specifically, what function does console.c export? How is this function used by printf.c?

`console.c`主要是在屏幕上输出内容，`printf.c`是根据不同的样式去打印输出，最终还是会调用`console.c`里面的内容

2.Explain the following from console.c:
{{< highlight c >}}
if (crt_pos >= CRT_SIZE) {
    int i;
    memmove(crt_buf, crt_buf + CRT_COLS, (CRT_SIZE - CRT_COLS) * sizeof(uint16_t));
    for (i = CRT_SIZE - CRT_COLS; i < CRT_SIZE; i++)
        crt_buf[i] = 0x0700 | ' ';
    crt_pos -= CRT_COLS;
}
{{< /highlight  >}}

代码是在`console.c`里面的`cga_putc`函数，而这部分是根据字符来输出什么内容的，但是当输出大于行的最大缓冲区时，会换行显示，所以这个代码片段就是这个作用

3.For the following questions you might wish to consult the notes for Lecture 2. These notes cover GCC's calling convention on the x86.

>Trace the execution of the following code step-by-step:

{{< highlight c >}}
int x = 1, y = 3, z = 4;
cprintf("x %d, y %x, z %d\n", x, y, z);
{{< /highlight  >}}

>In the call to cprintf(), to what does fmt point? To what does ap point?

>List (in order of execution) each call to cons_putc, va_arg, and vcprintf. For cons_putc, list its argument as well. For va_arg, list what ap points to before and after the call. For vcprintf list the values of its two arguments.

[Lecture 2](https://pdos.csail.mit.edu/6.828/2018/lec/l-x86.html).对`cprintf`函数（定义在`kern/printf.c`）而言，fmt是一个指针，指向"x %d, y %x, z %d\n"这个【格式字符串】。而ap指向后面的变量列表var_list的起始地址，本例中可以理解为x的地址.

总之是利用一个函数内的所有东西大家都挤在一个栈区这个特点。最后的变量先入栈，最前的变量最后入栈，于是就成了栈顶。既然我们知道了第一个指针fmt的值，就很容易找到紧挨着的下一个位置ap的值，顺着ap一路找下去就能遍历整个var_list了

4.Run the following code.
{{< highlight c >}}
    unsigned int i = 0x00646c72;
    cprintf("H%x Wo%s", 57616, &i);
{{< /highlight  >}}

5.What is the output? Explain how this output is arrived at in the step-by-step manner of the previous exercise. Here's an ASCII table that maps bytes to characters.

>The output depends on that fact that the x86 is little-endian. If the x86 were instead big-endian what would you set i to in order to yield the same output? Would you need to change 57616 to a different value?

>Here's a description of little- and big-endian and a more whimsical description.

输出`He110 World`由于是小端编码，所以`0x00646c72`对应存储的是`72` `6c` `64` 对应的ASCII码就是`r` `l` `d`, `57616`对应的16进制是`e110`，所以输出是`He110 World`

其实可以把这个代码复制到`kern/init.c`的`i386_init`函数中，去执行，也可以得到相同的答案。因为内核启动的时候，会在汇编中调用`i386_init`，所以在启动的时候，就可以看到运算结果

>In the following code, what is going to be printed after 'y='? (note: the answer is not a specific value.) Why does this happen?
{{< highlight c >}}
cprintf("x=%d y=%d", 3);
{{< /highlight  >}}

输出`x=3 y=随机数` 会把`x`所在地址+4Bit之后当作`y`的地址，然后取出这个数字打印，所以这个值是随机的。同样的也可以把代码贴进去验证

6.Let's say that GCC changed its calling convention so that it pushed arguments on the stack in declaration order, so that the last argument is pushed last. How would you have to change cprintf or its interface so that it would still be possible to pass it a variable number of arguments?

改变gcc压栈方式，in declaration order。 就是以声明的顺序，最早的变量最先入栈，到了栈底。这样栈顶指针，最后就指向了最后一个变量。此时可以添加一个计算变量总长度的参数`size_t n`。我们拿到栈顶指针`ap`之后，先 `(void *)fmt = ap - n`。就找到了格式化字符串的开头了。

## Exercise 9
>Exercise 9. Determine where the kernel initializes its stack, and exactly where in memory its stack is located. How does the kernel reserve space for its stack? And at which "end" of this reserved area is the stack pointer initialized to point to?

练习9主要是了解栈的初始化，栈的内存地址是位于哪里。主要是一个阅读内容。

x86栈指针(`esp`寄存器)指向当前正在使用栈的最低内存位置，保留区域中低于这个位置的都是可以使用的。入栈操作会先把栈指针变小，然后把这个值写入栈指针指向的位置。出栈动作是先从栈中读取数值，然后再把栈指针增加。在32位模式中，栈只能保存32位的数值，`esp`寄存器也总是被拆分成4个。不同的x86指令，例如`call`，是硬连接(hard-wired)使用栈指针寄存器的

## Exercise 10 
>Exercise 10. To become familiar with the C calling conventions on the x86, find the address of the test_backtrace function in obj/kern/kernel.asm, set a breakpoint there, and examine what happens each time it gets called after the kernel starts. How many 32-bit words does each recursive nesting level of test_backtrace push on the stack, and what are those words?

>Note that, for this exercise to work properly, you should be using the patched version of QEMU available on the tools page or on Athena. Otherwise, you'll have to manually translate all breakpoint and memory addresses to linear addresses.

练习10的主要目的是看函数调用的时候，到底会保存哪些数据，需要在`kern/init.c`的`test_backtrace`的地方打断点调试查看



## Exercise 11 
>Exercise 11. Implement the backtrace function as specified above. Use the same format as in the example, since otherwise the grading script will be confused. When you think you have it working right, run make grade to see if its output conforms to what our grading script expects, and fix it if it doesn't. After you have handed in your Lab 1 code, you are welcome to change the output format of the backtrace function any way you like.

>If you use read_ebp(), note that GCC may generate "optimized" code that calls read_ebp() before mon_backtrace()'s function prologue, which results in an incomplete stack trace (the stack frame of the most recent function call is missing). While we have tried to disable optimizations that cause this reordering, you may want to examine the assembly of mon_backtrace() and make sure the call to read_ebp() is happening after the function prologue.

就是实现`backtrace`函数，直接代码
{{< highlight c >}}
int
mon_backtrace(int argc, char **argv, struct Trapframe *tf)
{
    uint32_t ebp, *ptr_ebp;
    ebp = read_ebp();
    cprintf("Stack backtrace:\n");
    while (ebp != 0) {
        ptr_ebp = (uint32_t *)ebp;
        cprintf("\tebp %x  eip %x  args %08x %08x %08x %08x %08x\n", 
                ebp, ptr_ebp[1], ptr_ebp[2], ptr_ebp[3], ptr_ebp[4], ptr_ebp[5], ptr_ebp[6]);
        ebp = *ptr_ebp;
    }
    return 0;
}
{{< /highlight  >}}

主要是根据提示来改写 kern/monitor.c，重点用到的三个tricks：

利用read_ebp() 函数获取当前ebp值
利用 ebp 的初始值0判断是否停止
利用数组指针运算来获取 eip 以及 args

## Exercise 12
>Exercise 12. Modify your stack backtrace function to display, for each eip, the function name, source file name, and line number corresponding to that eip.

同样的也是实现，但是需要命令行输入，同时还要实现一个代码的二分查找

`kern/kdebug.c`中`debuginfo_eip`需要填充部分代码如下
{{< highlight c >}}
// Hint:
//	There's a particular stabs type used for line numbers.
//	Look at the STABS documentation and <inc/stab.h> to find
//	which one.
// Your code here.
stab_binsearch(stabs, &lline, &rline, N_SLINE, addr);
if (lline <= rline) {
    info->eip_line = stabs[lline].n_desc;
} else {
    return -1;
}
{{< /highlight  >}}

最后再实现一下`mon_backtrace`就可以了,同时更新一下`command`就可以了
{{< highlight c >}}
static struct Command commands[] = {
	{ "help", "Display this list of commands", mon_help },
	{ "kerninfo", "Display information about the kernel", mon_kerninfo },
	{"backtrace", "Backtrace monitor", mon_backtrace },
};

int
mon_backtrace(int argc, char **argv, struct Trapframe *tf)
{
	// Your code here.
    uint32_t ebp, *ptr_ebp;
    struct Eipdebuginfo info;
    ebp = read_ebp();
    cprintf("Stack backtrace:\n");
    while (ebp != 0) {
        ptr_ebp = (uint32_t *)ebp;
        cprintf("ebp %x  eip %x  args %08x %08x %08x %08x %08x\n", ebp, ptr_ebp[1], ptr_ebp[2], ptr_ebp[3], ptr_ebp[4], ptr_ebp[5], ptr_ebp[6]);
        if (debuginfo_eip(ptr_ebp[1], &info) == 0) {
            uint32_t fn_offset = ptr_ebp[1] - info.eip_fn_addr;
            cprintf("\t\t%s:%d: %.*s+%d\n", info.eip_file, info.eip_line,info.eip_fn_namelen,  info.eip_fn_name, fn_offset);
        }
        ebp = *ptr_ebp;
    }
    return 0;
}
{{< /highlight  >}}