package oc

import "github.com/google/go-github/github"

func LaunchMergeTests(repo *github.Repository)       {}
func LaunchAcceptTests(repo *github.Repository)      {}
func LaunchUnitTests(repo *github.Repository)        {}
func LaunchIntegrationTests(repo *github.Repository) {}
