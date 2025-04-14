package ca

import "testing"

func TestStops(t *testing.T) {
	s := ParseStops("../../testdata/cheeaun/stops.json")

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

func TestServices(t *testing.T) {
	s := ParseServices("../../testdata/cheeaun/services.json")
	t.Run("checkLength", func(t *testing.T) {
		if len(s) == 0 {
			t.Error("bad len")
		}
		t.Log(len(s))
	})
	t.Run("checkFirstService", func(t *testing.T) {
		if s["2"].Name != "Changi Village Ter ⇄ Kampong Bahru Ter" {
			t.Error("bad first service name")
		}
	})
	t.Run("checkFirstServiceRoute", func(t *testing.T) {
		r00 := s["2"].Route[0][0]
		if r00 != "99009" {
			t.Error("bad first service stop: dir: 0")
		}

		r10 := s["2"].Route[1][0]
		if r10 != "10499" {
			t.Error("bad first service stop: dir: 1")
		}
	})
}
