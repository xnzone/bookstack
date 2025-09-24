---
author: 黄健宏
title: UNWATCH
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21105
tags:
  - Redis
  - 自动过期
  - UNWATCH
---

# UNWATCH

**UNWATCH**

取消 [WATCH](../../11-transation/04-WATCH) 命令对所有 key 的监视。

如果在执行 [WATCH](../../11-transation/04-WATCH) 命令之后， [EXEC](../../11-transation/02-EXEC) 命令或 [DISCARD](../../11-transation/03-DISCARD) 命令先被执行了的话，那么就不需要再执行 [UNWATCH](#unwatch) 了。

因为 [EXEC](../../11-transation/02-EXECc) 命令会执行事务，因此 [WATCH](../../11-transation/04-WATCH) 命令的效果已经产生了；而 [DISCARD](../../11-transation/03-DISCARD) 命令在取消事务的同时也会取消所有对 key 的监视，因此这两个命令执行之后，就没有必要执行 [UNWATCH](#unwatch) 了。

**可用版本：**

>= 2.2.0

**时间复杂度：**

O(1)

**返回值：**

总是 `OK` 。

```bash
redis> WATCH lock lock_times
OK

redis> UNWATCH
OK
```