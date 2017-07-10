// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import "encoding/xml"

// The DialClient noun is meant to be used as a Dial.Noun and it specifies a
// client identifier to dial.
//
// You can use up to ten Client nouns within a Dial verb to simultaneously
// attempt a connection with many clients at once. The first client to accept
// the incoming connection is connected to the call and the other connection
// attempts are canceled. If you want to connect with multiple other clients
// simultaneously, read about the Conference noun.
type DialClient struct {
	XMLName              xml.Name            `xml:"Client"`
	ClientName           string              `xml:",chardata"`
	URL                  string              `xml:"url,attr,omitempty"`
	Method               string              `xml:"method,attr,omitempty"`
	StatusCallbackEvent  StatusCallbackEvent `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string              `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string              `xml:"statusCallbackMethod,attr,omitempty"`
}

// The DialConference noun is meant to be used as a Dial.Noun and it allows you
// to connect to a conference room. Much like how the DialNumber noun allows you
// to connect to another phone number, the DialConference noun allows you to connect
// to a named conference room and talk with the other callers who have also
// connected to that room. Conference is commonly used as a container for calls
// when implementing hold, transfer, and barge.
type DialConference struct {
	XMLName                       xml.Name                `xml:"Conference"`
	Name                          string                  `xml:",chardata"`
	Muted                         bool                    `xml:"muted,attr"`
	Beep                          ConfBeep                `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter        ConfStartOnEnterBool    `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit           bool                    `xml:"endConferenceOnExit,attr"`
	WaitURL                       string                  `xml:"waitUrl,attr,omitempty"`
	WaitMethod                    string                  `xml:"waitMethod,attr,omitempty"`
	MaxParticipants               uint16                  `xml:"maxParticipants,attr,omitempty"`
	Record                        ConfRecord              `xml:"record,attr,omitempty"`
	Region                        ConfRegion              `xml:"region,attr,omitempty"`
	Trim                          Trim                    `xml:"trim,attr,omitempty"`
	Whisper                       string                  `xml:"whisper,attr,omitempty"`
	StatusCallbackEvent           ConfStatusCallbackEvent `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback                string                  `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod          string                  `xml:"statusCallbackMethod,attr,omitempty"`
	RecordingStatusCallback       string                  `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string                  `xml:"recordingStatusCallbackMethod,attr,omitempty"`
}

// The DialNumber noun is meant to be used as a Dial.Noun and it specifies a
// phone number to dial. Using the noun's attributes you can specify particular
// behaviors that Twilio should apply when dialing the number.
type DialNumber struct {
	XMLName              xml.Name            `xml:"Number"`
	Number               string              `xml:",chardata"`
	SendDigits           string              `xml:"sendDigits,attr,omitempty"`
	URL                  string              `xml:"url,attr,omitempty"`
	Method               string              `xml:"method,attr,omitempty"`
	StatusCallbackEvent  StatusCallbackEvent `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string              `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string              `xml:"statusCallbackMethod,attr,omitempty"`
}

// The DialQueue noun is meant to be used as a Dial.Noun and it specifies a
// queue to dial. When dialing a queue, the caller will be connected with the
// first enqueued call in the specified queue. If the queue is empty, Dial will
// wait until the next person joins the queue or until the timeout duration is
// reached. If the queue does not exist, Dial will post an error status to its
// URL.
type DialQueue struct {
	XMLName             xml.Name `xml:"Queue"`
	QueueName           string   `xml:",chardata"`
	URL                 string   `xml:"url,attr,omitempty"`
	Method              string   `xml:"method,attr,omitempty"`
	ReservationSID      string   `xml:"reservationSid,attr,omitempty"`
	PostWorkActivitySID string   `xml:"postWorkActivitySid,attr,omitempty"`
}

// The DialSIM noun is meant to be used as a Dial.Noun and it specifies a
// Programmable Wireless SIM to dial.
type DialSIM struct {
	XMLName xml.Name `xml:"Sim"`
	SIM     string   `xml:",chardata"`
}

// The DialSIP noun is meant to be used as a Dial.Noun and it lets you set up
// VoIP sessions by using SIP -- Session Initiation Protocol. With this feature,
// you can send a call to any SIP endpoint.
type DialSIP struct {
	XMLName  xml.Name `xml:"Sip"`
	URI      string   `xml:",chardata"`
	Username string   `xml:"username,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`

	// URL is the call screening URL for the SIP call
	// With Method being the HTTP method used for hitting the URL
	URL    string `xml:"url,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`

	StatusCallbackEvent  StatusCallbackEvent `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string              `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string              `xml:"statusCallbackMethod,attr,omitempty"`

	//
	// Attributes shared from Dial verb
	//
	// Action                        string     `xml:"action,attr,omitempty"`
	// Method                        string     `xml:"method,attr,omitempty"`
	Timeout                       uint       `xml:"timeout,attr,omitempty"`
	HangupOnStar                  bool       `xml:"hangupOnStar,attr"`
	TimeLimit                     uint       `xml:"timeLimit,attr,omitempty"`
	CallerID                      string     `xml:"callerId,attr,omitempty"`
	Record                        DialRecord `xml:"record,attr,omitempty"`
	Trim                          Trim       `xml:"trim,attr,omitempty"`
	RecordingStatusCallback       string     `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string     `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	AnswerOnBridge                bool       `xml:"answerOnBridge,attr"`
	RingTone                      RingTone   `xml:"ringTone,attr,omitempty"`
}
