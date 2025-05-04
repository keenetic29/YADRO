package delivery

import (
	"ball-sorting/internal/domain"
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestStdoutOutput_WriteResult(t *testing.T) {
	tests := []struct {
		name     string
		result   *domain.Output
		expected string
	}{
		{
			name:     "can sort",
			result:   &domain.Output{CanSort: true},
			expected: "yes\n",
		},
		{
			name:     "cannot sort",
			result:   &domain.Output{CanSort: false},
			expected: "no\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			output := NewStdoutOutput(&buf, &buf)
			output.WriteResult(tt.result)

			if buf.String() != tt.expected {
				t.Errorf("got %q, want %q", buf.String(), tt.expected)
			}
		})
	}
}

func TestStdoutOutput_WriteError(t *testing.T) {
	var buf bytes.Buffer
	output := NewStdoutOutput(&buf, &buf)
	
	testErr := errors.New("test error") 
	output.WriteError(testErr)

	expected := "Ошибка: test error\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("got %q, want to contain %q", buf.String(), expected)
	}
}