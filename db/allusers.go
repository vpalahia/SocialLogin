package db

import (
	"fmt"

	"github.com/vpalahia/SocialLogin/types"
)

func Allusers() ([]types.User, error) {
	var userList []types.User

	fetchDetailsQuery := `SELECT id, name, email, meta FROM loggedinuser`
	err := db.Select(&userList, fetchDetailsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// for i, v := range userList {
	// 	fetchMetaQuery := `SELECT meta FROM loggedinuser WHERE email = $1`
	// 	row := db.QueryRow(fetchMetaQuery, v.Email)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, errors.New("Unable to fetch meta data")
	// 	}
	// 	var str string
	// 	err = row.Scan(&str)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, err
	// 	}
	// 	fmt.Println(str)
	// 	userList[i].MetaData = str
	// }

	return userList, nil
}
