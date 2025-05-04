package repository

import (
	"ball-sorting/internal/domain"
	"strings"
	"testing"
)

func TestStdinInput_ReadInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *domain.Input
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "2\n1 2\n3 4\n",
			want: &domain.Input{
				N: 2,
				Containers: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewStdinInput(strings.NewReader(tt.input))
			got, err := r.ReadInput()

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.N != tt.want.N {
				t.Errorf("Got N = %v, want %v", got.N, tt.want.N)
			}
		})
	}
}