import { useQuery } from "@tanstack/react-query";
import { useParams } from "@tanstack/react-router";
import { getChannels } from "~/domains/channels";
import { ChannelsPanel, PostForm, PostsPanel } from "./components";
import { mock } from "./mock";
import styles from "./posts.module.scss";

export const Posts = () => {
	const { serverId, channelId } = useParams({
		strict: false,
	});

	const { data: channels } = useQuery({
		queryKey: ["channels"],
		queryFn: getChannels(Number(serverId)),
	});

	return (
		<div className={styles.root}>
			<div className={styles.channels}>
				<ChannelsPanel
					channels={channels ?? []}
					currentChannelId={Number(channelId)}
				/>
			</div>
			<div className={styles.posts}>
				<PostsPanel channelName={mock.channels[0].name} posts={mock.posts} />
			</div>
			<div className={styles.form}>
				<PostForm onChange={() => {}} onSubmit={() => {}} />
			</div>
		</div>
	);
};
