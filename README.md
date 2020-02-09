# JSON Search
1. A simple command line application to search the data and return the results in a human readable format from JSON files
2. all searchable feilds will be retrived from Json file data automatically.
3. Output fields and order of the fields can be customized.
4. Type of the Search value can be **bool** **int** **string**
5. Using [go-prompt](https://github.com/c-bata/go-prompt) library for building powerful interactive prompts
6. Using [gojsonq](https://github.com/thedevsaddam/gojsonq) to do the json query
## Demo
![Demo](demo/jsonsearch.gif)
## Download the appliction
Download the binary file from [https://github.com/rocky0001/jsonsearch/releases](https://github.com/rocky0001/jsonsearch/releases)
## config.yaml
Put the config.yaml in the same folder as the application file,modifty the setttings per your requirement(__json files location__,__output fields__,__related records output fields__):

```
usersJsonFile:  "./users.json"
ticketsJsonFile: "./tickets.json"
organizationsJsonFile: "./organizations.json"
outputfields:
  users:
     - "_id"
     - "url"
     - "external_id"
     - "name"
     - "alias"
     - "created_at"
     - "active"
     - "verified"
     - "shared"
     - "locale"
     - "timezone"
     - "last_login_at"
     - "email"
     - "phone"
     - "signature"
     - "organization_id"
     - "tags"
     - "suspended"
     - "role"
  tickets:
     - "_id"
     - "url"
     - "external_id"
     - "created_at"
     - "type"
     - "subject"
     - "description"
     - "priority"
     - "status"
     - "submitter_id"
     - "assignee_id"
     - "organization_id"
     - "tags"
     - "has_incidents"
     - "due_at"
     - "via"
  organizations:
     - "_id"
     - "url"
     - "external_id"
     - "name"
     - "domain_names"
     - "created_at"
     - "details"
     - "shared_tickets"
     - "tags"
relatedoutputfields:
  users:
     - "name"
     - "alias"
     - "active"
     - "verified"
     - "shared"
     - "timezone"
     - "suspended"
     - "role"
  tickets:
     - "created_at"
     - "type"
     - "subject"
     - "description"
     - "priority"
     - "status"
     - "has_incidents"
     - "due_at"
     - "via"
  organizations:
     - "name"
     - "domain_names"
     - "details"
     - "shared_tickets"
```


## Run the application
Update the file permission if required and then run the application

### Some commands:

#### Exit the application:
For example:
```
quit
```

#### Select json file to do further query:
For example:
```
select users
```

#### List searchable fields:

For example:
List all the searchable fields for **all** json files:
```
list
```

List all the searchable fields for **users** json files:
```
list users
```

#### Start a new query from current json:
For example:
Search all the records from **users** where _id startsWith F:
```
users>>where _id startsWith F
```

Search all the records from **users** where tags has "Springville":
```
users>>where tags has Springville
```

Search all the records from **tickets** where has_incidents is **true**:
```
tickets>>where has_incidents = true
```

#### Filter records from last search results:
For example:
Filter  the records from **last search results** where _id startsWith F:
```
users>>where _id startsWith F
```

#### Returen to root menu:
For example:

```
users>>return
```

