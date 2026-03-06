package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com"

// Client - interface for post client
type Client interface {
	GetList() ([]Post, error)
}

type postClient struct{}

// NewClient - create client
func NewClient() Client {
	return &postClient{}
}

func (c *postClient) GetList() ([]Post, error) {
	client := http.Client{Timeout: time.Second * 5}

	res, err := client.Get(url + "/posts")
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal []Post: %w", err)
	}

	return posts, nil
}

// Post - struct for marshal/unmarshal get post json data
type Post struct {
	ID     int
	Title  string
	Body   string
	UserID int
}
