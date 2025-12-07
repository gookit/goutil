package ptr

import "testing"

func TestOf(t *testing.T) {
	ip := Of(10)
	if ip == nil || *ip != 10 {
		t.Fatalf("Of[int] failed, got: %v", ip)
	}

	sp := Of("hello")
	if sp == nil || *sp != "hello" {
		t.Fatalf("Of[string] failed, got: %v", sp)
	}

	bp := Of(true)
	if bp == nil || *bp != true {
		t.Fatalf("Of[bool] failed, got: %v", bp)
	}
}

func TestLegacyHelpers(t *testing.T) {
	i := 5
	ip := Int(i)
	if ip == nil || *ip != i {
		t.Fatalf("Int failed, got: %v", ip)
	}

	s := "world"
	sp := String(s)
	if sp == nil || *sp != s {
		t.Fatalf("String failed, got: %v", sp)
	}

	b := false
	bp := Bool(b)
	if bp == nil || *bp != b {
		t.Fatalf("Bool failed, got: %v", bp)
	}
}


