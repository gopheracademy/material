package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// StoreFiles stores file uploads at paths like /YYYY/MM/filename.ext
func StoreFiles(req *http.Request) (map[string]string, error) {
	err := req.ParseMultipartForm(1024 * 1024 * 4) // maxMemory 4MB
	if err != nil {
		return nil, err
	}

	ts := req.FormValue("timestamp") // timestamp in milliseconds since unix epoch

	if ts == "" {
		ts = fmt.Sprintf("%d", int64(time.Nanosecond)*time.Now().UnixNano()/int64(time.Millisecond)) // Unix() returns seconds since unix epoch
	}

	req.Form.Set("timestamp", ts)

	// To use for FormValue name:urlPath
	urlPaths := make(map[string]string)

	// get or create upload directory to save files from request
	pwd, err := os.Getwd()
	if err != nil {
		err := fmt.Errorf("Failed to locate current directory: %s", err)
		return nil, err
	}

	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, err
	}

	tm := time.Unix(int64(i/1000), int64(i%1000))

	urlPathPrefix := "api"
	uploadDirName := "uploads"

	uploadDir := filepath.Join(pwd, uploadDirName, fmt.Sprintf("%d", tm.Year()), fmt.Sprintf("%02d", tm.Month()))
	err = os.MkdirAll(uploadDir, os.ModeDir|os.ModePerm)

	// loop over all files and save them to disk
	for name, fds := range req.MultipartForm.File {
		filename := fds[0].Filename
		src, err := fds[0].Open()
		if err != nil {
			err := fmt.Errorf("Couldn't open uploaded file: %s", err)
			return nil, err

		}
		defer src.Close()

		// check if file at path exists, if so, add timestamp to file
		absPath := filepath.Join(uploadDir, filename)

		if _, err := os.Stat(absPath); !os.IsNotExist(err) {
			filename = fmt.Sprintf("%d-%s", time.Now().Unix(), filename)
			absPath = filepath.Join(uploadDir, filename)
		}

		// save to disk (TODO: or check if S3 credentials exist, & save to cloud)
		dst, err := os.Create(absPath)
		if err != nil {
			err := fmt.Errorf("Failed to create destination file for upload: %s", err)
			return nil, err
		}

		// copy file from src to dst on disk
		if _, err = io.Copy(dst, src); err != nil {
			err := fmt.Errorf("Failed to copy uploaded file to destination: %s", err)
			return nil, err
		}

		// add name:urlPath to req.PostForm to be inserted into db
		urlPath := fmt.Sprintf("/%s/%s/%d/%02d/%s", urlPathPrefix, uploadDirName, tm.Year(), tm.Month(), filename)

		urlPaths[name] = urlPath
	}

	return urlPaths, nil
}
