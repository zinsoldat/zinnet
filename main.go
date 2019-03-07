package main

func main() {
	config := &ServerConfig{
		Port: 3000,
	}

	server := NewServer(config)
	server.Start()
}
