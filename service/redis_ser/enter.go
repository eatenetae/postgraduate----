package redis_ser

const (
	articleLookPrefix    = "article_look"
	articleDiggPrefix    = "article_digg"
	articleCommentPrefix = "article_comment_count"
	commentDiggPrefix    = "comment_digg"
)

type RedisService struct {
}

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
