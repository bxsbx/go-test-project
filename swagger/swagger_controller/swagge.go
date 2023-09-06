package swagger_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type T struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

type T1 struct {
	Interface interface{} `json:"other"`
}

type Article struct {
	ID            uint32 `gorm:"primary_key" json:"id"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	CreatedOn     uint32 `json:"created_on"`
	ModifiedOn    uint32 `json:"modified_on"`
	DeletedOn     uint32 `json:"deleted_on"`
	IsDel         uint8  `json:"is_del"` //是否删除
	Title         string `json:"title"`  //标题
	Desc          string `json:"desc,omitempty"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	T             T      `json:"t"`  //结构体
	T1            T1     `json:"t1"` //结构体1
}

func NewArticle() Article {
	return Article{}
}

// @Summary	获取单个文章
// @Tags articles/uuuuu
// @Description 这里是描述
// @Produce	json
// @Param		id	path		int	true	"文章ID"
// @Response	500 {object}	string	"服务内部错误"
// @Response	200	{object}	Response{data=Article}
// @Router		/api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Article{Content: "ccqdqdqd"})
}

// @Summary	获取多个文章
// @Tags articles
// @Param		name		query		string	false	"文章名称"
// @Param		tag_id		query		int		false	"标签ID"
// @Param		state		query		int		false	"状态"
// @Param		page		query		int		false	"页码"
// @Param		page_size	query		int		false	"每页数量"
// @Success	200			{object}	Article	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "not 400")
}

// @Summary	创建文章
// @Tags articles
// @Produce	json
// @Param		tag_id			body		string	true	"标签ID"
// @Param		title			body		string	true	"文章标题"
// @Param		desc			body		string	false	"文章简述"
// @Param		cover_image_url	body		string	true	"封面图片地址"
// @Param		content			body		string	true	"文章内容"
// @Param		created_by		body		int		true	"创建者"
// @Param		state			body		int		false	"状态"
// @Success	200				{object}	Article	"成功"
// @Failure	400				{object}	string	"请求错误"
// @Failure	500				{object}	string	"内部错误"
// @Router		/api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "not 500")
}

// @Summary	更新文章
// @Produce	json
// @Tags uuuuu
// @Param		tag_id			body		string	false	"标签ID"
// @Param		title			body		string	false	"文章标题"
// @Param		desc			body		string	false	"文章简述"
// @Param		cover_image_url	body		string	false	"封面图片地址"
// @Param		content			body		string	false	"文章内容"
// @Param		modified_by		body		string	true	"修改者"
// @Success	200				{object}	Article	"成功"
// @Failure	400				{object}	string	"请求错误"
// @Failure	500				{object}	string	"内部错误"
// @Router		/api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "update")
}

// @Summary	删除文章
// @Produce	json
// @Tags uuuuu
// @Param		id	path		int		true	"文章ID"
// @Success	200	{string}	string	"成功"
// @Router		/api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	c.JSON(http.StatusHTTPVersionNotSupported, "not 505")
}
