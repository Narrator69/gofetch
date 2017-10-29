# gofetch
Simple http request library based on gorequest

### Usage

#### With settings
```content, _ := gofetch.Fetch{UserAgent: "Banana", Proxy: "http://127.0.0.1:8888"}.Go("GET", "http://vk.com", []gofetch.Header{}, "")```

#### With body
```content, _ := gofetch.Fetch{}.Go("POST", "http://vk.com", []gofetch.Header{}, "query=delta")```

#### With headers
```content, _ := gofetch.Fetch{}.Go("GET", "http://vk.com", []gofetch.Header{{"Referer", "Banana"}}, "query=delta")```

#### Minimalist
```content, _ := gofetch.Fetch{}.Go("GET", "http://vk.com", []gofetch.Header{}, "")```
