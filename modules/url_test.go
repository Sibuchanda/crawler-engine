package modules

import "testing"

func TestResolveURLNULL(t *testing.T) {
	result := resolveURL("", "")

	if result != "" {
		t.Errorf("resolveURL() => %v; expected '' (empty string)", result)
	}
}

func TestResolvedURLValidBase(t *testing.T) {
	base := "https://www.google.com"
	result := resolveURL("", base)

	if result != "" {
		t.Errorf("resolveURL() => %v; expected '' (empty string)", result)
	}
}

func TestResolveURLValidRef(t *testing.T) {
	ref := "https://www.google.com"
	result := resolveURL(ref, "")

	if result != ref {
		t.Errorf("resolveURL() => %v; expected %v", result, ref)
	}
}

func TestResolveURLInvalidBase(t *testing.T) {
	base := "abcd"
	ref := ""
	result := resolveURL(ref, base)

	if result != "" {
		t.Errorf("resolveURL() => %v; expected '' (empty string)", result)
	}
}

func TestResolveURLInvalidRef(t *testing.T) {
	base := ""
	ref := "abcd"
	result := resolveURL(ref, base)

	if result != "" {
		t.Errorf("resolveURL() => %v; expected '' (empty string)", result)
	}
}

func TestResolveURLValidRefValidBase(t *testing.T) {
	base := "https://www.google.com"
	ref := "https://example.com"
	result := resolveURL(ref, base)

	if result != ref {
		t.Errorf("resolveURL() => %v; expected %v", result, ref)
	}
}

func TestResolveURLValidRefInvalidBase(t *testing.T) {
	base := "abcd"
	ref := "https://example.com"
	result := resolveURL(ref, base)

	if result != ref {
		t.Errorf("resolveURL() => %v; expected %v", result, ref)
	}
}

func TestResolveURLInvalidRefValidBase(t *testing.T) {
	base := "https://example.com"
	ref := "abcd"
	result := resolveURL(ref, base)

	if result != (base + "/" + ref) {
		t.Errorf("resolveURL() => %v; expected %v", result, (base + "/" + ref))
	}
}

func TestResolveURLInvalidRefInvalidBase(t *testing.T) {
	base := "abcd"
	ref := "abcd"
	result := resolveURL(ref, base)

	if result != "" {
		t.Errorf("resolveURL() => %v; expected '' (empty string)", result)
	}
}
