import { createLazyFileRoute } from "@tanstack/react-router";
import { ServersPage } from "~/components/page";

export const Route = createLazyFileRoute("/servers")({
	component: () => <ServersPage />,
});
