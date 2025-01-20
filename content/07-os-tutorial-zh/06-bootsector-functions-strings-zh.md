---
author: xnzone 
title: 引导扇字符串函数
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 706
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：control structures, function calling, strings

**目标：学会使用汇编编码**

我们将进一步自定义引导扇区

在第7节课， 我们将从硬盘中开始读取引导扇区代码，这是在加载内核之前的最后一步。但是在这之前，我们需要用控制结构、函数调用、字符串编写相关代码。在跳转到硬盘和内核之前，我们必须适应这些概念

## Strings(字符串)

定义string类似于bytes，但是为了能够终止他们，使用一个null-byte终止（是的，类似于C）

```asm
mystring:
    db 'Hello, World', 0
```

注意，引用的文本会被汇编转换成ASCII，然而单个的0会被转换成空字节(null byte)`0x00`

## Control structures(控制结构)

我们已经使用了一个`jmp $`用于无限循环的控制结构

汇编跳转被定义成之前指令结果。例如

```asm
cmp ax, 4       ; 如果 ax = 4
je ax_is_four   ; 满足条件的处理，跳转到这个label进行处理
jmp else        ; else 做其他的事情
jmp endif       ; 最后，正常结束

ax_is_four:
    ......
    jmp endif

else:
    ......
    jmp endif ;通常不需要，但是还是在这个地方输出

endif:
```

在你的脑海里使用高级语言思考这个过程，然后使用汇编实现它

有很多`jmp`条件：if equal, if less than, 等等。谷歌上有更完美的解释

## Calling functions(函数调用)

正如你猜想的一样，调用一个函数就是跳转到某个label

棘手的部分是参数传递。有两个步骤用于解决参数传递：

1. 程序员知道他们共享一个指定的寄存器或者内存地址
2. 编写一小部分代码，创建一个函数调用，并且没有其他影响

步骤1是非常简单的，`al`(实际上是`ax`)寄存器可以用于参数传递

```asm
mov al, 'X'
jmp print
endprint:

...

print:
    mov ah, 0x0e    ; tty模式
    int 0x10        ; 假设 'al' 已经有所有的字符串
    jmp endprint    ; 之前已经有的label
```

如你所见，这个方法很快就让代码成长为面条式代码。当前的`print`函数，仅仅返回`endprint`，如果其他的函数调用它，会怎么样呢？代码就不能被复用

正确的解决方案提供了两点改进

- 保存返回地址，以便于可以更广泛的使用
- 保存当前寄存器，以便于子函数可以无副作用的修改

为了保存返回地址，需要CPU的帮助。不是使用一系列的`jmp`去调用子协程，而是使用`call`和`ret`

为了保存寄存器的数据，有一些特殊的栈操作`pusha`和`popa`,可以把所有寄存器都可以自动压入栈和恢复

## Including external files(引入外部文件)

假设你是一个程序员，不需要向你解释为啥这是一个好主意

代码是

```asm
%include "file.asm"
```

## Printing hex values(输出二进制数据)

下一节课程，我们将从硬盘读取，所以我们需要一些方式确保我们读取到正确的数据。文件`boot_sect_print_hex.asm` 扩展了`boot_sect_print.asm` 以便于可以输出二进制数据，不仅仅是ASCII

## Code!(编码)

开始编码。文件`boot_sect_print.asm`是一个子协程，被main文件引用了。使用一个循环打印字节到屏幕上。也包括一个函数来打印换行。熟悉的`\n`占用两个字节，换行符是`0x0A`,返回是`0x0D`。可以实验，如果移除返回字节，看一下会有什么影响

如上所述，`boot_sect_print_hex.asm`允许打印字节

主程序文件`boot_sect_main.asm` 加载一系列字符串和字节，调用`print`和`print_hex`，如果你已经理解上面的部分，你可以直接开始。