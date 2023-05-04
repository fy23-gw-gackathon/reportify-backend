package persistence

import (
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type OrganizationPersistence struct{}

func NewOrganizationPersistence() *OrganizationPersistence {
	return &OrganizationPersistence{}
}

func (p OrganizationPersistence) GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error) {
	var records []*model.Organization
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if limit != nil {
		db = db.Limit(*limit)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	var organizations []*entity.Organization
	for _, record := range records {
		organizations = append(organizations, &entity.Organization{
			ID:   record.ID,
			Name: record.Name,
		})
	}
	return organizations, nil
}

func (p OrganizationPersistence) GetMVV(ctx context.Context, organizationId string) (*entity.MVV, error) {
	var record model.MVV
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err :=  db.Table("mvvs").Where("organization_id = ?", organizationId).First(&record).Error; err != nil {
		return nil, err
	}
	return &entity.MVV{
		ID: record.ID,
		OrganizationID: record.OrganizationID,
		Mission: record.Mission,
		Vision: record.Vision,
		Value: record.Value,
	}, nil
}
