document.querySelector("form").addEventListener("submit", function(event) {
    // Prevent default form submission
    event.preventDefault();

    // Clear previous results
    const result = document.querySelector(".result");
    while (result.hasChildNodes()) {
        result.removeChild(result.lastChild);
    }

    // Get form data
    const formData = new FormData(this);

    // Convert form data to JSON object
    const jsonData = {};
    formData.forEach((value, key) => {
        jsonData[key] = value;
    });

    // Send POST request to Go server
    fetch("http://localhost:8080/data", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Failed to send data to server");
        }
    })
    .then(data => {
        // Handle response from server
        displayResults(data);
    })
    .catch(error => {
        console.error("Error:", error);
    });
});


document.querySelector(".swap").addEventListener("click", () => {
    const startPage = document.getElementById("start-page");
    const endPage = document.getElementById("end-page");
    [startPage.value, endPage.value] = [endPage.value, startPage.value];
});


const displayResults = (data) => {
    const result = document.querySelector(".result");
    
    const jumlahArtikelDiperiksa = document.createElement("p");
    jumlahArtikelDiperiksa.textContent = `Jumlah Artikel Diperiksa: ${data.jumlahArtikelDiperiksa}`
    result.appendChild(jumlahArtikelDiperiksa);
    
    const jumlahArtikelDilalui = document.createElement("p");
    jumlahArtikelDilalui.textContent = `Jumlah Artikel Dilalui: ${data.jumlahArtikelDilalui}`;
    result.appendChild(jumlahArtikelDilalui);

    const searchDuration = document.createElement("p");
    searchDuration.textContent = `Search Duration: ${data.searchDuration} ms`;
    result.appendChild(searchDuration);

    data.routes.forEach((rute, index) => {
        const div = document.createElement("div");
        
        const p = document.createElement("p");
        p.textContent = `Route ${index + 1}:`;
        div.appendChild(p);
        
        const ol = document.createElement("ol");
        rute.forEach(link => {
            const li = document.createElement("li");
            const a = document.createElement("a");
            a.href = link;
            a.textContent = link;
            li.appendChild(a);
            ol.appendChild(li);
        });
        div.appendChild(ol);
        result.appendChild(div)
    })
}
