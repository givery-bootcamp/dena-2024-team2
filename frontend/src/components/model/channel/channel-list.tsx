import clsx from "clsx";
import { IconButton, Text } from "~/components/ui";
import type { Channel } from "~/domains/channels";
import styles from "./channel-list.module.scss";

type Props = {
	channels: Channel[];
	currentChannelId?: number;
	onClickPlus?: () => void;
};

export const ChannelList = ({
	channels,
	currentChannelId,
	onClickPlus,
}: Props) => {
	// NOTE: serverと同じくここにこの関数作るのちょっと違和感あるけど一旦このままで
	const urlToChannel = (channel: Channel) =>
		`/servers/${channel.serverId}/channels/${channel.id}`;

	const handleClickPlus = () => {
		onClickPlus?.();
	};

	return (
		<ul className={styles.root}>
			{channels.map((channel) => (
				<a
					key={channel.id}
					href={urlToChannel(channel)}
					className={styles.link}
				>
					<li
						className={clsx(
							styles.item,
							currentChannelId === channel.id && styles.item__current,
						)}
					>
						<Text as="span" size="md">
							# {channel.name}
						</Text>
					</li>
				</a>
			))}
			{onClickPlus && (
				<li>
					<IconButton
						size="sm"
						color="inherit"
						icon="plus"
						onClick={handleClickPlus}
					/>
				</li>
			)}
		</ul>
	);
};
