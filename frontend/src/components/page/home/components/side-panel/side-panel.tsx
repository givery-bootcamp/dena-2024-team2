import { ChannelList } from "~/components/model/channel";
import { ServerList } from "~/components/model/server";
import type { Channel } from "~/domains/channels";
import type { Server } from "~/domains/servers";

type Props = {
	servers: Server[];
	channels: Channel[];
};

export const SidePanel = ({ servers, channels }: Props) => {
	return (
		<div>
			<ServerList servers={servers} />
			<ChannelList channels={channels} />
		</div>
	);
};
