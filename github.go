package main

type GitHubRepository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

type GitHubError struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Code     string `json:"code"`
		Field    string `json:"field"`
	} `json:"errors"`
}

type GitHubImport struct {
	VCSURL      string `json:"vcs_url"`
	VCS         string `json:"vcs"`
	VCSUsername string `json:"vcs_username"`
	VCSPassword string `json:"vcs_password"`
}
