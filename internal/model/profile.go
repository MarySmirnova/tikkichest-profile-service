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

func (p *Profile) Update(updateData *ProfileUpdate) {
	if updateData.Username != nil {
		p.Username = *updateData.Username
	}

	if updateData.Name != nil {
		p.Name = *updateData.Name
	}

	if updateData.Email != nil {
		p.Email = *updateData.Email
	}

	if updateData.Phone != nil {
		p.Phone = *updateData.Phone
	}

	if updateData.HashPassword != nil {
		p.HashPassword = *updateData.HashPassword
	}

	if updateData.Location != nil {
		if updateData.Location.Country != nil {
			p.Location.Country = *updateData.Location.Country
		}
		if updateData.Location.Town != nil {
			p.Location.Town = *updateData.Location.Town
		}
	}

	p.ChangeTime = updateData.ChangeTime
}
