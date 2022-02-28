package constant

const (
	LogInsert = 1 // LogInsert 记录日志时，插入的编号
	LogDelete = 2 // LogDelete 记录删除时，删除的编号
	LogUpdate = 3 // LogUpdate 记录更新时，更新的编号
	LogQuery  = 4 // LogQuery 记录查询时，查询的编号
)

const (
	TableComment      = "t_comment"       // 评论表
	TableBlog         = "t_blog"          // 博客表
	TableLog          = "t_log"           // 日志表
	TableUser         = "t_user"          // 用户表
	TableTag          = "t_tag"           // 标签表
	TableTagBlog      = "t_tag_blog"      // 标签和博客关联表
	TableClassify     = "t_classify"      // 分类表
	TableClassifyBlog = "t_classify_blog" // 分类和博客关联表
	TableSystem       = "t_system"        // 系统表
	TableLinks        = "t_links"         // 友情链接表
	TableReply        = "t_reply"         // 回复表
	TableStatistics   = "t_statistics"    // 统计表
)

const (
	CommentDetail      = 1 // CommentDetail 记录日志操作时，操作的评论对象
	BlogDetail         = 2 // BlogDetail 记录日志操作时，操作的博客评论对象
	LogDetail          = 3 // LogDetail 记录日志操作时，操作的日志评论对象
	UserDetail         = 4 // UserDetail 记录日志操作时，操作的用户评论对象
	TagDetail          = 5 // TagDetail 记录日志操作时，操作的标签评论对象
	TagBlogDetail      = 6 // TagBlogDetail 记录日志操作时，操作的标签和博客对象
	ClassifyDetail     = 7 // ClassifyDetail 记录日志操作时，操作的分类对象
	ClassifyBlogDetail = 8 // ClassifyBlogDetail 记录日志操作时，操作的分类和博客对象
)

const (
	ErrorMsg   = 0 // ErrorMsg 出现错误的返回值
	SuccessMsg = 1 // SuccessMsg 成功的时候的返回值
	NullMsg    = 2 // NullMsg 消息为空的时候的返回值
)

const (
	Delete = -1
	Normal = 1
	UnSee  = 2
)
