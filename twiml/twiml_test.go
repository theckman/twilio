// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, you can
// obtain one at https://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2017 Tim Heckman

package twiml

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func readFileString(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func TestEncodeResponse(t *testing.T) {
	simpleSay := &Say{Message: "Testing!"}
	fullSay := &Say{
		Message:  "Testing!",
		Language: LangEnglishUS,
		Loop:     2,
		Voice:    VoiceAlice,
	}
	sliceSimpleSay := []interface{}{simpleSay}
	sliceFullSay := []interface{}{fullSay}

	simpleRecord := &Record{}
	fullRecord := &Record{
		Action:      "https://example.org/action",
		Method:      "POST",
		Timeout:     3,
		FinishOnKey: FinishKeyAll,
		MaxLength:   350,
		PlayBeep:    true,
		Trim:        TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		Transcribe:                    true,
		TranscribeCallback:            "https://example.org/tc",
	}
	sliceSimpleRecord := []interface{}{simpleRecord}
	sliceFullRecord := []interface{}{fullRecord}

	simpleReject := &Reject{}
	fullReject := &Reject{Reason: RejectReasonRejected}
	sliceSimpleReject := []interface{}{simpleReject}
	sliceFullReject := []interface{}{fullReject}

	fullHangup := &Hangup{}
	sliceFullHangup := []interface{}{fullHangup}

	simplePlay := &Play{URL: "https://example.org/audio.mp3"}
	fullPlay := &Play{
		URL:    "https://example.org/audio.mp3",
		Loop:   2,
		Digits: "0w42*",
	}
	sliceSimplePlay := []interface{}{simplePlay}
	sliceFullPlay := []interface{}{fullPlay}

	simplePause := &Pause{}
	fullPause := &Pause{Length: 4}

	sliceSimplePause := []interface{}{simplePause}
	sliceFullPause := []interface{}{fullPause}

	simpleSms := &Sms{Message: "Test message!"}
	fullSms := &Sms{
		Message:        "Test message!",
		To:             "+14155555555",
		From:           "+14155555656",
		Action:         "https://example.org/action",
		Method:         "POST",
		StatusCallback: "https://example.org/scb",
	}

	sliceSimpleSms := []interface{}{simpleSms}
	sliceFullSms := []interface{}{fullSms}

	simpleRedirect := &Redirect{URL: "https://example.org/redirect"}
	fullRedirect := &Redirect{
		URL:    "https://example.org/redirect",
		Method: "POST",
	}

	sliceSimpleRedirect := []interface{}{simpleRedirect}
	sliceFullRedirect := []interface{}{fullRedirect}

	fullLeave := &Leave{}
	sliceFullLeave := []interface{}{fullLeave}

	simpleEnqueue := &Enqueue{QueueName: "test"}
	fullEnqueue := &Enqueue{
		QueueName:     "test",
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	fullEnqueueWithTask := &Enqueue{
		QueueName:     "test",
		Task:          `{"test":"obj"}`,
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	sliceSimpleEnqueue := []interface{}{simpleEnqueue}
	sliceFullEnqueue := []interface{}{fullEnqueue}
	sliceFullEnqueueWithTask := []interface{}{fullEnqueueWithTask}

	simpleGather := &Gather{}
	fullGather := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
	}
	fullGatherWithVerbs := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
		NestedVerbs: []interface{}{
			fullSay, fullPlay, fullPause,
		},
	}

	sliceSimpleGather := []interface{}{simpleGather}
	sliceFullGather := []interface{}{fullGather}
	sliceFullGatherWithVerbs := []interface{}{fullGatherWithVerbs}

	simpleDial := &Dial{Number: "415-555-5555"}
	fullDial := &Dial{
		Number:       "415-555-5555",
		Action:       "https://example.org/action",
		Method:       "POST",
		Timeout:      5,
		HangupOnStar: true,
		TimeLimit:    10,
		CallerID:     "+14155555555",
		Record:       DialRecordFromRingingDual,
		Trim:         TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneUSOld,
	}

	sliceSimpleDial := []interface{}{simpleDial}
	sliceFullDial := []interface{}{fullDial}

	simpleDialClient := &DialClient{ClientName: "Testing"}
	sdcDial := &Dial{Nouns: []interface{}{simpleDialClient}}
	sliceSimpleDialClient := []interface{}{sdcDial}

	fullDialClient := &DialClient{
		ClientName:           "Testing",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdcDial := &Dial{Nouns: []interface{}{fullDialClient}}
	sliceFullDialClient := []interface{}{fdcDial}

	simpleDialQueue := &DialQueue{QueueName: "Testing"}
	sdqDial := &Dial{Nouns: []interface{}{simpleDialQueue}}
	sliceSimpleDialQueue := []interface{}{sdqDial}

	fullDialQueue := &DialQueue{
		QueueName:           "Testing",
		URL:                 "https://example.org/url",
		Method:              "POST",
		ReservationSID:      "reservationSid",
		PostWorkActivitySID: "postWorkActivitySid",
	}
	fdqDial := &Dial{Nouns: []interface{}{fullDialQueue}}
	sliceFullDialQueue := []interface{}{fdqDial}

	fullDialSIM := &DialSIM{SIM: "Testing"}
	fdsDial := &Dial{Nouns: []interface{}{fullDialSIM}}
	sliceFullDialSIM := []interface{}{fdsDial}

	simpleDialNumber := &DialNumber{Number: "+14155555555"}
	sdnDial := &Dial{Nouns: []interface{}{simpleDialNumber}}
	sliceSimpleDialNumber := []interface{}{sdnDial}

	fullDialNumber := &DialNumber{
		Number:               "+14155555555",
		SendDigits:           "ww42",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdnDial := &Dial{Nouns: []interface{}{fullDialNumber}}
	sliceFullDialNumber := []interface{}{fdnDial}

	simpleDialConference := &DialConference{Name: "testConf"}
	sdconfDial := &Dial{Nouns: []interface{}{simpleDialConference}}
	sliceSimpleDialConference := []interface{}{sdconfDial}

	fullDialConference := &DialConference{
		Name:  "testConf",
		Muted: true,
		Beep:  ConfBeepTrue,
		StartConferenceOnEnter:        ConfStartOnEnterTrue,
		EndConferenceOnExit:           true,
		WaitURL:                       "https://example.org/wait",
		WaitMethod:                    "POST",
		MaxParticipants:               42, // because Twilio doesn't allow tree-fiddy
		Record:                        ConfRecordFromStart,
		Region:                        ConfRegionJapan,
		Trim:                          TrimSilence,
		Whisper:                       "testWhisper",
		StatusCallbackEvent:           ConfStatusCallbackAll,
		StatusCallbackMethod:          "POST",
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
	}
	fdconfDial := &Dial{Nouns: []interface{}{fullDialConference}}
	sliceFullDialConference := []interface{}{fdconfDial}

	simpleDialSIP := &DialSIP{URI: "Testing"}
	sdsipDial := &Dial{Nouns: []interface{}{simpleDialSIP}}
	sliceSimpleDialSIP := []interface{}{sdsipDial}

	fullDialSIP := &DialSIP{
		URI:                  "Testing",
		Username:             "testUser",
		Password:             "testPass",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
		Timeout:              42,
		HangupOnStar:         true,
		TimeLimit:            84,
		CallerID:             "theckman",
		Record:               DialRecordFromRingingDual,
		Trim:                 TrimSilence,
		RecordingStatusCallback:       "https://example.org/rscb",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneJapan,
	}
	fdsipDial := &Dial{Nouns: []interface{}{fullDialSIP}}
	sliceFullDialSIP := []interface{}{fdsipDial}

	dialWithNouns := &Dial{Nouns: []interface{}{fullDialSIP, fullDialQueue}}
	sliceDialWithNouns := []interface{}{dialWithNouns}

	anotherSimpleSay := &Say{Message: "Goodbye!"}
	anotherFullSay := &Say{
		Message:  "Goodbye!",
		Language: LangEnglishAustralia,
		Loop:     3,
		Voice:    VoiceAlice,
	}

	sliceSimple := []interface{}{
		simpleSay, simpleRecord, simpleReject, fullHangup,
		simplePlay, simplePause, simpleSms, simpleRedirect,
		fullLeave, simpleEnqueue, simpleGather, simpleDial,
		anotherSimpleSay,
	}

	sliceFull := []interface{}{
		fullSay, fullRecord, fullReject, fullHangup,
		fullPlay, fullPause, fullSms, fullRedirect,
		fullLeave, fullEnqueue, fullGather, fullDial,
		anotherFullSay,
	}

	//
	// Test loop
	//
	tests := []struct {
		desc        string
		in          *Response
		outfilePath string
	}{
		{"Response with one simple <Say> instruction", &Response{Verbs: sliceSimpleSay}, "simplesay.xml"},
		{"Response with one full <Say> instruction", &Response{Verbs: sliceFullSay}, "fullsay.xml"},
		{"Response with one simple <Record> instruction", &Response{Verbs: sliceSimpleRecord}, "simplerecord.xml"},
		{"Response with one full <Record> instruction", &Response{Verbs: sliceFullRecord}, "fullrecord.xml"},
		{"Response with one simple <Reject> instruction", &Response{Verbs: sliceSimpleReject}, "simplereject.xml"},
		{"Response with one full <Reject> instruction", &Response{Verbs: sliceFullReject}, "fullreject.xml"},
		{"Response with one full <Hangup> instruction", &Response{Verbs: sliceFullHangup}, "fullhangup.xml"},
		{"Response with one simple <Play> instruction", &Response{Verbs: sliceSimplePlay}, "simpleplay.xml"},
		{"Response with one full <Play> instruction", &Response{Verbs: sliceFullPlay}, "fullplay.xml"},
		{"Response with one simple <Pause> instruction", &Response{Verbs: sliceSimplePause}, "simplepause.xml"},
		{"Response with one full <Pause> instruction", &Response{Verbs: sliceFullPause}, "fullpause.xml"},
		{"Response with one simple <Sms> instruction", &Response{Verbs: sliceSimpleSms}, "simplesms.xml"},
		{"Response with one full <Sms> instruction", &Response{Verbs: sliceFullSms}, "fullsms.xml"},
		{"Response with one simple <Redirect> instruction", &Response{Verbs: sliceSimpleRedirect}, "simpleredirect.xml"},
		{"Response with one full <Redirect> instruction", &Response{Verbs: sliceFullRedirect}, "fullredirect.xml"},
		{"Response with one full <Leave> instruction", &Response{Verbs: sliceFullLeave}, "fullleave.xml"},
		{"Response with one simple <Enqueue> instruction", &Response{Verbs: sliceSimpleEnqueue}, "simpleenqueue.xml"},
		{"Response with one full <Enqueue> instruction", &Response{Verbs: sliceFullEnqueue}, "fullenqueue.xml"},
		{"Response with one full <Enqueue> instruction with a Task", &Response{Verbs: sliceFullEnqueueWithTask}, "fullenqueuewithtask.xml"},
		{"Response with one simple <Gather> instruction", &Response{Verbs: sliceSimpleGather}, "simplegather.xml"},
		{"Response with one full <Gather> instruction", &Response{Verbs: sliceFullGather}, "fullgather.xml"},
		{"Response with one full <Gather> instruction with verbs", &Response{Verbs: sliceFullGatherWithVerbs}, "fullgatherwithverbs.xml"},
		{"Response with one simple <Dial> instruction", &Response{Verbs: sliceSimpleDial}, "simpledial.xml"},
		{"Response with one full <Dial> instruction", &Response{Verbs: sliceFullDial}, "fulldial.xml"},
		{"Response with one simple <Dial><Client> instruction", &Response{Verbs: sliceSimpleDialClient}, "simpleDialClient.xml"},
		{"Response with one full <Dial><Client> instruction", &Response{Verbs: sliceFullDialClient}, "fullDialClient.xml"},
		{"Response with one simple <Dial><Queue> instruction", &Response{Verbs: sliceSimpleDialQueue}, "simpleDialQueue.xml"},
		{"Response with one full <Dial><Queue> instruction", &Response{Verbs: sliceFullDialQueue}, "fullDialQueue.xml"},
		{"Response with one full <Dial><Sim> instruction", &Response{Verbs: sliceFullDialSIM}, "fullDialSIM.xml"},
		{"Response with one simple <Dial><Number> instruction", &Response{Verbs: sliceSimpleDialNumber}, "simpleDialNumber.xml"},
		{"Response with one full <Dial><Number> instruction", &Response{Verbs: sliceFullDialNumber}, "fullDialNumber.xml"},
		{"Response with one simple <Dial><Conference> instruction", &Response{Verbs: sliceSimpleDialConference}, "simpleDialConference.xml"},
		{"Response with one full <Dial><Conference> instruction", &Response{Verbs: sliceFullDialConference}, "fullDialConference.xml"},
		{"Response with one simple <Dial><Sip> instruction", &Response{Verbs: sliceSimpleDialSIP}, "simpleDialSIP.xml"},
		{"Response with one full <Dial><Sip> instruction", &Response{Verbs: sliceFullDialSIP}, "fullDialSIP.xml"},
		{"Response with one <Dial> with multiple nouns instruction", &Response{Verbs: sliceDialWithNouns}, "fullDialWithNouns.xml"},
		{"Response with all simple instructions", &Response{Verbs: sliceSimple}, "simple.xml"},
		{"Response with all full instructions", &Response{Verbs: sliceFull}, "full.xml"},
	}

	for _, test := range tests {
		tdPath := filepath.Join("testdata", test.outfilePath)
		testExpectedOut, err := readFileString(tdPath)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nUnexpected error reading testdata file (%s): %s",
				test.desc, tdPath, err.Error(),
			)
			continue
		}

		b := &bytes.Buffer{}

		err = EncodeResponse(b, test.in)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nEncodeResponse() Unexpected Error: %s",
				test.desc, err,
			)
			continue
		}

		if out := b.String(); out != testExpectedOut {
			t.Errorf(
				"\nDescription: %s\nRendered XML (quoted with `):\n`%s`\n\nWant XML (quoted with `):\n`%s`",
				test.desc, out, testExpectedOut,
			)
			continue
		}
	}
}

