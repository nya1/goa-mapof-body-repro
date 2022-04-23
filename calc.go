package calcapi

import (
	calc "calcsvc/gen/calc"
	"context"
	"fmt"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	s.logger.Print("calc.multiply")

	s.logger.Print("provided query p.Anyobject " + fmt.Sprint(p.Anyobject))
	return
}
