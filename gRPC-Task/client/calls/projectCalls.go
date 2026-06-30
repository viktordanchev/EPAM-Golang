package calls

import (
	"context"
	"fmt"
	"log"
	pbProject "server/gen/pb/project"

	"google.golang.org/protobuf/types/known/emptypb"
)

func MakeProjectCalls(ctx context.Context, client pbProject.ProjectServiceClient) {
	project1 := createProject(ctx, client, "Project A", "First project")
	project2 := createProject(ctx, client, "Project B", "Second project")

	getProject(ctx, client, project1.ProjectId)
	getProject(ctx, client, project2.ProjectId)

	listProjects(ctx, client)

	deleteProject(ctx, client, project1.ProjectId)

	listProjects(ctx, client)
	updateProject(ctx, client, project2)
}

func createProject(ctx context.Context, client pbProject.ProjectServiceClient, name, description string) *pbProject.Project {
	project, err := client.CreateProject(ctx, &pbProject.Project{
		Name:        name,
		Description: description,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CREATED PROJECT:", project)
	return project
}

func updateProject(ctx context.Context, client pbProject.ProjectServiceClient, project *pbProject.Project) *pbProject.Project {
	project, err := client.UpdateProject(ctx, &pbProject.Project{
		ProjectId:   project.ProjectId,
		Name:        "UPDATED Second project",
		Description: project.Description,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UPDATED PROJECT:", project)
	return project
}

func getProject(ctx context.Context, client pbProject.ProjectServiceClient, projectID string) {
	project, err := client.GetProject(ctx, &pbProject.GetProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET PROJECT:", project)
}

func deleteProject(ctx context.Context, client pbProject.ProjectServiceClient, projectID string) {
	res, err := client.DeleteProject(ctx, &pbProject.DeleteProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETED PROJECT:", res, projectID)
}

func listProjects(ctx context.Context, client pbProject.ProjectServiceClient) {
	projects, err := client.ListProjects(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LIST PROJECTS:", projects)
}
