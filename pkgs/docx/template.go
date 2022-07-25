package docx

import (
	"archive/zip"
	"bytes"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"text/template"
)

type Template struct {
	reader   *zip.Reader
	template *template.Template
}

func ParseFile(path string) (Template, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msgf("can't open template file")
		return Template{}, err
	}
	return ParseBytes(data)
}

func ParseBytes(data []byte) (Template, error) {
	buf := bytes.NewReader(data)
	return Parse(buf, buf.Size())
}

func Parse(reader io.ReaderAt, size int64) (t Template, err error) {
	t.reader, err = zip.NewReader(reader, size)
	if err != nil {
		return t, err
	}
	content, err := readText(t.reader.File)
	if err != nil {
		return t, err
	}
	content = cleanText(content)
	t.template, err = template.New("").Option("missingkey=zero").Parse(content)
	return t, err
}

func (t *Template) Render(params interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := t.template.Execute(buf, params)
	if err != nil {
		return nil, err
	}
	return t.zipData(buf.Bytes())
}

func (t *Template) RenderPdf(params interface{}) ([]byte, error) {
	data, err := t.Render(params)
	if err != nil {
		return nil, err
	}

	return convertToPdf(data)
}

func (t *Template) zipData(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := zip.NewWriter(buf)

	for _, file := range t.reader.File {
		writer, err := w.Create(file.Name)
		if err != nil {
			return nil, err
		}
		reader, err := file.Open()
		if err != nil {
			return nil, err
		}
		if file.Name == "word/document.xml" {
			_, _ = writer.Write(data)
		} else {
			_, _ = writer.Write(streamToByte(reader))
		}
		_ = reader.Close()
	}
	w.Close()
	return buf.Bytes(), nil
}
