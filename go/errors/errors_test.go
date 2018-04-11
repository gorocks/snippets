package errors_test

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func a() error {
	return errors.New("I am an a error")
}

func b() error {
	return a()
}

func c() error {
	return b()
}

func open(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "failed to open %q", path)
	}
	return nil
}

func TestErrors(t *testing.T) {
	t.Logf("%+v", c())
	t.Logf("%v", c())
	t.Logf("%s", c())
	t.Logf("%q", c())
	t.Logf("%+v", open("a.txt"))
}
