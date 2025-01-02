---
author: 黄健宏
title: SCRIPT LOAD
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21203
tags:
  - Redis
  - Lua脚本
  - SCRIPT LOAD
---

# SCRIPT LOAD script

> 可用版本： >= 2.6.0
> 
> 时间复杂度： O(N) , `N` 为脚本的长度(以字节为单位)。

将脚本 `script` 添加到脚本缓存中，但并不立即执行这个脚本。

[EVAL script numkeys key [key …] arg [arg …]](eval.html#eval) 命令也会将脚本添加到脚本缓存中，但是它会立即对输入的脚本进行求值。

如果给定的脚本已经在缓存里面了，那么不做动作。

在脚本被加入到缓存之后，通过 EVALSHA 命令，可以使用脚本的 SHA1 校验和来调用这个脚本。

脚本可以在缓存中保留无限长的时间，直到执行 [SCRIPT FLUSH](script_flush.html#script-flush) 为止。

关于使用 Redis 对 Lua 脚本进行求值的更多信息，请参见 [EVAL script numkeys key [key …] arg [arg …]](eval.html#eval) 命令。

## 返回值

给定 `script` 的 SHA1 校验和。

## 代码示例

```bash
redis> SCRIPT LOAD "return 'hello moto'"
"232fd51614574cf0867b83d384a5e898cfd24e5a"

redis> EVALSHA 232fd51614574cf0867b83d384a5e898cfd24e5a 0
"hello moto"
```