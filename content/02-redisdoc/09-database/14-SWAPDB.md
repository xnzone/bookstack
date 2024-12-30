---
author: 黄健宏
title: SWAPDB
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 914
tags:
  - Redis
  - 数据库
  - SWAPDB
---

# SWAPDB db1 db2

> 版本要求： >= 4.0.0
> 
> 时间复杂度： O(1)

对换指定的两个数据库， 使得两个数据库的数据立即互换。

## 返回值

`OK`

## 代码示例

```bash
# 对换数据库 0 和数据库 1
redis> SWAPDB 0 1
OK
```