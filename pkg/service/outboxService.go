package service

import (
	"fmt"
	"time"

	"github.com/erolkaldi/agency/pkg/models"
	"github.com/erolkaldi/agency/pkg/repository"
)

type OutboxService struct {
	repository *repository.Repository
	smtp       *models.Smtp
}

func InitializeOutboxService(repository *repository.Repository, smtp *models.Smtp) *OutboxService {
	return &OutboxService{repository: repository, smtp: smtp}
}
func CreateOutBox(message string, rp *repository.Repository) (*models.OutBox, error) {
	outbox := &models.OutBox{Data: message, RecordState: models.PendingDelivery, BoxType: models.Mail, CreatedOn: time.Now()}
	return rp.SaveOutbox(outbox)
}

func GetOutboxItems(rp *repository.Repository) ([]models.OutBox, error) {
	return rp.GetOutBoxItems()
}

func (outboxService *OutboxService) OutboxPooling(rp *repository.Repository, smtp *models.Smtp) {
	for true {
		items, err := GetOutboxItems(rp)
		fmt.Println("Outbox item count:", len(items))
		if err == nil {
			for i := 0; i < len(items); i++ {
				items[i].RecordState = models.Processing
				rp.SaveOutbox(&items[i])
				if items[i].BoxType == models.Mail {
					var email models.Email
					StringToJson(items[i].Data, &email)
					err := SendEmail(email, *outboxService.smtp)
					if err != nil {
						items[i].AttemptsCount++
						items[i].Error = err.Error()
						items[i].LastAttemptOn = time.Now()
						items[i].RecordState = models.Delivered
						rp.SaveOutbox(&items[i])
						continue
					}
					items[i].AttemptsCount++
					items[i].ProcessedOn = time.Now()
					items[i].RecordState = models.Delivered
					rp.SaveOutbox(&items[i])
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
