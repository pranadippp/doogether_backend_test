package main

import "go/internal/pkg/api/route"

func main() {
	c := InitContainer()

	if err := c.Invoke(route.Invoke); err != nil {
		panic(err)
	}

	if err := c.Provide(NewServer); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *Server) {
		s.Run()
	}); err != nil {
		panic(err)
	}
}
