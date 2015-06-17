package tesseract

import (
	"os"
	"testing"
)

func TestReadText(t *testing.T) {
	tests := []struct {
		imgName  string
		expected string
	}{
		{"citibet01.jpg", "37233"},
		{"citibet02.jpg", "75526"},
		{"citibet03.jpg", "78442"},
	}
	for _, data := range tests {
		f, err := os.Open(data.imgName)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		text, err := ReadText(f)
		if err != nil {
			t.Error(err)
		}

		if text != data.expected {
			t.Errorf("expected %q, got %q", data.expected, text)
		}
	}
}

func TestReadTextFromFile(t *testing.T) {
	tests := []struct {
		imgName  string
		expected string
	}{
		{"citibet01.jpg", "37233"},
		{"citibet02.jpg", "75526"},
		{"citibet03.jpg", "78442"},
	}
	for _, data := range tests {
		text, err := ReadTextFromFile(data.imgName)
		if err != nil {
			t.Error(err)
		}

		if text != data.expected {
			t.Errorf("expected %q, got %q", data.expected, text)
		}
	}
}
