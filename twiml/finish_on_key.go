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

// FinishOnKey is a type for defining which digits will end a recording when
// pressed. This is a bit-flag type, with constants defined for each number.
// You use the bitwise OR operator (|) to enable specific buttons.
type FinishOnKey uint16

const (
	// FinishKeyNumber0 is the equivalent to someone pressing the number 0 on
	// their keypad to end a recording.
	FinishKeyNumber0 FinishOnKey = 1 << iota

	// FinishKeyNumber1 is the equivalent to someone pressing the number 1 on
	// their keypad to end a recording.
	FinishKeyNumber1

	// FinishKeyNumber2 is the equivalent to someone pressing the number 2 on
	// their keypad to end a recording.
	FinishKeyNumber2

	// FinishKeyNumber3 is the equivalent to someone pressing the number 3 on
	// their keypad to end a recording.
	FinishKeyNumber3

	// FinishKeyNumber4 is the equivalent to someone pressing the number 4 on
	// their keypad to end a recording.
	FinishKeyNumber4

	// FinishKeyNumber5 is the equivalent to someone pressing the number 5 on
	// their keypad to end a recording.
	FinishKeyNumber5

	// FinishKeyNumber6 is the equivalent to someone pressing the number 6 on
	// their keypad to end a recording.
	FinishKeyNumber6

	// FinishKeyNumber7 is the equivalent to someone pressing the number 7 on
	// their keypad to end a recording.
	FinishKeyNumber7

	// FinishKeyNumber8 is the equivalent to someone pressing the number 8 on
	// their keypad to end a recording.
	FinishKeyNumber8

	// FinishKeyNumber9 is the equivalent to someone pressing the number 9 on
	// their keypad to end a recording.
	FinishKeyNumber9

	// FinishKeyStar is the equivalent to someone pressing the star (*)
	// button on their keypad to end a recording.
	FinishKeyStar

	// FinishKeyPound is the equivalent to someone pressing the pound (#) button
	// on their keypad to end a recording.
	FinishKeyPound

	// FinishKeyNone is the equivalent of there being no finish key set. This
	// renders an empty string in the XML attribute.
	FinishKeyNone
)

// FinishKeyAll is the combination of all keys together
const FinishKeyAll = FinishKeyNumber0 | FinishKeyNumber1 | FinishKeyNumber2 |
	FinishKeyNumber3 | FinishKeyNumber4 | FinishKeyNumber5 | FinishKeyNumber6 |
	FinishKeyNumber7 | FinishKeyNumber8 | FinishKeyNumber9 | FinishKeyStar |
	FinishKeyPound

// MarshalXMLAttr implements the xml.MarshalerAttr interface.
func (f FinishOnKey) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{
		Name:  name,
		Value: f.String(),
	}

	return attr, nil
}

func (f FinishOnKey) String() string {
	if f == FinishKeyNone {
		return ""
	}

	buf := bufferPool.Get().(*bytes.Buffer)

	defer bufferPool.Put(buf)
	defer buf.Reset()

	if f&FinishKeyNumber1 == FinishKeyNumber1 {
		buf.WriteString("1")
	}

	if f&FinishKeyNumber2 == FinishKeyNumber2 {
		buf.WriteString("2")
	}

	if f&FinishKeyNumber3 == FinishKeyNumber3 {
		buf.WriteString("3")
	}

	if f&FinishKeyNumber4 == FinishKeyNumber4 {
		buf.WriteString("4")
	}

	if f&FinishKeyNumber5 == FinishKeyNumber5 {
		buf.WriteString("5")
	}

	if f&FinishKeyNumber6 == FinishKeyNumber6 {
		buf.WriteString("6")
	}

	if f&FinishKeyNumber7 == FinishKeyNumber7 {
		buf.WriteString("7")
	}

	if f&FinishKeyNumber8 == FinishKeyNumber8 {
		buf.WriteString("8")
	}

	if f&FinishKeyNumber9 == FinishKeyNumber9 {
		buf.WriteString("9")
	}

	if f&FinishKeyNumber0 == FinishKeyNumber0 {
		buf.WriteString("0")
	}

	if f&FinishKeyStar == FinishKeyStar {
		buf.WriteString("*")
	}

	if f&FinishKeyPound == FinishKeyPound {
		buf.WriteString("#")
	}

	return buf.String()
}
