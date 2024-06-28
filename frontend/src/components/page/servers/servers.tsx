import { ServerList } from "~/components/model/server";
import { Card, Text } from "~/components/ui";
import { mock } from "./mock";
import styles from "./servers.module.scss";

export const ServersPage = () => {
	return (
		<div className={styles.root}>
			<div className={styles.card}>
				<Card
					header={
						<Text as="h2" size="lg" bold>
							🪐 Servers
						</Text>
					}
					body={<ServerList servers={mock.servers} />}
				/>
			</div>
		</div>
	);
};
