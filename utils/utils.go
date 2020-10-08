package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/google/go-querystring/query"
)

var httpClient = &http.Client{}

func HttpGet(address string, v interface{}) ([]byte, error) {
	urlVal, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	values, _ := query.Values(v)
	urlVal.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlVal.String(), nil)
	if err != nil {
		return nil, err
	}
	a := req.URL.RequestURI()
	fmt.Println("URI:", a)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return nil, fmt.Errorf("statusCode: %d", statusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Sign(params map[string]string, appSecret string) string {
	var b strings.Builder

	keyList := make([]string, 0, len(params))
	for k := range params {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, key := range keyList {
		if val, ok := params[key]; ok {
			b.WriteString(fmt.Sprintf("%s%s", key, val))
		}
	}
	text := fmt.Sprintf("%s%s%s", appSecret, b.String(), appSecret)
	hash := md5.New()
	hash.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}
