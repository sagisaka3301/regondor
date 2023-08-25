package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// サインアップのエンドポイントで新しくユーザーを作成したとき、その情報をクライアントにレスポンスで返す際の型を定義。
type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
	Name  string `json:"name"`
}

type MypageResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}
