package tool

// GetAllTag 拼接一个or条件的字符串去查询标签的信息
func GetAllTag(tagList []string) string {
	tagStrList := ""
	for i, tag := range tagList {
		if i != len(tagList)-1 {
			tagStrList = tagStrList + "id = " + tag + " or "
		} else {
			tagStrList = tagStrList + "id = " + tag
		}
	}
	return tagStrList
}
