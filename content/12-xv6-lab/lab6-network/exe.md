---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1262
tags: ["xv6", "os", "network"]
---

## Exercise 1
>Exercise 1. Add a call to time_tick for every clock interrupt in kern/trap.c. Implement sys_time_msec and add it to syscall in kern/syscall.c so that user space has access to the time.

只是加一个中断，非常简单
{{< highlight c >}}
// trap.c的trap_dispatch函数
    case IRQ_OFFSET + IRQ_TIMER:
        time_tick();
        sched_yield();
        break;

// syscall.c
static int
sys_time_msec(void)
{
	// LAB 6: Your code here.
	return time_msec();
}

// syscall.c的syscall函数
	case SYS_time_msec:
		return sys_time_msec();
    
{{< /highlight  >}}

运行`make INIT_CFLAGS=-DTEST_NO_NS run-testtime-nox`可以看到从5倒数到0，再开始启动
{{< highlight bash >}}
starting count down: 5 4 3 2 1 0 
{{< /highlight  >}}

## Exercise 2
>Exercise 2. Browse Intel's Software Developer's Manual for the E1000. This manual /covers several closely related Ethernet controllers. QEMU emulates the 82540EM.

>You should skim over chapter 2 now to get a feel for the device. To write your driver, you'll need to be familiar with chapters 3 and 14, as well as 4.1 (though not 4.1's subsections). You'll also need to use chapter 13 as reference. The other chapters mostly cover components of the E1000 that your driver won't have to interact with. Don't worry about the details right now; just get a feel for how the document is structured so you can find things later.

>While reading the manual, keep in mind that the E1000 is a sophisticated device with many advanced features. A working E1000 driver only needs a fraction of the features and interfaces that the NIC provides. Think carefully about the easiest way to interface with the card. We strongly recommend that you get a basic driver working before taking advantage of the advanced features.

主要是阅读材料，如有需要，自行阅读


## Exercise 3
>Exercise 3. Implement an attach function to initialize the E1000. Add an entry to the pci_attach_vendor array in kern/pci.c to trigger your function if a matching PCI device is found (be sure to put it before the {0, 0, 0} entry that mark the end of the table). You can find the vendor ID and device ID of the 82540EM that QEMU emulates in section 5.2. You should also see these listed when JOS scans the PCI bus while booting.

>For now, just enable the E1000 device via pci_func_enable. We'll add more initialization throughout the lab.

>We have provided the kern/e1000.c and kern/e1000.h files for you so that you do not need to mess with the build system. They are currently blank; you need to fill them in for this exercise. You may also need to include the e1000.h file in other places in the kernel.

>When you boot your kernel, you should see it print that the PCI function of the E1000 card was enabled. Your code should now pass the pci attach test of make grade.

主要是跟前面文档介绍的一样，用`pci_func_enable`来启用`struct pci_func`，所以总的代码也比较简单了
{{< highlight c >}}
// kern/e1000.h

#define PCI_E1000_VENDOR_ID 0x8086
#define PCI_E1000_DEVICE_ID 0x100E

// kern/e1000.c
#include <kern/pci.h>
int
pci_func_attach(struct pci_func *pcif) 
{
    pci_func_enable(pcif);
    return 0;
}

// kern/pci.c
extern int pci_func_attach(struct pci_func *pcif);

struct pci_driver pci_attach_vendor[] = {
	{ PCI_E1000_VENDOR_ID, PCI_E1000_DEVICE_ID, &pci_func_attach},
	{ 0, 0, 0 },
};
{{< /highlight  >}}

运行`make grade`，可以看到下面输出就是成功
{{< highlight c >}}
testtime: OK (16.9s) 
pci attach: OK (10.2s) 
{{< /highlight  >}}

## Exercise 4
>Exercise 4. In your attach function, create a virtual memory mapping for the E1000's BAR 0 by calling mmio_map_region (which you wrote in lab 4 to support memory-mapping the LAPIC).

>You'll want to record the location of this mapping in a variable so you can later access the registers you just mapped. Take a look at the lapic variable in kern/lapic.c for an example of one way to do this. If you do use a pointer to the device register mapping, be sure to declare it volatile; otherwise, the compiler is allowed to cache values and reorder accesses to this memory.

