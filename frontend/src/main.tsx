import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./global.scss";

const rootElement = document.getElementById("root");
if (!rootElement) throw new Error("Failed to find the root Element");
const root = ReactDOM.createRoot(rootElement);
root.render(
	<React.StrictMode>
		<App />
	</React.StrictMode>,
);
