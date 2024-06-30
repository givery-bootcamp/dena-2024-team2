import { ChannelList } from "~/components/model/channel";
import { Card, Text } from "~/components/ui";
import type { Channel } from "~/domains/channels";

type Props = {
	channels: Channel[];
	currentChannelId: number;
	onClickPlus?: () => void;
};

export const ChannelsPanel = ({
	channels,
	currentChannelId,
	onClickPlus,
}: Props) => {
	const handleClickPlus = () => {
		onClickPlus?.();
	};

	return (
		<Card
			header={
				<Text as="h2" size="md" bold>
					📺 Channels
				</Text>
			}
			body={
				<ChannelList
					channels={channels}
					currentChannelId={currentChannelId}
					onClickPlus={handleClickPlus}
				/>
			}
		/>
	);
};
