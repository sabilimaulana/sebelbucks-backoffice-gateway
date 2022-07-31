package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/auth/pb"
)

type RegisterAdminRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterAdmin(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterAdminRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.RegisterAdmin(context.Background(), &pb.RegisterAdminRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(res.Status), &res)
}
