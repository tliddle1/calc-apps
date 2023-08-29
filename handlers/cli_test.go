package handler

import (
	"bytes"
	"errors"
	"testing"

	"github.com/tliddle1/calcy-lib/calcy"
)

func getTestHandler(calculator calcy.Calculator) CalcHandler {
	myWriter := &DummyWriter{}
	return New(myWriter, calculator)
}

func TestHandler(t *testing.T) {
	myWriter := &DummyWriter{}
	myHandler := New(myWriter, calcy.Addition{})

	err := myHandler.Handle([]string{"1", "2"})
	if err != nil {
		t.Error(err)
	}
	result := myWriter.Written
	if bytes.Compare(result, []byte{51, 10}) != 0 {
		t.Error("Expected a byte array of [51 10] and got", result)
	}
}

func TestHandlerBadInput(t *testing.T) {
	myHandler := getTestHandler(calcy.Addition{})

	err := myHandler.Handle([]string{"f", "2"})
	if err == nil {
		t.Error("Wanted an error but none was returned")
	} else if !errors.Is(err, ErrNumberParsing) {
		t.Error("Wanted error of type ErrNumberParsing but got", err.Error())
	}
}

func TestHandlerBadInput2(t *testing.T) {
	myHandler := getTestHandler(calcy.Addition{})

	err := myHandler.Handle([]string{"1", "c"})
	if err == nil {
		t.Error("Wanted an error but none was returned")
	} else if err.Error() == "strconv.Atoi: parsing \"c\": invalid syntax" {
		//fmt.Println("Error:", err.Error())
	}
}

func TestHandlerWriterError(t *testing.T) {
	myWriter := DummyWriter{}
	myHandler := CalcHandler{
		W:          &myWriter,
		Calculator: calcy.Addition{},
	}

	myWriter.Error = errors.New("test_error")
	err := myHandler.Handle([]string{"1", "c"})
	if err == nil {
		t.Error("Wanted an error but none was returned")
	} else if err.Error() == "test_error" {
		//fmt.Println("Error:", err.Error())
	}
}

type DummyWriter struct {
	Written []byte
	Error   error
}

func (this *DummyWriter) Write(p []byte) (n int, err error) {
	this.Written = p
	return len(p), this.Error
}
