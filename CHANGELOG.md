# CHANGELOG


## 0.3.0
### Added
- `ParseToFloat32(arg, name string) (float32, error){}`
- `ParseToFloat64(arg, name string) (float64, error){}`
- `ParseToInterface(pointer interface{}, args []string) error {}`
    - Allow Type in struct
        - int
        - float64
        - string
        - bool
        - argpas.Trigger
- new type `Trigger`
    - can use in struct as argument type
    - Trigger is useful os arg like `--Momoko` `--Help`.

### Updated
Rename function
- `ParseArg -> Parse`
- `ParseArgToInt -> ParseToInt`
- `ParseArgToStr -> ParseToStr`
- `ParseArgToBool -> ParseToBool`
## 0.2.0

### Added
- `ParseArgToBool(arg,name string) (bool,error){}`
    - Allow Values `True,TRUE,true,T,t,False,FALSE,false,F,f,0`
### Update
- Now space is allow in argument. parser will ignore it.
    - `--month = 12` This will work correctly right now

## 0.1.0

### Added
- `ParseArgToInt(arg, name string) (int, error){}`
- `ParseArgToStr(arg, name string) (string, error){}`
- `ParseArg(arg, name string) (string, error){}`