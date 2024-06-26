import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { Header } from "~/components/ui/header";
import "modern-css-reset";
import { QueryClientProvider } from "~/components/functional/query-client-provider";
import styles from "./__root.module.scss";

export const Route = createRootRoute({
	component: () => (
		<QueryClientProvider>
			<div className={styles.root}>
				<Header />
				<Outlet />
				<TanStackRouterDevtools />
			</div>
		</QueryClientProvider>
	),
});
