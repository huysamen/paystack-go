package enums

import (
	"encoding/json"
	"testing"
)

// Every enum in this package models a field Paystack controls and extends
// without notice. Rejecting an unrecognised value during UnmarshalJSON fails
// the decode of the entire payload the field appears in — a webhook event, a
// transaction, a settlement — over one informational string. That is what took
// charge.success webhooks down when capitec_pay shipped.
//
// So: unmarshalling must accept any string and preserve it verbatim. IsValid
// stays strict and is the way callers ask whether a value is one this version
// of the library models.
func TestEnums_AcceptUnknownValues(t *testing.T) {
	const unknown = "some_future_value_paystack_added"

	cases := []struct {
		name   string
		target json.Unmarshaler
	}{
		{"Bearer", new(Bearer)},
		{"Channel", new(Channel)},
		{"Currency", new(Currency)},
		{"DisputeCategory", new(DisputeCategory)},
		{"DisputeResolution", new(DisputeResolution)},
		{"DisputeSource", new(DisputeSource)},
		{"DisputeStatus", new(DisputeStatus)},
		{"Interval", new(Interval)},
		{"MandateAuthorizationStatus", new(MandateAuthorizationStatus)},
		{"MoMo", new(MoMo)},
		{"PageType", new(PageType)},
		{"RefundChannel", new(RefundChannel)},
		{"RefundStatus", new(RefundStatus)},
		{"SettlementStatus", new(SettlementStatus)},
		{"TerminalEventAction", new(TerminalEventAction)},
		{"TerminalEventType", new(TerminalEventType)},
		{"TerminalStatus", new(TerminalStatus)},
		{"TransactionSplitBearerType", new(TransactionSplitBearerType)},
		{"TransactionSplitType", new(TransactionSplitType)},
		{"TransferRecipientType", new(TransferRecipientType)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.target.UnmarshalJSON([]byte(`"` + unknown + `"`)); err != nil {
				t.Fatalf("%s: unknown value must not fail the decode, got: %v", tc.name, err)
			}

			got := stringerValue(t, tc.target)
			if got != unknown {
				t.Fatalf("%s: got %q, want the raw value preserved", tc.name, got)
			}

			if isValid(t, tc.target) {
				t.Fatalf("%s: an unmodelled value must report IsValid() == false", tc.name)
			}
		})
	}
}

// TestEnums_RejectNonStrings: tolerating unknown strings must not turn into
// tolerating malformed JSON. A number or object where a string belongs is a
// decode error, same as before.
func TestEnums_RejectNonStrings(t *testing.T) {
	targets := []struct {
		name   string
		target json.Unmarshaler
	}{
		{"Bearer", new(Bearer)},
		{"Channel", new(Channel)},
		{"Currency", new(Currency)},
		{"Interval", new(Interval)},
		{"SettlementStatus", new(SettlementStatus)},
		{"TransferRecipientType", new(TransferRecipientType)},
	}

	for _, tc := range targets {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.target.UnmarshalJSON([]byte(`{"not":"a string"}`)); err == nil {
				t.Fatalf("%s: expected an error for a non-string value", tc.name)
			}
		})
	}
}

// stringerValue reads the decoded value back through the enum's String method.
func stringerValue(t *testing.T, v any) string {
	t.Helper()
	s, ok := v.(interface{ String() string })
	if !ok {
		t.Fatalf("%T does not implement String()", v)
	}
	return s.String()
}

// isValid calls the enum's IsValid method, which every enum in this package
// exposes on its value (not pointer) receiver.
func isValid(t *testing.T, v any) bool {
	t.Helper()
	iv, ok := v.(interface{ IsValid() bool })
	if !ok {
		t.Fatalf("%T does not implement IsValid()", v)
	}
	return iv.IsValid()
}
