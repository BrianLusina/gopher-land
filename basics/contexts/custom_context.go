package contexts

import (
	"context"
	"time"
)

// detach is an example of a custom contest that implements the Context interface. This will not expire
// and is useful when passing contexts to other functions launched from a goroutine to handle asynchronous functions
// for example, a HTTP request doing some task and offloading the response to a publisher to publish a message to a broker
// This, ideally, should not add latency to the client request and therefore makes sense to launch the publication of the
// message in a separate goroutine passing in the context from the request. However, the context is cancelled when the
// response is returned to the client. This means that the publication of the message to the broker may not happen and this
// causes a subtle bug as no error may be returned or displayed. Using a custom context like this ensures that the async function
// to publish a message to a broker is not cancelled when the parent context is cancelled due to a response being returned to the client
// Using this could be something like:
// publish(detach{ ctx: r.Context() }, message). Where message is the message to publish, while publish is a function whose signature looks like:
// func publish(ctx context.Context, message Message) error { ... }
type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}
