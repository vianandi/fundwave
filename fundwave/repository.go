package fundwave

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Fundwave, error)
	FindByUserID(userID int) ([]Fundwave, error)
	FindByID(id int) (Fundwave, error)
	Save(fundwave Fundwave) (Fundwave, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Fundwave, error) {
	var fundwave []Fundwave

	err := r.db.Preload("FundwaveImages", "fundwave_images.is_primary = 1")Find(&fundwave).Error
	if err != nil {
		return fundwave, err
	}

	return fundwave, nil
}

func (r *repository) FindByUserID(userID int) ([]Fundwave, error) {
	var fundwave []Fundwave

	err := r.db.Where("user_id = ?", userID).Preload("FundwaveImages", "fundwave_images.is_primary = 1").Find(&fundwave).Error
	if err != nil {
		return fundwave, err
	}

	return fundwave, nil
}
