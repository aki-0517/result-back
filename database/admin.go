package database

import (
	"errors"
	"fmt"
	"go_blog/crypto"

	"gorm.io/gorm"
)

func init() {
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(Admin{})
}

func Signup(adminID, password string) (*Admin, error) {
	admin := Admin{}
	Db.Where("admin_id = ?", adminID).First(&admin)

	if admin.ID != 0 {
		err := errors.New("同一名のAdminIdが既に登録されています。")
		fmt.Println(err)
		return nil, err
	}

	encryptPw, err := crypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}
	admin = Admin{AdminID: adminID, HashedPassword: encryptPw}
	Db.Create(&admin)
	return &admin, nil
}

func Login(adminID, password string) (*Admin, error) {
	admin := Admin{}
	Db.Where("admin_id = ?", adminID).First(&admin)
	if admin.ID == 0 {
		err := errors.New("AdminIdが一致するユーザーが存在しません。")
		fmt.Println(err)
		return nil, err
	}

	err := crypto.CompareHashAndPassword(admin.HashedPassword, password)
	if err != nil {
		fmt.Println("パスワードが一致しませんでした。：", err)
		return nil, err
	}

	return &admin, nil
}
