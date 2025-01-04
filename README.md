# PDynIP: Dynamic DNS Updater

`PDynIP` is a lightweight utility to update DNS records dynamically when your public IP address changes. It supports both **Cloudflare** and **Namecheap** DNS providers and is packaged as a Docker container for easy deployment.

## Why is this here ?
I recently switched ISPs and lost my static IP, leaving me with a dynamic IP instead. This made accessing certain services at home, like VPNs, much more difficult while on the road. I explored dynamic DNS services, which most modern routers support out of the box, but I found them clunky and inconvenient. Additionally, I wanted to use my own domain, which I was already paying for. Many dynamic DNS providers charged extra for custom domains, and in my case, upgrading to a static IP would have been more cost-effective.

To address this, I created a solution tailored to my specific needs: a tool to update my Namecheap and Cloudflare hostnames with my current IP address. If other dynamic DNS services offer APIs, I’m happy to add support for them. Alternatively, feel free to contribute and add this functionality yourself!

---

## Features

- Automatically detects public IP changes.
- Updates DNS records with the new IP for specified domains and hostnames.
- Supports multiple hostnames under a single domain.
- Works with Cloudflare and Namecheap DNS APIs.
- Simple logging for easy monitoring.

---

## Configuration

The application is configured via environment variables:

| Variable            | Description                                        | Required for      |
|---------------------|----------------------------------------------------|-------------------|
| `PDYNIP_PROVIDER`   | The DNS provider (`cloudflare` or `namecheap`).    | All               |
| `PDYNIP_API_KEY`    | The API key or Dynamic DNS password.               | All               |
| `PDYNIP_EMAIL`      | Email associated with the Cloudflare account.      | Cloudflare only   |
| `PDYNIP_DOMAIN`     | The domain name to update (e.g., `example.com`).   | All               |
| `PDYNIP_HOSTNAMES`  | Comma-separated list of hostnames to update.       | All               |
| `PDYNIP_CHECK_INTERVAL` | Interval (in seconds) to check for IP changes. | All (optional)    |

---

## Docker Examples

### **Running with Namecheap**
```bash
docker run -d \
  -e PDYNIP_PROVIDER=namecheap \
  -e PDYNIP_API_KEY=<YOUR_NAMECHEAP_DDNS_PASSWORD> \
  -e PDYNIP_DOMAIN=example.com \
  -e PDYNIP_HOSTNAMES=www,api \
  -e PDYNIP_CHECK_INTERVAL=300 \
  --name pdynip-namecheap \
  ghcr.io/pradt/pdynip:latest
```
### **Running with Cloudflare**
```bash 
docker run -d \
  -e PDYNIP_PROVIDER=cloudflare \
  -e PDYNIP_API_KEY=<YOUR_CLOUDFLARE_API_KEY> \
  -e PDYNIP_EMAIL=<YOUR_CLOUDFLARE_EMAIL> \
  -e PDYNIP_DOMAIN=example.com \
  -e PDYNIP_HOSTNAMES=www,api \
  -e PDYNIP_CHECK_INTERVAL=300 \
  --name pdynip-cloudflare \
  ghcr.io/pradt/pdynip:latest
  ```

## Portainer Stack file example 
### Namecheap Stack Example
```yaml
version: "2.1"
services:
  pdynip-namecheap:
    image: ghcr.io/pradt/pdynip:latest
    container_name: pdynip-namecheap
    environment:
      - PDYNIP_PROVIDER=namecheap
      - PDYNIP_API_KEY=<YOUR_NAMECHEAP_DDNS_PASSWORD>
      - PDYNIP_DOMAIN=example.com
      - PDYNIP_HOSTNAMES=www,api
      - PDYNIP_CHECK_INTERVAL=300
    restart: always
```

### Cloudlare Stack Example
```yaml
version: "2.1"
services:
  pdynip-cloudflare:
    image: ghcr.io/pradt/pdynip:latest
    container_name: pdynip-cloudflare
    environment:
      - PDYNIP_PROVIDER=cloudflare
      - PDYNIP_API_KEY=<YOUR_CLOUDFLARE_API_KEY>
      - PDYNIP_EMAIL=<YOUR_CLOUDFLARE_EMAIL>
      - PDYNIP_DOMAIN=example.com
      - PDYNIP_HOSTNAMES=www,api
      - PDYNIP_CHECK_INTERVAL=300
    restart: always

```

## Logs and Monitoring
If you are not sure what's happening, please check out the logs it will tell you step by step what is happening. If you run into any problems I wouuld need to see this. 
The container logs provide real-time updates, making it easy to track the application's activity:

- **IP Change Detection**:  
  Example: `Detected IP change from 203.0.113.1 to 198.51.100.1`.
  
- **Successful Updates**:  
  Example: `Successfully updated hostname www to IP 198.51.100.1`.
  
- **Errors**:  
  Example: `Error updating DNS record for hostname www: API error: Invalid API key`.

To view the logs of a running container, use the following command:
```bash
docker logs -f <container-name>
```

Example: 
```bash
docker logs -f pdynip-namecheap
```

## Help and Support

The easiest way to run this is with Docker (instructions above). However, if you prefer to download the source code and compile it yourself, you’re welcome to do so! Just keep in mind that support for manually compiled setups may be limited.

If you're using Docker and encounter any issues, feel free to open an issue on the project’s repository. When reporting a problem, please include relevant log information—just make sure to omit any sensitive IP details if you’d like to keep them confidential.
