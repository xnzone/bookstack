---
author: 黄健宏
title: TYPE
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20902
tags:
  - Redis
  - 数据库
  - TYPE
---

# TYPE key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

返回 `key` 所储存的值的类型。

## 返回值

- `none` (key不存在)
    
- `string` (字符串)
    
- `list` (列表)
    
- `set` (集合)
    
- `zset` (有序集)
    
- `hash` (哈希表)
    
- `stream` （流）
    

## 代码示例

```bash
# 字符串

redis> SET weather "sunny"
OK

redis> TYPE weather
string

# 列表

redis> LPUSH book_list "programming in scala"
(integer) 1

redis> TYPE book_list
list

# 集合

redis> SADD pat "dog"
(integer) 1

redis> TYPE pat
set
```
