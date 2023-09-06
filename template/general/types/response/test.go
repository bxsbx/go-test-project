package response

type TestFuncResp struct {
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"` //是否删除
	Title      string `json:"title"`  //标题
	Desc       string `json:"desc,omitempty"`
	Content    string `json:"content"`
}
