---
author: xnzone 
title: 线性表 
date: 2023-02-22 10:04:00
image: https://s2.loli.net/2025/09/24/QCDLM3PdlaS4jgv.png
cover: false
weight: 902 
tags: ["数据结构", "算法", "线性表"]
---
## 线性表定义和基本操作
### 定义

具有相同数据类型的n个数据元素的有限序列

如果用L命名线性表，则一般表示为 `L = (a1,a2,...,ai,...,an)`,其中`a1`是表头，`an`是表尾；除第一个元素外，其他元素都有一个直接的前驱，除最后一个元素外，其他元素都有一个直接的后驱

### 基本操作
- `InitList(&L)` 初始化，构造一个空线性表，分配内存空间
- `DestroyList(&L)` 销毁，销毁线性表，并释放线性表L所占用的内存空间
- `Insert(&L, i, &e)` 插入操作，在表L的第i个位置插入元素e
- `Delete(&L, i, &e)` 删除操作，删除表L中第i个位置的元素，并用e返回删除元素的值
- `Locate(L, e)` 按值查找，在表L中查找具有给定值的元素位置
- `Get(L, i)` 按位置查找，在表L中查找第i个位置元素的值
- `Length(L)` 表长度，返回表L的长度
- `Print(L)` 输出操作，按前后顺序输出线性表所有元素的值
- `Empty(L)` 判空操作，若表为空，则返回true，否则返回false


## 顺序表

### 定义

用顺序存储的方式实现线性表顺序存储，逻辑上相邻的元素，物理位置上也相邻

- 优点：可随机存储，存储密度高
- 缺点：要求大片连续空间，改变容量不方便

### 顺序表特点
- 随机访问，O(1)时间内找到第i个元素
- 存储密度高
- 扩展容量不方便
- 插入和删除操作不方便，需要移动大量元素

### 静态分配实现

![顺序表静态分配实现](https://s2.loli.net/2025/09/24/HEXDhxMil2J5cY6.png)
![顺序表静态分配实现](https://s2.loli.net/2025/09/24/kstmX39KTIzd5Sy.png)

**如果数组满了怎么办**：

顺序表长度刚开始确定后就无法修改（存储空间是静态的），同时如果提前初始化太多，又会造成空间浪费。所以需要动态分配

**动态分配操作**

- C：`malloc`和`free`函数
- C++：`new`和`delete`关键字

tips：

- `L.data = (ElemType*)malloc(sizeof(ElemType*)*InitSize);`
- `malloc`函数返回一个指针，空间需要强制转为定义的数据元素类型指针
- `malloc`函数的参数，指明要分配多大的连续内存空间

### 动态分配实现

![顺序表动态分配](https://s2.loli.net/2025/09/24/Y4X3xuCIbWE6hck.png)

### 插入、删除和查找

- 插入和删除平均时间复杂度是O(n)
- 根据位置查找的平均时间复杂度是O(1)
- 根据元素查找的平均时间复杂度是O(n)

**插入**
![顺序表插入](https://s2.loli.net/2025/09/24/AOC9zHPFh6DYmSV.png)

增加i的合法性判断

![顺序表插入增加i的合法性判断](https://s2.loli.net/2025/09/24/sl7yELhvAguz6DW.png)

**删除**

![顺序表删除](https://s2.loli.net/2025/09/24/e2uOH8XwdgRvDZ5.png)

**查找**

按位置查找

![顺序表按位置查找](https://s2.loli.net/2025/09/24/IhYUBmCgryDzt9x.png)

按值查找

![顺序表按值查找](https://s2.loli.net/2025/09/24/oqUVHPlTBKyMsJE.png)

tips

- 结构类型的数据元素不能用==比较，但C++可以重载来实现


## 单链表

### 定义

每个节点除了存放数据元素外，还要存储指向下一个节点的指针

- 优点：不要求大量连续空间，改变容量方便
- 缺点：不可随机存取，要耗费空间存放指针

tips：

- typedef <数据类型> <别名>
- `typedef struct LNode LNode;`
- 之后可以用`LNode`代替`struct LNode`

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-base.jpg)

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-singer-list-create.jpg)

不带头节点的单链表

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-nohead.jpg)

带头节点的单链表

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-head.jpg)

区别：

- 不带头节点，写代码更麻烦
- 对第一个数据节点和后续节点的处理需要不同的代码逻辑
- 对空表和非空表的处理需要不同的代码逻辑
- 一般使用带头节点的单链表

### 插入、删除、查找和其他操作

- 按位置插入平均时间复杂度是O(n)
- 按位置删除平均时间复杂度是O(n)
- 按位置查找平均时间复杂度是O(n)
- 按节点插入平均时间复杂度是O(1)
- 按节点删除平均时间复杂度是O(1)
- 按节点查找平均时间复杂度是O(n)

**插入**

按位置插入-带头节点

- 在表L的第i个位置插入指定元素e
- 找到第i-1个节点，将新节点插入其后
- 带有头节点，插入更加方便

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-head-insert.jpg)

按位置插入-不带头节点

- 若`i != 1`，处理方法跟带头节点一样
- `int j = 1`而非带头节点的0

![](images/dsa/linear-list-single-list-nohead-insert.jpg)

指定节点的后插操作

- 按位置插入的代码一部分，按位置插入可以调用这个代码
- 封装代码，提高复用性

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-insert-next.jpg)


指定节点的前操作

- 方法1: 获取头节点，然后一步步找到指定节点的前驱
- 方法2: 将新节点连接到指定节点p的后继，接着指定节点p连接新节点s，将p中元素复制到s中，将p中元素覆盖为要插入的元素e

方法1
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-insert-prev-1.jpg)

方法2
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-insert-prev-2.jpg)


