---
author: 黄健宏
title: DEBUG SEGFAULT
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21807
tags:
  - Redis
  - 调试
  - DEBUG SEGFAULT
---

# DEBUG SEGFAULT

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

执行一个不合法的内存访问从而让 Redis 崩溃，仅在开发时用于 BUG 模拟。

## 返回值

无

## 代码示例

```bash
redis> DEBUG SEGFAULT
Could not connect to Redis at: Connection refused

not connected>
```