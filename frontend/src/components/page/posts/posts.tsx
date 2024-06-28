import { useQuery } from "@tanstack/react-query";
import { getChannels } from "~/domains/channels";
import { ChannelsPanel, PostForm, PostsPanel } from "./components";
import { mock } from "./mock";
import styles from "./posts.module.scss";

export const Posts = () => {
	const { data: channels } = useQuery({
		queryKey: ["channels"],
		queryFn: getChannels(1),
	});

	return (
		<div className={styles.root}>
			<div className={styles.channels}>
				<ChannelsPanel channels={channels ?? []} currentChannelId={1} />
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