func TestMarshalResponse(t *testing.T) {
	simpleSay := &Say{Message: "Testing!"}
	fullSay := &Say{
		Message:  "Testing!",
		Language: LangEnglishUS,
		Loop:     2,
		Voice:    VoiceAlice,
	}
	sliceSimpleSay := []interface{}{simpleSay}
	sliceFullSay := []interface{}{fullSay}

	simpleRecord := &Record{}
	fullRecord := &Record{
		Action:      "https://example.org/action",
		Method:      "POST",
		Timeout:     3,
		FinishOnKey: FinishKeyAll,
		MaxLength:   350,
		PlayBeep:    true,
		Trim:        TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		Transcribe:                    true,
		TranscribeCallback:            "https://example.org/tc",
	}
	sliceSimpleRecord := []interface{}{simpleRecord}
	sliceFullRecord := []interface{}{fullRecord}

	simpleReject := &Reject{}
	fullReject := &Reject{Reason: RejectReasonRejected}
	sliceSimpleReject := []interface{}{simpleReject}
	sliceFullReject := []interface{}{fullReject}

	fullHangup := &Hangup{}
	sliceFullHangup := []interface{}{fullHangup}

	simplePlay := &Play{URL: "https://example.org/audio.mp3"}
	fullPlay := &Play{
		URL:    "https://example.org/audio.mp3",
		Loop:   2,
		Digits: "0w42*",
	}
	sliceSimplePlay := []interface{}{simplePlay}
	sliceFullPlay := []interface{}{fullPlay}

	simplePause := &Pause{}
	fullPause := &Pause{Length: 4}

	sliceSimplePause := []interface{}{simplePause}
	sliceFullPause := []interface{}{fullPause}

	simpleSms := &Sms{Message: "Test message!"}
	fullSms := &Sms{
		Message:        "Test message!",
		To:             "+14155555555",
		From:           "+14155555656",
		Action:         "https://example.org/action",
		Method:         "POST",
		StatusCallback: "https://example.org/scb",
	}

	sliceSimpleSms := []interface{}{simpleSms}
	sliceFullSms := []interface{}{fullSms}

	simpleRedirect := &Redirect{URL: "https://example.org/redirect"}
	fullRedirect := &Redirect{
		URL:    "https://example.org/redirect",
		Method: "POST",
	}

	sliceSimpleRedirect := []interface{}{simpleRedirect}
	sliceFullRedirect := []interface{}{fullRedirect}

	fullLeave := &Leave{}
	sliceFullLeave := []interface{}{fullLeave}

	simpleEnqueue := &Enqueue{QueueName: "test"}
	fullEnqueue := &Enqueue{
		QueueName:     "test",
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	fullEnqueueWithTask := &Enqueue{
		QueueName:     "test",
		Task:          `{"test":"obj"}`,
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	sliceSimpleEnqueue := []interface{}{simpleEnqueue}
	sliceFullEnqueue := []interface{}{fullEnqueue}
	sliceFullEnqueueWithTask := []interface{}{fullEnqueueWithTask}

	simpleGather := &Gather{}
	fullGather := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
	}
	fullGatherWithVerbs := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
		NestedVerbs: []interface{}{
			fullSay, fullPlay, fullPause,
		},
	}

	sliceSimpleGather := []interface{}{simpleGather}
	sliceFullGather := []interface{}{fullGather}
	sliceFullGatherWithVerbs := []interface{}{fullGatherWithVerbs}

	simpleDial := &Dial{Number: "415-555-5555"}
	fullDial := &Dial{
		Number:       "415-555-5555",
		Action:       "https://example.org/action",
		Method:       "POST",
		Timeout:      5,
		HangupOnStar: true,
		TimeLimit:    10,
		CallerID:     "+14155555555",
		Record:       DialRecordFromRingingDual,
		Trim:         TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneUSOld,
	}

	sliceSimpleDial := []interface{}{simpleDial}
	sliceFullDial := []interface{}{fullDial}

	simpleDialClient := &DialClient{ClientName: "Testing"}
	sdcDial := &Dial{Nouns: []interface{}{simpleDialClient}}
	sliceSimpleDialClient := []interface{}{sdcDial}

	fullDialClient := &DialClient{
		ClientName:           "Testing",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdcDial := &Dial{Nouns: []interface{}{fullDialClient}}
	sliceFullDialClient := []interface{}{fdcDial}

	simpleDialQueue := &DialQueue{QueueName: "Testing"}
	sdqDial := &Dial{Nouns: []interface{}{simpleDialQueue}}
	sliceSimpleDialQueue := []interface{}{sdqDial}

	fullDialQueue := &DialQueue{
		QueueName:           "Testing",
		URL:                 "https://example.org/url",
		Method:              "POST",
		ReservationSID:      "reservationSid",
		PostWorkActivitySID: "postWorkActivitySid",
	}
	fdqDial := &Dial{Nouns: []interface{}{fullDialQueue}}
	sliceFullDialQueue := []interface{}{fdqDial}

	fullDialSIM := &DialSIM{SIM: "Testing"}
	fdsDial := &Dial{Nouns: []interface{}{fullDialSIM}}
	sliceFullDialSIM := []interface{}{fdsDial}

	simpleDialNumber := &DialNumber{Number: "+14155555555"}
	sdnDial := &Dial{Nouns: []interface{}{simpleDialNumber}}
	sliceSimpleDialNumber := []interface{}{sdnDial}

	fullDialNumber := &DialNumber{
		Number:               "+14155555555",
		SendDigits:           "ww42",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdnDial := &Dial{Nouns: []interface{}{fullDialNumber}}
	sliceFullDialNumber := []interface{}{fdnDial}

	simpleDialConference := &DialConference{Name: "testConf"}
	sdconfDial := &Dial{Nouns: []interface{}{simpleDialConference}}
	sliceSimpleDialConference := []interface{}{sdconfDial}

	fullDialConference := &DialConference{
		Name:  "testConf",
		Muted: true,
		Beep:  ConfBeepTrue,
		StartConferenceOnEnter:        ConfStartOnEnterTrue,
		EndConferenceOnExit:           true,
		WaitURL:                       "https://example.org/wait",
		WaitMethod:                    "POST",
		MaxParticipants:               42, // because Twilio doesn't allow tree-fiddy
		Record:                        ConfRecordFromStart,
		Region:                        ConfRegionJapan,
		Trim:                          TrimSilence,
		Whisper:                       "testWhisper",
		StatusCallbackEvent:           ConfStatusCallbackAll,
		StatusCallbackMethod:          "POST",
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
	}
	fdconfDial := &Dial{Nouns: []interface{}{fullDialConference}}
	sliceFullDialConference := []interface{}{fdconfDial}

	simpleDialSIP := &DialSIP{URI: "Testing"}
	sdsipDial := &Dial{Nouns: []interface{}{simpleDialSIP}}
	sliceSimpleDialSIP := []interface{}{sdsipDial}

	fullDialSIP := &DialSIP{
		URI:                  "Testing",
		Username:             "testUser",
		Password:             "testPass",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
		Timeout:              42,
		HangupOnStar:         true,
		TimeLimit:            84,
		CallerID:             "theckman",
		Record:               DialRecordFromRingingDual,
		Trim:                 TrimSilence,
		RecordingStatusCallback:       "https://example.org/rscb",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneJapan,
	}
	fdsipDial := &Dial{Nouns: []interface{}{fullDialSIP}}
	sliceFullDialSIP := []interface{}{fdsipDial}

	dialWithNouns := &Dial{Nouns: []interface{}{fullDialSIP, fullDialQueue}}
	sliceDialWithNouns := []interface{}{dialWithNouns}

	anotherSimpleSay := &Say{Message: "Goodbye!"}
	anotherFullSay := &Say{
		Message:  "Goodbye!",
		Language: LangEnglishAustralia,
		Loop:     3,
		Voice:    VoiceAlice,
	}

	sliceSimple := []interface{}{
		simpleSay, simpleRecord, simpleReject, fullHangup,
		simplePlay, simplePause, simpleSms, simpleRedirect,
		fullLeave, simpleEnqueue, simpleGather, simpleDial,
		anotherSimpleSay,
	}

	sliceFull := []interface{}{
		fullSay, fullRecord, fullReject, fullHangup,
		fullPlay, fullPause, fullSms, fullRedirect,
		fullLeave, fullEnqueue, fullGather, fullDial,
		anotherFullSay,
	}

	//
	// Test loop
	//
	tests := []struct {
		desc        string
		in          *Response
		outfilePath string
	}{
		{"Response with one simple <Say> instruction", &Response{Verbs: sliceSimpleSay}, "simplesay.xml"},
		{"Response with one full <Say> instruction", &Response{Verbs: sliceFullSay}, "fullsay.xml"},
		{"Response with one simple <Record> instruction", &Response{Verbs: sliceSimpleRecord}, "simplerecord.xml"},
		{"Response with one full <Record> instruction", &Response{Verbs: sliceFullRecord}, "fullrecord.xml"},
		{"Response with one simple <Reject> instruction", &Response{Verbs: sliceSimpleReject}, "simplereject.xml"},
		{"Response with one full <Reject> instruction", &Response{Verbs: sliceFullReject}, "fullreject.xml"},
		{"Response with one full <Hangup> instruction", &Response{Verbs: sliceFullHangup}, "fullhangup.xml"},
		{"Response with one simple <Play> instruction", &Response{Verbs: sliceSimplePlay}, "simpleplay.xml"},
		{"Response with one full <Play> instruction", &Response{Verbs: sliceFullPlay}, "fullplay.xml"},
		{"Response with one simple <Pause> instruction", &Response{Verbs: sliceSimplePause}, "simplepause.xml"},
		{"Response with one full <Pause> instruction", &Response{Verbs: sliceFullPause}, "fullpause.xml"},
		{"Response with one simple <Sms> instruction", &Response{Verbs: sliceSimpleSms}, "simplesms.xml"},
		{"Response with one full <Sms> instruction", &Response{Verbs: sliceFullSms}, "fullsms.xml"},
		{"Response with one simple <Redirect> instruction", &Response{Verbs: sliceSimpleRedirect}, "simpleredirect.xml"},
		{"Response with one full <Redirect> instruction", &Response{Verbs: sliceFullRedirect}, "fullredirect.xml"},
		{"Response with one full <Leave> instruction", &Response{Verbs: sliceFullLeave}, "fullleave.xml"},
		{"Response with one simple <Enqueue> instruction", &Response{Verbs: sliceSimpleEnqueue}, "simpleenqueue.xml"},
		{"Response with one full <Enqueue> instruction", &Response{Verbs: sliceFullEnqueue}, "fullenqueue.xml"},
		{"Response with one full <Enqueue> instruction with a Task", &Response{Verbs: sliceFullEnqueueWithTask}, "fullenqueuewithtask.xml"},
		{"Response with one simple <Gather> instruction", &Response{Verbs: sliceSimpleGather}, "simplegather.xml"},
		{"Response with one full <Gather> instruction", &Response{Verbs: sliceFullGather}, "fullgather.xml"},
		{"Response with one full <Gather> instruction with verbs", &Response{Verbs: sliceFullGatherWithVerbs}, "fullgatherwithverbs.xml"},
		{"Response with one simple <Dial> instruction", &Response{Verbs: sliceSimpleDial}, "simpledial.xml"},
		{"Response with one full <Dial> instruction", &Response{Verbs: sliceFullDial}, "fulldial.xml"},
		{"Response with one simple <Dial><Client> instruction", &Response{Verbs: sliceSimpleDialClient}, "simpleDialClient.xml"},
		{"Response with one full <Dial><Client> instruction", &Response{Verbs: sliceFullDialClient}, "fullDialClient.xml"},
		{"Response with one simple <Dial><Queue> instruction", &Response{Verbs: sliceSimpleDialQueue}, "simpleDialQueue.xml"},
		{"Response with one full <Dial><Queue> instruction", &Response{Verbs: sliceFullDialQueue}, "fullDialQueue.xml"},
		{"Response with one full <Dial><Sim> instruction", &Response{Verbs: sliceFullDialSIM}, "fullDialSIM.xml"},
		{"Response with one simple <Dial><Number> instruction", &Response{Verbs: sliceSimpleDialNumber}, "simpleDialNumber.xml"},
		{"Response with one full <Dial><Number> instruction", &Response{Verbs: sliceFullDialNumber}, "fullDialNumber.xml"},
		{"Response with one simple <Dial><Conference> instruction", &Response{Verbs: sliceSimpleDialConference}, "simpleDialConference.xml"},
		{"Response with one full <Dial><Conference> instruction", &Response{Verbs: sliceFullDialConference}, "fullDialConference.xml"},
		{"Response with one simple <Dial><Sip> instruction", &Response{Verbs: sliceSimpleDialSIP}, "simpleDialSIP.xml"},
		{"Response with one full <Dial><Sip> instruction", &Response{Verbs: sliceFullDialSIP}, "fullDialSIP.xml"},
		{"Response with one <Dial> with multiple nouns instruction", &Response{Verbs: sliceDialWithNouns}, "fullDialWithNouns.xml"},
		{"Response with all simple instructions", &Response{Verbs: sliceSimple}, "simple.xml"},
		{"Response with all full instructions", &Response{Verbs: sliceFull}, "full.xml"},
	}

	for _, test := range tests {
		tdPath := filepath.Join("testdata", test.outfilePath)
		testExpectedOut, err := readFileString(tdPath)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nUnexpected error reading testdata file (%s): %s",
				test.desc, tdPath, err.Error(),
			)
			continue
		}

		renderedXML, err := MarshalResponse(test.in)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nEncodeResponse() Unexpected Error: %s",
				test.desc, err,
			)
			continue
		}

		if out := string(renderedXML); out != testExpectedOut {
			t.Errorf(
				"\nDescription: %s\nRendered XML (quoted with `):\n`%s`\n\nWant XML (quoted with `):\n`%s`",
				test.desc, out, testExpectedOut,
			)
			continue
		}
	}
}

