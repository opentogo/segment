package segment

import (
	"testing"

	"github.com/allisson/go-assert"
)

func TestSegment(t *testing.T) {
	t.Run("verifying empty path", func(t *testing.T) {
		segment := NewSegment("", "")

		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/", segment.Current())
		assert.Equal(t, true, segment.Root())
	})

	t.Run("ignoring a piece of the path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "/foo")
		assert.Equal(t, "/bar/baz", segment.Current())
	})

	t.Run("checking root path", func(t *testing.T) {
		segment := NewSegment("/", "")

		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/", segment.Current())
		assert.Equal(t, true, segment.Root())
	})

	t.Run("checking if is from beginning of the path", func(t *testing.T) {
		segment := NewSegment("/", "")

		assert.Equal(t, true, segment.Init())
		assert.Equal(t, true, segment.Root())

		segment = NewSegment("/foo", "")

		assert.Equal(t, true, segment.Init())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "foo", segment.Extract())

		assert.Equal(t, false, segment.Init())
		assert.Equal(t, true, segment.Root())
	})

	t.Run("extracting values from path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "")

		assert.Equal(t, "foo", segment.Extract())
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "bar", segment.Extract())
		assert.Equal(t, "/foo/bar", segment.Previous())
		assert.Equal(t, "/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "baz", segment.Extract())
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())

		assert.Equal(t, "", segment.Extract())
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())
	})

	t.Run("retracting values to path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "")

		segment.Extract()
		segment.Extract()
		segment.Extract()

		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())

		assert.Equal(t, "baz", segment.Retract())
		assert.Equal(t, "/foo/bar", segment.Previous())
		assert.Equal(t, "/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "bar", segment.Retract())
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "foo", segment.Retract())
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, "", segment.Retract())
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())
	})

	t.Run("consuming values from path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "")

		assert.Equal(t, false, segment.Consume("bar"))
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, false, segment.Consume("fo"))
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, true, segment.Consume("foo"))
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, false, segment.Consume("foo"))
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, true, segment.Consume("bar"))
		assert.Equal(t, "/foo/bar", segment.Previous())
		assert.Equal(t, "/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, true, segment.Consume("baz"))
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())

		assert.Equal(t, false, segment.Consume("baz"))
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())
	})

	t.Run("restoring values to path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "")

		segment.Extract()
		segment.Extract()
		segment.Extract()

		assert.Equal(t, false, segment.Restore("foo"))
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())

		assert.Equal(t, true, segment.Restore("baz"))
		assert.Equal(t, "/foo/bar", segment.Previous())
		assert.Equal(t, "/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, true, segment.Restore("bar"))
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, true, segment.Restore("foo"))
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		assert.Equal(t, false, segment.Restore("foo"))
		assert.Equal(t, "", segment.Previous())
		assert.Equal(t, "/foo/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())
	})

	t.Run("capturing values from path", func(t *testing.T) {
		segment := NewSegment("/foo/bar/baz", "")
		captures := map[string]string{}

		segment.Capture("c1", captures)
		assert.Equal(t, 1, len(captures))
		assert.Equal(t, "/foo", segment.Previous())
		assert.Equal(t, "/bar/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		segment.Capture("c2", captures)
		assert.Equal(t, 2, len(captures))
		assert.Equal(t, "/foo/bar", segment.Previous())
		assert.Equal(t, "/baz", segment.Current())
		assert.Equal(t, false, segment.Root())

		segment.Capture("c3", captures)
		assert.Equal(t, 3, len(captures))
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())

		segment.Capture("c4", captures)
		assert.Equal(t, 3, len(captures))
		assert.Equal(t, "/foo/bar/baz", segment.Previous())
		assert.Equal(t, "", segment.Current())
		assert.Equal(t, true, segment.Root())

		assert.Equal(t, "foo", captures["c1"])
		assert.Equal(t, "bar", captures["c2"])
		assert.Equal(t, "baz", captures["c3"])
		assert.Equal(t, "", captures["c4"])
	})
}
