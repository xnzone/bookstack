---
author: xnzone 
title: xv6(Unix Like OS) Lab
date: 2021-09-10 10:23:32
image: /covers/xv6.png
cover: true
weight: 1
tags: ["xv6", "os"]
---

>xv6课程网址[6.828-2018](https://pdos.csail.mit.edu/6.828/2018/)

## Introduction
xv6 is a teaching operating system developed in the summer of 2006, which we ported xv6 to RISC-V for a new undergraduate class 6.1810.

## xv6 sources and text

The latest xv6 source and text are available via
{{< highlight bash >}}
git clone git://github.com/mit-pdos/xv6-riscv.git
{{< /highlight  >}}

and

{{< highlight bash >}}
git clone git://github.com/mit-pdos/xv6-riscv-book.git
{{< /highlight  >}}

## Develop
{{< highlight bash >}}
git clone https://pdos.csail.mit.edu/6.828/2018/jos.git
{{< /highlight  >}}

{{< highlight bash >}}
docker build -t xv6 . 
bash start.sh
{{< /highlight  >}}