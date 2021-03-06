package demoapp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		numWorker int
		numJob    int
		wantErr   bool
	}{
		{
			name:      "worker_1_job_1",
			numWorker: 1,
			numJob:    1,
			wantErr:   false,
		},
		{
			name:      "worker_2_job_10",
			numWorker: 2,
			numJob:    10,
			wantErr:   false,
		},
		{
			name:      "worker_4_job_1000",
			numWorker: 4,
			numJob:    1000,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := New(tt.numWorker, tt.numJob)
			require.NoError(t, err)
			require.NotNil(t, app)
			app.Close()

		})
	}
}
