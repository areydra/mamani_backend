1. run command => go env
    - if GOARCH value is not amd64, change it to linux by run command => set GOARCH=amd64
    - if GOOS=windows, change it to linux by run command => set GOOS=windows
2. main.go (top of func main)
    - put => utils.LoadEnv("production")
	- put => gin.SetMode(gin.ReleaseMode)
3. run command to build => go build -o mamani-backend main.go

