---
author: 黄健宏
title: PFMERGE
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20603
tags:
  - Redis
  - HyperLogLog
  - PFMERGE
---

# PFMERGE destkey sourcekey [sourcekey …]

> 可用版本： >= 2.8.9
> 
> 时间复杂度： O(N) ， 其中 N 为被合并的 HyperLogLog 数量， 不过这个命令的常数复杂度比较高。

将多个 HyperLogLog 合并（merge）为一个 HyperLogLog ， 合并后的 HyperLogLog 的基数接近于所有输入 HyperLogLog 的可见集合（observed set）的并集。

合并得出的 HyperLogLog 会被储存在 `destkey` 键里面， 如果该键并不存在， 那么命令在执行之前， 会先为该键创建一个空的 HyperLogLog 。

## 返回值

字符串回复：返回 `OK` 。

## 代码示例

```bash
redis> PFADD  nosql  "Redis"  "MongoDB"  "Memcached"
(integer) 1

redis> PFADD  RDBMS  "MySQL" "MSSQL" "PostgreSQL"
(integer) 1

redis> PFMERGE  databases  nosql  RDBMS
OK

redis> PFCOUNT  databases
(integer) 6
```
