package redpacket

import (
	"testing"
)

func TestDistributeCentsBasic(t *testing.T) {
	total := 10_00 // 10 元 -> 1000 分
	n := 7
	shares, err := DistributeCents(total, n)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(shares) != n {
		t.Fatalf("expected %d shares, got %d", n, len(shares))
	}

	sum := 0
	min := shares[0]
	max := shares[0]
	for _, v := range shares {
		if v <= 0 {
			t.Fatalf("share must be positive, got %d", v)
		}
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if sum != total {
		t.Fatalf("sum mismatch: want %d, got %d", total, sum)
	}
	if !(max < 2*min) {
		t.Fatalf("ratio constraint violated: max=%d min=%d", max, min)
	}
}

func TestDistributeCentsEdgeBaseOne(t *testing.T) {
	// 当 base==1 时，为满足 max < 2*min，必须无余额可再分配（rem==0）。
	total := 11
	n := 10
	_, err := DistributeCents(total, n)
	if err == nil {
		t.Fatalf("expected error for base==1 with non-zero remainder")
	}

	total = 10
	shares, err := DistributeCents(total, n)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, v := range shares {
		if v != 1 {
			t.Fatalf("expected all ones, got %v", shares)
		}
	}
}

func TestDistributeYuan(t *testing.T) {
	m := 12.34
	n := 4
	shares, err := DistributeYuan(m, n)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(shares) != n {
		t.Fatalf("len mismatch: got %d", len(shares))
	}
	sum := 0.0
	min := shares[0]
	max := shares[0]
	for _, v := range shares {
		if v <= 0 {
			t.Fatalf("share must be positive, got %f", v)
		}
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// 在“分”域比较，避免浮点误差影响
	gotCents := int(sum*100 + 0.5)
	wantCents := int(m*100 + 0.5)
	if gotCents != wantCents {
		t.Fatalf("sum mismatch cents: want %d got %d", wantCents, gotCents)
	}
	if !(max < 2*min) {
		t.Fatalf("ratio violated for yuan: max=%f min=%f", max, min)
	}
}
