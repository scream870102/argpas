# argpas

![badge](https://img.shields.io/github/v/tag/scream870102/argpas?color=E861A4&label=version&logo=github&logoColor=FFFFFF&style=flat)
![badge](https://img.shields.io/github/go-mod/go-version/scream870102/argpas?color=00ADD8&logo=go&logoColor=FFFFFF&style=flat)
![badge](https://img.shields.io/github/last-commit/scream870102/argpas?color=B682D2&label=last%20commit&style=flat)  

an utils to parse os arguments 👻
> argument and name is case insensitive

## Documentation

[Get the full document here](https://pkg.go.dev/github.com/scream870102/argpas)

## How to use

```go
var momoko string
arg := os.Args[1]
if result,err:= argpas.ParseToStr(arg,"StringArg");err!=nil{
    fmt.Println(err.Error())
}else{
    momoko = result
}
fmt.Println(momoko)
```

you can also define a struct which holds all arguments.  
And use `ParseToInterface` to deal them at the same time.

```go
type arguments struct {
 StringArg  string
 BoolArg    bool
 IntArg     int
 FloatArg   float64
 TriggerArg argpas.Trigger
}
func main(){
 var momoko arguments
 args := os.Args[1:]
 if err := argpas.ParseToInterface(&momoko, args); err != nil {
  fmt.Println(err.Error())
 }
 fmt.Println(momoko)
}

```

you can find the full example in `main/main.go`
