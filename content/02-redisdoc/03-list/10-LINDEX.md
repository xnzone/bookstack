---
author: 黄健宏
title: LINDEX
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 310
tags:
  - Redis
  - 列表
  - LINDEX
---

# LINDEX key index

> 可用版本： >= 1.0.0
> 
> 时间复杂度：O(N)， `N` 为到达下标 `index` 过程中经过的元素数量。因此，对列表的头元素和尾元素执行 [LINDEX](https://bookstack.xnzone.eu.org/02-redisdoc/03-list/10-lindex/) 命令，复杂度为O(1)。

返回列表 `key` 中，下标为 `index` 的元素。

下标(index)参数 `start` 和 `stop` 都以 `0` 为底，也就是说，以 `0` 表示列表的第一个元素，以 `1` 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 `-1` 表示列表的最后一个元素， `-2` 表示列表的倒数第二个元素，以此类推。

如果 `key` 不是列表类型，返回一个错误。

## 返回值

列表中下标为 `index` 的元素。 如果 `index` 参数的值不在列表的区间范围内(out of range)，返回 `nil` 。

## 代码示例

```shell
redis> LPUSH mylist "World"
(integer) 1

redis> LPUSH mylist "Hello"
(integer) 2

redis> LINDEX mylist 0
"Hello"

redis> LINDEX mylist -1
"World"

redis> LINDEX mylist 3        # index不在 mylist 的区间范围内
(nil)
```