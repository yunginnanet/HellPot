package heffalump

import (
	"errors"
	"io"
	"testing"

	"github.com/yunginnanet/HellPot/internal/util"
)

func TestHeffalump_WriteHell(t *testing.T) {
	sp := util.NewCappedSpeedometer(io.Discard, 55555)
	var err error
	var count int64
	for err == nil {
		var cnt int64
		cnt, err = DefaultHeffalump.WriteHell(sp.BufioWriter)
		t.Logf("written: %d", cnt)
		count += cnt
	}
	if !errors.Is(err, util.ErrLimitReached) {
		t.Errorf("expected %v, got %v", util.ErrLimitReached, err)
	} else {
		t.Logf("err: %v", err)
	}
	t.Logf("count: %d", count)
	t.Logf("rate: %f per second", sp.Rate())
}
