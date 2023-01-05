package models

type UserBasic struct {
	Id        string `bson:"_id"` // mongodb uuid
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Gender    int    `bson:"gender"` // 0-unknown, 1-male, 2-female
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

// 定义函数，返回结构体对应的集合名称，define a func return struct UserBasic collection name
func (UserBasic) CollectionName() string {
	return "user_basic"
}
