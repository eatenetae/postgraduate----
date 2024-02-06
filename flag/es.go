package flag

import "server/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()

}
