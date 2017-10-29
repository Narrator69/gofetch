# gofetch
Simple http request library based on gorequest

### Usage

#### With settings
```content, _ := gofetch.Fetch{UserAgent: "Banana", Proxy: "http://127.0.0.1:8888"}.Go("GET", "http://vk.com", []gofetch.Header{}, "")```

#### Without settings
```content, _ := gofetch.Fetch{}.Go("GET", "http://vk.com", []gofetch.Header{}, "")```
