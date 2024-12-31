---
author: 黄健宏
title: DECR
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20114
tags: ["Redis", "字符串", "DECR"]
---

# DECR key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

为键 `key` 储存的数字值减去一。

如果键 `key` 不存在， 那么键 `key` 的值会先被初始化为 `0` ， 然后再执行 `DECR` 操作。

如果键 `key` 储存的值不能被解释为数字， 那么 `DECR` 命令将返回一个错误。

本操作的值限制在 64 位(bit)有符号数字表示之内。

关于递增(increment) / 递减(decrement)操作的更多信息， 请参见 `INCR` 命令的文档。

## 返回值

`DECR` 命令会返回键 `key` 在执行减一操作之后的值。

## 代码示例

对储存数字值的键 `key` 执行 `DECR` 命令：

```shell
redis> SET failure_times 10
OK

redis> DECR failure_times
(integer) 9
```

对不存在的键执行 `DECR` 命令：

```shell
redis> EXISTS count
(integer) 0

redis> DECR count
(integer) -1
```