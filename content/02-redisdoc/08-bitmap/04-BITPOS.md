---
author: 黄健宏
title: BITPOS
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20804
tags:
  - Redis
  - 位图
  - BITPOS
---

# BITPOS key bit [start] [end]

> 可用版本： >= 2.8.7
> 
> 时间复杂度： O(N)，其中 N 为位图包含的二进制位数量

返回位图中第一个值为 `bit` 的二进制位的位置。

在默认情况下， 命令将检测整个位图， 但用户也可以通过可选的 `start` 参数和 `end` 参数指定要检测的范围。

## 返回值

整数回复。

## 代码示例

```bash
127.0.0.1:6379> SETBIT bits 3 1    # 1000
(integer) 0

127.0.0.1:6379> BITPOS bits 0
(integer) 0

127.0.0.1:6379> BITPOS bits 1
(integer) 3
```
