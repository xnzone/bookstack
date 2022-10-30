---
author: xnzone 
title: 苹果电脑编译
date: 1906-01-01
image: /covers/os-tutorial.png
cover: false 
weight: 4
tags: ["tutorial", "os"]
---


**目标：升级构建系统到EI Capitan**

如果你从最开始就按照这个指导做的，并且在升级了EI Capitan之后，发现Makefile不工作了，那么可以按照这篇文章，来升级的交叉编译环境

否则，直接跳到下一节

## 升级交叉编译环境

或多或少跟11节介绍的一样

首先，运行`brew upgrade` 把你的gcc升级到5.0（这边指导写的时候）

然后运行`xcode-select --install`来升级OSX命令行工具

一旦安装成功，找到你gcc安装包所在位置(记住不是clang)，并且添加到环境变量。例如

{{< highlight bash >}}
export CC=/usr/local/bin/gcc-5
export LD=/usr/local/bin/gcc-5
{{< /highlight  >}}

我们需要再次编译二进制构建工具和我们的交叉编译器gcc，输出targets和prefix环境变量

{{< highlight bash >}}
export PREFIX="/usr/local/i386elfgcc"
export TARGET=i386-elf
export PATH="$PREFIX/bin:$PATH"
{{< /highlight  >}}

## binutils

记住复制的时候格式问题，建议一行一行复制

{{< highlight bash >}}
mkdir /tmp/src
cd /tmp/src
curl -O http://ftp.gnu.org/gnu/binutils/binutils-2.24.tar.gz # If the link 404's, look for a more recent version
tar xf binutils-2.24.tar.gz
mkdir binutils-build
cd binutils-build
../binutils-2.24/configure --target=$TARGET --enable-interwork --enable-multilib --disable-nls --disable-werror --prefix=$PREFIX 2>&1 | tee configure.log
make all install 2>&1 | tee make.log
{{< /highlight  >}}

## gcc

{{< highlight bash >}}
cd /tmp/src
curl -O http://mirror.bbln.org/gcc/releases/gcc-4.9.1/gcc-4.9.1.tar.bz2
tar xf gcc-4.9.1.tar.bz2
mkdir gcc-build
cd gcc-build
../gcc-4.9.1/configure --target=$TARGET --prefix="$PREFIX" --disable-nls --disable-libssp --enable-languages=c --without-headers
make all-gcc 
make all-target-libgcc 
make install-gcc 
make install-target-libgcc
{{< /highlight  >}}

现在在这节课的文件夹内输入`make`， 检查所有的编译是否顺利