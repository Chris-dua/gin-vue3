package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
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

func findImageInDB(imageHash string) (bool, string) {
	var bannerModel models.BannerModel
	err := global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		//找到
		return true, bannerModel.Path
	}
	return false, ""
}
func UploadToDB(filePath string, imageHash string, fileName string) {
	global.DB.Create(&models.BannerModel{
		Path: filePath,
		Hash: imageHash,
		Name: fileName,
	})
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
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		//递归创建文件夹
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

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
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}
		filePath := path.Join(basePath, file.Filename)
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB，设定大小为：%dMB", size, global.Config.Upload.Size),
			})
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
		isFind, bannerPath := findImageInDB(imageHash)
		if isFind {
			resList = append(resList, FileUploadResponse{
				FileName:  bannerPath,
				IsSuccess: false,
				Msg:       "图片已经存在",
			})
			continue
		}

		err = context.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})

		// 图片入库
		UploadToDB(filePath, imageHash, fileName)

	}
	res.OkWithData(resList, context)

}
