package usecase

import (
	"ball-sorting/internal/domain"
	"ball-sorting/internal/repository"
	"strings"
	"testing"
)

type mockOutputWriter struct {
	result *domain.Output
	err    error
}

func (m *mockOutputWriter) WriteResult(result *domain.Output) {
	m.result = result
}

func (m *mockOutputWriter) WriteError(err error) {
	m.err = err
}

func TestSorter_Run(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     bool
		wantErr  bool
		errMsg   string
	}{
		{
			name:    "can sort",
			input:   "2\n2 1\n1 2\n",
			want:    true,
			wantErr: false,
		},
		{
			name:    "cannot sort",
			input:   "2\n1 2\n3 4\n",
			want:    false,
			wantErr: false,
		},
		{
			name:    "input error - invalid N",
			input:   "abc\n1 2\n3 4\n",
			wantErr: true,
			errMsg:  "некорректное значение N",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInput := repository.NewStdinInput(strings.NewReader(tt.input))
			mockOutput := &mockOutputWriter{}

			sorter := NewSorter(mockInput, mockOutput)
			err := sorter.Run()

			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if mockOutput.result.CanSort != tt.want {
					t.Errorf("got %v, want %v", mockOutput.result.CanSort, tt.want)
				}
			} else {
				if mockOutput.err == nil || !strings.Contains(mockOutput.err.Error(), tt.errMsg) {
					t.Errorf("expected error containing %q, got %v", tt.errMsg, mockOutput.err)
				}
			}
		})
	}
}