package kylo

import (
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
}

func assertBufferHas(t *testing.T, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
