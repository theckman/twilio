// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// BargeIn allows you to specify if Twilio should stop playing media from nested
// or verbs once Twilio receives speech or DTMF. Defaults to true.
type BargeIn uint8

const (
	// BargeInTrue sets BargeIn to true.
	BargeInTrue BargeIn = 1 << iota

	// BargeInFalse sets BargeIn to false.
	BargeInFalse
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (b BargeIn) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: b.String(),
	}

	return attr, nil
}

// Bool returns the boolean representation of the BargeIn value. If the value is
// not explicitly false, it's assumed true (to match Twilio's default).
func (b BargeIn) Bool() bool {
	if b == BargeInFalse {
		return false
	}

	return true
}

func (b BargeIn) String() string {
	switch b {
	case BargeInTrue:
		return "true"
	case BargeInFalse:
		return "false"
	default:
		return ""
	}
}
