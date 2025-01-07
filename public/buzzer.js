const audio = new Audio("/public/buzzer.mp3"); // Path to the sound file
document.addEventListener("htmx:wsAfterMessage", (e) => {
  if (
    !e.detail.elt["htmx-internal-data"].webSocket.socket.url.endsWith(
      "/buzzed-in-ws",
    )
  ) {
    return;
  }

  const buzzedIn = document.getElementById("buzzed-in-body");
  if (buzzedIn && buzzedIn.rows?.length > 0) {
    console.log("playing buzzer audio");
    audio.play();
  }
});
