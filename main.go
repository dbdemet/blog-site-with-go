package main

import (
	admin_models "goBlog/admin/models"
	"goBlog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	
	//(func delete)>>>db.Delete(&post,post.ID)	belirtilen adresteki, belirtilen ID ye ait veriyi siler.
	http.ListenAndServe(":8080", config.Routes())
}
