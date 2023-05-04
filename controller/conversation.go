package controller

import (
	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	conversationUseCase ConversationUseCase
	organizationUseCase OrganizationUseCase
}

func NewConversationController(c ConversationUseCase, o OrganizationUseCase) *ConversationController {
	return &ConversationController{
		conversationUseCase: c,
		organizationUseCase: o,
	}
}

type NewReportRequest struct {
	Report  string `json:"report"`
}

func (c ConversationController) SubmitReport(ctx *gin.Context) (interface{}, error){
	userID := "dfdsafdfafdfafda" //temporary
	organizationId := "12345678"
	var req NewReportRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	mvv, err := c.organizationUseCase.GetMVV(ctx, organizationId)
	if err != nil {
		return nil, err
	}
	resp, err := c.conversationUseCase.SendReport(ctx, userID, mvv, req.Report)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
