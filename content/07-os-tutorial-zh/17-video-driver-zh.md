---
author: xnzone 
title: 显示器驱动
date: 2021-09-10 10:23:32
image: https://s2.loli.net/2025/09/24/J8Y6zBqZmSn9NkW.jpg
cover: false
weight: 717
tags: ["tutorial", "os"]
---

你需要通过谷歌预先了解的概念：VGA character cells, screen offset

**目标：在屏幕上书写字符串**

最终，我们将能够在屏幕上输出文本。这节课包含一些比之前更多一点代码，所以让我们一步一步来

打开`drivers/screen.h`，你讲看到我们已经为VGA卡驱动定义了一些常量和三个公开函数，一个清屏，另外一对是写字符串，为"kernel print"准备的名为`kprint`函数

现在打开`drivers/screen.c`。以私有的用于`kprint`内核API的辅助函数声明为起始。

有两个之前就学过的`get`和`set_cursor_offset()`I/O端口程序，还有一个程序直接操作视频内存的函数，名为`print_char()`

最终，有三个小的辅助函数用于把行和列转换成偏移值

## kprint_at

`kprint_at` 可能被`col`和`row`的`-1`值调用，这表明在当前光标位置打印字符串。

最开始设置三个变量，col/row和偏移值。然后用`char*`和当前坐标值调用`print_char()`迭代。

注意`print_char`字节返回的是下个光标位置偏移值，我们在下次循环中重复使用它。

`kprint`是对`kprint_at`的基本封装

## print_char

像`kprint_at`一样，`print_char`允许行/列的值为`-1`。在这个例子中，使用`ports.c`程序，从硬件撤回光标位置。

`print_char`也处理回车。为此，我们把光标偏移值设置为下一行的第0列

记住VGA cells占用两个字节，一个是字节本身，另一个是属性

## kernel.c

我们新的内核最终能够输出字符串

正确的字符定位，多行，跨行都是测试通过的。最后，尝试在屏幕范围外书写，会发生什么呢？

在下一节，我们将学习如何滚动屏幕