package exec

import (
	"context"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dolong2/dcd-cli/api"
)

func UploadVolumeFile(ctx context.Context, workspaceId string, volumeId string, localFilePath string, targetPath string, createDirectory bool) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	param := make(map[string]string)
	param["path"] = targetPath
	param["createDirectory"] = strconv.FormatBool(createDirectory)

	_, uploadErr := api.SendPostMultipart(ctx, "/"+workspaceId+"/volume/"+volumeId+"/files", header, param, "file", filepath.Base(localFilePath), file)
	closeErr := file.Close()
	if uploadErr != nil {
		return uploadErr
	}
	return closeErr
}
