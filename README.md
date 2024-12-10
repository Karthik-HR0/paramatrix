<h1 align="center">paramatrix</h1>

> **paramatrix**  - Mining URLs from the hidden corners of Wayback Archives for bug hunting, fuzzing, and security probing.
---


  

![paramatrix png](https://github.com/user-attachments/assets/6775989d-222c-46cc-be5a-68bc0142a2e3)
---


`paramatrix` allows you to fetch URLs related to any domain or a list of domains from Wayback achives. It filters out "boring" URLs, allowing you to focus on the ones that matter the most.

## Overview
**Paramatrix** is a powerful tool designed for bug bounty hunters, penetration testers, and developers interested in web security. It fetches URLs associated with a specific domain (or list of domains) directly from the Wayback Machine. By filtering out "boring" URLs, it allows you to focus on those with parameters—ideal for fuzzing, security probing, and collecting insights.

### Key Features
- **Wayback Archive Mining**: Fetch URLs directly from the Internet Archive’s Wayback Machine.
- **URL Filtering**: Ignores common static files, such as `.jpg`, `.css`, etc., providing you with URLs that are more likely to be interesting for testing.
- **Customizable Placeholder**: Replace parameter values with a custom placeholder, making URLs fuzz-ready.
- **Retry Mechanism**: Automatically retries when connectivity issues occur.
- **User-Agent Rotation**: Rotates user-agents to minimize detection by anti-bot mechanisms.
- **Proxy Support**: Allows requests to be routed through a proxy if specified.

---


---

## installation 
```bash 
go install github.com/Karthik-HR0/paramatrix@latest
```

# Uses Examples

Get started with some practical scenarios for using Paramatrix:

Fetch URLs for a Single Domain:

```bash
paramatrix -d example.com
 ```

Fetch URLs for Multiple Domains from a File:

 ```bash
paramatrix -l domains.txt
 ```

Stream URLs in Real-Time to the Terminal:

``` bash
paramatrix -d example.com -s
 ```

Use a Proxy for Web Requests:

```bash
paramatrix -d example.com --proxy '127.0.0.1:8080'
```

Customize Placeholder for Parameter Values:

```bash
paramatrix -d example.com -p "<injection>"
```

<h3>Arguments</h3><table>
  <thead>
    <tr>
      <th>Argument</th>
      <th>Description</th>
      <th>Default</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>-d DOMAIN</code></td>
      <td>Specify a single domain to fetch URLs for</td>
      <td>None</td>
    </tr>
    <tr>
      <td><code>-l LIST</code></td>
      <td>Provide a file with a list of domains (one per line)</td>
      <td>None</td>
    </tr>
    <tr>
      <td><code>-s</code></td>
      <td>Stream URLs in real-time to the terminal</td>
      <td><code>false</code></td>
    </tr>
    <tr>
      <td><code>--proxy PROXY</code></td>
      <td>Use a specified proxy for HTTP requests</td>
      <td>None</td>
    </tr>
    <tr>
      <td><code>-p PLACEHOLDER</code></td>
      <td>Set a custom placeholder for parameter values</td>
      <td><code>FUZZ</code></td>
    </tr>
  </tbody>
</table>

---
```
./paramatrix



__________                         _____          __         .__
\______   \_____ ____________     /     \ _____ _/  |________|__|__  ___
 |     ___/\__  \\_  __ \__  \   /  \ /  \\__  \\   __\_  __ \  \  \/  /
 |    |     / __ \|  | \// __ \_/    Y    \/ __ \|  |  |  | \/  |>    <
 |____|    (____  /__|  (____  /\____|__  (____  /__|  |__|  |__/__/\_ \
                \/           \/         \/     \/                     \/


                                              with <3 by @Karthik-HR0

usage: paramatrix [-h] [-d DOMAIN] [-l LIST] [-s] [--proxy PROXY] [-p PLACEHOLDER]

Mining URLs from dark corners of Web Archives

options:
  -h, --help            show this help message and exit
  -d DOMAIN, --domain DOMAIN
                        Domain name to fetch related URLs for.
  -l LIST, --list LIST  File containing a list of domain names.
  -s, --stream          Stream URLs on the terminal.
  --proxy PROXY         Set the proxy address for web requests.
  -p PLACEHOLDER, --placeholder PLACEHOLDER
                        placeholder for parameter values
```
<p align="center">
Built with ❤️ by <a href="https://github.com/Karthik-HR0">@Karthik-HR0</a>
</p>
