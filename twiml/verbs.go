// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"encoding/xml"
)

// The Dial verb connects the current caller to another phone. If the called
// party picks up, the two parties are connected and can communicate until one
// hangs up. If the called party does not pick up, if a busy signal is received,
// or if the number doesn't exist, the dial verb will finish.
//
// When the dialed call ends, Twilio makes a GET or POST request to the 'action'
// URL if provided. Call flow will continue using the TwiML received in response
// to that request.
type Dial struct {
	XMLName                       xml.Name   `xml:"Dial"`
	Number                        string     `xml:",chardata"`
	Action                        string     `xml:"action,attr,omitempty"`
	Method                        string     `xml:"method,attr,omitempty"`
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
	Nouns                         []interface{}
}

// The Enqueue verb enqueues the current call in a call queue. Enqueued calls
// wait in hold music until the call is dequeued by another caller via the
// Dial verb or transfered out of the queue via the REST API or the Leave
// verb.
//
// The Enqueue verb will create a queue on demand, if the queue does not already
// exist. The default maximum length of the queue is 100. This can be modified
// using the REST API.
type Enqueue struct {
	XMLName       xml.Name `xml:"Enqueue"`
	QueueName     string   `xml:",chardata"`
	Task          string   `xml:"Task,omitempty"`
	Action        string   `xml:"action,attr,omitempty"`
	Method        string   `xml:"method,attr,omitempty"`
	WaitURL       string   `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod string   `xml:"waitUrlMethod,attr,omitempty"`
	WorkflowSID   string   `xml:"workflowSid,attr,omitempty"`
}

// The Gather verb collects digits or transcribes speech from a caller, when the
// caller is done entering digits or speaking, Twilio submits that data to the
// provided 'action' URL in an HTTP GET or POST request, just like a web browser
// submits data from an HTML form.
type Gather struct {
	XMLName                     xml.Name    `xml:"Gather"`
	Input                       GatherInput `xml:"input,attr,omitempty"`
	Action                      string      `xml:"action,attr,omitempty"`
	Method                      string      `xml:"method,attr,omitempty"`
	Timeout                     uint        `xml:"timeout,attr,omitempty"`
	FinishOnKey                 FinishOnKey `xml:"finishOnKey,attr,omitempty"`
	NumDigits                   uint        `xml:"numDigits,attr,omitempty"`
	PartialResultCallback       string      `xml:"partialResultCallback,attr,omitempty"`
	PartialResultCallbackMethod string      `xml:"partialResultCallbackMethod,attr,omitempty"`
	Language                    Language    `xml:"language,attr,omitempty"`
	Hints                       string      `xml:"hints,attr,omitempty"`
	BargeIn                     BargeIn     `xml:"bargeIn,attr,omitempty"`

	// NestedVerbs within Gather can only contain these three verb types: Say,
	// Play, and Pause. Other types may be rejected by the Twilio TwiML parser.
	NestedVerbs []interface{}
}

// The Hangup verb ends a call. If used as the first verb in a TwiML response it
// does not prevent Twilio from answering the call and billing your account. The
// only way to not answer a call and prevent billing is to use the Reject verb.
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

// The Leave verb transfers control of a call that is in a queue so that the
// caller exits the queue and execution continues with the next verb after the
// original Enqueue.
type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

// The Pause verb waits silently for a specific number of seconds. If Pause is
// the first verb in a TwiML document, Twilio will wait the specified number of
// seconds before picking up the call.
type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  uint     `xml:"length,attr,omitempty"`
}

// Play is play
type Play struct {
	XMLName xml.Name `xml:"Play"`
	URL     string   `xml:",chardata"`
	Loop    uint     `xml:"loop,attr,omitempty"`
	Digits  string   `xml:"digits,attr,omitempty"`
}

// The Record verb records the caller's voice and returns to you the URL of a
// file containing the audio recording. You can optionally generate text
// transcriptions of recorded calls by setting the Transcribe field of the
// Record struct to 'true'.
type Record struct {
	XMLName                       xml.Name    `xml:"Record"`
	Action                        string      `xml:"action,attr,omitempty"`
	Method                        string      `xml:"method,attr,omitempty"`
	Timeout                       uint        `xml:"timeout,attr,omitempty"`
	FinishOnKey                   FinishOnKey `xml:"finishOnKey,attr,omitempty"`
	MaxLength                     uint        `xml:"maxLength,attr,omitempty"`
	PlayBeep                      bool        `xml:"playBeep,attr,omitempty"`
	Trim                          Trim        `xml:"trim,attr,omitempty"`
	RecordingStatusCallback       string      `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string      `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	Transcribe                    bool        `xml:"transcribe,attr"`
	TranscribeCallback            string      `xml:"transcribeCallback,attr,omitempty"`
}

// The Redirect verb transfers control of a call to the TwiML at a different
// URL. All verbs after Redirect are unreachable and ignored.
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	URL     string   `xml:",chardata"`
	Method  string   `xml:"method,attr,omitempty"`
}

// The Reject verb rejects an incoming call to your Twilio number without
// billing you.
type Reject struct {
	XMLName xml.Name     `xml:"Reject"`
	Reason  RejectReason `xml:"reason,attr,omitempty"`
}

// The Say verb converts text to speech that is read back to the caller. Say is
// useful for development or saying dynamic text that is difficult to
// pre-record. The current verb offers different options for voices, each with
// its own supported set of languages and genders, so configure your TwiML
// depending on preferred gender and language combination.
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Message  string   `xml:",chardata"`
	Language Language `xml:"language,attr,omitempty"`
	Loop     uint     `xml:"loop,attr,omitempty"`
	Voice    Voice    `xml:"voice,attr,omitempty"`
}

// The Sms verb sends an SMS message to a phone number during a phone call.
type Sms struct {
	XMLName        xml.Name `xml:"Sms"`
	Message        string   `xml:",chardata"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         string   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
}
