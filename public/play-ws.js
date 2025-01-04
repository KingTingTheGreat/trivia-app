// this ws logic is only needed on the play/buzzer page
if (window.location.pathname.startsWith("/play")) {
  const PONG_DATA = JSON.stringify({
    name: "\x1F",
    HEADERS: {
      "HX-Current-URL": window.location.href,
    },
  });

  // const pongForm = document.getElementById("pong-form");
  function pongFunc(ws) {
    console.log("sending pong");
    ws.send(PONG_DATA);
    setTimeout(() => {
      console.log("recurse");
      pongFunc(ws);
    }, 5000);
  }

  document.addEventListener("htmx:wsOpen", (e) => {
    const ws = e.detail.event.currentTarget;
    try {
      ws.send(PONG_DATA);
    } catch (e) {
      console.error(e);
    }

    // console.log(e);
    // console.log(e.detail.socketWrapper);
    // console.log(Object.keys(e));
    // console.log(Object.getOwnPropertyNames(e));

    console.log("initializing connection");
    console.log("initiazed");

    pongFunc(ws);
  });

  document.addEventListener("htmx:wsClose", () => {
    window.location.replace("/");
    console.log("ws close event");
  });

  document.addEventListener("htmx:wsError", () => {
    window.location.replace("/");
    console.log("ws error event");
  });

  document.addEventListener("htmx:wsBeforeMessage", () => {
    console.log("buzzing in");
  });

  document.addEventListener("htmx:wsAfterMessage", (e) => {
    console.log("new buzzer", e.detail.message);
  });
}
