# CHANGELOG

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