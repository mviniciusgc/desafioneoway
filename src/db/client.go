package db

type ClientDB struct {
}

func InitializeClientBD() IClientDB {
	return &ClientDB{}
}
