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
  document.getElementById("error-message").innerText = errorMessage;
};

const getPassword = () => {
  return document.getElementById("password").value;
};

const getSelectedPlayer = () => {
  return document.getElementById("playerlist-dropdown").value;
};

const getAmount = () => {
  return document.getElementById("amount").value;
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

const updatePlayer = () => {
  const selectedPlayer = getSelectedPlayer();
  const amount = getAmount();
  fetchAuthEndpoint(
    `player?name=${selectedPlayer}&amount=${amount}&action=update`,
    "PUT",
  );
};

const clearPlayer = () => {
  const selectedPlayer = getSelectedPlayer();
  fetchAuthEndpoint(`player?name=${selectedPlayer}&action=clear`, "PUT");
};

const removePlayer = () => {
  const selectedPlayer = getSelectedPlayer();
  if (
    confirm(
      "Are you sure you want to delete this player? This action is permanent.",
    )
  ) {
    fetchAuthEndpoint(`player?name=${selectedPlayer}`, "DELETE");
  }
};

const resetBuzzers = () => {
  fetchAuthEndpoint("reset-buzzers", "POST");
};

const resetGame = () => {
  if (
    confirm(
      "Are you sure you want to reset the game? This action is permanent.",
    )
  ) {
    fetchAuthEndpoint("reset-game", "POST");
  }
};
