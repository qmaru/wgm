import React from "react"
import ReactDOM from "react-dom/client"

import "@fontsource-variable/google-sans-code"
import "@fontsource-variable/noto-sans-sc"

import "@/index.css"
import App from "@/App"

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement)
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
