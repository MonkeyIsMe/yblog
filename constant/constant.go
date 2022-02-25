package constant

const (
	LogInsert = 1 // LogInsert 记录日志时，插入的编号
	LogDelete = 2 // LogDelete 记录删除时，删除的编号
	LogUpdate = 3 // LogUpdate 记录更新时，更新的编号
	LogQuery  = 4 // LogQuery 记录查询时，查询的编号
)

const (
	TableComment      = "t_comment"       // TableComment 评论表表名
	TableBlog         = "t_blog"          // TableBlog 博客表表名
	TableLog          = "t_log"           // TableLog 日志表表名
	TableUser         = "t_user"          // TableUser 用户表表名
	TableTag          = "t_tag"           // TableTag 标签表表名
	TableTagBlog      = "t_tag_blog"      // TableTagBlog 标签和博客关联表表名
	TableClassify     = "t_classify"      // TableClassify 分类表表名
	TableClassifyBlog = "t_classify_blog" // TableClassifyBlog 分类和博客关联表表名
	TableSystem       = "t_system"        // TableSystem 系统信息表表名
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