>To test your mapping, try printing out the device status register (section 13.4.2). This is a 4 byte register that starts at byte 8 of the register space. You should get 0x80080783, which indicates a full duplex link is up at 1000 MB/s, among other things.

主要是为了映射关系，然后打印日志查看设备状态
{{< highlight c >}}
// kern/e1000.h
#include <inc/types.h>

#define E1000_STATUS   (0x00008/4)  /* Device Status - RO */
volatile uint32_t *e1000_bar0;

// kern/pic.c
#include <kern/pmap.h>
void
pci_func_enable(struct pci_func *f)
{
    // ... 
    	f->reg_base[regnum] = base;
		f->reg_size[regnum] = size;

        // 添加的内容
		if(regnum == 0) {
			e1000_bar0 = mmio_map_region(base, size); 
		}
    
    // ...
	cprintf("PCI function %02x:%02x.%d (%04x:%04x) enabled\n",
		f->bus->busno, f->dev, f->func,
		PCI_VENDOR(f->dev_id), PCI_PRODUCT(f->dev_id));

    // 添加的内容
    cprintf("e1000 device status: 0x%x\n", e1000_bar0[E1000_STATUS]);
}
{{< /highlight  >}}

运行`make qemu-nox`，可以看到输出的设备状态如下就是成功
{{< highlight bash >}}
e1000 device status: 0x80080783
{{< /highlight  >}}

## Exercise 5
>Exercise 5. Perform the initialization steps described in section 14.5 (but not its subsections). Use section 13 as a reference for the registers the initialization process refers to and sections 3.3.3 and 3.4 for reference to the transmit descriptors and transmit descriptor array.

>Be mindful of the alignment requirements on the transmit descriptor array and the restrictions on length of this array. Since TDLEN must be 128-byte aligned and each transmit descriptor is 16 bytes, your transmit descriptor array will need some multiple of 8 transmit descriptors. However, don't use more than 64 descriptors or our tests won't be able to test transmit ring overflow.

