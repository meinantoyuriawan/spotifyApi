package helper

import (
	b64 "encoding/base64"
)

func GenerateBasicToken() string {
	CLIENT_ID := " "
	CLIENT_SECRET := " "

	strSecret := CLIENT_ID + ":" + CLIENT_SECRET

	sEnc := b64.StdEncoding.EncodeToString([]byte(strSecret))

	Authorization := "Basic " + string(sEnc)

	return Authorization
}
