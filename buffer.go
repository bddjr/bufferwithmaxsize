package bufferwithmaxsize

import "sync"

type Buffer struct {
	buf     []byte
	maxSize int
	lock    sync.Mutex
}

func NewBuffer(maxSize int) *Buffer {
	if maxSize < 1 {
		panic(`bufferwithmaxsize NewBuffer: maxSize < 1`)
	}
	return &Buffer{
		buf:     []byte{},
		maxSize: maxSize,
	}
}

func (b *Buffer) Write(ib []byte) (int, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if len(b.buf) > b.maxSize {
		nb := make([]byte, b.maxSize)
		copy(nb, b.buf[len(b.buf)-b.maxSize:])
		b.buf = nb
	}

	if len(ib) >= b.maxSize {
		if len(b.buf) < b.maxSize {
			b.buf = make([]byte, b.maxSize)
		}
		copy(b.buf, ib[len(ib)-b.maxSize:])
	} else if len(ib) > 0 {
		b.buf = append(b.buf[max(len(b.buf)+len(ib)-b.maxSize, 0):], ib...)
	}

	return len(ib), nil
}

func (b *Buffer) Bytes() (out []byte) {
	b.lock.Lock()
	out = b.buf
	b.lock.Unlock()
	return
}

func (b *Buffer) Resize(maxSize int) {
	if maxSize < 1 {
		panic(`bufferwithmaxsize Buffer.Resize: maxSize < 1`)
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	b.maxSize = maxSize

	if len(b.buf) > b.maxSize {
		nb := make([]byte, b.maxSize)
		copy(nb, b.buf[len(b.buf)-b.maxSize:])
		b.buf = nb
	}
}
