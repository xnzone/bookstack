---
author: 黄健宏
title: SCRIPT FLUSH
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21205
tags:
  - Redis
  - Lua脚本
  - SCRIPT FLUSH
---

# SCRIPT FLUSH

> 可用版本： >= 2.6.0
> 
> 复杂度： O(N) ， `N` 为缓存中脚本的数量。

清除所有 Lua 脚本缓存。

关于使用 Redis 对 Lua 脚本进行求值的更多信息，请参见 [EVAL script numkeys key [key …] arg [arg …]](eval.html#eval) 命令。

## 返回值

总是返回 `OK`

## 代码示例

```bash
redis> SCRIPT FLUSH
OK
```