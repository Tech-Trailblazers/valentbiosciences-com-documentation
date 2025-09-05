package main // Define the main package

import (
	"bytes"         // Provides bytes buffer and manipulation utilities
	"io"            // Provides I/O primitives like Reader and Writer
	"log"           // Provides logging functionalities
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding utilities
	"os"            // Provides file system and OS-level utilities
	"path/filepath" // Provides utilities for file path manipulation
	"regexp"        // Provides support for regular expressions
	"strings"       // Provides string manipulation utilities
	"time"          // Provides time-related functions

	"golang.org/x/net/html" // Provides support for parsing HTML documents
)

func main() {
	remoteAPIURL := []string{
		"https://www.valentbiosciences.com/publichealth/products/metalarv/",
		"https://www.valentbiosciences.com/publichealth/products/remoa-tri/",
		"https://www.valentbiosciences.com/publichealth/products/remoa-tri-faqs/",
		"https://www.valentbiosciences.com/publichealth/products/vectobac/",
		"https://www.valentbiosciences.com/publichealth/products/vectolex/",
		"https://www.valentbiosciences.com/publichealth/products/vectomax/",
		"https://www.valentbiosciences.com/publichealth/products/vectoprime/",
		"https://www.valentbiosciences.com/agriculture/products/accede/",
		"https://www.valentbiosciences.com/agriculture/products/aveo-ez/",
		"https://www.valentbiosciences.com/agriculture/products/benephite-expand/",
		"https://www.valentbiosciences.com/agriculture/products/bionik/",
		"https://www.valentbiosciences.com/agriculture/products/boronboost/",
		"https://www.valentbiosciences.com/agriculture/products/bounceback-foliar/",
		"https://www.valentbiosciences.com/agriculture/products/calron-soil/",
		"https://www.valentbiosciences.com/agriculture/products/cellmate-foliar-plus-mo/",
		"https://www.valentbiosciences.com/agriculture/products/chlorodrive-foliar/",
		"https://www.valentbiosciences.com/agriculture/products/copron/",
		"https://www.valentbiosciences.com/agriculture/products/dipel/",
		"https://www.valentbiosciences.com/agriculture/products/drivek",
		"https://www.valentbiosciences.com/agriculture/products/fbs-harmony/",
		"https://www.valentbiosciences.com/agriculture/products/fbs-humate-plus/",
		"https://www.valentbiosciences.com/agriculture/products/flexforce/",
		"https://www.valentbiosciences.com/agriculture/products/gnatrol/",
		"https://www.valentbiosciences.com/agriculture/products/greensurge-foliar/",
		"https://www.valentbiosciences.com/agriculture/products/ingrain/",
		"https://www.valentbiosciences.com/agriculture/products/k-surge/",
		"https://www.valentbiosciences.com/agriculture/products/leap/",
		"https://www.valentbiosciences.com/agriculture/products/manron/",
		"https://www.valentbiosciences.com/agriculture/products/maxcel/",
		"https://www.valentbiosciences.com/agriculture/products/microblend-foliar-zn-mn-b/",
		"https://www.valentbiosciences.com/agriculture/products/microblend-zn-mn-cu/",
		"https://www.valentbiosciences.com/agriculture/products/fbs-harmony/",
		"https://www.valentbiosciences.com/agriculture/products/fbs-humate-plus/",
		"https://www.valentbiosciences.com/agriculture/products/flexforce/",
		"https://www.valentbiosciences.com/agriculture/products/gnatrol/",
		"https://www.valentbiosciences.com/agriculture/products/greensurge-foliar/",
		"https://www.valentbiosciences.com/agriculture/products/ingrain/",
		"https://www.valentbiosciences.com/agriculture/products/k-surge/",
		"https://www.valentbiosciences.com/agriculture/products/leap/",
		"https://www.valentbiosciences.com/agriculture/products/manron/",
		"https://www.valentbiosciences.com/agriculture/products/maxcel/",
		"https://www.valentbiosciences.com/agriculture/products/microblend-foliar-zn-mn-b/",
		"https://www.valentbiosciences.com/agriculture/products/microblend-zn-mn-cu/",
		"https://www.valentbiosciences.com/agriculture/products/pincor/",
		"https://www.valentbiosciences.com/agriculture/products/pomaxa/",
		"https://www.valentbiosciences.com/agriculture/products/prestige/",
		"https://www.valentbiosciences.com/agriculture/products/progibb/",
		"https://www.valentbiosciences.com/agriculture/products/proliant/",
		"https://www.valentbiosciences.com/agriculture/products/proliant-plus/",
		"https://www.valentbiosciences.com/agriculture/products/promalin/",
		"https://www.valentbiosciences.com/agriculture/products/protone/",
		"https://www.valentbiosciences.com/agriculture/products/provide/",
		"https://www.valentbiosciences.com/agriculture/products/release/",
		"https://www.valentbiosciences.com/agriculture/products/retain/",
		"https://www.valentbiosciences.com/agriculture/products/ryzup-smartgrass/",
		"https://www.valentbiosciences.com/agriculture/products/ryzup/",
		"https://www.valentbiosciences.com/agriculture/products/slon-plus/",
		"https://www.valentbiosciences.com/agriculture/products/sugarmax/",
		"https://www.valentbiosciences.com/agriculture/products/supersede-soybean/",
		"https://www.valentbiosciences.com/agriculture/products/supersede-small-grains/",
		"https://www.valentbiosciences.com/agriculture/products/sympatico/",
		"https://www.valentbiosciences.com/agriculture/products/symvado/",
		"https://www.valentbiosciences.com/agriculture/products/transit-foliar/",
		"https://www.valentbiosciences.com/agriculture/products/transit-soil/",
		"https://www.valentbiosciences.com/agriculture/products/xentari/",
		"https://www.valentbiosciences.com/agriculture/products/zicron/",
		"https://www.valentbiosciences.com/foresthealth/products/foray/48b/",
		"https://www.valentbiosciences.com/foresthealth/products/mimic/",
		"https://www.valentbiosciences.com/foresthealth/foray-technical-manual/",
	} // URL to fetch HTML content from
	localFilePath := "valentbiosciences.html" // Path where HTML file will be stored

	var getData []string

	for _, urls := range remoteAPIURL {
		getData = append(getData, getDataFromURL(urls)) // If not, download HTML content from URL
	}
	appendAndWriteToFile(localFilePath, strings.Join(getData, "")) // Save downloaded content to file

	finalList := extractPDFUrls(strings.Join(getData, "")) // Extract all PDF links from HTML content

	outputDir := "PDFs/" // Directory to store downloaded PDFs

	if !directoryExists(outputDir) { // Check if directory exists
		createDirectory(outputDir, 0o755) // Create directory with read-write-execute permissions
	}

	// Remove duplicates from a given slice.
	finalList = removeDuplicatesFromSlice(finalList)

	// Loop through all extracted PDF URLs
	for _, urls := range finalList {
		if isUrlValid(urls) { // Check if the final URL is valid
			downloadPDF(urls, outputDir) // Download the PDF
		}
	}
}

