# Changes made to original script

## Redirects

The most important feature when migrating blogging engines is that historical URLs are redirected
correctly.

The original script creates a `slug`.

## Output header format: use yaml instead of toml

The default `hugo new` command creates yaml headers, so I'll also use yaml.

## Handle alternative category syntax

In octopress one can use different array syntax in the header section.

Option 1:

```yaml
categories:
- foo
- bar
- baz
```

Option 2:

```yaml
categories: [ foo, bar, baz ]
```

The original script handles the first option. This script will also handle the second option.

## Other changes

TODO

## Example data

I am using some real posts from my previous octopress blog. The examples are located in folder
`example-input`.

Example output files are located in folder `example-output`. This folder is excluded from git (via
gitignore).

