import { useQuery } from "@tanstack/react-query";
import { ServerList } from "~/components/model/server";
import { Card, Text } from "~/components/ui";
import { getServers } from "~/domains/servers";
import styles from "./servers.module.scss";

export const ServersPage = () => {
	const { data: servers } = useQuery({
		queryKey: ["servers"],
		queryFn: getServers,
	});

	return (
		<div className={styles.root}>
			<div className={styles.card}>
				<Card
					header={
						<Text as="h2" size="lg" bold>
							🪐 Servers
						</Text>
					}
					body={<ServerList servers={servers ?? []} />}
				/>
			</div>
		</div>
	);
};
