# [Postshift API](https://post-shift.ru/api/) wrapper in golang

## Installing

```
go get github.com/DiGregory/golang-postshift-api
```

## Create new random mail
```go
mail, err := NewMail("","")
//lhnimqve05@post-shift.ru 3f5c2aa5e42e68fefb5e34a1afce3a84
```


## Create new mail with given address

```go
mail, err := NewMail("golang","post-shift.ru")
//golang@post-shift.ru 8d28fbba828f30ac3b7ee52b7008c1ea
```
