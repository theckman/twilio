// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"bytes"
	"encoding/xml"
	"io"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var xmlHeader = []byte(xml.Header)

// Response represents a full TwiML response. TwiML is used to instruct Twilio
// on what to do with a phone call.
type Response struct {
	XMLName xml.Name `xml:"Response"`
	Verbs   []interface{}
}

// EncodeResponse takes a *Response instance and encodes it, writing it to w.
func EncodeResponse(w io.Writer, r *Response) error {
	// get a new XML encoder for writing to the buffer
	// enable indenting of output
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")

	// write the xml header to the buffer
	if _, err := w.Write(xmlHeader); err != nil {
		return err
	}

	// encode the *Response in to XML and write it to the buffer
	if err := encoder.Encode(r); err != nil {
		return err
	}

	return nil
}

// MarshalResponse takes a *Response instance and renders it to XML.
func MarshalResponse(r *Response) ([]byte, error) {
	// get a new *bytes.Buffer from the pool
	buf := bufferPool.Get().(*bytes.Buffer)

	// defers are first in, last out...
	// so when we return, this will call buf.Reset() first
	defer bufferPool.Put(buf)
	defer buf.Reset()

	if err := EncodeResponse(buf, r); err != nil {
		return nil, err
	}

	// copy the byte slice from the *bytes.Buffer as any changes to the buffer
	// will impact the slice that .Bytes() returns
	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())

	return out, nil
}

// EncodeSlice takes a []inteface{}, allocates a *Response instances, and
// encodes it to w.
func EncodeSlice(w io.Writer, s []interface{}) error {
	return EncodeResponse(w, &Response{Verbs: s})
}

// MarshalSlice takes a []interface{}, allocates a *Response instance, and calls
// MarshalResponse with it.
func MarshalSlice(s []interface{}) ([]byte, error) {
	return MarshalResponse(&Response{Verbs: s})
}
