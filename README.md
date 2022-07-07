# gitlab-artifacts-cleaner

Golang CLI to delete GitLab stored artifacts

## TL;DR

When you need to wipe out a Gitlab project's artifacts, this is the tool for you. It will delete all the artifacts in a project. It will not delete the project itself. It will not delete the project's repository.

Because of GitLab API limitations, you need to know how many job pages you need to iterate through to delete all the artifacts. This is how GitLab API works. If you don't know it, just use the standard values:
  
  ```yaml
  per_page: 100
  page: 20
  ```

## Usage

Parameters:

- `--server`: GitLab server URL. Default: `https://gitlab.com`
- `--token`: Your GitLab personal access token
- `--project_id`: Project ID or path.
- `--per-page`: Number of jobs to fetch per page. Default: `100`
- `--page`: Page number. Default: `1`

```bash
gitlab-artifacts-cleaner --server https://gitlab.com --token <token> --project_id <project_id> --pages 10 --per_page 100
```

This is a work in progress. If you have any suggestions, please open an issue.
