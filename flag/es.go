package flag

import "gvb_server/models"

func ESCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
