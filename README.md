# Quiqup WheniWork API Client

The Quiqup WhenIWork API client is a golang client implemented using resty, it supports only a subset of the overall spec provided by wheniwork as their spec is broken, it uses an approach where openapi spec is still at the core primarily to be able to adopt the full spec if the main open api spec can be cleaned and fixed.

## In Development

Package is in active development

## Supported APIs

| Endpoint                              | Description                                 | Supported |
| ------------------------------------- | ------------------------------------------- | :-------: |
| /2/shifts                             | List Shifts                                 |    [x]    |
| /2/shifts/{id}                        | Get, Update, or Delete a single shift by ID |    [x]    |
| /2/times                              | List or Create Times                        |    [x]    |
| /2/times/{id}                         | Get, Update, or Delete a single time by ID  |    [x]    |
| /2/sites                              | List or Create Sites                        |    [x]    |
| /2/sites/{id}                         | Get, Update, or Delete Site                 |    [x]    |
| /2/shifts/bulk                        | Bulk Update Shifts                          |    [x]    |
| /2/shifts/eligible                    | List eligible users for OpenShift           |    [x]    |
| /2/shifts/notify                      | Notify shifts                               |    [x]    |
| /2/shifts/notify/{id}                 | Notify single shift                         |    [x]    |
| /2/shifts/publish                     | Publish Shifts                              |    [x]    |
| /2/shifts/unassign                    | Unassign/Release Shifts                     |    [x]    |
| /2/shifts/unpublish                   | Unpublish Shifts                            |    [ ]    |
| /2/shifts/{id}/assign                 | Assign multiple users to an OpenShift       |    [x]    |
| /2/shifts/{id}/history                | Fetch shift history                         |    [ ]    |
| /2/shifts/{id}/swapusers              | Swap users on a shift                       |    [ ]    |
| /2/users                              | List or Create Users                        |    [ ]    |
| /2/users/{id}                         | Get User                                    |    [ ]    |
| /2/users/avatar/{id}                  | Get User Avatar                             |    [ ]    |
| /2/users/bulkupdate                   | Bulk Update Users                           |    [ ]    |
| /2/users/invite                       | Invite Users                                |    [ ]    |
| /2/users/invite/{id}                  | Invite single user                          |    [ ]    |
| /2/users/profile                      | Get User Profile                            |    [ ]    |
| /2/locations                          | List or Create Locations (Schedules)        |    [ ]    |
| /2/locations/{id}                     | Get or Update Schedule (Location)           |    [ ]    |
| /2/account                            | Get Account                                 |    [ ]    |
| /2/account/{id}                       | Get Account by ID                           |    [ ]    |
| /2/account/{id}/admins                | List Admins                                 |    [ ]    |
| /2/account/settings                   | Update Account Settings                     |    [ ]    |
| /2/annotations                        | List or Create Annotations                  |    [ ]    |
| /2/annotations/{id}                   | Get or Update Annotation                    |    [ ]    |
| /2/availabilityevents                 | List or Create Availability Events          |    [ ]    |
| /2/availabilityevents/{id}            | Get or Update Availability Event            |    [ ]    |
| /2/availabilityevents/{id}/exceptions | Get or Update Availability Event Exceptions |    [ ]    |
| /2/blocks                             | List or Create Shift Templates (Blocks)     |    [ ]    |
| /2/blocks/{id}                        | Get or Update a Shift Template (Block)      |    [ ]    |
| /2/import                             | Create Import                               |    [ ]    |
| /2/import/{id}                        | Get Import                                  |    [ ]    |
| /2/import/{id}/finalize               | Finalize Import                             |    [ ]    |
| /2/import/{id}/preview                | Preview Import                              |    [ ]    |
| /2/payrolls                           | List Payrolls                               |    [ ]    |
| /2/payrolls/{id}                      | Get Payroll                                 |    [ ]    |
| /2/payrolls/{id}/stats                | Get Payroll Statistics                      |    [ ]    |
| /2/payrolls/notices                   | Payroll Notices                             |    [ ]    |
| /2/positions                          | List or Create Positions                    |    [ ]    |
| /2/positions/{id}                     | Get, Update, or Delete a Position           |    [ ]    |
| /2/positions/bulk                     | Bulk Update Positions                       |    [ ]    |
| /2/punch/state                        | Get Punch State                             |    [ ]    |
| /2/requests                           | List or Create Time Off Requests            |    [ ]    |
| /2/requests/{request_id}              | Get, Update, or Delete Time Off Request     |    [ ]    |
| /2/requesttypes                       | List Of Time Off Request Types              |    [ ]    |
| /2/requesttypes/{id}                  | Get or Delete Time Off Request Type         |    [ ]    |
| /2/swaps                              | List or Create Shift Swaps                  |    [ ]    |
| /2/swaps/{swap_id}                    | Get or Delete a specific shift request      |    [ ]    |
| /2/times/clockin                      | Clock In                                    |    [ ]    |
| /2/times/clockout                     | Clock Out                                   |    [ ]    |
| /2/templates                          | List or Create Schedule Templates           |    [ ]    |
| /2/templates/{id}                     | Get Schedule Template                       |    [ ]    |
| /2/timezones                          | List Timezones                              |    [ ]    |
| /2/timezones/{id}                     | Get Timezone                                |    [ ]    |
| /2/users/alerts                       | Get User Alerts                             |    [ ]    |
