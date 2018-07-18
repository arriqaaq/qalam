package qalam

import (
	"bufio"
	"github.com/lestrrat-go/strftime"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// How to make something thread safe? Research and implement

type (
	Qalam struct {
		fp *os.File
		// the location of the file
		location *strftime.Strftime

		// file name created by builder
		path string

		// time location
		tloc *time.Location

		// bufio size
		bufSize int

		// bufio writer
		bw *bufio.Writer

		mu sync.Mutex
	}
)

var (
	DefaultBufferSize = 4096
)

func New(location string) *Qalam {
	p, err := strftime.New(location)
	if err != nil {
		panic(err)
	}

	return &Qalam{
		location: p,
		tloc:     time.Local,
		bufSize:  DefaultBufferSize,
	}
}

func (q *Qalam) Likho(b []byte) (int, error) {
	return q.Write(b)
}

/*
	SetBufferSize set's the size of the buffer
	which is kept in memory before pushing to disk.
	Defaults to 4096, the default page size on older
	SSDs, can be set accordingly
*/
func (q *Qalam) SetBufferSize(b int) {
	q.bufSize = b
}

func (q *Qalam) initBuffer(path string) (err error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	bw := bufio.NewWriterSize(fp, q.bufSize)

	q.path = path
	q.fp = fp
	q.bw = bw
	return nil
}

func (q *Qalam) Write(b []byte) (int, error) {
	ct := time.Now()
	path := q.location.FormatString(ct.In(q.tloc))
	if q.path != path {
		if q.fp != nil {
			q.fp.Close()
		}

		err := q.initBuffer(path)
		if err != nil {
			return 0, err
		}
	}
	return q.write(b)
}

func (q *Qalam) bytesAvailable() int {
	return q.bw.Available()
}

func (q *Qalam) Writeln(b []byte) (int, error) {
	ct := time.Now()
	path := q.location.FormatString(ct.In(q.tloc))
	if q.path != path {
		if q.fp != nil {
			q.fp.Close()
		}
		err := q.initBuffer(path)
		if err != nil {
			return 0, err
		}
	}
	return q.writeln(b)
}

// Avoid data race when writing
func (q Qalam) write(b []byte) (int, error) {
	if q.bytesAvailable() < len(b) {
		q.bw.Flush()
	}
	return q.bw.Write(b)
}

// Avoid data race when writing
func (q Qalam) writeln(b []byte) (int, error) {
	if q.bytesAvailable() < len(b) {
		q.bw.Flush()
	}

	// Newline must always be appended
	q.bw.Write(b)
	return q.bw.Write([]byte("\n"))
}

/*
A successful close does not guarantee that the data has been successfully saved to disk,
as the kernel defers writes. It is not common for a file system to flush the buffers
when the stream is closed. If you need to be sure that the data is physically
stored use fsync(2). (It will depend on the disk hardware at this point.)
*/
func (q *Qalam) Close() {
	q.bw.Flush()
	q.fp.Sync()
	q.fp.Close()
}

func (q *Qalam) Likholn(b []byte) (int, error) {
	return q.Writeln(b)
}
