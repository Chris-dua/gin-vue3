package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`
}

func FindImageInDB(imageHash string) (bool, string) {
	var bannerModel models.BannerModel
	err := global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		//找到
		return true, bannerModel.Path
	}
	return false, ""
}
func UploadToDB(filePath string, imageHash string, fileName string, imageType ctype.ImageType) {
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: imageType,
	})
}

func MkdirDirectory() string {
	basePath := global.Config.Upload.Path
	_, err := os.ReadDir(basePath)
	if err != nil {
		//递归创建文件夹
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	return basePath
}

func AppendResList(fileName string, isSuccess bool, msg string) (res FileUploadResponse) {
	res = FileUploadResponse{
		FileName:  fileName,
		IsSuccess: isSuccess,
		Msg:       msg,
	}
	return res
}

// ImageUploadView 上传单张图片，返回图片URL
func (ImagesApi) ImageUploadView(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), context)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", context)
		return
	}
	// 判断路径是否存在，不存在则创建路径(可以用time.Time 进行创建)
	basePath := MkdirDirectory()
	/*
		1、不超过直接上传
		2、超过这个大小就进行压缩保存
		3、路径生成
	*/
	var resList []FileUploadResponse
	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		//后缀名
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !utils.InList(suffix, WhiteImageList) {
			msg := "非法文件"
			resList = append(resList, AppendResList(file.Filename, false, msg))
			continue
		}
		filePath := path.Join(basePath, file.Filename)
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			msg := fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB，设定大小为：%dMB", size, global.Config.Upload.Size)
			resList = append(resList, AppendResList(file.Filename, false, msg))
			continue
		}
		//保存到目录
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		// 依据哈希 去数据库中是否重复
		isFind, bannerPath := FindImageInDB(imageHash)
		if isFind {
			msg := "图片已经存在"
			resList = append(resList, AppendResList(bannerPath, false, msg))
			continue
		}
		if global.Config.QiNiu.Enable {
			qiNiuFilePath, errQiNiu := qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
			if errQiNiu != nil {
				global.Log.Error(err)
				continue
			}
			msg := "上传七牛成功"
			resList = append(resList, AppendResList(qiNiuFilePath, true, msg))
			UploadToDB(qiNiuFilePath, imageHash, fileName, ctype.QiNiu)
			continue
		}
		err = context.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			msg := err.Error()
			resList = append(resList, AppendResList(file.Filename, false, msg))
			continue
		}
		msg := "上传成功"
		resList = append(resList, AppendResList(file.Filename, true, msg))

		// 图片入库
		UploadToDB(filePath, imageHash, fileName, ctype.Local)

	}
	res.OkWithData(resList, context)

}
