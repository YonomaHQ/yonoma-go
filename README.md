## [Yonoma](https://yonoma.io/) Go Email Marketing SDK
### Install
```bash
go get github.com/YonomaHQ/yonoma-go
```
### Setup
First, you need to get an API key:
```go
client := yonoma.NewClient("api-key") 
```
### Usage
## Contacts
#### Create new contact
```go
contactData := map[string]any{
  "email" : "email@example.com",
  "status": "Subscribed" | "Unsubscribed"
  "data": {
    "firstName": string,
    "lastName": string,
    "phone": string,
    "gender": string,
    "address": string,
    "city": string,
    "state": string,
    "country": string,
    "zipcode": string,
  }
}
response, err := client.CreateContact("Contact Id", contactData) 
```
#### Update contact
```go
contactData := map[string]any{
    "status" : "Subscribed" | "Unsubscribed"
}
response, err := client.UnsubscribeContact("List Id", "Contact Id", contactData)
```
#### Add tag to contact
```go
contactData := map[string]any{
    "tag_id" : "Tag Id"
}
response, err := client.AddTag("Contact Id", contactData)
```
#### Remove tag from contact
```go
contactData := map[string]any{
    "tag_id" : "Tag Id"
}
response, err := client.RemoveContactTag("Contact Id", contactData)

```
## Managing Tags
#### Create a Tag
```go
tagData := map[string]any{
    "tag_name" : "Tag Name"
}
response, err := client.CreateTag(tagData)
```
#### Update a Tag
```go
tagData := map[string]any{
    "tag_name" : "Tag Name"
}
response, err := client.UpdateTag("Tag Id", tagData)
```
#### Delete a Tag
```go
response, err := client.DeleteTag("Tag Id")
```
#### Retrieve a Specific Tag
```go
response, err := client.RetrieveTag("Tag Id")
```
#### List All Tags
```go
response, err := client.ListTags()
```
## Managing Lists
#### Create a List
```go
listData := map[string]any{
    "list_name" : "List Name"
}
response, err := client.CreateList(listData)

```
#### Update a List
```go
listData := map[string]any{
    "list_name" : "List Name"
}
response, err := client.UpdateList("List Id", listData)
```
#### Delete a List
```go
response, err := client.DeleteList("List Id")
```
#### Retrieve a Specific List
```go
response, err := client.RetrieveList("List Id")
```
#### List All Lists
```go
response, err := client.ListLists()
```


