package constant

// LogOpDetail 日志操作细节
type LogOpDetail struct {
	OldInfo string // 更改前的信息
	NewInfo string // 更改后的信息
}

// BlogConfig 博客配置
type BlogConfig struct {
	IsFilterComment int // 是否开启评论过滤
	BlogCount       int // 每页展示的博客数
	CommentCount    int // 每页显示的评论数
	IsBlogerFilter  int // 是否展示博主个人信息
	BlogRankRule    int // 博客排序的规则，1为时间，2位评论，3位点击量，4为最后编辑的时间
}
