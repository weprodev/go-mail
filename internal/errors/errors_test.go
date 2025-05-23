package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weprodev/go-mail/mail"
)

func TestError_Error(t *testing.T) {
	tt := map[string]struct {
		input *Error
		debug bool
		want  string
	}{
		"Normal": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			false,
			fmt.Sprintf("%s: err", Prefix),
		},
		"Debug": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			true,
			fmt.Sprintf("%s: op: err", Prefix),
		},
		"Nil Operation": {
			&Error{Code: INTERNAL, Message: "test", Operation: "", Err: fmt.Errorf("err")},
			false,
			fmt.Sprintf("%s: err", Prefix),
		},
		"Nil Err": {
			&Error{Code: INTERNAL, Message: "test", Operation: "", Err: nil},
			false,
			fmt.Sprintf("%s: <internal> test", Prefix),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			defer func() { mail.Debug = false }()
			mail.Debug = test.debug
			assert.Equal(t, test.want, test.input.Error())
		})
	}
}

func TestError_Code(t *testing.T) {
	tt := map[string]struct {
		input error
		want  string
	}{
		"Normal": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			"internal",
		},
		"Nil Input": {
			nil,
			"",
		},
		"Nil Code": {
			&Error{Code: "", Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			"internal",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, Code(test.input))
		})
	}
}

func Test_Message(t *testing.T) {
	tt := map[string]struct {
		input error
		want  string
	}{
		"Normal": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			"test",
		},
		"Nil Input": {
			nil,
			"",
		},
		"Nil Message": {
			&Error{Code: "", Message: "", Operation: "op", Err: fmt.Errorf("err")},
			GlobalError,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, Message(test.input))
		})
	}
}

func TestError_ToError(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  *Error
	}{
		"Pointer": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
		},
		"Non Pointer": {
			Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
		},
		"Error": {
			fmt.Errorf("err"),
			&Error{Err: fmt.Errorf("err")},
		},
		"String": {
			"err",
			&Error{Err: fmt.Errorf("err")},
		},
		"Default": {
			nil,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, ToError(test.input))
		})
	}
}

func TestNew(t *testing.T) {
	want := fmt.Errorf("error")
	got := New("error")
	assert.Errorf(t, want, got.Error())
}
