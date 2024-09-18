package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/vector-ops/lifefolio/utils"
)

func TestGenPatientId(t *testing.T) {
	year, month, day := time.Now().Date()
	expected := fmt.Sprintf("LFPRT%d%d%d", year, month, day)
	got := utils.GenPatientId("Rockstar", "Teddy")

	if expected != got {
		t.Fatalf("expected: %s \ngot: %s", expected, got)
	}

}
