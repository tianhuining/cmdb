package http

import (
	"cmdb/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)


func (h *Hander)createHost(ctx *gin.Context){
	ins :=  host.NewHost()


    if	err := ctx.Bind(ins);err!=nil{
	//	ctx.Data(http.StatusBadRequest , "application/json" ,"")
		response.Failed(ctx.Writer , err)
		return
	}

	 ins, err := h.svc.CreateHost(ctx.Request.Context() , ins)
	 if err !=nil{
		 response.Failed(ctx.Writer , err)
		 return
	 }
	 response.Success(ctx.Writer ,ins )


}


