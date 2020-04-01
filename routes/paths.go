package routes

import "goreactapp/database/functions"

var (
	//CreateData slice to append current data to model
	CreateData = []route{
		{path: "/users", function: functions.CreateUser},
		{path: "/orgs", function: functions.CreateOrg},
		{path: "/clients", function: functions.CreateClient},
		{path: "/groups", function: functions.CreateGroup},
		{path: "/devs", function: functions.CreateDeveloper},
		{path: "/projects", function: functions.CreateProject},
		{path: "/managers", function: functions.CreateManager},
	}

	//ReadAllData slice to get all data from model
	ReadAllData = []route{
		{path: "/roles", function: functions.Roles},
		{path: "/users", function: functions.Users},
		{path: "/orgs", function: functions.Organizations},
		{path: "/clients", function: functions.Clients},
		{path: "/groups", function: functions.WorkGroups},
		{path: "/devs", function: functions.Developers},
		{path: "/status", function: functions.ProjectStatuses},
		{path: "/projects", function: functions.Projects},
		{path: "/projectsview", function: functions.ProjectsPreview},
		{path: "/managers", function: functions.Managers},
	}

	//ReadOne slice to get current data from model
	ReadOne = []route{
		{path: "/users/{id}", function: functions.User},
		{path: "/orgs/{id}", function: functions.Organization},
		{path: "/clients/{id}", function: functions.Client},
		{path: "/groups/{id}", function: functions.Group},
		{path: "/devs/{id}", function: functions.Developer},
		{path: "/projects/{id}", function: functions.Project},
		{path: "/managers/{id}", function: functions.Manager},
	}

	//UpdateOne slice to update current data from model
	UpdateOne = []route{
		{path: "/users/{id}", function: functions.UpdateUser},
		{path: "/orgs/{id}", function: functions.UpdateOrg},
		{path: "/clients/{id}", function: functions.UpdateClients},
		{path: "/groups/{id}", function: functions.UpdateGroups},
		{path: "/devs/{id}", function: functions.UpdateDevelopers},
		{path: "/projects/{id}", function: functions.UpdateProjects},
		{path: "/managers/{id}", function: functions.UpdateClients},
	}

	//DeleteOne slice to delete current data from model
	DeleteOne = []route{
		{path: "/users/{id}", function: functions.DeleteUser},
		{path: "/orgs/{id}", function: functions.DeleteOrg},
		{path: "/clients/{id}", function: functions.DeleteClient},
		{path: "/groups/{id}", function: functions.DeleteGroup},
		{path: "/devs/{id}", function: functions.DeleteDeveloper},
		{path: "/projects/{id}", function: functions.DeleteProject},
		{path: "/managers/{id}", function: functions.DeleteManager},
	}
)