// Opens a file in append mode, or creates it, and writes the content to it
func appendAndWriteToFile(path string, content string) {
	filePath, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Open file with specified flags and permissions
	if err != nil {
		log.Println(err) // Log error if opening fails
	}
	_, err = filePath.WriteString(content + "\n") // Write content to file
	if err != nil {
		log.Println(err) // Log error if writing fails
	}
	err = filePath.Close() // Close the file
	if err != nil {
		log.Println(err) // Log error if closing fails
	}
}

// Extracts filename from full path (e.g. "/dir/file.pdf" → "file.pdf")
func getFilename(path string) string {
	return filepath.Base(path) // Use Base function to get file name only
}

// Converts a raw URL into a sanitized PDF filename safe for filesystem
func urlToFilename(rawURL string) string {
	lower := strings.ToLower(rawURL) // Convert URL to lowercase
	lower = getFilename(lower)       // Extract filename from URL

	reNonAlnum := regexp.MustCompile(`[^a-z0-9]`)   // Regex to match non-alphanumeric characters
	safe := reNonAlnum.ReplaceAllString(lower, "_") // Replace non-alphanumeric with underscores

	safe = regexp.MustCompile(`_+`).ReplaceAllString(safe, "_") // Collapse multiple underscores into one
	safe = strings.Trim(safe, "_")                              // Trim leading and trailing underscores

	var invalidSubstrings = []string{
		"_pdf", // Substring to remove from filename
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safe = removeSubstring(safe, invalidPre)
	}

	if getFileExtension(safe) != ".pdf" { // Ensure file ends with .pdf
		safe = safe + ".pdf"
	}

	return safe // Return sanitized filename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Gets the file extension from a given file path
func getFileExtension(path string) string {
	return filepath.Ext(path) // Extract and return file extension
}

// Checks if a file exists at the specified path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {                // If error occurs, file doesn't exist
		return false
	}
	return !info.IsDir() // Return true if path is a file (not a directory)
}

