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

func TestDialRecord_String(t *testing.T) {
	tests := []struct {
		desc string
		in   DialRecord
		out  string
	}{
		{"Default (Zero Value) DialRecord should return empty-string", DialRecord(0), ""},
		{"DialDoNotRecord should return do-not-record", DialDoNotRecord, "do-not-record"},
		{"DialRecordFromAnswerMono should return record-from-answer", DialRecordFromAnswerMono, "record-from-answer"},
		{"DialRecordFromRingingMono should return record-from-ringing", DialRecordFromRingingMono, "record-from-ringing"},
		{"DialRecordFromAnswerDual should return record-from-answer-dual", DialRecordFromAnswerDual, "record-from-answer-dual"},
		{"DialRecordFromRingingDual should return record-from-ringing-dual", DialRecordFromRingingDual, "record-from-ringing-dual"},
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

func TestDialRecord_MarshalXMLAttr(t *testing.T) {
	attrName := xml.Name{Local: "trim"}

	tests := []struct {
		desc     string
		in       DialRecord
		inName   xml.Name
		outName  xml.Name
		outValue string
	}{
		{"Default (Zero Value) DialRecord should return empty-string", DialRecord(0), attrName, attrName, ""},
		{"DialDoNotRecord should return do-not-record", DialDoNotRecord, attrName, attrName, "do-not-record"},
		{"DialRecordFromAnswerMono should return record-from-answer", DialRecordFromAnswerMono, attrName, attrName, "record-from-answer"},
		{"DialRecordFromRingingMono should return record-from-ringing", DialRecordFromRingingMono, attrName, attrName, "record-from-ringing"},
		{"DialRecordFromAnswerDual should return record-from-answer-dual", DialRecordFromAnswerDual, attrName, attrName, "record-from-answer-dual"},
		{"DialRecordFromRingingDual should return record-from-ringing-dual", DialRecordFromRingingDual, attrName, attrName, "record-from-ringing-dual"},
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
