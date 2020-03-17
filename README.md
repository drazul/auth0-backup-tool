# Auth0 Backup Tool

This command line tool allow us to export all users from a given Auth0 database.
This tool also allow us to import the users backup to a given Auth0 database.

# Usage

```
Usage of auth0-backup-tool:
  -action string
        Action to perform. Can be 'import' or 'export'
  -client-id string
        Client ID of an application with user management rights
  -client-secret string
        Client secret of an application with user management rights
  -connection string
        Auth0 connection ID
  -domain string
        Auth0 domain
  -user-attributes string
        List of user attributes to export. Format: attr1,attr2,attr3
  -users-file string
        File path where to store the exported users or where to read the users to import (default "users-export.json")
```

# User export format with default attributes
```json
{
  "user_id": "auth0|5ce405424afaal03b7590dba",
  "nickname": "exampleuser",
  "name": "exampleuser@example.com",
  "email": "exampleuser@example.com",
  "email_verified": true,
  "created_at": "2019-05-21T14:03:46.999Z",
  "updated_at": "2019-07-19T08:10:00.408Z",
  "user_metadata": {
    "exampleKey": "exampleValue"
  },
  "app_metadata": {
    "exampleKey": "exampleValue"
  },
  "logins_count": 111,
  "last_login": "2019-07-19T08:10:00.408Z",
  "identities": [
    {
      "profileData": {
        "email": "exampleuser@example.com",
        "email_verified": true
      },
      "user_id": "5ce405424afaal03b7590dba",
      "provider": "auth0",
      "connection": "MyAuth0Database",
      "isSocial": false
    }
  ]
}
```
