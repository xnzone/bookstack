---
author: xnzone 
title: xv6环境
date: 1906-01-01
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1201
tags: ["xv6", "os"]
---


使用Docker
----------------
由于用macos编译太过于麻烦，所以就安装Docker来进行操作，参考[Docker](/01-xv6-lab/lab0-envs/02-docker)


编译运行
----------------
- `make qemu-nox` 无图形界面编译运行
- `make qemu-nox-gdb` 无图形界面调试运行
- `make run-name` 运行用户程序，例如`make run-hello`运行`user/hello.c`
- `make run-name-nox` `run-name-gdb``run-name-gdb-nox`等等
- `make V=1` 打印正在执行的每一条指令

如果提示需要加载安全路径，gdb在运行的时候，需要带上参数，命令如下
`gdb -iex 'add-auto-load-safe-path .'`


GDB使用
--------------
- `Ctrl-c` 关闭GDB
- `c`或者`continue` 断点执行
- `si` 或者 `stepi` 分步执行
- `b function`或者`b file:line` 给特定的行或者函数设置断点
- `b *addr` 在EIP地址设置断点
- `set print pretty` 对于数组和结构体启用格式化打印
- `info registers` 打印通用寄存器，如eip,eflags和段选择器
- `x/Nx addr` 显示16进制虚拟地址的前N个字，N默认是1，addr可以是任何表达式
- `x/Ni addr` 显示N装配说明，使用$eip,将会显示当前指针介绍
- `symbol-file file` 切换到标志文件
- `thread n` 切换线程，n从0开始
- `info threads` 列出所有线程信息，包括其中的状态和函数

QEMU使用
-------------
- `ctrl-a c`停止qemu
- `xp/Nx paddr` 显示十六进制物理地址的前N位，N默认是1
- `info registers` 显示所有内部寄存器状态，包括隐藏段信息的机器码，中断描述符表，和任务寄存器

比如下面的显示
```text
cs =0008 10000000 ffffffff 10cf9a00 DPL=0 CS32 [-R-]
```

1. **cs =0008** 代码选择器可见部分。我们使用段0x8.这也表明全局描述符表(0x8&4=0)和当前特权等级(CPL)为0x8&3=0
2. **10000000** 段基地址。线性地址=逻辑地址+段基地址
3. **ffffffff** 段限制。超过线性地址0xffffffff将会导致段无效错误
4. **10cf9a00** 原生标志段，在接下来的几个字段，QEMU将帮助我们解码
5. **DPL=0** 特权等级段。只有代码在特权等级0上运行的时候才能加载这个段 
6. **CS32** 32位代码段。其他值包括数据段`DS`(别和DS寄存器搞混)和局部描述符表(LDT)
7. **[-R-]** 表示这个段只读的

- `info mem` 显示虚拟地址和权限
- `info pg` 显示当前页表结构，输出和`info mem`类似，但是单独给出页入口和页表入口
- `make QEMUEXTRA='-d int' ...`打印所有中断日志