package qalam

import (
	"bufio"
	"github.com/lestrrat-go/strftime"
	"log"
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

func (q *Qalam) Write(b []byte) (int, error) {
	ct := time.Now()
	path := q.location.FormatString(ct.In(q.tloc))
	log.Println("path: ", path)
	if q.path != path {
		if q.fp != nil {
			log.Println("closing", q.path)
			q.fp.Close()
		}

		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return q.write(nil, err)
		}

		fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return q.write(nil, err)
		}
		q.fp = fp
		q.path = path
		q.bw = bufio.NewWriter(q.fp)
		q.bw = bufio.NewWriterSize(q.bw, q.bufSize)
	}
	return q.write(b, nil)
}

func (q *Qalam) write(b []byte, err error) (int, error) {
	if err == nil {
		bytesAvailable := q.bw.Available()
		log.Println(bytesAvailable, len(b))
		if bytesAvailable < len(b) {
			q.bw.Flush()
		}
		return q.bw.Write(b)
	}
	return 0, err
}
