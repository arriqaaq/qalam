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

func TestQalam_Write(t *testing.T) {
	type fields struct {
		fp       *os.File
		location *strftime.Strftime
		path     string
		tloc     *time.Location
		bufSize  int
		bw       *bufio.Writer
		mu       sync.RWMutex
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
