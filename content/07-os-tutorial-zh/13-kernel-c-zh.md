---
author: xnzone 
title: 内核系统C程序
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 713
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：C, object code, linker, disassemble

**目标：像使用汇编写的同样低级代码一样，使用C去编写**

## 编译

让我们看一下C编译器怎么编译代码，把它和汇编生成的机器码做比较

我们开始写一个简单包含函数的程序`function.c`，打开文件并实验

为了编译系统无关的代码，我们使用标记`--ffreestanding`。所以编译`function.c`命令为：

```bash
i386-elf-gcc -ffreestanding -c function.c -o function.o
```

用编译器查看生成的机器码

```bash
i386-elf-objdump -d function.o
```

我们认识其中的一些东西，不是吗？

## 链接

最后，为了生成二进制文件，我们需要链接器。这个步骤最重要的部分是学会高级语言怎么调用函数。我们的函数在内存里会被放到哪个位置？我们实际上并不知道。例如，我们把偏移量定为`0x0` 且使用生成不带label或metadata的机器码`binary`格式

```bash
i386-elf-ld -o function.bin -Ttext 0x0 -oformat binary function.o
```

*注意：可能会有warning，请忽略它*

现在使用`xxd`分别测试`function.o`和`function.bin`。你可以看到`.bin`文件时机器码，然而`.o`有很多调试信息，labels等

## 反编译

出于好奇，我们测试一下机器码

```bash
ndisasm -b 32 function.bin
```

## 补充

我鼓励你写更多小程序，比如

- 局部变量 `localvars.c`
- 函数调用 `functioncall.c`
- 指针 `pointers.c`

然后编译和反编译他们，并且测试最终的机器码。可以参考os-guide.pdf的解释。尝试解答这个问题：为什么`pointers.c`反编译的结果与你期望的不一样？"Hello"的ASCII码`0x48656c6c6f`在哪里？