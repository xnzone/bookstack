---
author: 黄健宏
title: LASTSAVE
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21304
tags:
  - Redis
  - 持久化
  - LASTSAVE
---

# LASTSAVE

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

返回最近一次 Redis 成功将数据保存到磁盘上的时间，以 UNIX 时间戳格式表示。

## 返回值

一个 UNIX 时间戳。

## 代码示例

```bash
redis> LASTSAVE
(integer) 1324043588
```