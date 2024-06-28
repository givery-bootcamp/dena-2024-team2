import { joinPaths } from "@tanstack/react-router";
import { IconButton, IconImage, Text } from "~/components/ui";
import type { Server } from "~/domains/servers";
import styles from "./server-list.module.scss";

type Props = {
	servers: Server[];
	onClickPlus?: () => void;
};

// TODO: 余裕があればコンテキストメニューを追加する
export const ServerList = ({ servers, onClickPlus }: Props) => {
	// NOTE: ここにこの関数作るのちょっと違和感あるけど一旦このままで
	// サーバーへのリンクを生成する
	const urlToServer = (server: Server) =>
		// NOTE: ホントは最後に開いてたチャンネルに飛ぶべきだけどちょっとめんどくさいので0にしてる
		joinPaths(["/servers", server.id, "channels", "0"]);

	const handleClickPlus = () => {
		onClickPlus?.();
	};
	return (
		<ul className={styles.root}>
			{servers.map((server) => (
				<a key={server.id} href={urlToServer(server)} className={styles.link}>
					<li className={styles.item}>
						<IconImage src={server.icon} size="sm" />
						<Text as="span" size="md">
							{server.name}
						</Text>
					</li>
				</a>
			))}
			{onClickPlus && (
				<li>
					<IconButton
						color="primary"
						icon="plus"
						size="lg"
						onClick={handleClickPlus}
					/>
				</li>
			)}
		</ul>
	);
};
