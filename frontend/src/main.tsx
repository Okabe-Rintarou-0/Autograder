import { enableMapSet } from 'immer';
import ReactDOM from "react-dom/client";
import App from "./App";
import { loadToken } from "./service/token";

function init() {
  loadToken();
  enableMapSet();
}

init();

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <App />
);