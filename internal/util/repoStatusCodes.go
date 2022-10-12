package util

type RepoStatusCode int8

const (
	RepoStatusSuccess  = 0
	RepoStatusNotFound = 1
	RepoStatusError    = 2
	RepoStatusExpired  = 3
)