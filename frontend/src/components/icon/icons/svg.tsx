import clsx from "clsx";
import type { IconProps } from "./icon-props";
import styles from "./icon.module.scss";

type Props = IconProps & {
	children: React.ReactNode;
};

export const Svg = ({ size, color, children }: Props) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			className={clsx(styles.svg, styles[color], styles[size])}
			viewBox="0 0 20 20"
		>
			<title>Icon</title>
			{children}
		</svg>
	);
};
