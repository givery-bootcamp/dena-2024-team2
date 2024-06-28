import { ChannelsPanel, PostForm, PostsPanel } from "./components";
import { mock } from "./mock";
import styles from "./posts.module.scss";

export const Posts = () => {
	return (
		<div className={styles.root}>
			<div className={styles.channels}>
				<ChannelsPanel channels={mock.channels} />
			</div>
			<div className={styles.posts}>
				<PostsPanel posts={mock.posts} />
			</div>
			<div className={styles.form}>
				<PostForm onChange={() => {}} onSubmit={() => {}} />
			</div>
		</div>
	);
};
