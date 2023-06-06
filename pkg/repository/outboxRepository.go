package repository

import "github.com/erolkaldi/agency/pkg/models"

func (rp *Repository) SaveOutbox(outbox *models.OutBox) (*models.OutBox, error) {
	err := rp.db.Save(&outbox).Error
	return outbox, err
}
func (rp *Repository) GetOutBoxItems() ([]models.OutBox, error) {
	var items []models.OutBox
	err := rp.db.Where(`record_state=0`).Find(&items).Error
	return items, err
}
