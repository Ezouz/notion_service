#This is the base of a Notion API connector for 42 Paris Bocal teamspace

Documentation partie js:
https://developers.notion.com/
https://github.com/makenotion/notion-sdk-js

Documentation partie go:
https://pkg.go.dev/github.com/dstotijn/go-notion
https://pkg.go.dev/github.com/dstotijn/go-notion#section-readme
https://www.datanovia.com/en/lessons/docker-compose-wait-for-container-using-wait-tool/docker-compose-wait-for-postgres-container-to-be-ready/

#Partie script

Documentation :
https://pypi.org/project/notionhq-client/


## Regenerate API
If any modification to `swagger.yml` has been made you need to regenerate the API:
```
make gen
```


## Regenerate Database model
If any modification to `/pkg/restapi/configure_hub42_api.go` has been made you need to regenerate the database model:

(This require golang installed)
```
go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent ./ent/schema/
```


## Swagger Documentation
Documentation is available at `/docs`
