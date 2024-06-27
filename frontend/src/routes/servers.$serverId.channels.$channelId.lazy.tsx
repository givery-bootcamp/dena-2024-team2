import { createLazyFileRoute } from "@tanstack/react-router";
import { Home } from "~/components/page/home/home";

export const Route = createLazyFileRoute(
	"/servers/$serverId/channels/$channelId",
)({
	component: () => <Home />,
});
