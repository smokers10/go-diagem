package etc

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func UploadFile(path string, filename string, base64Encoded string) (fullpath string, err error) {
	// buat directory jika tidak ada dan set permission
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	// decode file
	dec, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filepath.Join(path, filepath.Base(filename)))
	if err != nil {
		return "", err
	}
	defer file.Close()

	// write file di penyimpanan sementara (ram)
	if _, err := file.Write(dec); err != nil {
		return "", err
	}

	// commit file ke penyimpanan
	if err := file.Sync(); err != nil {
		return "", err
	}

	// beres
	return fmt.Sprint(path, filename), nil
}
