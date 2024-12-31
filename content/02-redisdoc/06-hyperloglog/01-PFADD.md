---
author: 黄健宏
title: PFADD
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20601
tags:
  - Redis
  - HyperLogLog
  - PFADD
---

# PFADD key element [element …]

> 可用版本： >= 2.8.9
> 
> 时间复杂度： 每添加一个元素的复杂度为 O(1) 。

将任意数量的元素添加到指定的 HyperLogLog 里面。

作为这个命令的副作用， HyperLogLog 内部可能会被更新， 以便反映一个不同的唯一元素估计数量（也即是集合的基数）。

如果 HyperLogLog 估计的近似基数（approximated cardinality）在命令执行之后出现了变化， 那么命令返回 `1` ， 否则返回 `0` 。 如果命令执行时给定的键不存在， 那么程序将先创建一个空的 HyperLogLog 结构， 然后再执行命令。

调用 [PFADD key element [element …]](#pfadd) 命令时可以只给定键名而不给定元素：

- 如果给定键已经是一个 HyperLogLog ， 那么这种调用不会产生任何效果；
    
- 但如果给定的键不存在， 那么命令会创建一个空的 HyperLogLog ， 并向客户端返回 `1` 。
    

要了解更多关于 HyperLogLog 数据结构的介绍知识， 请查阅 [PFCOUNT key [key …]](../../06-hyperloglog/02-PFCOUNT) 命令的文档。

## 返回值

整数回复： 如果 HyperLogLog 的内部储存被修改了， 那么返回 1 ， 否则返回 0 。

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
