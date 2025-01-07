var selectedPlayer = document.getElementById("playerlist-dropdown").value;

document.addEventListener("htmx:wsBeforeMessage", (e) => {
  if (
    !e.detail.elt["htmx-internal-data"].webSocket.socket.url.endsWith(
      "/players-ws",
    )
  ) {
    return;
  }
  selectedPlayer = document.getElementById("playerlist-dropdown").value;
});

document.addEventListener("htmx:wsAfterMessage", (e) => {
  if (
    !e.detail.elt["htmx-internal-data"].webSocket.socket.url.endsWith(
      "/players-ws",
    )
  ) {
    return;
  }

  // check if stored option is valid
  const options = document.getElementById("playerlist-dropdown").options;
  for (let i = 0; i < options.length; i++) {
    if (options[i].value === selectedPlayer) {
      document.getElementById("playerlist-dropdown").value = selectedPlayer;
      return;
    }
  }
  document.getElementById("playerlist-dropdown").value = "";
});