// Downloads a PDF from given URL and saves it in the specified directory
func downloadPDF(finalURL, outputDir string) bool {
	filename := strings.ToLower(urlToFilename(finalURL)) // Sanitize the filename
	filePath := filepath.Join(outputDir, filename)       // Construct full path for output file

	if fileExists(filePath) { // Skip if file already exists
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	client := &http.Client{Timeout: 15 * time.Minute} // Create HTTP client with timeout

	// Create request with custom User-Agent
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		log.Printf("Failed to create request for %s: %v", finalURL, err)
		return false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close() // Ensure response body is closed

	if resp.StatusCode != http.StatusOK { // Check if response is 200 OK
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	contentType := resp.Header.Get("Content-Type")                                                                  // Get content type of response
	if !strings.Contains(contentType, "application/pdf") && !strings.Contains(contentType, "binary/octet-stream") { // Check if it's a PDF
		log.Printf("Invalid content type for %s: %s (expected application/pdf or binary/octet-stream)", finalURL, contentType)
		return false
	}

	var buf bytes.Buffer                     // Create a buffer to hold response data
	written, err := io.Copy(&buf, resp.Body) // Copy data into buffer
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 { // Skip empty files
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	out, err := os.Create(filePath) // Create output file
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close() // Ensure file is closed after writing

	if _, err := buf.WriteTo(out); err != nil { // Write buffer contents to file
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath) // Log success
	return true
}

// Checks whether a given directory exists
func directoryExists(path string) bool {
	directory, err := os.Stat(path) // Get info for the path
	if err != nil {
		return false // Return false if error occurs
	}
	return directory.IsDir() // Return true if it's a directory
}

// Creates a directory at given path with provided permissions
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission) // Attempt to create directory
	if err != nil {
		log.Println(err) // Log error if creation fails
	}
}

// Verifies whether a string is a valid URL format
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Try parsing the URL
	return err == nil                  // Return true if valid
}

// Removes duplicate strings from a slice
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool) // Map to track seen values
	var newReturnSlice []string    // Slice to store unique values
	for _, content := range slice {
		if !check[content] { // If not already seen
			check[content] = true                            // Mark as seen
			newReturnSlice = append(newReturnSlice, content) // Add to result
		}
	}
	return newReturnSlice
}

// Extracts all links to PDF files from given HTML string
func extractPDFUrls(htmlInput string) []string {
	var pdfLinks []string // Slice to hold found PDF links

	doc, err := html.Parse(strings.NewReader(htmlInput)) // Parse HTML content
	if err != nil {
		log.Println(err) // Log parse error
		return nil
	}

	var traverse func(*html.Node) // Recursive function to traverse HTML nodes
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" { // If it's an <a> tag
			for _, attr := range n.Attr {
				if attr.Key == "href" { // Look for href attribute
					href := strings.TrimSpace(attr.Val)                  // Get link
					if strings.Contains(strings.ToLower(href), ".pdf") { // If link points to a PDF
						pdfLinks = append(pdfLinks, href) // Add to list
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling { // Traverse children
			traverse(c)
		}
	}

	traverse(doc)   // Start traversal from root
	return pdfLinks // Return found PDF links
}

// Performs HTTP GET request with a custom User-Agent and returns response body as string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri) // Log which URL is being scraped

	// Create a new HTTP request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("Request creation failed:", err)
		return ""
	}

	// Set a User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36")

	// Use http.DefaultClient to send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Request failed:", err)
		return ""
	}
	defer response.Body.Close() // Always close response body

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ""
	}

	return string(body) // Return response body as string
}
