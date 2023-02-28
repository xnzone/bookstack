---
author: xnzone 
title: 网络体系结构
date: 2023-02-27 10:04:00
image: /covers/network.jpg
cover: false
weight: 8
tags: ["王道408", "计算机网络", "网络体系结构"]
---

## 概念与功能

- 网络：网站系统
- 计算机网络：将一个分散的，具有独立功能的计算机系统，通过通信设备和线路连接起来，由功能完善的软件实现资源共享和信息传递的系统
- 计算机网络的功能：数据通信，资源共享，分布式处理，提高可靠性，负载均衡，...
- 发展阶段-第一阶段：ARPAnet(美国国防部的分散指挥系统) -> Internet(TCP/IP协议)
- 发展阶段-第二阶段：1985后，三层结构的因特网(校园网，地区网，主干网)
- 发展阶段-第三阶段：多层次的ISP结构(本地ISP，地区ISP，主干ISP，电信，移动，联通都是ISP)
- 组成：硬件，软件，协议(一系列规则和约定的集合)

### 工作方式

- 边缘部分（C/S方式:客户和服务；P2P方式：对等连接peer-to-peer)
- 核心部分（为边缘部分服务）

### 功能组成

- 通信子网（网络层，数据链路层，物理层）：各种传输介质，通信设备，对应的网络协议组成
- 资源子网（应用层，表示层，会话层）：实现资源共享/数据处理的设备和软件的集合
- 传输层（承上启下）

### 分类

**按分布分类**

- 广域网WAN(交换技术)
- 城域网MAN
- 局域网LAN(广播技术)
- 个人区域网PAN

**使用者分**

- 公用网
- 专用网

**交换技术分**

- 电路交换
- 报文交换
- 分组交换

**拓扑结构分**

- 总线型
- 星形
- 环形
- 网状型(常用于广域网)

**传输技术分**

- 广播式网络：共享公共信号通道
- 点对点网络：使用分组存储转发和路由选择机制

## 标准化工作及相关组织

### 标准分类

- 法定标准：由权威机构制订的正式的，合法的标准，如OSI模型
- 事实标准：某些公司的产品在竞争中占据了主流产生的标准，如TCP/IP协议

### 标准化工作

RFC要上升为Internet正式标准的四个阶段

- 因特网草案
- 建议标准
- 草案标准（现在已经取消了）
- 因特网标准

### 相关组织

- 国际标准化组织ISO: OSI模型，HDLC协议
- 国际电信联盟ITU：制定通信规则
- 电气和电子工程师协会IEEE：学术机构，IEEE802系列标准，5G
- Internet工程任务组IETF：负责因特网相关标准的制订 RFCYYYY

## 性能指标

- 速率：数据率，数据传输率，比特率。连接在计算机网络的主机在数字信道上传输数据位数的速率，单位有b/s,kb/s,Mb/s等，都是100倍关系
- 存储容量：1Byte=8bit;1kb=2^10B=1024B=1024*8b(一般用比特，不是字节)
- 带宽：表示网络的通信线路传送数据的能力。单位是b/s，链路带宽=1MB/s，意味着主机在1us内可向链路发送1bit数据
- 吞吐量：单位时间内，通过某个网络或信道或接口的数据量，单位b/s等。受到网络的带宽和网络的额定速率的限制，多个服务器的最高速率相加，最高不可能超过带宽
- 时延：从网络的一端传送到另一端所需的时间，单位是s，分为发送时延(传输时延),传播时延，排队时延，处理时延
- 发送时延：数据长度/信道带宽
- 传播时延：信道长度/电磁波在信道上的传播速率
- 排队时延：等待输入(或输出)链路的可用时间
- 处理时延：检错，找出口
- 时延带宽积：传播时延*带宽，单位b。又被称为以比特为单位的链路长度
- 往返时间RTT：从发送方发送数据开始，到发送方收到接收方的确认总共经历的时延
- RTT越大，在收到确认之前，可发送的数据越多
- RTT包括：往返传播时延，末端处理时间
- 利用率：分为信道利用率和网络利用率
- 信道利用率：有数据通过的时间/总时间
- 网络利用率：信道利用率加权平均值

## 分层结构、协议、接口、服务

### 为什么要分层

发送文件之前要完成的动作

- 发起通信的计算机必须将数据通信的通路进行激活
- 要告诉网络如何识别目的主机
- 发起通信的计算机要查明目的主机是否开机，并且与网络链接正常
- 发起通信的计算机要弄清楚，对方计算机中文件管理程序是否已经准备好工作
- 确保差错和意外可以解决
- ...

