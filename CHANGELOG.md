# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Migrated the following endpoints to OpenAPI 3.0:
  - `/2/account`, `/2/account/{id}`
  - `/2/locations`, `/2/locations/{id}`
  - `/2/users/profile`
  - `/2/users/invite`, `/2/users/invite/{id}`
  - `/2/users/bulkupdate`
  - `/2/users/avatar/{id}`
  - `/2/users/{id}`
  - `/2/users`
  - `/2/shifts/{id}/history`
  - `/2/shifts/{id}/assign`
  - `/2/shifts/{id}/swapusers`
  - `/2/shifts/unpublish`
  - `/2/shifts/unassign`
  - `/2/shifts/publish`
  - `/2/shifts/notify/{id}`
  - `/2/shifts/notify`
  - `/2/shifts/eligible`
  - `/2/shifts/bulk`
  - `/2/sites/{id}`
  - `/2/sites`
  - `/2/times/{id}`
  - `/2/times`
  - `/2/shifts/{id}`
  - `/2/shifts`
- Marked migrated endpoints as supported in the documentation.
- Updated and improved README and migration plan documentation.

### Changed

- General code and documentation updates to support OpenAPI 3.0 migration.

## [0.1.0] - 2024-06-10

### Added

- Initial open source release of the Quiqup WheniWork API Client (Go).
- MIT License.
