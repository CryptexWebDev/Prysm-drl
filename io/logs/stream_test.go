package logs

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func TestStreamServer_BackfillsMessages(t *testing.T) {
	ss := NewStreamServer()
	msgs := [][]byte{
		[]byte("foo"),
		[]byte("bar"),
		[]byte("buzz"),
	}
	for _, msg := range msgs {
		_, err := ss.Write(msg)
		require.NoError(t, err)
	}

	recentMessages := ss.GetLastFewLogs()
	require.DeepEqual(t, msgs, recentMessages)
}
