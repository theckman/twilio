// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// GatherInput allows you to define the type of input to gather from a caller.
// The constant values can be bitwise-OR'ed together to support `dtmf speech`
// input mode.
type GatherInput uint8

const (
	// GatherInputDTMF captures user in the form of dual tone multi frequency
	// inputs. This is the user pressing numbers on the keypad.
	GatherInputDTMF GatherInput = 1 << iota

	// GatherInputSpeech enables capturing user input in the form of speech.
	GatherInputSpeech
)

// GatherInputDTMFSpeech is a bitwise-OR of GatherInputDTMF and
// GatherInputSpeech as a convenience.
const GatherInputDTMFSpeech = GatherInputDTMF | GatherInputSpeech

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (g GatherInput) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: g.String(),
	}

	return attr, nil
}

func (g GatherInput) String() string {
	switch {
	case g == GatherInputDTMFSpeech:
		return "dtmf speech"
	case g&GatherInputDTMF == GatherInputDTMF:
		return "dtmf"
	case g&GatherInputSpeech == GatherInputSpeech:
		return "speech"
	default:
		return ""
	}
}
