package db

import (
	"fmt"

	"github.com/vpalahia/SocialLogin/types"
)

func User(id int) (types.User, error) {
	var userDetails []types.User

	fetchDetailsQuery := `SELECT id, name, email, meta FROM loggedinuser WHERE id = $1`
	err := db.Select(&userDetails, fetchDetailsQuery, id)
	if err != nil {
		fmt.Println(err)
		return types.User{}, err
	}
	if len(userDetails) > 0 {
		return userDetails[0], nil
	}

	return types.User{}, nil

}
