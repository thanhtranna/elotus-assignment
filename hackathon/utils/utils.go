package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

var (
	alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

func AddPrefixFileName(fileName string) string {
	names := strings.Split(fileName, ".")
	if len(names) < 2 {
		return ""
	}

	filename := names[0]
	prefix := RandomString(10)

	return fmt.Sprintf("%s-%s.%s", filename, prefix, names[1])
}

// hackathon/tmp/

func GetFileType(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return "", err
	}

	buf, _ := ioutil.ReadAll(f)

	fileType := http.DetectContentType(buf)

	return fileType, nil
}
