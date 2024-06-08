package protobuf

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	latestReleaseInfoUrl = "https://api.github.com/repos/protocolbuffers/protobuf/releases/latest"
	releaseDownloadUrl   = "https://github.com/protocolbuffers/protobuf/releases/download/%s/%s"
)

func DownloadProtobuf(version string) (string, error) {
	log.Println("INFO: Checking for existing protoc installation...")

	cachePath, err := getCachePath()
	if err != nil {
		return "", err
	}

	log.Printf("INFO: Protoc cache path: %s", cachePath)
	log.Printf("INFO: Searching for version %s...", version)

	if version == "" || version == "latest" {
		v, err := getLatestVersion()
		if err != nil {
			return "", err
		}

		version = v
	}

	version = normalizeVersion(version)
	tag := "v" + version

	log.Printf("INFO: Found version: %s", version)

	versionDir := path.Join(cachePath, tag)
	if _, err := os.Stat(versionDir); err == nil {
		log.Printf("INFO: Found existing installation at %s", versionDir)
		return versionDir, nil
	}

	osn := getOS()
	filename := getProtobufFilename(version, osn, getArch(osn))
	if filename == "" {
		return "", fmt.Errorf("unsupported platform")
	}

	log.Printf("INFO: Downloading %s...", filename)

	resp, err := http.Get(fmt.Sprintf(releaseDownloadUrl, tag, filename))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	protobufZip := path.Join(cachePath, filename)
	file, err := os.Create(protobufZip)
	if err != nil {
		return "", err
	}

	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return "", err
	}

	log.Printf("INFO: Extracting %s...", protobufZip)

	targetDir := path.Join(cachePath, tag)
	if err := unzip(protobufZip, targetDir); err != nil {
		return "", err
	}

	log.Printf("INFO: Cleaning up %s...", protobufZip)

	if err := os.Remove(protobufZip); err != nil {
		log.Printf("WARNING: failed to remove %s: %v (its a temp file)", protobufZip, err)
	}

	return targetDir, nil
}

// Returns win, osx, or linux, or empty string if unknown
func getOS() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "win"
	case "darwin":
		return "osx"
	case "linux":
		return "linux"
	default:
		return ""
	}
}

// Returns x64, x86, or empty string if unknown
func getArch(os string) string {
	switch arch := runtime.GOARCH; arch {
	case "amd64":
		fallthrough
	case "x86_64":
		if os == "win" {
			return "64"
		} else if os == "osx" || os == "linux" {
			return "-x86_64"
		}

		return ""
	case "386":
		fallthrough
	case "x86":
		if os == "win" {
			return "32"
		} else if os == "linux" {
			return "-x86_32"
		}

		return ""
	default:
		return ""
	}
}

func getProtobufFilename(version, os, arch string) string {
	if version == "" || os == "" || arch == "" {
		return ""
	}

	return fmt.Sprintf("protoc-%s-%s.zip", version, os+arch)
}

func normalizeVersion(version string) string {
	if !strings.HasPrefix(version, "v") {
		return version
	}

	return version[1:]
}

func getLatestVersion() (string, error) {
	resp, err := http.Get(latestReleaseInfoUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var releaseInfo struct {
		TagName string `json:"tag_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&releaseInfo); err != nil {
		return "", err
	}

	return releaseInfo.TagName, nil
}

func getCachePath() (string, error) {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	dir := path.Join(cachePath, "protolsd")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
