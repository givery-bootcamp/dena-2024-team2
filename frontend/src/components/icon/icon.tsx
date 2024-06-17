import type { IconProps } from "./icons/icon-props";
import { Plus } from "./icons/plus";

export const iconMap = (props: IconProps) => ({
	plus: <Plus {...props} />,
});

type Props = IconProps & {
	icon: keyof ReturnType<typeof iconMap>;
};

export const Icon = ({ icon, ...props }: Props) => {
	return iconMap(props)[icon];
};
