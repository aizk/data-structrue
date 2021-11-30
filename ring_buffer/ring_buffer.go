package ring_buffer

import "sync"

type Buffer struct {
	start int
	end int
	buf [][]byte
	cond *sync.Cond
	closed bool
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		start: 0,
		end: 0,
		buf: make([][]byte, size),
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Buffer) Cap() int {
	return cap(b.buf)
}

// (end - start + cap) % cap
// (end - start) 在负数时，这中间的不是数据，且结果为负值，此时 + cap 正好是数据数
// 使用的 buffer 长度
func (b *Buffer) usedBufferLen() int {
	return (b.end + cap(b.buf) - b.start) % cap(b.buf)
}

func (b *Buffer) Put(data []byte) bool {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	if b.closed {
		return false
	}

	// if there is only 1 free slot, we allocate more
	// 空间不够有两个选择：扩容、阻塞等待被消费
	if b.IsFull() {
		b.Allocate()
	}

	b.buf[b.end] = data
	// end move
	b.end = (b.end + 1) % cap(b.buf)
	b.cond.Signal()
	return true
}

func (b *Buffer) Pop() ([]byte, bool) {
	for {
		b.cond.L.Lock()
		if b.usedBufferLen() > 0 {
			data := b.buf[b.start]
			b.start = (b.start + 1) % b.Cap()
			b.cond.L.Unlock()
			return data, true
		}
		if b.closed {
			b.cond.L.Unlock()
			return nil, false
		}
		b.cond.Wait()
		b.cond.L.Unlock()
	}
}

// 差一个就满
func (b *Buffer) IsFull() bool {
	return (b.end + 1) % b.Cap() == b.start
}

func (b *Buffer) Allocate() {
	cap := b.Cap()
	buf := make([][]byte, cap * 2)
	if b.end < b.start {
		copy(buf, b.buf[b.start: cap])
		copy(buf[cap - b.start:], b.buf[0:b.end])
	} else {
		copy(buf, b.buf[b.start: b.end])
	}

	b.buf = buf
	b.start = 0
	b.end = cap - 1
	return
}
