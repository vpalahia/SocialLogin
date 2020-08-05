package db

import (
	"encoding/json"
	"fmt"

	"github.com/vpalahia/SocialLogin/types"
)

//InsertIntoDB function to insert into database
func InsertIntoDB(userinfo types.ThirdPartyResponse) error {
	meta, err := json.Marshal(userinfo.MetaData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if userinfo.Name == "" {
		fmt.Println("Name is not issued from third party")
	} else if userinfo.Id == 0 {
		fmt.Println("Email is not issued from third party")
	}

	insertQuery := `INSERT INTO loggedinuser (id, name, email, meta) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(insertQuery, userinfo.Id, userinfo.Name, userinfo.Email, string(meta))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
