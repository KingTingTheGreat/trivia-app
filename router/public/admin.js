if (window.location.pathname == "/admin") {
  document.addEventListener("htmx:afterRequest", (e) => {
    if (e.detail.xhr.responseText.length === 0) {
      document.getElementById("amount").value = "";
      document.getElementById("playerlist-dropdown").value = "";
    }
  });
}
