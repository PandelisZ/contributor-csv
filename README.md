# Contributor CSV

Use this to get a CSV dump of the [GitHub API Single User](https://developer.github.com/v3/users/#get-a-single-user) data on contributors to a repository.


## Usage

```sh
contributor-csv user repo >> repo.csv
```

The program outputs the generated csv to stout so it's important to pipe this to a file somewhere using `>>` or `>`


