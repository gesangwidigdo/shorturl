# ShortURL

Simple URL Shortener

### Request Body

**Endpoint:** `{{ root_url }}/api/url/shorten`

1. Without Custom Short URL

    ```js
    {
        "original_url": "https://www.example.com/",
    }
    ```

2. With Custom Short URL

    ```js
    {
        "original_url": "https://www.example.com/",
        "short_url": "example"
    }
    ```

### How to use shortened URL?

```
{{ root_url }}/:short_url
```
