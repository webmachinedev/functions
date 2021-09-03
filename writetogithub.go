package functions

import (
	"context"
	"errors"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v37/github"
)


func WriteToGithub(owner, repo, branch string, changes map[string]string, commitmessage, githubkey string) error {
	ctx := context.Background()
	token := oauth2.Token{AccessToken: githubkey}
	ts := oauth2.StaticTokenSource(&token)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ref, _, err := client.Git.GetRef(ctx, owner, repo, "heads/"+branch)
	if err != nil {
		return err
	}

	lastcommit, _, err := client.Git.GetCommit(ctx, owner, repo, *ref.Object.SHA)
	if err != nil {
		return err
	}

	var treeentries []*github.TreeEntry

	for path, contents := range changes {
		if contents == "" {
			return errors.New("delete not implemented")
		} else {
			blob := &github.Blob{
				Content: &contents,
			}
			blob, _, err = client.Git.CreateBlob(ctx, owner, repo, blob)
			if err != nil {
				return err
			}
	
			treeentrytype := "blob"
			mode := "100644"
			treeentry := github.TreeEntry{
				Path: &path,
				Type: &treeentrytype,
				Mode: &mode,
				SHA:  blob.SHA,
			}
			treeentries = append(treeentries, &treeentry)
		}
	}

	tree, _, err := client.Git.CreateTree(ctx, owner, repo, *lastcommit.Tree.SHA, treeentries)
	if err != nil {
		return err
	}

	commit := &github.Commit{
		Parents: []*github.Commit{lastcommit},
		Message: &commitmessage,
		Tree:    tree,
	}
	commit, _, err = client.Git.CreateCommit(ctx, owner, repo, commit)
	if err != nil {
		return err
	}

	refname := "heads/" + branch
	ref = &github.Reference{
		Ref: &refname,
		Object: &github.GitObject{
			SHA: commit.SHA,
		},
	}

	_, _, err = client.Git.UpdateRef(ctx, owner, repo, ref, false)

	return err
}
