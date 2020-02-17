package webframework

import "github.com/gin-gonic/gin"

type GinRouter struct {
	*gin.Engine
}

func NewGinWebFrameWork() *GinRouter {
	return &GinRouter{gin.Default()}
}

func (g *GinRouter) Run() {
	err := g.Engine.Run(":8080")
	if err != nil {
		//TODO
	}
}
