package webhook

import (
	"encoding/json"
	"fmt"
)

// Webhook event constants for all supported Paystack webhook events.
// These constants help prevent typos and make event handling more maintainable.
// Source: https://paystack.com/docs/payments/webhooks/ (as of 2025-07-31)
const (
	// Charge events
	EventChargeSuccess = "charge.success"

	// Dispute events
	EventChargeDisputeCreate  = "charge.dispute.create"
	EventChargeDisputeRemind  = "charge.dispute.remind"
	EventChargeDisputeResolve = "charge.dispute.resolve"

	// Customer identification events
	EventCustomerIdentificationFailed  = "customeridentification.failed"
	EventCustomerIdentificationSuccess = "customeridentification.success"

	// Dedicated Virtual Account (DVA) events
	EventDedicatedAccountAssignFailed  = "dedicatedaccount.assign.failed"
	EventDedicatedAccountAssignSuccess = "dedicatedaccount.assign.success"

	// Invoice events
	EventInvoiceCreate        = "invoice.create"
	EventInvoicePaymentFailed = "invoice.payment_failed"
	EventInvoiceUpdate        = "invoice.update"

	// Payment request events
	EventPaymentRequestPending = "paymentrequest.pending"
	EventPaymentRequestSuccess = "paymentrequest.success"

	// Refund events
	EventRefundFailed     = "refund.failed"
	EventRefundPending    = "refund.pending"
	EventRefundProcessed  = "refund.processed"
	EventRefundProcessing = "refund.processing"

	// Subscription events
	EventSubscriptionCreate        = "subscription.create"
	EventSubscriptionDisable       = "subscription.disable"
	EventSubscriptionExpiringCards = "subscription.expiring_cards"
	EventSubscriptionNotRenew      = "subscription.not_renew"

	// Transfer events
	EventTransferFailed   = "transfer.failed"
	EventTransferReversed = "transfer.reversed"
	EventTransferSuccess  = "transfer.success"
)

// Event represents a Paystack webhook event
type Event struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

// ParseEventData parses the webhook event data into the specified type
func ParseEventData[T any](event *Event) (*T, error) {
	var data T
	if err := json.Unmarshal(event.Data, &data); err != nil {
		return nil, fmt.Errorf("failed to parse event data: %w", err)
	}
	return &data, nil
}

// Convenience methods for parsing specific event types

// AsChargeSuccess parses the event data as ChargeSuccessEvent
func (e *Event) AsChargeSuccess() (*ChargeSuccessEvent, error) {
	return ParseEventData[ChargeSuccessEvent](e)
}

// AsCustomerIdentificationFailed parses the event data as CustomerIdentificationFailedEvent
func (e *Event) AsCustomerIdentificationFailed() (*CustomerIdentificationFailedEvent, error) {
	return ParseEventData[CustomerIdentificationFailedEvent](e)
}

// AsCustomerIdentificationSuccess parses the event data as CustomerIdentificationSuccessEvent
func (e *Event) AsCustomerIdentificationSuccess() (*CustomerIdentificationSuccessEvent, error) {
	return ParseEventData[CustomerIdentificationSuccessEvent](e)
}

// AsTransferSuccess parses the event data as TransferSuccessEvent
func (e *Event) AsTransferSuccess() (*TransferSuccessEvent, error) {
	return ParseEventData[TransferSuccessEvent](e)
}

// AsTransferFailed parses the event data as TransferFailedEvent
func (e *Event) AsTransferFailed() (*TransferFailedEvent, error) {
	return ParseEventData[TransferFailedEvent](e)
}

// AsTransferReversed parses the event data as TransferReversedEvent
func (e *Event) AsTransferReversed() (*TransferReversedEvent, error) {
	return ParseEventData[TransferReversedEvent](e)
}

// AsSubscriptionCreate parses the event data as SubscriptionCreateEvent
func (e *Event) AsSubscriptionCreate() (*SubscriptionCreateEvent, error) {
	return ParseEventData[SubscriptionCreateEvent](e)
}

// AsInvoiceCreate parses the event data as InvoiceCreateEvent
func (e *Event) AsInvoiceCreate() (*InvoiceCreateEvent, error) {
	return ParseEventData[InvoiceCreateEvent](e)
}

// AsRefundProcessed parses the event data as RefundProcessedEvent
func (e *Event) AsRefundProcessed() (*RefundProcessedEvent, error) {
	return ParseEventData[RefundProcessedEvent](e)
}

// AsChargeDispute parses the event data as ChargeDisputeEvent
func (e *Event) AsChargeDispute() (*ChargeDisputeEvent, error) {
	return ParseEventData[ChargeDisputeEvent](e)
}

// AsDedicatedAccount parses the event data as DedicatedAccountEvent
func (e *Event) AsDedicatedAccount() (*DedicatedAccountEvent, error) {
	return ParseEventData[DedicatedAccountEvent](e)
}

// AsInvoicePaymentFailed parses the event data as InvoicePaymentFailedEvent
func (e *Event) AsInvoicePaymentFailed() (*InvoicePaymentFailedEvent, error) {
	return ParseEventData[InvoicePaymentFailedEvent](e)
}

// AsInvoiceUpdate parses the event data as InvoiceUpdateEvent
func (e *Event) AsInvoiceUpdate() (*InvoiceUpdateEvent, error) {
	return ParseEventData[InvoiceUpdateEvent](e)
}

// AsPaymentRequest parses the event data as PaymentRequestEvent
func (e *Event) AsPaymentRequest() (*PaymentRequestEvent, error) {
	return ParseEventData[PaymentRequestEvent](e)
}

// AsRefundFailed parses the event data as RefundFailedEvent
func (e *Event) AsRefundFailed() (*RefundFailedEvent, error) {
	return ParseEventData[RefundFailedEvent](e)
}

// AsRefundPending parses the event data as RefundPendingEvent
func (e *Event) AsRefundPending() (*RefundPendingEvent, error) {
	return ParseEventData[RefundPendingEvent](e)
}

// AsSubscriptionDisable parses the event data as SubscriptionDisableEvent
func (e *Event) AsSubscriptionDisable() (*SubscriptionDisableEvent, error) {
	return ParseEventData[SubscriptionDisableEvent](e)
}

// AsSubscriptionNotRenew parses the event data as SubscriptionNotRenewEvent
func (e *Event) AsSubscriptionNotRenew() (*SubscriptionNotRenewEvent, error) {
	return ParseEventData[SubscriptionNotRenewEvent](e)
}

// AsSubscriptionExpiringCards parses the event data as SubscriptionExpiringCardsEvent
func (e *Event) AsSubscriptionExpiringCards() (*SubscriptionExpiringCardsEvent, error) {
	return ParseEventData[SubscriptionExpiringCardsEvent](e)
}
