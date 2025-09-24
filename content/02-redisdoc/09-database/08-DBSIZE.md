---
author: 黄健宏
title: DBSIZE
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20908
tags:
  - Redis
  - 数据库
  - DBSIZE
---

# DBSIZE

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

返回当前数据库的 key 的数量。

## 返回值

当前数据库的 key 的数量。

## 代码示例

```bash
redis> DBSIZE
(integer) 5

redis> SET new_key "hello_moto"     # 增加一个 key 试试
OK

redis> DBSIZE
(integer) 6
```