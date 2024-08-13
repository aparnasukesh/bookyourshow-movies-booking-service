package theatres

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateTheaterType(ctx context.Context, theaterType TheaterType) error
	DeleteTheaterTypeByID(ctx context.Context, id int) error
	DeleteTheaterTypeByName(ctx context.Context, name string) error
	FindTheatertypeByName(ctx context.Context, name string) (*TheaterType, error)
	GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error)
	GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error)
	UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error
	ListTheaterTypes(ctx context.Context) ([]TheaterType, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindTheatertypeByName(ctx context.Context, name string) (*TheaterType, error) {
	theaterType := &TheaterType{}
	res := r.db.Where("theater_type_name ILIKE ?", name).First(&theaterType)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		} else if res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return theaterType, nil

}

func (r *repository) CreateTheaterType(ctx context.Context, theaterType TheaterType) error {
	if err := r.db.Create(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterTypeByID(ctx context.Context, id int) error {
	theaterType := &TheaterType{}
	if err := r.db.Where("id =?", id).Delete(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterTypeByName(ctx context.Context, name string) error {
	theaterType := &TheaterType{}
	if err := r.db.Where("theater_type_name ILIKE ?", name).Delete(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error) {
	theatertype := TheaterType{}
	if err := r.db.Where("id =?", id).First(&theatertype).Error; err != nil {
		return nil, err
	}
	return &theatertype, nil
}

func (r *repository) GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error) {
	theaterType := &TheaterType{}
	if err := r.db.Where("theater_type_name ILIKE ?", name).First(&theaterType).Error; err != nil {
		return nil, err
	}
	return theaterType, nil
}

func (r *repository) ListTheaterTypes(ctx context.Context) ([]TheaterType, error) {
	theaterTypes := []TheaterType{}
	if err := r.db.Find(&theaterTypes).Error; err != nil {
		return nil, err
	}
	return theaterTypes, nil
}

func (r *repository) UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&TheaterType{}).Where("id = ?", id).Updates(theaterType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
