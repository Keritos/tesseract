package tesseract

import (
	"os"
	"testing"
)

var tests = []struct {
	imgName  string
	expected string
}{
	{"img01.jpg", "37233"},
	{"img02.jpg", "75526"},
	{"img03.jpg", "78442"},
}

func TestReadText(t *testing.T) {
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
