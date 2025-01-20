---
author: xnzone 
title: 内核系统交叉编译
date: 2021-09-10 10:23:32
image: /covers/os-tutorial-zh.jpg
cover: false
weight: 712
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：cross-compiler

**目标：搭建构造内核的开发环境**

如果你使用的是Mac，你需要按照这个流程做下去。否则，对于一些接下来的课程，你可能需要等待一些时间。不管怎样，一旦我们跳转在一个更高级的编程语言开发，比如C，[阅读为什么](/02-os-tutorial/04-kernel/11-kernel-crosscompiler/)，你需要一个交叉编译环境

我将根据[OSDev wiki](/02-os-tutorial/04-kernel/11-kernel-crosscompiler/)介绍

## 需要的包

首先，安装必须的包。在linux上，使用你的包分发。在Mac上，如果你没有在lesson 00的时候安装过的话，安装brew，并且用`brew install`安装下面的包

- gmp
- mpfr
- libmpc
- gcc

是的，我们需要`gcc`去构建我们的交叉编译，尤其是在Mac上，`gcc`已经被`clang`替代了

一旦安装了，找到你安装的gcc（记住，不是clang）的位置，并且添加到环境变量，例如

```bash
export CC=/usr/local/bin/gcc-4.9
export LD=/usr/local/bin/gcc-4.9
```

我们需要构建二进制包和一个gcc交叉编译器，我们需要把这些放到`/usr/local/i386elfgcc`。所以让我们把这些放到环境变量吧。可以根据你的喜好来修改他们

```bash
export PREFIX="/usr/local/i386elfgcc"
export TARGET=i386-elf
export PATH="$PREFIX/bin:$PATH"
```

## binutils(二进制工具包)

记住：注意从网上复制的文本格式问题。我建议你一行一行复制

```bash
mkdir /tmp/src
cd /tmp/src
curl -O http://ftp.gnu.org/gnu/binutils/binutils-2.24.tar.gz # 如果404，查找一个最近的版本
tar xf binutils-2.24.tar.gz
mkdir binutils-build
cd binutils-build
../binutils-2.24/configure --target=$TARGET --enable-interwork --enable-multilib --disable-nls --disable-werror --prefix=$PREFIX 2>&1 | tee configure.log
make all install 2>&1 | tee make.log
```

## gcc

```bash
cd /tmp/src
curl -O https://ftp.gnu.org/gnu/gcc/gcc-4.9.1/gcc-4.9.1.tar.bz2
tar xf gcc-4.9.1.tar.bz2
mkdir gcc-build
cd gcc-build
../gcc-4.9.1/configure --target=$TARGET --prefix="$PREFIX" --disable-nls --disable-libssp --enable-languages=c --without-headers
make all-gcc 
make all-target-libgcc 
make install-gcc 
make install-target-libgcc
```

你应该已经有了所有GNU二进制工具包和`/usr/local/i386elfgcc/bin`编译器，`i386-elf-`前缀是为了避免与系统编译器和二进制工具包冲突

你可能想要把`$PATH`加入到`.bashrc`。从现在开始，在这个教程里面，当使用交叉编译的gcc的时候，我们将明确的使用前缀

译者注：

我在编译gcc过程中，遇到的一些问题，现在针对问题给出一些解决方案：

- 解压压缩包会失败，解决方案是在windows下载好，解压后再上传到linux
- 编译gcc的时候会出现`configure: error: Building GCC requires GMP 4.2+, MPFR 2.4.0+ and MPC 0.8.0+.`。这个可以按照提示，在命令行添加几个命令就可以；或者也可以安装一下`libmpc-dev`也可以解决
- 编译的时候出现`make: execvp: <YOUR FILE NAME>: Permission denied …` 。需要赋权，执行命令`chmod 755 /path/to/yourfile`