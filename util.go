package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/text/encoding/ianaindex"
)

func Unmarshal(b []byte) (*Scene, error) {
	var data Scene
	decoder := xml.NewDecoder(bytes.NewBuffer(b))
	decoder.CharsetReader = func(charset string, reader io.Reader) (io.Reader, error) {
		enc, err := ianaindex.IANA.Encoding(charset)
		if err != nil {
			return nil, fmt.Errorf("charset %s: %s", charset, err.Error())
		}
		if enc == nil {
			// Assume it's compatible with (a subset of) UTF-8 encoding
			// Bug: https://github.com/golang/go/issues/19421
			return reader, nil
		}
		return enc.NewDecoder().Reader(reader), nil
	}

	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
