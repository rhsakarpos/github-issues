package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)
import "golang.org/x/oauth2"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "447004fb5ba373f29410096ccdfddf709c3ccbc3"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	//repos, _, _ := client.Repositories.List(ctx, "", nil)
	client.Repositories.List(ctx, "", nil)
	//fmt.Println(repos)

	var bearray = make([]string, 0)
	bearray = append(bearray, "YIG")
	bearray = append(bearray, "AWS")
	bearray = append(bearray, "GCP")
	bearray = append(bearray, "Azure")
	bearray = append(bearray, "HWS")

	var yiglabels = make([]string,2)
	label0 := "Daito"
	label1 := "YIG"
	yiglabels[0] = label0
	yiglabels[1] = label1
	//milestone := 1
	milestone := 4
	var yigowner = "sfzeng"
	fmt.Println(yigowner)
	var yigarray = make([]string, 0)
	yigarray = append(yigarray, "Put object/Put object range")
	yigarray = append(yigarray,  "Put bucket")
	yigarray = append(yigarray,  "Get object")
	yigarray = append(yigarray,  "Delete object")
	yigarray = append(yigarray, "Delete bucket")
	yigarray = append(yigarray, "Head bucket")
	yigarray = append(yigarray, "Head object")
	yigarray = append(yigarray, "Post object (Browser-Based Uploads Using POST)")
	yigarray = append(yigarray, "Put object part")
	yigarray = append(yigarray, "Put bucket policy")
	yigarray = append(yigarray, "Get bucket policy")
	yigarray = append(yigarray, "Delete bucket policy")
	yigarray = append(yigarray, "Put object ACL")
	yigarray = append(yigarray, "Get object ACL")
	yigarray = append(yigarray, "Put bucket ACL")
	yigarray = append(yigarray, "Get bucket ACL")
	yigarray = append(yigarray, "Copy object")
	yigarray = append(yigarray, "New multipart upload")
	yigarray = append(yigarray, "Abort multipart upload")
	yigarray = append(yigarray, "Complete multipart upload")
	yigarray = append(yigarray, "List bucket multipart")
	yigarray = append(yigarray, "Enable Redis")

	var cnlabels = make([]string,1)
	cnlabel0 := "Daito"
	cnlabels[0] = cnlabel0
	var cnowner = "sfzeng"
	fmt.Println(cnowner)
	var cnarray = make([]string, 0)
	cnarray = append(cnarray, "Put object/Put object range")
	cnarray = append(cnarray, "Put Bucket Lifecycle")
	cnarray = append(cnarray, "Get Bucket Lifecycle")
	cnarray = append(cnarray, "Delete Bucket Lifecycle")
	cnarray = append(cnarray, "Schedule")
	cnarray = append(cnarray, "Migration APIs(if any)")
	cnarray = append(cnarray, "Requirements or pre-requisites for gelato installer")
	cnarray = append(cnarray, "Register Backend")
	cnarray = append(cnarray, "List Backends")
	cnarray = append(cnarray, "Get Bucket")
	cnarray = append(cnarray, "List Buckets")
	cnarray = append(cnarray, "Get bucket location")

	var inlabels = make([]string,1)
	inlabel0 := "Daito"
	inlabels[0] = inlabel0
	//var inowner = "nguptaopensds"
	//var inowner = "rhsakarpos"
	var inowner = "anvithks"
	fmt.Println(inowner)
	var inarray = make([]string, 0)
	/*inarray = append(inarray, "Enable bucket versioning")
	inarray = append(inarray, "Disable bucket versioning")
	inarray = append(inarray, "Put bucket CORS")
	inarray = append(inarray, "Get bucket CORS")
	inarray = append(inarray, "Delete bucket CORS")*/
	/*inarray = append(inarray, "PUT bucket SSE")
	inarray = append(inarray, "GET bucket SSE")
	inarray = append(inarray, "DELETE bucket SSE")
	inarray = append(inarray, "PUT bucket SSE-CK")
	inarray = append(inarray, "GET bucket SSE-CK")
	inarray = append(inarray, "DELETE bucket SSE-CK")
	inarray = append(inarray, "PUT bucket SSE-KMS")
	inarray = append(inarray, "GET bucket SSE-KMS")
	inarray = append(inarray, "DELETE bucket SSE-KMS")*/
	//inarray = append(inarray, "Support Storage Server Encryption on bucket from dashboard")
	inarray = append(inarray, "Support CORS configuration on bucket from dashboard")
	inarray = append(inarray, "Support Object version-ing from dashboard")
	inarray = append(inarray, "Support Access Control List configuration on bucket from dashboard")
	inarray = append(inarray, "Support Policy configuration on bucket from dashboard")
	inarray = append(inarray, "Support Copy Object from dashboard")

	for _, fv := range inarray{
		//for _, bv := range bearray{
			var title = fv// + "-" + bv

			var issue = github.IssueRequest{
				Title:     &title,
				Body:      nil,
				Labels:    &inlabels,
				Assignee:  &inowner,
				State:     nil,
				Milestone: &milestone,
				Assignees: nil,
			}
			retIssue, resp, err := client.Issues.Create(ctx,
				"opensds",
				"opensds-dashboard",
				&issue)
			fmt.Println(retIssue)
			fmt.Println(resp)
			fmt.Println(err)
		//}
	}

/*
YIG

   Put object/Put object range
   Put bucket
   Get object
   Delete object
   Delete bucket
   Head bucket
   Head object
   Post object (Browser-Based Uploads Using POST)
   Put object part
   Put bucket policy
   Get bucket policy
   Delete bucket policy
   Put object ACL
   Get object ACL
   Put bucket ACL
   Get bucket ACL
   Copy object
   New multipart upload
   Abort multipart upload
   Complete multipart upload
   List bucket multipart
   Enable Redis

sfzeng

   Put Bucket Lifecycle
   Get Bucket Lifecycle
   Delete Bucket Lifecycle
   Schedule
   Migration APIs(if any)
   Requirements or pre-requisites for gelato installer
   Register Backend
   List Backends
   Get Bucket
   List Buckets
   Get bucket location

NP

   Enable bucket versioning
   Disable bucket versioning
   Put bucket CORS
   Get bucket CORS
   Delete bucket CORS
   PUT bucket SSE
   GET bucket SSE
   DELETE bucket SSE
   PUT bucket SSE-CK
   GET bucket SSE-CK
   DELETE bucket SSE-CK
   PUT bucket SSE-KMS
   GET bucket SSE-KMS
   DELETE bucket SSE-KMS
 */
}