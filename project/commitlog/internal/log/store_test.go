package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const filePattern = "store_append_read_test"

var (
	write = []byte("Hello")
	width = uint64(len(write)) + recordsWidth
)

func TestStoreAppendRead(t *testing.T) {
	// This is useful when you need a file that won’t conflict with other filenames and will be removed after use
	// If dir string "" it will create temp FS of OS, for example /tmp in Linux
	f, err := os.CreateTemp("", filePattern)
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := NewStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	// Create store again and test reading from it to verify that our service will recover its state after a restart.
	s, err = NewStore(f)
	require.NoError(t, err)
	testRead(t, s)
	err = s.Close()
	require.NoError(t, err)
}

func testAppend(t *testing.T, s *Store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		bytesWritten, pos, err := s.Append(write)
		require.NoError(t, err)
		// t.Log(pos + bytesWritten)
		require.Equal(t, pos+bytesWritten, width*i)
	}
}

func testRead(t *testing.T, s *Store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ {
		// t.Log(pos)
		v, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, v, write)
		pos = width * i
	}

}

func testReadAt(t *testing.T, s *Store) {
	t.Helper()
	var off int64
	for i := uint64(1); i < 4; i++ {
		b := make([]byte, recordsWidth)
		n, err := s.File.ReadAt(b, off)
		require.NoError(t, err)
		off += int64(n)
		r := make([]byte, enc.Uint64(b))
		w, err := s.File.ReadAt(r, off)
		require.NoError(t, err)
		require.Equal(t, r, write)
		require.Equal(t, w, len(r))
		off += int64(w)
	}
}

func TestStoreClose(t *testing.T) {
	f, err := os.CreateTemp("", filePattern)
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := NewStore(f)
	require.NoError(t, err)
	_, _, err = s.Append(write)
	require.NoError(t, err)
	_, beforeSize, err := openFile(f.Name())
	t.Log(beforeSize)
	require.NoError(t, err)

	// Close will flush the data from buffer hence after closing size of file should be greater than before size
	err = s.Close()
	require.NoError(t, err)
	_, afterSize, err := openFile(f.Name())
	t.Log(afterSize)
	require.NoError(t, err)
	require.True(t, afterSize > beforeSize)
}
func openFile(name string) (*os.File, int64, error) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fi.Size(), nil
}
