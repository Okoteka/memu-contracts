package test

import (
	"g.nas.loc/Okoteka/memu-contracts/gen/go/downloader/v1"
	"os"
	"os/exec"
	"testing"
)

func TestProtoFilesExist(t *testing.T) {
	files := []string{
		"../gen/go/downloader/v1/downloader_service.pb.go",
		"../gen/go/downloader/v1/downloader_service_grpc.pb.go",
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("File doesnt't exist: %s", file)
		}
	}
}

func TestTypesDefined(t *testing.T) {
	var _ downloaderv1.DownloadRequest
	var _ downloaderv1.DownloadResponse
	var _ downloaderv1.DownloaderServiceServer
	var _ downloaderv1.DownloaderServiceClient
}

func TestGoBuildSucceeds(t *testing.T) {
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = "../gen/go"

	if output, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("Compilation fail:\n%s", output)
	}
}

func TestTypesFields(t *testing.T) {
	req := &downloaderv1.DownloadRequest{
		Url: "https://example.com/file.mp4",
	}

	if req.Url == "" {
		t.Error("Request should have URL field")
	}

	resp := &downloaderv1.DownloadResponse{
		Status:          downloaderv1.DownloadResponse_STATUS_PROGRESS,
		ProgressPercent: 50,
	}

	if resp.Status != downloaderv1.DownloadResponse_STATUS_PROGRESS {
		t.Error("Enum value should be accessible")
	}

}
