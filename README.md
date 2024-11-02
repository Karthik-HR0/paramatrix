# Paramatrix

Mining URLs from the depths of Web Archives for bug hunting, fuzzing, and advanced probing.


---

# Overview

Paramatrix is a powerful tool for bug bounty hunters, penetration testers, and developers interested in web security. It retrieves URLs for a given domain or list of domains directly from the Wayback Machine, filters out "boring" URLs, and allows you to focus on the ones with parameters—ideal for fuzzing, security probing, or gathering insights.

# Key Features

• Wayback Archive Mining: Fetch URLs directly from the Internet Archive’s Wayback Machine.

• URL Filtering: Ignores common static files, like .jpg and .css, so you get URLs that are more likely to be interesting.

• Customizable Placeholder: Replace parameter values with a custom placeholder to make URLs fuzz-ready.

• Retry Mechanism: Automatically retries in case of connectivity issues.

• User-Agent Rotation: Uses multiple user-agents to avoid simple bot detection.

• Proxy Support: Route requests through a proxy if desired.



---

# Installation

To get started with Paramatrix, make sure you have Go installed. Then, clone the repository and build it:
```git clone https://github.com/Karthik-HR0/paramatrix ``
cd paramatrix
go build -o paramatrix
mv paramatrix /bin/bash```


Alternatively, install directly using:

``` go install github.com/Karthik-HR0/paramatrix@latest
```

# Usage

Paramatrix provides flexible options to fetch and process URLs:

``` paramatrix -d example.com ```

# Command-Line Options

-d DOMAIN: Specify a domain to fetch URLs for.

-l LIST: Provide a file containing a list of domains (one per line).

-s: Stream URLs directly to the terminal for real-time inspection.

--proxy PROXY: Use a specified proxy for HTTP requests.

-p PLACEHOLDER: Replace parameter values with a custom placeholder (default is "FUZZ") 



# Examples

Get started with some practical usage scenarios for Paramatrix:

Fetch URLs for a Single Domain:

``` paramatrix -d example.com ```

Fetch URLs for Multiple Domains from a File:

``` paramatrix -l domains.txt ```

Stream URLs in Real-Time to Terminal:

``` paramatrix -d example.com -s ```

Use a Proxy for Web Requests:

``` paramatrix -d example.com --proxy '127.0.0.1:8080' ```

Customize Placeholder for Parameter Values:

``` paramatrix -d example.com -p '<injection>' ```



---

# Contributing

We welcome contributions! Feel free to submit feature suggestions, report issues, or open pull requests. Let’s make Paramatrix even more powerful for the bug hunting community!
