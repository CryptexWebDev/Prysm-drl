package bazel_test

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/build/bazel"
)

func TestBuildWithBazel(t *testing.T) {
	if !bazel.BuiltWithBazel() {
		t.Error("not built with Bazel")
	}
}
