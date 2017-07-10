// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// ConfStartOnEnterBool tells a conference to start when this participant joins
// the conference, if it is not already started. This is true by default. If
// this is false and the participant joins a conference that has not started,
// they are muted and hear background music until a participant joins where
// startConferenceOnEnter is true. This is useful for implementing moderated
// conferences.
type ConfStartOnEnterBool uint8

const (
	// ConfStartOnEnterTrue sets ConfStartOnEnterBool to true.
	ConfStartOnEnterTrue ConfStartOnEnterBool = 1 << iota

	// ConfStartOnEnterFalse sets ConfStartOnEnterBool to false.
	ConfStartOnEnterFalse
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (s ConfStartOnEnterBool) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: s.String(),
	}

	return attr, nil
}

// Bool returns the boolean representation of the ConfStartOnEnterBool value. If the value is
// not explicitly false, it's assumed true (to match Twilio's default).
func (s ConfStartOnEnterBool) Bool() bool {
	if s == ConfStartOnEnterFalse {
		return false
	}

	return true
}

func (s ConfStartOnEnterBool) String() string {
	switch s {
	case ConfStartOnEnterTrue:
		return "true"
	case ConfStartOnEnterFalse:
		return "false"
	default:
		return ""
	}
}
