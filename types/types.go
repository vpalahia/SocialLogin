package types

type ThirdPartyResponse struct {
	Id       uint64   `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Email    string   `json:"email" db:"email"`
	MetaData MetaData `json:"meta" db:"meta"`
}

type MetaData struct {
	Github   bool
	Linkedin bool
	Twitter  bool
}

type User struct {
	Id       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	MetaData string `json:"meta" db:"meta"`
}

type LinkedinResp struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
}

type LinkedUser struct {
	Elements []Handle `json:"elements"`
}

type Handle struct {
	Handle Email `json:"handle~"`
}
type Email struct {
	Email string `json:"emailAddress"`
}
