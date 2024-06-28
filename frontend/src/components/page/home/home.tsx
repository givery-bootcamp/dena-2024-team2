import { SidePanel } from "./components/side-panel";
import { mock } from "./mock";

export const Home = () => {
	return (
		<div>
			<h1>Home</h1>
			<div>
				<SidePanel servers={mock.servers} channels={mock.channels} />
			</div>
		</div>
	);
};
