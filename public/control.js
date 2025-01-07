const pathname = window.location.pathname;
if (["/control"].includes(pathname)) {
  console.log("control or home page");
  const errorMessage = document.getElementById("error-message");
  document.addEventListener("htmx:afterRequest", (e) => {
    if (e.detail.successful) {
      errorMessage.innerText = "";
    } else {
      errorMessage.innerText = e.detail.xhr.responseText;
    }
  });
}

const setError = (errorMessage) => {
  document.getElementById("error").innerText = errorMessage;
};

const getPassword = () => {
  return document.getElementById("password").value;
};

const getSelectedPlayer = () => {
  return document.getElementById("playerlist-dropdown").value;
};

const fetchAuthEndpoint = (endpoint, method) => {
  const password = getPassword();
  if (password === "") {
    setError("password is required");
    return;
  }
  const pw = (endpoint.includes("?") ? "&" : "?") + `password=${password}`;
  fetch(`/auth/${endpoint}${pw}`, {
    method: method,
  })
    .then((res) => res.text())
    .then((msg) => {
      if (msg != "success") {
        setError(msg);
      } else {
        setError("");
        document.getElementById("playerlist-dropdown").value = "";
        document.getElementById("amount").value = 0;
      }
    });
};
