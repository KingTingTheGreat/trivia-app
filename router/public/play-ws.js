// this ws logic is only needed on the play/buzzer page

function handleWSCloseAndError() {
  document.getElementById("player-data").innerHTML =
    `<div class="m-8">${loading}</div>`;
  new Promise((resolve) => setTimeout(resolve, 500)).then(() => {
    // console.log("reloading page after ws close/error");
    // const name = window.location.pathname.split("/").pop();
    // fetch(`/player?name=${name}`).then((res) => {
    //   if (res.status === 200) {
    //     window.location.reload();
    //   } else {
    window.location.href = "/";
    // }
    // });
  });
}

if (window.location.pathname.startsWith("/play")) {
  document.addEventListener("htmx:wsClose", () => {
    console.log("ws close event");
    handleWSCloseAndError();
  });

  // document.addEventListener("htmx:wsError", () => {
  //   console.log("ws error event");
  //   handleWSCloseAndError();
  // });

  document.addEventListener("htmx:wsBeforeMessage", () => {
    console.log("buzzing in");
  });

  document.addEventListener("htmx:wsAfterMessage", (e) => {
    console.log("new buzzer", e.detail.message);
  });
}
