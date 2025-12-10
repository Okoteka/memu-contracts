// test/test_import.go
// Этот файл проверяет что сгенерированный код можно импортировать
// Он НЕ генерируется в CI, а лежит в репозитории

package test

import (
	downloader "g.nas.loc/Okoteka/memu-contracts/downloader/v1"
	"os"
	"os/exec"
	"testing"
)

func TestProtoFilesExist(t *testing.T) {
	files := []string{
		"../gen/downloader/v1/downloader_service.pb.go",
		"../gen/downloader/v1/downloader_service_grpc.pb.go",
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("File doesnt't exist: %s", file)
		}
	}
}

func TestTypesDefined(t *testing.T) {
	var _ downloader.DownloadRequest
	var _ downloader.DownloadResponse
	var _ downloader.DownloaderServiceServer
	var _ downloader.DownloaderServiceClient
}

func TestGoBuildSucceeds(t *testing.T) {
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = "../gen"

	if output, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("Compilation fail:\n%s", output)
	}
}

func TestTypesFields(t *testing.T) {
	req := &downloader.DownloadRequest{
		Url: "https://example.com/file.mp4",
	}

	if req.Url == "" {
		t.Error("Request should have URL field")
	}

	resp := &downloader.DownloadResponse{
		Status:          downloader.DownloadResponse_STATUS_PROGRESS,
		ProgressPercent: 50,
	}

	if resp.Status != downloader.DownloadResponse_STATUS_PROGRESS {
		t.Error("Enum value should be accessible")
	}

}
