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
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "user_id",
					In:          "query",
					Description: "The user id to filter by",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeInteger}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "start",
					In:          "query",
					Description: "The start of the filter range.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "end",
					In:          "query",
					Description: "The end of the filter range.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "unpublished",
					In:          "query",
					Description: "Whether or not to include unpublished shifts. Requires supervisor rights.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "include_open",
					In:          "query",
					Description: "Whether or not to include open shifts from the user's assigned Schedules.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "include_onlyopen",
					In:          "query",
					Description: "Whether or not to include only open shifts from the user's assigned Schedules.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "include_allopen",
					In:          "query",
					Description: "Whether or to include open shifts across All Schedules. Requires \"Manager or Admin access\" level.\nCommon practice is to combine allopen with one of the other inclusion options.\n",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "deleted",
					In:          "query",
					Description: "Whether to include a list of shift IDs (\"deleted_ids\") that were deleted during the passed time window.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "include_swaps",
					In:          "query",
					Description: "Whether or not to include swap requests.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "limit",
					In:          "query",
					Description: "Maximum number of results to return.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeInteger}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "all_locations",
					In:          "query",
					Description: "Whether to include data from all locations. Shifts are marked as \"readonly\" if not a manager user.\n\nIf this option is included in addition to the `location_id` option, all shifts linked to other\nlocations, through users in other locations, will also be included.\n",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "location_id",
					In:          "query",
					Description: "One or more location IDs by which to limit results\n\n_Also see `all_locations` above_\n",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "shift_sort",
					In:          "query",
					Description: "True to sort results by user_id, false to sort by shift time. Missing for default sort",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "include_repeating_shifts_to",
					In:          "query",
					Description: "End date to include repeating shifts in series, if applicable",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}}},
				}},
				&openapi3.ParameterRef{Value: &openapi3.Parameter{
					Name:        "trim_openshifts",
					In:          "query",
					Description: "Setting to true will work w/ the Allow Partial Openshifts feature to display trimmed start/end times for users that can take a conflicting openshift based on the account settings.",
					Schema:      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}},
				}},
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
