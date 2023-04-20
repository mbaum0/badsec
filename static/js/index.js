const body = document.querySelector("body");
const container = document.querySelector(".container");
const fetchButton = document.getElementById("fetch-button");
const setButton = document.getElementById("set-button");
const keyInput = document.getElementById("key");
const contentInput = document.getElementById("content");
const darkModeToggle = document.getElementById("dark-mode-toggle");

fetchButton.addEventListener("click", () => {
    const key = keyInput.value;
    const requestBody = {
        key: key
    };
    fetch(`/api/secret`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
    })
        .then(response => response.json())
        .then(body => {
            contentInput.value = `${body.content}`;
        })
        .catch(error => {
            console.error(error);
            contentInput.value = "An error occurred while fetching the secret.";
        });
});

setButton.addEventListener("click", () => {
    const key = keyInput.value;
    const content = contentInput.value;
    const requestBody = {
        key: key,
        content: content
    };
    fetch(`/api/secret`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
    })
        .then(response => response.json())
        .then(body => {
            contentInput.value = `${body.content}`;
        })
        .catch(error => {
            console.error(error);
            contentInput.value = "An error occurred while setting the secret.";
        });
});

darkModeToggle.addEventListener("change", () => {
    if (darkModeToggle.checked) {
        body.classList.add("dark-mode")
    } else {
        body.classList.remove("dark-mode")
    }
});