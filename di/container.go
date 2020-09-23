package di

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam0/pbi-auction/rest"
)

type Container struct{}

func (Container) Router() *gin.Engine {
	return rest.NewRouter()
}
