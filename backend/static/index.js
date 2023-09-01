const shortenBtn = document.getElementById("shortenBtn");
const urlInput = document.getElementById("url");
const shortenedLinkSection = document.getElementById("shortenedLinkSection");
const shortenedLink = document.getElementById("shortenedLink");

shortenBtn.addEventListener("click", () => {
  const url = urlInput.value;
  if (url.trim() !== "") {
    // Send request to backend API to shorten the URL
    fetch("/shorten", {
      method: "POST",
      body: JSON.stringify({ longUrl: url }),
      headers: { "Content-Type": "application/json" },
    })
      .then((response) => response.json())
      .then((data) => {
        shortenedLink.textContent = data.newUrl;
        shortenedLink.href = data.newUrl;
        shortenedLinkSection.classList.remove("hidden");
      });
  } else {
    alert("Please enter a valid URL.");
  }
});
