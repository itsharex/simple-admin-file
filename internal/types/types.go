// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// required : true
	// min : 0
	Page uint64 `json:"page" validate:"required,number,gt=0"`
	// Page size | 单页数据行数
	// required : true
	// max : 100000
	PageSize uint64 `json:"pageSize" validate:"required,number,lt=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic path UUID request | 基础UUID地址参数请求
// swagger:model UUIDPathReq
type UUIDPathReq struct {
	// ID
	// Required: true
	Id string `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id *uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id *string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The request params of setting boolean status | 设置状态参数
// swagger:model StatusCodeReq
type StatusCodeReq struct {
	// ID
	Id string `json:"id"`
	// Status code | 状态码
	Status uint64 `json:"status" validate:"number"`
}

// The data when upload finished | 上传完成数据
// swagger:model UploadInfo
type UploadInfo struct {
	// File name | 文件名称
	Name string `json:"name"`
	// File path | 文件路径
	Url string `json:"url"`
}

// The response data when upload finished | 上传完成返回的数据
// swagger:model UploadResp
type UploadResp struct {
	BaseDataInfo
	// The  data when upload finished | 上传完成数据
	Data UploadInfo `json:"data"`
}

// Update file information params | 更新文件信息参数
// swagger:model UpdateFileReq
type UpdateFileReq struct {
	// ID
	// Required : true
	ID string `json:"id"`
	// File name | 文件名
	// Required : true
	Name string `json:"name" validate:"max=50"`
}

// Get file list params | 获取文件列表参数
// swagger:model FileListReq
type FileListReq struct {
	PageInfo
	// File type | 文件类型
	// max length : 10
	FileType *uint8 `json:"fileType,optional" validate:"omitempty,max=10"`
	// File name | 文件名
	// max length : 50
	FileName *string `json:"fileName,optional" validate:"omitempty,max=50"`
	// Create date period | 创建日期时间段
	Period []string `json:"period,optional"`
}

// The response data of file information | 文件信息数据
// swagger:model FileInfo
type FileInfo struct {
	BaseUUIDInfo
	// User's UUID | 用户的UUID
	UserUUID *string `json:"userUUID"`
	// File name | 文件名
	Name *string `json:"name"`
	// File type | 文件类型
	FileType *uint8 `json:"fileType"`
	// File size | 文件大小
	Size *uint64 `json:"size"`
	// File path | 文件路径
	Path *string `json:"path"`
	// File public status | 文件公开状态
	// false private true public | false 私人, true公开
	Status *uint8 `json:"status"`
	// The public URL | 公开访问的链接
	PublicPath *string `json:"publicPath"`
}

// The response data of file information list | 文件信息列表数据
// swagger:model FileListResp
type FileListResp struct {
	BaseDataInfo
	// The file list data | 文件信息列表数据
	Data FileListInfo `json:"data"`
}

// swagger:model FileListInfo
type FileListInfo struct {
	BaseListInfo
	// The file list data | 文件信息列表数据
	Data []FileInfo `json:"data"`
}
