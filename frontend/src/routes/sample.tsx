import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { getChannels } from "~/domains/channels/get-channels";

// sampleページなのでベタ書き
const SamplePage = () => {
	const query = useQuery({ queryKey: ["sample"], queryFn: getChannels });

	return (
		<div>
			<h1>Data Fetch Check</h1>
			<div>
				{query.data?.map((channel) => (
					<div key={channel.id}>{channel.name}</div>
				))}
			</div>
		</div>
	);
};
export const Route = createFileRoute("/sample")({
	component: () => <SamplePage />,
});
