# API Migration Task Tracker

This document tracks the migration of API routes from `spec/original-spec.json` to `spec/apispec.yml`. During this migration, we do not
need to look at any other files in the codebase just [original spec](./spec/original-spec.json) and [api spec](./spec/apispec.yml)

## Instructions

- Each route to be migrated is listed as a task.
- Check off each task after:
  1. Migrating the route to `apispec.yml`. (ensure it is openapi 3.0.0 compliant)
  2. Make sure the new route has a default response
  3. Ensuring the project compiles and tests pass.
  4. Verifying the route works as expected.
- Add notes or issues as needed.

---

## Migration Tasks

- [x] **/2/shifts**: List Shifts (already migrated)
- [x] **/2/shifts/{id}**: Get, Update, or Delete a single shift by ID
- [ ] **/2/times**: List or Create Times
- [ ] **/2/times/{id}**: Get, Update, or Delete a single time by ID
- [ ] **/2/times/clockin**: Clock In
- [ ] **/2/times/clockout**: Clock Out
- [ ] **/2/shifts/publish**: Publish Shifts
- [ ] **/2/shifts/unpublish**: Unpublish Shifts
- [ ] **/2/blocks**: List or Create Shift Templates (Blocks)
- [ ] **/2/blocks/{id}**: Get or Update a Shift Template (Block)
- [ ] **/2/templates/{id}**: Get Schedule Template
- [ ] **/2/annotations/{id}**: Get or Update Annotation
- [ ] **/2/requests/{request_id}**: Get, Update, or Delete Time Off Request
- [ ] **/2/requesttypes**: List Of Time Off Request Types
- [ ] **/2/requesttypes/{id}**: Get or Delete Time Off Request Type
- [ ] **/2/swaps/{swap_id}**: Get or Delete a specific shift request
- [ ] **/2/positions/{id}**: Get, Update, or Delete a Position
- [ ] **/2/locations/{id}**: Get or Update Schedule (Location)
- [ ] **/2/sites/{id}**: Get, Update, or Delete Site
- [ ] **/2/payrolls**: List Payrolls
- [ ] **/2/payrolls/{id}**: Get Payroll
- [ ] **/2/users/{id}**: Get User
- [ ] **/2/users/profile**: Get User Profile
- [ ] **/2/users/invite/{id}**: Invite single user
- [ ] **/2/account**: Get Account
- [ ] **/2/account/settings**: Update Account Settings
- [ ] **/2/account/{id}/admins**: List Admins
- [ ] **/2/import**: Create Import
- [ ] **/2/import/{id}**: Get Import
- [ ] **/2/import/{id}/preview**: Preview Import
- [ ] **/2/import/{id}/finalize**: Finalize Import
- [ ] **/2/timezones**: List Timezones
- [ ] **/2/timezones/{id}**: Get Timezone
- [ ] **/v3/scheduled-breaks**: List Scheduled Shift Breaks
- [ ] **/v3/shift-breaks**: List Shift Breaks
- [ ] **/v3/shift-breaks/{id}**: Get Shift Break
- [ ] **/v3/shift-break-audits**: List Shift Break - Paid Rest records (Deprecated)
- [ ] **/v3/shift-break-audits/{id}**: Get or Delete Shift Break - Paid Rest Record (Deprecated)
- [ ] **/v3/auto-assign**: List Auto Scheduled shifts

<!-- Add notes or mark tasks as complete as you progress -->
