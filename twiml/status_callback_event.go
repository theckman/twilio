// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"bytes"
	"encoding/xml"
)

// StatusCallbackEvent allows you to specify which events Twilio should webhook
// on. If you'd like to have the callback fire for multiple event types, you can
// use a bitwise-OR to select multiple event types.
type StatusCallbackEvent uint8

const (
	// StatusCallbackInitiated is the event for when a call is started, before
	// it starts to ring.
	StatusCallbackInitiated StatusCallbackEvent = 1 << iota

	// StatusCallbackRinging is the event for when a call starts to ring.
	StatusCallbackRinging

	// StatusCallbackAnswered is the event for when the call is answered.
	StatusCallbackAnswered

	// StatusCallbackCompleted is the event for when the call is finished.
	StatusCallbackCompleted
)

// StatusCallbackAll is a combination of all StatusCallbackEvents, for endpoints
// that want to receive a webhook from Twilio for all call status events.
const StatusCallbackAll = StatusCallbackInitiated | StatusCallbackRinging | StatusCallbackAnswered | StatusCallbackCompleted

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (s StatusCallbackEvent) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: s.String(),
	}

	return attr, nil
}

func (s StatusCallbackEvent) String() string {
	if s == StatusCallbackEvent(0) {
		return ""
	}

	buf := bufferPool.Get().(*bytes.Buffer)

	defer bufferPool.Put(buf)
	defer buf.Reset()

	if s&StatusCallbackInitiated == StatusCallbackInitiated {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("initiated")
	}

	if s&StatusCallbackRinging == StatusCallbackRinging {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("ringing")
	}

	if s&StatusCallbackAnswered == StatusCallbackAnswered {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("answered")
	}

	if s&StatusCallbackCompleted == StatusCallbackCompleted {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("completed")
	}

	return buf.String()
}
