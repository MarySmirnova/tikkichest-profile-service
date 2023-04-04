package model

type Profile struct {
	Username     string
	Name         string
	Email        string
	Phone        string
	Location     Location
	HashPassword string
	IsCreator    bool
	CreateTime   int64
	ChangeTime   int64
}

type Location struct {
	Country string
	Town    string
}

type ProfileUpdate struct {
	Username     *string
	Name         *string
	Email        *string
	Phone        *string
	Location     *LocationUpdate
	HashPassword *string
	ChangeTime   int64
}

type LocationUpdate struct {
	Country *string
	Town    *string
}
