import { SidePanel } from "./components/side-panel";
import { mock } from "./mock";

export const Posts = () => {
	return (
		<div>
			<h1>Posts</h1>
			<div>
				<SidePanel servers={mock.servers} channels={mock.channels} />
			</div>
		</div>
	);
};
