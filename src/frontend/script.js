document.querySelector("form").addEventListener("submit", function(event) {
    // Prevent default form submission
    event.preventDefault();

    // Clear previous results
    const result = document.querySelector(".result");
    while (result.hasChildNodes()) {
        result.removeChild(result.lastChild);
    }

    // Check if start page and end page are the same
    const startPageValue = document.getElementById("start-page").value;
    const endPageValue = document.getElementById("end-page").value;
    if (startPageValue === endPageValue) {
        const div = document.createElement("div");
        div.classList.add("summary");
        const p = document.createElement("p");
        p.textContent = "Start page and end page must be different!";
        div.appendChild(p);
        result.appendChild(div);
        return;
    }

    // Display searching message
    const div = document.createElement("div");
    div.classList.add("summary");
    const p = document.createElement("p");
    p.textContent = "Searching...";
    div.appendChild(p);
    result.appendChild(div);

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
        console.log(data);
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

    // remove all result messages
    while (result.hasChildNodes()) {
        result.removeChild(result.lastChild);
    }

    const divSummary = document.createElement("div");
    divSummary.classList.add("summary");
    
    const jumlahArtikelDiperiksa = document.createElement("p");
    jumlahArtikelDiperiksa.textContent = `Jumlah Artikel Diperiksa: ${data.jumlahArtikelDiperiksa}`
    divSummary.appendChild(jumlahArtikelDiperiksa);
    
    const jumlahArtikelDilalui = document.createElement("p");
    jumlahArtikelDilalui.textContent = `Jumlah Artikel Dilalui: ${data.jumlahArtikelDilalui}`;
    divSummary.appendChild(jumlahArtikelDilalui);

    const searchDuration = document.createElement("p");
    searchDuration.textContent = `Search Duration: ${data.searchDuration} ms`;
    divSummary.appendChild(searchDuration);
    result.appendChild(divSummary);

    const divRoutes = document.createElement("div");
    divRoutes.classList.add("routes");


    data.routes.forEach((routes, index) => {
        const divRoute = document.createElement("div");
        divRoute.classList.add("route");
        
        const h3 = document.createElement("h3");
        h3.textContent = `Route ${index + 1}:`;
        divRoute.appendChild(h3);
        
        const ol = document.createElement("ol");
        for (const [link, title] of Object.entries(routes)) {
            const li = document.createElement("li");
            const a = document.createElement("a");
            a.setAttribute("target", "_blank");
            a.href = link;
            a.textContent = title;
            li.appendChild(a);
            ol.appendChild(li);
        };
        
        divRoute.appendChild(ol);
        divRoutes.appendChild(divRoute)
    });

    result.appendChild(divRoutes);
}
