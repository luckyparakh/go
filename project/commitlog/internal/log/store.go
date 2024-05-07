package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian
)

const (
	recordsWidth = 8 // defines the number of bytes used to store the record’s length
)

type Store struct {
	mu   sync.Mutex
	size uint64
	buf  *bufio.Writer
	*os.File
}

func NewStore(f *os.File) (*Store, error) {

	// It calls os.Stat(name string) to get the file’s current size, in case we’re re-creating the store from a file that has existing
	// data, which would happen if, for example, our service had restarted.
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	size := uint64(fi.Size())
	return &Store{
		File: f,
		size: size,
		mu:   sync.Mutex{},
		/*
			Write to the buffered writer instead of directly to the file to reduce the number of system calls and improve performance.
			If a user wrote a lot of small records, this would help a lot.
		*/
		buf: bufio.NewWriter(f),
	}, nil
}

// Append appends a byte slice to the store and returns the number of bytes written, the position of the appended data, and any error encountered.
// The function locks the store, writes the length of the record to the buffer using binary.Write, writes the record data to the buffer
// using s.buf.Write, updates the size of the store, and returns the number of bytes written, the position of the appended data, and any error encountered.
// If an error occurs while writing the length of the record or the record data, the function returns 0, 0, and the error.
// The function uses a defer statement to unlock the store before returning.
func (s *Store) Append(p []byte) (uint64, uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// position from where writing in file will start
	pos := s.size

	// Write len of record, so that while reading we know how many bytes to read
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil { // Why uint64
		return 0, 0, err
	}

	noBytesWritten, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}
	noBytesWritten += recordsWidth
	s.size += uint64(noBytesWritten)
	return uint64(noBytesWritten), pos, nil
}

// Read reads a record from the store at the specified position and returns the record data as a byte slice.
// The function locks the store, flushes the buffer, reads the length of the record from the file at the specified position,
// reads the record data from the file at the position after the length, and returns the record data and any error encountered.
// If an error occurs while flushing the buffer or reading the record length or data, the function returns nil and the error.
// The function uses a defer statement to unlock the store before returning.
func (s *Store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Writes buffered data into underlying writer
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	// Find length of record
	recordLen := make([]byte, recordsWidth)
	if _, err := s.File.ReadAt(recordLen, int64(pos)); err != nil {
		return nil, err
	}

	// Read number of bytes as per recordLen from File
	// enc.Uint64 is used as recordLen is binary number so it should read in Big-endian way
	record := make([]byte, enc.Uint64(recordLen))
	if _, err := s.File.ReadAt(record, int64(pos+recordsWidth)); err != nil {
		return nil, err
	}
	// This record will escape as it will live beyond the lifetime of current func call
	return record, nil
}

// It implements io.ReaderAt on the store type
func (s *Store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Writes buffered data into underlying writer
	if err := s.buf.Flush(); err != nil {
		return 0, err
	}

	// Read len(p) byte into p start from offset location of the File
	return s.File.ReadAt(p, off)
}

// It persists any buffered data before closing the file.
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Writes buffered data into underlying writer
	if err := s.buf.Flush(); err != nil {
		return err
	}
	return s.File.Close()
}
