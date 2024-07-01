import { useQuery } from "@tanstack/react-query";
import { useParams } from "@tanstack/react-router";
import { useState } from "react";
import { getChannels } from "~/domains/channels";
import { getPosts } from "~/domains/posts";
import {
	ChannelsPanel,
	NewChannelDialog,
	PostForm,
	PostsPanel,
} from "./components";
import { useChannel, usePost } from "./hooks";
import styles from "./posts.module.scss";

export const Posts = () => {
	const { serverId, channelId } = useParams({
		strict: false,
	});

	const { status, handleChangeContent, handleClickSubmit } = usePost(
		Number(serverId),
		Number(channelId),
	);

	const { handleChangeName, handleCreateChannel } = useChannel(
		Number(serverId),
	);

	const [open, setOpen] = useState(false);
	const handleCloseDialog = () => {
		setOpen(false);
	};
	const handleClickPlus = () => {
		setOpen(true);
	};

	const { data: channels } = useQuery({
		queryKey: ["channels"],
		queryFn: getChannels(Number(serverId)),
	});

	const { data: posts } = useQuery({
		queryKey: ["posts"],
		queryFn: getPosts(Number(channelId)),
		refetchInterval: 1000 * 10,
	});

	return (
		<div className={styles.root}>
			<div className={styles.channels}>
				<ChannelsPanel
					channels={channels ?? []}
					currentChannelId={Number(channelId)}
					onClickPlus={handleClickPlus}
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
			<NewChannelDialog
				open={open}
				onChangeName={handleChangeName}
				onSubmit={handleCreateChannel}
				onClose={handleCloseDialog}
			/>
		</div>
	);
};
