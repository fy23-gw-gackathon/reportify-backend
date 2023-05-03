package persistence

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"time"
)

type OrganizationPersistence struct{}

func NewOrganizationPersistence() *OrganizationPersistence {
	return &OrganizationPersistence{}
}

type Organization struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (p OrganizationPersistence) GetOrganizations(ctx context.Context, limit *int, offset *int) ([]*entity.Organization, error) {
	var records []*Organization
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
