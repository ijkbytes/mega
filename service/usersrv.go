package service

import "github.com/ijkbytes/mega/model"

type userService struct {
}

var User = &userService{}

func (srv *userService) AddUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (srv *userService) UpdateUser(user *model.User) error {
	if err := db.Update(user).Error; err != nil {
		return err
	}

	return nil
}

func (srv *userService) GetUserByName(name string) *model.User {
	ret := &model.User{}
	if err := db.Where("`user_name` = ?", name).First(ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *userService) GetUser(id uint64) *model.User {
	ret := &model.User{}
	if err := db.Where("`id` = ?", id).First(ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *userService) UsersCount() (count int, err error) {
	err = db.Model(&model.User{}).Count(&count).Error
	return
}
