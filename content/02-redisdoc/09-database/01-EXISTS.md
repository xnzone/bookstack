---
author: 黄健宏
title: EXISTS
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20901
tags:
  - Redis
  - 数据库
  - EXISTS
---

# EXISTS key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

检查给定 `key` 是否存在。

## 返回值

若 `key` 存在，返回 `1` ，否则返回 `0` 。

## 代码示例

```bash
redis> SET db "redis"
OK

redis> EXISTS db
(integer) 1

redis> DEL db
(integer) 1

redis> EXISTS db
(integer) 0
```