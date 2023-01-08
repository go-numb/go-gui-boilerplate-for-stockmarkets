# Desktop App for Requests

![Screenshot](https://github.com/go-numb/go-gui-boilerplate-for-stockmarkets/blob/master/frontend/src/assets/screenshot.png)

- Go Serve the Vue static dir.
- Vue3 build to dist.
- Choose Electron or Lorca views GUI, at go served localhost.
- Blank Methods, choose go api or axios in html.


## Cross Compile
※ **Describe absolute paths with import filepath**  
※ **abs, err := filepath.Abs("./")**
``` cmd
$ GOOS=darwin GOARCH=amd64
$ go build main.app main.go // for Mac
$ GOOS=windows GOARCH=amd64
$ go build main.go // for Windows
```