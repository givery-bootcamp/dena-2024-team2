import { Text } from "~/components/ui";
import type { Post } from "~/domains/posts";
import styles from "./post-card.module.scss";

type Props = {
	post: Post;
};

export const PostCard = ({ post }: Props) => {
	return (
		<div className={styles.root}>
			<div className={styles.header}>
				<span className={styles.header__name}>
					<Text as="span" size="md" bold>
						{post.user.name}
					</Text>
				</span>
				<Text as="span" size="sm">
					{post.createdAt}
				</Text>
			</div>
			<div className={styles.body}>{post.content}</div>
		</div>
	);
};
