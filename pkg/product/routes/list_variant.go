package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/product/pb"
)

func ListVariant(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ListVariant(ctx, &empty.Empty{})
	if err != nil {
		ctx.AbortWithError(int(res.Status), err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
