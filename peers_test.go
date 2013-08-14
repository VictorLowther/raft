package raft

import (
	"testing"
	"strconv"
)

func TestQuorum(t *testing.T) {
	for _, tuple := range []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 4},
		{7, 4},
		{8, 5},
		{9, 5},
		{10, 6},
		{11, 6},
	} {
		pm := peerMap{}
		for i := 0; i < tuple.n; i++ {
			n := strconv.Itoa(i)
			pm[n] = nonresponsivePeer(n)
		}
		if expected, got := tuple.expected, pm.quorum(); expected != got {
			t.Errorf("Quorum of %d: expected %d, got %d", tuple.n, expected, got)
		}
	}
}
