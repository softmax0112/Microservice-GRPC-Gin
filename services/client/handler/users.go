package handler

import (
	"context"
	"fmt"
	"log"
	usersv "micr-go/services/users/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler struct{}

var user UserHandler

func (u *UserHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":3000"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return conn
}

func FindAllUser(c *gin.Context) {
	conn := user.connGrpc()
	defer conn.Close()

	u := usersv.NewUsersClient(conn)
	response, err := u.FindAll(context.Background(), &usersv.FindAllRequest{Page: "1", PerPage: "2"})
	if err != nil {
		log.Fatalf("Error when calling User FindAllRequest: %s", err)
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": response,
	})

}