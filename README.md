# argpas

![](https://img.shields.io/github/v/tag/scream870102/argpas?color=E861A4&label=version&logo=github&logoColor=FFFFFF&style=flat)
![](https://img.shields.io/github/go-mod/go-version/scream870102/argpas?color=00ADD8&logo=go&logoColor=FFFFFF&style=flat)
![](https://img.shields.io/github/last-commit/scream870102/argpas?color=B682D2&label=last%20commit&style=flat)

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
