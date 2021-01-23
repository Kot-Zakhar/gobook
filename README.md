# GoBook

## 1. Pure Go

The goal of first task is to create a simple http server to operate on notes.

## Objects and operations

### Note object

A **note** is a an object of the corresponding structure:

```json
{
  category: Category
  subject: text
  created_at: date-time
  updated_at: date-time
  text: text
}
```

**Category** is an enumeration of "text tags". Once user has created a category - it should be accessible and reusable. Category is a plain text tag, which should be created on **note** create/update. There should be single instance of **category** for each identical text.

### Operations on Note

**Notes** can be created, updated, deleted and retrieved both via ID or all **notes** at once. User should be able to query **notes** by category as well.

User should be able to list all available **categories** (the ones,  which he had already created) and update them. In case the **category** is not used it should be deletable as well.

## Architecture requirements

There is no special requirement on architecture at the moment. The only ones are:

1. DB should be a separate process (either on local machine, in container or in cloud - local machine is preferable now)
2. There should NOT be service discovery (do not try to implement your own service discovery - too time consuming; just provide a db address to your app explicitly on each run)

## The way to present

In order to verify that everything works and demo it some preconfigured request should be created. They could be a simple *curl* requests, *postman-*defined or, if you have enough time, terminal cli (prefer Golang to implement)

## Technical requirements

- As much code as possible should be covered with UTs (if you find a way to launch DB in CI you can implement tests even with DB involved - this is not necessary and should be implemented last)
- Setup CI to your repo and use one of git-branching strategies (you free to choose, but describe your choice)
- Different frameworks could be used during implementation (but you'll have to provide points why to use them)
- It would be cool if you will be able to separate your work into multiple PRs and send each of them for review