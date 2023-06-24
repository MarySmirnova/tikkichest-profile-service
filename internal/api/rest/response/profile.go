package response

import "github.com/MarySmirnova/tikkichest-profile-service/internal/model"

type Profile struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	Town      string `json:"town"`
	IsCreator bool   `json:"is_creator"`
}

type ProfileShortDetailed struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsCreator bool   `json:"is_creator"`
}

type ProfilePage struct {
	Profiles []ProfileShortDetailed `json:"profiles"`
	Page     Page                   `json:"page"`
}

func ProfileFromDBModel(profile *model.Profile) Profile {
	if profile == nil {
		return Profile{}
	}
	return Profile{
		Username:  profile.Username,
		Name:      profile.Name,
		Email:     profile.Email,
		Phone:     profile.Phone,
		Country:   profile.Location.Country,
		Town:      profile.Location.Town,
		IsCreator: profile.IsCreator,
	}
}

func ShortProfileFromDBModel(profile *model.Profile) ProfileShortDetailed {
	if profile == nil {
		return ProfileShortDetailed{}
	}

	return ProfileShortDetailed{
		Username:  profile.Username,
		Name:      profile.Name,
		Email:     profile.Email,
		IsCreator: profile.IsCreator,
	}
}

func ProfilePageFromDBModel(profiles []*model.Profile, limit int, nextFrom string) ProfilePage {
	respProfiles := make([]ProfileShortDetailed, 0, len(profiles))
	for _, p := range profiles {
		respProfiles = append(respProfiles, ShortProfileFromDBModel(p))
	}

	return ProfilePage{
		Profiles: respProfiles,
		Page:     createPage(len(respProfiles), limit, nextFrom),
	}
}
