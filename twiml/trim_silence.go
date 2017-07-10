// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// Trim lets you specify whether to trim leading and trailing silence from your
// audio files.
type Trim uint8

const (
	// TrimSilence is the default and instructs Twilio to trim silence from
	// recordings.
	TrimSilence Trim = 1 << iota

	// DoNotTrimSilence instructs Twilio to not trim silence from recordings.
	DoNotTrimSilence
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (t Trim) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: t.String(),
	}

	return attr, nil
}

func (t Trim) String() string {
	switch t {
	case TrimSilence:
		return "trim-silence"
	case DoNotTrimSilence:
		return "do-not-trim"
	default:
		return ""
	}
}
