package counter

import "module06/internal/app/services/post"

// PostCount - count posts
func PostCount(client post.Client) (int, error) {
	posts, err := client.GetList()
	if err != nil {
		return 0, err
	}

	return len(posts), nil
}
