```graphql
mutation test {
  store1: store(entry: {
    key: "test",
    value: "test1"
  }) {
    key
  },
  store2: store(entry: {
    key: "test2",
    value: "test4"
  }) {
    value
  }
}

query test1 {
  get1: get(key: "test") {
    value
  },
  get2 :get(key: "test2") {
    key,
  }
}
```
