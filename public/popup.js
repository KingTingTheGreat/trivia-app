if (["/control", "/buzzed-in", "/game"].includes(window.location.pathname)) {
  showPopup();
  document.addEventListener("click", hidePopup);
}

function showPopup() {
  const popup = document.getElementById("popup");
  const overlay = document.getElementById("overlay");
  popup.classList.remove("hidden");
  overlay.classList.remove("hidden");
}

function hidePopup() {
  const popup = document.getElementById("popup");
  const overlay = document.getElementById("overlay");

  popup.classList.add("hidden");
  overlay.classList.add("hidden");
}
