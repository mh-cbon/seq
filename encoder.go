// Package seq almost implements `RFC7464`
// to realize `application/json-seq` streaming.
package seq

import (
	"encoding/json"
	"io"
)

//RFC7464 provides an encoder compatible with the aforementionned spec.
func RFC7464(out io.Writer) *Encoder {
	return &Encoder{
		EncoderWriter: EncoderWriter{
			BytesEncoder: json.NewEncoder(out),
			Writer:       out,
		},
		PrefixSeparator: []byte{30},
		// SuffixSeparator: []byte{10}, // dont set it, golang json encoder already does it.
	}
}

//BytesEncoder encodes any kind of data to []byte.
type BytesEncoder interface {
	Encode(interface{}) error
}

//EncoderWriter is something that can Encode and Write on the same underlying storage.
type EncoderWriter struct {
	BytesEncoder
	io.Writer
}

// Encoder of anykind of data wrapped with prefix and/or suffis spearators.
//TODO: It does not escape separators from the encoded data.
type Encoder struct {
	EncoderWriter
	PrefixSeparator []byte
	SuffixSeparator []byte
	Dest            io.Writer
}

//Encode anything using wand wraps it with separators.
func (s *Encoder) Encode(v interface{}) (err error) {
	writers := []writer{
		s.writeRaw(s.PrefixSeparator),
		s.writeMarshalled(v),
		s.writeRaw(s.SuffixSeparator),
	}
	for _, write := range writers {
		if write == nil {
			continue
		}
		if err := write(); err != nil {
			return err
		}
	}
	return nil
}

//writer reduces write operation to an error, or nil on success
type writer func() error

//writeRaw []byte on the underlying storage.
func (s *Encoder) writeRaw(v []byte) func() error {
	if v == nil || len(v) == 0 {
		return nil
	}
	return func() error {
		if n, err := s.EncoderWriter.Write(v); err != nil {
			return err
		} else if n != len(v) {
			return io.ErrShortWrite
		}
		return nil
	}
}

//writeMarshalled anyhting.
func (s *Encoder) writeMarshalled(v interface{}) func() error {
	return func() error {
		return s.EncoderWriter.Encode(v)
	}
}
