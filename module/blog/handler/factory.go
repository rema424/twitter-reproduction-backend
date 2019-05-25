package handler

type (
	// Factory ...
	Factory interface {
		Article() ArticleHandler
	}

	factoryImpl struct {
		articleHandler ArticleHandler
	}
)

// NewFactory ...
func NewFactory(articleHandler ArticleHandler) Factory {
	return &factoryImpl{
		articleHandler: articleHandler,
	}
}

func (f *factoryImpl) Article() ArticleHandler {
	return f.articleHandler
}
