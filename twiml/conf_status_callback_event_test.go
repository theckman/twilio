// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"encoding/xml"
	"testing"
)

func TestConfStatusCallbackEvent_String(t *testing.T) {
	tests := []struct {
		desc string
		in   ConfStatusCallbackEvent
		out  string
	}{
		{"Default (Zero Value) ConfStatusCallbackEvent should return empty-string", ConfStatusCallbackEvent(0), ""},
		{"ConfStatusCallbackEnd should return the end value", ConfStatusCallbackEnd, "end"},
		{"ConfStatusCallbackEnd should return the end value", ConfStatusCallbackEnd, "end"},
		{"ConfStatusCallbackJoin should return the join value", ConfStatusCallbackJoin, "join"},
		{"ConfStatusCallbackLeave should return the leave value", ConfStatusCallbackLeave, "leave"},
		{"ConfStatusCallbackMute should return the mute value", ConfStatusCallbackMute, "mute"},
		{"ConfStatusCallbackHold should return the hold value", ConfStatusCallbackHold, "hold"},
		{"ConfStatusCallbackSpeaker should return the speaker value", ConfStatusCallbackSpeaker, "speaker"},
		{"ConfStatusCallbackStart | ConfStatusCallbackEnd | ConfStatusCallbackJoin should return the start, end, and join values", ConfStatusCallbackStart | ConfStatusCallbackEnd | ConfStatusCallbackJoin, "start end join"},
		{"ConfStatusCallbackAll should return all values", ConfStatusCallbackAll, "start end join leave mute hold speaker"},
	}

	for _, test := range tests {
		if out := test.in.String(); out != test.out {
			t.Errorf(
				"\nDescription: %s\nFinishOnKey(%d).String() = %q; want %q",
				test.desc, test.in, out, test.out,
			)
		}
	}
}

func TestConfStatusCallbackEvent_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       ConfStatusCallbackEvent
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) ConfStatusCallbackEvent should return empty-string", ConfStatusCallbackEvent(0), attrName, attrName, ""},
		{"ConfStatusCallbackEnd should return the end value", ConfStatusCallbackEnd, attrName, attrName, "end"},
		{"ConfStatusCallbackEnd should return the end value", ConfStatusCallbackEnd, attrName, attrName, "end"},
		{"ConfStatusCallbackJoin should return the join value", ConfStatusCallbackJoin, attrName, attrName, "join"},
		{"ConfStatusCallbackLeave should return the leave value", ConfStatusCallbackLeave, attrName, attrName, "leave"},
		{"ConfStatusCallbackMute should return the mute value", ConfStatusCallbackMute, attrName, attrName, "mute"},
		{"ConfStatusCallbackHold should return the hold value", ConfStatusCallbackHold, attrName, attrName, "hold"},
		{"ConfStatusCallbackSpeaker should return the speaker value", ConfStatusCallbackSpeaker, attrName, attrName, "speaker"},
		{"ConfStatusCallbackStart | ConfStatusCallbackEnd | ConfStatusCallbackJoin should return the start, end, and join values", ConfStatusCallbackStart | ConfStatusCallbackEnd | ConfStatusCallbackJoin, attrName, attrName, "start end join"},
		{"ConfStatusCallbackAll should return all values", ConfStatusCallbackAll, attrName, attrName, "start end join leave mute hold speaker"},
	}

	var out xml.Attr
	var err error

	for _, test := range tests {
		out, err = test.in.MarshalXMLAttr(test.inName)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v) Error: %s",
				test.desc, test.in, test.inName, err,
			)
		}

		if out.Name.Space != test.outName.Space || out.Name.Local != test.outName.Local {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v).Name = %#v; want %#v",
				test.desc, test.in, test.inName, out.Name, test.outName,
			)
		}

		if out.Value != test.outValue {
			t.Errorf(
				"\nDescription: %s\nVoice(%d).MarshalAttr(%#v).Value = %q; want %q",
				test.desc, test.in, test.inName, out.Value, test.outValue,
			)
		}
	}
}
