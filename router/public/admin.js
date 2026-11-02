// clear the selected player and amount after a successful admin action.
function resetAdminForm(e) {
  const { elt, successful, xhr } = e.detail;
  if (!successful || xhr.responseText !== "") return;
  if (!elt.closest("#admin-form")) return;

  const form = document.getElementById("admin-form");
  form.querySelector("#amount").value = "";
  form.querySelector("#playerlist-dropdown").value = "";
}

if (document.getElementById("admin-form")) {
  document.addEventListener("htmx:afterRequest", resetAdminForm);
}
