---
author: xnzone 
title: "进程互斥与同步"
date: 2022-09-10 10:23:32
image: /covers/os.png
cover: false
weight: 9 
tags: ["os", "操作系统"]
---

两个或多个进程读写一些共享数据(一般是驻内存或共享文件)，最终的结果取决于当时运行的是哪个程序，这种场景被称为**竞态条件**

为了避免竞态条件，最关键的处理就是在同一个时间内，只允许一个进程使用共享数据(共享内存，共享文件或者其他的贡献的东西)，通常来说就是**互斥**。根据互斥的解决方案，设计了一系列关于操作系统同步的方案

## 互斥
互斥方案在现行的解决方案中，主要分为软件和硬件两种方案。软件方案主要是**Peterson算法**，硬件方案主要是`TSL`指令

### Peterson算法
Peterson算法的主要代码结构如下

```c
#define FALSE 0
#define TRUE 1
#define N 2 /* 进程数量 */

int turn; /* 谁的轮次 */
int interested[N]; /* 所有变量初始化为0 */

void enter_region(int process) { /* 进程是0或者1 */
    int other;
    other = 1 - process; /* 其他进程的数量，跟进程是相反数 */
    interested[process] = TRUE; /* 表明正在运行 */
    turn = process; /* 标记运行的进程 */
    while(turn == process && interested[other] == TRUE); /* 循环 */
}

void leave_region(int process) {
    interested[process] = FALSE;
}
```

主要解决了两个进程之间互斥的问题，用一个`turn`变量来表明是谁的轮次。在使用共享变量（即进入其临界区）之前，各个进程使用其进程号0或1作为参数来调用enter_region。该调用在需要时将使进程等待，直到能安全地进入临界区。在完成对共享变量的操作之后，进程将调用leave_region，表示操作已完成，若其他的进程希望进入临界区，则现在就可以进入。

### TSL指令
`TSL`指令全称为`Test and Set Lock`，即测试并加锁，这是一个原子操作

`TSL RX, LOCK`这条硬件指令来帮助实现互斥。主要过程是，把内存的`lock`读到寄存器`RX`，然后保存一个非0值到`lock`的内存地址，这个操作让处理器无法访问这块内存，直到指令完成

一个相似的指令是`XCHG`

## 同步
同步其实就是利用上述两种互斥的方案来保证共享资源每次只有一个进程在使用，所以根据互斥的方案，进程或线程间同步的主要有信号量、互斥量、条件变量和读写锁

### 信号量
信号量用一个整型变量来累计唤醒次数，信号量为0表示没有唤醒操作的线程或进程，信号量为正数时表示被唤醒进程或线程的数量。

同时信号量有两种操作，`down`和`up`(也可以称为`sleep`和`wakeup`)。`down`的操作是检查值是否大于0，大于0则减一，若该值为0，则将进程睡眠，此时`down`操作并未结束，检查数值，修改变量以及睡眠操作都是一个原子操作；`up`操作是对信号量加1，同时唤醒一个进程或线程

### 互斥量
互斥量是简化版的信号量，通常也被称为互斥锁，信号量实现容易且有效，因此经常用作线程间同步。

互斥量是一个只有两个状态的变量(解锁和加锁)，通常用一个整型(当然也可以用一个二进制位)表示，0表示解锁，其他值表示加锁。当一个线程或进程要访问共享资源时，调用`mutex_lock`，如果调用成功，则可以访问共享资源，如果失败，则线程被阻塞，直到共享资源的线程完成并调用`mutex_unlock`为止，多个进程或线程都在等待这个互斥量，则会随机选择一个并允许它获取锁

互斥量非常简单，通常使用`TSL`或`XCHG`指令实现

### 条件变量
条件变量一般是用于等待，允许线程或进程由于一些未达到的条件而阻塞。绝大部分情况下，条件变量会和互斥量一起使用。

条件变量一般有两个函数`pthread_cond_wait`和`pthread_cond_signal`比较重要，一个是用来阻塞以等待，一个是向另一个进程发送信号来唤醒，如果唤醒多个线程，则可以调用`pthread_cond_broadcast`

条件变量和互斥量一起使用的模式一般为：让一个线程锁住一个互斥量，然后当它不能获取它期待的结果时，等待一个条件变量，最后另一个线程会唤醒它，继续执行

条件变量不会存在内存中，如果将一个信号量传递给一个没有线程在等待的条件变量，那么这个信号会丢失，必须小心这种情况发生

### 读写锁
互斥锁是无论读写，都只有一个线程或进程进入共享资源。而读写锁，是把读操作和写操作分开。只要没有线程在修改，那么对于任意数目的线程，都是可以拥有读的访问权，只有在其他线程读或修改某个特定数据时，当前线程才可以修改

获取一个读写锁用于读的称为共享锁，获取一个读写锁用于写的称为独占锁，所以读写锁是一个共享-独占锁

可以使用互斥锁和条件变量来实现读写锁


## 小结
线程和进程的互斥和同步是分不开的，有互斥就有同步，互斥是从访问一个共享资源开始的，当多个线程和进程访问同一个共享资源时，就会有竞态条件，从而需要互斥的解决，互斥是同步的底层实现，只有从底层实现了互斥，才有同步的解决方案。文章主要介绍了互斥和同步的一些方案，一个很好的例子是生产者-消费者的例子，可以尝试用这几种同步方案去解决