package main

import (
	"github.com/ganeshrvel/go-mtpfs/mtp"
	"github.com/ganeshrvel/go-mtpx"
	"os"
)

type verifyMtpSessionMode struct {
	skipDeviceChangeCheck bool
}

type deviceContainer struct {
	dev        *mtp.Device
	deviceInfo *mtp.DeviceInfo
	locked     bool
}

type MakeDirectoryInput struct {
	StorageId uint32 `json:"storageId"`
	FullPath  string `json:"fullPath"`
}

type FileExistsInput struct {
	StorageId uint32   `json:"storageId"`
	Files     []string `json:"Files"`
}

type DeleteFileInput struct {
	StorageId uint32   `json:"storageId"`
	Files     []string `json:"Files"`
}

type RenameFileInput struct {
	StorageId   uint32 `json:"storageId"`
	FullPath    string `json:"fullPath"`
	NewFileName string `json:"newFileName"`
}

type WalkInput struct {
	StorageId           uint32 `json:"storageId"`
	FullPath            string `json:"fullPath"`
	Recursive           bool   `json:"recursive"`
	SkipDisallowedFiles bool   `json:"skipDisallowedFiles"`
	SkipHiddenFiles     bool   `json:"skipHiddenFiles"`
}

type UploadFilesInput struct {
	StorageId       uint32   `json:"storageId"`
	Sources         []string `json:"sources"`
	Destination     string   `json:"destination"`
	PreprocessFiles bool     `json:"preprocessFiles"`
}

type DownloadFilesInput struct {
	StorageId       uint32   `json:"storageId"`
	Sources         []string `json:"sources"`
	Destination     string   `json:"destination"`
	PreprocessFiles bool     `json:"preprocessFiles"`
}

type UploadPreprocessContainer struct {
	fi       *os.FileInfo
	fullPath string
}

type DownloadPreprocessContainer struct {
	fi *mtpx.FileInfo
}

type ProgressContainer struct {
	pInfo *mtpx.ProgressInfo
}
