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

// ConfStatusCallbackEvent allows you to specify if Twilio lets you specify whether a
// notification beep is played to the conference when a participant joins or
// leaves the conference.
type ConfStatusCallbackEvent uint16

const (
	// ConfStatusCallbackStart is when the conference has begun and audio is
	// being mixed between all participants. This occurs when there is at least
	// one participant in the conference and a participant with
	// startConferenceOnEnter="true" joins.
	ConfStatusCallbackStart ConfStatusCallbackEvent = 1 << iota

	// ConfStatusCallbackEnd is when the last participant has left the
	// conference or a participant with endConferenceOnExit="true" leaves the
	// conference.
	ConfStatusCallbackEnd

	// ConfStatusCallbackJoin is when a participant has joined the conference.
	ConfStatusCallbackJoin

	// ConfStatusCallbackLeave is when a participant has left the conference.
	ConfStatusCallbackLeave

	// ConfStatusCallbackMute is when a participant has been muted or unmuted.
	ConfStatusCallbackMute

	// ConfStatusCallbackHold is for when a participant has been held or unheld.
	ConfStatusCallbackHold // Hold me closer, Tony Danza

	// ConfStatusCallbackSpeaker is for when a participant has started or
	// stopped speaking.
	ConfStatusCallbackSpeaker
)

// ConfStatusCallbackAll is a constant value that encompasses all ConfStatusCallbackEvent values.
const ConfStatusCallbackAll = ConfStatusCallbackStart | ConfStatusCallbackEnd | ConfStatusCallbackJoin |
	ConfStatusCallbackLeave | ConfStatusCallbackMute | ConfStatusCallbackHold | ConfStatusCallbackSpeaker

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (s ConfStatusCallbackEvent) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: s.String(),
	}

	return attr, nil
}

func (s ConfStatusCallbackEvent) String() string {
	if s == ConfStatusCallbackEvent(0) {
		return ""
	}

	buf := bufferPool.Get().(*bytes.Buffer)

	defer bufferPool.Put(buf)
	defer buf.Reset()

	if s&ConfStatusCallbackStart == ConfStatusCallbackStart {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("start")
	}

	if s&ConfStatusCallbackEnd == ConfStatusCallbackEnd {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("end")
	}

	if s&ConfStatusCallbackJoin == ConfStatusCallbackJoin {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("join")
	}

	if s&ConfStatusCallbackLeave == ConfStatusCallbackLeave {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("leave")
	}

	if s&ConfStatusCallbackMute == ConfStatusCallbackMute {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("mute")
	}

	if s&ConfStatusCallbackHold == ConfStatusCallbackHold {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("hold")
	}

	if s&ConfStatusCallbackSpeaker == ConfStatusCallbackSpeaker {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString("speaker")
	}

	return buf.String()
}
