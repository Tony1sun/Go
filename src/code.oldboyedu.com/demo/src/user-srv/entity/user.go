package entity

import "code.oldboyedu.com/demo/src/share/pb"

// 结构体定义
type User struct {
	Id      int32  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
}

func (u *User) ToPorotoUser() *pb.User {
	return &pb.User{
		Id:      u.Id,
		Name:    u.Name,
		Address: u.Address,
		Phone:   u.Phone,
	}
}
