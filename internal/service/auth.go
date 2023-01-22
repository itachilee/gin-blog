package service

import "errors"

type AuthRequest struct {
	Id       int    `json:"id" from:"id" binding:"required"`
	Username string `json:"username" from:"username" binding:"required"`
}

func (svc *Service) CheckAuth(params *AuthRequest) error {
	auth, err := svc.dao.GetAuth(params.Id, params.Username)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
