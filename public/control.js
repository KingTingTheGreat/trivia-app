if (window.location.pathname == "/control") {
  document.addEventListener("htmx:afterRequest", (e) => {
    if (e.detail.xhr.responseText.length === 0) {
      document.getElementById("amount").value = "";
      document.getElementById("playerlist-dropdown").value = "";
    }
  });
}
