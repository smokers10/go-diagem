package etc

import (
	"io/ioutil"
	"net/http"
)

type ReqStruct struct {
	Endpoint string
	Method   string
	Key      string
	Payload  string
}

func ReqAPI(rs *ReqStruct) (string, error) {
	req, err := http.NewRequest(rs.Method, rs.Endpoint, nil)
	if err != nil {
		return "", err
	}

	if rs.Key != "" {
		req.Header.Add("key", rs.Key)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
