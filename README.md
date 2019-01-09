# Bitbucket to GitHub import

A very basic tool to import your private repositories from [Bitbucket](https://developer.atlassian.com/bitbucket/api/2/reference/) to [GitHub](https://developer.github.com/v3/).
Check the code and adjust to your needs.

At the time of writing, the [GitHub Source Imports API](https://developer.github.com/v3/migrations/source_imports/) is in public preview, it can change any time.

## Install
go get github.com/andreiavrammsd/bitbucket-github-import

## Usage
Create a `.env` file from [template](.env.dist).

```
go run *.go
```
