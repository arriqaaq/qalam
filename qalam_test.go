package qalam

import (
	"bufio"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/lestrrat-go/strftime"
)

func TestNew(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name string
		args args
		want *Qalam
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_Likho(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			got, err := q.Likho(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Qalam.Likho() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Qalam.Likho() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_SetBufferSize(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			q.SetBufferSize(tt.args.b)
		})
	}
}

func TestQalam_initBuffer(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			if err := q.initBuffer(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Qalam.initBuffer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQalam_Write(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			got, err := q.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Qalam.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Qalam.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_bytesAvailable(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			if got := q.bytesAvailable(); got != tt.want {
				t.Errorf("Qalam.bytesAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_Writeln(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			got, err := q.Writeln(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Qalam.Writeln() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Qalam.Writeln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_write(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			got, err := q.write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Qalam.write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Qalam.write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_writeln(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			got, err := q.writeln(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Qalam.writeln() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Qalam.writeln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQalam_Close(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Qalam{
				fp:       tt.fields.fp,
				location: tt.fields.location,
				path:     tt.fields.path,
				tloc:     tt.fields.tloc,
				bufSize:  tt.fields.bufSize,
				bw:       tt.fields.bw,
				mu:       tt.fields.mu,
			}
			q.Close()
		})
	}
}
