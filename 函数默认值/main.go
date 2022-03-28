package main

import "fmt"

const (
	CommonCart = "common"
	BuyNowCart = "buyNow"
)

type dialOptions struct {
	insecure  bool
	transport string
}

type ClientConn struct {
	host string
	opts dialOptions
}

type DialOption interface {
	apply(options *dialOptions)
}

//emptyDialOption kong de
type emptyDialOption struct{}

func (emptyDialOption) apply(options *dialOptions) {}

// useful
type funcDialOption struct {
	f func(options *dialOptions)
}

func (fdo *funcDialOption) apply(do *dialOptions) {
	fdo.f(do)
}

func NewFuncDialOption(f func(options *dialOptions)) DialOption {
	return &funcDialOption{
		f: f,
	}
}

func WithInsecure() DialOption {
	return NewFuncDialOption(func(options *dialOptions) {
		options.insecure = true
	})
}

func WithTransType(kind string) DialOption {
	return NewFuncDialOption(func(options *dialOptions) {
		options.transport = kind
	})
}

func NewConnection(host string, opts ...DialOption) *ClientConn {
	client := &ClientConn{
		host: host,
	}

	for _, opt := range opts {
		opt.apply(&client.opts)
	}
	return client
}
func main() {
	opts := []DialOption{
		WithInsecure(),
		WithTransType("tcp"),
	}

	dd := NewConnection("192", opts...)
	fmt.Println(dd.host, "  ", dd.opts)

	TestPlayer()
}
