package response

import "fmt"

type Page struct {
	NextPage string `json:"next_page,omitempty"`
	Count    int    `json:"count"`
}

func createPage(countObjects, limit int, nextFrom string) Page {
	if countObjects < limit {
		return Page{
			Count: countObjects,
		}
	}

	return Page{
		NextPage: fmt.Sprintf("/profile?limit=%d&from=%s", limit, nextFrom),
		Count:    countObjects,
	}
}
