package envconfig

import (
	"context"
	"os"
)

// Process processes the struct using the environment. See ProcessWith for a
// more customizable version.
func Process(ctx context.Context, i interface{}) error {
	return ProcessWith(ctx, i, OsParser())
}

type Lookuper interface {
	Lookup(key string) (string, bool)
}

type osLookuper struct{}

func (o *osLookuper) Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

func OsParser() Lookuper {
	return new(osLookuper)
}

func ProcessWith(ctx context.Context, i interface{}, l Lookuper) error {
	return nil
}