**删除**

按位置删除-带头节点

- 删除操作，删除表L中第i个位置的元素，并用e返回删除元素的值
- 找到第i-1个节点，将其指针指向第i+1个节点，并释放第i个节点


![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-delete-locate.jpg)

按指定节点删除

- 删除节点p，需要修改前驱节点的next指针，和指定节点的前插操作一样
- 方法1: 传入头指针，循环着p的前驱节点
- 方法2: 偷天换日，类似于节点的前插

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-delete-node.jpg)

tips：

- 如果要删除节点p是最后一个节点，只能从表头开始一次寻找前驱

**查找**

- 按位置查找操作，获取表L中第i个位置元素的值
- 实际上是单链表的插入中找到i-1部分就是按位置查找

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-locate.jpg)

- 按元素查找，跟按位置查找类似

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-node.jpg)


**长度**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-length.jpg)


**创建**

尾插法

- 每次插入元素都插入到链表尾部
- 方法1：每次都从头开始往后便利，时间复杂度为O(n^2)
- 方法2: 增加一个尾指针r，每次插入都让r指向新的表尾节点，时间复杂度为O(n)

方法1
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-tail-create-1.jpg)

方法2
![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-tail-create-2.jpg)


头插法

- 每次插入元素到表头
- 头插法和单链表后插操作是一样的
- `L->next = NULL`可以防止野指针

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-single-list-head-create.jpg)

tips:

- 头插法和尾插法：核心就是初始化操作，指定节点的后插操作
- 个指向表尾节点的指针
- 头插法重要应用：链表的逆置

## 双链表

- 单链表：无法逆向检索
- 双链表：可进可退，但存储密度更低一点

**定义**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-double-list-def.jpg)

**初始化-带头节点**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-double-list-init.jpg)

**插入**

- 如果p节点为最后一个节点，产生空指针的问题
- 注意指针的修改顺序

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-double-list-insert.jpg)

**删除**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-double-list-delete.jpg)


**遍历**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-double-list-range.jpg)


## 循环链表



### 循环单链表

- 表尾节点的next指针指向头节点
- 从一个节点出发可以找到其他任何节点

**初始化**

- 从头节点找到尾部，时间复杂度为O(n)
- 如果需要频繁访问表头、表尾，可以让L指向表尾元素(插入、删除时可能需要修改L)
- 从尾部找到头部，时间复杂度尾O(1)



![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-cycle-single-list-init.jpg)

### 循环双链表

- 表头节点的prior指向表尾节点
- 表尾节点next指向头节点

**初始化**

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-cycle-double-list-init.jpg)



### 插入和删除

**插入**

- 如果p节点尾最后一个节点，因为next节点尾null，p->next->prior=s 会产生空指针的问题
- 循环链表规避是因为最后节点的next节点为头节点，因此不会发生此问题

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-cycle-list-insert.jpg)

**删除**

- 如何判空
- 如何判断节点p是否表尾/表头节点
- 如何在表头、表中、表尾插入/删除一个节点

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-cycle-list-delete.jpg)


## 静态链表

- 分配一整片连续内存空间，各个节点集中安置
- 每个节点由两部分组成：data(数据元素)和next(游标)
- 0号节点充当头节点，不具体存放数据
- 游标为-1表示已经到达表尾
- 游标充当指针，表示下个节点存放位置

### 定义

- 用数组的方式实现链表
- 优点：增删操作不需要移动大量元素
- 缺点：不能随机存储，只能从头开始着；容量固定不可变
- 场景：不支持指针的低级语言；元素数量固定不变的场景

方法1

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-static-list-create-1.jpg)

方法2

![](https://jihulab.com/xnzone/earth-bear/-/raw/master/linear-list-static-list-create-2.jpg)

### 基本操作

**初始化**

- 把a[0]的next设为-1
- 把其他节点的next设为一个特殊值来表示节点空闲，如-2

**插入**

- 找到一个空节点，存入数据
- 从头节点出发找到位序为i-1的节点
- 修改新节点的next
- 修改i-1号节点的next

**删除**

- 从表头出发找到前驱节点
- 修改前驱节点的游标
- 被删除节点的next设为-2
