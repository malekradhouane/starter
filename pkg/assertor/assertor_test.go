package assertor

import (
	"errors"
	"strings"
	"testing"
)

func TestAssertor_Assert(t *testing.T) {
	t.Parallel()

	type args struct {
		ok     bool
		errMsg string
		args   []any
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{ok: 2 == 2, errMsg: "right", args: []interface{}{}}, //nolint
			want: true,
		},
		{
			args: args{ok: 2 == 3, errMsg: "wrong", args: []interface{}{}},
			want: false,
		},
	}

	assertor := New()

	for i := range testCases {
		test := testCases[i]
		t.Run(test.name,
			func(t *testing.T) {
				t.Parallel()

				if got := assertor.Assert(test.args.ok, test.args.errMsg, test.args.args...); got != test.want {
					t.Errorf("Assertor.Assert() = %v, want %v", got, test.want)
				}
			})
	}
}

func TestAssertor_Validate(t *testing.T) {
	const (
		wrong1 = "wrong1"
		wrong2 = "wrong2"
		sep1   = "; "
		sep2   = "-"
	)

	t.Parallel()

	assertor := New()
	assertor.Assert(2 == 2, "ok") //nolint

	err := assertor.Validate()
	if err != nil {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, nil)
	}

	expected := ErrValidate.Error() + ": 2 unsatisfied requirement(s): " + strings.Join([]string{wrong1, wrong2}, sep1)

	assertor.Assert(2 == 3, wrong1) //nolint
	assertor.Assert(2 == 4, wrong2) //nolint

	err = assertor.Validate()
	if !errors.Is(err, ErrValidate) {
		t.Errorf("errors.Is(ErrValidate) should be true")
	}

	if err.Error() != expected {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, expected)
	}

	assertor.Separator = sep2
	expected = ErrValidate.Error() + ": 2 unsatisfied requirement(s): " + strings.Join([]string{wrong1, wrong2}, sep2)

	err = assertor.Validate()
	if err.Error() != expected {
		t.Errorf("Assertor.Validate() error = %v, wantErr %v", err, expected)
	}
}

func TestAssertor_Example(t *testing.T) {
	input1 := "123"
	chatEnabled := "true"
	chatServer := "chat.myserver.com:12345"

	v := New()

	v.Assert(len(input1) == 3, "string must be 3 characters long, not %d", len(input1))

	if v.Assert(chatEnabled == "true", "") {
		v.Assert(chatServer != "", "chat server address is missing")
	}

	if err := v.Validate(); err != nil {
		t.Errorf("Assertor.Validate() error = %v, want not error", err)
	}
}
