package ca

import "testing"

func TestStops(t *testing.T) {
	s := ParseStops()

	t.Run("checkLength", func(t *testing.T) {
		if len(s) == 0 {
			t.Error("bad len")
		}
		t.Log(len(s))
	})
	t.Run("firstEntryDesc", func(t *testing.T) {
		if s["10009"].Desc != "Bt Merah Int" {
			t.Error("bad first entry Desc")
		}
	})
	t.Run("firstEntryLat", func(t *testing.T) {
		if s["10009"].Lat != 1.2821 {
			t.Error("bad first entry Lat")
		}
	})
}
