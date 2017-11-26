package clock

import (
	"testing"
	"time"
)

var (
	oneAM                    = From(time.Date(0, 0, 0, 1, 0, 0, 0, time.UTC))
	twoAM                    = From(time.Date(0, 0, 0, 2, 0, 0, 0, time.UTC))
	twoTwentyNineFiftyNineAM = From(time.Date(0, 0, 0, 2, 29, 59, 0, time.UTC))
	twoThirtyAM              = From(time.Date(0, 0, 0, 2, 30, 0, 0, time.UTC))
	twoThirtyThirtyAM        = From(time.Date(0, 0, 0, 2, 30, 30, 0, time.UTC))
)

func TestBefore(t *testing.T) {
	if twoAM.Before(oneAM) {
		t.Errorf("%s should not be before %s", twoAM, oneAM)
	}

	if twoThirtyThirtyAM.Before(twoThirtyAM) {
		t.Errorf("%s should not be before %s", twoThirtyThirtyAM, twoThirtyAM)
	}
}

func TestAfter(t *testing.T) {
	if oneAM.After(twoAM) {
		t.Errorf("%s should not be after %s", oneAM, twoAM)
	}

	if twoThirtyAM.After(twoThirtyThirtyAM) {
		t.Errorf("%s should not be after %s", twoThirtyAM, twoThirtyThirtyAM)
	}
}

func TestDiff(t *testing.T) {
	d := oneAM.Diff(twoAM)
	d2 := twoThirtyAM.Diff(twoThirtyThirtyAM)
	d3 := twoThirtyAM.Diff(twoTwentyNineFiftyNineAM)
	d4 := twoTwentyNineFiftyNineAM.Diff(twoThirtyAM)

	if d != time.Hour {
		t.Errorf("%s diff %s should be 1h, not %d", twoAM, oneAM, d)
	}

	if d2 != 30*time.Second {
		t.Errorf("%s diff %s should be 30s, not %d", twoThirtyAM, twoThirtyThirtyAM, d2)
	}

	if d3 != time.Second {
		t.Errorf("%s diff %s should be 1s, not %d", twoThirtyAM, twoTwentyNineFiftyNineAM, d3)
	}

	if d4 != time.Second {
		t.Errorf("%s diff %s should be 1s, not %d", twoTwentyNineFiftyNineAM, twoThirtyAM, d4)
	}
}
