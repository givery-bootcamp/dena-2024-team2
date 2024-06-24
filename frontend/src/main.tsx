import { RouterProvider, createRouter } from "@tanstack/react-router";
import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import "./global.scss";
import { routeTree } from "./routeTree.gen";

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
	interface Register {
		router: typeof router;
	}
}

const rootElement = document.getElementById("root");
if (!rootElement) throw new Error("Failed to find the root Element");
const root = ReactDOM.createRoot(rootElement);
root.render(
	<StrictMode>
		<RouterProvider router={router} />
	</StrictMode>,
);
