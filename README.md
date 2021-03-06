# argpas

an utils to parse os arguments 👻
> argument and name is case insensitive

```go
ParseToInterface(pointer interface{},args []string) error{}
Parse(arg, name string) (string, error){}
ParseToInt(arg, name string) (int, error){}
ParseToStr(arg, name string) (string, error){}
ParseToBool(arg, name string) (bool, error){}
ParseToFloat32(arg, name string) (float32, error){}
ParseToFloat64(arg, name string) (float64, error){}
```


you can find the example in `main/main.go`
