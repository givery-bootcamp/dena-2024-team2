import { createFileRoute } from "@tanstack/react-router";
import { LoginPage } from "~/components/page";

export const Route = createFileRoute("/login")({
	component: () => <LoginPage />,
});
