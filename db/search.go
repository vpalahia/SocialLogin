package db

import (
	"fmt"

	"github.com/vpalahia/SocialLogin/types"
)

func Search(email string) ([]types.User, error) {
	var searchedUserDetails []types.User
	fetchDetailsQuery := fmt.Sprintf("SELECT id, name, email, meta FROM loggedinuser WHERE email like '%%%s%%'", email)
	fmt.Println(fetchDetailsQuery)
	err := db.Select(&searchedUserDetails, fetchDetailsQuery)
	if err != nil {
		fmt.Println(err)
		return []types.User{}, err
	}

	return searchedUserDetails, nil
}
