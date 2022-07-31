package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/auth/pb"
)

type LoginAdminRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginAdmin(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginAdminRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.LoginAdmin(context.Background(), &pb.LoginAdminRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
