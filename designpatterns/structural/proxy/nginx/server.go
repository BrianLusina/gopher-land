package nginx

type server interface {
	handleRequest(string, string) (int, string)
}
