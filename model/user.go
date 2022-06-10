package model

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Pwd       string `json:"pwd"`
	Create_ts int64  `json:"create_ts"`
	Update_ts int64  `json:"update_ts"`
}
