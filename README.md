- `psql -U root`
- app から db に繋げる場合は、container 名を host に指定しないとだめなので、container_name で固定しておいたほうが確実に繋げられる
- hot reload にhttps://github.com/cosmtrek/air

```bash
gobin() {
  $(go env GOPATH)/bin/$1
}
```

`gobin air`

# postgres

- https://eng-entrance.com/postgresql-role
