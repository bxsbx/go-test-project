package gin

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 设置gin模式
	gin.SetMode(GetEnvMode())
	var Router = gin.New()
	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 错误处理
	Router.Use(middleware.GinRecovery(true))
	// 准备
	Router.Use(middleware.Prepare())
	// 日志
	Router.Use(middleware.Logging())
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// 没有对应路由和对应方法时的处理
	Router.NoRoute(response.NoRouter)
	Router.NoMethod(response.NoMethod)

	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	contentCommunityRouter := router.RouterGroupApp.ContentCommunity
	personalCenterRouter := router.RouterGroupApp.PersonalCenter
	creativeCenterRouter := router.RouterGroupApp.CreativeCenter
	supermarketRouter := router.RouterGroupApp.CaseSupermarket
	acitivityRouter := router.RouterGroupApp.ActivityTopics
	evaluationRouter := router.RouterGroupApp.CaseEvaluation
	openApiRouter := router.RouterGroupApp.OpenApi
	portalApiRuter := router.RouterGroupApp.PortalApi
	voteApiRouter := router.RouterGroupApp.VoteApi

	// 公共路由组
	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)        // 注册基础功能路由 不做鉴权
		portalApiRuter.InitPortalApiRouter(PublicGroup) // 门户相关路由 不做鉴权
		systemRouter.InitUserRouter(PublicGroup)        // 注册用户路由
	}

	// 私有路由组
	PrivateGroup := Router.Group("")
	// TODO:注释授权
	PrivateGroup.Use(middleware.JWTAuth())
	{
		// 权限相关api
		systemRouter.InitAuthorityRouter(PrivateGroup)          // 注册角色路由
		systemRouter.InitJwtRouter(PrivateGroup)                // jwt相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)             // 权限相关路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
		//systemRouter.InitSystemRouter(PrivateGroup)             // system相关路由

		// 业务api
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		contentCommunityRouter.InitCommentRouter(PrivateGroup)      // 评论功能路由
		contentCommunityRouter.InitContentRouter(PrivateGroup)      // 教研内容功能路由
		contentCommunityRouter.InitRateRouter(PrivateGroup)         // 教研内容评价功能路由
		contentCommunityRouter.InitUserSpaceRouter(PrivateGroup)    // 教研内容用户空间功能相关路由
		contentCommunityRouter.InitUserRouter(PrivateGroup)         //教研内容用户功能路由
		contentCommunityRouter.InitHotWordsRouter(PrivateGroup)     //广场-搜索热词路由
		personalCenterRouter.InitPersonalCenterRouter(PrivateGroup) // 个人中心路由
		creativeCenterRouter.InitCreativeCenterRouter(PrivateGroup) // 创作中心
		supermarketRouter.InitCaseSupermarketRouter(PrivateGroup)   // 案例超市
		acitivityRouter.InitActivityRouter(PrivateGroup)            // 活动专题
		evaluationRouter.InitCaseEvaluationRouter(PrivateGroup)     // 课题评比
		openApiRouter.InitOpenApiRouter(PrivateGroup)               // 开放接口 主要给其他项目调用
		voteApiRouter.InitOpenApiRouter(PrivateGroup)               // 投票模块
	}

	global.GVA_LOG.Info("router register success")
	// 解析路由
	pattern.Parse(Router)
	return Router
}
