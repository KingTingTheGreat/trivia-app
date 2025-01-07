document.body.addEventListener("htmx:afterRequest", function (e) {
  const status = e.detail.xhr.status;
  const msg = e.detail.xhr.message;
  console.log(e, status, msg);
  if (status >= 300 && status < 400) {
    console.log("htmx redirect");
    let location = e.detail.xhr.getResponseHeader("Location");
    if (location) {
      window.location.replace(location);
    }
  }
});
