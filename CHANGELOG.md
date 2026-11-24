# Changelog

All notable changes to Constellation are documented here. The current version is v1.1.0.

## [1.1.0] - 2026-11-24

### Added

- Hot-reload development server in `lib/dev`. Run `go run ./lib/dev` to rebuild and refresh open browser tabs when `.go`, `.css`, or `.md` files change.
- Example API and web endpoints under `src/api/` and `src/web/`.
- Element helper example in the getting-started guide.

### Changed

- Moved the framework packages from `src/` to `lib/` and updated imports to `constellation/lib`.
- Updated the README, user guides, and contributor notes for the new layout.

### Fixed

- Permission checks run before endpoint handlers and return `403 Forbidden` when access is denied.

## [1.0.0] - 2026-11-24

### Added

- Initial standard-library framework with declarative endpoints, HTML helpers, an embedded stylesheet, Tailwind CSS support, and Markdown documentation served at `/docs`.