>For the TCTL.COLD, you can assume full-duplex operation. For TIPG, refer to the default values described in table 13-77 of section 13.4.34 for the IEEE 802.3 standard IPG (don't use the values in the table in section 14.5).


需要在`kern/e1000.c`中添加一个初始化函数。根据上面的提示，应该来说不是特别难
{{< highlight c >}}
// 除以4 作为uint32_t[]使用
#define E1000_TDBAL (0x03800/4) /* TX Descriptor Base Address Low - RW */
#define E1000_TDBAH (0x03804/4) /* TX Descriptor Base Address High - RW */
#define E1000_TDLEN (0x03808/3) /* TX Descriptor Length - RW */
#define E1000_TDH (0x3810/4) /* TX Descriptor Head - RW */
#define E1000_TDT (0x3818/4) /* TX Descripotr Tail - RW */
#define E1000_TCTL (0x00400/4)

#define E1000_TCTL_EN     0x00000002    /* enable tx */
#define E1000_TCTL_PSP    0x00000008    /* pad short packets */
#define E1000_TCTL_CT     0x00000ff0    /* collision threshold */
#define E1000_TCTL_COLD   0x003ff000    /* collision distance */

#define E1000_TIPG     0x00410  /* TX Inter-packet gap -RW */

#define E1000_TXD_CMD_RS 0x08000000
#define E1000_TXD_STAT_DD 0x00000001

struct tx_desc {
    uint64_t addr;
    uint16_t length;
    uint8_t cso;
    uint8_t cmd;
    uint8_t status;
    uint8_t css;
    uint16_t sepcial;
} __attribute__ ((aligned (16)));

struct tx_desc tdesc[64];

static void init_tx() {
    // 分配缓存空间，每页处理2MTU
    size_t tdesc_length = sizeof(tdesc) / sizeof(struct tx_desc);
    for(int i = 0; i < tdesc_length; i += 2) {
        tdesc[i].addr = page2pa(page_alloc(ALLOC_ZERO));
        tdesc[i].cmd |= (E1000_TXD_CMD_RS >> 24); // 设置RS
        tdesc[i].status |= E1000_TXD_STAT_DD; // 清除DD

        tdesc[i + 1].addr = tdesc[i].addr + PGSIZE / 2;
        tdesc[i + 1].cmd |= (E1000_TXD_CMD_RS >> 24);
        tdesc[i + 1].status |= E1000_TXD_STAT_DD;
    }
    // 初始化
    e1000_bar0[E1000_TDBAH] = 0; // 高位为0
    e1000_bar0[E1000_TDBAL] = PADDR(tdesc); // 物理地址
    e1000_bar0[E1000_TDLEN] = sizeof(tdesc); // 大小
    e1000_bar0[E1000_TDH] = 0x0; // 硬件负责更新
    e1000_bar0[E1000_TDT] = 0x0; // 软件负责更新
    e1000_bar0[E1000_TCTL] = ((0x40 << 12)&E1000_TCTL_COLD)|E1000_TCTL_PSP|E1000_TCTL_EN; // 启用tx
    e1000_bar0[E1000_TIPG] = 10; // IEEE 802.3 标准IPG
}
{{< /highlight  >}}

同时在之前的`pci_func_attach`要添加初始化
{{< highlight c >}}

int
pci_func_attach(struct pci_func *pcif) 
{
    pci_func_enable(pcif);
    init_tx();
    return 0;
}
{{< /highlight  >}}

运行`make E1000_DEBUG=TXERR,TX qemu-nox`，可以看到有`e1000: tx disabled`的信息出现就是成功

## Exercise 6
>Exercise 6. Write a function to transmit a packet by checking that the next descriptor is free, copying the packet data into the next descriptor, and updating TDT. Make sure you handle the transmit queue being full.

就是实现发送函数
{{< highlight c >}}
// kern/e1000.h
size_t e1000_transmit(const void* buffer, size_t size);
{{< /highlight  >}}

{{< highlight c >}}
size_t e1000_transmit(const void *buffer, size_t size) {
    uint32_t current = e1000_bar0[E1000_TDT];
    if (tdesc[current].status & E1000_TXD_STAT_DD){
        tdesc[current].status &= ~E1000_TXD_STAT_DD;
        void *addr = (void *)KADDR((uint32_t)tdesc[current].addr);
        size_t length = MIN(size, MTU);
        memcpy(addr, buffer, length);
        tdesc[current].cmd |= (E1000_TXD_CMD_EOP >> 24);      // End of Packet
        tdesc[current].length = length;
        // update tail pointer, inform network card
        uint32_t next = current + 1;
        e1000_bar0[E1000_TDT] = next % (sizeof(tdesc) / sizeof(struct tx_desc));
        return length;
    }
    else{
        // require for re-transmission
        cprintf("lost packet 0x%x\n", buffer);
        return 0;
    }
}
{{< /highlight  >}}

## Exercise 7
>Exercise 7. Add a system call that lets you transmit packets from user space. The exact interface is up to you. Don't forget to check any pointers passed to the kernel from user space.

添加系统调用，这里就比较简单了
{{< highlight c >}}
// inc/lib.h
int sys_send(const void *buffer, size_t length);
// inc/syscall.h
enum {
    // 添加一个枚举
    	SYS_send,
}
// lib/syscall.c
int sys_send(const void *buffer, size_t length) {
	return syscall(SYS_send, (uint32_t)buffer, length, 0, 0, 0, 0);
}

// kern/syscall.c
#include <kern/e1000.h>

static int sys_send(const void* buffer, size_t length) {
	user_mem_assert(curenv, buffer, length, PTE_U);
	return (int)e1000_transmit(buffer, length);
}

int32_t
syscall(uint32_t syscallno, uint32_t a1, uint32_t a2, uint32_t a3, uint32_t a4, uint32_t a5)
{
    // 添加一个系统调用
    	case SYS_send:
		return sys_send((const void*)a1, a2);
}
{{< /highlight  >}}

## Exercise 8
>Exercise 8. Implement net/output.c.

{{< highlight c >}}
void
output(envid_t ns_envid)
{
	binaryname = "ns_output";

	// LAB 6: Your code here:
	// 	- read a packet from the network server
	//	- send the packet to the device driver
	uint32_t reqno;
	uint32_t whom;
	int perm;
	int r;

	while(1) {
		reqno = ipc_recv((int32_t*)&whom, (void*)&nsipcbuf, &perm);
		char *ptr = nsipcbuf.pkt.jp_data;
		size_t total = (size_t)nsipcbuf.pkt.jp_len;
		if(reqno == NSREQ_OUTPUT) {
		retry:
			while((r = sys_send((const void*)ptr, total)) == 0){
				sys_yield();
			}
			if(r < total) {
				ptr += r;
				total -= r;
				cprintf("Send %d bytes, remaining %d bytes to send\n", r, total);
				goto retry;
			}
		}
	}
}
{{< /highlight  >}}

使用`make grade`可以看到`testoutput`成功

## Question
>1.How did you structure your transmit implementation? In particular, what do you do if the transmit ring is full?

transmit的结构体前面有

## Exercise 9
>Exercise 9. Read section 3.2. You can ignore anything about interrupts and checksum offloading (you can return to these sections if you decide to use these features later), and you don't have to be concerned with the details of thresholds and how the card's internal caches work.

就是一个阅读材料的练习，自行阅读

## Exercise 10
>Exercise 10. Set up the receive queue and configure the E1000 by following the process in section 14.4. You don't have to support "long packets" or multicast. For now, don't configure the card to use interrupts; you can change that later if you decide to use receive interrupts. Also, configure the E1000 to strip the Ethernet CRC, since the grade script expects it to be stripped.

>By default, the card will filter out all packets. You have to configure the Receive Address Registers (RAL and RAH) with the card's own MAC address in order to accept packets addressed to that card. You can simply hard-code QEMU's default MAC address of 52:54:00:12:34:56 (we already hard-code this in lwIP, so doing it here too doesn't make things any worse). Be very careful with the byte order; MAC addresses are written from lowest-order byte to highest-order byte, so 52:54:00:12 are the low-order 32 bits of the MAC address and 34:56 are the high-order 16 bits.

>The E1000 only supports a specific set of receive buffer sizes (given in the description of RCTL.BSIZE in 13.4.22). If you make your receive packet buffers large enough and disable long packets, you won't have to worry about packets spanning multiple receive buffers. Also, remember that, just like for transmit, the receive queue and the packet buffers must be contiguous in physical memory.

>You should use at least 128 receive descriptors

## Exercise 11
>Exercise 11. Write a function to receive a packet from the E1000 and expose it to user space by adding a system call. Make sure you handle the receive queue being empty.

## Exercise 12
>Exercise 12. Implement net/input.c.


跟发送类似，都是添加一些函数和系统调用。这里就不多赘述了

最后`make grade`,可以看到`testinput`成功

## Exercise 13
>Exercise 13. The web server is missing the code that deals with sending the contents of a file back to the client. Finish the web server by implementing send_file and send_data.

{{< highlight c >}}
static int
send_data(struct http_request *req, int fd)
{
	// LAB 6: Your code here.
	struct Stat stat;
	int r;
	if((r = fstat(fd, &stat)) < 0) {
		panic("FSTAT: %e", r);
	}
	int length = stat.st_size;
	char buffer[BUFFSIZE];
	int offset;
	for(offset = 0; offset < length; offset += BUFFSIZE) {
		memset(buffer, 0, BUFFSIZE);
		if((r = readn(fd, buffer, BUFFSIZE)) < 0) {
			panic("readn: %e", r);
		}
		write(req->sock, buffer, BUFFSIZE);
	}
	int remain = length - (offset - BUFFSIZE);
	memset(buffer, 0, BUFFSIZE);
	if((r = readn(fd, buffer, remain)) < 0) {
		panic("readn: %e", r);
	}
	write(req->sock, buffer, remain);
	return r;
}

static int
send_file(struct http_request *req)
{
	int r;
	off_t file_size = -1;
	int fd;

	// open the requested url for reading
	// if the file does not exist, send a 404 error using send_error
	// if the file is a directory, send a 404 error using send_error
	// set file_size to the size of the file

	// LAB 6: Your code here.
	struct Stat stat;
	if((r = stat(req->url, &stat)) < 0 || stat.st_isdir == 1) {
		return send_error(req, 404);
	}
	if((fd = open(req->url, O_RDONLY)) < 0) {
		goto end;
	}

	if ((r = send_header(req, 200)) < 0)
		goto end;

	if ((r = send_size(req, stat.st_size)) < 0)
		goto end;

	if ((r = send_content_type(req)) < 0)
		goto end;

	if ((r = send_header_fin(req)) < 0)
		goto end;

	r = send_data(req, fd);

end:
	close(fd);
	return r;
}
{{< /highlight  >}}

最后运行`make grade`，可以看到全部成功