---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: /covers/xv6.png
cover: false 
weight: 2
tags: ["xv6", "os", "shell"]
---

## Exercise 1
>Exercise 1. i386_init identifies the file system environment by passing the type ENV_TYPE_FS to your environment creation function, env_create. Modify env_create in env.c, so that it gives the file system environment I/O privilege, but never gives that privilege to any other environment.

>Make sure you can start the file environment without causing a General Protection fault. You should pass the "fs i/o" test in `make grade`.

这个就非常简单，就判断下类型，然后是`ENV_TYPE_FS`的话，就设置`IOPL`

{{< highlight c >}}
void
env_create(uint8_t *binary, enum EnvType type)
{

	// If this is the file server (type == ENV_TYPE_FS) give it I/O privileges.
	// LAB 5: Your code here.
	struct Env *penv;
	int r = env_alloc(&penv, 0);
	if(r < 0) {
		panic("env_create %e", r);
	}
	penv->env_type = type;
	if(type == ENV_TYPE_FS) {
		penv->env_tf.tf_eflags |= FL_IOPL_MASK;
	}
	load_icode(penv, binary);
}
{{< /highlight  >}}

运行`make grade`，看到下列的输出，就算成功
{{< highlight bash >}}
internal FS tests [fs/test.c]: OK (9.2s) 
  fs i/o: OK 
{{< /highlight  >}}

## Question
>1. Do you have to do anything else to ensure that this I/O privilege setting is saved and restored properly when you subsequently switch from one environment to another? Why?

没有，因为是由硬件保存的

## Exercise 2
>Exercise 2. Implement the bc_pgfault and flush_block functions in fs/bc.c. bc_pgfault is a page fault handler, just like the one your wrote in the previous lab for copy-on-write fork, except that its job is to load pages in from the disk in response to a page fault. When writing this, keep in mind that (1) addr may not be aligned to a block boundary and (2) ide_read operates in sectors, not blocks.

