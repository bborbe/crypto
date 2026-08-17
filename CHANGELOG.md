# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## Unreleased

- Bump go toolchain to 1.26.6
- Bump golang.org/x/mod to v0.40.0 (CVE-2026-56864, CVE-2026-56865)
- Bump golang.org/x/net to v0.58.0, golang.org/x/sys to v0.47.0, golang.org/x/text to v0.41.0, golang.org/x/tools to v0.49.0
## v1.0.5

- Update Go to 1.26.6 and update dependencies

## v1.0.4

- fix: Decrypt returns an error instead of panicking when the value is shorter than the GCM nonce
- add regression specs for short and empty Decrypt input
- docs: add License section to README

## v1.0.3

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v1.0.2

- Bump go toolchain to 1.26.5
- Bump github.com/bborbe/errors to v1.5.16

## v1.0.1

- Update ginkgo to v2.32.0 and gomega to v1.42.1
- Bump indirect golang.org/x dependencies
- Clean up import formatting in test files and mocks

## v1.0.0

- Initial Version
