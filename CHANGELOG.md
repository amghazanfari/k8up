# Backup as a service

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.0.7] - 2018-09-13
### Added
- Global repo password. The operator can use a global password for Restic repositories for backups which don't specify an own password.

## [v0.0.6] - 2018-09-12
### Added
- Manage service accounts, roles and rolebindings in namespaces for backup commands. Please see the README chapter `Installation changes` about the new permissions necessary for the operator service account.
- Global cluster state. The operator now knows when a prune job is running to a shared repository and waits before starting the next backup/prune.
- Global S3 credentials. The operator can use global S3 credentials to be used with backups which don't specify own credentials.
### Fixed
- Ignore if the service account, roles or rilebindings already exist in the namespace.

## [v0.0.5] - 2018-08-15
### Added
- Consistent backups -> it is now possible to set backup commands on pods. This changes the prerequisites, please consult the README chapter `Installation changes` for more infos.
### Fixed
- Crash when a running job got deleted
- Skip backup if there are no backup commands and no PVCs

## [v0.0.4] - 2018-08-06
### Changes
- Skip backup if no suitable PVCs are found
- Minor refactoring
- When the backup definition is deleted it will also cleanup all jobs and pods

## [v0.0.3] - 2018-07-27
### Changed
- CRD status, start and end are now timestamps
- Print namespace wenn a backup is created

## [v0.0.2] - 2018-07-27
### Added
- Vscode launch config
- Read credentials and password from OpenShift secrets
- Specifiy default Prometheus push gateway for the operator
- Read the Prometheus push gateway url from the CRD, has precedence over the operator settings
### Changed
- Use alpine for operator image
- Moved the Restic wrapper to its own project: https://git.vshn.net/vshn/wrestic

## [v0.0.1]
### Changed
- Initial version

[unreleased]: https://git.vshn.net/vshn/baas/compare/v0.0.5...master
[v0.0.5]: https://git.vshn.net/vshn/baas/compare/v0.0.4...v0.0.5
[v0.0.4]: https://git.vshn.net/vshn/baas/compare/v0.0.3...v0.0.4
[v0.0.3]: https://git.vshn.net/vshn/baas/compare/v0.0.2...v0.0.3
[v0.0.2]: https://git.vshn.net/vshn/baas/compare/v0.0.1...v0.0.2
[v0.0.1]: https://git.vshn.net/vshn/baas/tree/v0.0.1