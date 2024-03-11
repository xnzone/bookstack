---
author: 黄健宏
title: INCRBYFLOAT
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 313
tags: ["Redis", "字符串", "INCRBYFLOAT"]
---

# INCRBYFLOAT key increment

> 可用版本： >= 2.6.0
> 
> 时间复杂度： O(1)

为键 `key` 储存的值加上浮点数增量 `increment` 。

如果键 `key` 不存在， 那么 `INCRBYFLOAT` 会先将键 `key` 的值设为 `0` ， 然后再执行加法操作。

如果命令执行成功， 那么键 `key` 的值会被更新为执行加法计算之后的新值， 并且新值会以字符串的形式返回给调用者。

无论是键 `key` 的值还是增量 `increment` ， 都可以使用像 `2.0e7` 、 `3e5` 、 `90e-2` 那样的指数符号(exponential notation)来表示， 但是， **执行 INCRBYFLOAT 命令之后的值**总是以同样的形式储存， 也即是， 它们总是由一个数字， 一个（可选的）小数点和一个任意长度的小数部分组成（比如 `3.14` 、 `69.768` ，诸如此类)， 小数部分尾随的 `0` 会被移除， 如果可能的话， 命令还会将浮点数转换为整数（比如 `3.0` 会被保存成 `3` ）。

此外， 无论加法计算所得的浮点数的实际精度有多长， `INCRBYFLOAT` 命令的计算结果最多只保留小数点的后十七位。

当以下任意一个条件发生时， 命令返回一个错误：

- 键 `key` 的值不是字符串类型(因为 Redis 中的数字和浮点数都以字符串的形式保存，所以它们都属于字符串类型）；
    
- 键 `key` 当前的值或者给定的增量 `increment` 不能被解释(parse)为双精度浮点数。
    

## 返回值

在加上增量 `increment` 之后， 键 `key` 的值。

## 代码示例

{{< highlight shell >}}
redis> GET decimal
"3.0"

redis> INCRBYFLOAT decimal 2.56
"5.56"

redis> GET decimal
"5.56"
{{< /highlight >}}