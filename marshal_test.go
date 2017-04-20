package xmp

import (
	"fmt"
	"testing"
	"time"
)

func TestAsPOD(t *testing.T) {
	created, _ := time.Parse(time.RFC3339, "2016-04-03T14:29:11-04:00")
	expect := &POD{
		Title:       "What Climate Change Means for Massachusetts",
		Description: "This fact sheet provides a concise overview of the observed and projected effects and impacts of climate change on Massachusetts.",
		Created:     &created,
	}
	packet, err := Unmarshal([]byte(xmpExample))
	if err != nil {
		t.Error(err.Error())
		return
	}

	pod := packet.AsPOD()

	if err := ComparePod(expect, pod); err != nil {
		t.Errorf(err.Error())
	}
}

func ComparePod(a, b *POD) error {
	if a == nil && b != nil || a != nil && b == nil {
		return fmt.Errorf("nil mismatch, %s != %s", a, b)
	}

	return nil
}
