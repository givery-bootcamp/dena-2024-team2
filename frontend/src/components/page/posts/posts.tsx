import { useQuery } from "@tanstack/react-query";
import { useParams } from "@tanstack/react-router";
import { getChannels } from "~/domains/channels";
import { getPosts } from "~/domains/posts";
import { ChannelsPanel, PostForm, PostsPanel } from "./components";
import { usePost } from "./hooks/use-post";
import styles from "./posts.module.scss";

export const Posts = () => {
	const { serverId, channelId } = useParams({
		strict: false,
	});

	const { status, handleChangeContent, handleClickSubmit } = usePost(
		Number(serverId),
		Number(channelId),
	);

	const { data: channels } = useQuery({
		queryKey: ["channels"],
		queryFn: getChannels(Number(serverId)),
	});

	const { data: posts } = useQuery({
		queryKey: ["posts"],
		queryFn: getPosts(Number(channelId)),
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
				<PostsPanel
					channelName={
						channels?.find((c) => c.id === Number(channelId))?.name ?? ""
					}
					posts={posts ?? []}
				/>
			</div>
			<div className={styles.form}>
				<PostForm
					onChange={handleChangeContent}
					onSubmit={handleClickSubmit}
					disableSubmit={status === "pending"}
				/>
			</div>
		</div>
	);
};
