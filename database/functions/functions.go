package functions

import (
	"github.com/elgris/sqrl"
)

var postgres = sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)

var (
	//#--ROLES--#//
	SelectRoles = postgres.Select("*").From("roles") //+
	//#--USERS--#//
	SelectUsers = postgres.Select( //+
		"user_login",
		"first_name",
		"middle_name",
		"sur_name",
		"birth_date",
		"phone_num",
		"email_addr",
		"user_image_src",
		"role_id",
		"role_name").
		From("users").
		Join("roles ON user_role = role_id").
		Where("logical_delete_status = false")
	//#--ORGANISATIONS--#//
	SelectOrgs = postgres.Select("*").From("organisations") //+
	//#--CLIENTS--#//
	SelectClients = postgres.Select(
		"user_login",
		"sur_name",
		"first_name",
		"middle_name",
		"birth_date",
		"phone_num",
		"email_addr",
		"user_image_src",
		"role_id",
		"role_name",
		"organisation_id",
		"full_name",
		"short_name",
		"organisation_desc",
		"organisation_image_src").
		From("clients").
		Join("users ON user_login = client_login").
		Join("roles on user_role = role_id").
		Join("organisations on client_organisation_id = organisation_id")
	//#--GROUPS--#//
	SelectWorkGroups = postgres.Select("*").From("workgroups") //+
	//#--DEVELOPERS--#//
	SelectDevs = postgres.Select(
		"user_login",
		"sur_name",
		"first_name",
		"middle_name",
		"birth_date",
		"phone_num",
		"email_addr",
		"user_image_src",
		"outsource_spec",
		"role_id",
		"role_name").
		From("developers").
		Join("users ON user_login = developer_login").
		Join("roles ON user_role = role_id")
	/*SelectDevGroups, _, _ = postgres.Select(
		"list_id",
		"developer_login"
	)*/
	//#--STATUS--#//
	SelectStatus = postgres.Select("*").From("status") //+
	//#--PROJECTS--#//
	SelectProjects = postgres.Select(
		"project_id", "project_name", "project_info",
		"workgroup_id", "workgroup_name",
		"status.status_id", "status.status_name",
		"mn.user_login", "mn.sur_name", "mn.first_name", "mn.middle_name", "mn.birth_date", "mn.phone_num", "mn.email_addr", "mn.user_image_src",
		"ms.user_login as user_login1", "ms.sur_name as sur_name1", "ms.first_name as first_name1", "ms.middle_name as middle_name1", "ms.birth_date as birth_date1", "ms.phone_num as phone_num1", "ms.email_addr as email_addr1", "ms.user_image_src as user_image_src1",
		"organisations.organisation_id", "organisations.short_name", "organisations.organisation_image_src").
		From("projects").
		Join("workgroups ON workgroups.workgroup_id = projects.project_workgroup_id").
		Join("clients ON clients.client_login = projects.project_client_login").
		Join("managers ON managers.manager_login = projects.project_manager_login").
		Join("users mn on projects.project_manager_login = mn.user_login").
		Join("users ms on projects.project_client_login = ms.user_login").
		Join("organisations on clients.client_organisation_id = organisations.organisation_id").
		Join("status ON status.status_id = projects.project_status_id")
	SelectProjectsPreview = postgres.Select(
		"project_id",
		"project_name",
		"organisation_id",
		"project_status_id",
		"status_name",
		"organisation_image_src",
		`(
            SELECT  COUNT(*)
              FROM  (
                  SELECT *FROM tasks
                   WHERE  task_stage_id IN (
                       SELECT  stage_id 
                         FROM  stages 
                        WHERE  stage_module_id IN (
                            SELECT  module_id 
                              FROM  modules
                             WHERE  module_project_id = project_id
                        )
                       )
              ) AS tasks_2
		) AS tasks_all`,
		`(
            SELECT  COUNT(*)
              FROM  (
                  SELECT *FROM tasks
                   WHERE  task_stage_id IN (
                       SELECT  stage_id 
                         FROM  stages 
                        WHERE  stage_module_id IN (
                            SELECT  module_id 
                              FROM  modules
                             WHERE  module_project_id = project_id
                        )
                       ) AND task_status_id = 4
              ) AS tasks_2
		) AS tasks_finished`,
		`(
            SELECT  COUNT(*)
              FROM  (
                  SELECT *FROM issues_list
                   WHERE  task_id IN (
                       SELECT  task_id 
                         FROM  tasks
                        WHERE  task_stage_id IN (
                            SELECT  stage_id 
                              FROM  stages 
                             WHERE  stage_module_id IN (
                                SELECT  module_id 
                                  FROM  modules
                                 WHERE  module_project_id = project_id
                        )
                       )
                   )
              ) AS issues_2
		) AS project_issues`,
		"start_date",
		"finish_date",
		"start_date_fact",
		"finish_date_fact").
		From("projects").
		Join("clients ON project_client_login = client_login").
		Join("organisations ON organisation_id = client_organisation_id").
		Join("status ON status_id = project_status_id")
	//#--MANAGERS--#//
	SelectManagers = postgres.
			Select("user_login", "user_surname", "user_name", "user_midname", "user_birthdate", "user_phone", "user_email", "role_id", "role_name").
			From("managers").
			Join("users ON users.user_login = managers.manager_user_login").
			Join("roles ON users.user_role = role_id")
)