>The flush_block function should write a block out to disk if necessary. flush_block shouldn't do anything if the block isn't even in the block cache (that is, the page isn't mapped) or if it's not dirty. We will use the VM hardware to keep track of whether a disk block has been modified since it was last read from or written to disk. To see whether a block needs writing, we can just look to see if the PTE_D "dirty" bit is set in the uvpt entry. (The PTE_D bit is set by the processor in response to a write to that page; see 5.2.4.3 in chapter 5 of the 386 reference manual.) After writing the block to disk, flush_block should clear the PTE_D bit using sys_page_map.

>Use `make grade` to test your code. Your code should pass "check_bc", "check_super", and "check_bitmap".

主要实现`fs/bc.c`中的`bc_pgfault`和`flush_block`。主要实现如下
{{< highlight c >}}
static void
bc_pgfault(struct UTrapframe *utf)
{
	void *addr = (void *) utf->utf_fault_va;
	uint32_t blockno = ((uint32_t)addr - DISKMAP) / BLKSIZE;
	int r;

	// Check that the fault was within the block cache region
	if (addr < (void*)DISKMAP || addr >= (void*)(DISKMAP + DISKSIZE))
		panic("page fault in FS: eip %08x, va %08x, err %04x",
		      utf->utf_eip, addr, utf->utf_err);

	// Sanity check the block number.
	if (super && blockno >= super->s_nblocks)
		panic("reading non-existent block %08x\n", blockno);

	// LAB 5: you code here:
	addr = ROUNDDOWN(addr, PGSIZE);
	sys_page_alloc(0, addr, PTE_P|PTE_W|PTE_U);
	if((r = ide_read(blockno * BLKSECTS, addr, BLKSECTS)) < 0) {
		panic("ide_read: %e", r);
	}

	// block from disk
	if ((r = sys_page_map(0, addr, 0, addr, uvpt[PGNUM(addr)] & PTE_SYSCALL)) < 0)
		panic("in bc_pgfault, sys_page_map: %e", r);

	// Check that the block we read was allocated. (exercise for
	// the reader: why do we do this *after* reading the block
	// in?)
	if (bitmap && block_is_free(blockno))
		panic("reading free block %08x\n", blockno);
}
{{< /highlight  >}}

{{< highlight c >}}
void
flush_block(void *addr)
{
	uint32_t blockno = ((uint32_t)addr - DISKMAP) / BLKSIZE;

	if (addr < (void*)DISKMAP || addr >= (void*)(DISKMAP + DISKSIZE))
		panic("flush_block of bad va %08x", addr);

	// LAB 5: Your code here.
	addr = ROUNDDOWN(addr, PGSIZE);
	if(!va_is_mapped(addr) || !va_is_dirty(addr)) {
		return;
	}
	int r;
	if((r = ide_write(blockno * BLKSECTS, addr, BLKSECTS)) < 0) {
		panic("ide_write: %e", r);
	}
	if((r = sys_page_map(0, addr, 0, addr, uvpt[PGNUM(addr)] & PTE_SYSCALL)) < 0) {
		panic("sys_page_map: %e", r);
	}
}
{{< /highlight  >}}

使用`make grade`，如果出现以下内容，则说明是成功的
{{< highlight bash >}}
check_bc: OK 
check_super: OK 
check_bitmap: OK 
{{< /highlight  >}}

## Exercise 3
>Exercise 3. Use free_block as a model to implement alloc_block in fs/fs.c, which should find a free disk block in the bitmap, mark it used, and return the number of that block. When you allocate a block, you should immediately flush the changed bitmap block to disk with flush_block, to help file system consistency.

>Use make grade to test your code. Your code should now pass "alloc_block".

具体实现看下面：
{{< highlight c >}}
int
alloc_block(void)
{
	// The bitmap consists of one or more blocks.  A single bitmap block
	// contains the in-use bits for BLKBITSIZE blocks.  There are
	// super->s_nblocks blocks in the disk altogether.

	// LAB 5: Your code here.
	for(uint32_t blockno = 1; blockno < super->s_nblocks; blockno++) {
		if(!block_is_free(blockno)) continue;
		bitmap[blockno/32] &= ~(1 << (blockno % 32));
		flush_block(&bitmap[blockno / 32]);
		return blockno;
	}
	return -E_NO_DISK;
}
{{< /highlight  >}}

使用`make grade`，如果输出以下结果，则说明成功
{{< highlight bash >}}
alloc_block: OK 
{{< /highlight  >}}

## Exercise 4
>Exercise 4. Implement file_block_walk and file_get_block. file_block_walk maps from a block offset within a file to the pointer for that block in the struct File or the indirect block, very much like what pgdir_walk did for page tables. file_get_block goes one step further and maps to the actual disk block, allocating a new one if necessary.

>Use make grade to test your code. Your code should pass "file_open", "file_get_block", and "file_flush/file_truncated/file rewrite", and "testfile".

主要是根据提示来实现两个函数`file_block_walk`和`file_get_block`

{{< highlight c >}}
static int
file_block_walk(struct File *f, uint32_t filebno, uint32_t **ppdiskbno, bool alloc)
{
       // LAB 5: Your code here.
	if(filebno >= NDIRECT + NINDIRECT) return -E_INVAL;

	uint32_t *addr = NULL;
	if(filebno < NDIRECT) {
		addr = &(f->f_direct[filebno]);
	} else {
		int r;
		if(f->f_indirect== 0) {
			if(alloc) {
				r = alloc_block();
				if(r < 0) return r;
				memset(diskaddr(r), 0, BLKSIZE);
				f->f_indirect = r;
			} else {
				return -E_NOT_FOUND;
			}
		}
		uint32_t *indir = (uint32_t*)diskaddr(f->f_indirect);
		addr = indir + (filebno - NDIRECT);
	}
	*ppdiskbno = addr;
	return 0;

}
{{< /highlight  >}}

{{< highlight c >}}
int
file_get_block(struct File *f, uint32_t filebno, char **blk)
{
       // LAB 5: Your code here.
	uint32_t *addr = NULL;
	int r = file_block_walk(f, filebno, &addr, 1);
	if(r < 0) return r;
	if((*addr) == 0) {
		r = alloc_block();
		if(r < 0) return r;
		else *addr = r;
	}
	*blk = (char*)diskaddr(*addr);
	flush_block(*blk);
	return 0;
}
{{< /highlight  >}}

运行`make grade`，出现以下内容，则说明正确
{{< highlight bash >}}
file_open: OK 
  file_get_block: OK 
  file_flush/file_truncate/file rewrite: OK 
testfile: OK (8.4s) 
{{< /highlight  >}}

## Exercise 5
>Exercise 5. Implement serve_read in fs/serv.c.

>serve_read's heavy lifting will be done by the already-implemented file_read in fs/fs.c (which, in turn, is just a bunch of calls to file_get_block). serve_read just has to provide the RPC interface for file reading. Look at the comments and code in serve_set_size to get a general idea of how the server functions should be structured.

>Use make grade to test your code. Your code should pass "serve_open/file_stat/file_close" and "file_read" for a score of 70/150.

按照要求来实现，应该是以下代码
{{< highlight c >}}
int
serve_read(envid_t envid, union Fsipc *ipc)
{
	struct Fsreq_read *req = &ipc->read;
	struct Fsret_read *ret = &ipc->readRet;

	if (debug)
		cprintf("serve_read %08x %08x %08x\n", envid, req->req_fileid, req->req_n);

	// Lab 5: Your code here:
	int r;
	struct OpenFile *o;
	if ((r = openfile_lookup(envid, req->req_fileid, &o)) < 0)
		return r;
	
	if ((r = file_read(o->o_file, ret->ret_buf, req->req_n, o->o_fd->fd_offset)) < 0)
		return r;

	o->o_fd->fd_offset += r;
	return r;
}
{{< /highlight  >}}

## Exercise 6
>Exercise 6. Implement serve_write in fs/serv.c and devfile_write in lib/file.c.

>Use make grade to test your code. Your code should pass "file_write", "file_read after file_write", "open", and "large file" for a score of 90/150.

同样，写就很简单了
{{< highlight c >}}
int
serve_write(envid_t envid, struct Fsreq_write *req)
{
	if (debug)
		cprintf("serve_write %08x %08x %08x\n", envid, req->req_fileid, req->req_n);

	// LAB 5: Your code here.
	int r;
	struct OpenFile *o;
	if ((r = openfile_lookup(envid, req->req_fileid, &o)) < 0)
		return r;
	
	if ((r = file_write(o->o_file, req->req_buf, req->req_n, o->o_fd->fd_offset)) < 0)
		return r;

	o->o_fd->fd_offset += r;
	return r;
}
{{< /highlight  >}}

同时，还需要实现`lib/file.c`中的`devfile_write`函数。参考`devfile_read`函数就很容易实现
{{< highlight c >}}
static ssize_t
devfile_write(struct Fd *fd, const void *buf, size_t n)
{
	// LAB 5: Your code here
	int r;
	
	fsipcbuf.write.req_fileid = fd->fd_file.id;
	fsipcbuf.write.req_n = n;
	memmove(fsipcbuf.write.req_buf, buf, n);
	if((r = fsipc(FSREQ_WRITE, NULL)) < 0) {
		return r;
	}
	assert(r <= n);
	assert(r <= sizeof(fsipcbuf.write.req_buf));
	return r;
}
{{< /highlight  >}}


运行`make grade`,看到以下内容就是成功
{{< highlight bash >}}
testfile: OK (10.1s) 
  serve_open/file_stat/file_close: OK 
  file_read: OK 
  file_write: OK 
  file_read after file_write: OK 
  open: OK 
  large file: OK 
{{< /highlight  >}}

## Exercise 7
>Exercise 7. spawn relies on the new syscall sys_env_set_trapframe to initialize the state of the newly created environment. Implement sys_env_set_trapframe in kern/syscall.c (don't forget to dispatch the new system call in syscall()).

>Test your code by running the user/spawnhello program from kern/init.c, which will attempt to spawn /hello from the file system.

>Use make grade to test your code.

根据要求，就是实现`kern/syscall.c`中的`sys_env_set_trapframe`函数。所以根据这个也就很好实现了
{{< highlight c >}}
static int
sys_env_set_trapframe(envid_t envid, struct Trapframe *tf)
{
	// LAB 5: Your code here.
	// Remember to check whether the user has supplied us with a good
	// address!
	int r;
	struct Env *env = NULL;
	if((r = envid2env(envid, &env, 1)) != 0) return r;

	user_mem_assert(env, tf, sizeof(struct Trapframe), PTE_U);
	env->env_tf = *tf;
	env->env_tf.tf_cs = GD_UT | 3;
	env->env_tf.tf_eflags |= FL_IF;
	env->env_tf.tf_eflags &= ~FL_IOPL_MASK;
	return 0;
}
{{< /highlight  >}}

同时，这个是个系统调用，需要在`syscall`函数中添加上述的系统调用
{{< highlight c >}}
	case SYS_env_set_trapframe:
		return sys_env_set_trapframe(a1, (struct Trapframe*)a2);
{{< /highlight  >}}

最后`make grade`，出现下面的情况，说明代码成功
{{< highlight bash >}}
spawn via spawnhello: OK (8.6s) 
    (Old jos.out.spawn failure log removed)
Protection I/O space: OK (8.3s) 
{{< /highlight  >}}

## Exercise 8
>Exercise 8. Change duppage in lib/fork.c to follow the new convention. If the page table entry has the PTE_SHARE bit set, just copy the mapping directly. (You should use PTE_SYSCALL, not 0xfff, to mask out the relevant bits from the page table entry. 0xfff picks up the accessed and dirty bits as well.)

>Likewise, implement copy_shared_pages in lib/spawn.c. It should loop through all page table entries in the current process (just like fork did), copying any page mappings that have the PTE_SHARE bit set into the child process.

先要修改`lib/fork.c`中的`duppage`函数，需要验证`PTE_SHARE`标识位。这个比较简单
{{< highlight c >}}
static int
duppage(envid_t envid, unsigned pn)
{
	int r;

	// LAB 4: Your code here.
	int perm = PTE_P | PTE_U;	// at least PTE_P and PTE_U
	// envid_t curenvid = sys_getenvid();

	int is_wr = (uvpt[pn] & PTE_W) == PTE_W;
	int is_cow = (uvpt[pn] & PTE_COW) == PTE_COW;
	int is_shared = (uvpt[pn] & PTE_SHARE); // PTE_SHARE验证
	void *addr = (void *)(pn * PGSIZE);
	if ((is_wr || is_cow) && !is_shared)
	{
		// create new mapping
		if ((r = sys_page_map(0, addr, envid, addr, perm | PTE_COW)) != 0)
			panic("sys_page_map, %e", r);
		if ((r = sys_page_map(0, addr, 0, addr, perm | PTE_COW)) != 0)
			panic("sys_page_map, %e", r);
	}
	else
	{
		if (is_shared)
			perm = PTE_SYSCALL & uvpt[pn];
		// only remap child without PTE_COW
		if ((r = sys_page_map(0, addr, envid, addr, perm)) != 0)
			panic("sys_page_map, %e", r);
	}
	return 0;
}
{{< /highlight  >}}

然后完成`lib/spawn.c`中的`copy_shared_pages`函数。主要是遍历页表，然后根据是否是`PTE_SHARE`页面，进行页面映射(也就是拷贝到子进程中，文档中说的比较清楚)
{{< highlight c >}}
static int
copy_shared_pages(envid_t child)
{
	// LAB 5: Your code here.
	int r;
	void *addr = 0;
	for(uint32_t i = 0; i < UTOP / PGSIZE; i++) {
		addr = (void*)(i * PGSIZE);
		if((uvpd[PDX(addr)] & PTE_P) && (uvpt[i] & PTE_P)) {
			if(uvpt[i] & PTE_SHARE) {
				if((r = sys_page_map(0, addr, child, addr, PTE_SYSCALL & uvpt[i])) != 0){
					panic("copy_shared_pages: %e", r);
				}
			}
		}
	}
	return 0;
}
{{< /highlight  >}}

最后运行`make run-testpteshare-nox`， 出现以下内容，则说明代码表现正常
{{< highlight bash >}}
fork handles PTE_SHARE right
spawn handles PTE_SHARE right
{{< /highlight  >}}

运行`make run-testfdsharing-nox`， 出现以下内容，说明文件描述符拷贝正常
{{< highlight bash >}}
read in child succeeded
read in parent succeeded
{{< /highlight  >}}

最后可以运行`make grade`， 可以出现以下内容
{{< highlight bash >}}
spawn via spawnhello: OK (9.3s) 
Protection I/O space: OK (8.7s) 
PTE_SHARE [testpteshare]: OK (9.3s) 
    (Old jos.out.pte_share failure log removed)
PTE_SHARE [testfdsharing]: OK (9.1s) 
{{< /highlight  >}}

## Exercise 9
>Exercise 9. In your kern/trap.c, call kbd_intr to handle trap IRQ_OFFSET+IRQ_KBD and serial_intr to handle trap IRQ_OFFSET+IRQ_SERIAL.

这部分在`kern/trap.c`中的`trap_dispatch`的`switch-case`添加以下内容就可以了
{{< highlight c >}}
        // 串口和键盘，需要直接返回，不能进行下面的trapframe操作
		case IRQ_OFFSET + IRQ_KBD:
			kbd_intr();
			return;
		case IRQ_OFFSET + IRQ_SERIAL:
			serial_intr();
			return;
{{< /highlight  >}}

运行`make run-testkbd-nox`，可以输入内容，然后echo同样的内容，则说明成功

{{< highlight bash >}}
Type a line: a
a
Type a line: b
b
Type a line: c
c
Type a line: d
d
Type a line: e
e
Type a line: f
f
Type a line: g
g
{{< /highlight  >}}

## Exercise 10
> Exercise 10. The shell doesn't support I/O redirection. It would be nice to run `sh <script` instead of having to type in all the commands in the script by hand, as you did above. Add I/O redirection for < to user/sh.c.

>Test your implementation by typing sh `<script` into your shell

>Run make run-testshell to test your shell. testshell simply feeds the above commands (also found in fs/testshell.sh) into the shell and then checks that the output matches fs/testshell.key.


运行`make run-testshell-nox`，可以看到以下内容
{{< highlight bash >}}
running sh -x < testshell.sh | cat
shell ran correctly
{{< /highlight  >}}

最后`make grade`， 可以看到所有的都成功了