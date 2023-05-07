package persistence

import (
	"encoding/json"
	"errors"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/persistence/model"
	"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
)

type ReportPersistence struct {
	*redis.Client
}

func NewReportPersistence(cli *redis.Client) *ReportPersistence {
	return &ReportPersistence{cli}
}

func (p ReportPersistence) GetReports(ctx context.Context, organizationID string) ([]*entity.Report, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var records []*model.Report
	if err := db.Preload("User").Where("organization_id = ?", organizationID).Find(&records).Error; err != nil {
		return nil, err
	}
	var reports []*entity.Report
	for _, report := range records {
		reports = append(reports, report.ToEntity())
	}
	return reports, nil
}

func (p ReportPersistence) GetReport(ctx context.Context, organizationID *string, reportID string) (*entity.Report, error) {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	var record *model.Report
	if organizationID != nil {
		db = db.Where("organization_id = ?", *organizationID)
	}
	if err := db.Preload("User").First(&record, "id = ?", reportID).Error; err != nil {
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

func (p ReportPersistence) UpdateReviewBody(ctx context.Context, reportID string, reviewBody string) error {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	err := db.Model(&model.Report{}).Where("id = ?", reportID).Update("review_body", reviewBody).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.NewError(http.StatusNotFound, err)
		}
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return entity.NewError(http.StatusConflict, err)
		}
		return err
	}
	return nil
}

func (p ReportPersistence) DispatchReport(ctx context.Context, reportID, body string) error {
	payload, err := json.Marshal(&driver.Message{
		ID:   reportID,
		Body: body,
	})
	if err != nil {
		return err
	}
	return p.Client.Publish(ctx, driver.JobQueueKey, string(payload)).Err()
}
