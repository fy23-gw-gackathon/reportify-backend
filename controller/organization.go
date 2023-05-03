package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reportify-backend/entity"
)

type OrganizationController struct {
	OrganizationUseCase
}

func NewOrganizationController(organizationUseCase OrganizationUseCase) *OrganizationController {
	return &OrganizationController{organizationUseCase}
}

type GetOrganizationsQuery struct {
	Limit  *int `form:"limit"`
	Offset *int `form:"offset"`
}

type GetOrganizationsResponse struct {
	Organizations []*entity.Organization `json:"organizations"`
}

// GetOrganizations godoc
// @Summary       get organizations
// @Description    get organizations
// @Tags         Organization
// @Accept       json
// @Produce       json
// @Param        limit  query     string false  "limit"
// @Param        offset query     string false  "offset"
// @Success       200       {object}   GetOrganizationsResponse
// @Failure       400       {object}   entity.ErrorResponse
// @Failure       404       {object}   entity.ErrorResponse
// @Router       /organizations [get]
func (c *OrganizationController) GetOrganizations(ctx *gin.Context) (interface{}, error) {
	var query GetOrganizationsQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	if query.Limit == nil && query.Offset != nil {
		return nil, entity.NewError(http.StatusBadRequest, errors.New("cannot use offset query without limit query"))
	}
	orgs, err := c.OrganizationUseCase.GetOrganizations(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	return GetOrganizationsResponse{Organizations: orgs}, nil
}
