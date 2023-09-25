package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	FilePath = os.TempDir()

	tests := []struct {
		name string
		ex   string
		want bool
	}{
		{
			name: "Test GetConfigFile with yaml",
			ex:   "yaml",
			want: true,
		},
		{
			name: "Test GetConfigFile with Json",
			ex:   "json",
			want: true,
		},
		{
			name: "Test GetConfigFile with wrong extension",
			ex:   "txt",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeConfigFile(tt.ex)
			got, _ := GetConfigFile()
			if got != tt.want {
				t.Errorf("GetConfigFile() got = %v, want %v", got, tt.want)
			}

			os.Remove(filepath.Join(FilePath, fmt.Sprintf("%s.%s", Filename, tt.ex)))

		})
	}
}
