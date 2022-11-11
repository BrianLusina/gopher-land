package nginx

// Nginx will serve as a proxy to the server
type Nginx struct {
	app                *App
	maxAllowedRequests int
	rateLimiter        map[string]int
}

func NewNginxServer(maxAllowedRequests int) *Nginx {
	return &Nginx{
		app:                &App{},
		maxAllowedRequests: maxAllowedRequests,
		rateLimiter:        make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.app.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequests {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
