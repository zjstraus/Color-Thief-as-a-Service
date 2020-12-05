# Color Thief as-a-Service
This is a very simple HTTP server that allows you to utilize workflows similar to [Color Thief](https://lokeshdhakar.com/projects/color-thief/)
in environments that may not be able to do image processing natively. It presents an HTTP POST endpoint with a JSON
body containing a URL of an image and will return an array of 10 colors as RGB values.

# Usage
Start the server and make POST requests to `/` with the below format; the response format is also shown below. The 
internal HTTP server defaults to serving on port 8080, but that can be overridden with the `port` flag.

## Examples
### Default port (8080)
```
./ctass.exe
2020/11/26 15:26:06 Starting HTTP listener on port 8080
```

### Custom port
```
./ctass.exe -port 8082
2020/11/26 15:28:04 Starting HTTP listener on port 8082
```

## Request Format
```json5
{
  // URL of an image
  "url": "https://i.scdn.co/image/ab67616d0000b2734f4b14080b3f4c8db28ee67f"
}
```

## Response Format
```json5
{
  // Colors contained in the image
  "palette": [
    {
      "r": 60,
      "g": 36,
      "b": 60
    },
    {
      "r": 104,
      "g": 74,
      "b": 106
    },
    {
      "r": 110,
      "g": 81,
      "b": 112
    },
    {
      "r": 111,
      "g": 82,
      "b": 113
    },
    {
      "r": 116,
      "g": 87,
      "b": 118
    },
    {
      "r": 167,
      "g": 136,
      "b": 188
    },
    {
      "r": 27,
      "g": 23,
      "b": 26
    },
    {
      "r": 28,
      "g": 24,
      "b": 27
    },
    {
      "r": 34,
      "g": 27,
      "b": 33
    },
    {
      "r": 59,
      "g": 35,
      "b": 59
    }
  ]
}
```