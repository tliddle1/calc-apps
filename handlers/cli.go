package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/tliddle1/calcy-lib/calcy"
)

type CalcHandler struct {
	W          io.Writer
	Calculator calcy.Calculator
}

func New(w io.Writer, c calcy.Calculator) CalcHandler {
	return CalcHandler{
		W:          w,
		Calculator: c,
	}
}

func (this CalcHandler) Handle(args []string) error {
	// Parse the command line arguments into integers
	num1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		return fmt.Errorf("%w: %w", ErrNumberParsing, err1)
	}
	num2, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		return fmt.Errorf("%w: %w", ErrNumberParsing, err2)
	}
	result := this.Calculator.Calculate(num1, num2)
	_, err := fmt.Fprintln(this.W, result)
	return err
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

var ErrNumberParsing = errors.New("error parsing string as number")
