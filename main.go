package main

import (
	"api-wrapper/wrapper"
	"encoding/json"
	"fmt"
)

func main() {
	// DEVELOPER WILL DO THE FOLLOWING:
	client := wrapper.NewClient(
		"client_id_pSdgBAaNfvRCcFmqK682tIhXOyD5iEV0rJx3nskQ|client_secret_7GfSOVxD5seYHRC8o2X0grlvZtd9azmhMpn3U1Ju",
		"127.0.0.1",
		"|5beb505292571b00a14dd31a",
	)

	data := client.GetUsers(map[string]interface{}{
		"per_page": 3,
		"page":     2,
	})
	// data := client.GetUser("5bec6ebebaabfc00ab168fa0", false)
	// var newUserData = string(`{
	// 	"logins": [
	// 		{
	// 			"email": "test4@synapsefi.com"
	// 		}
	// 	],
	// 	"phone_numbers": [
	// 		"901.111.1111",
	// 		"test@synapsefi.com"
	// 	],
	// 	"legal_names": [
	// 		"Test User"
	// 	],
	// 	"extra": {
	// 		"supp_id": "122eddfgbeafrfvbbb",
	// 		"cip_tag":1,
	// 		"is_business": false
	// 	}
	// }`)
	// data := client.CreateUser(newUserData)
	// data := client.GetClientTransactions()
	// data := client.GetSubscriptions()
	// data := client.CreateSubscription(string(`{
	// 	"scope": [
	// 		"USERS|POST",
	// 		"USER|PATCH",
	// 		"NODES|POST",
	// 		"NODE|PATCH",
	// 		"TRANS|POST",
	// 		"TRAN|PATCH"
	// 	],
	// 	"url": "https://requestb.in/zp216zzp"
	// }`))
	// data := client.GetSubscription("5bf2665a7e874a00a9cb32df")
	// data := client.UpdateSubscription("5bf2665a7e874a00a9cb32df", string(`{
	// 	"is_active": true
	// }`))
	// data := client.GetNodes("5bec6ebebaabfc00ab168fa0")
	// data := client.GetInstitutions()
	// data := client.GetPublicKey("OAUTH|POST")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	payload, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
	}

	// fmt.Println(data["UserList"])
	fmt.Println(string(payload))
}
