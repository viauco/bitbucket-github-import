package main

type BitbucketRepositoriesResponse struct {
	Next         string                `json:"next"`
	Repositories []BitbucketRepository `json:"values"`
}

type BitbucketRepository struct {
	SCM         string `json:"scm"`
	Name        string `json:"name"`
	Private     bool   `json:"is_private"`
	Description string `json:"description"`
	Links       struct {
		HTML struct {
			HREF string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}
