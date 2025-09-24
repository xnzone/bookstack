---
author: xnzone
title: xv6(Unix Like OS) Lab
date: 2021-09-10 10:23:32
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: true
weight: 12
tags:
  - xv6
  - os
---

>xv6课程网址[6.828-2018](https://pdos.csail.mit.edu/6.828/2018/)

## Introduction
xv6 is a teaching operating system developed in the summer of 2006, which we ported xv6 to RISC-V for a new undergraduate class 6.1810.

## xv6 sources and text

The latest xv6 source and text are available via
```bash
git clone git://github.com/mit-pdos/xv6-riscv.git
```

and

```bash
git clone git://github.com/mit-pdos/xv6-riscv-book.git
```

## Develop
```bash
git clone https://pdos.csail.mit.edu/6.828/2018/jos.git
```

```bash
docker build -t xv6 . 
bash start.sh
```