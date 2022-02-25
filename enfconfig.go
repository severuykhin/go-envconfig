package envconfig

import (
	"context"
	"os"
)

const (
	ErrInvalidEnvvarName  = Error("invalid environment variable name")
	ErrInvalidMapItem     = Error("invalid map item")
	ErrLookuperNil        = Error("lookuper cannot be nil")
	ErrMissingKey         = Error("missing key")
	ErrMissingRequired    = Error("missing required value")
	ErrNoInitNotPtr       = Error("field must be a pointer to have noinit")
	ErrNotPtr             = Error("input must be a pointer")
	ErrNotStruct          = Error("input must be a struct")
	ErrPrefixNotStruct    = Error("prefix is only valid on struct types")
	ErrPrivateField       = Error("cannot parse private fields")
	ErrRequiredAndDefault = Error("field cannot be required and have a default value")
	ErrUnknownOption      = Error("unknown option")
)

// Error is a custom error type for errors returned by envconfig.
type Error string

// Error implements error interface.
func (e Error) Error() string {
	return string(e)
}

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

	if l == nil {
		return ErrLookuperNil
	}

	return nil
}
