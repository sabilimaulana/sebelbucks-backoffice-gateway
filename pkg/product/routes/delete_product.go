package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/product/pb"
)

func DeleteProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DeleteProduct(ctx, &pb.DeleteProductRequest{Id: id})
	if err != nil {
		ctx.AbortWithError(int(res.Status), err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
