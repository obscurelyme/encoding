package pom_test

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/obscurelyme/encoding/pom"
)

//go:embed pom.xml
var testPomFile []byte

func setup(tmpDir string) error {
	return os.WriteFile(fmt.Sprintf("%s/pom.xml", tmpDir), testPomFile, 0644)
}

func readPomFile(tmpDir string) ([]byte, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/pom.xml", tmpDir))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestPom(t *testing.T) {
	tmpDir := t.TempDir()
	setup(tmpDir)

	t.Run("Should read a basic pom file", func(t *testing.T) {
		data, err := readPomFile(tmpDir)
		if err != nil {
			t.Errorf("Expected no errors reading pom file, but found: %s", err.Error())
			return
		}

		var pomModel pom.Model

		err = xml.Unmarshal(data, &pomModel)
		if err != nil {
			t.Errorf("Expected no errors unmarshalling pom file, but found: %s", err.Error())
			return
		}
	})
}
