async function shortenUrl() {
    const url = document.getElementById('urlInput').value.trim();
    const resultDiv = document.getElementById('result');

    if (!url) {
        resultDiv.innerHTML = '<span style="color: red;">Please enter a URL!</span>';
        return;
    }

    // Show loading state
    resultDiv.innerHTML = 'Shortening...';

    try {
        // Call the Go backend API
        // Because we are on the same port, we can just use a relative path!
        const response = await fetch(`/shorten?url=${encodeURIComponent(url)}`);
        
        if (!response.ok) {
            throw new Error('Server returned an error');
        }

        const shortUrl = await response.text();
        
        // Display the result as a clickable link
        resultDiv.innerHTML = `Short URL: <a href="${shortUrl}" target="_blank">${shortUrl}</a>`;
        
    } catch (error) {
        resultDiv.innerHTML = `<span style="color: red;">Error: ${error.message}</span>`;
    }
}