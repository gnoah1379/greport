package docx

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// readText reads text from a word document
func readText(files []*zip.File) (text string, err error) {
	var documentFile *zip.File
	documentFile, err = retrieveWordDoc(files)
	if err != nil {
		return text, err
	}
	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil {
		return text, err
	}

	text = string(streamToByte(documentReader))
	return
}

// retrieveWordDoc fetches a word document.
func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/document.xml" {
			file = f
		}
	}
	if file == nil {
		err = errors.New("document.xml file not found")
	}
	return
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// normalize fixes quotation marks in documnet
func normalizeQuotes(in rune) rune {
	switch in {
	case '“', '”':
		return '"'
	case '‘', '’':
		return '\''
	}
	return in
}

// cleans template tagged text of all brakets
func normalizeAll(text string) string {
	brakets := regexp.MustCompile("<.*?>")
	quotes := regexp.MustCompile("&quot;")
	text = brakets.ReplaceAllString(text, "")
	text = quotes.ReplaceAllString(text, "\"")
	return strings.Map(normalizeQuotes, text)
}

func cleanText(text string) string {
	braketFinder := regexp.MustCompile("{{.*?}}")
	return braketFinder.ReplaceAllStringFunc(text, normalizeAll)
}

func convertToPdf(data []byte) ([]byte, error) {
	tempDir, err := os.MkdirTemp("", "greport")
	if err != nil {
		return nil, err
	}
	//defer os.RemoveAll(tempDir)
	filename := filepath.Join(tempDir, uuid.NewString())
	if err = os.WriteFile(filename+".docx", data, 0666); err != nil {
		log.Error().Err(err).Msg("write tmp file error")
		return nil, err
	}
	log.Debug().Msgf("convert docx to pdf file temp: %s", filename)
	out, err := exec.Command("lowriter",
		"--invisible",
		"--convert-to",
		"pdf:writer_pdf_Export",
		"--outdir",
		tempDir,
		filename+".docx").Output()
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("convert docx to pdf output: %s", out)
	result, err := os.ReadFile(filename + ".pdf")
	if err != nil {
		return nil, err
	}
	return result, nil
}
