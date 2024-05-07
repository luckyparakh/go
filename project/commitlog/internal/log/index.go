package log

import (
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size())

	// os.Truncate is used to truncate the file to Segment.MaxIndexBytes bytes. If file.txt is larger than Segment.MaxIndexBytes bytes,
	// it will be truncated to Segment.MaxIndexBytes bytes, and any data beyond Segment.MaxIndexBytes bytes will be discarded.
	// we resize them now is that, once they’re memory-mapped, we can’t resize them, so it’s now or never
	if err := os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}

	// It maps the file descriptor of the idx.file to memory using the gommap.Map function,
	// These are the memory protection flags. gommap.PROT_READ allows the mapped memory to be read, and gommap.PROT_WRITE allows it to be written to.
	// This is the mapping type flaggommap.MAP_SHARED creates a shared mapping, which means that changes to the mapped
	// memory are carried through to the underlying file and are visible to other processes mapping the same file.
	if idx.mmap, err = gommap.Map(idx.file.Fd(), gommap.PROT_READ|gommap.PROT_WRITE, gommap.MAP_SHARED); err != nil {
		return nil, err
	}
	return idx, nil
}

func (i *index) Close() error {

	// This line attempts to sync the memory-mapped file (mmap) with the physical storage.
	// gommap.MS_SYNC is a flag that ensures that the changes are written immediately to the physical storage
	err := i.mmap.Sync(gommap.MS_SYNC)
	if err != nil {
		return err
	}

	// This line attempts to sync the os.File (file) with the physical storage. This ensures that any changes made to the file are written to the disk
	err = i.file.Sync()
	if err != nil {
		return err
	}

	// This line truncates or extends the file to the specified size. If the file’s current size is larger than the specified size, the extra data is discarded.
	// If the file’s current size is smaller than the specified size, the file is extended and the extended part reads as null bytes.
	/*
		When we start our service, the service needs to know the offset to set on the
		next record appended to the log. The service learns the next record’s offset
		by looking at the last entry of the index, a simple process of reading the last
		12 bytes of the file. However, we mess up this process when we grow the files
		so we can memory-map them. (The reason we resize them now is that, once
		they’re memory-mapped, we can’t resize them, so it’s now or never.) We grow
		the files by appending empty space at the end of them, so the last entry is no
		longer at the end of the file—instead, there’s some unknown amount of space
		between this entry and the file’s end. This space prevents the service from
		restarting properly. That’s why we shut down the service by truncating the
		index files to remove the empty space and put the last entry at the end of the
		file once again. This graceful shutdown returns the service to a state where
		it can restart properly and efficiently.

	*/
	err = i.file.Truncate(int64(i.size))
	if err != nil {
		return err
	}
	return i.file.Close()
}

func (i *index) Read(inOff int64) (out uint32, pos uint64, err error) {
	/*
		Index file looks like
		0      1       2       3        n-1
		off,pos|off,pos|off,pos|--------|off,pos
	*/

	// If the size of the file is 0, it returns an io.EOF error.
	if i.size == 0 {
		return 0, 0, io.EOF
	}

	// If inOff is -1, it calculates the index of the last entry by dividing the size of the file by the width of each entry and subtracting 1.
	// 1 is subtracted because 0 is always the offset of the index’s first entry.
	if inOff == -1 {
		out = uint32((i.size / entWidth) - 1)
	} else {
		out = uint32(inOff)
	}
	pos = uint64(out) * entWidth
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}
	out = enc.Uint32(i.mmap[pos : pos+offWidth])
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth])
	return out, pos, nil
}
func (i *index) Write(off uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+entWidth {
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += entWidth
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
