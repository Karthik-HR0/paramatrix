package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

const banner = `


__________                         _____          __         .__        
\______   \_____ ____________     /     \ _____ _/  |________|__|__  ___
 |     ___/\__  \\_  __ \__  \   /  \ /  \\__  \\   __\_  __ \  \  \/  /
 |    |     / __ \|  | \// __ \_/    Y    \/ __ \|  |  |  | \/  |>    < 
 |____|    (____  /__|  (____  /\____|__  (____  /__|  |__|  |__/__/\_ \
                \/           \/         \/     \/                     \/
     
     
                                              with <3 by @Karthik-HR0 
                                              `                                                             
                                                                                         

var (
	hardcodedExtensions = []string{
		".jpg", ".jpeg", ".png", ".gif", ".pdf", ".svg", ".json",
		".css", ".js", ".webp", ".woff", ".woff2", ".eot", ".ttf",
		".otf", ".mp4", ".txt",
	}

	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
	}

	maxRetries = 3
)

type paramatrix struct {
	client      *http.Client
	proxy       string
	placeholder string
}

func Newparamatrix(proxy, placeholder string) *paramatrix {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Fatalf("Invalid proxy URL: %v", err)
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	return &paramatrix{
		client:      client,
		proxy:       proxy,
		placeholder: placeholder,
	}
}

func (ps *paramatrix) fetchWithRetry(waybackURI string) ([]string, error) {
	var urls []string
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		req, err := http.NewRequest("GET", waybackURI, nil)
		if err != nil {
			lastErr = err
			continue
		}

		// Set random User-Agent
		req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])

		resp, err := ps.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
			time.Sleep(5 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			continue
		}

		urls = strings.Split(string(body), "\n")
		return urls, nil
	}

	return nil, fmt.Errorf("failed after %d retries, last error: %v", maxRetries, lastErr)
}

func (ps *paramatrix) ProcessDomain(domain string, streamOutput bool) {
	log.Printf("\033[94m[INFO]\033[94m Fetching URLs for \033[36m%s\033[94m", domain)

	waybackURI := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s/*&output=txt&collapse=urlkey&fl=original&page=/", domain)
	urls, err := ps.fetchWithRetry(waybackURI)
	if err != nil {
		log.Printf("Error fetching URLs for %s: %v", domain, err)
		return
	}

	log.Printf("\033[94m[INFO]\033[94m Found \033[32m%d\033[94m URLs for \033[36m%s\033[94m", len(urls), domain)

	cleanedURLs := ps.cleanURLs(urls)
	log.Printf("\033[94m[INFO]\033[94m Cleaning URLs for \033[36m%s\033[94m", domain)
	log.Printf("\033[94m[INFO]\033[94m Found \033[32m%d\033[94m URLs after cleaning", len(cleanedURLs))
	log.Printf("\033[94m[INFO]\033[94m Extracting URLs with parameters")

	if err := os.MkdirAll("results", 0755); err != nil {
		log.Fatalf("Error creating results directory: %v", err)
	}

	resultFile := path.Join("results", domain+".txt")
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalf("Error creating result file: %v", err)
	}
	defer file.Close()

	for _, u := range cleanedURLs {
		if strings.Contains(u, "?") {
			fmt.Fprintln(file, u)
			if streamOutput {
				fmt.Println(u)
			}
		}
	}

	log.Printf("\033[94m[INFO]\033[94m Saved cleaned URLs to \033[36m%s\033[94m", resultFile)
}

func (ps *paramatrix) cleanURLs(urls []string) []string {
	seen := make(map[string]bool)
	var cleanedURLs []string

	for _, u := range urls {
		if hasExtension(u) {
			continue
		}

		cleanedURL := ps.cleanURL(u)
		if cleanedURL != "" && !seen[cleanedURL] {
			cleanedURLs = append(cleanedURLs, cleanedURL)
			seen[cleanedURL] = true
		}
	}

	return cleanedURLs
}

func (ps *paramatrix) cleanURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	// Remove standard ports
	if (parsedURL.Scheme == "http" && parsedURL.Port() == "80") ||
		(parsedURL.Scheme == "https" && parsedURL.Port() == "443") {
		host := parsedURL.Hostname()
		parsedURL.Host = host
	}

	// Clean query parameters
	query := parsedURL.Query()
	for k := range query {
		query.Set(k, ps.placeholder)
	}
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String()
}

func hasExtension(url string) bool {
	for _, ext := range hardcodedExtensions {
		if strings.HasSuffix(url, ext) {
			return true
		}
	}
	return false
}
func main() {
    // Remove timestamp from log output
    log.SetFlags(0)

    fmt.Printf("\033[94m%s\033[94m\n", banner)

    domain := flag.String("d", "", "Domain name to fetch related URLs for")
    listFile := flag.String("l", "", "File containing a list of domain names")
    streamOutput := flag.Bool("s", false, "Stream URLs on the terminal")
    proxy := flag.String("proxy", "", "Set the proxy address for web requests")
    placeholder := flag.String("p", "FUZZ", "Placeholder for parameter values")
    flag.Parse()

    if *domain == "" && *listFile == "" {
        fmt.Println(`usage: paramatrix [-h] [-d DOMAIN] [-l LIST] [-s] [--proxy PROXY] [-p PLACEHOLDER]

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
`)
        os.Exit(1)
    }

    if *domain != "" && *listFile != "" {
        log.Fatal("Please provide either the -d option or the -l option, not both")
    }

    spider := Newparamatrix(*proxy, *placeholder)

    if *domain != "" {
        spider.ProcessDomain(*domain, *streamOutput)
    }

    if *listFile != "" {
        domains := readDomainList(*listFile)
        for _, domain := range domains {
            spider.ProcessDomain(domain, *streamOutput)
        }
    }
}

func readDomainList(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var domains []string
	seen := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		domain = strings.ToLower(domain)
		domain = strings.TrimPrefix(domain, "https://")
		domain = strings.TrimPrefix(domain, "http://")
		if domain != "" && !seen[domain] {
			domains = append(domains, domain)
			seen[domain] = true
		}
	}

	return domains
}