因此需要分层

### 分层的原则

- 每层相对独立
- 每层界面自然清晰，易于理解
- 每层都采用最合适的技术来实现
- 保持下层和上层的独立性
- 分层结构应该促进标准化工作

**实体**：第n层的活动元素被称为n层实体，同一层的实体叫对等实体

**协议**：为进行网络中对等实体数据交换而建立的规则、标准或约定

- 语法：规定传输数据的格式
- 语义：规定所要完成的功能
- 同步：规定各种操作的顺序

**接口**：上层使用下层服务的入口

**服务**：下层为相邻层提供的功能调用

**SDU数据服务单元**：为完成用户所要求的功能而应传送的数据（上一层真正有意义的信息）

**PCI协议控制信息**：控制协议操作的信息（控制信息）

**PDU协议数据单元**：对等层次之间传送的数据单元（上面两种信息之和）

PCI+SDU=PDU：紧接着PDU作为下一层的SDU，是一个类似于金字塔的结构

## OSI参考模型

OSI是7层模型，是法定标准，但现在基本上都是用TCP/IP协议，理论是成功的，但市场是失败的

- 应用层
- 表示层
- 会话层
- 传输层
- 网络层
- 数据链路层
- 物理层

### 通信过程

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-osi-transport.jpg)

### 应用层

面向用户的所有能和用户交互产生网络流量的程序就是应用层

典型的服务：

- 文件传输(FTP)
- 电子邮件(SMTP)
- 万维网(HTTP)
- ...

### 表示层

用于处理在两个通信系统中交换信息的表示方式（语法和语义）

功能：

- 数据格式变换
- 数据加密/解密
- 数据压缩和恢复

协议：JPEG，ASCII

### 会话层

向表示层/用户进程提供建立连接，并在连接上有序的传输数据。这就是会话，也被称为建立同步(SYN)

功能：

- 建立、管理、终止会话
- 通过校验点可使会话在通信失效时，从校验点/同步点继续恢复通信，实现数据同步（传输大文件）

协议：ADSP，ASP

### 传输层

负责主机中两个进程的通信，即端到端的通信。传输单位是报文段或用户数据报

功能：

- 可靠传输、不可靠传输
- 差错控制
- 流量控制
- 复用分用

复用：多个应用进程可同时使用下面运输车的服务

分用：运输层把收到的信息分别交付给上面应用层相应的进程

协议：UDP，TCP

### 网络层

主要任务是把分组从源端传送目的端，为分组交换网上的不同主机提供通信服务。网络层传输单位是数据报。

功能：

- 路由选择：选择合适的路由（最佳路径）
- 流量控制：协调发送端和接收端的速度
- 差错控制
- 拥塞控制：通过一定措施缓解所有节点都来不及接收分组的状态

协议：IP，IPX，ICMP，IGMP，ARP，RARP，OSPF

### 数据链路层

主要任务是把网络层传下来的数据报组装成帧。数据链路层/链路层的传输单位是帧。

功能：

- 成帧（定义帧的开始和结束）
- 差错控制 帧错+位错
- 流量控制
- 访问（接入）控制 控制对信道的访问

协议：SDLC HDLC PPP STP

### 物理层

主要任务是在物理媒体上实现比特流的透明传输。传输单位是比特。

透明传输：指的是不敢是什么样的比特组合，都应当能够在链路上传送。

功能：

- 定义接口特性
- 定义传输模式 单工、半双工、双工
- 定义传输速率
- 比特同步
- 比特编码

协议：Rj45 802.3

## TCP/IP参考模型

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-tcp-ip.jpg)

### 与OSI参考模型差异

**相同点**

- 都分层
- 基于独立协议栈的概念
- 可以实现异构网络互联

**不同点**

- OSI定义的三点
- OSI先出现，参考模型先于协议发明，不偏向协议
- TCP/IP设计之初就考虑到异构互联问题，将IP作为重要层次

面向连接分为三个阶段，1. 建立连接；2. 建立连接后传输数据；3.数据传输后关闭连接

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-osi-tcpip-diff.jpg)

### 五层参考模型

- 综合了OSI和TCP/IP的优点
- 应用层支持网络应用
- 传输层负责进程与进程的数据传输
- 网络层负责源主机到目的主机的数据分组与路由转发
- 数据链路层负责把网络层传下来的数据报封装成帧
- 物理层负责比特传输

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/network-five.jpg)