//
// DIVIDE
//

func TestEncodeSlice(t *testing.T) {
	simpleSay := &Say{Message: "Testing!"}
	fullSay := &Say{
		Message:  "Testing!",
		Language: LangEnglishUS,
		Loop:     2,
		Voice:    VoiceAlice,
	}
	sliceSimpleSay := []interface{}{simpleSay}
	sliceFullSay := []interface{}{fullSay}

	simpleRecord := &Record{}
	fullRecord := &Record{
		Action:      "https://example.org/action",
		Method:      "POST",
		Timeout:     3,
		FinishOnKey: FinishKeyAll,
		MaxLength:   350,
		PlayBeep:    true,
		Trim:        TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		Transcribe:                    true,
		TranscribeCallback:            "https://example.org/tc",
	}
	sliceSimpleRecord := []interface{}{simpleRecord}
	sliceFullRecord := []interface{}{fullRecord}

	simpleReject := &Reject{}
	fullReject := &Reject{Reason: RejectReasonRejected}
	sliceSimpleReject := []interface{}{simpleReject}
	sliceFullReject := []interface{}{fullReject}

	fullHangup := &Hangup{}
	sliceFullHangup := []interface{}{fullHangup}

	simplePlay := &Play{URL: "https://example.org/audio.mp3"}
	fullPlay := &Play{
		URL:    "https://example.org/audio.mp3",
		Loop:   2,
		Digits: "0w42*",
	}
	sliceSimplePlay := []interface{}{simplePlay}
	sliceFullPlay := []interface{}{fullPlay}

	simplePause := &Pause{}
	fullPause := &Pause{Length: 4}

	sliceSimplePause := []interface{}{simplePause}
	sliceFullPause := []interface{}{fullPause}

	simpleSms := &Sms{Message: "Test message!"}
	fullSms := &Sms{
		Message:        "Test message!",
		To:             "+14155555555",
		From:           "+14155555656",
		Action:         "https://example.org/action",
		Method:         "POST",
		StatusCallback: "https://example.org/scb",
	}

	sliceSimpleSms := []interface{}{simpleSms}
	sliceFullSms := []interface{}{fullSms}

	simpleRedirect := &Redirect{URL: "https://example.org/redirect"}
	fullRedirect := &Redirect{
		URL:    "https://example.org/redirect",
		Method: "POST",
	}

	sliceSimpleRedirect := []interface{}{simpleRedirect}
	sliceFullRedirect := []interface{}{fullRedirect}

	fullLeave := &Leave{}
	sliceFullLeave := []interface{}{fullLeave}

	simpleEnqueue := &Enqueue{QueueName: "test"}
	fullEnqueue := &Enqueue{
		QueueName:     "test",
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	fullEnqueueWithTask := &Enqueue{
		QueueName:     "test",
		Task:          `{"test":"obj"}`,
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	sliceSimpleEnqueue := []interface{}{simpleEnqueue}
	sliceFullEnqueue := []interface{}{fullEnqueue}
	sliceFullEnqueueWithTask := []interface{}{fullEnqueueWithTask}

	simpleGather := &Gather{}
	fullGather := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
	}
	fullGatherWithVerbs := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
		NestedVerbs: []interface{}{
			fullSay, fullPlay, fullPause,
		},
	}

	sliceSimpleGather := []interface{}{simpleGather}
	sliceFullGather := []interface{}{fullGather}
	sliceFullGatherWithVerbs := []interface{}{fullGatherWithVerbs}

	simpleDial := &Dial{Number: "415-555-5555"}
	fullDial := &Dial{
		Number:       "415-555-5555",
		Action:       "https://example.org/action",
		Method:       "POST",
		Timeout:      5,
		HangupOnStar: true,
		TimeLimit:    10,
		CallerID:     "+14155555555",
		Record:       DialRecordFromRingingDual,
		Trim:         TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneUSOld,
	}

	sliceSimpleDial := []interface{}{simpleDial}
	sliceFullDial := []interface{}{fullDial}

	simpleDialClient := &DialClient{ClientName: "Testing"}
	sdcDial := &Dial{Nouns: []interface{}{simpleDialClient}}
	sliceSimpleDialClient := []interface{}{sdcDial}

	fullDialClient := &DialClient{
		ClientName:           "Testing",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdcDial := &Dial{Nouns: []interface{}{fullDialClient}}
	sliceFullDialClient := []interface{}{fdcDial}

	simpleDialQueue := &DialQueue{QueueName: "Testing"}
	sdqDial := &Dial{Nouns: []interface{}{simpleDialQueue}}
	sliceSimpleDialQueue := []interface{}{sdqDial}

	fullDialQueue := &DialQueue{
		QueueName:           "Testing",
		URL:                 "https://example.org/url",
		Method:              "POST",
		ReservationSID:      "reservationSid",
		PostWorkActivitySID: "postWorkActivitySid",
	}
	fdqDial := &Dial{Nouns: []interface{}{fullDialQueue}}
	sliceFullDialQueue := []interface{}{fdqDial}

	fullDialSIM := &DialSIM{SIM: "Testing"}
	fdsDial := &Dial{Nouns: []interface{}{fullDialSIM}}
	sliceFullDialSIM := []interface{}{fdsDial}

	simpleDialNumber := &DialNumber{Number: "+14155555555"}
	sdnDial := &Dial{Nouns: []interface{}{simpleDialNumber}}
	sliceSimpleDialNumber := []interface{}{sdnDial}

	fullDialNumber := &DialNumber{
		Number:               "+14155555555",
		SendDigits:           "ww42",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdnDial := &Dial{Nouns: []interface{}{fullDialNumber}}
	sliceFullDialNumber := []interface{}{fdnDial}

	simpleDialConference := &DialConference{Name: "testConf"}
	sdconfDial := &Dial{Nouns: []interface{}{simpleDialConference}}
	sliceSimpleDialConference := []interface{}{sdconfDial}

	fullDialConference := &DialConference{
		Name:  "testConf",
		Muted: true,
		Beep:  ConfBeepTrue,
		StartConferenceOnEnter:        ConfStartOnEnterTrue,
		EndConferenceOnExit:           true,
		WaitURL:                       "https://example.org/wait",
		WaitMethod:                    "POST",
		MaxParticipants:               42, // because Twilio doesn't allow tree-fiddy
		Record:                        ConfRecordFromStart,
		Region:                        ConfRegionJapan,
		Trim:                          TrimSilence,
		Whisper:                       "testWhisper",
		StatusCallbackEvent:           ConfStatusCallbackAll,
		StatusCallbackMethod:          "POST",
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
	}
	fdconfDial := &Dial{Nouns: []interface{}{fullDialConference}}
	sliceFullDialConference := []interface{}{fdconfDial}

	simpleDialSIP := &DialSIP{URI: "Testing"}
	sdsipDial := &Dial{Nouns: []interface{}{simpleDialSIP}}
	sliceSimpleDialSIP := []interface{}{sdsipDial}

	fullDialSIP := &DialSIP{
		URI:                  "Testing",
		Username:             "testUser",
		Password:             "testPass",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
		Timeout:              42,
		HangupOnStar:         true,
		TimeLimit:            84,
		CallerID:             "theckman",
		Record:               DialRecordFromRingingDual,
		Trim:                 TrimSilence,
		RecordingStatusCallback:       "https://example.org/rscb",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneJapan,
	}
	fdsipDial := &Dial{Nouns: []interface{}{fullDialSIP}}
	sliceFullDialSIP := []interface{}{fdsipDial}

	dialWithNouns := &Dial{Nouns: []interface{}{fullDialSIP, fullDialQueue}}
	sliceDialWithNouns := []interface{}{dialWithNouns}

	anotherSimpleSay := &Say{Message: "Goodbye!"}
	anotherFullSay := &Say{
		Message:  "Goodbye!",
		Language: LangEnglishAustralia,
		Loop:     3,
		Voice:    VoiceAlice,
	}

	sliceSimple := []interface{}{
		simpleSay, simpleRecord, simpleReject, fullHangup,
		simplePlay, simplePause, simpleSms, simpleRedirect,
		fullLeave, simpleEnqueue, simpleGather, simpleDial,
		anotherSimpleSay,
	}

	sliceFull := []interface{}{
		fullSay, fullRecord, fullReject, fullHangup,
		fullPlay, fullPause, fullSms, fullRedirect,
		fullLeave, fullEnqueue, fullGather, fullDial,
		anotherFullSay,
	}

	//
	// Test loop
	//
	tests := []struct {
		desc        string
		in          []interface{}
		outfilePath string
	}{
		{"Response with one simple <Say> instruction", sliceSimpleSay, "simplesay.xml"},
		{"Response with one full <Say> instruction", sliceFullSay, "fullsay.xml"},
		{"Response with one simple <Record> instruction", sliceSimpleRecord, "simplerecord.xml"},
		{"Response with one full <Record> instruction", sliceFullRecord, "fullrecord.xml"},
		{"Response with one simple <Reject> instruction", sliceSimpleReject, "simplereject.xml"},
		{"Response with one full <Reject> instruction", sliceFullReject, "fullreject.xml"},
		{"Response with one full <Hangup> instruction", sliceFullHangup, "fullhangup.xml"},
		{"Response with one simple <Play> instruction", sliceSimplePlay, "simpleplay.xml"},
		{"Response with one full <Play> instruction", sliceFullPlay, "fullplay.xml"},
		{"Response with one simple <Pause> instruction", sliceSimplePause, "simplepause.xml"},
		{"Response with one full <Pause> instruction", sliceFullPause, "fullpause.xml"},
		{"Response with one simple <Sms> instruction", sliceSimpleSms, "simplesms.xml"},
		{"Response with one full <Sms> instruction", sliceFullSms, "fullsms.xml"},
		{"Response with one simple <Redirect> instruction", sliceSimpleRedirect, "simpleredirect.xml"},
		{"Response with one full <Redirect> instruction", sliceFullRedirect, "fullredirect.xml"},
		{"Response with one full <Leave> instruction", sliceFullLeave, "fullleave.xml"},
		{"Response with one simple <Enqueue> instruction", sliceSimpleEnqueue, "simpleenqueue.xml"},
		{"Response with one full <Enqueue> instruction", sliceFullEnqueue, "fullenqueue.xml"},
		{"Response with one full <Enqueue> instruction with a Task", sliceFullEnqueueWithTask, "fullenqueuewithtask.xml"},
		{"Response with one simple <Gather> instruction", sliceSimpleGather, "simplegather.xml"},
		{"Response with one full <Gather> instruction", sliceFullGather, "fullgather.xml"},
		{"Response with one full <Gather> instruction with verbs", sliceFullGatherWithVerbs, "fullgatherwithverbs.xml"},
		{"Response with one simple <Dial> instruction", sliceSimpleDial, "simpledial.xml"},
		{"Response with one full <Dial> instruction", sliceFullDial, "fulldial.xml"},
		{"Response with one simple <Dial><Client> instruction", sliceSimpleDialClient, "simpleDialClient.xml"},
		{"Response with one full <Dial><Client> instruction", sliceFullDialClient, "fullDialClient.xml"},
		{"Response with one simple <Dial><Queue> instruction", sliceSimpleDialQueue, "simpleDialQueue.xml"},
		{"Response with one full <Dial><Queue> instruction", sliceFullDialQueue, "fullDialQueue.xml"},
		{"Response with one full <Dial><Sim> instruction", sliceFullDialSIM, "fullDialSIM.xml"},
		{"Response with one simple <Dial><Number> instruction", sliceSimpleDialNumber, "simpleDialNumber.xml"},
		{"Response with one full <Dial><Number> instruction", sliceFullDialNumber, "fullDialNumber.xml"},
		{"Response with one simple <Dial><Conference> instruction", sliceSimpleDialConference, "simpleDialConference.xml"},
		{"Response with one full <Dial><Conference> instruction", sliceFullDialConference, "fullDialConference.xml"},
		{"Response with one simple <Dial><Sip> instruction", sliceSimpleDialSIP, "simpleDialSIP.xml"},
		{"Response with one full <Dial><Sip> instruction", sliceFullDialSIP, "fullDialSIP.xml"},
		{"Response with one <Dial> with multiple nouns instruction", sliceDialWithNouns, "fullDialWithNouns.xml"},
		{"Response with all simple instructions", sliceSimple, "simple.xml"},
		{"Response with all full instructions", sliceFull, "full.xml"},
	}

	for _, test := range tests {
		tdPath := filepath.Join("testdata", test.outfilePath)
		testExpectedOut, err := readFileString(tdPath)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nUnexpected error reading testdata file (%s): %s",
				test.desc, tdPath, err.Error(),
			)
			continue
		}

		b := &bytes.Buffer{}

		err = EncodeSlice(b, test.in)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nEncodeResponse() Unexpected Error: %s",
				test.desc, err,
			)
			continue
		}

		if out := b.String(); out != testExpectedOut {
			t.Errorf(
				"\nDescription: %s\nRendered XML (quoted with `):\n`%s`\n\nWant XML (quoted with `):\n`%s`",
				test.desc, out, testExpectedOut,
			)
			continue
		}
	}
}

