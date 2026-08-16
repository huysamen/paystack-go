package enums

import (
	"encoding/json"
	"testing"
)

func TestChannelUnmarshalJSON_KnownValues(t *testing.T) {
	for _, want := range AllChannels() {
		t.Run(want.String(), func(t *testing.T) {
			var got Channel
			if err := json.Unmarshal([]byte(`"`+want.String()+`"`), &got); err != nil {
				t.Fatalf("unmarshal %q: unexpected error: %v", want, err)
			}
			if got != want {
				t.Fatalf("unmarshal %q: got %q", want, got)
			}
		})
	}
}

func TestChannelUnmarshalJSON_CapitecPay(t *testing.T) {
	var got Channel
	if err := json.Unmarshal([]byte(`"capitec_pay"`), &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != ChannelCapitecPay {
		t.Fatalf("got %q, want %q", got, ChannelCapitecPay)
	}
	if !got.IsValid() {
		t.Fatal("capitec_pay should be a valid channel")
	}
}

// Paystack adds channels without notice, so an unrecognised value must not fail
// the decode of the whole payload it appears in. The raw value is preserved so
// callers can log or persist it.
func TestChannelUnmarshalJSON_UnknownValueIsPreserved(t *testing.T) {
	var got Channel
	if err := json.Unmarshal([]byte(`"some_future_channel"`), &got); err != nil {
		t.Fatalf("unknown channel must not error, got: %v", err)
	}
	if got.String() != "some_future_channel" {
		t.Fatalf("got %q, want raw value preserved", got)
	}
	if got.IsValid() {
		t.Fatal("unknown channel should report IsValid() == false")
	}
}

func TestChannelUnmarshalJSON_EmptyString(t *testing.T) {
	var got Channel
	if err := json.Unmarshal([]byte(`""`), &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Fatalf("got %q, want empty", got)
	}
}

func TestChannelUnmarshalJSON_NonStringStillErrors(t *testing.T) {
	var got Channel
	if err := json.Unmarshal([]byte(`123`), &got); err == nil {
		t.Fatal("expected error for non-string channel")
	}
}

// A charge payload carrying an unknown channel must still decode: this is the
// production failure that motivated the change.
func TestChannelUnmarshalJSON_EmbeddedInStruct(t *testing.T) {
	var payload struct {
		Reference string  `json:"reference"`
		Channel   Channel `json:"channel"`
	}

	if err := json.Unmarshal([]byte(`{"reference":"ref_123","channel":"capitec_pay"}`), &payload); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if payload.Reference != "ref_123" {
		t.Fatalf("reference: got %q", payload.Reference)
	}
	if payload.Channel != ChannelCapitecPay {
		t.Fatalf("channel: got %q", payload.Channel)
	}
}

func TestChannelMarshalJSON_RoundTrip(t *testing.T) {
	for _, c := range append(AllChannels(), Channel("some_future_channel"), Channel("")) {
		b, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("marshal %q: unexpected error: %v", c, err)
		}
		var got Channel
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatalf("round trip %q: unexpected error: %v", c, err)
		}
		if got != c {
			t.Fatalf("round trip: got %q, want %q", got, c)
		}
	}
}

func TestAllChannelsIncludesCapitecPay(t *testing.T) {
	for _, c := range AllChannels() {
		if c == ChannelCapitecPay {
			return
		}
	}
	t.Fatal("AllChannels() must include capitec_pay")
}
