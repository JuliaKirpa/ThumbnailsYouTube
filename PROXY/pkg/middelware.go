package pkg

import "net/url"

func ParceURL(argument string) (string, string, error) {

	url, err := url.Parse(argument)
	if err != nil {
		return "", "", err
	}
	q := url.Query()
	id := q["v"]

	return id[0], "https://img.youtube.com/vi/" + id[0] + "/0.jpg", nil
}
