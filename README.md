<h1 align="center">
    paramatrix
  <br>
</h1>

<h4 align="center">  Mining URLs from dark corners of Web Archives for bug hunting/fuzzing/further probing </h4>

<p align="center">
  <a href="#about">📖 About</a> •
  <a href="#installation">🏗️ Installation</a> •
  <a href="#usage">⛏️ Usage</a> •
  <a href="#examples">🚀 Examples</a> •
  <a href="#contributing">🤝 Contributing</a> •
</p>


![paramatrix](https://github.com/user-attachments/assets/92e2b685-285d-4151-a95d-2fd7ca21b56f)

## About

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

## Installation

To install `paramatrix`, follow these steps:

```sh
git clone https://github.com/Karthik-HR0/paramatrix
cd paramatrix
go build -o paramatrix
sudo cp -r paramatrix /usr/bin/
```
# Direct installation 
``` go install github.com/Karthik-HR0/paramatrix@latest```
`` cd go/bin
sudo cp -r paramatrix /usr/bin/ ``

## Usage

Paramatrix offers flexible options for fetching and processing URLs:

` paramatrix -d example.com `

Command-Line Options

-d DOMAIN: Specify a single domain to fetch URLs for.

-l LIST: Provide a file with a list of domains (one per line).

-s: Stream URLs to the terminal in real-time.

--proxy PROXY: Use a specified proxy for HTTP requests.

-p PLACEHOLDER: Set a custom placeholder for parameter values in URLs (default is "FUZZ").



---

# Examples

Get started with some practical scenarios for using Paramatrix:

Fetch URLs for a Single Domain:

` paramatrix -d example.com `

Fetch URLs for Multiple Domains from a File:

` paramatrix -l domains.txt `

Stream URLs in Real-Time to the Terminal:

` paramatrix -d example.com -s `

Use a Proxy for Web Requests:

` paramatrix -d example.com --proxy '127.0.0.1:8080' `

Customize Placeholder for Parameter Values:

`paramatrix -d example.com -p '<injection>'`



---

## 🤝 Contributing

We welcome contributions from the community! To contribute, please fork the repository, make your changes, and submit a pull request. Feel free to report issues or suggest new features to enhance Paramatrix.

Happy hacking and happy hunting!
