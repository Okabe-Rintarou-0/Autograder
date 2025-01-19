import ReactDOM from "react-dom/client";
import App from "./App";
import React from "react";
import { loadToken } from "./service/token";

loadToken();

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <App />
);