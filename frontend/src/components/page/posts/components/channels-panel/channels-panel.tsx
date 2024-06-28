import { ChannelList } from "~/components/model/channel";
import { Card, Text } from "~/components/ui";
import type { Channel } from "~/domains/channels";

type Props = {
	channels: Channel[];
};

export const ChannelsPanel = ({ channels }: Props) => {
	return (
		<Card
			header={
				<Text as="h2" size="md" bold>
					Channels
				</Text>
			}
			body={<ChannelList channels={channels} />}
		/>
	);
};
