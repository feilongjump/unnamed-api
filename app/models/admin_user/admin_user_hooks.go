package admin_user

import (
	"unnamed-api/pkg/hash"

	"gorm.io/gorm"
)

func (adminUser *AdminUser) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(adminUser.Password) {
		adminUser.Password = hash.BcryptHash(adminUser.Password)
	}

	return
}

// func (adminUser *AdminUser) BeforeCreate(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) AfterCreate(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) AfterUpdate(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) AfterSave(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) BeforeDelete(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) AfterDelete(tx *gorm.DB) (err error) {}
// func (adminUser *AdminUser) AfterFind(tx *gorm.DB) (err error) {}
