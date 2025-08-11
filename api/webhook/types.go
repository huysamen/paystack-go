package webhook

import (
	"encoding/json"

	"github.com/huysamen/paystack-go/types"
)

const (
	EventChargeSuccess = "charge.success"

	EventChargeDisputeCreate           = "charge.dispute.create"
	EventChargeDisputeRemind           = "charge.dispute.remind"
	EventChargeDisputeResolve          = "charge.dispute.resolve"
	EventCustomerIdentificationFailed  = "customeridentification.failed"
	EventCustomerIdentificationSuccess = "customeridentification.success"
	EventDedicatedAccountAssignFailed  = "dedicatedaccount.assign.failed"
	EventDedicatedAccountAssignSuccess = "dedicatedaccount.assign.success"
	EventInvoiceCreate                 = "invoice.create"
	EventInvoicePaymentFailed          = "invoice.payment_failed"
	EventInvoiceUpdate                 = "invoice.update"
	EventPaymentRequestPending         = "paymentrequest.pending"
	EventPaymentRequestSuccess         = "paymentrequest.success"
	EventRefundFailed                  = "refund.failed"
	EventRefundPending                 = "refund.pending"
	EventRefundProcessed               = "refund.processed"
	EventRefundProcessing              = "refund.processing"
	EventSubscriptionCreate            = "subscription.create"
	EventSubscriptionDisable           = "subscription.disable"
	EventSubscriptionExpiringCards     = "subscription.expiring_cards"
	EventSubscriptionNotRenew          = "subscription.not_renew"
	EventTransferFailed                = "transfer.failed"
	EventTransferReversed              = "transfer.reversed"
	EventTransferSuccess               = "transfer.success"
)

type Event struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func ParseEventData[T any](event *Event) (*T, error) {
	var data T
	if err := json.Unmarshal(event.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (e *Event) AsChargeSuccess() (*ChargeSuccessEvent, error) {
	return ParseEventData[ChargeSuccessEvent](e)
}

func (e *Event) AsCustomerIdentificationFailed() (*CustomerIdentificationFailedEvent, error) {
	return ParseEventData[CustomerIdentificationFailedEvent](e)
}

func (e *Event) AsCustomerIdentificationSuccess() (*CustomerIdentificationSuccessEvent, error) {
	return ParseEventData[CustomerIdentificationSuccessEvent](e)
}

func (e *Event) AsTransferSuccess() (*TransferSuccessEvent, error) {
	return ParseEventData[TransferSuccessEvent](e)
}

func (e *Event) AsTransferFailed() (*TransferFailedEvent, error) {
	return ParseEventData[TransferFailedEvent](e)
}

func (e *Event) AsTransferReversed() (*TransferReversedEvent, error) {
	return ParseEventData[TransferReversedEvent](e)
}

func (e *Event) AsSubscriptionCreate() (*SubscriptionCreateEvent, error) {
	return ParseEventData[SubscriptionCreateEvent](e)
}

func (e *Event) AsInvoiceCreate() (*InvoiceCreateEvent, error) {
	return ParseEventData[InvoiceCreateEvent](e)
}

func (e *Event) AsRefundProcessed() (*RefundProcessedEvent, error) {
	return ParseEventData[RefundProcessedEvent](e)
}

func (e *Event) AsChargeDispute() (*ChargeDisputeEvent, error) {
	return ParseEventData[ChargeDisputeEvent](e)
}

func (e *Event) AsDedicatedAccount() (*DedicatedAccountEvent, error) {
	return ParseEventData[DedicatedAccountEvent](e)
}

func (e *Event) AsInvoicePaymentFailed() (*InvoicePaymentFailedEvent, error) {
	return ParseEventData[InvoicePaymentFailedEvent](e)
}

func (e *Event) AsInvoiceUpdate() (*InvoiceUpdateEvent, error) {
	return ParseEventData[InvoiceUpdateEvent](e)
}

func (e *Event) AsPaymentRequest() (*PaymentRequestEvent, error) {
	return ParseEventData[PaymentRequestEvent](e)
}

func (e *Event) AsRefundFailed() (*RefundFailedEvent, error) {
	return ParseEventData[RefundFailedEvent](e)
}

func (e *Event) AsRefundPending() (*RefundPendingEvent, error) {
	return ParseEventData[RefundPendingEvent](e)
}

func (e *Event) AsSubscriptionDisable() (*SubscriptionDisableEvent, error) {
	return ParseEventData[SubscriptionDisableEvent](e)
}

func (e *Event) AsSubscriptionNotRenew() (*SubscriptionNotRenewEvent, error) {
	return ParseEventData[SubscriptionNotRenewEvent](e)
}

func (e *Event) AsSubscriptionExpiringCards() (*SubscriptionExpiringCardsEvent, error) {
	// Special-case array payload
	var arr []map[string]any
	if err := json.Unmarshal(e.Data, &arr); err == nil {
		return &SubscriptionExpiringCardsEvent{Entries: types.NewMetadata(map[string]any{"entries": arr})}, nil
	}
	return ParseEventData[SubscriptionExpiringCardsEvent](e)
}
