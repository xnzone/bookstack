---
author: 黄健宏
title: TIME
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21605
tags:
  - Redis
  - 客户端与服务器
  - TIME
---

# TIME

> 可用版本： >= 2.6.0
> 
> 时间复杂度： O(1)

返回当前服务器时间。

## 返回值

一个包含两个字符串的列表： 第一个字符串是当前时间(以 UNIX 时间戳格式表示)，而第二个字符串是当前这一秒钟已经逝去的微秒数。

## 代码示例

```bash
redis> TIME
1) "1332395997"
2) "952581"

redis> TIME
1) "1332395997"
2) "953148"
```