console.log("ping-pong.js loaded");
if (
  ["/game", "/control", "/leaderboard", "/buzzed-in"].includes(
    window.location.pathname,
  ) ||
  window.location.pathname.startsWith("/play/")
) {
  const PONG_DATA = JSON.stringify({
    name: "\x1F",
    HEADERS: {
      "HX-Current-URL": window.location.href,
    },
  });

  function pongFunc(ws, i = 0) {
    if (i === 0) console.log("starting pong");

    if (ws.readyState !== WebSocket.OPEN) {
      console.log("ws not open, stopping pong");
      return;
    }

    ws.send(PONG_DATA);

    if (i === 0) console.log("sent first pong");

    setTimeout(() => {
      pongFunc(ws, i + 1);
    }, 5000);
  }

  document.addEventListener("htmx:wsOpen", (e) => {
    const ws = e.detail.event.currentTarget;
    try {
      pongFunc(ws);
    } catch (e) {
      console.error(e);
    }
  });
}
