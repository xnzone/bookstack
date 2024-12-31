---
author: 黄健宏
title: PFCOUNT
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20602
tags:
  - Redis
  - HyperLogLog
  - PFCOUNT
---

# PFCOUNT key [key …]

> 可用版本： >= 2.8.9
> 
> 时间复杂度： 当命令作用于单个 HyperLogLog 时， 复杂度为 O(1) ， 并且具有非常低的平均常数时间。 当命令作用于 N 个 HyperLogLog 时， 复杂度为 O(N) ， 常数时间也比处理单个 HyperLogLog 时要大得多。

当 [PFCOUNT key [key …]](#pfcount) 命令作用于单个键时， 返回储存在给定键的 HyperLogLog 的近似基数， 如果键不存在， 那么返回 `0` 。

当 [PFCOUNT key [key …]](#pfcount) 命令作用于多个键时， 返回所有给定 HyperLogLog 的并集的近似基数， 这个近似基数是通过将所有给定 HyperLogLog 合并至一个临时 HyperLogLog 来计算得出的。

通过 HyperLogLog 数据结构， 用户可以使用少量固定大小的内存， 来储存集合中的唯一元素 （每个 HyperLogLog 只需使用 12k 字节内存，以及几个字节的内存来储存键本身）。

命令返回的可见集合（observed set）基数并不是精确值， 而是一个带有 0.81% 标准错误（standard error）的近似值。

举个例子， 为了记录一天会执行多少次各不相同的搜索查询， 一个程序可以在每次执行搜索查询时调用一次 [PFADD key element [element …]](../../06-hyperloglog/01-PFADD) ， 并通过调用 [PFCOUNT key [key …]](#pfcount) 命令来获取这个记录的近似结果。

## 返回值

整数回复： 给定 HyperLogLog 包含的唯一元素的近似数量。

## 代码示例

```bash
redis> PFADD  databases  "Redis"  "MongoDB"  "MySQL"
(integer) 1

redis> PFCOUNT  databases
(integer) 3

redis> PFADD  databases  "Redis"    # Redis 已经存在，不必对估计数量进行更新
(integer) 0

redis> PFCOUNT  databases    # 元素估计数量没有变化
(integer) 3

redis> PFADD  databases  "PostgreSQL"    # 添加一个不存在的元素
(integer) 1

redis> PFCOUNT  databases    # 估计数量增一
4
```
