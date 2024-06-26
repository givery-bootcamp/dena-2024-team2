import clsx from "clsx";
import styles from "./icon-image.module.scss";

type Props = {
	src: string;
	size?: "sm" | "md" | "lg";
};

export const IconImage = ({ src, size = "md" }: Props) => {
	return (
		<div className={clsx(styles.root, styles[size])}>
			<img className={styles.icon} src={src} alt="icon" />
		</div>
	);
};
