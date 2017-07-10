// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// DialRecord lets you record both legs of a call within the associated Dial verb. Recordings are available in two options: mono-channel or dual-channel.
type DialRecord uint8

const (
	// DialDoNotRecord disables recording.
	DialDoNotRecord DialRecord = 1 << iota

	// DialRecordFromAnswerMono is the mono-channel recording from the call being
	// answered.
	DialRecordFromAnswerMono

	// DialRecordFromRingingMono is the mono-channel recording from the call
	// starting to ring.
	DialRecordFromRingingMono

	// DialRecordFromAnswerDual is the dual-channel recording from the call being
	// answered.
	DialRecordFromAnswerDual

	// DialRecordFromRingingDual is the dual-channel recording from the call
	// starting to ring.
	DialRecordFromRingingDual
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (d DialRecord) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: d.String(),
	}

	return attr, nil
}

func (d DialRecord) String() string {
	switch d {
	case DialDoNotRecord:
		return "do-not-record"
	case DialRecordFromAnswerMono:
		return "record-from-answer"
	case DialRecordFromRingingMono:
		return "record-from-ringing"
	case DialRecordFromAnswerDual:
		return "record-from-answer-dual"
	case DialRecordFromRingingDual:
		return "record-from-ringing-dual"
	default:
		return ""
	}
}