func TestMarshalSlice(t *testing.T) {
	simpleSay := &Say{Message: "Testing!"}
	fullSay := &Say{
		Message:  "Testing!",
		Language: LangEnglishUS,
		Loop:     2,
		Voice:    VoiceAlice,
	}
	sliceSimpleSay := []interface{}{simpleSay}
	sliceFullSay := []interface{}{fullSay}

	simpleRecord := &Record{}
	fullRecord := &Record{
		Action:      "https://example.org/action",
		Method:      "POST",
		Timeout:     3,
		FinishOnKey: FinishKeyAll,
		MaxLength:   350,
		PlayBeep:    true,
		Trim:        TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		Transcribe:                    true,
		TranscribeCallback:            "https://example.org/tc",
	}
	sliceSimpleRecord := []interface{}{simpleRecord}
	sliceFullRecord := []interface{}{fullRecord}

	simpleReject := &Reject{}
	fullReject := &Reject{Reason: RejectReasonRejected}
	sliceSimpleReject := []interface{}{simpleReject}
	sliceFullReject := []interface{}{fullReject}

	fullHangup := &Hangup{}
	sliceFullHangup := []interface{}{fullHangup}

	simplePlay := &Play{URL: "https://example.org/audio.mp3"}
	fullPlay := &Play{
		URL:    "https://example.org/audio.mp3",
		Loop:   2,
		Digits: "0w42*",
	}
	sliceSimplePlay := []interface{}{simplePlay}
	sliceFullPlay := []interface{}{fullPlay}

	simplePause := &Pause{}
	fullPause := &Pause{Length: 4}

	sliceSimplePause := []interface{}{simplePause}
	sliceFullPause := []interface{}{fullPause}

	simpleSms := &Sms{Message: "Test message!"}
	fullSms := &Sms{
		Message:        "Test message!",
		To:             "+14155555555",
		From:           "+14155555656",
		Action:         "https://example.org/action",
		Method:         "POST",
		StatusCallback: "https://example.org/scb",
	}

	sliceSimpleSms := []interface{}{simpleSms}
	sliceFullSms := []interface{}{fullSms}

	simpleRedirect := &Redirect{URL: "https://example.org/redirect"}
	fullRedirect := &Redirect{
		URL:    "https://example.org/redirect",
		Method: "POST",
	}

	sliceSimpleRedirect := []interface{}{simpleRedirect}
	sliceFullRedirect := []interface{}{fullRedirect}

	fullLeave := &Leave{}
	sliceFullLeave := []interface{}{fullLeave}

	simpleEnqueue := &Enqueue{QueueName: "test"}
	fullEnqueue := &Enqueue{
		QueueName:     "test",
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	fullEnqueueWithTask := &Enqueue{
		QueueName:     "test",
		Task:          `{"test":"obj"}`,
		Action:        "https://example.org/action",
		Method:        "POST",
		WaitURL:       "https://example.org/wait",
		WaitURLMethod: "GET",
		WorkflowSID:   "WWtesting",
	}

	sliceSimpleEnqueue := []interface{}{simpleEnqueue}
	sliceFullEnqueue := []interface{}{fullEnqueue}
	sliceFullEnqueueWithTask := []interface{}{fullEnqueueWithTask}

	simpleGather := &Gather{}
	fullGather := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
	}
	fullGatherWithVerbs := &Gather{
		Input:                       GatherInputDTMFSpeech,
		Action:                      "https://example.org/action",
		Method:                      "POST",
		Timeout:                     5,
		FinishOnKey:                 FinishKeyStar | FinishKeyPound,
		NumDigits:                   42,
		PartialResultCallback:       "https://example.org/prc",
		PartialResultCallbackMethod: "POST",
		Language:                    LangEnglishUS,
		Hints:                       "bacon ipsum, other stuff",
		BargeIn:                     BargeInFalse,
		NestedVerbs: []interface{}{
			fullSay, fullPlay, fullPause,
		},
	}

	sliceSimpleGather := []interface{}{simpleGather}
	sliceFullGather := []interface{}{fullGather}
	sliceFullGatherWithVerbs := []interface{}{fullGatherWithVerbs}

	simpleDial := &Dial{Number: "415-555-5555"}
	fullDial := &Dial{
		Number:       "415-555-5555",
		Action:       "https://example.org/action",
		Method:       "POST",
		Timeout:      5,
		HangupOnStar: true,
		TimeLimit:    10,
		CallerID:     "+14155555555",
		Record:       DialRecordFromRingingDual,
		Trim:         TrimSilence,
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneUSOld,
	}

	sliceSimpleDial := []interface{}{simpleDial}
	sliceFullDial := []interface{}{fullDial}

	simpleDialClient := &DialClient{ClientName: "Testing"}
	sdcDial := &Dial{Nouns: []interface{}{simpleDialClient}}
	sliceSimpleDialClient := []interface{}{sdcDial}

	fullDialClient := &DialClient{
		ClientName:           "Testing",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdcDial := &Dial{Nouns: []interface{}{fullDialClient}}
	sliceFullDialClient := []interface{}{fdcDial}

	simpleDialQueue := &DialQueue{QueueName: "Testing"}
	sdqDial := &Dial{Nouns: []interface{}{simpleDialQueue}}
	sliceSimpleDialQueue := []interface{}{sdqDial}

	fullDialQueue := &DialQueue{
		QueueName:           "Testing",
		URL:                 "https://example.org/url",
		Method:              "POST",
		ReservationSID:      "reservationSid",
		PostWorkActivitySID: "postWorkActivitySid",
	}
	fdqDial := &Dial{Nouns: []interface{}{fullDialQueue}}
	sliceFullDialQueue := []interface{}{fdqDial}

	fullDialSIM := &DialSIM{SIM: "Testing"}
	fdsDial := &Dial{Nouns: []interface{}{fullDialSIM}}
	sliceFullDialSIM := []interface{}{fdsDial}

	simpleDialNumber := &DialNumber{Number: "+14155555555"}
	sdnDial := &Dial{Nouns: []interface{}{simpleDialNumber}}
	sliceSimpleDialNumber := []interface{}{sdnDial}

	fullDialNumber := &DialNumber{
		Number:               "+14155555555",
		SendDigits:           "ww42",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
	}
	fdnDial := &Dial{Nouns: []interface{}{fullDialNumber}}
	sliceFullDialNumber := []interface{}{fdnDial}

	simpleDialConference := &DialConference{Name: "testConf"}
	sdconfDial := &Dial{Nouns: []interface{}{simpleDialConference}}
	sliceSimpleDialConference := []interface{}{sdconfDial}

	fullDialConference := &DialConference{
		Name:  "testConf",
		Muted: true,
		Beep:  ConfBeepTrue,
		StartConferenceOnEnter:        ConfStartOnEnterTrue,
		EndConferenceOnExit:           true,
		WaitURL:                       "https://example.org/wait",
		WaitMethod:                    "POST",
		MaxParticipants:               42, // because Twilio doesn't allow tree-fiddy
		Record:                        ConfRecordFromStart,
		Region:                        ConfRegionJapan,
		Trim:                          TrimSilence,
		Whisper:                       "testWhisper",
		StatusCallbackEvent:           ConfStatusCallbackAll,
		StatusCallbackMethod:          "POST",
		RecordingStatusCallback:       "https://example.org/rsc",
		RecordingStatusCallbackMethod: "POST",
	}
	fdconfDial := &Dial{Nouns: []interface{}{fullDialConference}}
	sliceFullDialConference := []interface{}{fdconfDial}

	simpleDialSIP := &DialSIP{URI: "Testing"}
	sdsipDial := &Dial{Nouns: []interface{}{simpleDialSIP}}
	sliceSimpleDialSIP := []interface{}{sdsipDial}

	fullDialSIP := &DialSIP{
		URI:                  "Testing",
		Username:             "testUser",
		Password:             "testPass",
		URL:                  "https://example.org/url",
		Method:               "POST",
		StatusCallbackEvent:  StatusCallbackAll,
		StatusCallback:       "https://example.org/scb",
		StatusCallbackMethod: "POST",
		Timeout:              42,
		HangupOnStar:         true,
		TimeLimit:            84,
		CallerID:             "theckman",
		Record:               DialRecordFromRingingDual,
		Trim:                 TrimSilence,
		RecordingStatusCallback:       "https://example.org/rscb",
		RecordingStatusCallbackMethod: "POST",
		AnswerOnBridge:                true,
		RingTone:                      RingToneJapan,
	}
	fdsipDial := &Dial{Nouns: []interface{}{fullDialSIP}}
	sliceFullDialSIP := []interface{}{fdsipDial}

	dialWithNouns := &Dial{Nouns: []interface{}{fullDialSIP, fullDialQueue}}
	sliceDialWithNouns := []interface{}{dialWithNouns}

	anotherSimpleSay := &Say{Message: "Goodbye!"}
	anotherFullSay := &Say{
		Message:  "Goodbye!",
		Language: LangEnglishAustralia,
		Loop:     3,
		Voice:    VoiceAlice,
	}

	sliceSimple := []interface{}{
		simpleSay, simpleRecord, simpleReject, fullHangup,
		simplePlay, simplePause, simpleSms, simpleRedirect,
		fullLeave, simpleEnqueue, simpleGather, simpleDial,
		anotherSimpleSay,
	}

	sliceFull := []interface{}{
		fullSay, fullRecord, fullReject, fullHangup,
		fullPlay, fullPause, fullSms, fullRedirect,
		fullLeave, fullEnqueue, fullGather, fullDial,
		anotherFullSay,
	}

	//
	// Test loop
	//
	tests := []struct {
		desc        string
		in          []interface{}
		outfilePath string
	}{
		{"Response with one simple <Say> instruction", sliceSimpleSay, "simplesay.xml"},
		{"Response with one full <Say> instruction", sliceFullSay, "fullsay.xml"},
		{"Response with one simple <Record> instruction", sliceSimpleRecord, "simplerecord.xml"},
		{"Response with one full <Record> instruction", sliceFullRecord, "fullrecord.xml"},
		{"Response with one simple <Reject> instruction", sliceSimpleReject, "simplereject.xml"},
		{"Response with one full <Reject> instruction", sliceFullReject, "fullreject.xml"},
		{"Response with one full <Hangup> instruction", sliceFullHangup, "fullhangup.xml"},
		{"Response with one simple <Play> instruction", sliceSimplePlay, "simpleplay.xml"},
		{"Response with one full <Play> instruction", sliceFullPlay, "fullplay.xml"},
		{"Response with one simple <Pause> instruction", sliceSimplePause, "simplepause.xml"},
		{"Response with one full <Pause> instruction", sliceFullPause, "fullpause.xml"},
		{"Response with one simple <Sms> instruction", sliceSimpleSms, "simplesms.xml"},
		{"Response with one full <Sms> instruction", sliceFullSms, "fullsms.xml"},
		{"Response with one simple <Redirect> instruction", sliceSimpleRedirect, "simpleredirect.xml"},
		{"Response with one full <Redirect> instruction", sliceFullRedirect, "fullredirect.xml"},
		{"Response with one full <Leave> instruction", sliceFullLeave, "fullleave.xml"},
		{"Response with one simple <Enqueue> instruction", sliceSimpleEnqueue, "simpleenqueue.xml"},
		{"Response with one full <Enqueue> instruction", sliceFullEnqueue, "fullenqueue.xml"},
		{"Response with one full <Enqueue> instruction with a Task", sliceFullEnqueueWithTask, "fullenqueuewithtask.xml"},
		{"Response with one simple <Gather> instruction", sliceSimpleGather, "simplegather.xml"},
		{"Response with one full <Gather> instruction", sliceFullGather, "fullgather.xml"},
		{"Response with one full <Gather> instruction with verbs", sliceFullGatherWithVerbs, "fullgatherwithverbs.xml"},
		{"Response with one simple <Dial> instruction", sliceSimpleDial, "simpledial.xml"},
		{"Response with one full <Dial> instruction", sliceFullDial, "fulldial.xml"},
		{"Response with one simple <Dial><Client> instruction", sliceSimpleDialClient, "simpleDialClient.xml"},
		{"Response with one full <Dial><Client> instruction", sliceFullDialClient, "fullDialClient.xml"},
		{"Response with one simple <Dial><Queue> instruction", sliceSimpleDialQueue, "simpleDialQueue.xml"},
		{"Response with one full <Dial><Queue> instruction", sliceFullDialQueue, "fullDialQueue.xml"},
		{"Response with one full <Dial><Sim> instruction", sliceFullDialSIM, "fullDialSIM.xml"},
		{"Response with one simple <Dial><Number> instruction", sliceSimpleDialNumber, "simpleDialNumber.xml"},
		{"Response with one full <Dial><Number> instruction", sliceFullDialNumber, "fullDialNumber.xml"},
		{"Response with one simple <Dial><Conference> instruction", sliceSimpleDialConference, "simpleDialConference.xml"},
		{"Response with one full <Dial><Conference> instruction", sliceFullDialConference, "fullDialConference.xml"},
		{"Response with one simple <Dial><Sip> instruction", sliceSimpleDialSIP, "simpleDialSIP.xml"},
		{"Response with one full <Dial><Sip> instruction", sliceFullDialSIP, "fullDialSIP.xml"},
		{"Response with one <Dial> with multiple nouns instruction", sliceDialWithNouns, "fullDialWithNouns.xml"},
		{"Response with all simple instructions", sliceSimple, "simple.xml"},
		{"Response with all full instructions", sliceFull, "full.xml"},
	}

	for _, test := range tests {
		tdPath := filepath.Join("testdata", test.outfilePath)
		testExpectedOut, err := readFileString(tdPath)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nUnexpected error reading testdata file (%s): %s",
				test.desc, tdPath, err.Error(),
			)
			continue
		}

		renderedXML, err := MarshalSlice(test.in)

		if err != nil {
			t.Errorf(
				"\nDescription: %s\nEncodeResponse() Unexpected Error: %s",
				test.desc, err,
			)
			continue
		}

		if out := string(renderedXML); out != testExpectedOut {
			t.Errorf(
				"\nDescription: %s\nRendered XML (quoted with `):\n`%s`\n\nWant XML (quoted with `):\n`%s`",
				test.desc, out, testExpectedOut,
			)
			continue
		}
	}
}
