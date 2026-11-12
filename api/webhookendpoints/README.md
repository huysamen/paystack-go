# Webhook Endpoints API

Send webhook events to more than one URL, and choose the events each URL receives. Each endpoint has its own name, URL, events and status; Paystack sends each event to every enabled endpoint subscribed to it (and, separately, to the webhook URL set on the dashboard).

> **Pilot.** Webhook endpoints are in pilot. An integration outside it gets `403 Forbidden` on every call. Request access from support@paystack.com.

Test and live mode have separate endpoints: the test secret key manages test endpoints, the live key live ones. A restricted key needs `webhooks:read` for `EventsCatalog`, `List` and `Fetch`, and `webhooks:write` for the rest.

## Available Operations

- **EventsCatalog** - Events an endpoint can subscribe to, by category
- **List** - Endpoints on the integration, without their events (cursor pages)
- **Fetch** - One endpoint with its events
- **Create** - Create an endpoint
- **Update** - Change name, URL, events or status
- **Delete** - Remove an endpoint
- **Ping** - Send a signed test event and report the answer

## Quick Examples

### Create an endpoint

```go
import "github.com/huysamen/paystack-go/api/webhookendpoints"

req := webhookendpoints.NewCreateRequestBuilder(
    "Order fulfillment",
    "https://example.com/webhooks/paystack",
    []string{"charge.success", "charge.failed"},
)

rsp, err := client.WebhookEndpoints.Create(ctx, *req)
if err != nil {
    return err
}
if !rsp.IsSuccess() {
    return rsp.Err() // 409: another endpoint has this URL
}

code := rsp.Data.WebhookURLCode.String() // store it: WHK_...
```

### List every endpoint

```go
builder := webhookendpoints.NewListRequestBuilder().Limit(50)

for {
    rsp, err := client.WebhookEndpoints.List(ctx, *builder)
    if err != nil {
        return err
    }

    for _, ep := range rsp.Data {
        fmt.Println(ep.WebhookURLCode, ep.WebhookURL, ep.Enabled)
    }

    if rsp.Meta == nil || !rsp.Meta.Next.Valid {
        break
    }

    builder = webhookendpoints.NewListRequestBuilder().Limit(50).Next(rsp.Meta.Next.String())
}
```

### Fetch with events

```go
rsp, err := client.WebhookEndpoints.Fetch(ctx, "WHK_1a2b3c4d")
// rsp.Data.Events is grouped by category; EventNames() flattens and sorts it.
events := rsp.Data.EventNames()
```

### Move an endpoint, change its events, pause it

```go
req := webhookendpoints.NewUpdateRequestBuilder().
    URL("https://new.example.com/webhooks/paystack").
    Events([]string{"charge.success"}) // replaces the whole list

rsp, err := client.WebhookEndpoints.Update(ctx, "WHK_1a2b3c4d", *req)

// Pause without losing URL or events; Enabled(true) resumes.
rsp, err = client.WebhookEndpoints.Update(ctx, "WHK_1a2b3c4d", *webhookendpoints.NewUpdateRequestBuilder().Enabled(false))
```

### Ping

```go
rsp, err := client.WebhookEndpoints.Ping(ctx, "WHK_1a2b3c4d")
// rsp.Data.OK: the URL answered 2xx. rsp.Data.HTTPStatus is null when Paystack
// could not reach it. Limit: 10 pings every 60 seconds per mode (429 beyond).
```

### Delete

```go
rsp, err := client.WebhookEndpoints.Delete(ctx, "WHK_1a2b3c4d")
```

## Notes

- URLs must be https, publicly reachable, up to 2,048 characters, with no username or password, and unique within the environment (scheme and host compared ignoring case).
- `[]string{"*"}` subscribes to every event in the catalog at that moment; events Paystack adds later need another `Update` with `"*"`. `"*"` can't be combined with named events.
- Create and Update answer `events` as a flat list; Fetch answers them grouped by category.
