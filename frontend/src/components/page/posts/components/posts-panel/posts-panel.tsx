import { PostList } from "~/components/model/posts";
import { Card, Text } from "~/components/ui";
import type { Post } from "~/domains/posts";
type Props = {
	posts: Post[];
};

export const PostsPanel = ({ posts }: Props) => {
	return (
		<Card
			header={
				<Text as="h2" size="md" bold>
					# channel1
				</Text>
			}
			body={<PostList posts={posts} />}
		/>
	);
};
