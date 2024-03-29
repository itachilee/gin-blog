package upload

import (
	"fmt"
	"github.com/itachilee/ginblog/global"
	"github.com/itachilee/ginblog/pkg/file"
	"github.com/itachilee/ginblog/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return global.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImagePath() string {
	return global.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return global.AppSetting.RuntimeRootPath + GetImagePath()

}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range global.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	log.Printf("%s checkimageext fail %s", fileName, ext)
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Printf("%d rather %d", size, global.AppSetting.ImageMaxSize)
	return size <= global.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {

		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
