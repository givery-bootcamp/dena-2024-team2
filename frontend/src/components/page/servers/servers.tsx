import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { ServerList } from "~/components/model/server";
import { Card, Text } from "~/components/ui";
import { getServers } from "~/domains/servers";
import { NewServerDialog } from "./components";
import { useServer } from "./hooks";
import styles from "./servers.module.scss";

export const ServersPage = () => {
	const { data: servers } = useQuery({
		queryKey: ["servers"],
		queryFn: getServers,
	});

	const [open, setOpen] = useState(false);

	const handleClickPlus = () => {
		setOpen(true);
	};
	const handleClose = () => {
		setOpen(false);
	};

	const { handleChangeName, handleChangeIcon, handleCreateSever } = useServer();

	return (
		<div className={styles.root}>
			<div className={styles.card}>
				<Card
					header={
						<Text as="h2" size="lg" bold>
							🪐 Servers
						</Text>
					}
					body={
						<ServerList servers={servers ?? []} onClickPlus={handleClickPlus} />
					}
				/>
			</div>
			<NewServerDialog
				open={open}
				onClose={handleClose}
				onSubmit={handleCreateSever}
				onChangeName={handleChangeName}
				onChangeIcon={handleChangeIcon}
			/>
		</div>
	);
};
