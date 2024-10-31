package config

import (
	admin "goBlog/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMIN
	//Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
    r.GET("/admin/add-new",admin.Dashboard{}.NewItem)
	r.POST("/admin/add",admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id",admin.Dashboard{}.Delete)
    r.GET("/admin/edit/:id",admin.Dashboard{}.Edit)   

	//SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath",http.Dir("uploads"))
	return r
}
