import clsx from "clsx";
import type { ComponentProps } from "react";
import { Icon } from "../icon";
import type { IconProps } from "../icon/icons/icon-props";
import styles from "./icon-button.module.scss";

type Props = IconProps & {
	icon: ComponentProps<typeof Icon>["icon"];
	onClick?: () => void;
};

export const IconButton = ({ icon, onClick, size, color }: Props) => {
	const handleClick = () => {
		onClick?.();
	};
	return (
		<button
			className={clsx(styles.root, styles.button)}
			type="button"
			onClick={handleClick}
		>
			<Icon icon={icon} size={size} color={color} />
		</button>
	);
};
