package fundwave

import "gorm.io/gorm"

// Repository interface defines the methods expected
type Repository interface {
	FindAll() ([]Fundwave, error)
	FindByUserID(userID int) ([]Fundwave, error)
	FindByID(ID int) (Fundwave, error)
	Save(fundwave Fundwave) (Fundwave, error) // Add this method to the interface
}

// repository struct implements the Repository interface
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// FindAll retrieves all fundwaves
func (r *repository) FindAll() ([]Fundwave, error) {
	var fundwave []Fundwave
	err := r.db.Preload("FundwaveImages", "fundwave_images.is_primary = 1").Find(&fundwave).Error
	if err != nil {
		return fundwave, err
	}
	return fundwave, nil
}

// FindByUserID retrieves fundwaves by user ID
func (r *repository) FindByUserID(userID int) ([]Fundwave, error) {
	var fundwave []Fundwave
	err := r.db.Where("user_id = ?", userID).Preload("FundwaveImages", "fundwave_images.is_primary = 1").Find(&fundwave).Error
	if err != nil {
		return fundwave, err
	}
	return fundwave, nil
}

// FindByID retrieves a fundwave by ID
func (r *repository) FindByID(ID int) (Fundwave, error) {
	var fundwave Fundwave
	err := r.db.Preload("User").Preload("FundwaveImages").Where("id = ?", ID).Find(&fundwave).Error
	if err != nil {
		return fundwave, err
	}
	return fundwave, nil
}

// Save adds or updates a fundwave in the database
func (r *repository) Save(fundwave Fundwave) (Fundwave, error) {
	err := r.db.Save(&fundwave).Error
	if err != nil {
		return fundwave, err
	}
	return fundwave, nil
}
