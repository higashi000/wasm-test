# wasm-test

## Usage
```
$ git clone https://github.com/higashi000/wasm-test
$ cd wasm-test
```

### fish<br>
```
$ env GOOS=js GOARCH=wasm go build -o main.wasm
$ set -x WHICHGO (go env GOROOT)
$ cp $WHICHGO/misc/wasm/wasm_exec.js .
```
### bash<br>
```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ WHICHGO=`go env GOROOT`
$ cp $WHICHGO/misc/wasm/wasm_exec.js .
```


### execution
```
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

Please access http://localhost:8080.
