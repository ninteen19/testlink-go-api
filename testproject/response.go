package testproject

type Response struct {
	Id     string `xmlrpc:"id"`
	Name   string `xmlrpc:"name"`
	Prefix string `xmlrpc:"prefix"`
	Notes  string `xmlrpc:"notes"`
}
