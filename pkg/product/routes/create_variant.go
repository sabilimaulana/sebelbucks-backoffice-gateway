package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/product/pb"
)

type CreateVariantRequestBody struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Description string `json:"description"`
}

func CreateVariant(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateVariantRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateVariant(ctx, &pb.CreateVariantRequest{
		Name:        body.Name,
		Price:       body.Price,
		Stock:       body.Stock,
		Description: body.Description,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
