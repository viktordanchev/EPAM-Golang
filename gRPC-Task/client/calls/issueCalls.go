package calls

import (
	"context"
	"fmt"
	"log"

	pbIssue "server/gen/pb/issue"

	"google.golang.org/protobuf/types/known/emptypb"
)

func MakeIssueCalls(ctx context.Context, client pbIssue.IssueServiceClient) {
	issue1 := createIssue(ctx, client,
		"Login bug",
		"Cannot login with correct credentials",
		pbIssue.IssueType_BUG,
		pbIssue.IssuePriority_CRITICAL,
	)

	issue2 := createIssue(ctx, client,
		"Add dark mode",
		"Feature request for dark mode",
		pbIssue.IssueType_FEATURE,
		pbIssue.IssuePriority_MAJOR,
	)

	getIssue(ctx, client, issue1.IssueId)
	getIssue(ctx, client, issue2.IssueId)

	listIssues(ctx, client)

	updateIssue(ctx, client, issue1)

	deleteIssue(ctx, client, issue1.IssueId)

	listIssues(ctx, client)
}

func createIssue(
	ctx context.Context,
	client pbIssue.IssueServiceClient,
	summary, description string,
	issueType pbIssue.IssueType,
	priority pbIssue.IssuePriority,
) *pbIssue.Issue {

	issue, err := client.CreateIssue(ctx, &pbIssue.Issue{
		Summary:     summary,
		Description: description,
		Type:        issueType,
		Priority:    priority,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CREATED ISSUE:", issue)
	return issue
}

func updateIssue(ctx context.Context, client pbIssue.IssueServiceClient, issue *pbIssue.Issue) {
	issue.Description = "UPDATED: " + issue.Description
	issue.Status = pbIssue.IssueStatus_IN_PROGRESS

	updated, err := client.UpdateIssue(ctx, issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UPDATED ISSUE:", updated)
}

func getIssue(ctx context.Context, client pbIssue.IssueServiceClient, issueID string) {
	issue, err := client.GetIssue(ctx, &pbIssue.GetIssueRequest{
		IssueId: issueID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET ISSUE:", issue)
}

func deleteIssue(ctx context.Context, client pbIssue.IssueServiceClient, issueID string) {
	_, err := client.DeleteIssue(ctx, &pbIssue.DeleteIssueRequest{
		IssueId: issueID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETED ISSUE:", issueID)
}

func listIssues(ctx context.Context, client pbIssue.IssueServiceClient) {
	issues, err := client.ListIssues(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LIST ISSUES:", issues)
}
