package persistence

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	"reportify-backend/entity"
	"reportify-backend/infrastructure/driver"
	"reportify-backend/infrastructure/persistence/model"
)

type ReportPersistence struct{}

func NewReportPersistence() *ReportPersistence {
	return &ReportPersistence{}
}

func (p ReportPersistence) GetReports(ctx context.Context, organizationCode, userID string) ([]*entity.Report, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var records []*model.Report
	if err := db.Preload("User.Organizations", "code = ?", organizationCode).Find(&records).Error; err != nil {
		return nil, err
	}
	var reports []*entity.Report
	for _, report := range records {
		reports = append(reports, report.ToEntity())
	}
	return reports, nil
}

func (p ReportPersistence) GetReport(ctx context.Context, organizationCode, reportID string) (*entity.Report, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Report
	if err := db.Preload("User.Organizations", "code = ?", organizationCode).First(&record, "id = ?", reportID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.NewError(http.StatusNotFound, err)
		}
		return nil, err
	}
	return record.ToEntity(), nil
}

func (p ReportPersistence) CreateReport(ctx context.Context, organizationID, userID string, body string, task []entity.Task) (*entity.Report, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	reportID := generateID().String()
	var taskRecords []*model.Task
	for _, taskRecord := range task {
		taskRecords = append(taskRecords, &model.Task{
			ID:         generateID().String(),
			Name:       taskRecord.Name,
			ReportID:   reportID,
			StartedAt:  taskRecord.StartedAt,
			FinishedAt: taskRecord.FinishedAt,
		})
	}
	record := &model.Report{
		ID:             reportID,
		Body:           body,
		ReviewBody:     nil,
		UserID:         userID,
		OrganizationID: organizationID,
		Tasks:          taskRecords,
	}
	if err := db.Create(record).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return nil, entity.NewError(http.StatusConflict, err)
		}
	}
	return record.ToEntity(), nil
}
