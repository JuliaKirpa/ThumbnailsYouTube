package pkg

import "net/url"

func ParceURL(argument []string) (string, error) {
	url, err := url.Parse(argument[1])
	if err != nil {
		return "", err
	}

	q := url.Query()
	id := q["v"]
	return "https://img.youtube.com/vi/" + id[0] + "/0.jpg", nil
}