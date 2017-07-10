// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// Voice is the voices that are available as part of the Twilio Text to Speech
// engine using in calls. The default voice is Alice as it has better support
// for languages.
type Voice uint8

const (
	// VoiceDefault is the default vault for which Voice to use. This
	// effectively renders an empty string / no value to have the default of the
	// API be used.
	VoiceDefault Voice = iota

	// VoiceAlice is the default voice, named Alice. It has the best support for
	// languages.
	VoiceAlice

	// VoiceMan is the legacy male voice. It only supports the legacy languages.
	VoiceMan

	// VoiceWoman is the legacy female voice. It only supports the legacy languages.
	VoiceWoman
)

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (v Voice) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: v.String(),
	}

	return attr, nil
}

func (v Voice) String() string {
	switch v {
	case VoiceDefault:
		return ""
	case VoiceAlice:
		return "alice"
	case VoiceMan:
		return "man"
	case VoiceWoman:
		return "woman"
	default:
		return "unknown"
	}
}
