package api

import (
	"net/http"

	"github.com/PhamDuyKhang/userplayboar/internal/app/api/handler"
	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/middleware"

	"github.com/PhamDuyKhang/userplayboar/internal/app/logicpool"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"github.com/teera123/gin"
)

const (
	post   = http.MethodPost
	get    = http.MethodGet
	delete = http.MethodDelete
	put    = http.MethodPut
)

var (
	logger = glog.New().WithPrefix("api")
)

type (
	//Router hold all information of a router
	Router struct {
		desc     string
		path     string
		method   string
		endPoint gin.HandlerFunc
	}
)

//Init init  gin router instand for RESTfull service
func Init(cf *conf.Config) http.Handler {
	logicpool.NewLogicPool(cf)
	h := handler.NewHello(logicpool.HlSrv)
	crud := handler.NewCRUD(logicpool.EmployeeSvr)
	r := []Router{
		{
			desc:     "ping/pong service",
			path:     "/ping",
			method:   get,
			endPoint: ping,
		},
		{
			desc:     "hello service",
			path:     "/hello",
			method:   get,
			endPoint: h.Hello,
		},
		{
			desc:     "get employee",
			path:     "/employee",
			method:   get,
			endPoint: crud.FindEmployee,
		},
		{
			desc:     "create new employee",
			path:     "/employee",
			method:   post,
			endPoint: crud.AddNewEmployee,
		},
		{
			desc:     "update employee",
			path:     "/employee",
			method:   put,
			endPoint: crud.UpdateEmployee,
		},
		{
			desc:     "update employee",
			path:     "/employee",
			method:   delete,
			endPoint: crud.DeleteEmployee,
		},
	}
	ro := gin.New()
	// ro.Use(middleware.AcceptLang)
	ro.Use(middleware.Logging)
	v1 := ro.Group("/v1")
	for _, rou := range r {
		v1.Handle(rou.method, rou.path, rou.endPoint)
	}
	return ro

}
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}
