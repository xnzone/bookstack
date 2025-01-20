---
author: xnzone 
title: 断点调试
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 715
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：monolithic kernel, microkernel, debugger, gdb

**目标：暂停和组织你的代码一小会儿。然后学习怎么用gdb调试内核**

可能你没有意识到，但是你i经有你自己的运行内核了

然而，它确实非常小，仅仅只打印一个'X'。现在是时候停下来把代码组织到文件夹了，为之后的代码创建一个可伸缩的Makefile， 并且开始考虑策略

看一下新的目录结构。在前面的课程中，大多数的文件已经被符号链接了，所以我们必须在某些点修改，移除符号链接并且创建一个新文件可能是个好主意

而且，从现在开始，大多数的代码都是C，我们将利用qemu可以打开gdb连接的能力。首先，让我们安装一个交叉编译的`gdb`，因为OSX使用不能兼容ELF文件的`lldb`(Homebrew仓库也没有`gdb`)

```bash
cd /tmp/src
curl -O http://ftp.rediris.es/mirror/GNU/gdb/gdb-7.8.tar.gz
tar xf gdb-7.8.tar.gz
mkdir gdb-build
cd gdb-build
export PREFIX="/usr/local/i386elfgcc"
export TARGET=i386-elf
../gdb-7.8/configure --target="$TARGET" --prefix="$PREFIX" --program-prefix=i386-elf-
make
make install
```

检查Makefile文件`make debug`。这个目标使用构建`kernel.elf`，这是一个对象文件(不是二进制)，带有很多在内核中生成的链接符号，感谢gcc的`-g`标签。使用`xxd`测试它，你可以看到一些字符串。实际上，正确测试对象文件中的字符串方法是`strings kernel.elf`

我们可以利用qemu非常酷的功能。输入`make debug`，并在gdb的shell上：

- 在`kernel.c:main()`设置一个断点: `b main`
- 运行操作系统: `continue`
- 运行两步代码: `next` 然后 `next`。你可以看到'X'在屏幕上，但是检查qemu屏幕，还没有打印
- 让我们看看再video内存里有什么：`print *video_memory`。有一个来自'Landed in 32-bit Protected Mode'的'L'
- 让我们确保`video_memory`指向正确的地址：print video_memory
- `next` 打印我们的'X'
- 确保：`print *video_memory` 看一下qemu的屏幕。

现在阅读`gdb`的教程是一个好机会，学习更高级的用法，比如 `info registers` , 这个在以后的工作中能节省大量的时间

你可能注意到了，自从这个教程开始，我们还没讨论我们将写个什么样的内核。它饿能是一个单内核，因为它更容易设计和实现，毕竟这是我们第一个操作系统。在未来，可能我们会加一个课程"15-b"，用于微内核设计，谁知道呢。

## 译者注

用7.8的版本编译会报错，无法解决，故采用的7.6的版本

```bash
cd /tmp/src
curl -O http://ftp.gnu.org/gnu/gdb/gdb-7.6.tar.gz
tar xf gdb-7.6.tar.gz
mkdir gdb-build
cd gdb-build
export PREFIX="/usr/local/i386elfgcc"
export TARGET=i386-elf
../gdb-7.8/configure --target="$TARGET" --prefix="$PREFIX" --program-prefix=i386-elf- --disable-werror
make
make install
```

如果报`no termcap library found` 错误，执行`sudo apt-get install libncurses5-dev` 即可。提示没有权限，前面添加`sudo`