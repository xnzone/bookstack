---
author: xnzone 
title: 数据链路层 
date: 2023-02-27 10:04:00
image: /covers/network.jpg
cover: false
math: true
weight: 10
tags: ["王道408", "计算机网络","数据链路层"]
---

## 功能概述

结点：主机、路由

链路：网络中两个结点之间的物理通道

数据链路：网络中两个结点之间的逻辑通道

帧：链路层的协议数据单元，封装网络层数据报

### 功能

负责通过一条链路从一个结点向另一个物理链路直接相连的相邻结点传送数据报

为网络层提供服务，无确认无连接服务，有确认无连接服务，有确认面向连接服务（有连接一定有确认）

链路管理，即连接的建立、维持、释放（用于面向连接的服务）

组帧

流量控制

差错控制（帧错/位错）

## 封装成帧和透明传输

### 封装成帧

封装成帧就是在一段数据的前后部分添加首部和尾部，这样就形成了一个帧

接收端在收到物理层上交的比特流后，就能根据首部和尾部的标记，从收到的比特流中识别帧的开始和结束

首部和尾部包含许多控制信息，他们的一个重要作用：帧定界（确定帧的界限）

帧同步：接收方应当能从接收到的二进制比特流中区分出帧的起始和终止

组帧的四种方式：1.字符计数法，2.字符(节)填充法，3.零比特填充法，4.违规编码法

### 透明传输

透明传输是指不管所传数据是什么养的比特组合，都应当能够在链路上传送

因此，链路层就看不见有什么妨碍数据传输的东西

当所传数据中的比特组合恰巧与某一个控制信息完全一样时，就必须采取适当的措施，使接收方不会将这样的数据误认为使某种控制信息。这样才能保证数据链路层的传输时透明的

### 字符计数法

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-byte-cnt.jpg)

### 字符填充法

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-byte-fill-1.jpg)

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-byte-fill-2.jpg)

ESC就是一个转义字符

### 零比特填充

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-zero-fill.jpg)

### 违规编码法

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-mess-code.jpg)

## 差错控制(检错编码)

### 差错来源

传输中的差错都是由于噪声引起的，分为全局性和局部性噪音

全局性：由于线路本身电器特性所产生的随机噪声(热噪声)，是信道固有的，随机存在的

解决办法：提高信噪比来减少或避免干扰(对传感器下手)

局部性：外界特定的短暂原因所造成的冲击噪声，是产生差错的主要原因

解决办法：通常利用编码技术来解决

### 差错的分类

- 位错：比特位出错，1变0，0变1
- 帧错：丢失，重复，失序

### 位错的差错控制

**检错编码**

- 奇偶校验码：奇校验-1的个数为奇数，偶校验-1的个数位偶数
- CRC循环冗余码：FCS的生成以及接收端CRC校验都是由硬件实现，处理很迅速，因此不会延误数据传输

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-crc-1.jpg)

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-link-crc-2.jpg)

**纠错编码**

- 海明校验码
