---
author: 黄健宏
title: GEOADD
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20701
tags:
  - Redis
  - 地理位置
  - GEOADD
---

# GEOADD key longitude latitude member [longitude latitude member …]

> 可用版本： >= 3.2.0
> 
> 时间复杂度： 每添加一个元素的复杂度为 O(log(N)) ， 其中 N 为键里面包含的位置元素数量。

将给定的空间元素（纬度、经度、名字）添加到指定的键里面。 这些数据会以有序集合的形式被储存在键里面， 从而使得像 `GEORADIUS` 和 `GEORADIUSBYMEMBER` 这样的命令可以在之后通过位置查询取得这些元素。

`GEOADD` 命令以标准的 `x,y` 格式接受参数， 所以用户必须先输入经度， 然后再输入纬度。 `GEOADD` 能够记录的坐标是有限的： 非常接近两极的区域是无法被索引的。 精确的坐标限制由 EPSG:900913 / EPSG:3785 / OSGEO:41001 等坐标系统定义， 具体如下：

- 有效的经度介于 -180 度至 180 度之间。
    
- 有效的纬度介于 -85.05112878 度至 85.05112878 度之间。
    

当用户尝试输入一个超出范围的经度或者纬度时， `GEOADD` 命令将返回一个错误。

## 返回值

新添加到键里面的空间元素数量， 不包括那些已经存在但是被更新的元素。

## 代码示例

```bash
redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2

redis> GEODIST Sicily Palermo Catania
"166274.15156960039"

redis> GEORADIUS Sicily 15 37 100 km
1) "Catania"

redis> GEORADIUS Sicily 15 37 200 km
1) "Palermo"
2) "Catania"
```
