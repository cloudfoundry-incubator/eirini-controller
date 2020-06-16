package util

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/gomega" //nolint:golint,stylecheck
)

func WriteTempFile(content []byte, fileName string) string {
	configFile, err := ioutil.TempFile("", fileName)
	Expect(err).ToNot(HaveOccurred())
	defer configFile.Close()

	err = ioutil.WriteFile(configFile.Name(), content, os.ModePerm)
	Expect(err).ToNot(HaveOccurred())
	return configFile.Name()
}
