import type { Post } from "~/domains/posts";
import { PostCard } from "./post-card";
import styles from "./post-list.module.scss";

type Props = {
	posts: Post[];
};

export const PostList = ({ posts }: Props) => {
	return (
		<ul className={styles.list}>
			{posts.map((post) => (
				<li key={post.id}>
					<PostCard post={post} />
				</li>
			))}
		</ul>
	);
};
