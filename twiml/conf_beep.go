// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// ConfBeep allows you to specify if Twilio lets you specify whether a
// notification beep is played to the conference when a participant joins or
// leaves the conference.
type ConfBeep uint8

const (
	// ConfBeepTrue sets ConfBeep to true.
	ConfBeepTrue ConfBeep = 1 << iota

	// ConfBeepFalse sets ConfBeep to false.
	ConfBeepFalse
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (b ConfBeep) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: b.String(),
	}

	return attr, nil
}

// Bool returns the boolean representation of the ConfBeep value. If the value is
// not explicitly false, it's assumed true (to match Twilio's default).
func (b ConfBeep) Bool() bool {
	if b == ConfBeepFalse {
		return false
	}

	return true
}

func (b ConfBeep) String() string {
	switch b {
	case ConfBeepTrue:
		return "true"
	case ConfBeepFalse:
		return "false"
	default:
		return ""
	}
}
