const buzzedIn = document.getElementById("buzzed-in");

const audio = new Audio("/public/buzzer.mp3"); // Path to the sound file
document.addEventListener("htmx:wsAfterMessage", (e) => {
  if (!e.detail.message.includes('<tbody id="buzzed-in">')) {
    return;
  }
  if (buzzedIn && buzzedIn.rows?.length > 0) {
    console.log("playing buzzer audio");
    audio.play();
  }
});
