package constant

const (
	LogInsert = 1
	LogDelete = 2
	LogUpdate = 3
	LogQuery  = 4
)

const (
	TableComment      = "t_comment"
	TableBlog         = "t_blog"
	TableLog          = "t_log"
	TableUser         = "t_user"
	TableTag          = "t_tag"
	TableTagBlog      = "t_tag_blog"
	TableClassify     = "t_classify"
	TableClassifyBlog = "t_classify_blog"
	TableSystem       = "t_system"
	TableLinks        = "t_links"
)

const (
	CommentDetail      = 1
	BlogDetail         = 2
	LogDetail          = 3
	UserDetail         = 4
	TagDetail          = 5
	TagBlogDetail      = 6
	ClassifyDetail     = 7
	ClassifyBlogDetail = 8
)

const (
	Delete = -1
	Normal = 1
	UnSee  = 2
)
