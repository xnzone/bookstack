---
author: 黄健宏
title: EVALSHA
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21202
tags:
  - Redis
  - Lua脚本
  - EVALSHA
---

# EVALSHA sha1 numkeys key [key …] arg [arg …]

> 可用版本： >= 2.6.0
> 
> 时间复杂度： 根据脚本的复杂度而定。

根据给定的 sha1 校验码，对缓存在服务器中的脚本进行求值。

将脚本缓存到服务器的操作可以通过 [SCRIPT LOAD script](script_load.html#script-load) 命令进行。

这个命令的其他地方，比如参数的传入方式，都和 [EVAL script numkeys key [key …] arg [arg …]](eval.html#eval) 命令一样。

```bash
redis> SCRIPT LOAD "return 'hello moto'"
"232fd51614574cf0867b83d384a5e898cfd24e5a"

redis> EVALSHA "232fd51614574cf0867b83d384a5e898cfd24e5a" 0
"hello moto"
```