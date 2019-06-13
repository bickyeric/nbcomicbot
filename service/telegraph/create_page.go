package telegraph

import (
	"encoding/json"
	"os"

	"github.com/bickyeric/arumba/connection"
)

type createPageParam struct {
	AccessToken string        `json:"access_token"`
	Title       string        `json:"title"`
	Author      string        `json:"author_name"`
	AuthorURL   string        `json:"author_url"`
	Content     []interface{} `json:"content"`
}

type attrs struct {
	Href string `json:"href,omitempty"`
	Src  string `json:"src,omitempty"`
}

type node struct {
	Tag      string `json:"tag"`
	Attrs    *attrs `json:"attrs,omitempty"`
	Children []node `json:"children,omitempty"`
}

type createPageResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Path        string `json:"path"`
		URL         string `json:"url"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      string `json:"author_name"`
		AuthorURL   string `json:"author_url"`
		Views       int    `json:"views"`
		CanEdit     bool   `json:"can_edit"`
	} `json:"result"`
}

type PageCreator interface {
	Perform(source, title string, images []string) (string, error)
}

type createPage struct {
	network connection.NetworkInterface
	token   string
}

// NewCreatePage ...
func NewCreatePage() PageCreator {
	return createPage{
		network: connection.NewNetwork("https://api.telegra.ph/"),
		token:   os.Getenv("TELEGRAPH_ACCESS_TOKEN"),
	}
}

// Perform ...
func (cp createPage) Perform(source, title string, images []string) (string, error) {
	param := cp.buildParam(source, title, images)
	jsonParam, _ := json.Marshal(param)

	body, err := cp.network.POST("createPage", jsonParam)
	if err != nil {
		return "", err
	}

	createPageResponse := createPageResponse{}
	if err := json.Unmarshal(body, &createPageResponse); err != nil {
		return "", err
	}

	return createPageResponse.Result.URL, nil
}

func (cp createPage) buildParam(source, title string, images []string) (param createPageParam) {
	param.AccessToken = cp.token
	param.Title = title
	param.Author = source
	param.AuthorURL = "https://t.me/arumba_channel"

	i := 1
	for _, link := range images {
		param.Content = append(param.Content, node{
			Tag: "figure",
			Children: []node{
				node{
					Tag: "img",
					Attrs: &attrs{
						Src: link,
					},
				},
			},
		})
		if i == 3 {
			param.Content = append(param.Content, cp.footer())
			i = 0
		} else {
			i++
		}
	}

	return param
}

func (createPage) footer() interface{} {
	return map[string]interface{}{
		"tag": "p",
		"children": []interface{}{
			"Check out the ",
			map[string]interface{}{
				"tag": "a",
				"attrs": map[string]interface{}{
					"href": "https://t.me/arumba_bot",
				},
				"children": []interface{}{
					map[string]interface{}{
						"tag":      "strong",
						"children": []string{"@Arumba"},
					},
				},
			},
			" for more comics and episodes...",
		},
	}
}
