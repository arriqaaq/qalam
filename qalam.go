package qalam

import (
	"bufio"
	"github.com/lestrrat-go/strftime"
	"os"
	"path/filepath"
	"time"
)

type (
	Qalam struct {
		fp *os.File
		// the location of the file
		location *strftime.Strftime

		// file name created by builder
		path string

		// time location
		tloc *time.Location

		// Add prom stats
		// Add zap logger later
		// bufio size
		bufSize int
		// bufio writer
		bw *bufio.Writer
	}
)

var (
	now               = time.Now()
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

func (q *Qalam) Writeln(b []byte) (int, error) {
	ct := time.Now()
	path := q.location.FormatString(ct.In(q.tloc))
	if q.path != path {
		if q.fp != nil {
			q.fp.Close()
		}

		err := q.initBuffer(path)
		if err != nil {
			return q.write(nil, err, false)
		}
		q.path = path
	}
	return q.write(b, nil, true)
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

	bw := bufio.NewWriter(fp)
	bw = bufio.NewWriterSize(bw, q.bufSize)
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
			return q.write(nil, err, false)
		}
		q.path = path
	}
	return q.write(b, nil, false)
}

func (q *Qalam) write(b []byte, err error, newline bool) (int, error) {
	if err == nil {
		bytesAvailable := q.bw.Available()
		if bytesAvailable < len(b) {
			q.bw.Flush()
		}

		if newline {
			q.bw.Write(b)
			return q.bw.Write([]byte("\n"))
		}
		return q.bw.Write(b)

	}
	return 0, err
}
