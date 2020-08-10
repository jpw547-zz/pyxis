package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jpw547/pyxis/server/structs"
)

// BusinessSearch executes a business search endpoint request against the Yelp Fusion API.
func BusinessSearch(latitude float64, longitude float64, params map[string][]string) {
	addr := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?latitude=%f&longitude=%f", latitude, longitude)

	if len(params) > 0 {
		for k, v := range params {
			addr += fmt.Sprintf("&%s=%s", k, strings.Join(v, ","))
		}
	}

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		fmt.Printf("failed to make yelp request: %s", err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("YELP_API_KEY")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("failed to execute yelp request: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	switch {
	case err != nil:
		return
	case resp.StatusCode != http.StatusOK:
		return
	case body == nil:
		return
	}

	var yResp structs.YelpResponse

	err = json.Unmarshal(body, &yResp)
	if err != nil {
		fmt.Printf("failed to unmarshal the response from yelp: %s", err)
		return
	}

	for _, b := range yResp.Businesses {
		fmt.Println(b.Name)
	}
}
