# CSP Testing Site - Go Implementation

This is a Go web application that replicates the functionality of the original Node.js Express application for testing Content Security Policy (CSP) headers.

## Features

- **70+ endpoints** with different CSP configurations
- **Login endpoints** (`/login1` to `/login19`) with various CSP policies
- **Login endpoints with nonce** (`/login1x` to `/login19x`) for nonce-based CSP testing
- **CSP test endpoints** (`/csp1`, `/csp2`, etc.) for comprehensive CSP testing
- **Interactive HTML pages** for visual CSP testing

## Project Structure

```
├── main.go          # Main Go server implementation
├── index.html       # Primary test page with navigation
├── index1.html      # Alternative test page for specific endpoints
├── go.mod           # Go module definition
└── README.md        # This file
```

## Endpoints Overview

### Basic Login Endpoints (`/login1` - `/login19`)
Each endpoint sets different CSP policies to test various security configurations:
- `script-src`, `default-src`, `style-src` directives
- `'self'`, `'unsafe-inline'`, `'unsafe-eval'`, `'strict-dynamic'` keywords
- Different combinations for comprehensive testing

### Nonce-based Login Endpoints (`/login1x` - `/login19x`)
Similar to basic login endpoints but include nonce values in CSP policies for testing nonce-based script execution.

### CSP Test Endpoints (`/csp1`, `/csp2`, `/csp8`-`/csp28`, `/csp40`-`/csp46`)
Specialized endpoints for testing:
- Cache control headers (`no-store`)
- Case sensitivity in CSP directives
- Alternative HTML files
- Custom header names
- Various CSP policy combinations

## Running the Application

### Prerequisites
- Go 1.24.5 or later

### Environment Variables
- `PORT`: Server port (default: 3000)
- `HOST`: Server host (default: localhost)

### Start the Server

```bash
# Using default settings
go run main.go

# With custom port and host
PORT=8080 HOST=0.0.0.0 go run main.go
```

### Build and Run

```bash
# Build the application
go build -o csp-test-site main.go

# Run the built binary
PORT=3000 HOST=localhost ./csp-test-site
```

## Testing CSP Policies

1. **Visit the main page**: `http://localhost:3000/`
2. **Navigate to different endpoints** using the provided links
3. **Open browser DevTools** to observe:
   - CSP headers in Network tab
   - CSP violations in Console tab
   - Security warnings and errors

### Example Tests

- **Test inline scripts**: Visit `/login3` (allows `'unsafe-inline'`)
- **Test eval()**: Visit `/csp1` (allows `'unsafe-eval'`)
- **Test nonce**: Visit `/login1x` (includes nonce in CSP)
- **Test strict-dynamic**: Visit `/login5` (uses `'strict-dynamic'`)

## Key Differences from Node.js Version

1. **Static file serving**: Uses `http.ServeFile()` instead of Express static middleware
2. **Routing**: Uses `http.HandleFunc()` instead of Express routes
3. **Headers**: Uses `w.Header().Set()` instead of `res.setHeader()`
4. **Environment variables**: Uses `os.Getenv()` instead of `process.env`
5. **Server startup**: Uses `http.ListenAndServe()` instead of Express `app.listen()`

## CSP Policy Examples

### Basic Script Restriction
```
Content-Security-Policy: script-src 'self'; object-src 'none'
```

### Allow Inline Scripts
```
Content-Security-Policy: script-src 'unsafe-inline'; object-src 'none'
```

### Nonce-based Policy
```
Content-Security-Policy: script-src 'nonce-abc123'; object-src 'none'
```

### Strict Dynamic
```
Content-Security-Policy: script-src 'strict-dynamic'; object-src 'none'
```

## Development

### Adding New Endpoints

1. Create a new handler function or use existing handler functions
2. Register the route using `http.HandleFunc()`
3. Define the CSP policy for the endpoint
4. Test the endpoint with different browsers

### Handler Functions

- `homeHandler()`: Basic home page with default CSP
- `loginHandler()`: Generic handler with custom CSP policy
- `cspHandlerWithCache()`: Handler with CSP + cache control
- `cspHandlerWithCacheOnly()`: Handler with only cache control
- `cspHandlerWithCustomHeader()`: Handler with custom header name
- `cspHandlerWithDifferentFile()`: Handler serving alternative HTML file

## Browser Compatibility

Tested with:
- Chrome/Chromium
- Firefox
- Safari
- Edge

## Security Notes

⚠️ **Warning**: This application is designed for testing CSP policies and includes deliberately insecure configurations. Do not use these configurations in production environments.

## License

This project is for educational and testing purposes.
