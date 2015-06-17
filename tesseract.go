package tesseract

import (
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/harrydb/go/img/grayscale"
)

// ExecutablePath should be tesseract.exe path if tesseract.exe is not in the PATH
var ExecutablePath string

// ReadTextFromFile read text from the file. It internally calls ReadText after reading the file.
func ReadTextFromFile(f string) (string, error) {
	r, err := os.Open(f)
	if err != nil {
		return "", err
	}
	defer r.Close()

	return ReadText(r)
}

// ReadText read text from the given io.Reader r. It converts to grayscale first before pass it to tesseract.
// It writes grayscale image and output text file to the os.TempFile.
func ReadText(r io.Reader) (string, error) {
	grayImg, err := covertGrayscale(r)

	outfile, err := ioutil.TempFile("", "ghost-tesseract-out-")
	defer outfile.Close()
	if err != nil {
		return "", err
	}

	if err = runOcr(grayImg.Name(), outfile.Name()); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(outfile.Name() + ".txt")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(bytes)), nil
}

func runOcr(in string, out string) error {
	exePath := ExecutablePath
	if exePath == "" {
		path, err := exec.LookPath("tesseract.exe")
		if err != nil {
			return err
		}

		exePath = path
	}

	ocr, err := filepath.Abs(exePath)
	if err != nil {
		return err
	}

	cmd := exec.Command(ocr, in, out)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func covertGrayscale(r io.Reader) (*os.File, error) {
	src, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}

	gray := grayscale.Convert(src, grayscale.ToGrayLuminance)
	grayImg, err := ioutil.TempFile("", "tesseract-gray-")
	defer grayImg.Close()
	if err != nil {
		return nil, err
	}

	err = jpeg.Encode(grayImg, gray, &jpeg.Options{Quality: 80})
	if err != nil {
		return nil, err
	}

	return grayImg, nil
}
