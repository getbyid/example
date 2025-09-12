# example

GoReleaser example using GitHub.

[How I build and ship Go binaries with GoReleaser and GitHub Actions](https://gitgist.com/posts/goreleaser-and-github-actions/)

## Usage

To create a new release:

+ Tag your commit: `git tag v1.0.0`
+ Push the tag: `git push origin v1.0.0`

GitHub Actions will automatically:

+ Build binaries for all platforms
+ Generate checksums
+ Create a GitHub release
+ Update your Homebrew formula

You can also trigger releases manually through the GitHub Actions interface using the `workflow_dispatch` event.

## Best Practices

+ Always test builds locally first: `goreleaser release --snapshot --clean`
+ Use semantic versioning for tags
+ Review the generated release notes before final publication
