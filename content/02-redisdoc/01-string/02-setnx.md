---
author: 黄健宏
title: SETNX
date: 2024-03-07 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2102 
tags: ["Redis", "字符串", "SETNX"]
---

# SETNX key value

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

只在键 `key` 不存在的情况下， 将键 `key` 的值设置为 `value` 。

若键 `key` 已经存在， 则 `SETNX` 命令不做任何动作。

`SETNX` 是『SET if Not eXists』(如果不存在，则 SET)的简写。

## 返回值

命令在设置成功时返回 `1` ， 设置失败时返回 `0` 。

## 代码示例

```shell
redis> EXISTS job                # job 不存在
(integer) 0

redis> SETNX job "programmer"    # job 设置成功
(integer) 1

redis> SETNX job "code-farmer"   # 尝试覆盖 job ，失败
(integer) 0

redis> GET job                   # 没有被覆盖
"programmer"
```