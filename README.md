# CSP Testing Site - Go Implementation

This is a Go web application that replicates the functionality of the original Node.js Express application for testing Content Security Policy (CSP) headers.

## Features

- **70+ endpoints** with different CSP configurations
- **Login endpoints** (`/login1` to `/login19`) with various CSP policies
- **Login endpoints with nonce** (`/login1x` to `/login19x`) for nonce-based CSP testing
- **CSP test endpoints** (`/csp1`, `/csp2`, etc.) for comprehensive CSP testing
- **Interactive HTML pages** for visual CSP testing
- **Automated CI/CD** with GitHub Actions
- **Docker support** for containerized deployment

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

## Deployment

This project includes automated deployment to Linode using GitHub Actions and manual deployment scripts.

### GitHub Actions Deployment

The repository includes a complete CI/CD pipeline that automatically:

1. **Tests** the application
2. **Builds** for Linux AMD64
3. **Deploys** to your Linode server

#### Setup GitHub Actions

1. **Fork/Clone** this repository to your GitHub account

2. **Set up GitHub Secrets** in your repository settings:
   ```
   LINODE_HOST=your-server-ip
   LINODE_USER=your-username
   LINODE_SSH_PRIVATE_KEY=your-private-key
   ```

3. **Generate SSH Key Pair** (if you don't have one):
   ```bash
   ssh-keygen -t rsa -b 4096 -C "github-actions"
   ```

4. **Add Public Key** to your Linode server:
   ```bash
   # On your Linode server
   mkdir -p ~/.ssh
   echo "your-public-key" >> ~/.ssh/authorized_keys
   chmod 600 ~/.ssh/authorized_keys
   chmod 700 ~/.ssh
   ```

5. **Add Private Key** to GitHub Secrets as `LINODE_SSH_PRIVATE_KEY`

6. **Push to main/master branch** to trigger deployment

#### GitHub Actions Workflow

The workflow (`/.github/workflows/deploy.yml`) includes:

- **Test Stage**: Runs Go tests and linting
- **Build Stage**: Compiles for Linux and creates deployment package
- **Deploy Stage**: Deploys to Linode with zero-downtime
- **Health Check**: Verifies deployment success

### Manual Deployment

For manual deployments, use the included deployment script:

```bash
# Deploy to production
./deploy.sh production

# Deploy to staging
./deploy.sh staging

# Build locally only
./deploy.sh local
```

#### Setup Manual Deployment

1. **Copy environment files**:
   ```bash
   cp .env.production.example .env.production
   cp .env.staging.example .env.staging
   ```

2. **Configure your servers**:
   ```bash
   # Edit .env.production
   DEPLOY_HOST=your-linode-ip
   DEPLOY_USER=your-username
   DEPLOY_PORT=8080
   ```

3. **Setup SSH access** to your servers

4. **Run deployment**:
   ```bash
   ./deploy.sh production
   ```

### Docker Deployment

#### Build and Run Locally

```bash
# Build the image
docker build -t csp-test-site .

# Run the container
docker run -p 3000:8080 csp-test-site
```

#### Docker Compose

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Linode Server Setup

#### Initial Server Setup

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y nginx certbot python3-certbot-nginx

# Create application user
sudo useradd -m -s /bin/bash cspuser
sudo usermod -aG sudo cspuser

# Setup firewall
sudo ufw allow ssh
sudo ufw allow http
sudo ufw allow https
sudo ufw enable
```

#### Nginx Configuration

```bash
# Copy nginx configuration
sudo cp nginx.conf /etc/nginx/nginx.conf

# Test configuration
sudo nginx -t

# Restart nginx
sudo systemctl restart nginx
sudo systemctl enable nginx
```

#### SSL Certificate (Optional)

```bash
# Get SSL certificate
sudo certbot --nginx -d your-domain.com

# Auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

### Monitoring and Maintenance

#### Check Application Status

```bash
# Check service status
sudo systemctl status csp-test-site

# View logs
sudo journalctl -u csp-test-site -f

# Restart service
sudo systemctl restart csp-test-site
```

#### Health Check

```bash
# Manual health check
curl http://your-server:8080/

# Automated health check (included in deployment)
```

#### Backup Management

The deployment automatically:
- Creates backups before each deployment
- Keeps the last 5 backups
- Cleans up old deployment artifacts

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `HOST` | Server host | `localhost` |
| `DEPLOY_HOST` | Target server IP | - |
| `DEPLOY_USER` | SSH username | - |
| `DEPLOY_PORT` | Application port on server | `8080` |

### Troubleshooting

#### Deployment Issues

```bash
# Check GitHub Actions logs
# Go to Actions tab in your repository

# Check server logs
sudo journalctl -u csp-test-site -n 50

# Check nginx logs
sudo tail -f /var/log/nginx/error.log
```

#### Common Issues

1. **SSH Connection Failed**:
   - Verify SSH key is correctly added to GitHub Secrets
   - Check server firewall settings
   - Ensure user has sudo privileges

2. **Service Won't Start**:
   - Check binary permissions: `sudo chmod +x /opt/csp-test-site/csp-test-site`
   - Verify port availability: `sudo netstat -tlnp | grep 8080`
   - Check systemd logs: `sudo journalctl -u csp-test-site`

3. **Health Check Failed**:
   - Verify application is listening on correct port
   - Check firewall rules
   - Test local connectivity: `curl localhost:8080`

## License

This project is for educational and testing purposes.
