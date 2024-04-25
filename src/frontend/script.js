document.getElementById('wiki-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent default form submission

    // Get form data
    const formData = new FormData(this);

    // Convert form data to JSON object
    const jsonData = {};
    formData.forEach((value, key) => {
        jsonData[key] = value;
    });

    // Send POST request to Go server
    fetch('http://localhost:8080/data', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonData)
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Failed to send data to server');
        }
    })
    .then(data => {
        // Handle response from server
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});