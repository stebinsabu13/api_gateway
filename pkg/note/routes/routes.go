package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/note/pb"
)

type CreateNoteReqBody struct {
	Note string `json:"note" binding:"required"`
}

type DeleteReqBody struct {
	Id uint32 `json:"id" binding:"required"`
}

func CreateNote(c *gin.Context, notecli pb.NoteServiceClient) {
	var body CreateNoteReqBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	userid, _ := c.Get("userId")
	res, err := notecli.CreateNote(context.Background(), &pb.CreateNoteRequest{
		Uid:  userid.(int64),
		Note: body.Note,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
	}
	c.JSON(http.StatusOK, &res)
}

func ListAllNote(c *gin.Context, notecli pb.NoteServiceClient) {
	userid, _ := c.Get("userId")
	res, err := notecli.ListAllNote(context.Background(), &pb.ListAllNoteRequest{
		Id: userid.(int64),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
	}
	c.JSON(http.StatusOK, &res)
}

func DeleteNote(c *gin.Context, notecli pb.NoteServiceClient) {
	var body DeleteReqBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	userid, _ := c.Get("userId")
	res, err := notecli.DeleteNote(context.Background(), &pb.DeleteNoteRequest{
		Uid: userid.(int64),
		Id:  body.Id,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
	}
	c.JSON(http.StatusOK, &res)
}
