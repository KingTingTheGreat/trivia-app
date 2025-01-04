document.body.addEventListener("htmx:afterRequest", function (e) {
  const status = e.detail.xhr.status;
  console.log(status);
  if (status >= 300 && status < 400) {
    console.log("htmx redirect");
    let location = e.detail.xhr.getResponseHeader("Location");
    if (location) {
      window.location.replace(location);
    }
  }
});
