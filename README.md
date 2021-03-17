# Google Workspace credentials validator

## Usage

```bash
$ docker build -t gw-validator .
$ docker run --rm -v "${PWD}/file.json":"/app/creds.json" gw-validator \
    -jwtfile "/app/creds.json" \
    -delegated_account "account@example.com" \
    -scopes "https://www.googleapis.com/auth/admin.reports.audit.readonly" \
    -endpoint "https://www.googleapis.com/admin/reports/v1/activity/users/all/applications/admin"
```