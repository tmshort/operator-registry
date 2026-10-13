# Default branch migration plan: `master` to `main`

This plan covers `operator-framework/operator-registry` and the references found in `openshift/operator-framework-tooling` and `openshift/release`. Keep `master` available during the transition.

## 1. Prepare this repository

Update the following files in `operator-framework/operator-registry`:

- `.github/workflows/go-verdiff.yaml`: add `main` to the pull request branch trigger. Keep `master` in the trigger during the transition.
- `.github/workflows/goreleaser.yaml`: add `main` to the push branch trigger. Keep `master` temporarily so pushes there continue to run the workflow.
- `.github/PULL_REQUEST_TEMPLATE.md`: tell contributors to rebase on `main`.

Other `master` references found in this repository point to other projects' branches or use “master” as a domain term; they are unrelated to this rename.

## 2. Update OpenShift release configuration

Update `openshift/release/core-services/prow/02_config/operator-framework/operator-registry/_prowconfig.yaml`:

- Add `main` to the Tide query that currently includes `master`, preserving the same merge labels and requirements for pull requests targeting `main`.
- Add `main` to the excluded branch list in the broad query, alongside `master`, to preserve the current query separation.
- Retain `master` in both places during the transition. Remove it from the relevant lists only when retiring the branch.

The repository-specific ci-operator config and job directories contain only `OWNERS` files; no registry-specific job definitions were found there. Other `operator-registry` matches in the release repository are image names or refer to unrelated repositories.

## 3. Check `openshift/operator-framework-tooling`

No change is currently indicated specifically for this repository's rename:

- The `upstreamBranch = "master"` constant and related fetches in `pkg/v0/cmd.go` refer to `operator-framework/operator-lifecycle-manager`, not `operator-registry`.
- For dependency repositories including `operator-registry`, the tool resolves versions from OLM's `go.mod` and fetches the resulting tag or commit.
- Existing downstream branch selection already prefers `main` when available and falls back to `master`; it concerns `openshift/operator-framework-olm`.

Recheck this conclusion if the tooling's dependency sync behavior changes.

## 4. Rename and verify

1. Create `main` at the current `master` tip and set `main` as the GitHub default branch.
2. Keep `master` available and synchronized during the agreed transition period.
3. Confirm pull request checks and Tide merge behavior work for `main`, and pushes to `main` trigger GoReleaser.
4. Update contributor guidance and branch-specific links as they are encountered.
5. When the transition is complete, remove `master` from workflow triggers and Prow rules, then retire the branch.

## Audit scope

The review covered branch references in this repository's `.github` configuration and repository-specific references in the local checkouts of `openshift/operator-framework-tooling` and `openshift/release`. No files were changed as part of the audit.
