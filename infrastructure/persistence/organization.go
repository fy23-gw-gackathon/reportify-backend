package persistence

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"
)

type OrganizationPersistence struct{}

func NewOrganizationPersistence() *OrganizationPersistence {
	return &OrganizationPersistence{}
}

func (p OrganizationPersistence) GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error) {
	var record *model.User
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Preload("Organizations").First(&record, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	var organizations []*entity.Organization
	for _, organization := range record.Organizations {
		organizations = append(organizations, organization.ToEntity())
	}
	return organizations, nil
}

func (p OrganizationPersistence) GetOrganization(ctx context.Context, organizationID string) (*entity.Organization, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Organization
	if err := db.Where("id = ?", organizationID).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	return record.ToEntity(), nil
}

func (p OrganizationPersistence) UpdateOrganization(ctx context.Context, organizationID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Organization
	if err := db.Where("id = ?", organizationID).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	record.Name = organizationName
	record.Code = organizationCode
	record.Mission = mission
	record.Vision = vision
	record.Value = value
	if err := db.Save(&record).Error; err != nil {
		return nil, err
	}
	return record.ToEntity(), nil
}
