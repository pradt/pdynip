# PDynIP: Dynamic DNS Updater

`PDynIP` is a lightweight utility to update DNS records dynamically when your public IP address changes. It supports both **Cloudflare** and **Namecheap** DNS providers and is packaged as a Docker container for easy deployment.

## Why is this here ?
I changed ISPs can I no longer had a static IP, instead I had a dynamic IP. This made calling specific services at home a lot a more difficult when I'm on the road such as VPNs. I had a look at dynamic ip services, which most routers now support ootb I found them to be really clunky. I also just wanted to use my own domain which I was already paying for (Dynamic IP services charged extra for this, and getting a static IP was going to cheaper in my case).

Anyway I made this for my specific usecase, to update namecheap and cloudflare hostanems with my  IP. I'm happy to add other services if they have an API or if you want to add this feature please go right ahead. 

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

The best way to run this is via docker, see above. If you want to download the source and compile it yourself then you are welcome to do so. Please note that support would be limited in this case. 

If you are running it via Docker, and you run into any problems then open an issue. You should include the log information ( careful to ommit any ip information if you want it to be confidential). 
