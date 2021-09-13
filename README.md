# Requirements

Necessary programs to run the api

```
- go
```

# Run

```
» make run
```

# Concerns

- The memory repository should use a sync map
- If disk storage was instead used, some operations like fetching and updating the `expiration_time` should be done inside a
  transaction
- Parameter sanitization and validation is not done. The api should check the `name` and `snippet` length. It also
  should validate the `expires_in` parameter (check that is a positive value lower than a given `max_expiration_time`)
- The user can give all the likes that he wants to a snippet.
- The user has access to all snippets and not only the ones created by him

# Decisions

- Return snippet if update expiration operation fails
- Can't create another snippet while another one is still using that name
- Have a controller - service - repository code organization. The controller handles the requests and converts them into
  domain structs. The service handles the business logic (in this case the expiration time decisions). The repository
  layer stores the snippets

# Future work

- Improve errors (add more custom errors, be more precise with errors)
- Add testing
- Add disk storage