package spec

import "github.com/getkin/kin-openapi/openapi3"

func Document() *openapi3.T {
	spec := &openapi3.T{
		OpenAPI: "3.1.0",
		Info: &openapi3.Info{
			Version:     "1.0.0",
			Title:       "When I Work API Documentation",
			Description: "The When I Work API is thorough, flexible, and restful. Its methods are logically grouped and follow standard conventions. Make a selection from the left to jump to the method group you would like to know more about.\n\nWhen designing your integration, When I Work recommends leveraging our Webhooks subscriptions if you plan to regularly pull data to sync records in your data store. This may be preferable to using our API for tasks like staying up to date about shifts or time entries in your account. Frequent large API requests may run into rate limitations.\nFind out more about Webhooks at our [Help Center](https://help.wheniwork.com/articles/webhooks-reference/) or contact our [Customer Care team](mailto:support@wheniwork.com) for assistance.\n\nFor more information about obtaining an API key, or general API questions, please refer to the [Help Center](https://help.wheniwork.com/articles/api-services-reference-guide/).\n",
		},
		Servers: []*openapi3.Server{
			{
				URL:         "https://api.wheniwork.com",
				Description: "Production",
			},
		},
		Security: openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{
				"W-Token": []string{},
			},
		},
		Tags:  tags(),
		Paths: paths(),
	}
	return spec
}

func paths() *openapi3.Paths {
	return openapi3.NewPaths(
		openapi3.WithPath("/2/shifts", shiftPaths()),
		openapi3.WithPath("/2/shifts/{id}", shiftIDPaths()),
	)
}

// For paths matching /2/shifts
func shiftPaths() *openapi3.PathItem {
	return &openapi3.PathItem{
		Get: &openapi3.Operation{
			Summary:     "List Shifts",
			Description: "Fetch a list of shifts based on a set of filters",
			Tags:        []string{"Shifts"},
			Parameters: openapi3.Parameters{
				openapi3.NewQueryParameter("user_id", "The user id to filter by", "integer", 201),
				openapi3.NewQueryParameter("start", "The start of the filter range.", "string", "now"),
				openapi3.NewQueryParameter("end", "The end of the filter range.", "string", "now + 3 days"),
				openapi3.NewQueryParameter("unpublished", "Whether or not to include unpublished shifts. Requires supervisor rights.", "boolean", false),
			},
		},
	}
}

// For paths matching /2/shifts/{id}
func shiftIDPaths() *openapi3.PathItem {
	return &openapi3.PathItem{
		Get: &openapi3.Operation{},
	}
}

func tags() []*openapi3.Tag {
	return []*openapi3.Tag{
		{
			Name:        "Shifts",
			Description: "Shifts provide the basis for scheduling. Many other objects, including Schedules (aka Locations), Positions, Sites, Users, Tasks, and Tags all link through Shifts.\n\nFor more information about how to use Shifts, visit the [Help Center](https://help.wheniwork.com/article-categories/scheduling-your-team/).\n",
		},
		{
			Name:        "Accounts",
			Description: "Accounts (aka Workplaces) are objects that define a business account with When I Work. Each user is associated with an account enabling them to access Shifts or other data.\n\nFor more information about Accounts, visit the [Help Center](https://help.wheniwork.com/article-categories/account-settings/).\n",
		},
		{
			Name:        "Users",
			Description: "A person on the When I Work platform is associated with a two tier record.   The persons email/password is associated to a person_id.\nFor each Workplace the person belongs to a user_id record exists. This record links back to the person_id so only a single instance of the email/password exists. If a user updates their email or password this applies to all related instances of the user.\n\nFor more information about Users, visit the [Help Center](https://help.wheniwork.com/articles/user-access-privileges-computer/).\n",
		},
		{
			Name:        "Sites",
			Description: "Sites communicate additional information about a given shift to the recipient. Typical usage is when you schedule employees for shifts that are at different physical sites (addresses) compared to the Schedule where the shift is assigned. A Job site can include additional detail (address or coordinates) that informs the user where this particular shift is to take place. An additional use case would be if you want to use an additional labor reporting tier and having the Job Site be a project or cost center value.\n\n For more information about using Sites, visit the [Help Center](https://help.wheniwork.com/articles/scheduling-shifts-at-job-sites-computer/).\n",
		},
	}
}
