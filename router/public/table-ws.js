if (
  ["/admin", "/leaderboard", "/buzzed-in", "/game"].includes(
    window.location.pathname,
  )
) {
  htmx.config.wsReconnectDelay = function () {
    return 500; // return value in milliseconds
  };

  const setLoading = () => {
    console.log("setLoading()");
    document.getElementById("leaderboard-table").innerHTML = loading;
    document.getElementById("buzzed-in-table").innerHTML = loading;
  };

  document.addEventListener("htmx:wsConnecting", (e) => {
    console.log("starting ws connection");
  });

  document.addEventListener("htmx:wsAfterMessage", (e) => {
    // console.log("new message from websocket", e.detail.message);
  });

  document.addEventListener("htmx:wsClose", (e) => {
    setLoading();
  });

  document.addEventListener("htmx:wsError", (e) => {
    setLoading();
  });
}
