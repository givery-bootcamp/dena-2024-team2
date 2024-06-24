import { createLazyFileRoute } from "@tanstack/react-router";
import { SpaceDetailPage } from "~/components/page";

export const Route = createLazyFileRoute("/spaces/$spaceId")({
	component: () => <SpaceDetailPage />,
});
