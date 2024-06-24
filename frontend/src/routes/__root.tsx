import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { Header } from "~/components/ui/header";
import "modern-css-reset";
import styles from "./__root.module.scss";

export const Route = createRootRoute({
	component: () => (
		<div className={styles.root}>
			<Header />
			<Outlet />
			<TanStackRouterDevtools />
		</div>
	),
});
