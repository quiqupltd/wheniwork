# API Migration Task Tracker

This document tracks the migration of all API routes from `spec/original-spec.json` to `spec/apispec.yml`.

**Instructions:**

- Each API route is listed below. For each route follow these steps DO NO SKIP ANY STEPS:

  1. Create a new branch using feat/<apislug>
  2. Migrate the route to `apispec.yml` (ensure OpenAPI 3.0.0 compliance).
  3. Migrate all referenced schemas (e.g., request/response bodies) to `apispec.yml`.
  4. Ensure the new route has a default response.
  5. Use multi-line YAML for any multi-line text fields (see example in this file).
  6. Run `task generate` to verify the code generator works (no errors).
  7. Mark the task as complete in this table.
  8. Update the main [readme](README.md) for supported APIs
  9. Commit all files using -am to make sure to add all changes
  10. Use gh command line tool to open a new pr

- Add notes or issues as needed in the Notes column.

---

## Migration Tasks

| Endpoint                                  | Status | Notes |
| ----------------------------------------- | :----: | ----- |
| **/2/sites**                              |  [x]   |       |
| **/2/sites/{id}**                         |  [x]   |       |
| **/2/shifts**                             |  [x]   |       |
| **/2/shifts/{id}**                        |  [x]   |       |
| **/2/shifts/bulk**                        |  [x]   |       |
| **/2/shifts/eligible**                    |  [x]   |       |
| **/2/shifts/notify**                      |  [ ]   |       |
| **/2/shifts/notify/{id}**                 |  [ ]   |       |
| **/2/shifts/publish**                     |  [ ]   |       |
| **/2/shifts/unassign**                    |  [ ]   |       |
| **/2/shifts/unpublish**                   |  [ ]   |       |
| **/2/shifts/{id}/assign**                 |  [ ]   |       |
| **/2/shifts/{id}/history**                |  [ ]   |       |
| **/2/shifts/{id}/swapusers**              |  [ ]   |       |
| **/2/users**                              |  [ ]   |       |
| **/2/users/{id}**                         |  [ ]   |       |
| **/2/users/avatar/{id}**                  |  [ ]   |       |
| **/2/users/bulkupdate**                   |  [ ]   |       |
| **/2/users/invite**                       |  [ ]   |       |
| **/2/users/invite/{id}**                  |  [ ]   |       |
| **/2/users/profile**                      |  [ ]   |       |
| **/2/locations**                          |  [ ]   |       |
| **/2/locations/{id}**                     |  [ ]   |       |
| **/2/account**                            |  [ ]   |       |
| **/2/account/{id}**                       |  [ ]   |       |
| **/2/account/{id}/admins**                |  [ ]   |       |
| **/2/account/settings**                   |  [ ]   |       |
| **/2/annotations**                        |  [ ]   |       |
| **/2/annotations/{id}**                   |  [ ]   |       |
| **/2/availabilityevents**                 |  [ ]   |       |
| **/2/availabilityevents/{id}**            |  [ ]   |       |
| **/2/availabilityevents/{id}/exceptions** |  [ ]   |       |
| **/2/blocks**                             |  [ ]   |       |
| **/2/blocks/{id}**                        |  [ ]   |       |
| **/2/import**                             |  [ ]   |       |
| **/2/import/{id}**                        |  [ ]   |       |
| **/2/import/{id}/finalize**               |  [ ]   |       |
| **/2/import/{id}/preview**                |  [ ]   |       |
| **/2/payrolls**                           |  [ ]   |       |
| **/2/payrolls/{id}**                      |  [ ]   |       |
| **/2/payrolls/{id}/stats**                |  [ ]   |       |
| **/2/payrolls/notices**                   |  [ ]   |       |
| **/2/positions**                          |  [ ]   |       |
| **/2/positions/{id}**                     |  [ ]   |       |
| **/2/positions/bulk**                     |  [ ]   |       |
| **/2/punch/state**                        |  [ ]   |       |
| **/2/requests**                           |  [ ]   |       |
| **/2/requests/{request_id}**              |  [ ]   |       |
| **/2/requesttypes**                       |  [ ]   |       |
| **/2/requesttypes/{id}**                  |  [ ]   |       |
| **/2/swaps**                              |  [ ]   |       |
| **/2/swaps/{swap_id}**                    |  [ ]   |       |
| **/2/times**                              |  [x]   |       |
| **/2/times/{id}**                         |  [x]   |       |
| **/2/times/clockin**                      |  [ ]   |       |
| **/2/times/clockout**                     |  [ ]   |       |
| **/2/templates**                          |  [ ]   |       |
| **/2/templates/{id}**                     |  [ ]   |       |
| **/2/timezones**                          |  [ ]   |       |
| **/2/timezones/{id}**                     |  [ ]   |       |
| **/2/users/alerts**                       |  [ ]   |       |
