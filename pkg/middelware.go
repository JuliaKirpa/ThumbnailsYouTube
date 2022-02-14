package pkg

import "net/url"

func ParceURL(argument []string) (map[string]string, error) {
	response := make(map[string]string)

	for _, value := range argument {
		url, err := url.Parse(value)
		if err != nil {
			return nil, err
		}
		q := url.Query()
		id := q["v"]

		response[id[0]] = "https://img.youtube.com/vi/" + id[0] + "/0.jpg"
	}

	return response, nil
}
