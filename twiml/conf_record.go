// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// ConfRecord lets you record an entire conference.
type ConfRecord uint8

const (
	// ConfDoNotRecord disables recording.
	ConfDoNotRecord ConfRecord = 1 << iota

	// ConfRecordFromStart tells Twilio to record the conference from when the
	// first two participants are bridged. The hold music is never recorded.
	ConfRecordFromStart
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (r ConfRecord) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: r.String(),
	}

	return attr, nil
}

func (r ConfRecord) String() string {
	switch r {
	case ConfDoNotRecord:
		return "do-not-record"
	case ConfRecordFromStart:
		return "record-from-start"
	default:
		return ""
	}
}
