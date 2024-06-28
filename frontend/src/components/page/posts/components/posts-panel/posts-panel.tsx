import { PostList } from "~/components/model/posts";
import { Card, Text } from "~/components/ui";
import type { Post } from "~/domains/posts";
type Props = {
	posts: Post[];
	channelName: string;
};

export const PostsPanel = ({ posts, channelName }: Props) => {
	return (
		<Card
			header={
				<Text as="h2" size="md" bold>
					# {channelName}
				</Text>
			}
			body={<PostList posts={posts} />}
		/>
	);
};
