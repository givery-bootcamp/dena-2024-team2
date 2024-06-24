import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/spaces/$spaceId")({
	component: () => <div>Hello /spaces$spaceId!</div>,
});
