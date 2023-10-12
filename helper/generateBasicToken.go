package helper

import (
	b64 "encoding/base64"
)

func GenerateBasicToken() string {
	CLIENT_ID := GetClientID()
	CLIENT_SECRET := GetClientSecret()

	strSecret := CLIENT_ID + ":" + CLIENT_SECRET

	sEnc := b64.StdEncoding.EncodeToString([]byte(strSecret))

	Authorization := "Basic " + string(sEnc)

	return Authorization
}
