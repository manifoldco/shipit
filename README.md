# Shipit
Create release issues in GitHub

## Should be:
a bot, yup, one day I will be


### build

`make bootstap`
`make build`

### run

Set up GITHUB_TOKEN env var in your local, or use [manifold](https://www.manifold.co) and do it
the right way

`GITHUB_TOKEN=<cool_token_here> ./bin/shipit`
or
`manifold run ./bin/shipit`

### Args
`-o` owner of the repo (i.e manifoldco)
`-r` repository name (i.e. engineering)
`-v` version tag to use when making the ticket (i.e. v2.0.2)
`-f` filename of the template (optional: tmpl/issue.txt)

## Example

`GITHUB_TOKEN=<cool_token_here> ./bin/shipit -o manifoldco -r engineering -v 2.0.2 -f tmpl/issue.txt`
