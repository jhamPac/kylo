package kylo

import (
	"context"
	"strings"
	"testing"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("normal reader test", func(t *testing.T) {
		rdr := strings.NewReader("abcdefghijk")
		got := make([]byte, 3)
		_, err := rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "abc")

		_, err = rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "def")
	})

	t.Run("behaves like a normal reader", func(t *testing.T) {
		ctx, _ := context.WithCancel(context.Background())
		rdr := NewCancellableReader(ctx, strings.NewReader("abcdefghijk"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "abc")

		_, err = rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "def")
	})

	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		rdr := NewCancellableReader(ctx, strings.NewReader("abcdefghijk"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)

		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "abc")

		cancel()

		n, err := rdr.Read(got)

		if err == nil {
			t.Error("expected an error after cancellation but didnt get one")
		}

		if n > 0 {
			t.Errorf("expected 0 bytes to be read after cancellation but %d were read", n)
		}
	})
}

func assertBufferHas(t *testing.T, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
