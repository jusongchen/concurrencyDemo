package demoapp

import (
	"context"
	"testing"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			Run(ctx, tt.numWorker, tt.numJob)

		})
	}
}
